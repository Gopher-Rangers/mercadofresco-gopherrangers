SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema mercado-fresco
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema mercado-fresco
-- -----------------------------------------------------
-- DROP DATABASE `mercado-fresco`;
CREATE SCHEMA IF NOT EXISTS `mercado-fresco` DEFAULT CHARACTER SET utf8;
USE `mercado-fresco` ;

-- -----------------------------------------------------
-- Table `mercado-fresco`.`buyers`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mercado-fresco`.`buyers` (
                                                         `id` SERIAL,
                                                         `card_number_id` VARCHAR(45) NOT NULL,
                                                         `first_name` VARCHAR(45) NOT NULL,
                                                         `last_name` VARCHAR(45) NOT NULL,
                                                         PRIMARY KEY (`id`))
    ENGINE = InnoDB;

-- -----------------------------------------------------
-- Table `mercado-fresco`.`warehouse`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mercado-fresco`.`warehouse` (
                                                            `id` SERIAL,
                                                            `warehouse_code` VARCHAR(20) NOT NULL,
                                                            `address` VARCHAR(80) NOT NULL,
                                                            `telephone` VARCHAR(15) NOT NULL,
                                                            `locality_id` BIGINT UNSIGNED,
                                                            PRIMARY KEY (`id`))
    ENGINE = InnoDB;

-- -----------------------------------------------------
-- Table `mercado-fresco`.`employees`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mercado-fresco`.`employees` (
                                                            `id` SERIAL,
                                                            `card_number_id` VARCHAR(45) NOT NULL,
                                                            `first_name` VARCHAR(45) NOT NULL,
                                                            `last_name` VARCHAR(45) NOT NULL,
                                                            `warehouse_id` BIGINT UNSIGNED,
                                                            PRIMARY KEY (`id`))
    ENGINE = InnoDB;

-- -----------------------------------------------------
-- Table `mercado-fresco`.`sellers`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mercado-fresco`.`sellers` (
                                                          `id` SERIAL,
                                                          `cid` INT(11) UNSIGNED NOT NULL,
                                                          `company_name` VARCHAR(80) NOT NULL,
                                                          `address` VARCHAR(80) NOT NULL,
                                                          `telephone` VARCHAR(15) NOT NULL,
                                                          `locality_id` BIGINT UNSIGNED,
                                                          PRIMARY KEY (`id`))
    ENGINE = InnoDB;

-- -----------------------------------------------------
-- Table `mercado-fresco`.`products`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mercado-fresco`.`products` (
                                                            `id` SERIAL,
                                                            `product_code` VARCHAR(15) NULL,
                                                            `description` VARCHAR(45) NULL,
                                                            `width` DECIMAL(2) NULL,
                                                            `height` DECIMAL(2) NULL,
                                                            `length` DECIMAL(2) NULL,
                                                            `net_weight` DECIMAL(2) NULL,
                                                            `expiration_rate` VARCHAR(15) NULL,
                                                            `recommended_freezing_temperature` DECIMAL(2) NULL,
                                                            `freezing_rate` DECIMAL(2) NULL,
                                                            `product_type_id` BIGINT UNSIGNED,
                                                            `seller_id` BIGINT UNSIGNED,
                                                            PRIMARY KEY (`id`)
) ENGINE = InnoDB;

-- -----------------------------------------------------
-- Table `mercado-fresco`.`users`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mercado-fresco`.`users` (
                                                        `id` SERIAL,
                                                        `username` VARCHAR(255) NOT NULL,
                                                        `password` VARCHAR(255) NOT NULL,
                                                        PRIMARY KEY (`id`)
) ENGINE = InnoDB;

-- -----------------------------------------------------
-- Table `mercado-fresco`.`user_rol`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mercado-fresco`.`rol` (
                                                      `id` SERIAL,
                                                      `rol_name` VARCHAR(255) NOT NULL,
                                                      `description` VARCHAR(255) NOT NULL,
                                                      PRIMARY KEY (`id`)
) ENGINE = InnoDB;

-- -----------------------------------------------------
-- Table `mercado-fresco`.`user_rol`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mercado-fresco`.`user_rol` (
                                                           `usuario_id` BIGINT UNSIGNED,
                                                           `rol_id` BIGINT UNSIGNED
) ENGINE = InnoDB;

-- -----------------------------------------------------
-- Table `mercado-fresco`.`countries`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mercado-fresco`.`countries` (
                                                            `id` SERIAL,
                                                            `country_name` VARCHAR(255) NULL,
                                                            PRIMARY KEY (`id`)
) ENGINE = InnoDB;

-- -----------------------------------------------------
-- Table `mercado-fresco`.`provinces`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mercado-fresco`.`provinces` (
                                                            `id` SERIAL,
                                                            `province_name` VARCHAR(255) NULL,
                                                            `id_country` BIGINT UNSIGNED,
                                                            PRIMARY KEY (`id`)
) ENGINE = InnoDB;

-- -----------------------------------------------------
-- Table `mercado-fresco`.`localities`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mercado-fresco`.`localities` (
                                                             `id` SERIAL,
                                                             `locality_name` VARCHAR(255) NULL,
                                                             `province_id` BIGINT UNSIGNED,
                                                             PRIMARY KEY (`id`)
) ENGINE = InnoDB;

-- -----------------------------------------------------
-- Table `mercado-fresco`.`product_types`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mercado-fresco`.`product_types` (
                                                                `id` SERIAL,
                                                                `description` VARCHAR(255) NULL,
                                                                PRIMARY KEY (`id`)
) ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `mercado-fresco`.`inbound_orders`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mercado-fresco`.`inbound_orders` (
                                                                 `id` SERIAL,
                                                                 `order_date` DATETIME(6),
                                                                 `order_number` VARCHAR(255),
                                                                 `employee_id` BIGINT UNSIGNED,
                                                                 `product_batch_id` BIGINT UNSIGNED,
                                                                 `warehouse_id` BIGINT UNSIGNED,
                                                                 PRIMARY KEY (`id`)
) ENGINE = InnoDB;

-- -----------------------------------------------------
-- Table `mercado-fresco`.`product_batches`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mercado-fresco`.`product_batches` (
                                                                  `id` SERIAL,
                                                                  `batch_number` INT,
                                                                  `current_quantity` INT,
                                                                  `current_temperature` INT,
                                                                  `due_date` DATETIME(6),
                                                                  `initial_quantity` INT,
                                                                  `manufacturing_date` DATETIME(6),
                                                                  `manufacturing_hour` INT,
                                                                  `minimum_temperature` INT,
                                                                  `product_id` BIGINT UNSIGNED,
                                                                  `section_id` BIGINT UNSIGNED,
                                                                  PRIMARY KEY (`id`)
) ENGINE = InnoDB;

-- -----------------------------------------------------
-- Table `mercado-fresco`.`product_records`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mercado-fresco`.`product_records` (
                                                                  `id` SERIAL,
                                                                  `last_update_date` DATETIME(6),
                                                                  `purchase_price` DECIMAL(19,2),
                                                                  `sale_price` DECIMAL(19,2),
                                                                  `product_id` BIGINT UNSIGNED,
                                                                  PRIMARY KEY (`id`)
) ENGINE = InnoDB;

-- -----------------------------------------------------
-- Table `mercado-fresco`.`order_status`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mercado-fresco`.`order_status` (
                                                               `id` SERIAL,
                                                               `description` VARCHAR(255),
                                                               PRIMARY KEY (`id`)
) ENGINE = InnoDB;

-- -----------------------------------------------------
-- Table `mercado-fresco`.`carriers`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mercado-fresco`.`carriers` (
                                                           `id` SERIAL,
                                                           `cid` VARCHAR(255),
                                                           `company_name` VARCHAR(255),
                                                           `address` VARCHAR(255),
                                                           `telephone` VARCHAR(255),
                                                           `locality_id` BIGINT UNSIGNED,
                                                           PRIMARY KEY (`id`)
) ENGINE = InnoDB;

-- -----------------------------------------------------
-- Table `mercado-fresco`.`purchase_orders`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mercado-fresco`.`purchase_orders` (
                                                                  `id` SERIAL,
                                                                  `order_number` VARCHAR(255),
                                                                  `order_date` DATETIME(6),
                                                                  `tracking_code` VARCHAR(255),
                                                                  `buyer_id` BIGINT UNSIGNED,
                                                                  `carrier_id` BIGINT UNSIGNED,
                                                                  `order_status_id` BIGINT UNSIGNED,
                                                                  `werehouse_id` BIGINT UNSIGNED,
                                                                  PRIMARY KEY (`id`)
) ENGINE = InnoDB;

-- -----------------------------------------------------
-- Table `mercado-fresco`.`order_details`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mercado-fresco`.`order_details` (
                                                                `id` SERIAL,
                                                                `clean_lines_status` VARCHAR(255),
                                                                `quantity` INT(11),
                                                                `temperature` DECIMAL(19,2),
                                                                `product_record_id` BIGINT UNSIGNED,
                                                                `purchase_order_id` BIGINT UNSIGNED,
                                                                PRIMARY KEY (`id`)
) ENGINE = InnoDB;

-- -----------------------------------------------------
-- Table `mercado-fresco`.`section`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `mercado-fresco`.`section` (
                                                          `id` SERIAL,
                                                          `section_number` INT(11) NULL,
                                                          `current_temperature` INT(11) NULL,
                                                          `minimum_temperature` INT(11) NULL,
                                                          `current_capacity` INT(11) NULL,
                                                          `minimum_capacity` INT(11) NULL,
                                                          `maximum_capacity` INT(11) NULL,
                                                          `warehouse_id` BIGINT UNSIGNED,
                                                          `product_type_id` BIGINT UNSIGNED,
                                                          PRIMARY KEY (`id`)
) ENGINE = InnoDB;

ALTER TABLE `mercado-fresco`.`user_rol` ADD CONSTRAINT `FK_USER_ROL_USER` FOREIGN KEY (`usuario_id`) REFERENCES `mercado-fresco`.`users`(`id`);
ALTER TABLE `mercado-fresco`.`user_rol` ADD CONSTRAINT `FK_USER_ROL_ROL` FOREIGN KEY (`rol_id`) REFERENCES `mercado-fresco`.`rol`(`id`);

ALTER TABLE `mercado-fresco`.`provinces` ADD CONSTRAINT `FK_PROVINCE_COUNTRY` FOREIGN KEY (`id_country`) REFERENCES `mercado-fresco`.`countries`(`id`);

ALTER TABLE `mercado-fresco`.`localities` ADD CONSTRAINT `FK_LOCALITIES_PROVINCE` FOREIGN KEY (`province_id`) REFERENCES `mercado-fresco`.`provinces`(`id`);

ALTER TABLE `mercado-fresco`.`carriers` ADD CONSTRAINT `FK_CARRIERS_LOCALITY` FOREIGN KEY (`locality_id`) REFERENCES `mercado-fresco`.`localities`(`id`);

ALTER TABLE `mercado-fresco`.`product_records` ADD CONSTRAINT `FK_PRODUCT_RECORDS_PRODUCT` FOREIGN KEY (`product_id`) REFERENCES `mercado-fresco`.`products`(`id`);

ALTER TABLE `mercado-fresco`.`order_details` ADD CONSTRAINT `FK_ORDER_DETAILS_PRODUCT_RECORD` FOREIGN KEY (`product_record_id`) REFERENCES `mercado-fresco`.`product_records`(`id`);
ALTER TABLE `mercado-fresco`.`order_details` ADD CONSTRAINT `FK_ORDER_DETAILS_PURCHASE_ORDER` FOREIGN KEY (`purchase_order_id`) REFERENCES `mercado-fresco`.`purchase_orders`(`id`);

ALTER TABLE `mercado-fresco`.`inbound_orders` ADD CONSTRAINT `FK_INBOUND_ORDERS_EMPLOYEE` FOREIGN KEY (`employee_id`) REFERENCES `mercado-fresco`.`employees`(`id`);
ALTER TABLE `mercado-fresco`.`inbound_orders` ADD CONSTRAINT `FK_INBOUND_ORDERS_PRODUCT_BATCH` FOREIGN KEY (`product_batch_id`) REFERENCES `mercado-fresco`.`product_batches`(`id`);
ALTER TABLE `mercado-fresco`.`inbound_orders` ADD CONSTRAINT `FK_INBOUND_ORDERS_WAREHOUSE` FOREIGN KEY (`warehouse_id`) REFERENCES `mercado-fresco`.`warehouse`(`id`);

ALTER TABLE `mercado-fresco`.`product_batches` ADD CONSTRAINT `FK_PRODUCT_BATCHES_PRODUCT` FOREIGN KEY (`product_id`) REFERENCES `mercado-fresco`.`products`(`id`);
ALTER TABLE `mercado-fresco`.`product_batches` ADD CONSTRAINT `FK_PRODUCT_BATCHES_SECTION` FOREIGN KEY (`section_id`) REFERENCES `mercado-fresco`.`section`(`id`);

ALTER TABLE `mercado-fresco`.`purchase_orders` ADD CONSTRAINT `FK_PURCHASE_ORDERS_BUYER` FOREIGN KEY (`buyer_id`) REFERENCES `mercado-fresco`.`buyers`(`id`);
ALTER TABLE `mercado-fresco`.`purchase_orders` ADD CONSTRAINT `FK_PURCHASE_ORDERS_CARRIER` FOREIGN KEY (`carrier_id`) REFERENCES `mercado-fresco`.`carriers`(`id`);
ALTER TABLE `mercado-fresco`.`purchase_orders` ADD CONSTRAINT `FK_PURCHASE_ORDERS_WAREHOUSE` FOREIGN KEY (`werehouse_id`) REFERENCES `mercado-fresco`.`warehouse`(`id`);
ALTER TABLE `mercado-fresco`.`purchase_orders` ADD CONSTRAINT `FK_PURCHASE_ORDERS_STATUS_ORDER` FOREIGN KEY (`order_status_id`) REFERENCES `mercado-fresco`.`order_status`(`id`);

ALTER TABLE `mercado-fresco`.`sellers` ADD CONSTRAINT `FK_SELLER_LOCALITY` FOREIGN KEY (`locality_id`) REFERENCES `mercado-fresco`.`localities`(`id`);

ALTER TABLE `mercado-fresco`.`products` ADD CONSTRAINT `FK_PRODUCT_PRODUCT_TYPE` FOREIGN KEY (`product_type_id`) REFERENCES `mercado-fresco`.`product_types`(`id`);
ALTER TABLE `mercado-fresco`.`products` ADD CONSTRAINT `FK_PRODUCT_SELLER` FOREIGN KEY (`seller_id`) REFERENCES `mercado-fresco`.`sellers`(`id`);

ALTER TABLE `mercado-fresco`.`employees` ADD CONSTRAINT `FK_EMPLOYEE_WAREHOUSE` FOREIGN KEY (`warehouse_id`) REFERENCES `mercado-fresco`.`warehouse`(`id`);

ALTER TABLE `mercado-fresco`.`warehouse` ADD CONSTRAINT `FK_WAREHOUSE_LOCALITY` FOREIGN KEY (`locality_id`) REFERENCES `mercado-fresco`.`localities`(`id`);

ALTER TABLE `mercado-fresco`.`section` ADD CONSTRAINT `FK_SECTION_WAREHOUSE` FOREIGN KEY (`warehouse_id`) REFERENCES `mercado-fresco`.`warehouse`(`id`);
ALTER TABLE `mercado-fresco`.`section` ADD CONSTRAINT `FK_SECTION_PRODUCT` FOREIGN KEY (`product_type_id`) REFERENCES `mercado-fresco`.`product_types`(`id`);
