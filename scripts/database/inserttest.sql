/*
初始化测试表用
*/
INSERT INTO `users` (`id`, `name`, `password`, `authority`, `user_info`) VALUES (1, "admin", "admin", "admin", "{}");
INSERT INTO `users` (`id`, `name`, `password`, `authority`, `user_info`) VALUES (2, "user", "user", "user", "{}");
INSERT INTO `online_judge`.`emails` (`email`, `user_id`) VALUES ("admin@qq.com", 1);
INSERT INTO `online_judge`.`emails` (`email`, `user_id`) VALUES ("user@qq.com", 2);
INSERT INTO `problems` (`name`,`description`,`is_hide_to_user`,`is_robot_problem`,`judgeer_info`) VALUES ('a+b','{"problemDescription":"如题","inputDescription":"每一次输入包含两个数","outputDescription":"输出一个数","inputCase":"1 2","outputCase":"3","tips":"","timeLimit":1000,"memoryLimit":1000,"realTimeLimit":1000,"fileSizeLimit":1000}',false,false,'{"problemPath":"problems/APlusB/problem.json"}');
INSERT INTO `contests` (`id`, `name`, `start_time`, `end_time`, `contest_info`) VALUES (1, "练习专用", "2020-12-11 14:00:00", "2022-12-11 16:00:00", "{}");
INSERT INTO `languages` (`id`, `language`) VALUES (1, "gcc");
INSERT INTO `languages` (`id`, `language`) VALUES (2, "g++");
INSERT INTO `languages` (`id`, `language`) VALUES (3, "java");
INSERT INTO `languages` (`id`, `language`) VALUES (4, "python2");
INSERT INTO `languages` (`id`, `language`) VALUES (5, "python3");
INSERT INTO `languages` (`id`, `language`) VALUES (6, "golang");
INSERT INTO `online_judge`.`contests_support_languages` (`contest_id`, `language_id`) VALUES (1, 1);
INSERT INTO `users_join_contests` (`user_id`, `contest_id`) VALUES (1, 1);
INSERT INTO `contests_has_problems` (`contest_id`, `problem_id`) VALUES (1, 1);
INSERT INTO `submits` (`id`, `submit_state`, `submit_time`, `problem_id`, `contest_id`,`language_id`,  `user_id`, `submit_info`, `is_error`, `submit_code`) VALUES (1, "Accept", "2021-01-08 12:00:00", 1, 1, 1, 1, "{}", false, "");
