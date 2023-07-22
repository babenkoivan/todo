DROP TABLE IF EXISTS `users`;

CREATE TABLE `users`
(
    `id`         int unsigned NOT NULL AUTO_INCREMENT,
    `username`   varchar(255) NOT NULL,
    `password`   varchar(255) NOT NULL,
    `created_at` timestamp    NOT NULL DEFAULT NOW(),
    PRIMARY KEY (`id`),
    UNIQUE (`username`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci;