CREATE DATABASE my_db;

GRANT ALL PRIVILEGES ON my_db.* TO 'backend'@'%';
FLUSH PRIVILEGES;

USE my_db;

CREATE TABLE users (
	name 	 VARCHAR(20),
	password VARCHAR(20)
);

SHOW TABLES;

INSERT INTO users
	(name, password)
VALUES
	('debhole', 'my_pw'),
	('kamistr', 'kavi1234'),
	('tacocat', 'austin9000');
