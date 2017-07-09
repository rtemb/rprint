package apiv1

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	rc "github.com/rtemb/rprint/receiptCustom"
	rs "github.com/rtemb/rprint/receiptSchema"
	"github.com/takama/router"
)

// CreateCustom prins custom receipt
// not implemented yet !!
func CreateCustom(c *router.Control) {
	var Rc rc.Receipt
	data, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(data, &Rc)

	currDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Can't get current app dirr: %s", err)
	}

	filePath := currDir + "/receipts/"
	fileName := strconv.FormatInt(time.Now().UnixNano(), 10)
	ext := ".pdf"

	Rc.Print(filePath + fileName + ext)

	response := make(map[string]string)
	response["link"] = c.Request.Host + "/v1/pdf/" + fileName
	c.Code(http.StatusOK).Body(response)
}

//CreateReceipt print receipt and put it to filesystem
// in ./receipts folder
func CreateReceipt(c *router.Control) {
	var Rs rs.ReceiptData
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatalf("Error due reading request body: %s", err)
	}
	json.Unmarshal(data, &Rs)

	currDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Can't get current app dirr: %s", err)
	}

	filePath := currDir + "/receipts/"
	fileName := strconv.FormatInt(time.Now().UnixNano(), 10)
	ext := ".pdf"

	receiptsDir := currDir + "/receipts"
	os.MkdirAll(receiptsDir, 0700)

	Rs.Print(filePath + fileName + ext)

	response := make(map[string]string)
	response["link"] = c.Request.Host + "/v1/pdf/" + fileName
	c.Code(http.StatusOK).Body(response)
}

func GiveFile(c *router.Control) {
	file, _ := ioutil.ReadFile("receipts/" + c.Get(":docName") + ".pdf")
	http.ServeContent(c.Writer, c.Request, "myfile", time.Now(), bytes.NewReader(file))
}
