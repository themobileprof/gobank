CREATE TABLE accounts (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `user_id` INT NOT NULL,
    `account_number` VARCHAR(10) NOT NULL,
    `balance` DECIMAL(10, 2) DEFAULT 0.00,
    `status` ENUM('active', 'inactive', 'blacklisted') DEFAULT 'active',
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT FK_UserAccount FOREIGN KEY (`user_id`) REFERENCES users(`id`)
    ON DELETE CASCADE  
    ON UPDATE CASCADE  
);