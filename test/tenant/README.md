# 업체(테넌트) 분리 시나리오 테스트

두 점검업체(A·B)와 건물주 업체, 사용자·건물·설비·점검·청구를 넣은 로컬 DB 에 서버를 띄우고,
A 사용자가 B 데이터에 접근·수정하지 못하는지 HTTP 로 확인한다 (`router/tenant.go`, `router/signup.go`, `router/hardening.go`).

**운영·개발 DB 에 돌리지 말 것** — 사용자 상태를 바꾸고 데이터를 지운다.

```sh
# 1) 로컬 MariaDB (스키마는 개발 DB 에서 --no-data 로 덤프한 것 + view.sql)
docker run -d --name zkeep-testdb -e MARIADB_ROOT_PASSWORD=test -e MARIADB_DATABASE=zkeep \
  -e MARIADB_USER=zkeep -e MARIADB_PASSWORD=test -p 127.0.0.1:33306:3306 mariadb:10.11
docker exec -i zkeep-testdb mariadb -uroot -ptest zkeep < schema.sql   # DEFINER 제거한 덤프
docker exec -i zkeep-testdb mariadb -uroot -ptest zkeep < view.sql
docker exec -i zkeep-testdb mariadb -uroot -ptest zkeep < test/tenant/seed.sql

# 2) 서버: 로컬 DB 를 가리키는 .env.yml 이 있는 별도 폴더에서 실행 (develop, port 19303, db 127.0.0.1:33306 zkeep/test)
JWT_SECRET=local-test-secret ./zkeep

# 3) 테스트 (시드 직후 1회 — 가입 테스트가 계정을 만든다)
python3 test/tenant/scenario.py   # 마지막 줄 FAILS 0
```
