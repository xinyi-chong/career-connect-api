-- ----------------------------
-- Table structure for Account
-- ----------------------------
CREATE TABLE `Account` (
  `id` char(36) NOT NULL COMMENT 'Account ID',
  `email` varchar(255) NOT NULL COMMENT 'Email',
  `password` char(255) NOT NULL COMMENT 'Password',
  `status` varchar(255) NOT NULL,
  `create_at` datetime DEFAULT NULL COMMENT 'Created Time',
  `update_at` datetime DEFAULT NULL COMMENT 'Updated Time',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for Media
-- ----------------------------
CREATE TABLE `Media` (
  `id` char(36) NOT NULL,
  `url` varchar(255) NOT NULL,
  `key` varchar(255) NOT NULL,
  `create_at` datetime DEFAULT NULL COMMENT 'Created Time',
  `update_at` datetime DEFAULT NULL COMMENT 'Updated Time',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for User
-- ----------------------------
CREATE TABLE `User` (
  `id` char(36) NOT NULL UNIQUE COMMENT 'User ID',
  `account_id` char(36) NOT NULL UNIQUE,
  `firstname` varchar(255) NOT NULL,
  `lastname` varchar(255) NOT NULL,
  `nationality` varchar(255) NOT NULL,
  `profile_picture_id` char(36) DEFAULT NULL,
  `create_at` datetime DEFAULT NULL COMMENT 'Created Time',
  `update_at` datetime DEFAULT NULL COMMENT 'Updated Time',
  PRIMARY KEY (`id`),
  FOREIGN KEY (`account_id`) REFERENCES `Account`(`id`) ON DELETE CASCADE,
  FOREIGN KEY (`profile_picture_id`) REFERENCES `Media`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for Company
-- ----------------------------
CREATE TABLE `Company` (
  `id` char(36) NOT NULL COMMENT 'Company ID',
  `account_id` char(36) NOT NULL UNIQUE,
  `name` varchar(255) NOT NULL,
  `description` varchar(255) DEFAULT NULL,
  `industry` varchar(255) DEFAULT NULL,
  `tag` varchar(255) DEFAULT NULL,
  `address` varchar(255) DEFAULT NULL,
  `website` varchar(255) DEFAULT NULL,
  `city` varchar(255) DEFAULT NULL,
  `size` varchar(255) DEFAULT NULL,
  `contact` varchar(255) DEFAULT NULL,
  `logo_id` char(36) DEFAULT NULL,
  `create_at` datetime DEFAULT NULL COMMENT 'Created Time',
  `update_at` datetime DEFAULT NULL COMMENT 'Updated Time',
  PRIMARY KEY (`id`),
  FOREIGN KEY (`account_id`) REFERENCES `Account`(`id`) ON DELETE CASCADE,
  FOREIGN KEY (`logo_id`) REFERENCES `Media`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for Resume
-- ----------------------------
CREATE TABLE `Resume` (
  `id` char(36) NOT NULL,
  `user_id` char(36) NOT NULL,
  `media_id` char(36) NOT NULL,
  `create_at` datetime DEFAULT NULL COMMENT 'Created Time',
  `update_at` datetime DEFAULT NULL COMMENT 'Updated Time',
  PRIMARY KEY (`id`),
  FOREIGN KEY (`user_id`) REFERENCES `User`(`id`) ON DELETE CASCADE,
  FOREIGN KEY (`media_id`) REFERENCES `Media`(`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for Certificate
-- ----------------------------
CREATE TABLE `Certificate` (
  `id` char(36) NOT NULL COMMENT 'ID',
  `user_id` char(36) NOT NULL COMMENT 'User ID',
  `name` varchar(255) NOT NULL COMMENT 'User email',
  `create_at` datetime DEFAULT NULL COMMENT 'Created Time',
  `update_at` datetime DEFAULT NULL COMMENT 'Updated Time',
  PRIMARY KEY (`id`),
  FOREIGN KEY (`user_id`) REFERENCES `User`(`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for Role
-- ----------------------------
CREATE TABLE `Role` (
  `id` char(36) NOT NULL UNIQUE,
  `name` varchar(255) NOT NULL,
  `status` varchar(255) NOT NULL,
  `create_at` datetime DEFAULT NULL COMMENT 'Created Time',
  `update_at` datetime DEFAULT NULL COMMENT 'Updated Time',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for CompanyAccounts
-- ----------------------------
CREATE TABLE `CompanyAccounts` (
  `user_id` char(36) NOT NULL,
  `company_id` char(36) NOT NULL,
  `role_id` char(36) NOT NULL,
  `create_at` datetime DEFAULT NULL COMMENT 'Created Time',
  `update_at` datetime DEFAULT NULL COMMENT 'Updated Time',
  PRIMARY KEY (`user_id`,`company_id`),
  FOREIGN KEY (`user_id`) REFERENCES `User`(`id`) ON DELETE CASCADE,
  FOREIGN KEY (`company_id`) REFERENCES `Company`(`id`) ON DELETE CASCADE,
  FOREIGN KEY (`role_id`) REFERENCES `Role`(`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for CompanyPlan
-- ----------------------------
CREATE TABLE `CompanyPlan` (
  `id` char(36) NOT NULL,
  `name` varchar(255) NOT NULL,
  `create_at` datetime DEFAULT NULL COMMENT 'Created Time',
  `update_at` datetime DEFAULT NULL COMMENT 'Updated Time',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for CompanySubscription
-- ----------------------------
CREATE TABLE `CompanySubscription` (
  `id` char(36) NOT NULL UNIQUE,
  `company_id` char(36) NOT NULL,
  `company_plan_id` char(36) NOT NULL,
  `status` varchar(255) NOT NULL,
  `expiry` datetime DEFAULT NULL,
  `create_at` datetime DEFAULT NULL COMMENT 'Created Time',
  `update_at` datetime DEFAULT NULL COMMENT 'Updated Time',
  PRIMARY KEY (`id`),
  FOREIGN KEY (`company_id`) REFERENCES `Company`(`id`) ON DELETE CASCADE,
  FOREIGN KEY (`company_plan_id`) REFERENCES `CompanyPlan`(`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for Education
-- ----------------------------
CREATE TABLE `Education` (
  `id` char(36) NOT NULL UNIQUE,
  `user_id` char(36) NOT NULL,
  `start_date` datetime NOT NULL,
  `end_date` datetime DEFAULT NULL,
  `institute_id` char(36) DEFAULT NULL,
  `institute_string` varchar(255) DEFAULT NULL,
  `level` varchar(255) DEFAULT NULL,
  `programme` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `create_at` datetime DEFAULT NULL COMMENT 'Created Time',
  `update_at` datetime DEFAULT NULL COMMENT 'Updated Time',
  PRIMARY KEY (`id`),
  FOREIGN KEY (`institute_id`) REFERENCES `Company`(`id`),
  FOREIGN KEY (`user_id`) REFERENCES `User`(`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for Experience
-- ----------------------------
CREATE TABLE `Experience` (
  `id` char(36) NOT NULL UNIQUE,
  `user_id` char(36),
  `start_date` datetime NOT NULL,
  `end_date` datetime DEFAULT NULL,
  `is_present` boolean NOT NULL,
  `description` varchar(255) DEFAULT NULL,
  `title` varchar(255) NOT NULL,
  `company_id` char(36) DEFAULT NULL,
  `company_string` varchar(255) DEFAULT NULL,
  `create_at` datetime DEFAULT NULL COMMENT 'Created Time',
  `update_at` datetime DEFAULT NULL COMMENT 'Updated Time',
  PRIMARY KEY (`id`),
  FOREIGN KEY (`user_id`) REFERENCES `User`(`id`) ON DELETE CASCADE,
  FOREIGN KEY (`company_id`) REFERENCES `Company`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for Feature
-- ----------------------------
CREATE TABLE `Feature` (
  `id` char(36) NOT NULL UNIQUE,
  `name` varchar(255) NOT NULL,
  `create_at` datetime DEFAULT NULL COMMENT 'Created Time',
  `update_at` datetime DEFAULT NULL COMMENT 'Updated Time',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for JobQuestion
-- ----------------------------
CREATE TABLE `JobQuestion` (
  `id` char(36) NOT NULL,
  `question` varchar(255) NOT NULL,
  `create_at` datetime DEFAULT NULL COMMENT 'Created Time',
  `update_at` datetime DEFAULT NULL COMMENT 'Updated Time',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for Job
-- ----------------------------
CREATE TABLE `Job` (
  `id` char(36) NOT NULL,
  `created_by` char(36) NOT NULL,
  `updated_by` char(36) NOT NULL,
  `title` varchar(255) NOT NULL,
  `company_id` char(36) NOT NULL,
  `tag` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `level` varchar(255) DEFAULT NULL,
  `salary` varchar(255) DEFAULT NULL,
  `posted_at` datetime DEFAULT NOW(),
  `location` varchar(255) DEFAULT NULL,
  `is_remote` boolean DEFAULT 0,
  `is_hybrid` boolean DEFAULT 0,
  `expiry` datetime DEFAULT NULL,
  `job_question_id` char(36) DEFAULT NULL,
  `status` varchar(255) NOT NULL,
  `create_at` datetime DEFAULT NULL COMMENT 'Created Time',
  `update_at` datetime DEFAULT NULL COMMENT 'Updated Time',
  PRIMARY KEY (`id`),
  FOREIGN KEY (`created_by`) REFERENCES `User`(`id`),
  FOREIGN KEY (`created_by`) REFERENCES `Company`(`id`),
  FOREIGN KEY (`updated_by`) REFERENCES `User`(`id`),
  FOREIGN KEY (`updated_by`) REFERENCES `Company`(`id`),
  FOREIGN KEY (`company_id`) REFERENCES `Company`(`id`) ON DELETE CASCADE,
  FOREIGN KEY (`job_question_id`) REFERENCES `JobQuestion`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for Notification
-- ----------------------------
CREATE TABLE `Notification` (
  `id` char(36) NOT NULL,
  `recipient_id` char(36) NOT NULL,
  `redirect` varchar(255) DEFAULT NULL,
  `title` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `status` varchar(255) NOT NULL,
  `create_at` datetime DEFAULT NULL COMMENT 'Created Time',
  `update_at` datetime DEFAULT NULL COMMENT 'Updated Time',
  PRIMARY KEY (`id`),
  FOREIGN KEY (`recipient_id`) REFERENCES `User`(`id`) ON DELETE CASCADE,
  FOREIGN KEY (`recipient_id`) REFERENCES `Company`(`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for Permission
-- ----------------------------
CREATE TABLE `Permission` (
  `id` char(36) NOT NULL UNIQUE,
  `role_id` char(36) NOT NULL,
  `feature_id` char(36) NOT NULL,
  `allow` boolean NOT NULL,
  `create_at` datetime DEFAULT NULL COMMENT 'Created Time',
  `update_at` datetime DEFAULT NULL COMMENT 'Updated Time',
  PRIMARY KEY (`id`),
  FOREIGN KEY (`role_id`) REFERENCES `Role`(`id`) ON DELETE CASCADE,
  FOREIGN KEY (`feature_id`) REFERENCES `Feature`(`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for Activity
-- ----------------------------
CREATE TABLE `Activity` (
  `id` char(36) NOT NULL,
  `name` varchar(255) NOT NULL,
  `create_at` datetime DEFAULT NULL COMMENT 'Created Time',
  `update_at` datetime DEFAULT NULL COMMENT 'Updated Time',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for Application
-- ----------------------------
CREATE TABLE `Application` (
  `id` char(36) NOT NULL,
  `job_id` char(36) NOT NULL,
  `user_id` char(36) NOT NULL,
  `answer` varchar(255) DEFAULT NULL,
  `resume_id` char(36) DEFAULT NULL,
  `apply_at` datetime DEFAULT NOW(),
  `activity_id` varchar(255),
  `create_at` datetime DEFAULT NULL COMMENT 'Created Time',
  `update_at` datetime DEFAULT NULL COMMENT 'Updated Time',
  PRIMARY KEY (`id`),
  FOREIGN KEY (`job_id`) REFERENCES `Job`(`id`) ON DELETE CASCADE,
  FOREIGN KEY (`user_id`) REFERENCES `User`(`id`) ON DELETE CASCADE,
  FOREIGN KEY (`activity_id`) REFERENCES `Activity`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for ApplicationChatMessage
-- ----------------------------
CREATE TABLE `ApplicationChatMessage` (
  `id` char(36) NOT NULL,
  `name` varchar(255) NOT NULL,
  `message` varchar(255) NOT NULL,
  `sender_id` char(36) NOT NULL,
  `application_id`  char(36) NOT NULL,
  `create_at` datetime DEFAULT NULL COMMENT 'Created Time',
  `update_at` datetime DEFAULT NULL COMMENT 'Updated Time',
  PRIMARY KEY (`id`),
  FOREIGN KEY (`sender_id`) REFERENCES `User`(`id`) ON DELETE CASCADE,
  FOREIGN KEY (`sender_id`) REFERENCES `Company`(`id`) ON DELETE CASCADE,
  FOREIGN KEY (`application_id`) REFERENCES `Application`(`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for Schedule
-- ----------------------------
CREATE TABLE `Schedule` (
  `id` char(36) NOT NULL,
  `application_id` char(36) DEFAULT NULL,
  `start_time` datetime NOT NULL,
  `end_time` datetime NOT NULL,
  `title` varchar(255) NOT NULL,
  `company_id` char(36) NOT NULL,
  `user_id` char(36) NOT NULL,
  `location` varchar(255) NOT NULL,
  `link` varchar(255) DEFAULT NULL,
  `status` varchar(255) NOT NULL,
  `create_at` datetime DEFAULT NULL COMMENT 'Created Time',
  `update_at` datetime DEFAULT NULL COMMENT 'Updated Time',
  PRIMARY KEY (`id`),
  FOREIGN KEY (`company_id`) REFERENCES `Company`(`id`) ON DELETE CASCADE,
  FOREIGN KEY (`user_id`) REFERENCES `User`(`id`) ON DELETE CASCADE,
  FOREIGN KEY (`application_id`) REFERENCES `Application`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for Skill
-- ----------------------------
CREATE TABLE `Skill` (
  `id` char(36) NOT NULL,
  `user_id` char(36) NOT NULL,
  `name` varchar(255) NOT NULL,
  `create_at` datetime DEFAULT NULL COMMENT 'Created Time',
  `update_at` datetime DEFAULT NULL COMMENT 'Updated Time',
  PRIMARY KEY (`id`),
  FOREIGN KEY (`user_id`) REFERENCES `User`(`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for SuggestSkill
-- ----------------------------
CREATE TABLE `SuggestSkill` (
  `id` char(36) NOT NULL UNIQUE,
  `name` varchar(255) NOT NULL,
  `category` varchar(255) NOT NULL,
  `create_at` datetime DEFAULT NULL COMMENT 'Created Time',
  `update_at` datetime DEFAULT NULL COMMENT 'Updated Time',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for UserPlan
-- ----------------------------
CREATE TABLE `UserPlan` (
  `id` char(36) NOT NULL,
  `name` varchar(255) NOT NULL,
  `create_at` datetime DEFAULT NULL COMMENT 'Created Time',
  `update_at` datetime DEFAULT NULL COMMENT 'Updated Time',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for UserSubscription
-- ----------------------------
CREATE TABLE `UserSubscription` (
  `id` char(36) NOT NULL,
  `user_id` char(36) NOT NULL UNIQUE,
  `user_plan_id` char(36) NOT NULL,
  `status` varchar(255) NOT NULL,
  `expiry` datetime DEFAULT NULL,
  `create_at` datetime DEFAULT NULL COMMENT 'Created Time',
  `update_at` datetime DEFAULT NULL COMMENT 'Updated Time',
  PRIMARY KEY (`id`),
  FOREIGN KEY (`user_id`) REFERENCES `User`(`id`) ON DELETE CASCADE,
  FOREIGN KEY (`user_plan_id`) REFERENCES `UserPlan`(`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

SET FOREIGN_KEY_CHECKS = 1;
