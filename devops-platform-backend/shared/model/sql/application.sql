USE `devops-platform`;

CREATE TABLE IF NOT EXISTS `application` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id，全局唯一',
    `application_name` varchar(128) NOT NULL DEFAULT '' COMMENT '应用名，不可重复',
    `dev_responsible_id` bigint unsigned NOT NULL DEFAULT 0 COMMENT '应用开发负责人，关联user表主键',
    `test_responsible_id` bigint unsigned NOT NULL DEFAULT 0 COMMENT '应用测试负责人，关联user表主键',
    `project_id` bigint unsigned NOT NULL DEFAULT 0 COMMENT '应用所属项目，关联project表主键',
    `business_line_id` bigint unsigned NOT NULL DEFAULT 0 COMMENT '应用所属业务线，关联business_line表主键',
    `environment` tinyint unsigned NOT NULL DEFAULT 1 COMMENT '应用环境，0-生产环境 1-测试环境 2-预发环境',
    `description` varchar(1024) NOT NULL DEFAULT '' COMMENT '应用描述',
    `git_link` varchar(1024) NOT NULL DEFAULT '' COMMENT '应用关联的git仓库链接',
    `jenkins_link` varchar(1024) NOT NULL DEFAULT '' COMMENT '应用关联的jenkins链接',
    `kubernetes_link` varchar(1024) NOT NULL DEFAULT '' COMMENT '应用关联的k8s链接',
    `log_link` varchar(1024) NOT NULL DEFAULT '' COMMENT '应用关联的日志库链接',
    `config_link` varchar(1024) NOT NULL DEFAULT '' COMMENT '应用关联的配置信息链接',
    `status` tinyint unsigned NOT NULL DEFAULT 1 COMMENT '应用状态，0-生产环境 1-测试环境 2-预发环境',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'gorm维护，创建时间',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'gorm维护，更新时间',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'gorm维护，删除时间',
    PRIMARY KEY (`id`),
    KEY `idx_application_name` (
        `application_name`,
        `deleted_at`
    ) COMMENT '应用名索引',
    KEY `idx_dev_responsible_id` (`dev_responsible_id`) COMMENT '开发负责人索引',
    KEY `idx_test_responsible_id` (`dev_responsible_id`) COMMENT '测试负责人索引',
    KEY `idx_project_id` (`project_id`) COMMENT '项目关联索引',
    KEY `idx_business_line_id` (`business_line_id`) COMMENT '业务线关联索引'
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '应用表';