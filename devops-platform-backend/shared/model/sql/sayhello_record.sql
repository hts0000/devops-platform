CREATE DATABASE `devops-platform` CHARACTER SET 'utf8mb4' COLLATE 'utf8mb4_general_ci';

USE `devops-platform`;

CREATE TABLE `sayhello_record` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id，全局唯一',
    `name` varchar(128) NOT NULL DEFAULT '' COMMENT '用户名，不可重复',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'gorm维护，创建时间',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'gorm维护，更新时间',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'gorm维护，删除时间',
    PRIMARY KEY (`id`),
    KEY `idx_name` (`name`, `deleted_at`) COMMENT 'User name index'
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'sayhello记录表';