/**
  Create table tb_port as a initial version of the database.
 */

CREATE TABLE `tb_port`
( `id` INT UNSIGNED NOT NULL AUTO_INCREMENT
    , `created_at` DATETIME     NOT NULL DEFAULT NOW()
    , `updated_at` DATETIME     NOT NULL DEFAULT NOW() ON UPDATE NOW()

    , `slug` varchar(5) NOT NULL
    , `name` VARCHAR(40) NOT NULL
    , `city` VARCHAR(40) NOT NULL
    , `country` VARCHAR(80) NOT NULL
    , `alias` VARCHAR(80) NULL
    , `regions` VARCHAR(80) NULL
    , `coordinates` POINT NOT NULL
    , `province` VARCHAR(30) NOT NULL
    , `timezone` VARCHAR(30) NOT NULL
    , `unlocs` VARCHAR(20) NOT NULL
    , `code` VARCHAR(5) NULL
    , PRIMARY KEY (`id`)
    , UNIQUE INDEX `uix_tb_port_code` (`slug` ASC)
);