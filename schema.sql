CREATE SCHEMA IF NOT EXISTS ed_platform;

	CREATE TABLE IF NOT EXISTS users (
		id         SERIAL NOT NULL PRIMARY KEY,
		email      VARCHAR(120) NOT NULL UNIQUE,
		password   VARCHAR(255) NOT NULL,
		first_name VARCHAR(255),
		last_name  VARCHAR(255),
		created_timestamp TIMESTAMP DEFAULT NOW(),
		updated_timestamp TIMESTAMP
	);