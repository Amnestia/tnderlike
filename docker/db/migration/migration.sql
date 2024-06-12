CREATE TABLE IF NOT EXISTS account(
	id BIGSERIAL PRIMARY KEY NOT NULL,
	acc_email VARCHAR(255) NOT NULL UNIQUE
	CHECK(acc_email ~ '^[a-zA-Z0-9.!#$%&''*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$'),
	acc_password VARCHAR(1024) NOT NULL,
	created_at TIMESTAMP NOT NULL default NOW(),
	created_by VARCHAR(255) NOT NULL,
	updated_at TIMESTAMP NOT NULL default NOW(),
	updated_by VARCHAR(255) NOT NULL,
	deleted_at TIMESTAMP,
	deleted_by VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS profile(
	id BIGSERIAL PRIMARY KEY NOT NULL,
	account_id BIGINT NOT NULL REFERENCES account(id) ON UPDATE NO ACTION ON DELETE NO ACTION,
	fullname VARCHAR(255) NOT NULL,
	location VARCHAR(255) NOT NULL,
	photoURL VARCHAR(255) NOT NULL,
	bio VARCHAR(1024) NOT NULL,
	verified BOOL NOT NULL DEFAULT FALSE,
	unlimited_swipe BOOL NOT NULL DEFAULT FALSE,
	created_at TIMESTAMP NOT NULL default NOW(),
	created_by VARCHAR(255) NOT NULL,
	updated_at TIMESTAMP NOT NULL default NOW(),
	updated_by VARCHAR(255) NOT NULL,
	deleted_at TIMESTAMP,
	deleted_by VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS history(
	id BIGSERIAL PRIMARY KEY NOT NULL,
	viewer_id BIGINT NOT NULL REFERENCES account(id) ON UPDATE NO ACTION ON DELETE NO ACTION,
	viewed_id BIGINT NOT NULL REFERENCES account(id) ON UPDATE NO ACTION ON DELETE NO ACTION,
	action INT NOT NULL,
	created_at TIMESTAMP NOT NULL default NOW(),
	created_by VARCHAR(255) NOT NULL,
	updated_at TIMESTAMP NOT NULL default NOW(),
	updated_by VARCHAR(255) NOT NULL,
	deleted_at TIMESTAMP,
	deleted_by VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS package_snapshot(
	id BIGSERIAL PRIMARY KEY NOT NULL,
	name VARCHAR(255) NOT NULL,
	price INT NOT NULL,
	created_at TIMESTAMP NOT NULL default NOW(),
	created_by VARCHAR(255) NOT NULL,
	updated_at TIMESTAMP NOT NULL default NOW(),
	updated_by VARCHAR(255) NOT NULL,
	deleted_at TIMESTAMP,
	deleted_by VARCHAR(255)
);


CREATE TABLE IF NOT EXISTS package(
	id BIGSERIAL PRIMARY KEY NOT NULL,
	name VARCHAR(255) NOT NULL,
	price INT NOT NULL,
	created_at TIMESTAMP NOT NULL default NOW(),
	created_by VARCHAR(255) NOT NULL,
	updated_at TIMESTAMP NOT NULL default NOW(),
	updated_by VARCHAR(255) NOT NULL,
	deleted_at TIMESTAMP,
	deleted_by VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS payment_id(
	id BIGSERIAL PRIMARY KEY NOT NULL,
	account_id BIGINT NOT NULL REFERENCES account(id) ON UPDATE NO ACTION ON DELETE NO ACTION,
	package_snapshot_id BIGINT NOT NULL REFERENCES package_snapshot(id) ON UPDATE NO ACTION ON DELETE NO ACTION,
	payment_id VARCHAR(255) NOT NULL,
	created_at TIMESTAMP NOT NULL default NOW(),
	created_by VARCHAR(255) NOT NULL,
	updated_at TIMESTAMP NOT NULL default NOW(),
	updated_by VARCHAR(255) NOT NULL,
	deleted_at TIMESTAMP,
	deleted_by VARCHAR(255)
);
