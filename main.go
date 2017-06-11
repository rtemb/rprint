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
	addr := os.Getenv("SERVICE_ADDR")
	if len(addr) == 0  {
		log.Fatal("Required parameter SERVICE_ADDR is not set")
	}

	r := router.New()
	r.Logger = logger

	r.GET("/", root)
	r.GET("/all", GetAllReceipts)
	r.POST("/create", CreateReceipt)
	r.GET("/pdf/:docName", giveFile)
	
	log.Info("Service starting up... ", addr, ":", port)
	r.Listen(addr + ":" + port)

}

type Receipt struct {
	
    Id		int		`json:"Id"`
    Name	string	`json:"Name"`
    Price	float32	`json:"Price"`
    Bill	string	`json:"Bill"`
}