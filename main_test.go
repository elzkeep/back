package main_test

import (
	"testing"

	"zkeep/global/pdf"
)

func TestExcel(t *testing.T) {

	//xRatio := 595.28 / 210.0
	//yRatio := 841.89 / 297.0

	doc := pdf.NewPdf()
	doc.SetFontWeight(pdf.Medium)
	doc.SetFontSize(19)

	doc.AddPage()

	y := 16.0
	mx := 20.0

	width := 595.28 - mx*2

	grey50 := pdf.FromRGB(55, 55, 55)
	grey100 := pdf.FromRGB(114, 114, 114)
	grey180 := pdf.FromRGB(181, 181, 181)
	grey160 := pdf.FromRGB(161, 161, 161)
	grey230 := pdf.FromRGB(237, 237, 237)

	unit := 595.28 / 8
	doc.FillRect(mx, y, unit, 3, grey100)
	doc.FillRect(mx+unit, y, unit, 3, grey180)
	doc.FillRect(mx+unit*2, y, unit, 3, grey230)

	y += 3

	doc.FillRect(mx, y, width, 0.5, grey160)

	doc.SetStrokeColor(grey50)
	doc.SetFillColor(grey50)

	y += 30

	doc.SetTextColor(grey50)
	doc.TextOut(mx, y, width, 25, "전기설비 점검결과 기록표", pdf.Center|pdf.Top)

	fullFilename := "test.pdf"
	doc.Save(fullFilename)
}
