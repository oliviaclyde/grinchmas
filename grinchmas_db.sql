DROP DATABASE IF EXISTS grinchmas;

CREATE DATABASE grinchmas;

USE grinchmas;

CREATE TABLE events (
    id int PRIMARY KEY NOT NULL AUTO_INCREMENT,
    event_description varchar(255)
);

INSERT INTO events (id, event_description) VALUES (1, 'go to in-laws for Christmas Eve dinner');