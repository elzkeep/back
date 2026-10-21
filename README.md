# zkeep back

지킴E(전기안전관리 대행업체 업무 시스템)의 API 서버. Go + Fiber + MariaDB. 관리 웹(`front` 빌드)도 함께 서빙한다.

- 무엇을 하나 · 데이터 모델 · 인증/업체 분리 · 현황: **[docs/PROJECT.md](docs/PROJECT.md)**
- 시장 조사: [docs/MARKET.md](docs/MARKET.md)

## 실행

```sh
cp .env.yml.example .env.yml          # DB·SMS·메일 값 채우기 (또는 DB_PASSWORD 등 환경변수)
export JWT_SECRET=<32자 이상 임의 문자열>   # 필수 (운영은 없으면 기동 거부)
go run main.go                        # :9303 develop
go test ./...
```

| 경로 | 내용 |
|---|---|
| `router/router.go`, `controllers/rest`, `models` | **생성 코드** (`make model`/`make router`, `model.json` 기준) — 직접 고치지 않는다 |
| `router/{auth,tenant,signup,hardening,account}.go` | 인증·업체 분리·공개 가입 API·응답 비밀번호 제거·탈퇴 |
| `controllers/api` | 커스텀 API (청구·지로·엑셀·보고서 PDF·대시보드) |
| `view.sql` | 통계·목록 뷰 — 바꾸면 운영 DB 에 직접 적용 |
| `test/tenant` | 두 업체 시나리오 테스트 (로컬 DB) |

배포: `make` 타깃(`Dockerfile_admin` → `netb.co.kr:5000/zkeep_admin` → `dev.zkeep.space`). 비밀값(`.env.yml`, `model.json`, `certs/`)은 저장소에 올리지 않는다.
