CREATE TABLE users
(
    id               INT          NOT NULL AUTO_INCREMENT,
    name             VARCHAR(255) NOT NULL,
    email            VARCHAR(255) NOT NULL,
    password         VARCHAR(255) NOT NULL,
    gender           enum('male', 'female') NOT NULL,
    date_of_birthday TIMESTAMP         NOT NULL DEFAULT now(),
    created_at TIMESTAMP         NOT NULL DEFAULT now(),
    updated_at TIMESTAMP         NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP NULL,
    PRIMARY KEY (id)
)ENGINE = InnoDB;