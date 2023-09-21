package global

import (
	"fmt"
	"image"
	"log"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/360EntSecGroup-Skylar/excelize"
	humanize "github.com/dustin/go-humanize"
	uuid "github.com/google/uuid"
)

type Excel struct {
	File  *excelize.File
	Width []int
	Align []string
	Cols  int
	Rows  int

	Pos      int
	Height   float64
	Filename string
	Path     string
}

func NewExcel(title string, header []string, width []int, align []string, headerFont float64, bodyFont float64) *Excel {
	var item Excel

	item.Width = width
	item.Align = align
	item.Cols = len(header)
	item.Pos = 0
	item.Rows = 0

	item.File = excelize.NewFile()

	for i, value := range header {
		t := fmt.Sprintf("%v", rune('A'+i))
		item.File.SetColWidth("Sheet1", t, t, float64(width[i])*0.8)
		item.HeaderCell(value)
	}

	return &item
}

func (p *Excel) SetHeight(height float64) {
	p.Height = height
}

func (p *Excel) Save(filename string) string {
	if filename == "" {
		p.Filename = GetTempFilename()
	} else {
		p.Path = fmt.Sprintf("webdata/temp/%v", uuid.New().String())
		os.Mkdir(p.Path, 0755)
		p.Filename = fmt.Sprintf("%v/%v", p.Path, filename)
	}

	p.File.SaveAs(p.Filename)

	return p.Filename
}

func (p *Excel) Remove() {
	if p.Path != "" {
		os.RemoveAll(p.Path)
	} else {
		os.Remove(p.Filename)
	}
}

func (p *Excel) HeaderCell(str string) {
	if p.Pos == 0 {
		p.File.SetRowHeight("Sheet1", p.Rows+1, 30)
	}

	style, _ := p.File.NewStyle(`{"alignment":{"horizontal":"center","vertical":"center"},"border":[{"type":"left","color":"000000","style":1},{"type":"top","color":"000000","style":1},{"type":"bottom","color":"000000","style":1},{"type":"right",   "color":"000000","style":1}],"fill":{"type":"pattern","pattern":1,"color":["#CCCCCC"]},"number_format":0,"lang":"ko-kr"}`)

	t := fmt.Sprintf("%v%v", rune('A'+p.Pos), p.Rows+1)
	p.File.SetCellValue("Sheet1", t, str)
	p.File.SetCellStyle("Sheet1", t, t, style)

	p.Pos++

	if p.Pos == p.Cols {
		p.Pos = 0
		p.Rows++
	}
}

func (p *Excel) Cell(str string) string {
	if p.Pos == 0 && p.Rows > 0 {
		p.File.SetRowHeight("Sheet1", p.Rows+1, p.Height)
	}

	align := "center"
	if p.Align[p.Pos] == "L" {
		align = "left"
	} else if p.Align[p.Pos] == "R" {
		align = "right"
	}

	style, _ := p.File.NewStyle(`{"alignment":{"horizontal":"` + align + `","vertical":"center"},"border":[{"type":"left","color":"000000","style":1},{"type":"top","color":"000000","style":1},{"type":"bottom","color":"000000","style":1},{"type":"right",   "color":"000000","style":1}],"fill":{"type":"pattern","pattern":1,"color":["#FFFFFF"]},"number_format":0,"lang":"ko-kr"}`)

	t := fmt.Sprintf("%v%v", rune('A'+p.Pos), p.Rows+1)
	p.File.SetCellValue("Sheet1", t, str)
	p.File.SetCellStyle("Sheet1", t, t, style)

	p.Pos++

	if p.Pos == p.Cols {
		p.Pos = 0
		p.Rows++
	}

	return t
}

func (p *Excel) CellInt(value int) {
	p.Cell(fmt.Sprintf("%v", value))
}

func (p *Excel) CellPrice(value int) {
	p.Cell(fmt.Sprintf("₩ %v", humanize.Comma(int64(value))))
}

func (p *Excel) CellImage(filename string) {
	t := p.Cell("")

	width := 100.0
	height := 100.0

	xScale := 0.2
	yScale := 0.2

	xOffset := 10
	yOffset := 10

	log.Println(xScale, yScale)

	if reader, err := os.Open(filename); err == nil {
		defer reader.Close()
		im, _, err := image.DecodeConfig(reader)
		if err != nil {
		}
		fmt.Printf("%d %d\n", im.Width, im.Height)

		xScale = width / float64(im.Width)
		yScale = height / float64(im.Height)
	} else {
		fmt.Println("Impossible to open the file:", err)
	}

	if err := p.File.AddPicture("Sheet1", t, filename, fmt.Sprintf(`{"x_scale": %v, "y_scale": %v, "x_offset":%v, "y_offset":%v}`, xScale, yScale, xOffset, yOffset)); err != nil {
		fmt.Println(err)
	}
}
