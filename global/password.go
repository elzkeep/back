package global

// 비밀번호 저장 — bcrypt. 예전 계정은 평문으로 저장돼 있어 로그인할 때 평문 비교로 맞으면 해시로 바꿔 저장한다(router/auth.go).

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/hex"
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

// 입력 비밀번호 길이 상한 — 이보다 긴 입력은 비교 없이 거부 (bcrypt 전처리 비용)
const MaxPasswordInput = 256

var dummyHashOnce sync.Once
var dummyHash []byte

// BurnPasswordCheck 는 결과 없이 bcrypt 비교 한 번만큼 시간을 쓴다.
// 없는 아이디·평문 계정도 해시 계정과 같은 시간이 걸리게 해 응답 시간으로 계정 존재를 알 수 없게 한다
func BurnPasswordCheck(input string) {
	dummyHashOnce.Do(func() {
		dummyHash, _ = bcrypt.GenerateFromPassword([]byte("timing-equalizer"), bcrypt.DefaultCost)
	})
	if len(input) > MaxPasswordInput {
		input = input[:MaxPasswordInput]
	}
	bcrypt.CompareHashAndPassword(dummyHash, []byte(input))
}

// CheckPassword 는 저장값과 입력을 비교한다. upgrade = 평문 저장값이라 해시로 바꿔야 함
func CheckPassword(stored string, input string) (ok bool, upgrade bool) {
	if stored == "" || input == "" || len(input) > MaxPasswordInput {
		BurnPasswordCheck("")
		return false, false
	}

	if IsPasswordHash(stored) {
		return bcrypt.CompareHashAndPassword([]byte(stored), []byte(input)) == nil, false
	}

	// 예전 평문 계정 — 비교 시간을 해시 계정과 맞춘다
	BurnPasswordCheck(input)
	if subtle.ConstantTimeCompare([]byte(stored), []byte(input)) == 1 {
		return true, true
	}

	return false, false
}

var lockedPasswordOnce sync.Once
var lockedPasswordHash string

// LockedPasswordHash 는 비밀번호를 알 수 없는 계정용 해시 (외부 엑셀 연동으로 만든 계정).
// 아무도 모르는 임의 값의 해시라 관리자가 비밀번호를 지정하기 전에는 로그인할 수 없다 (한 번만 계산)
func LockedPasswordHash() string {
	lockedPasswordOnce.Do(func() {
		buf := make([]byte, 32)
		if _, err := rand.Read(buf); err != nil {
			panic(err)
		}
		lockedPasswordHash = HashPassword(hex.EncodeToString(buf))
	})
	return lockedPasswordHash
}
