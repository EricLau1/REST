create database crud_rest
default character set utf8 
default collate utf8_general_ci;

use crud_rest;

create table tb_produto (
id int auto_increment,
descricao varchar(100) not null,
quantidade int default 0,
valor decimal(7,2),
constraint tb_produto_id_pk primary key(id)
)
engine = InnoDB
default charset = utf8;

insert into tb_produto (descricao, quantidade, valor) values
('Notebook', 10, 2500),
('Celular', 20, 7500),
('Roteador', 15, 300);