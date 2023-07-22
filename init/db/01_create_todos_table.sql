DROP TABLE IF EXISTS `todos`;

CREATE TABLE `todos`
(
    `id`         int unsigned NOT NULL AUTO_INCREMENT,
    `user_id`    int unsigned NOT NULL,
    `task`       text      NOT NULL,
    `created_at` timestamp NOT NULL DEFAULT NOW(),
    PRIMARY KEY (`id`),
    KEY `idx_user_id` (`user_id`),
    CONSTRAINT `fk_todos_users` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_0900_ai_ci;