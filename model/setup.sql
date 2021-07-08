-- 数据库
CREATE DATABASE ps_db;

-- 用户表
CREATE TABLE user_t(
    id       SERIAL PRIMARY KEY,
    name     VARCHAR(64),
    email    VARCHAR(64),
    password VARCHAR(256),
    token    VARCHAR(64)
);

-- 文章表
CREATE TABLE artical_t(
    id    SERIAL PRIMARY KEY,
    title VARCHAR(64),
    category VARCHAR(8),
    content VARCHAR(1024),
);

