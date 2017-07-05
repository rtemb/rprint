package receiptCustom

import (
	"github.com/Sirupsen/logrus"
	gofpdf "github.com/jung-kurt/gofpdf"
	"math"
	"os"
	"strconv"
	"unicode/utf8"
)

var log = logrus.New()

type Receipt struct {
	PageConfig   *PageConfig
	Instructions []Instruction
}

type PageConfig struct {
	Orientation string `json:"Orientation"`
	Format      string `json:"Format"`
	FontStyle   string `json:"FontStyle"`
}

type Instruction struct {
	Type       string `json:"Type"`
	Value      string `json:"Value"`
	LineConfig *LineConfig
}

type LineConfig struct {
	FontSize float64 `json:"FontSize"`
	Width    float64 `json:"Width"`
	Height   float64 `json:"Height"`
	NewLine  int     `json:"NewLine"`
	Align    string  `json:"Align"`
}

func (receipt *Receipt) Print(ffName string) {
	currDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Can't get current app dirr: %s", err)
	}
	pdf := gofpdf.New(receipt.PageConfig.Orientation, "mm", receipt.PageConfig.Format, "")
	pdf.SetFontLocation(currDir + "/fonts")
	pdf.AddFont("Helvetica", "", "helvetica_1251.json")
	pdf.AddPage()
	pdf.SetFont("Helvetica", receipt.PageConfig.FontStyle, 16)
	for _, instr := range receipt.Instructions {
		if instr.Type == "text" {
			printLine(instr.Value, pdf, instr.LineConfig.Width, instr.LineConfig.Height, instr.LineConfig.NewLine, instr.LineConfig.Align)
		} else if instr.Type == "nl" {
			lines, _ := strconv.ParseFloat(instr.Value, 64)
			pdf.Ln(lines)
		}
	}

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
