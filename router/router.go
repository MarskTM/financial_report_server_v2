package router

import (
	"net/http"
	"phenikaa/controller"
	"phenikaa/infrastructure"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/golang/glog"

	"github.com/go-chi/render"

	// _ "phenikaa/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

func Router() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.URLFormat)
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Compress(6, "application/json"))
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Use(middleware.Timeout(time.Duration(5 * time.Second)))
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*", "http://localhost:5173"}, // Use this to allow specific origin hosts
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},

		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)

	// Api swagger for developer mode
	r.Get("/api/v1/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(infrastructure.GetHTTPSwagger()),
		httpSwagger.DocExpansion("none"),
	))

	// Declare controller
	accessController := controller.NewAccessController()
	userController := controller.NewUserController()
	basicQueryController := controller.NewBasicQueryController()
	advanceFilterController := controller.NewAdvanceFilterController()
	// seedController := controller.NewSeedController()
	tidingController := controller.NewTidingController()

	documentController := controller.NewDocumentController()

	r.Route("/api/v1", func(router chi.Router) {
		// Ping the API
		router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("pong"))
		})

		// Public routes
		router.Post("/login", accessController.Login)
		router.Post("/logout", accessController.Logout)
		router.Post("/refresh", accessController.Refresh)
		router.Post("/users/register", userController.Register)
		router.Post("/users/forgot-password", userController.ForgotPassword)
		router.Post("/users/check-email-exact", userController.CheckEmailExact)

		// Private routes
		router.Group(func(protectRouter chi.Router) {
			// protectRouter.Use(jwtauth.Authenticator)
			// protectRouter.Use(internalMiddle.Authenticator)
			// protectRouter.Use(jwtauth.Verifier(infrastructure.GetEncodeAuth()))

			protectRouter.Route("/users", func(userRouter chi.Router) {
				userRouter.Put("/reset-password", userController.ResetPassword)
				userRouter.Put("/change-password", userController.ChangePassowrd)
				userRouter.Get("/all", userController.GetAllUsers)
				userRouter.Put("/update-state", userController.UpdateUserState)
				userRouter.Put("/update-role", userController.UpdateUserRole)
			})

			protectRouter.Route("/basic-query", func(accessRouter chi.Router) {
				accessRouter.Post("/", basicQueryController.Upsert)
				accessRouter.Delete("/", basicQueryController.Delete)
			})

			protectRouter.Route("/advance-filter", func(accessRouter chi.Router) {
				accessRouter.Post("/", advanceFilterController.Filter)
			})

			protectRouter.Route("/financial-report", func(financial chi.Router) {
				financial.Post("/upload", documentController.ImportReportData)
				financial.Post("/export", documentController.ExportReportData)
				financial.Delete("/delete", documentController.DeleteHistoryReport)
			})

			protectRouter.Route("/tiding", func(tidingRouter chi.Router) {
				tidingRouter.Get("/all", tidingController.GetAll)
                tidingRouter.Post("/create", tidingController.Create)
				tidingRouter.Post("/update", tidingController.Update)
				tidingRouter.Post("/delete", tidingController.Delete)
			})

			// Serve static files from the media directory
			fullpath := infrastructure.GetRootPath() + "/" + infrastructure.GetStoragePath()
			glog.V(1).Info("Full path: ", fullpath)
			fs := http.StripPrefix("/api/v1/media/", http.FileServer(http.Dir(fullpath)))
			protectRouter.Handle("/media/*", fs)
		})
	})

	return r
}
