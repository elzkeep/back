package api

import (
	"testing"
	"zkeep/models"
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
