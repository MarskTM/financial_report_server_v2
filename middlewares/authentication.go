package middlewares

import (
	"fmt"
	"net/http"
	"phenikaa/controller"
	"phenikaa/infrastructure"
	"strings"

	"github.com/golang/glog"
)

func Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// cookie, err := r.Cookie("access_token")
		// if err != nil {
		// 	http.Error(w, "Access token cookie not found", http.StatusUnauthorized)
		// 	return
		// }
		// glog.V(3).Info("Cookie", cookie.Value)

		// Xác thực token
		// if err := verifyJWT(cookie.Value); err != nil {
		// 	http.Error(w, "Invalid token", http.StatusUnauthorized)
		// 	return
		// }
		// accessToken := cookie.Value

		authorization := r.Header.Get("Authorization")
		if authorization == "" {
			controller.BadRequestResponse(w, r, fmt.Errorf("authorization header not found"))
			return
		}
		authorizationBearer := strings.Split(authorization, " ")[1]
		accessToken := strings.Split(authorizationBearer, ";")[0]

		glog.V(3).Info("Cookie", accessToken)

		accessClaims, errDecodeToken := controller.GetAndDecodeToken(accessToken)
		if errDecodeToken != nil {
			controller.UnauthorizedResponse(w, r, errDecodeToken)
			return
		}

		// Check authentication
		accessUuid, ok := accessClaims["access_uuid"].(string)
		if !ok {
			controller.UnauthorizedResponse(w, r, fmt.Errorf("can't parse access uuid from token"))
			return
		}

		if index, isExist := infrastructure.FetchAuth(accessUuid); isExist != nil || index == 0 {
			controller.UnauthorizedResponse(w, r, fmt.Errorf("Unauthorized"))
			return
		}
		next.ServeHTTP(w, r)
	})
}
