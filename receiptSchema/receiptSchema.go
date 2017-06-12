package receiptCustom

import (
	// "errors"
	"path/filepath"
	"os"

	gofpdf "github.com/jung-kurt/gofpdf"
	"github.com/Sirupsen/logrus"
	"strconv"
)

const Default_Schema = "default"

var log = logrus.New()

type ReceiptData struct {
	Schema		string `json:"schema"`
	ReceiptS	*ReceiptS
}

type ReceiptS struct {
	MPlaceName		string `json:"MPlaceName"`
	MPlaceAddress	string `json:"MPlaceAddress"`
	MPlaceINN		string `json:"MPlaceINN"`
	OperationType	string `json:"OperationType"`
	Items			[]*Item
	TaxPercent		string `json:"TaxPercent"`
	Total			float64 `json:"Total"`
	FiscalNumber	string `json:"FiscalNumber"`
	Date			string `json:"Date"`
}

type Item struct {
	Name		string	`json:"Name"`
	Quantity	float64 `json:"Quantity"`
	Price		float64 `json:"Price"`
}

func (rd *ReceiptData) Print(ffName string) {
	if rd.Schema != "" {
		printDefaultReceipt(rd, ffName)
	}
}

func printDefaultReceipt(pd *ReceiptData, ffName string) {
	currDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	// rWidth := 100
	pdf := gofpdf.New("P", "mm", "A5", "")
	pdf.SetFontLocation(currDir + "/fonts")
    pdf.AddFont("Helvetica", "", "helvetica_1251.json")
    pdf.AddPage()
    pdf.SetFont("Helvetica", "", 16)
    // tr := pdf.UnicodeTranslatorFromDescriptor("cp1251")
	// pdf.CellFormat(125, 50, tr(pd.ReceiptS.MPlaceName), "", 1, "C", false, 0, "")
	pdf.Ln(4)
	printLine(pd.ReceiptS.MPlaceName, pdf, 0, 7, 2, "C")
	pdf.Ln(1)
	printLine(pd.ReceiptS.MPlaceAddress, pdf, 0, 7, 2, "C")
	pdf.Ln(3)
	printLine("TIN:", pdf, 0, 7, 0, "L")
	printLine(pd.ReceiptS.MPlaceINN, pdf, 0, 7, 2, "R")
	pdf.Ln(1)
	printLine("Opeartion Type:", pdf, 0, 7, 0, "L")
	printLine(pd.ReceiptS.OperationType, pdf, 0, 7, 2, "R")
	pdf.Ln(4)
	printLine("Purchased items:", pdf, 0, 7, 2, "L")
	pdf.Ln(1)
	for _, Item := range pd.ReceiptS.Items {
		printLine(Item.Name, pdf, 40, 7, 0, "L")
		printLine(strconv.FormatFloat(Item.Quantity, 'f', 2, 64), pdf, 5, 7, 0, "L")
		printLine(strconv.FormatFloat(Item.Price, 'f', 2, 64), pdf, 0, 7, 2, "R")
		pdf.Ln(1)
	}
	pdf.Ln(4)
	printLine("Tax Percent:", pdf, 0, 7, 0, "L")
	printLine(pd.ReceiptS.TaxPercent, pdf, 0, 7, 2, "R")
	pdf.Ln(1)
	printLine("TOTAL:", pdf, 0, 7, 0, "L")
	printLine(strconv.FormatFloat(pd.ReceiptS.Total, 'f', 2, 64), pdf, 0, 7, 2, "R")
	pdf.Ln(4)
	printLine("Fiscal number:", pdf, 0, 7, 0, "L")
	printLine(pd.ReceiptS.FiscalNumber, pdf, 0, 7, 2, "R")
	pdf.Ln(1)
	printLine("Purchase Date:", pdf, 0, 7, 0, "L")
	printLine(pd.ReceiptS.Date, pdf, 0, 7, 2, "R")
    err := pdf.OutputFileAndClose(ffName)
    if err != nil {
		log.Fatal("Can't write receipt to pdf file")
    }
}

func printLine(str string, pdf *gofpdf.Fpdf, y float64, x float64, nl int, align string) {
	tr := pdf.UnicodeTranslatorFromDescriptor("cp1251")
	if len(str) > 50 {
		pdf.CellFormat(y, x, tr(str[:50]), "", nl, align, false, 0, "")	
	}
	pdf.CellFormat(y, x, tr(str), "", nl, align, false, 0, "")	
}

