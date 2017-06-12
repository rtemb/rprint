package main

import (
	"os"

	"github.com/takama/router"
	"github.com/Sirupsen/logrus"

) 

type Receipts []Receipt

var log = logrus.New()

func main() {
	port := os.Getenv("SERVICE_PORT")
	if len(port) == 0  {
		log.Fatal("Required parameter SERVICE_PORT is not set")
	}

	r := router.New()
	r.Logger = logger

	r.GET("/", root)
	r.GET("/test", GetAllReceipts)
	r.POST("/create", CreateReceipt)
	r.POST("/createN", CreateReceiptN)
	r.GET("/pdf/:docName", giveFile)
	
	log.Info("Service starting up... http://127.0.0.1:" , port)
	r.Listen(":" + port)

}

type Receipt struct {
	
    Id		int		`json:"Id"`
    Name	string	`json:"Name"`
    Price	float32	`json:"Price"`
    Bill	string	`json:"Bill"`
}