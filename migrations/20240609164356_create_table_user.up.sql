CREATE TABLE user (
    id INT AUTO_INCREMENT PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    phone_number VARCHAR(255) NOT NULL,
    is_email_verified BOOLEAN DEFAULT FALSE,
    email_verified_at TIMESTAMP(6) DEFAULT NULL,
    is_phone_number_verified BOOLEAN DEFAULT FALSE,
    phone_number_verified_at TIMESTAMP(6) DEFAULT NULL,
    deactivated_at TIMESTAMP(6) DEFAULT NULL,
    last_login TIMESTAMP(6) DEFAULT NULL,
    created_at TIMESTAMP(6) NOT NULL,
    created_by INT NOT NULL,
    updated_at TIMESTAMP(6) NOT NULL,
    updated_by INT NOT NULL
) ENGINE = InnoDB;