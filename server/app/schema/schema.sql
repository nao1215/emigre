CREATE DATABASE IF NOT EXISTS emigre;
USE emigre;

CREATE TABLE `user` (
    `id` VARCHAR(191) NOT NULL COMMENT 'ID is user id. ULID（Universally Unique Lexicographically Sortable Identifier）',
    `name` VARCHAR(20) NOT NULL COMMENT 'Name is user name',
    `biography` VARCHAR(300) NOT NULL COMMENT 'Biography is user self introduction',
    `email` VARCHAR(256) NOT NULL COMMENT 'Email is user email address',
    `created_at` DATETIME NOT NULL COMMENT 'CreatedAt is the date that record was created',
    `updated_at` DATETIME NOT NULL COMMENT 'UpdatedAt is the date record was updated',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARACTER SET utf8mb4 COMMENT='user table';