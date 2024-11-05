USE `devops-platform`;

CREATE TABLE IF NOT EXISTS `business_line` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id，全局唯一',
    `business_line_name` varchar(128) NOT NULL DEFAULT '' COMMENT '业务线名，不可重复',
    `responsible_id` bigint unsigned NOT NULL DEFAULT 0 COMMENT '业务线负责人，关联user表主键',
    `description` varchar(1024) NOT NULL DEFAULT '' COMMENT '业务线描述',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'gorm维护，创建时间',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'gorm维护，更新时间',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'gorm维护，删除时间',
    PRIMARY KEY (`id`),
    KEY `idx_business_line_name` (
        `business_line_name`,
        `deleted_at`
    ) COMMENT '业务线名索引',
    KEY `idx_responsible_id` (`responsible_id`) COMMENT '负责人索引'
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '业务线表';