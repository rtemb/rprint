package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/jung-kurt/gofpdf"
	"github.com/gorilla/mux"
	rc "gitlab.com/rtemb/receipt-print/receiptCustom"
	rs "gitlab.com/rtemb/receipt-print/receiptSchema"
)

// root derictory
func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Processing URL %s...", r.URL.Path)
}

// logger provides a log of requests
func logger(w http.ResponseWriter, r *http.Request) {
	remoteAddr := r.Header.Get("X-Forwarded-For")
	if remoteAddr == "" {
		remoteAddr = r.RemoteAddr
	}
	log.Infof("%s %s %s", remoteAddr, r.Method, r.URL.Path)
}

// GetAllReceipts returns all receipts
func GetAllReceipts(w http.ResponseWriter, r *http.Request) {
	receipts := Receipts{
		Receipt{Name: "item1", Price: 1.99, Bill: "18490000009984"},
		Receipt{Name: "item2", Price: 4.50, Bill: "18490000009985"},
	}
	responseWithJSON(w, http.StatusOK, receipts);
}

// CreateCustom prins custom receipt
// not implemented yet !!
func CreateCustom(w http.ResponseWriter, r *http.Request) {
	var Rc rc.PdfDocument
	data, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(data, &Rc)
	// todo add printing
	responseWithJSON(w, http.StatusOK, Rc)
}

//CreateReceipt print receipt and put it to filesystem
// in ./receipts folder
func CreateReceipt(w http.ResponseWriter, r *http.Request) {
	var Rs rs.ReceiptData
	data, err := ioutil.ReadAll(r.Body)
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
	response["link"] = r.Host + "/pdf/" + fileName
	responseWithJSON(w, http.StatusOK, response)
}

func giveFile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	file, _ := ioutil.ReadFile("receipts/" + vars["docName"]+ ".pdf")
	http.ServeContent(w, r, "myfile", time.Now(), bytes.NewReader(file))
}

func CreatePdfFile(receipt Receipt) string {
	currDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	filePath := currDir + "/receipts/"
	fileName := strconv.FormatInt(time.Now().UnixNano(), 10)
	ext := ".pdf"

	receiptId := strconv.Itoa(receipt.Id)
	itemName := string(receipt.Name)
	itemPrice := strconv.FormatFloat(float64(receipt.Price), 'f', 2, 64)
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
	pdf.CellFormat(0, 7, tr("Id чека: ")+receiptId, "", 1, "L", false, 0, "")
	pdf.CellFormat(0, 7, tr("Наименование товара: "+itemName), "", 1, "L", false, 0, "")
	pdf.CellFormat(0, 7, tr("Цена: ")+itemPrice, "", 1, "L", false, 0, "")
	pdf.CellFormat(0, 7, tr("Номер транзакции: ")+receiptBill, "", 1, "L", false, 0, "")
	pdf.ImageOptions(currDir+"/images/qrcode.png", 75, 0, 50, 50, true, gofpdf.ImageOptions{ImageType: "PNG", ReadDpi: false}, 0, "")
	err := pdf.OutputFileAndClose(filePath + fileName + ext)
	if err != nil {
		log.Println(err)
	}
	return fileName
}
