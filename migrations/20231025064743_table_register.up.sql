CREATE TABLE IF NOT EXISTS `register` (
  `account_id` BIGINT(20),
  `course_id` VARCHAR(50),
  `semester_id` VARCHAR(50),
  `class_id` VARCHAR(50),
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  `created_by` INT,
  `is_canceled` BOOLEAN DEFAULT false

) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;