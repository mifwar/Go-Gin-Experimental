CREATE TABLE discounts(
    `id` INT NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL,
    `quantity` INT NOT NULL,
    `remaining_quantity` INT NOT NULL,
    `type` VARCHAR (255) NOT NULL,
    `start_date` TIMESTAMP NOT NULL,
    `end_date` TIMESTAMP NOT NULL,
    `created_by` INT NULL,
    `updated_by` INT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NULL,
    `deleted_at` TIMESTAMP NULL,
    PRIMARY KEY(`id`),
    INDEX idx_discounts_name(`name`),
    INDEX idx_discounts_type(`type`),
    INDEX idx_discounts_start_date(`start_date`),
    INDEX idx_discounts_end_date(`end_date`),
    INDEX idx_users_created_by(`created_by`),
    INDEX idx_users_updated_by(`updated_by`),
    UNIQUE KEY discounts_name_unique(`name`),
    CONSTRAINT FK_discounts_created_by FOREIGN KEY (`created_by`) REFERENCES admins(`id`) ON DELETE SET NULL,
    CONSTRAINT FK_discounts_updated_by FOREIGN KEY (`updated_by`) REFERENCES admins(`id`) ON DELETE SET NULL
)