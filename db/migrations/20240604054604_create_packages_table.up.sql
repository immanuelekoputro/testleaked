CREATE TABLE packages
(
    id                    INT          NOT NULL AUTO_INCREMENT,
    package_name          VARCHAR(255) NOT NULL,
    package_price         INT          NOT NULL,
    package_duration_days INT          NOT NULL,
    status                BOOLEAN DEFAULT false,
    created_at TIMESTAMP         NOT NULL DEFAULT now(),
    updated_at TIMESTAMP         NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP NULL,
    PRIMARY KEY (id)
)ENGINE = InnoDB;