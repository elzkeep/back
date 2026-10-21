package global

// 비밀번호 저장 — bcrypt. 예전 계정은 평문으로 저장돼 있어 로그인할 때 평문 비교로 맞으면 해시로 바꿔 저장한다(router/auth.go).

import (
	"crypto/subtle"
	"strings"
	"sync"

	"golang.org/x/crypto/bcrypt"
)

func IsPasswordHash(stored string) bool {
	return strings.HasPrefix(stored, "$2a$") || strings.HasPrefix(stored, "$2b$") || strings.HasPrefix(stored, "$2y$")
}

// HashPassword 는 평문을 bcrypt 로 바꾼다. 이미 해시면 그대로
func HashPassword(passwd string) string {
	if passwd == "" || IsPasswordHash(passwd) {
		return passwd
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}

	return string(hash)
}

// CheckPassword 는 저장값과 입력을 비교한다. upgrade = 평문 저장값이라 해시로 바꿔야 함
func CheckPassword(stored string, input string) (ok bool, upgrade bool) {
	if stored == "" || input == "" {
		return false, false
	}

	if IsPasswordHash(stored) {
		return bcrypt.CompareHashAndPassword([]byte(stored), []byte(input)) == nil, false
	}

	if subtle.ConstantTimeCompare([]byte(stored), []byte(input)) == 1 {
		return true, true
	}

	return false, false
}

var defaultPasswordOnce sync.Once
var defaultPasswordHash string

// DefaultPasswordHash 는 엑셀 일괄등록으로 만든 계정의 초기 비밀번호("0000") 해시 (한 번만 계산)
func DefaultPasswordHash() string {
	defaultPasswordOnce.Do(func() {
		defaultPasswordHash = HashPassword("0000")
	})
	return defaultPasswordHash
}
