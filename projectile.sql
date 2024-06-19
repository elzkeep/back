select * from company_tb a, (select c_id as c_customercompany, count(*) as c_buildingcount, sum(cu_contractprice) as c_contractprice from company_tb, building_tb, customer_tb where c_id = b_company and b_id = cu_building group by c_id) as b where a.c_id = b.c_customercompany

/

select * from building_tb, customer_tb where b_id = cu_building

/

select * from customer_tb

/

select * from company_tb

/

create view customercompany_vw as 
select * from company_tb a, (select c_id as c_customercompany, count(*) as c_buildingcount, sum(cu_contractprice) as c_contractprice from company_tb, building_tb, customer_tb where c_id = b_company and b_id = cu_building group by c_id) as b where a.c_id = b.c_customercompany;

/

drop view customercompany_vw;
create view customercompany_vw as
select cu_company as c_company, c_id, c_name, c_companyno, c_ceo, c_address, c_addressetc, c_type, c_status, c_billingname, c_billingtel, c_billingemail, c_date, 
count(*) as c_buildingcount, sum(cu_contractprice) as c_contractprice from company_tb, building_tb, customer_tb where c_id = b_company and b_id = cu_building group by
cu_company, c_id, c_name, c_companyno, c_ceo, c_address, c_addressetc, c_type, c_status, c_billingname, c_billingtel, c_billingemail, c_date;

/

select cu_company as c_company, c_id, c_name, c_companyno, c_ceo, c_address, c_addressetc, c_type, c_status, c_billingname, c_billingtel, c_billingemail, c_date, 
count(*) as c_buildingcount, sum(cu_contractprice) as c_contractprice from company_tb, building_tb, customer_tb where c_id = b_company and b_id = cu_building group by
cu_company, c_id, c_name, c_companyno, c_ceo, c_address, c_addressetc, c_type, c_status, c_billingname, c_billingtel, c_billingemail, c_date;

/

drop view billinglist_vw;
create view billinglist_vw as
select billing_tb.*, b_name as bi_buildingname, c_billingname as bi_billingname, c_billingtel as bi_billingtel, c_billingemail as bi_billingemail from billing_tb, building_tb, company_tb where bi_building = b_id and b_company = c_id;

/

select c_id, c_name, c_companyno, c_ceo, c_address, c_addressetc, c_tel, c_email, c_date, cc_company as c_company, count(*) as c_buildingcount, sum(cu_contractprice + cu_contractvat) as c_contractprice
from company_tb
left join customercompany_tb on c_id = cc_customer
left join building_tb on c_id = b_company
left outer join customer_tb on b_id = cu_building4
group by c_id, c_name, c_companyno, c_ceo, c_address, c_addressetc, c_tel, c_email, c_date, cc_company

/

drop view companylist_vw;
create view companylist_vw as 
select c_id, c_name, c_companyno, c_ceo, c_address, c_addressetc, c_tel, c_email, c_date, cc_company as c_company, count(*) as c_buildingcount, sum(cu_contractprice + cu_contractvat) as c_contractprice from company_tb, customercompany_tb, building_tb, customer_tb where c_id = cc_customer and c_id = b_company and b_id = cu_building
group by c_id, c_name, c_companyno, c_ceo, c_address, c_addressetc, c_tel, c_email, c_date, cc_company;


/

select * from company_tb where c_id = 5006

/

desc license_bt

/

select * from company_tb

/

select * from company_tb where c_name like '%동양%'

/

select * from customercompany_tb where cc_company = 1436;

/

delete from  customercompany_tb where cc_company = 9683;

/

desc customer_tb

/

delete from customer_tb where cu_company = 9683;

/

