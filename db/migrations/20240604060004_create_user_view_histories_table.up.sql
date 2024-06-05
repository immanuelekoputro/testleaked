CREATE TABLE user_view_histories
(
    id              INT      NOT NULL AUTO_INCREMENT,
    visitor_user_id INT      NOT NULL,
    host_user_id    INT      NOT NULL,
    visitor_action     ENUM('like','pass'),
    is_superlike    BOOLEAN DEFAULT false,
    created_at TIMESTAMP         NOT NULL DEFAULT now(),
    updated_at TIMESTAMP         NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP NULL,
    PRIMARY KEY (id)
) ENGINE = InnoDB;