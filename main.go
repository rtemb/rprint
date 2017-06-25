package main

import (
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/takama/router"
)

type Receipts []Receipt

var log = logrus.New()

func main() {
	port := os.Getenv("SERVICE_PORT")
	appUrl := os.Getenv("HEROKU_APP_URL")
	if len(appUrl) == 0 {
		appUrl = "127.0.0.1"
	}
	
	if len(port) == 0 {
		log.Fatal("Required parameter SERVICE_PORT is not set")
	}

	r := router.New()
	r.Logger = logger

	r.GET("/", root)
	r.GET("/test", GetAllReceipts)
	r.POST("/create", CreateReceipt)
	r.POST("/createcustom", CreateCustom)
	r.GET("/pdf/:docName", giveFile)

	log.Info("Service starting up... ", appUrl + ":" + port)
	r.Listen(appUrl + ":" + port)

}

type Receipt struct {
	Id    int     `json:"Id"`
	Name  string  `json:"Name"`
	Price float32 `json:"Price"`
	Bill  string  `json:"Bill"`
}
