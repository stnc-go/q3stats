select count(alias_id), from match_player_stats group by match_player_stats.match_id ;
select * from match_player_stats where match_id = 1;
select matches.id, matches.date, matches.map, count(match_player_stats.alias_id) from match_player_stats inner join matches on match_player_stats.match_id = matches.id group by matches.id order by matches.date desc;
select alias_id, score, player_aliases.alias from match_player_stats inner join player_aliases on match_player_stats.alias_id = player_aliases.id where match_id = 33 order by score desc limit 1;
select * from match_player_stats where alias_id = 9;
select * from player_aliases where id = 9;
select match_player_weapon_stats.type, match_player_weapon_stats.hits, match_player_weapon_stats.shots, match_player_weapon_stats.kills 
from match_player_weapon_stats 
inner join match_player_stats on match_player_stats.id = match_player_weapon_stats.match_player_stats_id where alias_id = 9 group by type;
select match_player_weapon_stats.type, match_player_weapon_stats.hits, match_player_weapon_stats.shots, match_player_weapon_stats.kills 
from match_player_weapon_stats 
inner join match_player_stats on match_player_stats.id = match_player_weapon_stats.match_player_stats_id where alias_id = 9;

-- select weapon stats by player
select match_player_weapon_stats.type, sum(match_player_weapon_stats.hits), sum(match_player_weapon_stats.shots), sum(match_player_weapon_stats.kills) 
from match_player_weapon_stats where match_player_weapon_stats.match_player_stats_id in (select id from match_player_stats where alias_id = 1)
group by match_player_weapon_stats.type;

-- select items stats by player
select match_player_items_stats.type, sum(match_player_items_stats.pickups), sum(match_player_items_stats.time)
from match_player_items_stats where match_player_items_stats.match_player_stats_id in (select id from match_player_stats where alias_id = 1)
group by match_player_items_stats.type;

select id from match_player_stats where alias_id = 9;

-- player matches
select count(*) from match_player_stats where alias_id = 9;
-- player frags, deaths, suicides
select sum(kills), sum(deaths), sum(suicides) from match_player_stats where alias_id = 1;
-- player hits shots
select sum(hits), sum(shots) from match_player_weapon_stats where match_player_stats_id in (select id from match_player_stats where alias_id = 9);
-- player wins
select count(*) from matches where winner_alias_id = 1;
select * from match_player_stats where match_id = 216;
update 

select * from player_aliases;
select * from players;
select * from player_aliases where player_id = 6;
update player_aliases set player_id = 6 where id in (1);
update player_aliases set player_id = 0 where player_id = 6;
select name from players where id in (select player_id from player_aliases where id = 10);

-- match statistics
select 
id, alias, players.id, 
score, kills,  deaths, 
suicides, damage_given, 
damage_taken, health_total, armor_total 
from match_player_stats  
inner join player_aliases 
on match_player_stats.alias_id = player_aliases.id 
inner join players
on player_aliases.player_id = players.id
where match_id = 1 
order by 
score 
desc;

select * from match_player_weapon_stats where match_player_stats_id in (select id from match_player_stats where match_id = 1) order by match_player_stats_id;

select * from match_player_weapon_stats inner join match_player_stats 
on match_player_weapon_stats.match_player_stats_id = match_player_stats.id 
where match_player_stats.match_id = 1 order by score desc;