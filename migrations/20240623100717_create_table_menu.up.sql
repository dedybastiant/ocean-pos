CREATE TABLE menu (
    id INT AUTO_INCREMENT PRIMARY KEY,
    category_id INT NOT NULL,
    name VARCHAR(255) NOT NULL,
    default_price INT NOT NULL,
    deactivated_at TIMESTAMP(6) DEFAULT NULL,
    created_at TIMESTAMP(6) NOT NULL,
    created_by INT NOT NULL,
    updated_at TIMESTAMP(6) NOT NULL,
    updated_by INT NOT NULL,
    FOREIGN KEY (category_id) REFERENCES category(id)
) ENGINE = InnoDB;