package main

import (
	"os"

	"github.com/Sirupsen/logrus"
	apiv1 "github.com/rtemb/rprint/apiv1"
	"github.com/takama/router"
)

var log = logrus.New()

const defaultAppPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		log.Warn("Required env variable PORT is not set! Used default PORT")
		port = defaultAppPort
	}

	r := router.New()
	r.Logger = logger
	r.GET("/", Root)
	r.POST("/v1/create", apiv1.CreateReceipt)
	r.POST("/v1/createcustom", apiv1.CreateCustom)
	r.GET("/v1/pdf/:docName", apiv1.GiveFile)

	// Readiness and liveness probes for Kubernetes
	r.GET("/info", Info)
	r.GET("/healthz", Healthz)

	log.Info("Service run at PORT: " + port)
	r.Listen(":" + port)
}
