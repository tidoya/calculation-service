CREATE TABLE users (
    id serial PRIMARY KEY,
    name varchar(255) not null,
    username varchar(255) not null unique,
    password_hash varchar(255) not null
);