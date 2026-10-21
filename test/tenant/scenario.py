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
    s,j=req('GET','/api/jwt?'+urllib.parse.urlencode({'loginid':i,'passwd':p})); return j.get('token')
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
s,j=req('GET','/api/company', A); check('A 업체목록=자기+연결 건물주', ids(j)==[1,3], ids(j))
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
print('FAILS', fails)
