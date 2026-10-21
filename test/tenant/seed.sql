SET SESSION sql_mode='';
INSERT INTO company_tb (c_id,c_name,c_companyno,c_type,c_egiropasswd,c_date) VALUES
 (1,'A점검','111-11-11111',1,'egiro-secret-A',now()),(2,'B점검','222-22-22222',1,'egiro-secret-B',now()),
 (3,'X건물주','333-33-33333',2,'',now()),(4,'Y건물주','444-44-44444',2,'',now()),(5,'미연결건물주','555-55-55555',2,'',now());
INSERT INTO customercompany_tb (cc_id,cc_company,cc_customer,cc_date) VALUES (1,1,3,now()),(2,2,4,now());
INSERT INTO department_tb (de_id,de_name,de_company,de_date) VALUES (1,'A팀',1,now()),(2,'B팀',2,now());
INSERT INTO user_tb (u_id,u_loginid,u_passwd,u_name,u_email,u_level,u_approval,u_status,u_company,u_department,u_date) VALUES
 (1,'a_admin','pw-a-admin','A관리자','a@a.com',3,3,1,1,1,now()),(2,'a_user','pw-a-user','A직원','a2@a.com',1,3,1,1,1,now()),
 (3,'b_admin','pw-b-admin','B관리자','b@b.com',3,3,1,2,2,now()),(4,'root','pw-root','전체관리자','r@r.com',4,3,1,1,0,now());
INSERT INTO building_tb (b_id,b_name,b_company,b_date) VALUES (10,'X빌딩',3,now()),(20,'Y빌딩',4,now()),(30,'X신축',3,now());
INSERT INTO customer_tb (cu_id,cu_company,cu_building,cu_user,cu_status,cu_contractprice,cu_contractvat,cu_date) VALUES (1,1,10,2,1,100000,10000,now()),(2,2,20,3,1,200000,20000,now());
INSERT INTO facility_tb (f_id,f_category,f_name,f_building,f_date) VALUES (100,1,'X수전',10,now()),(200,1,'Y수전',20,now());
INSERT INTO report_tb (r_id,r_title,r_company,r_user,r_building,r_checkdate,r_date) VALUES (1000,'A점검',1,2,10,now(),now()),(2000,'B점검',2,3,20,now(),now());
INSERT INTO data_tb (d_id,d_title,d_report,d_company,d_date) VALUES (1,'A데이터',1000,1,now()),(2,'B데이터',2000,2,now());
INSERT INTO item_tb (i_id,i_title,i_report,i_data,i_date) VALUES (1,'A항목',1000,1,now()),(2,'B항목',2000,2,now());
INSERT INTO billing_tb (bi_id,bi_title,bi_price,bi_status,bi_company,bi_building,bi_billdate,bi_date) VALUES (1,'A청구',110000,3,1,10,'2026-10-01',now()),(2,'B청구',990000,3,2,20,'2026-10-01',now());
INSERT INTO license_tb (l_id,l_user,l_licensecategory,l_licenselevel,l_date) VALUES (1,2,1,1,now()),(2,3,1,1,now());
INSERT INTO licensecategory_tb (lc_id,lc_name,lc_date) VALUES (1,'전기기사',now());
INSERT INTO licenselevel_tb (ll_id,ll_name,ll_date) VALUES (1,'특급',now());
INSERT INTO notice_tb (n_id,n_title,n_date) VALUES (1,'공지',now());
