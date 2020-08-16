DROP DATABASE IF EXISTS test_database;
CREATE DATABASE test_database;
use test_database;

CREATE TABLE IF NOT EXISTS `products`
(
    `id`     int(11) NOT NULL AUTO_INCREMENT,
    `name`   varchar(500),
    `status` tinyint(1),
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8
  AUTO_INCREMENT = 1;

CREATE TABLE `test_table`
(
    `id`   int(11),
    `name` varchar(500)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;
