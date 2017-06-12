package main

import (
	"io/ioutil"
	"path/filepath"
	"encoding/json"
	"os"
	"time"
	"bytes"
	"strconv"	
	"fmt"
	"net/http"

	"github.com/takama/router"
	"github.com/jung-kurt/gofpdf"
	rn "gitlab.com/rtemb/receipt-print/receiptCustom"
	rs "gitlab.com/rtemb/receipt-print/receiptSchema"
)

// root derictory
func root(c *router.Control) {
	fmt.Fprintf(c.Writer, "Processing URL %s...", c.Request.URL.Path)
}

// logger provides a log of requests
func logger(c *router.Control) {
	remoteAddr := c.Request.Header.Get("X-Forwarded-For")
	if remoteAddr == "" {
		remoteAddr = c.Request.RemoteAddr
	}
	log.Infof("%s %s %s", remoteAddr, c.Request.Method, c.Request.URL.Path)
}

// Return all receipts
func GetAllReceipts(c *router.Control) {
	receipts := Receipts{
		Receipt{Name: "item1", Price: 1.99, Bill: "18490000009984"},
		Receipt{Name: "item2", Price: 4.50, Bill: "18490000009985"},
	}
	c.Code(http.StatusOK).Body(receipts)
}

func CreateReceiptN(c *router.Control) {
	var Rn rn.PdfDocument
	data, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(data, &Rn)
	c.Code(http.StatusOK).Body(Rn)
}

func CreateReceipt(c *router.Control) {
	var Rs rs.ReceiptData
	data, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(data, &Rs)

	currDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	filePath := currDir + "/receipts/" 
	fileName := strconv.FormatInt(time.Now().UnixNano(), 10)
	ext := ".pdf"

	receiptsDir := filepath.Join(currDir, "receipts")
	os.MkdirAll(receiptsDir, 0700)

	Rs.Print(filePath + fileName + ext)
	
	response := make(map[string]string)
	response["link"] =  c.Request.Host + "/pdf/" + fileName
	c.Code(http.StatusOK).Body(response)
}

func giveFile(c *router.Control) {
	file, _ := ioutil.ReadFile("receipts/" + c.Get(":docName") + ".pdf")
	http.ServeContent(c.Writer, c.Request, "myfile", time.Now(), bytes.NewReader(file))
}

func CreatePdfFile(receipt Receipt) string {
	currDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	filePath := currDir + "/receipts/" 
	fileName := strconv.FormatInt(time.Now().UnixNano(), 10)
	ext := ".pdf"

	receiptId	:= strconv.Itoa(receipt.Id)
	itemName	:= string(receipt.Name)
	itemPrice	:= strconv.FormatFloat(float64(receipt.Price), 'f', 2, 64)
	receiptBill := string(receipt.Bill)

	pdf := gofpdf.New("P", "mm", "A5", "")
	pdf.SetFontLocation(currDir + "/fonts")
    pdf.AddFont("Helvetica", "", "helvetica_1251.json")
    pdf.AddPage()
    pdf.SetFont("Helvetica", "", 16)
    tr := pdf.UnicodeTranslatorFromDescriptor("cp1251")
	pdf.CellFormat(67, 10, tr("Чек за покупку на сайте "), "", 0, "L", false, 0, "")
	pdf.SetTextColor(6, 69, 173)
	pdf.CellFormat(0, 10, "www.rtemb.com", "", 1, "L", false, 0, "http://www.rtemb.com")
	pdf.SetTextColor(0, 0, 0)
	pdf.CellFormat(0, 7, tr("Id чека: ") + receiptId, "", 1, "L", false, 0, "")
	pdf.CellFormat(0, 7, tr("Наименование товара: " + itemName), "", 1, "L", false, 0, "")
	pdf.CellFormat(0, 7, tr("Цена: ") + itemPrice, "", 1, "L", false, 0, "")
	pdf.CellFormat(0, 7, tr("Номер транзакции: ") + receiptBill, "", 1, "L", false, 0, "")
	pdf.ImageOptions(currDir + "/images/qrcode.png", 75, 0, 50, 50, true, gofpdf.ImageOptions{ImageType: "PNG", ReadDpi: false}, 0, "")
    err := pdf.OutputFileAndClose(filePath + fileName + ext)
    if err != nil {
        log.Println(err)
    }
	return fileName
}