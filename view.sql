drop view billinglist_vw;
create view billinglist_vw as
select billing_tb.*, b_name as bi_buildingname, c_billingname as bi_billingname, c_billingtel as bi_billingtel, c_billingemail as bi_billingemail from billing_tb, building_tb, company_tb where bi_building = b_id and b_company = c_id;

drop view userlist_vw;
create view userlist_vw as select *, ifnull((select sum(b_score) from building_tb, customer_tb where b_id = cu_building and cu_user = u_id), 0) as u_totalscore from user_tb;


drop view statisticsyear_vw;
create view statisticsyear_vw as
select
0 as bi_id,
date_format(bi_billdate, '%Y') as bi_duration,
count(bi_id) as bi_total,
count(bi_id)* bi_price as bi_totalprice,
now() as bi_billdate
from billing_tb where bi_status = 2
group by date_format(bi_billdate, '%Y');


drop view statisticsmonth_vw;
create view statisticsmonth_vw as
select
0 as bi_id,
date_format(bi_billdate, '%Y') as bi_year,
date_format(bi_billdate, '%Y-%m') as bi_duration,
count(bi_id) as bi_total,
count(bi_id)* bi_price as bi_totalprice,
now() as bi_billdate
from billing_tb where bi_status = 2
group by date_format(bi_billdate, '%Y'), date_format(bi_billdate, '%Y-%m');

drop view statisticsday_vw;
create view statisticsday_vw as
select
0 as bi_id,
date_format(bi_billdate, '%Y-%m') as bi_month,
date_format(bi_billdate, '%Y-%m-%d') as bi_duration,
count(bi_id) as bi_total,
count(bi_id)* bi_price as bi_totalprice,
now() as bi_billdate
from billing_tb where bi_status = 2
group by date_format(bi_billdate, '%Y-%m'), date_format(bi_billdate, '%Y-%m-%d');

drop view calendarcompanylist_vw;
create view calendarcompanylist_vw as
select r_company as r_id, r_company, date_format(r_checkdate, '%Y-%m') as r_month, date_format(r_checkdate, '%Y-%m-%d') as r_day, r_status, count(*) as r_count, now() as r_checkdate from report_tb group by r_company, date_format(r_checkdate, '%Y-%m'), date_format(r_checkdate, '%Y-%m-%d'), r_status;


/*
drop view customercompany_vw;
create view customercompany_vw as
select cu_company as c_company, c_id, c_name, c_companyno, c_ceo, c_address, c_addressetc, c_type, c_status, c_date, 
count(*) as c_buildingcount, sum(cu_contractprice) as c_contractprice from company_tb, building_tb, customer_tb where c_id = b_company and b_id = cu_building group by
cu_company, c_id, c_name, c_companyno, c_ceo, c_address, c_addressetc, c_type, c_status, c_date;
*/

drop view companylist_vw;
create view companylist_vw as 
select c_id, c_name, c_companyno, c_ceo, c_address, c_addressetc, c_tel, c_email, c_date, cc_company as c_company, count(*) as c_buildingcount, sum(cu_contractprice + cu_contractvat) as c_contractprice from company_tb, customercompany_tb, building_tb, customer_tb where c_id = cc_customer and c_id = b_company and b_id = cu_building
group by c_id, c_name, c_companyno, c_ceo, c_address, c_addressetc, c_tel, c_email, c_date, cc_company;
