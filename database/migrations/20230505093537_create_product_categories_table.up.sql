CREATE TABLE product_categories(
    `id` INT NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL,
    `created_by` INT NULL,
    `updated_by` INT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NULL,
    `deleted_at` TIMESTAMP NULL,
    PRIMARY KEY(`id`),
    UNIQUE KEY product_categories_name_unique(`name`),
    INDEX idx_discounts_name(`name`),
    INDEX idx_users_created_by(`created_by`),
    INDEX idx_users_updated_by(`updated_by`),
    CONSTRAINT FK_product_categories_created_by FOREIGN KEY (`created_by`) REFERENCES admins(`id`) ON DELETE SET NULL,
    CONSTRAINT FK_product_categories_updated_by FOREIGN KEY (`updated_by`) REFERENCES admins(`id`) ON DELETE SET NULL
)