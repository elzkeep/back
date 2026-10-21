package api

import (
	"bytes"
	"encoding/binary"
	"hash/crc32"
	"image"
	"image/png"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"zkeep/config"
	"zkeep/models"

	"github.com/signintech/gopdf"
)

func TestSelectLabel(t *testing.T) {
	// 단일 구분: 점검구분 제목 + 항목 제목
	if got := selectLabel("전압 및 전류 계측값", models.Item{Title: "Main 전압/전류", Value: 1}); got != "380v/220v" {
		t.Errorf("single = %v", got)
	}
	// multi 구분: 사용자가 바꾼 제목이어도 항목 제목으로 찾는다
	if got := selectLabel("1층 차단기", models.Item{Title: "차단기 종류", Value: 2}); got != "MCCB" {
		t.Errorf("multi = %v", got)
	}
	// 모르는 항목은 번호 그대로
	if got := selectLabel("x", models.Item{Title: "없는 항목", Value: 3}); got != "3" {
		t.Errorf("unknown = %v", got)
	}
}

func TestReasonActionText(t *testing.T) {
	if reasonText(models.Item{Reason: 7}) != "오염" || reasonText(models.Item{Reason: 1, Reasontext: "직접"}) != "직접" {
		t.Error("reason")
	}
	if actionText(models.Item{Action: 4}) != "수리 및 정비요함" || actionText(models.Item{Action: 99}) != "" {
		t.Error("action")
	}
}

func TestRasterizeSignature(t *testing.T) {
	svg := `<svg width="300" height="120" xmlns="http://www.w3.org/2000/svg"><path d="M 20 80 C 40 20, 60 20, 80 80" stroke="#000" stroke-width="4" fill="none"/></svg>`
	img, err := rasterizeSignature(svg, 300, 120)
	if err != nil || img.Bounds().Dx() == 0 {
		t.Fatalf("rasterize: %v", err)
	}
	if _, err := rasterizeSignature("", 10, 10); err == nil {
		t.Error("empty svg must fail")
	}
	if uploadPath("../etc/passwd") != "" || uploadPath("/etc/passwd") != "" {
		t.Error("uploadPath traversal")
	}
}

// 헤더만 거대한 PNG(압축 폭탄)는 풀지 않고 거부, 정상 사진은 넣는다
func TestReportImageLimits(t *testing.T) {
	dir := t.TempDir()
	old := config.UploadPath
	config.UploadPath = dir
	defer func() { config.UploadPath = old }()

	// 65535×65535 RGBA 라고 주장하는 PNG (IDAT 없음 — 크기 확인 단계에서 거부돼야 한다)
	var bomb bytes.Buffer
	bomb.Write([]byte("\x89PNG\r\n\x1a\n"))
	ihdr := make([]byte, 13)
	binary.BigEndian.PutUint32(ihdr[0:], 65535)
	binary.BigEndian.PutUint32(ihdr[4:], 65535)
	ihdr[8], ihdr[9] = 8, 6
	writeChunk := func(typ string, data []byte) {
		binary.Write(&bomb, binary.BigEndian, uint32(len(data)))
		bomb.WriteString(typ)
		bomb.Write(data)
		crc := crc32.NewIEEE()
		crc.Write([]byte(typ))
		crc.Write(data)
		binary.Write(&bomb, binary.BigEndian, crc.Sum32())
	}
	writeChunk("IHDR", ihdr)
	writeChunk("IEND", nil)
	os.WriteFile(filepath.Join(dir, "bomb.png"), bomb.Bytes(), 0644)

	img := image.NewRGBA(image.Rect(0, 0, 400, 300))
	var ok bytes.Buffer
	png.Encode(&ok, img)
	os.WriteFile(filepath.Join(dir, "ok.png"), ok.Bytes(), 0644)

	pdf := &gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()
	w := &reportWriter{pdf: pdf}

	if w.image("bomb.png", 0, 0, 100, 100) {
		t.Error("bomb png must be rejected")
	}
	if !w.image("ok.png", 0, 0, 100, 100) {
		t.Error("normal png must be embedded")
	}
	w.images = reportMaxImages
	if w.image("ok.png", 0, 0, 100, 100) {
		t.Error("image count limit")
	}
	if _, err := rasterizeSignature("<svg>"+strings.Repeat(" ", reportMaxSignature)+"</svg>", 10, 10); err == nil {
		t.Error("huge signature must be rejected")
	}
}
