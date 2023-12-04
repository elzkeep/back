select * from game_tb, gameuser_tb a, (select gu_faction, max(gu_score) as score from gameuser_tb group by gu_faction) as b
where g_status = 4 and g_id = a.gu_game and a.gu_faction = b.gu_faction and a.gu_score = b.score;

/

select a.g_id, a.g_count, b.g_faction, b.g_score, a.gu_user, u_name from (select * from game_tb, gameuser_tb, user_tb where g_status = 4 and g_id = gu_game and u_id = gu_user) a, (select g_count, gu_faction as g_faction, max(gu_score) as g_score from game_tb, gameuser_tb where g_status = 4 and g_id = gu_game group by g_count, gu_faction) b
       where a.g_count = b.g_count and a.gu_faction = b.g_faction and a.gu_score = b.g_score order by g_count, g_faction, g_score;

/

select * from gameuser_tb, game_tb where g_id = gu_game and g_status = 4 and g_count = 2 and gu_faction = 44;
