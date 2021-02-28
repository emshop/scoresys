package scheme
 
//sc_user_info 用户信息
const sc_user_info=`
	DROP TABLE IF EXISTS sc_user_info;
	CREATE TABLE IF NOT EXISTS sc_user_info (
		uid int  not null auto_increment comment '用户编号' ,
		name varchar(64)  not null  comment '姓名' ,
		score int default 0 not null  comment '分数' ,
		status tinyint default 0 not null  comment '状态' ,
		create_time datetime default current_timestamp not null  comment '创建时间' 
		,primary key (uid)
	) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='用户信息'`