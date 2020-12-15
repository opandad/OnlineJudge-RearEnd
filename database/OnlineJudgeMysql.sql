/*
作者：唐仕航
功能：建立oj的mysql数据库
日期：2020.12.4 16：22
版本：0.0.1
*/
create database online_judge;
use online_judge;

-- 问题集合，这里只是做个索引，问题储存是放在mongodb中
CREATE TABLE IF NOT EXISTS problem_set_table(
	`problem_id` int unsigned auto_increment,
	`problem_name` TINYTEXT,
	`is_robot_problem` bool,
	`problem_accept_number` int unsigned,
	`problem_fail_number` int unsigned,
	primary key(`problem_id`)
)DEFAULT CHARSET=utf8;

create table IF NOT EXISTS user_table(
	`user_email` char(30),
	`user_name` char(20),
	`user_password` char(30),
	`user_authority` int,
	primary key(`user_email`)
);

create table IF NOT EXISTS contest_table(
	`contest_id` int unsigned auto_increment,
	`contest_name` tinytext,
	`contest_start_time` datetime,
	`contest_end_time` datetime, -- 可能需要修改
	`is_official_match` bool,
	primary key(`contest_id`)
);


create table IF NOT EXISTS submit_status_table(
	`submit_id` int unsigned,
	`problem_id` int unsigned,
	`submit_state` char(50),
	`return_result` text,
	`submit_time` datetime,
	`language` char(50),
	`contest_id` int unsigned,
	`user_email` char(30),
	`run_time` smallint unsigned,
	primary key(`submit_id`)
);


create table IF NOT EXISTS support_language_table(
	`language` char(50),
	`submit_command` text,
	primary key(`language`)
);


-- 有bug，主键唯一bug，需要修改
create table IF NOT EXISTS offical_contest_users_table(
	`contest_id` int unsigned,
	`user_email` char(30),
	primary key(`contest_id`)
);

-- 有bug，主键唯一bug，需要修改
create table IF NOT EXISTS contest_problem_table(
	`contest_id` int unsigned,
	`problem_id` int unsigned,
	primary key(`contest_id`)
);
