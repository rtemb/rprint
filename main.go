package main

import (
	"os"

	"encoding/json"
	"github.com/Sirupsen/logrus"
	info "github.com/rtemb/rprint/info"
	"github.com/rtemb/rprint/version"
	"github.com/takama/router"
	"net/http"
)

type Receipts []Receipt

var log = logrus.New()

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		log.Fatal("Required env variable PORT is not set")
	}

	r := router.New()
	r.Logger = logger
	r.GET("/", root)
	r.GET("/test", GetAllReceipts)
	r.POST("/create", CreateReceipt)
	r.POST("/createcustom", CreateCustom)
	r.GET("/pdf/{docName}", giveFile)

	// Readiness and liveness probes for Kubernetes
	r.GET("/info", func(c *router.Control) {
		info.Info(c, version.RELEASE, version.REPO, version.COMMIT)
	})
	r.GET("/healthz", func(c *router.Control) {
		c.Code(http.StatusOK).Body(http.StatusText(http.StatusOK))
	})

	log.Info("Service started up at port: " + port)
	r.Listen(":" + port)
}

type Receipt struct {
	Id    int     `json:"Id"`
	Name  string  `json:"Name"`
	Price float32 `json:"Price"`
	Bill  string  `json:"Bill"`
}

func responseWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
