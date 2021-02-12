CREATE TABLE IF NOT EXISTS `comment` (
    `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'Primary Key: Unique post ID.',
    `post_id` int(11) NOT NULL DEFAULT '0' COMMENT 'The post.post_id that owns this comment.',
    `name` varchar(255) NOT NULL DEFAULT '' COMMENT 'The name of this comment.',
    `email` varchar(255) NOT NULL DEFAULT '' COMMENT 'The email author comment of this post.',
    `body` longtext,
    PRIMARY KEY (`id`),
    KEY `post_id` (`post_id`),
    KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE utf8mb4_general_ci;
