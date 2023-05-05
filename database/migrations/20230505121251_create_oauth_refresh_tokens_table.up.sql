CREATE TABLE oauth_refresh_tokens(
    `id` INT NOT NULL AUTO_INCREMENT,
    `oauth_access_token_id` INT NULL,
    `user_id` INT NULL,
    `token` VARCHAR(255) NULL,
    `expired_at` TIMESTAMP NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NULL,
    `deleted_at` TIMESTAMP NULL,
    PRIMARY KEY(`id`),
    INDEX idx_oauth_refresh_tokens_oauth_access_token_id(`oauth_access_token_id`),
    INDEX idx_oauth_refresh_tokens_user_id(`user_id`),
    CONSTRAINT FK_oauth_refresh_tokens_oauth_access_token_id FOREIGN KEY (`oauth_access_token_id`) REFERENCES oauth_access_tokens(`id`) ON DELETE SET NULL
)
