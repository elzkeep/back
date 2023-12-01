drop view gamelist_vw;
create view gamelist_vw as select game_tb.*, gu_user as g_gameuser from game_tb, gameuser_tb where g_id = gu_game;

drop view statisticsfaction_vw;
create view statisticsfaction_vw as select gu_faction, sum(if(gu_rank = 1, 1, 0))as gu_rank1, sum(if(gu_rank = 2, 1, 0)) as gu_rank2, sum(if(gu_rank = 3, 1, 0)) as gu_rank3, sum(if(gu_rank = 4, 1, 0)) as gu_rank4, sum(if(gu_rank = 5, 1, 0)) as gu_rank5, count(*) as gu_count, convert(avg(gu_score), int) as gu_avg from game_tb, gameuser_tb where g_status = 4 and g_id = gu_game group by gu_faction;


drop view statisticscolor_vw;
create view statisticscolor_vw as select gu_color, sum(if(gu_rank = 1, 1, 0))as gu_rank1, sum(if(gu_rank = 2, 1, 0)) as gu_rank2, sum(if(gu_rank = 3, 1, 0)) as gu_rank3, sum(if(gu_rank = 4, 1, 0)) as gu_rank4, sum(if(gu_rank = 5, 1, 0)) as gu_rank5, count(*) as gu_count, convert(avg(gu_score), int) as gu_avg from game_tb, gameuser_tb where g_status = 4 and g_id = gu_game group by gu_color;


drop view statistics_vw;
create view statistics_vw as select gu_faction, gu_color, sum(if(gu_rank = 1, 1, 0))as gu_rank1, sum(if(gu_rank = 2, 1, 0)) as gu_rank2, sum(if(gu_rank = 3, 1, 0)) as gu_rank3, sum(if(gu_rank = 4, 1, 0)) as gu_rank4, sum(if(gu_rank = 5, 1, 0)) as gu_rank5, count(*) as gu_count, convert(avg(gu_score), int) as gu_avg from game_tb, gameuser_tb where g_status = 4 and g_id = gu_game group by gu_faction, gu_color;


drop view ranklist_vw;
create view ranklist_vw as
select gu_user as r_id, u_name as r_name, sum(gu_elo) + 1000.0 as r_elo, avg(gu_score) as r_score, count(*) as  r_count, sum(if(gu_rank = 1, 1, 0)) as r_rank1, sum(if(gu_rank = 2, 1, 0)) as r_rank2, sum(if(gu_rank = 3, 1, 0)) as r_rank3, sum(if(gu_rank = 4, 1, 0)) as r_rank4, sum(if(gu_rank = 5, 1, 0)) as r_rank5 from gameuser_tb, user_tb where gu_user = u_id group by gu_user, u_name order by sum(gu_elo) desc;
