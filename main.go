package main

import (
	"log"
	"net/http"
	"time"

	"phenikaa/infrastructure"
	"phenikaa/router"

	"github.com/golang/glog"
)

func main() {
	// go run cron job

	// go run main
	glog.V(1).Info("Environment: ", infrastructure.GetEnvironments())
	glog.V(1).Info("Server URL: ", infrastructure.GetHTTPURL())
	glog.V(1).Info("Database name: ", infrastructure.GetDBName())
	glog.V(1).Infof("Server running at port: %+v\n", infrastructure.GetAppPort())

	s := &http.Server{
		Addr:           ":" + infrastructure.GetAppPort(),
		Handler:        router.Router(),
		ReadTimeout:    6000 * time.Second,
		WriteTimeout:   6000 * time.Second,
		MaxHeaderBytes: 1 << 30,
	}
	log.Fatal(s.ListenAndServe())
	// glog.Flush()
}
