drop database if exists tokentask ;
create database tokentask character set utf8;

use tokentask

create table t_user(
	user_id int primary key auto_increment,
	user_name varchar(100) not null,
	password varchar(20) not null,
	address varchar(100)
);

create table t_tasks(
	task_id int primary key,
	issue_user varchar(100),
	task_user varchar(100),
	bonus  int not null,
	task_name varchar(200),
	status smallint,
	comment varchar(200)
);