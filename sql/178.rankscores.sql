-- use leetcode;
-- create table if not exists scores (
--     `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
--     `score` float(3,2) NOT NULL DEFAULT '0.00',
--     PRIMARY KEY (`id`)
-- );
-- insert into scores (id, score) values ('1', '3.5'),
-- ('2', '3.65'),
-- ('3', '4.0'),
-- ('4', '3.85'),
-- ('5', '4.0'),
-- ('6', '3.65');


-- sql
select a.score as score, 
    (select count(distinct(b.score)) from scores b where b.score >= a.score) as rank 
from scores a 
order by a.score desc;
