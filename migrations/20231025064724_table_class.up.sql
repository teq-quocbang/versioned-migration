CREATE TABLE IF NOT EXISTS `class` (
  `id` BIGINT(20) AUTO_INCREMENT,
  `course_id` varchar(50),
  `semester_id` varchar(50),
  `credits` INT,
  `max_slot` INT,
  `current_slot` INT,
  `can_cancel` BOOLEAN,
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  `created_by` INT, 
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_by` INT,

  PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;