package router

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func newAuthTestApp() *fiber.App {
	app := fiber.New()
	app.Use(JwtAuthRequired)
	app.All("/*", func(c *fiber.Ctx) error {
		return c.SendString("ok")
	})

	return app
}

func TestJwtAuthRequired(t *testing.T) {
	app := newAuthTestApp()

	cases := []struct {
		method string
		path   string
		status int
	}{
		{"GET", "/management/billing", 200},
		{"GET", "/", 200},
		{"GET", "/api/billing", 401},
		{"POST", "/api/billinghistory/deposit", 401},
		{"GET", "/api/webnotice", 200},
		{"POST", "/api/user/findid", 200},
		{"POST", "/api/webjoin", 200},
		{"GET", "/api/webjoin", 401},
		{"PUT", "/api/webjoin", 401},
		{"DELETE", "/api/webjoin", 401},
		{"POST", "/api/webquestion", 200},
		{"GET", "/api/webquestion", 401},
		{"DELETE", "/api/webquestion", 401},
	}

	for _, v := range cases {
		res, err := app.Test(httptest.NewRequest(v.method, v.path, nil))
		if err != nil {
			t.Fatal(err)
		}

		if res.StatusCode != v.status {
			t.Errorf("%v %v: status %v, want %v", v.method, v.path, res.StatusCode, v.status)
		}
	}
}

func TestJwtAuthRequiredWithToken(t *testing.T) {
	app := newAuthTestApp()

	req := httptest.NewRequest("GET", "/api/billing", nil)
	req.Header.Set("Authorization", "Bearer invalid")
	res, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != 401 {
		t.Errorf("invalid token: status %v, want 401", res.StatusCode)
	}
}
