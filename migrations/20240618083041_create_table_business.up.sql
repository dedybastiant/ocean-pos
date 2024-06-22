CREATE TABLE business (
    id INT AUTO_INCREMENT PRIMARY KEY,
    owner_user_id INT NOT NULL,
    email VARCHAR(255) NOT NULL,
    phone_number VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    verified_at TIMESTAMP(6) DEFAULT NULL,
    deactivated_at TIMESTAMP(6) DEFAULT NULL,
    created_at TIMESTAMP(6) NOT NULL,
    created_by INT NOT NULL,
    updated_at TIMESTAMP(6) NOT NULL,
    updated_by INT NOT NULL,
    FOREIGN KEY (owner_user_id) REFERENCES user(id)
) ENGINE = InnoDB;