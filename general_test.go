package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Sirupsen/logrus"
	"github.com/rtemb/rprint/apiv1"
	"github.com/takama/router"
	"io/ioutil"
)

// buffer is a special variable to save log messages
var buffer bytes.Buffer

func init() {
	log.Out = &buffer
	log.Formatter = &logrus.JSONFormatter{}
}

// TestLogger checks if logger handler works correctly
func TestLogger(t *testing.T) {
	r := router.New()
	r.Logger = logger

	ts := httptest.NewServer(r)
	defer ts.Close()

	_, err := http.Get(ts.URL + "/")
	if err != nil {
		t.Fatal(err)
	}

	formated := struct {
		Level string `json:"level"`
		Msg   string `json:"msg"`
		Time  string `json:"time"`
	}{}
	err = json.NewDecoder(&buffer).Decode(&formated)
	if err != nil {
		t.Fatal(err)
	}

	msgParts := strings.Split(formated.Msg, " ")
	if len(msgParts) != 3 {
		t.Fatalf("Wrong message was logged: %s", formated.Msg)
	}
}

// TestHandler is the simplest test: check base (/) URL
func TestHandlerRoot(t *testing.T) {
	r := router.New()
	r.GET("/", apiv1.Root)

	ts := httptest.NewServer(r)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/")
	if err != nil {
		t.Fatal(err)
	}

	greeting, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		t.Fatal(err)
	}

	expectedGreeting := "Processing URL /..."
	testingGreeting := strings.Trim(string(greeting), " \n")
	if testingGreeting != expectedGreeting {
		t.Fatalf(
			"Wrong greeting '%s', expected '%s'",
			testingGreeting, expectedGreeting,
		)
	}
}

// TestHandler is the simplest test: check base (/) URL
func TestHandlerReceipt(t *testing.T) {
	r := router.New()
	r.POST("/v1/create", apiv1.CreateReceipt)

	ts := httptest.NewServer(r)
	defer ts.Close()

	var jsonStr = []byte(`{ "Schema": "default", "ReceiptS": {  "MPlaceName": "Exmaple header",  "MPlaceAddress": "www.example.com",  "MPlaceINN": "00000111111239990",  "OperatinType": "Sell",  "Items": [   {    "Name": "Raincoat",    "Quantity": 1.000,    "Price": 100.0   },    {    "Name": "Black Hat",    "Quantity": 1.000,    "Price": 33.0   },   {    "Name": "Gloves",    "Quantity": 1.000,    "Price": 15.0   }  ],  "TaxPercent": "18%",  "Total": 148.0,  "FiscalNumber": "000000000011198",  "Date": "2017-06-11 23:21:11" }}`)
	res, err := http.Post(ts.URL+"/v1/create", "text/plain", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Fatal(err)
	}
	Link := struct {
		Link string `json:"link"`
	}{}
	json.Unmarshal(body, &Link)

	testingLink := strings.Trim(string(Link.Link), " \n")
	if len(testingLink) == 0 {
		t.Fatalf("No link recived. Recived link: '%s'", testingLink)
	}
}
