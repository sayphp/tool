-- 创建测试用数据库
CREATE DATABASE `test`;

-- 创建测试用数据表

CREATE TABLE `tblUser` (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `username` VARCHAR(50) NOT NULL,
    `password` VARCHAR(50) NOT NULL,
    `email` VARCHAR(50) NOT NULL,
    PRIMARY KEY (id)
);

-- 测试数据
INSERT INTO tblUser(`username`, `password`, `email`) VALUES ("test", "123456", "liuxiao@zuoyebang.com");