package router

import "testing"

func TestClassifyRoute(t *testing.T) {
	cases := []struct {
		method string
		rest   []string
		want   routeShape
	}{
		{"GET", nil, shapeList},
		{"HEAD", []string{"12"}, shapeRead},
		{"HEAD", nil, shapeList},
		{"GET", []string{"count"}, shapeList},
		{"GET", []string{"12"}, shapeRead},
		{"GET", []string{"get", "companybuilding", "3"}, shapeParamGet},
		{"GET", []string{"search", "1"}, shapeCustom},
		{"GET", []string{"initdata"}, shapeCustom},
		{"POST", nil, shapeInsert},
		{"POST", []string{"batch"}, shapeInsertBatch},
		{"POST", []string{"make"}, shapeCustom},
		{"PUT", nil, shapeUpdate},
		{"PUT", []string{"remark"}, shapeFieldUpdate},
		{"DELETE", nil, shapeDelete},
		{"DELETE", []string{"batch"}, shapeDeleteBatch},
		{"DELETE", []string{"bybuilding"}, shapeDeleteBy},
	}

	for _, c := range cases {
		if got := classifyRoute(c.method, c.rest); got != c.want {
			t.Errorf("%v %v = %v, want %v", c.method, c.rest, got, c.want)
		}
	}
}

func TestToID(t *testing.T) {
	if toID(float64(3)) != 3 || toID("7") != 7 || toID(map[string]interface{}{"id": float64(9)}) != 9 || toID(nil) != 0 {
		t.Error("toID")
	}
}

func TestStripSecrets(t *testing.T) {
	v := map[string]interface{}{
		"item":  map[string]interface{}{"passwd": "x", "name": "a"},
		"items": []interface{}{map[string]interface{}{"extra": map[string]interface{}{"company": map[string]interface{}{"egiropasswd": "y"}}}},
	}
	stripSecrets(v)

	if v["item"].(map[string]interface{})["passwd"] != "" {
		t.Error("passwd")
	}
	inner := v["items"].([]interface{})[0].(map[string]interface{})["extra"].(map[string]interface{})["company"].(map[string]interface{})
	if inner["egiropasswd"] != "" {
		t.Error("egiropasswd")
	}
}

func TestRequiresTenantAdmin(t *testing.T) {
	if !requiresTenantAdmin("/api/external") || !requiresTenantAdmin("/api/user/upload/a.xlsx") || requiresTenantAdmin("/api/user") || requiresTenantAdmin("/api/externalx") {
		t.Error("requiresTenantAdmin")
	}
}

func TestNormalizeKeys(t *testing.T) {
	out, ok := normalizeKeys(map[string]interface{}{"Building": 1.0, "id": 2.0})
	if !ok || out["building"] != 1.0 || out["id"] != 2.0 {
		t.Errorf("normalize = %v", out)
	}
	if _, ok := normalizeKeys(map[string]interface{}{"building": 1.0, "Building": 2.0}); ok {
		t.Error("case collision must fail")
	}
	if _, ok := normalizeKeys(map[string]interface{}{"\u212Aey": 1.0}); ok {
		t.Error("non-ascii key must fail")
	}
}
