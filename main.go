package main

import (
	"fmt"
    "log"
    "net/http"
	"encoding/json"
	"time"
	"io/ioutil"
 	"strconv"
	"bytes"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
    "github.com/jung-kurt/gofpdf"
) 

type Receipts []Receipt

func main() {
	fmt.Println("service start at: http://localhost:8081")
	handleRequests()
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/all", GetAllReceipts)
	router.HandleFunc("/create", CreateReceipt)
	router.HandleFunc("/pdf/{docName}", giveFile)
	log.Fatal(http.ListenAndServe(":8081", router))
}


func giveFile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
    docName := vars["docName"]
	file, _ := ioutil.ReadFile("receipts/" + docName + ".pdf")
	http.ServeContent(w, r, "myfile", time.Now(), bytes.NewReader(file))
}

func CreateReceipt(w http.ResponseWriter, r *http.Request) {
	var receipt Receipt
	response := make(map[string]string)
	data, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(data, &receipt)

	currDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	receiptsDir := filepath.Join(currDir, "receipts")
	os.MkdirAll(receiptsDir, 0700)

	fileName := CreatePdfFile(receipt)
	response["link"] = "http://localhost:8081/pdf/" + fileName
	responseWithJSON(w, http.StatusOK, response);
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

	pdf := gofpdf.New("P", "mm", "A4", "")
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

func GetAllReceipts(w http.ResponseWriter, r *http.Request) {
	receipts := Receipts{
		Receipt{Name: "item1", Price: 1.99, Bill: "18490000009984"},
		Receipt{Name: "item2", Price: 4.50, Bill: "18490000009985"},
	}
	responseWithJSON(w, http.StatusOK, receipts)
}

func responseWithError(w http.ResponseWriter, code int, message string) {
    responseWithJSON(w, code, map[string]string{"error": message})
}

func responseWithJSON(w http.ResponseWriter, code int, payload interface{}) {
    response, _ := json.Marshal(payload)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    w.Write(response)
}

type Receipt struct {
    Id		int		`json:"Id"`
    Name	string	`json:"Name"`
    Price	float32	`json:"Price"`
    Bill	string	`json:"Bill"`
}