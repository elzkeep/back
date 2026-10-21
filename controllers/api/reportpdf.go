package api

// 점검 보고서 PDF — 「전기설비 점검결과 기록표」(doc/report.pdf 시안)를 점검 데이터로 만든다.
//   머리글(점검자·점검지·고객·점검일) → 1. 정격용량 → 2. 점검 결과(설비 구분별 항목·판정)
//   → 3. 부적합·요주의 항목(사유·조치·사진) → 4. 종합의견 → 점검 업체·서명(점검자/관리담당자)
// 선택형 항목은 번호만 저장되므로 앱 정의에서 뽑은 reportSelectLabels 로 글자를 찾는다.
// 서명은 앱(hand_signature)이 SVG 로 저장한다 → oksvg 로 래스터화해 넣는다.

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"zkeep/config"
	"zkeep/models"
	itemmodel "zkeep/models/item"

	"github.com/signintech/gopdf"
	"github.com/srwiley/oksvg"
	"github.com/srwiley/rasterx"
)

var reportTopcategoryNames = map[int]string{
	1: "저압설비", 2: "고압설비", 3: "변전설비", 4: "부하설비", 5: "발전설비", 6: "기타안전설비",
	7: "태양광 발전", 8: "전기차 충전기", 9: "ESS", 10: "UPS", 11: "연료전지", 12: "풍력발전",
	13: "수력발전", 14: "적외선 열화상",
}

// 앱 components/status.dart 의 사유·조치 선택지 (번호로 저장)
var reportReasonLabels = []string{"없음", "직접입력", "누전", "단락", "단선", "작동불량", "파손", "오염", "노후", "누유", "누수", "균열", "발열", "이상음"}
var reportActionLabels = []string{"없음", "직접입력", "점검중", "조치완료", "수리 및 정비요함", "교체요함", "설치요함"}

const (
	reportFontRegular = "regular"
	reportFontBold    = "bold"

	pageWidth    = 595.28
	pageHeight   = 841.89
	marginX      = 40.0
	marginTop    = 40.0
	marginBottom = 50.0
	contentWidth = pageWidth - marginX*2
)

type reportData struct {
	Report      models.Report
	Building    *models.Building
	Customer    *models.Company // 건물주
	Company     *models.Company // 점검 업체
	Inspector   *models.User
	Manager     string // 관리담당자 (계약의 담당자명)
	Datas       []models.Data
	Items       map[int64][]models.Item // data id → items
	Topcategory []int
}

func loadReportData(conn interface{}, report *models.Report) *reportData {
	d := &reportData{Report: *report, Items: map[int64][]models.Item{}}

	d.Building = models.NewBuildingManager(conn).Get(report.Building)
	if d.Building != nil {
		d.Customer = models.NewCompanyManager(conn).Get(d.Building.Company)
	}
	d.Company = models.NewCompanyManager(conn).Get(report.Company)
	d.Inspector = models.NewUserManager(conn).Get(report.User)

	if customer := models.NewCustomerManager(conn).GetByCompanyBuilding(report.Company, report.Building); customer != nil {
		d.Manager = customer.Managername
	}

	d.Datas = models.NewDataManager(conn).Find([]interface{}{
		models.Where{Column: "report", Value: report.Id, Compare: "="},
		models.Ordering("d_topcategory, d_order, d_id"),
	})

	items := models.NewItemManager(conn).Find([]interface{}{
		models.Where{Column: "report", Value: report.Id, Compare: "="},
		models.Ordering("i_data, i_order, i_id"),
	})
	for _, v := range items {
		d.Items[v.Data] = append(d.Items[v.Data], v)
	}

	seen := map[int]bool{}
	for _, v := range d.Datas {
		if !seen[v.Topcategory] {
			seen[v.Topcategory] = true
			d.Topcategory = append(d.Topcategory, v.Topcategory)
		}
	}
	sort.Ints(d.Topcategory)

	return d
}

func selectLabel(dataTitle string, item models.Item) string {
	if labels, ok := reportSelectLabels[dataTitle+"|"+item.Title]; ok {
		if v, ok := labels[item.Value]; ok {
			return v
		}
	}

	// multi 구분은 사용자가 제목을 바꾸므로 항목 제목만으로도 찾는다 (유일할 때)
	var found map[int]string
	for key, labels := range reportSelectLabels {
		if strings.HasSuffix(key, "|"+item.Title) {
			if found != nil {
				found = nil
				break
			}
			found = labels
		}
	}
	if found != nil {
		if v, ok := found[item.Value]; ok {
			return v
		}
	}

	return fmt.Sprintf("%v", item.Value)
}

func labelAt(labels []string, i int) string {
	if i <= 0 || i >= len(labels) {
		return ""
	}
	return labels[i]
}

func reasonText(item models.Item) string {
	if item.Reason == 1 || item.Reasontext != "" && item.Reason == 0 {
		return item.Reasontext
	}
	return labelAt(reportReasonLabels, item.Reason)
}

func actionText(item models.Item) string {
	if item.Action == 1 || item.Actiontext != "" && item.Action == 0 {
		return item.Actiontext
	}
	return labelAt(reportActionLabels, item.Action)
}

type reportWriter struct {
	pdf *gopdf.GoPdf
	y   float64
}

func (w *reportWriter) font(bold bool, size float64) {
	if bold {
		w.pdf.SetFont(reportFontBold, "", size)
	} else {
		w.pdf.SetFont(reportFontRegular, "", size)
	}
}

func (w *reportWriter) newPage() {
	w.pdf.AddPage()
	// 시안의 상단 띠
	w.pdf.SetFillColor(110, 110, 110)
	w.pdf.RectFromUpperLeftWithStyle(marginX-12, 22, 100, 3, "F")
	w.pdf.SetFillColor(180, 180, 180)
	w.pdf.RectFromUpperLeftWithStyle(marginX+88, 22, 100, 3, "F")
	w.pdf.SetStrokeColor(160, 160, 160)
	w.pdf.SetLineWidth(0.5)
	w.pdf.Line(marginX-12, 25, pageWidth-marginX+12, 25)
	w.y = marginTop
}

// ensure 는 h 만큼 자리가 없으면 새 쪽으로 넘긴다
func (w *reportWriter) ensure(h float64) {
	if w.y+h > pageHeight-marginBottom {
		w.newPage()
	}
}

func (w *reportWriter) text(x, y, width, height float64, s string, align int) {
	w.pdf.SetXY(x, y)
	w.pdf.CellWithOption(&gopdf.Rect{W: width, H: height}, s, gopdf.CellOption{Align: align})
}

// lines 는 width 안에 들어가게 줄을 나눈다
func (w *reportWriter) lines(s string, width float64) []string {
	result := []string{}
	for _, para := range strings.Split(strings.ReplaceAll(s, "\r", ""), "\n") {
		if para == "" {
			result = append(result, "")
			continue
		}
		parts, err := w.pdf.SplitText(para, width)
		if err != nil || len(parts) == 0 {
			result = append(result, para)
			continue
		}
		result = append(result, parts...)
	}
	return result
}

func (w *reportWriter) section(title string) {
	w.ensure(40)
	w.y += 8
	w.font(true, 12)
	w.pdf.SetTextColor(30, 30, 30)
	w.text(marginX, w.y, contentWidth, 18, title, gopdf.Left|gopdf.Middle)
	w.y += 20
	w.pdf.SetFillColor(110, 110, 110)
	w.pdf.RectFromUpperLeftWithStyle(marginX, w.y, 100, 3, "F")
	w.pdf.SetStrokeColor(160, 160, 160)
	w.pdf.SetLineWidth(0.5)
	w.pdf.Line(marginX, w.y+3, pageWidth-marginX, w.y+3)
	w.y += 10
}

func (w *reportWriter) subsection(title string) {
	w.ensure(30)
	w.y += 4
	w.font(true, 10.5)
	w.pdf.SetTextColor(30, 30, 30)
	w.text(marginX, w.y, contentWidth, 16, title, gopdf.Left|gopdf.Middle)
	w.y += 17
	w.pdf.SetStrokeColor(200, 200, 200)
	w.pdf.SetLineWidth(0.5)
	w.pdf.Line(marginX, w.y, pageWidth-marginX, w.y)
	w.y += 4
}

func fileExists(p string) bool {
	info, err := os.Stat(p)
	return err == nil && !info.IsDir()
}

// uploadPath 는 업로드 파일명(하위 폴더 포함 가능)을 실제 경로로 바꾼다. 경로 이탈은 막는다
func uploadPath(name string) string {
	name = strings.TrimSpace(name)
	if name == "" || strings.Contains(name, "..") || path.IsAbs(name) {
		return ""
	}
	return filepath.Join(config.UploadPath, filepath.FromSlash(name))
}

// image 는 사진을 box 안에 비율을 지켜 넣는다. 실패하면 false
func (w *reportWriter) image(file string, x, y, boxW, boxH float64) bool {
	p := uploadPath(file)
	if p == "" || !fileExists(p) {
		return false
	}

	f, err := os.Open(p)
	if err != nil {
		return false
	}
	cfg, _, err := image.DecodeConfig(f)
	f.Close()
	if err != nil || cfg.Width == 0 || cfg.Height == 0 {
		return false
	}

	scale := boxW / float64(cfg.Width)
	if s := boxH / float64(cfg.Height); s < scale {
		scale = s
	}
	iw, ih := float64(cfg.Width)*scale, float64(cfg.Height)*scale

	return w.pdf.Image(p, x+(boxW-iw)/2, y+(boxH-ih)/2, &gopdf.Rect{W: iw, H: ih}) == nil
}

// signature 는 SVG 서명을 box 안에 그린다
func (w *reportWriter) signature(svg string, x, y, boxW, boxH float64) bool {
	img, err := rasterizeSignature(svg, int(boxW*3), int(boxH*3))
	if err != nil {
		return false
	}
	return w.pdf.ImageFrom(img, x, y, &gopdf.Rect{W: boxW, H: boxH}) == nil
}

func rasterizeSignature(svg string, width, height int) (image.Image, error) {
	if !strings.Contains(svg, "<svg") {
		return nil, errors.New("not svg")
	}

	icon, err := oksvg.ReadIconStream(bytes.NewBufferString(svg), oksvg.IgnoreErrorMode)
	if err != nil {
		return nil, err
	}

	vw, vh := icon.ViewBox.W, icon.ViewBox.H
	if vw <= 0 || vh <= 0 {
		return nil, errors.New("empty viewbox")
	}

	// 비율 유지
	scale := float64(width) / vw
	if s := float64(height) / vh; s < scale {
		scale = s
	}
	tw, th := int(vw*scale), int(vh*scale)
	if tw <= 0 || th <= 0 {
		return nil, errors.New("empty size")
	}
	icon.SetTarget(0, 0, float64(tw), float64(th))

	img := image.NewRGBA(image.Rect(0, 0, tw, th))
	draw.Draw(img, img.Bounds(), &image.Uniform{C: color.White}, image.Point{}, draw.Src)
	scanner := rasterx.NewScannerGV(tw, th, img, img.Bounds())
	icon.Draw(rasterx.NewDasher(tw, th, scanner), 1)

	return img, nil
}

func itemLine(dataTitle string, item models.Item) (label string, value string) {
	switch item.Type {
	case itemmodel.TypeText:
		v := strings.TrimSpace(item.Content)
		if v != "" && item.Unit != "" {
			v += " " + item.Unit
		}
		return item.Title, v
	case itemmodel.TypeSelect:
		return item.Title, selectLabel(dataTitle, item)
	case itemmodel.TypeStatus:
		label = item.Title
		if label == "" {
			label = "상태"
		}
		return label, itemmodel.GetStatus(item.Status)
	}

	return item.Title, strings.TrimSpace(item.Content)
}

func isDefect(item models.Item) bool {
	return item.Type == itemmodel.TypeStatus && (item.Status == itemmodel.StatusWarning || item.Status == itemmodel.StatusDanger)
}

// BuildReportPdf 는 점검 보고서 PDF 바이트를 만든다
func BuildReportPdf(d *reportData) ([]byte, error) {
	pdf := &gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})

	if err := pdf.AddTTFFont(reportFontRegular, "./fonts/NotoSansKR-Regular.ttf"); err != nil {
		return nil, fmt.Errorf("font: %v", err)
	}
	if err := pdf.AddTTFFont(reportFontBold, "./fonts/NotoSansKR-Bold.ttf"); err != nil {
		return nil, fmt.Errorf("font: %v", err)
	}

	w := &reportWriter{pdf: pdf}
	w.newPage()

	// 제목
	w.font(true, 20)
	pdf.SetTextColor(20, 20, 20)
	w.text(marginX, w.y+6, contentWidth, 30, "전기설비 점검결과 기록표", gopdf.Center|gopdf.Middle)
	w.y += 52

	// 머리글
	inspector, place, customer := "", "", ""
	if d.Inspector != nil {
		inspector = d.Inspector.Name
	}
	if d.Building != nil {
		place = d.Building.Name
	}
	if d.Customer != nil {
		customer = d.Customer.Name
	}
	checkdate := d.Report.Checkdate
	if len(checkdate) >= 10 {
		checkdate = checkdate[:10]
	}
	heads := [][2]string{{"점검자", inspector}, {"점검지", place}, {"고객명", customer}, {"점검일", checkdate}}
	colW := contentWidth / 4
	for i, h := range heads {
		x := marginX + colW*float64(i)
		w.font(true, 9.5)
		pdf.SetTextColor(60, 60, 60)
		w.text(x, w.y, 40, 16, h[0]+" :", gopdf.Left|gopdf.Middle)
		w.font(false, 9.5)
		pdf.SetTextColor(20, 20, 20)
		w.text(x+40, w.y, colW-44, 16, h[1], gopdf.Left|gopdf.Middle)
	}
	w.y += 20
	if d.Report.Title != "" {
		w.font(false, 9)
		pdf.SetTextColor(100, 100, 100)
		w.text(marginX, w.y, contentWidth, 14, d.Report.Title, gopdf.Left|gopdf.Middle)
		w.y += 14
	}

	// 1. 정격용량
	w.section("1. 정격용량")
	if d.Building != nil {
		caps := [][2]string{
			{"계약전력", fmt.Sprintf("%v kW", d.Building.Contractvolumn)},
			{"수전설비", fmt.Sprintf("%v kW", d.Building.Receivevolumn)},
			{"발전설비", fmt.Sprintf("%v kW", d.Building.Generatevolumn)},
			{"태양광", fmt.Sprintf("%v kW", d.Building.Sunlightvolumn)},
		}
		for i, c := range caps {
			x := marginX + colW*float64(i)
			w.font(false, 9)
			pdf.SetTextColor(90, 90, 90)
			w.text(x, w.y, 50, 16, c[0], gopdf.Left|gopdf.Middle)
			pdf.SetTextColor(20, 20, 20)
			w.text(x+50, w.y, colW-54, 16, c[1], gopdf.Left|gopdf.Middle)
		}
		w.y += 20
	}

	// 2. 점검 결과
	w.section("2. 점검 결과")
	defects := []struct {
		category string
		data     string
		item     models.Item
	}{}

	if len(d.Datas) == 0 {
		w.font(false, 9.5)
		pdf.SetTextColor(120, 120, 120)
		w.text(marginX, w.y, contentWidth, 16, "입력된 점검 항목이 없습니다.", gopdf.Left|gopdf.Middle)
		w.y += 20
	}

	labelW := 130.0
	valueX := marginX + labelW + 10
	valueW := contentWidth - labelW - 10

	for _, top := range d.Topcategory {
		category := reportTopcategoryNames[top]
		if category == "" {
			category = fmt.Sprintf("설비 %v", top)
		}
		w.subsection(category)

		for _, data := range d.Datas {
			if data.Topcategory != top {
				continue
			}
			items := d.Items[data.Id]

			rows := [][2]string{}
			for _, item := range items {
				label, value := itemLine(data.Title, item)
				if label == "" && value == "" {
					continue
				}
				rows = append(rows, [2]string{label, value})
				if isDefect(item) {
					defects = append(defects, struct {
						category string
						data     string
						item     models.Item
					}{category, data.Title, item})
				}
			}
			if len(rows) == 0 {
				rows = append(rows, [2]string{"", "-"})
			}

			w.font(false, 9)
			titleLines := w.lines(data.Title, labelW)
			height := float64(len(rows)) * 15
			if h := float64(len(titleLines)) * 13; h > height {
				height = h
			}
			w.ensure(height + 8)

			w.font(true, 9.5)
			pdf.SetTextColor(30, 30, 30)
			for i, l := range titleLines {
				w.text(marginX, w.y+float64(i)*13, labelW, 13, l, gopdf.Left|gopdf.Middle)
			}

			for i, r := range rows {
				y := w.y + float64(i)*15
				w.font(false, 9)
				pdf.SetTextColor(70, 70, 70)
				w.text(valueX, y, valueW-90, 15, r[0], gopdf.Left|gopdf.Middle)
				if r[1] == "부적합" || r[1] == "요주의" {
					w.font(true, 9)
					pdf.SetTextColor(200, 40, 40)
				} else {
					pdf.SetTextColor(20, 20, 20)
				}
				w.text(valueX+valueW-90, y, 90, 15, r[1], gopdf.Right|gopdf.Middle)
			}

			w.y += height + 4
			pdf.SetStrokeColor(225, 225, 225)
			pdf.SetLineWidth(0.4)
			pdf.Line(valueX, w.y, pageWidth-marginX, w.y)
			w.y += 4
		}
	}

	// 3. 부적합·요주의 항목
	w.section("3. 부적합·요주의 항목")
	if len(defects) == 0 {
		w.font(false, 9.5)
		pdf.SetTextColor(120, 120, 120)
		w.text(marginX, w.y, contentWidth, 16, "해당 없음", gopdf.Left|gopdf.Middle)
		w.y += 20
	} else {
		cols := []float64{80, 105, 50, 85, 95, contentWidth - 415}
		heads := []string{"구분", "항목", "판정", "사유", "조치사항", "사진"}
		drawHead := func() {
			w.font(true, 9)
			pdf.SetTextColor(40, 40, 40)
			x := marginX
			for i, h := range heads {
				w.text(x, w.y, cols[i], 16, h, gopdf.Center|gopdf.Middle)
				x += cols[i]
			}
			w.y += 18
			pdf.SetStrokeColor(120, 120, 120)
			pdf.SetLineWidth(0.6)
			pdf.Line(marginX, w.y, pageWidth-marginX, w.y)
		}
		w.ensure(40)
		drawHead()

		for _, def := range defects {
			rowH := 70.0
			if def.item.Image == "" {
				rowH = 34
			}
			if w.y+rowH > pageHeight-marginBottom {
				w.newPage()
				drawHead()
			}

			label := def.item.Title
			if label == "" {
				label = def.data
			} else {
				label = def.data + " · " + label
			}
			cells := []string{def.category, label, itemmodel.GetStatus(def.item.Status), reasonText(def.item), actionText(def.item)}

			w.font(false, 8.5)
			x := marginX
			for i, cell := range cells {
				if i == 2 {
					pdf.SetTextColor(200, 40, 40)
				} else {
					pdf.SetTextColor(30, 30, 30)
				}
				ls := w.lines(cell, cols[i]-6)
				top := w.y + (rowH-float64(len(ls))*12)/2
				for j, l := range ls {
					w.text(x+3, top+float64(j)*12, cols[i]-6, 12, l, gopdf.Center|gopdf.Middle)
				}
				x += cols[i]
			}

			if def.item.Image != "" {
				if !w.image(def.item.Image, x+4, w.y+4, cols[5]-8, rowH-8) {
					pdf.SetTextColor(150, 150, 150)
					w.text(x, w.y, cols[5], rowH, "사진 없음", gopdf.Center|gopdf.Middle)
				}
			} else {
				pdf.SetTextColor(150, 150, 150)
				w.text(x, w.y, cols[5], rowH, "-", gopdf.Center|gopdf.Middle)
			}

			w.y += rowH
			pdf.SetStrokeColor(210, 210, 210)
			pdf.SetLineWidth(0.4)
			pdf.Line(marginX, w.y, pageWidth-marginX, w.y)
		}
		w.y += 6
	}

	// 4. 종합의견 (+ 대표 사진)
	w.section("4. 종합의견")
	w.font(false, 9.5)
	pdf.SetTextColor(30, 30, 30)
	content := strings.TrimSpace(d.Report.Content)
	if content == "" {
		content = "-"
	}
	for _, l := range w.lines(content, contentWidth) {
		w.ensure(14)
		w.text(marginX, w.y, contentWidth, 14, l, gopdf.Left|gopdf.Middle)
		w.y += 14
	}
	if d.Report.Image != "" {
		w.ensure(170)
		if w.image(d.Report.Image, marginX, w.y+6, 220, 160) {
			w.y += 172
		}
	}

	// 업체·서명
	if w.y+10 > pageHeight-marginBottom-80 {
		w.newPage()
	}
	w.y = pageHeight - marginBottom - 80
	company, address := "", ""
	if d.Company != nil {
		company = d.Company.Name
		address = strings.TrimSpace(d.Company.Address + " " + d.Company.Addressetc)
	}
	w.font(true, 18)
	pdf.SetTextColor(20, 20, 20)
	w.text(marginX, w.y+8, 260, 28, company, gopdf.Left|gopdf.Middle)
	w.font(false, 8.5)
	pdf.SetTextColor(80, 80, 80)
	w.text(marginX, w.y+42, 260, 14, address, gopdf.Left|gopdf.Middle)

	manager := d.Manager
	signs := []struct {
		label, name, svg string
	}{{"점검자", inspector, d.Report.Sign1}, {"관리담당자", manager, d.Report.Sign2}}
	sx := marginX + 290
	for i, s := range signs {
		y := w.y + float64(i)*34
		w.font(true, 9.5)
		pdf.SetTextColor(30, 30, 30)
		w.text(sx, y+6, 60, 22, s.label, gopdf.Left|gopdf.Middle)
		w.font(false, 9.5)
		w.text(sx+60, y+6, 70, 22, s.name, gopdf.Left|gopdf.Middle)
		if !w.signature(s.svg, sx+130, y, 85, 32) {
			pdf.SetTextColor(150, 150, 150)
			w.text(sx+130, y+6, 85, 22, "(서명)", gopdf.Right|gopdf.Middle)
		}
	}
	pdf.SetFillColor(110, 110, 110)
	pdf.RectFromUpperLeftWithStyle(pageWidth-marginX-90, pageHeight-marginBottom+14, 102, 3, "F")
	pdf.SetStrokeColor(160, 160, 160)
	pdf.SetLineWidth(0.5)
	pdf.Line(marginX-12, pageHeight-marginBottom+17, pageWidth-marginX+12, pageHeight-marginBottom+17)

	return pdf.GetBytesPdf(), nil
}
