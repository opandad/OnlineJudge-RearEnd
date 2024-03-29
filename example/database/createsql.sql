-- MySQL Script generated by MySQL Workbench
-- Wed Apr  7 22:33:04 2021
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
  `name` VARCHAR(45) NOT NULL,
  `description` JSON NULL,
  `is_hide_to_user` TINYINT NOT NULL,
  `is_robot_problem` TINYINT NOT NULL,
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
  UNIQUE INDEX `id_UNIQUE` (`id` ASC) VISIBLE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `online_judge`.`contests`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `online_judge`.`contests` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NOT NULL,
  `start_time` DATETIME NOT NULL,
  `end_time` DATETIME NOT NULL,
  `contest_info` JSON NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC) VISIBLE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `online_judge`.`submits`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `online_judge`.`submits` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `submit_state` VARCHAR(45) NOT NULL,
  `submit_time` DATETIME NOT NULL,
  `problem_id` INT UNSIGNED NOT NULL,
  `user_id` INT UNSIGNED NOT NULL,
  `contest_id` INT NULL,
  `language_id` INT NULL,
  `is_error` TINYINT NOT NULL,
  `submit_info` JSON NULL,
  `submit_code` LONGTEXT NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_submit_problems1_idx` (`problem_id` ASC) VISIBLE,
  UNIQUE INDEX `id_UNIQUE` (`id` ASC) VISIBLE,
  INDEX `fk_submit_user1_idx` (`user_id` ASC) VISIBLE,
  CONSTRAINT `fk_submit_problems1`
    FOREIGN KEY (`problem_id`)
    REFERENCES `online_judge`.`problems` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_submit_user1`
    FOREIGN KEY (`user_id`)
    REFERENCES `online_judge`.`users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `online_judge`.`languages`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `online_judge`.`languages` (
  `id` INT UNSIGNED NOT NULL,
  `language` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC) VISIBLE,
  UNIQUE INDEX `language_UNIQUE` (`language` ASC) VISIBLE)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `online_judge`.`users_join_contests`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `online_judge`.`users_join_contests` (
  `user_id` INT UNSIGNED NOT NULL,
  `contest_id` INT UNSIGNED NOT NULL,
  PRIMARY KEY (`user_id`, `contest_id`),
  INDEX `fk_users_join_contests_contests1_idx` (`contest_id` ASC) VISIBLE,
  INDEX `fk_users_join_contests_users1_idx` (`user_id` ASC) VISIBLE,
  CONSTRAINT `fk_users_join_contests_users1`
    FOREIGN KEY (`user_id`)
    REFERENCES `online_judge`.`users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_users_join_contests_contests1`
    FOREIGN KEY (`contest_id`)
    REFERENCES `online_judge`.`contests` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `online_judge`.`contests_has_problems`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `online_judge`.`contests_has_problems` (
  `contest_id` INT UNSIGNED NOT NULL,
  `problem_id` INT UNSIGNED NOT NULL,
  PRIMARY KEY (`contest_id`, `problem_id`),
  INDEX `fk_contests_has_problems_problems1_idx` (`problem_id` ASC) VISIBLE,
  INDEX `fk_contests_has_problems_contests1_idx` (`contest_id` ASC) VISIBLE,
  CONSTRAINT `fk_contests_has_problems_contests1`
    FOREIGN KEY (`contest_id`)
    REFERENCES `online_judge`.`contests` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_contests_has_problems_problems1`
    FOREIGN KEY (`problem_id`)
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
  CONSTRAINT `fk_email_users1`
    FOREIGN KEY (`user_id`)
    REFERENCES `online_judge`.`users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `online_judge`.`contests_support_languages`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `online_judge`.`contests_support_languages` (
  `contest_id` INT UNSIGNED NOT NULL,
  `language_id` INT UNSIGNED NOT NULL,
  PRIMARY KEY (`contest_id`, `language_id`),
  INDEX `fk_contests_support_languages_languages1_idx` (`language_id` ASC) VISIBLE,
  INDEX `fk_contests_support_languages_contests1_idx` (`contest_id` ASC) VISIBLE,
  CONSTRAINT `fk_contests_support_languages_contests1`
    FOREIGN KEY (`contest_id`)
    REFERENCES `online_judge`.`contests` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `fk_contests_support_languages_languages1`
    FOREIGN KEY (`language_id`)
    REFERENCES `online_judge`.`languages` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `online_judge`.`teams`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `online_judge`.`teams` (
  `team` VARCHAR(40) NOT NULL,
  `user_id` INT UNSIGNED NOT NULL,
  PRIMARY KEY (`team`),
  INDEX `fk_teams_users1_idx` (`user_id` ASC) VISIBLE,
  UNIQUE INDEX `team_UNIQUE` (`team` ASC) VISIBLE,
  UNIQUE INDEX `users_id_UNIQUE` (`user_id` ASC) VISIBLE,
  CONSTRAINT `fk_teams_users1`
    FOREIGN KEY (`user_id`)
    REFERENCES `online_judge`.`users` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
