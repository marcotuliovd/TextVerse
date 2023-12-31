CREATE DATABASE IF NOT EXISTS textverse;
USE textverse;

DROP TABLE IF EXISTS usuarios;

CREATE TABLE usuarios(
    id SERIAL PRIMARY KEY,
    nome varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    senha varchar(50) not null,
    create_at timestamp default current_timestamp()
)