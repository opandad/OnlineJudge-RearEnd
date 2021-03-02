-- MySQL Script generated by MySQL Workbench
-- Thu Feb 25 21:30:20 2021
-- Model: New Model    Version: 1.0
-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema online_judge
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema online_judge
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `online_judge` DEFAULT CHARACTER SET utf8mb4 ;
USE `online_judge` ;

-- -----------------------------------------------------
-- Table `online_judge`.`problems`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `online_judge`.`problems` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NULL,
  `description` JSON NULL,
  `is_hide_to_user` TINYINT NULL,
  `is_robot_problem` TINYINT NULL,
  `judgeer_info` JSON NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC) VISIBLE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `online_judge`.`users`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `online_judge`.`users` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(10) NOT NULL,
  `password` TEXT NOT NULL,
  `authority` VARCHAR(20) NOT NULL,
  `user_info` JSON NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC) VISIBLE,
  UNIQUE INDEX `name_UNIQUE` (`name` ASC) VISIBLE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `online_judge`.`contests`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `online_judge`.`contests` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NULL,
  `start_time` DATETIME NULL,
  `end_time` DATETIME NULL,
  `contest_info` JSON NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC) VISIBLE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `online_judge`.`language`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `online_judge`.`language` (
  `id` INT NOT NULL,
  `language` VARCHAR(45) NOT NULL,
  `run_cmd` VARCHAR(45) NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC) VISIBLE,
  UNIQUE INDEX `language_UNIQUE` (`language` ASC) VISIBLE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `online_judge`.`submits`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `online_judge`.`submits` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `submit_state` VARCHAR(45) NULL,
  `run_time` INT UNSIGNED NULL,
  `submit_time` DATETIME NULL,
  `problem_id` INT UNSIGNED NOT NULL,
  `contest_id` INT UNSIGNED NOT NULL,
  `user_id` INT UNSIGNED NOT NULL,
  `language_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_submit_problems1_idx` (`problem_id` ASC) VISIBLE,
  UNIQUE INDEX `id_UNIQUE` (`id` ASC) VISIBLE,
  INDEX `fk_submit_contest1_idx` (`contest_id` ASC) VISIBLE,
  INDEX `fk_submit_user1_idx` (`user_id` ASC) VISIBLE,
  INDEX `fk_submit_language1_idx` (`language_id` ASC) VISIBLE,
  CONSTRAINT `fk_submit_problems1`
    FOREIGN KEY (`problem_id`)
    REFERENCES `online_judge`.`problems` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_submit_contest1`
    FOREIGN KEY (`contest_id`)
    REFERENCES `online_judge`.`contests` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_submit_user1`
    FOREIGN KEY (`user_id`)
    REFERENCES `online_judge`.`users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_submit_language1`
    FOREIGN KEY (`language_id`)
    REFERENCES `online_judge`.`language` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `online_judge`.`users_join_contests`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `online_judge`.`users_join_contests` (
  `users_id` INT UNSIGNED NOT NULL,
  `contests_id` INT UNSIGNED NOT NULL,
  PRIMARY KEY (`users_id`, `contests_id`),
  INDEX `fk_users_join_contests_contests1_idx` (`contests_id` ASC) VISIBLE,
  INDEX `fk_users_join_contests_users1_idx` (`users_id` ASC) VISIBLE,
  CONSTRAINT `fk_users_has_contests_users1`
    FOREIGN KEY (`users_id`)
    REFERENCES `online_judge`.`users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_users_has_contests_contests1`
    FOREIGN KEY (`contests_id`)
    REFERENCES `online_judge`.`contests` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `online_judge`.`contests_has_problems`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `online_judge`.`contests_has_problems` (
  `contests_id` INT UNSIGNED NOT NULL,
  `problems_id` INT UNSIGNED NOT NULL,
  PRIMARY KEY (`contests_id`, `problems_id`),
  INDEX `fk_contests_has_problems_problems1_idx` (`problems_id` ASC) VISIBLE,
  INDEX `fk_contests_has_problems_contests1_idx` (`contests_id` ASC) VISIBLE,
  CONSTRAINT `fk_contests_has_problems_contests1`
    FOREIGN KEY (`contests_id`)
    REFERENCES `online_judge`.`contests` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_contests_has_problems_problems1`
    FOREIGN KEY (`problems_id`)
    REFERENCES `online_judge`.`problems` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `online_judge`.`emails`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `online_judge`.`emails` (
  `email` VARCHAR(40) NOT NULL,
  `user_id` INT UNSIGNED NOT NULL,
  PRIMARY KEY (`email`),
  INDEX `fk_email_user1_idx` (`user_id` ASC) VISIBLE,
  UNIQUE INDEX `email_UNIQUE` (`email` ASC) VISIBLE,
  UNIQUE INDEX `user_id_UNIQUE` (`user_id` ASC) VISIBLE,
  CONSTRAINT `fk_email_users1`
    FOREIGN KEY (`user_id`)
    REFERENCES `online_judge`.`users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;

CREATE USER 'online_judge_admin' IDENTIFIED BY 'qweasd';

GRANT ALL ON `online_judge`.* TO 'online_judge_admin';
CREATE USER 'online_judge_user' IDENTIFIED BY 'qweasd';

GRANT SELECT ON TABLE `online_judge`.* TO 'online_judge_user';
GRANT SELECT, INSERT, TRIGGER ON TABLE `online_judge`.* TO 'online_judge_user';

SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
