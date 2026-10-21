# 지킴E (zkeep) — 프로젝트 설명

전기안전관리 대행업체를 위한 업무 시스템. 점검업체가 고객 건물의 전기설비를 등록하고, 현장에서 정기 점검표를 작성하고,
「전기설비 점검결과 기록표」를 발행하고, 지로로 청구·수금하는 일을 한곳에서 한다.

> 저장소 3개: [`back`](https://github.com/elzkeep/back) (Go API + 웹 서빙) · [`front`](https://github.com/elzkeep/front) (업체 관리 웹) · [`app`](https://github.com/elzkeep/app) (현장 점검 앱)
> 서비스 도메인 `zkeep.space` (운영), `dev.zkeep.space` (개발 서버)

## 1. 누가 쓰나

| 사용자 | 레벨 | 쓰는 곳 | 하는 일 |
|---|---|---|---|
| 현장 점검자(안전관리자) | 일반(1)·팀장(2) | 앱 | 오늘 점검 일정, 점검표 작성(사진·서명), 설비 등록, 보고서 공유 |
| 점검업체 관리자 | 관리자(3) | 웹 `/management` | 고객·건물·계약, 직원·팀·자격증, 점검 관리, 청구·지로·수금, 매출 통계 |
| 플랫폼 운영자 | 전체관리자(4) | 웹 `/admin` | 전체 업체·사용자, 공지·FAQ·가입문의 |
| 비로그인 | — | 웹·앱 | 회원가입(점검업체 / 업체 직원), 가입문의 |

업체 = 테넌트. 점검업체(`company.type = 1`)마다 데이터가 분리되고, 건물주 업체(`type = 2`)는 점검업체가 고객으로 등록한다.

## 2. 구성

```
[현장 앱 Flutter] ──┐
                    ├── HTTPS ──> [back: Go Fiber :9303] ──> MariaDB (zkeep)
[관리 웹 Vue 3] ────┘               ├─ /api/*  REST (JWT)
                                    ├─ /       관리 웹 정적 파일 (front 빌드)
                                    └─ /webdata 업로드 파일
```

| | 기술 | 비고 |
|---|---|---|
| back | Go 1.20+, Fiber v2, MariaDB, gopdf, excelize | 모델·CRUD 컨트롤러·라우터 대부분이 `buildtool-model`/`buildtool-router_fiber` 생성 코드 (`model.json` 기준) |
| front | Vue 3, Vite, TypeScript, Element Plus, Vuex | API 모델(`src/models`)도 생성 코드 |
| app | Flutter 3.44, GetX, `common_control`(사내 공용 위젯·HTTP) | 웹 빌드도 가능 (`zkeep.space` 에 배포) |

### 생성 코드와 수기 코드 (back)
- 생성: `models/*.go`, `models/<테이블>/`(enum), `controllers/rest/*.go`, `router/router.go`
  — 다시 생성하면 덮어쓴다. 보안 검사는 생성 파일 밖(`router/tenant.go` 등)에 둔다
- 수기: `controllers/api/*.go`(커스텀 로직), `router/{auth,tenant,signup,hardening,account}.go`, `services/`, `global/`, `config/`

## 3. 데이터 모델 (요약)

| 영역 | 테이블 | 설명 |
|---|---|---|
| 업체·사용자 | `company`, `user`, `department`, `license`(직원 자격증), `companylicense`(업체 자격), `licensecategory`·`licenselevel` | 업체 = 테넌트, 사용자 레벨·승인·상태 |
| 고객 | `customercompany`(점검업체 ↔ 건물주 연결), `building`(건물·수전/발전 용량·관리점수), `customer`(계약: 업체+건물, 금액·주기·담당) | 건물의 `b_company` 는 건물주 업체 |
| 설비 | `facility`(건물별 수전·발전·태양광·ESS·UPS 등, `value1..25`) | |
| 점검 | `report`(점검: 업체·건물·점검자·일시·서명·종합의견), `data`(설비 구분별 묶음), `item`(항목: 측정값·선택값·판정·사유·조치·사진) | 판정: 적합/부적합/요주의/해당없음 |
| 청구 | `billing`(청구), `billinghistory`(입금), `giro`(e지로 수납) | 지로 OCR 고지서 출력 |
| 통계 뷰 | `statisticsday/month/year_vw`, `billinglist_vw`, `companylist_vw` 등 | 정의는 `view.sql` (업체별 집계) |
| 웹사이트 | `notice`, `webnotice`, `webfaq`, `webjoin`(가입문의), `webquestion`(문의) | |

## 4. 주요 기능

**현장(앱)** 오늘 점검 일정 · 점검 목록 · 점검표 작성(저압·고압·변전·부하·발전·기타·태양광·충전기·ESS·UPS·연료전지·풍력·수력·열화상, 사진·서명) · 설비 등록 · 고객 조회 · 보고서 PDF 공유 · 마이페이지(자격증·이력) · 회원가입·탈퇴

**관리(웹)** 대시보드(공지·현황·점검 캘린더) · 고객현황·건물/계약(엑셀 일괄등록) · 설비 · 점검관리 · 매출(청구 생성) · 지로 고지서 · 수금(입금 처리) · 매출 현황(엑셀) · 직원·팀·자격증 · 업체정보

**서버** 점검 보고서 PDF 생성(`/api/report/download/:id`) · 지로 고지서 PDF · 엑셀 양식·일괄등록 · 관리점수 계산 · SMS·메일

## 5. 인증과 업체 분리

- 로그인 `POST /api/jwt` → JWT(30일, `JWT_SECRET` 필수). 요청마다 DB 에서 사용자 레벨·업체·상태·승인을 다시 읽는다
- 비밀번호 bcrypt (평문으로 저장돼 있던 예전 계정은 로그인 시 전환)
- 비로그인 공개: `POST /api/jwt`, `/api/signup/*`(가입·중복확인·업체/팀 선택지), 공지·문의 등록, 자격 종류 — 나머지 `/api/*` 는 토큰 필요
- 업체 분리(`router/tenant.go`): 전체관리자가 아니면 목록은 자기 업체로 고정, id 접근·쓰기는 소유 확인, 참조(건물·점검·사용자)도 확인,
  응답의 비밀번호 제거. 엑셀 일괄등록·메일·SMS 는 업체 관리자 이상
- 검증: `test/tenant/scenario.py` (두 업체 시나리오 98개)

## 6. 실행·배포

```sh
# back — .env.yml.example 을 .env.yml 로 복사해 값 채움 (또는 DB_PASSWORD 등 환경변수)
JWT_SECRET=<32자 이상> go run main.go            # :9303, develop 모드
JWT_SECRET=... ./main --mode production

# front
pnpm install && pnpm dev                          # :4000 → API localhost:9303
pnpm release                                      # 운영 빌드 → back 의 dist

# app
flutter run --dart-define=API_BASE_URL=http://localhost:9303
```

배포: back `Makefile` — `Dockerfile_admin`(관리 웹 포함) 이미지 `netb.co.kr:5000/zkeep_admin` → `dev.zkeep.space` docker-compose.
DB 뷰 변경(`view.sql`)은 운영 DB 에 직접 적용한다.

## 7. 현황 (2026-10-21)

2023-09 시작, 2024-07 「수금 관리 작업중」 이후 중단 → 2026-10 재개. 기능은 대부분 있으나 운영 전 정리 중.

| 단계 | 내용 | PR |
|---|---|---|
| 보안 1·2 | 공개 API 축소·가입 API, 업체 분리, 레벨 상승·승인 강제 | back#3, front#3, app#3 |
| 앱 빌드 | Android 툴체인, 주소검색 대체, common_control 메서드 | app#9, common_control#23 |
| 보고서 PDF | 점검 데이터로 기록표 생성 | back#5 |
| 비밀값·인증 | 비밀 파일 추적 해제, JWT 키 필수, bcrypt, POST 로그인 | back#7, front#4, app#6 |
| 출시 준비 | 앱 이름·아이콘·권한·서명·탈퇴 | back#9, app#8 |

**남은 일**
- 공개 저장소에 노출됐던 비밀값 교체 (DB·SMS·Gmail 비밀번호, SSL 인증서), 운영 `JWT_SECRET` 설정
- 번들 ID 확정·스토어 계정·개인정보처리방침 (`app/docs/RELEASE.md`)
- 업체 컬럼 없는 목록(업체·건물·설비)의 업체 필터를 SQL 로 (현재 응답 후처리 — 페이지네이션 부정확)
- 세금계산서 화면(껍데기), 아이디·비밀번호 찾기, 관리자 웹 메뉴, 레벨 1·2 웹 접근
- 점검 항목 정의를 서버 스키마로 (네이티브 앱 전환 계획 `app/docs/NATIVE_PLAN.md`)

관련 문서: [시장 조사](MARKET.md) · [앱 출시 체크리스트](https://github.com/elzkeep/app/blob/main/docs/RELEASE.md) · [네이티브 전환 계획](https://github.com/elzkeep/app/blob/main/docs/NATIVE_PLAN.md)
