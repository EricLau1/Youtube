create database mydb;

use mydb;

create table if not exists users (
id int auto_increment primary key,
nickname varchar(15) not null unique,
email varchar(40) not null unique,
password varchar(60) not null,
created_at timestamp default current_timestamp(),
updated_at timestamp default current_timestamp()    
);

create table if not exists  avatars (
id int auto_increment primary key,
user int not null unique,
image longblob not null,
name varchar(255) default 'image',
constraint avatars_user_fk foreign key(user) 
references users(id)
on delete cascade
on update cascade
);

create table if not exists feedbacks (
id int auto_increment primary key,
user int not null,
comment varchar(255),
constraint feedbacks_user_fk foreign key(user)
references users(id)
on delete cascade
on update cascade
);
