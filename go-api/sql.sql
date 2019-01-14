create database test;

create table users(
id int auto_increment primary key,
name varchar(100) not null,
email varchar(100) not null unique,
password varchar(100) not null,
createdAt timestamp default current_timestamp()
)
engine=InnoDB
default charset = utf8;
