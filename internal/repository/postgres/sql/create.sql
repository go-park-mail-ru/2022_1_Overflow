
DROP SCHEMA IF EXISTS overflow CASCADE ;
CREATE SCHEMA overflow

CREATE TABLE overflow.users (
    id serial  not null primary key,
    first_name varchar(30) not null,
    last_name varchar(30) not null,
    password varchar(30) not null,
    email varchar(234) not null unique
);
CREATE TABLE overflow.mails (
    id serial not null primary key,
    client_id serial not null,
    sender varchar(234) not null ,
    addressee varchar(234) not null ,
    date date not null,
    theme varchar(30),
    text text not null ,
    files varchar(30),
    read bool DEFAULT FALSE,
    foreign key (client_id) references overflow.users(id) on delete cascade
);