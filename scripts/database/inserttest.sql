/*
初始化测试表用
*/
INSERT INTO `users` (`id`, `name`, `password`, `authority`, `user_info`) VALUES (1, "abc", "abc", "user", "{}");
INSERT INTO `users` (`id`, `name`, `password`, `authority`, `user_info`) VALUES (2, "json", "json", "admin", '{
    "hello":"helloworld",
    "numb":1000,
    "nihao":"nihao"
}');
INSERT INTO `online_judge`.`emails` (`email`, `user_id`) VALUES ("abc@qq.com", 1);
INSERT INTO `online_judge`.`emails` (`email`, `user_id`) VALUES ("jsontest@qq.com", 2);
INSERT INTO `contests` (`id`, `name`, `start_time`, `end_time`, `contest_info`) VALUES (1, "contesttest", "2020-12-11 14:00:00", "2020-12-11 16:00:00", NULL);
INSERT INTO `languages` (`id`, `language`, `run_cmd`) VALUES (1, "g++", NULL);
INSERT INTO `problems` (`id`, `name`, `description`, `is_hide_to_user`,`is_robot_problem`, `judgeer_info`) VALUES (1, "helloworld", NULL, false, false, NULL);
INSERT INTO `problems` (`id`, `name`, `description`, `is_hide_to_user`,`is_robot_problem`, `judgeer_info`) VALUES (2, "c++", NULL, false, false, NULL);
INSERT INTO `submits` (`id`, `submit_state`, `run_time`, `submit_time`, `problem_id`, `contest_id`, `user_id`, `languages_id`, `submit_info`, `is_error`) VALUES (1, "ac", 1000, "2021-01-08 12:00:00", 1, 1, 1, 1, "{}", false);
INSERT INTO `users_join_contests` (`users_id`, `contests_id`) VALUES (1, 1);
INSERT INTO `contests_has_problems` (`contests_id`, `problems_id`) VALUES (1, 1);
