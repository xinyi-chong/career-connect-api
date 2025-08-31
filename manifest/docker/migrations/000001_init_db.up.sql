-- ----------------------------
-- Table structure for account
-- ----------------------------
CREATE TABLE `account` (
  `id` char(36) NOT NULL COMMENT 'account ID',
  `email` varchar(255) NOT NULL COMMENT 'Email',
  `password` char(255) NOT NULL COMMENT 'Password',
  `status` varchar(255) NOT NULL,
  `create_at` datetime DEFAULT NULL COMMENT 'Created Time',
  `update_at` datetime DEFAULT NULL COMMENT 'Updated Time',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for media
-- ----------------------------
CREATE TABLE `media` (
  `id` char(36) NOT NULL,
  `url` varchar(255) NOT NULL,
  `key` varchar(255) NOT NULL,
  `create_at` datetime DEFAULT NULL COMMENT 'Created Time',
  `update_at` datetime DEFAULT NULL COMMENT 'Updated Time',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for user
-- ----------------------------
CREATE TABLE `user` (
  `id` char(36) NOT NULL UNIQUE COMMENT 'user ID',
  `account_id` char(36) NOT NULL UNIQUE,
  `firstname` varchar(255) NOT NULL,
  `lastname` varchar(255) NOT NULL,
  `nationality` varchar(255) NOT NULL,
  `profile_picture_id` char(36) DEFAULT NULL,
  `create_at` datetime DEFAULT NULL COMMENT 'Created Time',
  `update_at` datetime DEFAULT NULL COMMENT 'Updated Time',
  PRIMARY KEY (`id`),
  FOREIGN KEY (`account_id`) REFERENCES `account`(`id`) ON DELETE CASCADE,
  FOREIGN KEY (`profile_picture_id`) REFERENCES `media`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for company
-- ----------------------------
CREATE TABLE `company` (
  `id` char(36) NOT NULL COMMENT 'company ID',
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
  FOREIGN KEY (`account_id`) REFERENCES `account`(`id`) ON DELETE CASCADE,
  FOREIGN KEY (`logo_id`) REFERENCES `media`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for Resume
-- ----------------------------
CREATE TABLE `resume` (
  `id` char(36) NOT NULL,
  `user_id` char(36) NOT NULL,
  `media_id` char(36) NOT NULL,
  `create_at` datetime DEFAULT NULL COMMENT 'Created Time',
  `update_at` datetime DEFAULT NULL COMMENT 'Updated Time',
  PRIMARY KEY (`id`),
  FOREIGN KEY (`user_id`) REFERENCES `user`(`id`) ON DELETE CASCADE,
  FOREIGN KEY (`media_id`) REFERENCES `media`(`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for Certificate
-- ----------------------------
CREATE TABLE `certificate` (
  `id` char(36) NOT NULL COMMENT 'ID',
  `user_id` char(36) NOT NULL COMMENT 'user ID',
  `name` varchar(255) NOT NULL COMMENT 'user email',
  `create_at` datetime DEFAULT NULL COMMENT 'Created Time',
  `update_at` datetime DEFAULT NULL COMMENT 'Updated Time',
  PRIMARY KEY (`id`),
  FOREIGN KEY (`user_id`) REFERENCES `user`(`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for role
-- ----------------------------
CREATE TABLE `role` (
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
CREATE TABLE `company_accounts` (
  `user_id` char(36) NOT NULL,
  `company_id` char(36) NOT NULL,
  `role_id` char(36) NOT NULL,
  `create_at` datetime DEFAULT NULL COMMENT 'Created Time',
  `update_at` datetime DEFAULT NULL COMMENT 'Updated Time',
  PRIMARY KEY (`user_id`,`company_id`),
  FOREIGN KEY (`user_id`) REFERENCES `user`(`id`) ON DELETE CASCADE,
  FOREIGN KEY (`company_id`) REFERENCES `company`(`id`) ON DELETE CASCADE,
  FOREIGN KEY (`role_id`) REFERENCES `role`(`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for company_plan
-- ----------------------------
CREATE TABLE `company_plan` (
  `id` char(36) NOT NULL,
  `name` varchar(255) NOT NULL,
  `create_at` datetime DEFAULT NULL COMMENT 'Created Time',
  `update_at` datetime DEFAULT NULL COMMENT 'Updated Time',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for CompanySubscription
-- ----------------------------
CREATE TABLE `company_subscription` (
  `id` char(36) NOT NULL UNIQUE,
  `company_id` char(36) NOT NULL,
  `company_plan_id` char(36) NOT NULL,
  `status` varchar(255) NOT NULL,
  `expiry` datetime DEFAULT NULL,
  `create_at` datetime DEFAULT NULL COMMENT 'Created Time',
  `update_at` datetime DEFAULT NULL COMMENT 'Updated Time',
  PRIMARY KEY (`id`),
  FOREIGN KEY (`company_id`) REFERENCES `company`(`id`) ON DELETE CASCADE,
  FOREIGN KEY (`company_plan_id`) REFERENCES `company_plan`(`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for Education
-- ----------------------------
CREATE TABLE `education` (
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
  FOREIGN KEY (`institute_id`) REFERENCES `company`(`id`),
  FOREIGN KEY (`user_id`) REFERENCES `user`(`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for Experience
-- ----------------------------
CREATE TABLE `experience` (
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
  FOREIGN KEY (`user_id`) REFERENCES `user`(`id`) ON DELETE CASCADE,
  FOREIGN KEY (`company_id`) REFERENCES `company`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for feature
-- ----------------------------
CREATE TABLE `feature` (
  `id` char(36) NOT NULL UNIQUE,
  `name` varchar(255) NOT NULL,
  `create_at` datetime DEFAULT NULL COMMENT 'Created Time',
  `update_at` datetime DEFAULT NULL COMMENT 'Updated Time',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for Job
-- ----------------------------
CREATE TABLE `job` (
  `id` char(36) NOT NULL,
  `created_by` char(36) NOT NULL,
  `created_by_type` ENUM('user', 'company'),
  `updated_by` char(36) NOT NULL,
  `updated_by_type` ENUM('user', 'company'),
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
  `status` varchar(255) NOT NULL,
  `create_at` datetime DEFAULT NULL COMMENT 'Created Time',
  `update_at` datetime DEFAULT NULL COMMENT 'Updated Time',
  PRIMARY KEY (`id`),
  FOREIGN KEY (`company_id`) REFERENCES `company`(`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for job_question
-- ----------------------------
CREATE TABLE `job_question` (
  `id` char(36) NOT NULL,
  `question` varchar(255) NOT NULL,
  `job_id` char(36) NOT NULL,
  `create_at` datetime DEFAULT NULL COMMENT 'Created Time',
  `update_at` datetime DEFAULT NULL COMMENT 'Updated Time',
  PRIMARY KEY (`id`),
  FOREIGN KEY (`job_id`) REFERENCES `job`(`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for Notification
-- ----------------------------
CREATE TABLE `notification` (
  `id` char(36) NOT NULL,
  `recipient_id` char(36) NOT NULL,
  `redirect` varchar(255) DEFAULT NULL,
  `title` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `status` varchar(255) NOT NULL,
  `create_at` datetime DEFAULT NULL COMMENT 'Created Time',
  `update_at` datetime DEFAULT NULL COMMENT 'Updated Time',
  PRIMARY KEY (`id`),
  FOREIGN KEY (`recipient_id`) REFERENCES `account`(`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for Permission
-- ----------------------------
CREATE TABLE `permission` (
  `id` char(36) NOT NULL UNIQUE,
  `role_id` char(36) NOT NULL,
  `feature_id` char(36) NOT NULL,
  `allow` boolean NOT NULL,
  `create_at` datetime DEFAULT NULL COMMENT 'Created Time',
  `update_at` datetime DEFAULT NULL COMMENT 'Updated Time',
  PRIMARY KEY (`id`),
  FOREIGN KEY (`role_id`) REFERENCES `role`(`id`) ON DELETE CASCADE,
  FOREIGN KEY (`feature_id`) REFERENCES `feature`(`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for activity
-- ----------------------------
CREATE TABLE `activity` (
  `id` char(36) NOT NULL,
  `name` varchar(255) NOT NULL,
  `create_at` datetime DEFAULT NULL COMMENT 'Created Time',
  `update_at` datetime DEFAULT NULL COMMENT 'Updated Time',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for application
-- ----------------------------
CREATE TABLE `application` (
  `id` char(36) NOT NULL,
  `job_id` char(36) NOT NULL,
  `user_id` char(36) NOT NULL,
  `answer` varchar(255) DEFAULT NULL,
  -- `resume_id` char(36) DEFAULT NULL,
  `apply_at` datetime DEFAULT NOW(),
  `activity_id` varchar(255),
  `create_at` datetime DEFAULT NULL COMMENT 'Created Time',
  `update_at` datetime DEFAULT NULL COMMENT 'Updated Time',
  PRIMARY KEY (`id`),
  FOREIGN KEY (`job_id`) REFERENCES `job`(`id`) ON DELETE CASCADE,
  FOREIGN KEY (`user_id`) REFERENCES `user`(`id`) ON DELETE CASCADE,
  FOREIGN KEY (`activity_id`) REFERENCES `activity`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for application_file
-- ----------------------------
CREATE TABLE `application_file` (
  `id` char(36) NOT NULL,
  `application_id` char(36) NOT NULL,
  `media_id` char(36) NOT NULL,
  `file_type` varchar(255) NOT NULL,
  `create_at` datetime DEFAULT NULL COMMENT 'Created Time',
  `update_at` datetime DEFAULT NULL COMMENT 'Updated Time',
  PRIMARY KEY (`id`),
  FOREIGN KEY (`application_id`) REFERENCES `application`(`id`) ON DELETE CASCADE,
  FOREIGN KEY (`media_id`) REFERENCES `media`(`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for ApplicationChatMessage
-- ----------------------------
CREATE TABLE `application_chat_message` (
  `id` char(36) NOT NULL,
  `name` varchar(255) NOT NULL,
  `message` varchar(255) NOT NULL,
  `sender_id` char(36) NOT NULL,
  `application_id`  char(36) NOT NULL,
  `create_at` datetime DEFAULT NULL COMMENT 'Created Time',
  `update_at` datetime DEFAULT NULL COMMENT 'Updated Time',
  PRIMARY KEY (`id`),
  FOREIGN KEY (`sender_id`) REFERENCES `account`(`id`) ON DELETE CASCADE,
  FOREIGN KEY (`application_id`) REFERENCES `application`(`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for Schedule
-- ----------------------------
CREATE TABLE `schedule` (
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
  FOREIGN KEY (`company_id`) REFERENCES `company`(`id`) ON DELETE CASCADE,
  FOREIGN KEY (`user_id`) REFERENCES `user`(`id`) ON DELETE CASCADE,
  FOREIGN KEY (`application_id`) REFERENCES `application`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for Skill
-- ----------------------------
CREATE TABLE `skill` (
  `id` char(36) NOT NULL,
  `user_id` char(36) NOT NULL,
  `name` varchar(255) NOT NULL,
  `create_at` datetime DEFAULT NULL COMMENT 'Created Time',
  `update_at` datetime DEFAULT NULL COMMENT 'Updated Time',
  PRIMARY KEY (`id`),
  FOREIGN KEY (`user_id`) REFERENCES `user`(`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for SuggestSkill
-- ----------------------------
CREATE TABLE `suggest_skill` (
  `id` char(36) NOT NULL UNIQUE,
  `name` varchar(255) NOT NULL,
  `category` varchar(255) NOT NULL,
  `create_at` datetime DEFAULT NULL COMMENT 'Created Time',
  `update_at` datetime DEFAULT NULL COMMENT 'Updated Time',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for user_plan
-- ----------------------------
CREATE TABLE `user_plan` (
  `id` char(36) NOT NULL,
  `name` varchar(255) NOT NULL,
  `create_at` datetime DEFAULT NULL COMMENT 'Created Time',
  `update_at` datetime DEFAULT NULL COMMENT 'Updated Time',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for UserSubscription
-- ----------------------------
CREATE TABLE `user_subscription` (
  `id` char(36) NOT NULL,
  `user_id` char(36) NOT NULL UNIQUE,
  `user_plan_id` char(36) NOT NULL,
  `status` varchar(255) NOT NULL,
  `expiry` datetime DEFAULT NULL,
  `create_at` datetime DEFAULT NULL COMMENT 'Created Time',
  `update_at` datetime DEFAULT NULL COMMENT 'Updated Time',
  PRIMARY KEY (`id`),
  FOREIGN KEY (`user_id`) REFERENCES `user`(`id`) ON DELETE CASCADE,
  FOREIGN KEY (`user_plan_id`) REFERENCES `user_plan`(`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

SET FOREIGN_KEY_CHECKS = 1;
