CREATE TABLE products(
    `id` INT NOT NULL AUTO_INCREMENT,
    `product_category_id` INT NULL,
    `title` VARCHAR(255) NULL,
    `description` text NULL,
    `image` VARCHAR(255) NULL,
    `video` VARCHAR(255) NULL,
    `created_by` INT NULL,
    `updated_by` INT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NULL,
    `deleted_at` TIMESTAMP NULL,
    PRIMARY KEY(`id`),
    INDEX idx_products_product_category_id(`product_category_id`),
    INDEX idx_products_title(`title`),
    INDEX idx_products_description(`description`),
    CONSTRAINT FK_products_product_category_id FOREIGN KEY (`product_category_id`) REFERENCES product_categories(`id`) ON DELETE SET NULL
)