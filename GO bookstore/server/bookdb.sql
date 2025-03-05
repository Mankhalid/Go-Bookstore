CREATE DATABASE bookdb;
USE	bookdb;
CREATE table books(
id INT auto_increment Primary KEY,
title varchar(255) NOT NULL,
author varchar(255) NOT NULL,
isbn varchar(100) NOT NULL unique
);
USE bookdb;
SHOW TABLES;
INSERT INTO books (title, author, isbn) VALUES 
('Test Book', 'Test Author', '1234567890');

