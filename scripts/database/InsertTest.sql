/*
初始化测试表用
*/
INSERT INTO `users` (`id`, `name`, `password`, `authority`, `user_info`) VALUES (1, "abc", "abc", "user", NULL);
INSERT INTO `users` (`id`, `name`, `password`, `authority`, `user_info`) VALUES (2, "json", "json", "json", '{
    "hello":"helloworld",
    "numb":1000,
    "nihao":"nihao"
}');
INSERT INTO `online_judge`.`emails` (`email`, `user_id`) VALUES ("abc@qq.com", 1);
INSERT INTO `online_judge`.`emails` (`email`, `user_id`) VALUES ("jsontest@qq.com", 2);
INSERT INTO `contests` (`id`, `name`, `start_time`, `duration`, `contest_info`) VALUES (1, "contesttest", "2020-12-11 14:00:00", "05:00", NULL);
INSERT INTO `language` (`id`, `language`, `run_cmd`) VALUES (1, "g++", NULL);
INSERT INTO `problems` (`id`, `name`, `description`, `accept`, `fail`, `is_robot_problem`, `judgeer_info`) VALUES (1, "helloworld", NULL, 0, 0, false, NULL);
INSERT INTO `submits` (`id`, `submit_state`, `run_time`, `submit_time`, `problems_id`, `contest_id`, `user_id`, `language_id`) VALUES (1, "ac", 1000, "2021-01-08 12:00:00", 1, 1, 1, 1);
INSERT INTO `users_join_contests` (`users_id`, `contests_id`) VALUES (1, 1);
INSERT INTO `contests_has_problems` (`contests_id`, `problems_id`) VALUES (1, 1);