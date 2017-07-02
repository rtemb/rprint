package receiptCustom

import (
	"os"

	"github.com/Sirupsen/logrus"
	gofpdf "github.com/jung-kurt/gofpdf"
	"math"
	"strconv"
	"strings"
	"unicode/utf8"
)

// DefaultSchema Default receipt schema
const DefaultSchema = "default"

var log = logrus.New()

// ReceiptData - data of receipt
type ReceiptData struct {
	Schema   string `json:"schema"`
	ReceiptS *ReceiptS
}

// ReceiptS default receipt
type ReceiptS struct {
	MPlaceName    string `json:"MPlaceName"`
	MPlaceAddress string `json:"MPlaceAddress"`
	MPlaceINN     string `json:"MPlaceINN"`
	OperationType string `json:"OperationType"`
	Items         []*Item
	TaxPercent    string  `json:"TaxPercent"`
	Total         float64 `json:"Total"`
	FiscalNumber  string  `json:"FiscalNumber"`
	Date          string  `json:"Date"`
}

// Item - item that was sold
type Item struct {
	Name     string  `json:"Name"`
	Quantity float64 `json:"Quantity"`
	Price    float64 `json:"Price"`
}

// Print - function which execute printing process
func (rd *ReceiptData) Print(ffName string) {
	if rd.Schema != "" {
		printDefaultReceipt(rd, ffName)
	}
}

func printDefaultReceipt(pd *ReceiptData, ffName string) {
	currDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Can't get current app dirr: %s", err)
	}
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
	pdf.Ln(1)
	printLine("Purchased items:", pdf, 0, 7, 2, "L")
	pdf.Ln(1)
	delimeter("-", pdf)
	for idx, Item := range pd.ReceiptS.Items {
		printLine(strconv.Itoa(idx+1)+". ", pdf, 5, 7, 0, "L")
		printLine(Item.Name, pdf, 0, 7, 1, "L")
		printLine("Quantity: "+strconv.FormatFloat(Item.Quantity, 'f', 2, 64), pdf, 0, 7, 1, "L")
		printLine("Price: "+strconv.FormatFloat(Item.Price, 'f', 2, 64), pdf, 0, 7, 2, "L")
		delimeter("-", pdf)
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
	errSave := pdf.OutputFileAndClose(ffName)
	if errSave != nil {
		log.Fatalf("Can't write receipt to pdf file: %s", errSave)
	}
}

// printLine - Print line of text
// str - string to print
// pfs - *gofpdf.Fpdf package
// w - width of printing area
// h - height of printing aria
// nl - 1 - print new line after string, 0 - print next string on same line, 2 - below
// align - align of string: "C" center, "L" left, "R" rigth
func printLine(str string, pdf *gofpdf.Fpdf, w float64, h float64, nl int, align string) {
	tr := pdf.UnicodeTranslatorFromDescriptor("cp1251")
	charPerline := 46
	length := utf8.RuneCountInString(str)
	startPos := 0
	endPos := charPerline
	if length > charPerline {
		lines := int(math.Ceil(float64(length) / float64(charPerline)))
		// h = float64(lines) * h
		for idx := 0; idx < lines; idx++ {
			startPos = idx * charPerline
			endPos = (idx + 1) * charPerline
			if endPos >= length {
				endPos = endPos - (endPos - length)
			}
			pdf.CellFormat(w, h, tr(str[startPos:endPos]), "", 1, align, false, 0, "")
		}
	} else {
		pdf.CellFormat(w, h, tr(str), "", nl, align, false, 0, "")
	}
}

func delimeter(delimeter string, pdf *gofpdf.Fpdf) {
	str := strings.Repeat(delimeter, 70)
	pdf.CellFormat(0, 7, str, "", 1, "L", false, 0, "")
}
