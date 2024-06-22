CREATE TABLE category (
    id INT AUTO_INCREMENT PRIMARY KEY,
    business_id INT,
    name VARCHAR(255) NOT NULL,
    deactivated_at TIMESTAMP(6) DEFAULT NULL,
    created_at TIMESTAMP(6) NOT NULL,
    created_by INT NOT NULL,
    updated_at TIMESTAMP(6) NOT NULL,
    updated_by INT NOT NULL,
    FOREIGN KEY (business_id) REFERENCES business(id)
) ENGINE = InnoDB;