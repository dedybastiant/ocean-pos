CREATE TABLE store_menu (
    id INT AUTO_INCREMENT PRIMARY KEY,
    store_id INT NOT NULL,
    menu_id INT NOT NULL,
    store_price INT NOT NULL,
    is_available BOOLEAN NOT NULL,
    deactivated_at TIMESTAMP(6) DEFAULT NULL,
    created_at TIMESTAMP(6) NOT NULL,
    created_by INT NOT NULL,
    updated_at TIMESTAMP(6) NOT NULL,
    updated_by INT NOT NULL,
    FOREIGN KEY (store_id) REFERENCES store(id),
    FOREIGN KEY (menu_id) REFERENCES menu(id)
) ENGINE = InnoDB;