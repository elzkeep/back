drop view billinglist_vw;
create view billinglist_vw as
select billing_tb.*, b_name as bi_buildingname, c_billingname as bi_billingname, c_billingtel as bi_billingtel, c_billingemail as bi_billingemail from billing_tb, building_tb, company_tb where bi_building = b_id and b_company = c_id;

drop view userlist_vw;
create view userlist_vw as select *, ifnull((select sum(b_score) from building_tb, customer_tb where b_id = cu_building and cu_user = u_id), 0) as u_totalscore from user_tb;
