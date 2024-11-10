USE `devops-platform`;

CREATE TABLE IF NOT EXISTS `project` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id，全局唯一',
    `project_name` varchar(128) NOT NULL DEFAULT '' COMMENT '项目名，不可重复',
    `prd_responsible_id` bigint unsigned NOT NULL DEFAULT 0 COMMENT '项目产品负责人，关联user表主键',
    `dev_responsible_id` bigint unsigned NOT NULL DEFAULT 0 COMMENT '项目开发负责人，关联user表主键',
    `test_responsible_id` bigint unsigned NOT NULL DEFAULT 0 COMMENT '项目测试负责人，关联user表主键',
    `sre_responsible_id` bigint unsigned NOT NULL DEFAULT 0 COMMENT '项目sre负责人，关联user表主键',
    `dba_responsible_id` bigint unsigned NOT NULL DEFAULT 0 COMMENT '项目dba负责人，关联user表主键',
    `business_line_id` bigint unsigned NOT NULL DEFAULT 0 COMMENT '项目所属业务线，关联business_line表主键',
    `description` varchar(1024) NOT NULL DEFAULT '' COMMENT '项目描述',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'gorm维护，创建时间',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'gorm维护，更新时间',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'gorm维护，删除时间',
    PRIMARY KEY (`id`),
    KEY `idx_project_name` (`project_name`, `deleted_at`) COMMENT '项目名索引',
    KEY `idx_prd_responsible_id` (`prd_responsible_id`) COMMENT '产品负责人索引',
    KEY `idx_dev_responsible_id` (`dev_responsible_id`) COMMENT '开发负责人索引',
    KEY `idx_test_responsible_id` (`test_responsible_id`) COMMENT '测试负责人索引',
    KEY `idx_sre_responsible_id` (`sre_responsible_id`) COMMENT 'sre负责人索引',
    KEY `idx_dba_responsible_id` (`dba_responsible_id`) COMMENT 'dba负责人索引',
    KEY `idx_business_line_id` (`business_line_id`) COMMENT '业务线关联索引'
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '项目表';