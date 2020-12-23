/*
功能：建立oj的mysql数据库
日期：2020.12.16
版本：0.0.3
*/
create database online_judge;
use online_judge;


CREATE TABLE IF NOT EXISTS problems(
	`id` int unsigned auto_increment,
	`name` TINYTEXT,			-- 问题名称
	`problem_description` json,		-- 包含下面注释的几个模块
	/*
	`description` text，            	-- 问题描述
        `input_description` text,       	-- 输入描述
        `output_description` text,      	-- 输出描述
        `input_sample_data` text,       	-- 输入样例
        `output_sample_data` text,      	-- 输出样例
        `tips` text,                   		-- 提示
	`author`
	*/
	`accept` int unsigned,			-- 通过数
	`fail` int unsigned,			-- 未通过数（计算百分比做数据统计用）
	`judgeer_info` json, 			-- 负责储存判题地址，因为可能是爬虫下来的题目，存储格式为以下注释
	/*
	-- 先判定是否为爬虫，如为爬虫分两步走
	judgger:xxxOJ/local
	oj_problem_id:1000

	language:{				-- 必定存储的语言及判题方式，或着直接爬
		--sample
		"g++":1000			-- 单位（毫秒）
	}					-- 存储语言和时间，到本地判题
	*/
	primary key(`id`)
)ENGINE=INNODB;

-- 没有使用check检查
create table IF NOT EXISTS user(
	`email` char(30),
	`name` char(20),
	`password` char(30) not null,
	`authority` int unsigned,
	primary key(`email`)
)ENGINE=INNODB;

create table IF NOT EXISTS contest(
	`id` int unsigned auto_increment,
	`name` tinytext,
	`start_time` datetime,
	`duration` time,
	`contest_info` json,			-- 负责开关，以及比赛用户过滤，存储注释内的内容
	`is_official_contest` boolean,
	/*
	-- 判断两步走
	-- true存储大量相关信息
		password:			-- 注意加密
		user:[]				-- 存储比赛用户
		organizer:			-- 举办方
		sponsor:			-- 赞助方

	--false不存储任何信息
	*/
	`problem_id` json,			-- 负责存储题目编号方便读取
	primary key(`id`)
)ENGINE=INNODB;

create table IF NOT EXISTS submit_status(
	`submit_id` int unsigned auto_increment,
	`problem_id` int unsigned not null,
	`submit_state` char(50),
	`return_result` text,
	`submit_time` datetime,
	`language` char(50),
	`contest_id` int unsigned,
	`user_email` char(30) not null,
	`run_time` smallint unsigned,
	primary key(`submit_id`)
)ENGINE=INNODB;
