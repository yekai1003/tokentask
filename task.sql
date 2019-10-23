create database tokentask character set utf8;

use tokentask

create table t_user(name varchar(30), pass varchar(16),remark varchar(100));

insert into t_user values('yekai', 'admin','boss'),('fuhongxue','yekai','xiaodi');

create table t_task(
	task_id int primary key auto_increment,
	user_name varchar(30) not null,
	task_user varchar(30),
	bonus int ,
	status varchar(20),
	remark varchar(100),
	comment varchar(200)
);


delete from t_task;

insert into t_task(user_name, task_user, bonus, status, remark, comment)
values('路小佳', '叶开', 1000, '进行中', '洗澡',null);


insert into t_task(user_name, task_user, bonus, status, remark, comment)
values('万马堂', '路小佳', 10000, '任务完结', '杀傅红雪','任务没有执行完成，没有出手');

insert into t_task(user_name, task_user, bonus, status, remark, comment)
values('李寻欢', '叶开', 0, '进行中', '寻找真相',null);