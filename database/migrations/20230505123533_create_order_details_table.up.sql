CREATE TABLE order_details(
    `id` INT NOT NULL AUTO_INCREMENT,
    `product_id` INT NULL,
    `price` INT NULL,
    `order_id` INT NULL,
    `created_by` INT NULL,
    `updated_by` INT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NULL,
    `deleted_at` TIMESTAMP NULL,
    PRIMARY KEY(`id`),
    INDEX idx_order_details_product_id(`product_id`),
    INDEX idx_order_details_order_id(`order_id`),
    CONSTRAINT FK_order_details_product_id FOREIGN KEY (`product_id`) REFERENCES products(`id`) ON DELETE SET NULL,
    CONSTRAINT FK_order_details_order_id FOREIGN KEY (`order_id`) REFERENCES orders(`id`) ON DELETE SET NULL
)