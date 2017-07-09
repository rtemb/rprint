package main

import (
	"os"

	"github.com/Sirupsen/logrus"
	apiv1 "github.com/rtemb/rprint/apiv1"
	"github.com/takama/router"
)

var log = logrus.New()

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		log.Fatal("Required env variable PORT is not set")
	}

	r := router.New()
	r.Logger = logger
	r.GET("/", apiv1.Root)
	r.POST("/v1/create", apiv1.CreateReceipt)
	r.POST("/v1/createcustom", apiv1.CreateCustom)
	r.GET("/v1/pdf/:docName", apiv1.GiveFile)

	// Readiness and liveness probes for Kubernetes
	r.GET("/info", apiv1.Info)
	r.GET("/healthz", apiv1.Healthz)

	log.Info("Service started up at port: " + port)
	r.Listen(":" + port)
}
