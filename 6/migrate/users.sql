CREATE TABLE IF NOT EXISTS `user` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `name` varchar(255) NOT NULL DEFAULT '',
    `username` varchar(255) NOT NULL DEFAULT '',
    `email` varchar(255) NOT NULL DEFAULT '',
    `phone` varchar(255) NOT NULL DEFAULT '',
    `website` varchar(255) NOT NULL DEFAULT '',
    PRIMARY KEY (`id`),
    KEY `name` (`name`),
    KEY `username` (`username`),
    KEY `email` (`email`),
    KEY `phone` (`phone`),
    KEY `website` (`website`)
) ENGINE = InnoDB CHARSET=utf8mb4 COLLATE utf8mb4_general_ci;