package main

import (
	"os"

	"net/http"
	"encoding/json"
	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
)

type Receipts []Receipt

var log = logrus.New()

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		log.Fatal("Required env variable PORT is not set")
	}

	r := mux.NewRouter()
	r.HandleFunc("/", root)
	r.HandleFunc("/test", GetAllReceipts)
	r.HandleFunc("/create", CreateReceipt)
	r.HandleFunc("/createcustom", CreateCustom)
	r.HandleFunc("/pdf/{docName}", giveFile)

	log.Info("Service started up at port: " + port)
	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	log.Fatal(http.ListenAndServe("0.0.0.0:" + port, loggedRouter))
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
