import json, urllib.request, urllib.parse, sys
import os
B=os.environ.get('ZKEEP_URL','http://localhost:19303')
def req(method, path, token=None, body=None):
    data = json.dumps(body).encode() if body is not None else None
    r = urllib.request.Request(B+path, data=data, method=method)
    r.add_header('Content-Type','application/json')
    if token: r.add_header('Authorization','Bearer '+token)
    try:
        with urllib.request.urlopen(r) as resp: return resp.status, json.loads(resp.read() or b'{}')
    except urllib.error.HTTPError as e:
        try: return e.code, json.loads(e.read() or b'{}')
        except Exception: return e.code, {}
def login(i,p):
    s,j=req('POST','/api/jwt',body={'loginid':i,'passwd':p}); return j.get('token')
A=login('a_admin','pw-a-admin'); AU=login('a_user','pw-a-user'); BB=login('b_admin','pw-b-admin'); R=login('root','pw-root')
assert A and AU and BB and R, 'login failed'
fails=0
def check(name, cond, info=''):
    global fails
    print(('PASS ' if cond else 'FAIL ')+name+('' if cond else f'  -> {info}'))
    if not cond: fails+=1
def ids(j): return sorted(x['id'] for x in (j.get('items') or []))
def secrets(j): s=json.dumps(j,ensure_ascii=False); return ('pw-' in s) or ('egiro-secret' in s)

# 1. 공개 API 축소
s,j=req('GET','/api/user'); check('비로그인 GET /api/user 차단', s==401, s)
s,j=req('GET','/api/user?loginid=a_admin&passwd=pw-a-admin'); check('비로그인 passwd 필터 차단', s==401, s)
s,j=req('POST','/api/user',body={'loginid':'evil','passwd':'x','level':4}); check('비로그인 POST /api/user 차단', s==401, s)
s,j=req('GET','/api/company'); check('비로그인 GET /api/company 차단', s==401, s)
s,j=req('GET','/api/department'); check('비로그인 GET /api/department 차단', s==401, s)
s,j=req('GET','/api/licensecategory'); check('비로그인 자격종류 조회 허용', s==200 and ids(j)==[1], (s,j))
s,j=req('POST','/api/signup/check',body={'loginid':'a_admin','email':'none@x.com','companyno':'111-11-11111'}); check('가입 중복확인', j.get('loginid') is True and j.get('email') is False and j.get('companyno') is True, j)
s,j=req('GET','/api/signup/companies'); check('가입 업체목록은 점검업체 id/name만', sorted(x['id'] for x in j['items'])==[1,2] and set(j['items'][0].keys())=={'id','name'}, j)
s,j=req('POST','/api/signup/user',body={'user':{'loginid':'newbie','passwd':'p','name':'신입','email':'n@n.com','company':1,'department':1,'level':4},'licenses':[{'licensecategory':1,'licenselevel':1}]}); check('직원 가입(레벨 무시)', j.get('code')=='ok', j)
nid=j.get('user')
s,j=req('POST','/api/signup/user',body={'user':{'loginid':'bad','passwd':'p','name':'x','company':3}}); check('건물주 업체로 직원 가입 거부', j.get('code')=='error', j)
s,j=req('POST','/api/signup/company',body={'company':{'name':'C점검','companyno':'999-99-99999'},'user':{'loginid':'c_admin','passwd':'p','name':'C대표','email':'c@c.com'}}); check('업체 가입', j.get('code')=='ok', j)
s,j=req('GET','/api/user/%d'%nid, R); check('가입 직원은 일반·승인대기', j['item']['level']==1 and j['item']['approval']==1, j.get('item'))
s,j=req('GET','/api/license?user=%d'%nid, R); check('가입 자격증 저장', len(j.get('items') or [])==1, j)

# 2. 목록 업체 고정
s,j=req('GET','/api/customer', A); check('A 고객목록=A것만', ids(j)==[1], ids(j))
s,j=req('GET','/api/customer?company=2', A); check('company=2 지정해도 A것만', ids(j)==[1], ids(j))
s,j=req('GET','/api/billing', A); check('A 청구목록', ids(j)==[1], ids(j))
s,j=req('GET','/api/report', A); check('A 점검목록', ids(j)==[1000], ids(j))
s,j=req('GET','/api/user', A); check('A 사용자목록', 3 not in ids(j) and 1 in ids(j), ids(j))
s,j=req('GET','/api/company', A); check('A 업체목록=자기+연결 건물주(type 0 포함)', ids(j)==[1,3,6], ids(j))
s,j=req('GET','/api/building', A); check('A 건물목록', ids(j)==[10,30], ids(j))
s,j=req('GET','/api/facility', A); check('A 설비목록', ids(j)==[100], ids(j))
s,j=req('GET','/api/item', A); check('A 점검항목', ids(j)==[1], ids(j))
s,j=req('GET','/api/license', A); check('A 자격증(B 것 없음)', 2 not in ids(j) and 1 in ids(j), ids(j))
s,j=req('GET','/api/statisticsmonth', A); tot=[x['totalprice'] for x in j.get('items') or []]; check('A 월매출=A것만', tot==[110000], j)
s,j=req('GET','/api/customer', R); check('전체관리자는 전체', ids(j)==[1,2], ids(j))

# 3. id 접근
for path in ['/api/customer/2','/api/building/20','/api/facility/200','/api/report/2000','/api/item/2','/api/user/3','/api/company/2','/api/billing/2','/api/license/2']:
    s,j=req('GET',path,A); check('A가 B의 '+path+' 조회 거부', s==403, (s,j))
s,j=req('GET','/api/customer/1', A); check('A 자기 고객 조회', s==200 and j['item']['id']==1, (s,j))
s,j=req('GET','/api/company/1', A); check('업체 응답에 e지로 비밀번호 없음', s==200 and not secrets(j), j)
s,j=req('GET','/api/user/1', A); check('사용자 응답에 비밀번호 없음', not secrets(j), j)
s,j=req('GET','/api/report', A); check('점검 목록 Extra 에도 비밀번호 없음', not secrets(j), json.dumps(j,ensure_ascii=False)[:300])

# 4. 쓰기
s,j=req('PUT','/api/customer',A,{'id':2,'company':2,'building':20}); check('A가 B 고객 수정 거부', s==403, s)
s,j=req('PUT','/api/customer/remark',A,{'id':2,'remark':'hack'}); check('A가 B 고객 필드수정 거부', s==403, s)
s,j=req('PUT','/api/customer/company',A,{'id':1,'company':2}); check('업체 이동 필드수정 거부', s==403, s)
s,j=req('DELETE','/api/facility',A,{'id':200}); check('A가 B 설비 삭제 거부', s==403, s)
s,j=req('DELETE','/api/facility/batch',A,[{'id':100},{'id':200}]); check('batch에 B 섞이면 거부', s==403, s)
s,j=req('DELETE','/api/facility/bybuilding',A,{'building':20}); check('B 건물 설비 일괄삭제 거부', s==403, s)
s,j=req('POST','/api/report',A,{'title':'x','building':20,'user':2}); check('B 건물로 점검 생성 거부', s==403, (s,j))
s,j=req('POST','/api/customer',A,{'company':2,'building':10,'user':2}); cid=j.get('id')
s2,j2=req('GET','/api/customer?building=10',R); check('POST 시 company를 A로 강제', any(x['company']==1 for x in j2['items']) and not any(x['company']==2 for x in j2['items']), j2)
s,j=req('POST','/api/facility',A,{'name':'새설비','building':10}); check('A 건물에 설비 추가', s==200 and j.get('code')=='ok', (s,j))
s,j=req('POST','/api/customercompany',A,{'customer':4}); check('B의 건물주 연결 거부', s==403, s)
s,j=req('POST','/api/customercompany',A,{'customer':2}); check('다른 점검업체 연결 거부', s==403, s)
s,j=req('POST','/api/customercompany',A,{'customer':5}); check('미연결 건물주 연결 허용', s==200, (s,j))
s,j=req('PUT','/api/notice',A,{'id':1,'title':'x'}); check('공지 수정은 전체관리자만', s==403, s)
s,j=req('GET','/api/giro',A); check('e지로 수납 목록은 전체관리자만', s==403, s)

# 5. 권한 상승
s,j=req('POST','/api/user',A,{'loginid':'r2','passwd':'p','level':4}); check('업체관리자가 전체관리자 생성 거부', s==403, s)
s,j=req('PUT','/api/user/level',AU,{'id':2,'level':3}); check('일반직원 자기 레벨 상승 거부', s==403, s)
s,j=req('PUT','/api/user',AU,{'id':2,'loginid':'a_user','name':'바뀐이름','level':3,'approval':3,'status':1,'company':1}); 
s2,j2=req('GET','/api/user/2',R); check('일반직원 전체수정 시 레벨 유지·이름 변경', j2['item']['level']==1 and j2['item']['name']=='바뀐이름', j2['item'])
s,j=req('PUT','/api/user/name',AU,{'id':1,'name':'x'}); check('일반직원이 다른 사용자 수정 거부', s==403, s)
s,j=req('GET','/api/external',AU); check('일반직원 엑셀 일괄처리 거부', s==403, s)
s,j=req('POST','/api/sms/index',AU,{'to':'010','message':'x'}); check('일반직원 SMS 발송 거부', s==403, s)
s2,j2=req('GET','/api/company/1',R)
s,j=req('PUT','/api/company',A,{'id':1,'name':'A점검(수정)','companyno':'111-11-11111','type':1,'egiropasswd':''})
import subprocess
v=subprocess.run(['docker','exec',os.environ.get('ZKEEP_DB_CONTAINER','zkeep-testdb'),'mariadb','-uroot','-ptest','zkeep','-N','-e','select c_egiropasswd from company_tb where c_id=1'],capture_output=True,text=True).stdout.strip()
check('업체 수정 시 빈 e지로 비밀번호는 기존 값 유지', v=='egiro-secret-A', v)

# 6. 정지된 사용자 토큰
subprocess.run(['docker','exec',os.environ.get('ZKEEP_DB_CONTAINER','zkeep-testdb'),'mariadb','-uroot','-ptest','zkeep','-e','update user_tb set u_status=2 where u_id=3'])
s,j=req('GET','/api/customer',BB); check('사용 중지된 사용자 토큰 거부', s==401, s)
subprocess.run(['docker','exec',os.environ.get('ZKEEP_DB_CONTAINER','zkeep-testdb'),'mariadb','-uroot','-ptest','zkeep','-e','update user_tb set u_status=1 where u_id=3'])

# 7. 커스텀 API 보강
s,j=req('GET','/api/billing/search/0?company=2',A); check('청구 검색 company=2 지정해도 A것만', all(x['company']==1 for x in j.get('items') or []) and len(j.get('items') or [])>=1, j)
s,j=req('POST','/api/dataitem/process',A,[{'data':{'id':2,'report':1000,'topcategory':1,'title':'t','company':2},'items':[{'id':2,'title':'x'}]}])
v=subprocess.run(['docker','exec',os.environ.get('ZKEEP_DB_CONTAINER','zkeep-testdb'),'mariadb','-uroot','-ptest','zkeep','-N','-e','select count(*) from data_tb where d_id=2 and d_company=2; select count(*) from data_tb where d_report=1000 and d_company=1; select count(*) from item_tb where i_id=2 and i_report=2000'],capture_output=True,text=True).stdout.split()
check('점검 일괄저장이 B 데이터·항목을 덮지 않음', v==['1','2','1'], v)
s,j=req('POST','/api/dataitem/process',A,[{'data':{'report':2000,'topcategory':1,'title':'t'},'items':[]}]); check('B 점검에 일괄저장 거부', j.get('code')=='error', j)
s,j=req('GET','/api/company/2',A); check('다른 점검업체 정보 조회 거부(재확인)', s==403, s)
subprocess.run(['docker','exec',os.environ.get('ZKEEP_DB_CONTAINER','zkeep-testdb'),'mariadb','-uroot','-ptest','zkeep','-e',"insert into customercompany_tb (cc_company,cc_customer,cc_date) values (1,2,now())"])
s,j=req('GET','/api/company/2',A); check('점검업체를 연결해도 정보 조회 거부', s==403, s)
import urllib.request as U
def upload(name, data):
    bnd='----x'; body=('--%s\r\nContent-Disposition: form-data; name="file"; filename="%s"\r\nContent-Type: application/octet-stream\r\n\r\n'%(bnd,name)).encode()+data+('\r\n--%s--\r\n'%bnd).encode()
    r=U.Request(B+'/api/upload/index',data=body,method='POST'); r.add_header('Content-Type','multipart/form-data; boundary='+bnd); r.add_header('Authorization','Bearer '+A)
    return json.loads(U.urlopen(r).read())
check('html 업로드 거부', upload('x.html',b'<script>alert(1)</script>').get('code')=='error')
check('png 업로드 허용', upload('x.png',b'\x89PNG').get('code')=='ok')

# 8. 리뷰 반영 (대소문자 키·HEAD·승인·상위 레벨·뷰 id 0·Content-Type)
s,j=req('PUT','/api/user',A,{'id':1,'loginid':'a_admin','name':'A관리자','Level':4,'company':1})
v=subprocess.run(['docker','exec',os.environ.get('ZKEEP_DB_CONTAINER','zkeep-testdb'),'mariadb','-uroot','-ptest','zkeep','-N','-e','select u_level from user_tb where u_id=1'],capture_output=True,text=True).stdout.strip()
check('대문자 Level 키로 전체관리자 상승 불가', v=='3', (s,j,v))
s,j=req('POST','/api/customer',A,{'Building':20,'user':2}); check('대문자 Building 키로 B 건물 참조 거부', s==403, (s,j))
s,j=req('POST','/api/customer',A,{'building':10,'Building':20}); check('대소문자만 다른 중복 키 거부', s==403, s)
s,j=req('HEAD','/api/customer/2',A); check('HEAD 로 B 고객 조회 거부', s==403, s)
s,j=req('HEAD','/api/customer/1',A); check('HEAD 로 자기 고객 조회 허용', s==200, s)
s,j=req('PUT','/api/user/passwd',A,{'id':4,'passwd':'hacked'}); check('업체관리자가 같은 업체 전체관리자 비밀번호 변경 거부', s==403, s)
s,j=req('PUT','/api/user/name',A,{'id':4,'name':'x'}); check('업체관리자가 전체관리자 정보 수정 거부', s==403, s)
s,j=req('GET','/api/statisticsmonth/0',A); check('통계 뷰 id 0 조회 거부', s==403, s)
s,j=req('GET','/api/company/6',A); check('c_type 0 건물주(연결됨) 조회 허용', s==200 and j['item']['id']==6, (s,j))
s,j=req('GET','/api/building/sum',A); check('업체 컬럼 없는 테이블 합계 거부', s==403, s)
r=urllib.request.Request(B+'/api/customer',data=b'{"building":10}',method='POST'); r.add_header('Content-Type','text/plain'); r.add_header('Authorization','Bearer '+A)
try: code=urllib.request.urlopen(r).status
except urllib.error.HTTPError as e: code=e.code
check('JSON 아닌 본문 쓰기 거부', code==403, code)
s,j=req('DELETE','/api/facility/batch',A,[None]); check('[null] 배치 거부(서버 오류 없이)', s==403, s)
s,j=req('POST','/api/jwt',body={'loginid':'newbie','passwd':'p'}); check('승인 대기 계정 로그인 거부', j.get('code')=='error', j)
subprocess.run(['docker','exec',os.environ.get('ZKEEP_DB_CONTAINER','zkeep-testdb'),'mariadb','-uroot','-ptest','zkeep','-e','update user_tb set u_approval=2 where u_id=2'])
s,j=req('GET','/api/customer',AU); check('거절된 계정의 기존 토큰 거부', s==401, s)
subprocess.run(['docker','exec',os.environ.get('ZKEEP_DB_CONTAINER','zkeep-testdb'),'mariadb','-uroot','-ptest','zkeep','-e','update user_tb set u_approval=3 where u_id=2'])
codes=[]
for i in range(12):
    s,j=req('POST','/api/signup/company',body={'company':{'name':'RL%d'%i,'companyno':'rl-%d'%i},'user':{'loginid':'rl%d'%i,'passwd':'p','name':'x'}}); codes.append(s)
check('가입 요청 횟수 제한(시간당 10회)', 429 in codes and codes.count(200)<=10, codes)

# 9. 4단계: 로그인·비밀번호 해시
DB=os.environ.get('ZKEEP_DB_CONTAINER','zkeep-testdb')
def sql(q): return subprocess.run(['docker','exec',DB,'mariadb','-uroot','-ptest','zkeep','-N','-e',q],capture_output=True,text=True).stdout.strip()
s,j=req('GET','/api/jwt?loginid=a_admin&passwd=pw-a-admin'); check('GET 로그인 차단(비밀번호 URL 노출)', s==405, (s,j))
check('평문 계정은 로그인 후 bcrypt 로 전환', sql("select left(u_passwd,4) from user_tb where u_id=1")=='$2a$')
check('전환 후에도 같은 비밀번호로 로그인', login('a_admin','pw-a-admin') is not None)
s,j=req('POST','/api/jwt',body={'loginid':'a_admin','passwd':'wrong'}); check('틀린 비밀번호 거부', j.get('code')=='error', j)
check('가입 계정 비밀번호는 해시로 저장', sql("select left(u_passwd,4) from user_tb where u_loginid='c_admin'")=='$2a$')
s,j=req('POST','/api/user',A,{'loginid':'made_by_admin','passwd':'pw-new','name':'신규','level':1,'status':1,'approval':3})
check('관리자가 만든 계정 비밀번호 해시', sql("select left(u_passwd,4) from user_tb where u_loginid='made_by_admin'")=='$2a$')
check('관리자가 만든 계정 로그인', login('made_by_admin','pw-new') is not None)
s,j=req('PUT','/api/user/passwd',A,{'id':int(sql("select u_id from user_tb where u_loginid='made_by_admin'")),'passwd':'pw-changed'})
check('비밀번호 변경도 해시·로그인', sql("select left(u_passwd,4) from user_tb where u_loginid='made_by_admin'")=='$2a$' and login('made_by_admin','pw-changed') is not None)
import base64
tok=login('a_admin','pw-a-admin'); payload=json.loads(base64.urlsafe_b64decode(tok.split('.')[1]+'=='))
import time as _t
check('토큰 유효기간 30일', 29*86400 < payload['exp']-_t.time() <= 30*86400+60, payload['exp']-_t.time())

# 9-2. 비밀번호 해시는 인증 뒤, 권한 있는 쓰기만
s,j=req('PUT','/api/user/passwd',AU,{'id':2,'passwd':'pw-a-user-2'})
check('일반 직원 자기 비밀번호 변경 → 해시·로그인', sql("select left(u_passwd,4) from user_tb where u_id=2")=='$2a$' and login('a_user','pw-a-user-2') is not None)
s,j=req('POST','/api/user/batch',A,[{'loginid':'bulk%d'%i,'passwd':'p','name':'x','level':1} for i in range(201)])
check('일괄 등록 200건 초과 거부(413)', s==413, s)
s,j=req('POST','/api/user/batch',body=[{'loginid':'anon%d'%i,'passwd':'p'} for i in range(5)])
check('비로그인 일괄 등록 401 (해시 전에 거부)', s==401, s)

# 9-3. 비밀번호 검색 필터 무시·로그인 실패 응답 통일
sql("set session sql_mode=''; insert into user_tb (u_loginid,u_passwd,u_name,u_level,u_approval,u_status,u_company,u_date) values ('legacy','pw-legacy','예전계정',1,3,1,1,now())")
s,j=req('GET','/api/user/count',A); total=j.get('total')
s,j=req('GET','/api/user/count?passwd=pw-legacy',A); check('?passwd= 필터 무시(평문 비밀번호 추측 불가)', j.get('total')==total and total and total>1, (j,total))
s,j=req('GET','/api/user?passwd=pw-legacy',A); check('목록 ?passwd= 필터 무시', len(j.get('items') or [])==total, len(j.get('items') or []))
s1,j1=req('POST','/api/jwt',body={'loginid':'no-such-user','passwd':'x'}); s2,j2=req('POST','/api/jwt',body={'loginid':'a_admin','passwd':'wrong'})
check('없는 아이디·틀린 비밀번호 응답 동일', j1==j2 and j1.get('message')=='invalid login', (j1,j2))
check('예전 평문 계정 로그인 후 해시 전환', login('legacy','pw-legacy') is not None and sql("select left(u_passwd,4) from user_tb where u_loginid='legacy'")=='$2a$')
# 10. 회원 탈퇴
s,j=req('POST','/api/account/withdraw',AU,{'passwd':'wrong'}); check('탈퇴: 비밀번호 틀리면 거부', j.get('code')=='error', j)
s,j=req('POST','/api/account/withdraw',A,{'passwd':'pw-a-admin'}); check('탈퇴: 직원이 남은 마지막 관리자는 거부', j.get('code')=='error' and '관리자' in j.get('message',''), j)
s,j=req('POST','/api/account/withdraw',AU,{'passwd':'pw-a-user-2'}); check('탈퇴: 일반 직원 탈퇴', j.get('code')=='ok', j)
check('탈퇴 후 로그인 불가', login('a_user','pw-a-user-2') is None)
check('탈퇴 계정 개인정보 삭제', sql("select concat(u_name,'|',u_email,'|',u_status) from user_tb where u_id=2")=='탈퇴회원||2')
s,j=req('GET','/api/customer',AU); check('탈퇴 계정 기존 토큰 거부', s==401, s)
s,j=req('POST','/api/account/withdraw',body={'passwd':'x'}); check('탈퇴: 비로그인 거부', s==401, s)
s,j=req('POST','/api/account/withdraw',A,{'passwd':'x'*300}); check('탈퇴: 너무 긴 비밀번호 거부', j.get('code')=='error', j)
# 같은 업체 관리자 둘이 동시에 탈퇴해도 관리자가 0명이 되지 않는다
sql("set session sql_mode=''; insert into user_tb (u_loginid,u_passwd,u_name,u_level,u_approval,u_status,u_company,u_date) values ('a_admin2','pw-a-admin2','A관리자2',3,3,1,1,now())")
A2=login('a_admin2','pw-a-admin2')
import threading
res={}
def go(name,tok,pw): res[name]=req('POST','/api/account/withdraw',tok,{'passwd':pw})[1]
t1=threading.Thread(target=go,args=('a1',A,'pw-a-admin')); t2=threading.Thread(target=go,args=('a2',A2,'pw-a-admin2'))
t1.start(); t2.start(); t1.join(); t2.join()
oks=[k for k,v in res.items() if v.get('code')=='ok']
check('동시 탈퇴: 관리자 한 명만 탈퇴되고 한 명은 남는다', len(oks)==1 and sql("select count(*) from user_tb where u_company=1 and u_level=3 and u_status=1")=='1', (res, sql("select count(*) from user_tb where u_company=1 and u_level=3 and u_status=1")))
print('FAILS', fails)
