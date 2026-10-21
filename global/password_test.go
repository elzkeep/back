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
	long := make([]byte, MaxPasswordInput+1)
	for i := range long {
		long[i] = 'a'
	}
	if ok, _ := CheckPassword(string(long), string(long)); ok {
		t.Error("too long input must be rejected")
	}
}
