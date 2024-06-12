FROM postgres:latest

ADD docker/db/migration/migration.sql /docker-entrypoint-initdb.d/

