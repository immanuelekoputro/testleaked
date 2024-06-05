CREATE TABLE user_subscribes
(
    id         INT      NOT NULL AUTO_INCREMENT,
    user_id    INT      NOT NULL,
    package_id INT      NOT NULL,
    start_date DATETIME NOT NULL,
    end_date   DATETIME NOT NULL,
    created_at TIMESTAMP         NOT NULL DEFAULT now(),
    updated_at TIMESTAMP         NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP NULL,
    PRIMARY KEY (id)
)ENGINE = InnoDB;