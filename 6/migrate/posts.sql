CREATE TABLE IF NOT EXISTS `post` (
    `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'Primary Key: Unique post ID.',
    `title`  varchar(255) NOT NULL DEFAULT '' COMMENT 'The title of this post.',
    `user_id` int(11) NOT NULL DEFAULT '0' COMMENT 'The users.user_id that owns this post.',
    `body` longtext,
    PRIMARY KEY (`id`),
    KEY `user_id` (`user_id`),
    KEY `title` (`title`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE utf8mb4_general_ci;