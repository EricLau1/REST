create database test;

create table login(
id int auto_increment primary key,
nome varchar(25),
senha varchar(50),
created_at timestamp default current_timestamp()
)
default charset = utf8;

insert into login (nome, senha) values ('admin', '@admin');

create table produto(
id int auto_increment primary key,
descricao varchar(100),
preco double,
quantidade int,
created_at timestamp default current_timestamp()
)
default charset = utf8;

insert into produto (descricao, preco, quantidade) values 
('Apple', 2.00, 350),
('Kiwi', 5.00, 900),
('Banana', 0.9, 100);