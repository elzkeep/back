package global

import "testing"

func TestPassword(t *testing.T) {
	hash := HashPassword("secret")
	if !IsPasswordHash(hash) || HashPassword(hash) != hash {
		t.Fatal("hash")
	}
	if ok, up := CheckPassword(hash, "secret"); !ok || up {
		t.Error("hash compare")
	}
	if ok, _ := CheckPassword(hash, "wrong"); ok {
		t.Error("wrong password accepted")
	}
	if ok, up := CheckPassword("plain", "plain"); !ok || !up {
		t.Error("legacy plain must pass and ask upgrade")
	}
	if ok, _ := CheckPassword("", ""); ok {
		t.Error("empty")
	}
}
