CREATE DATABASE IF NOT EXISTS `golang_beginner` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
USE `golang_beginner`;
CREATE USER 'golang_beginner'@'%' IDENTIFIED BY 'ymVz4U6PBmq51HPP';
GRANT ALL PRIVILEGES ON *.* TO 'golang_beginner'@'%';
GRANT ALL PRIVILEGES ON `golang_beginner`.* TO 'golang_beginner'@'%';
FLUSH PRIVILEGES;