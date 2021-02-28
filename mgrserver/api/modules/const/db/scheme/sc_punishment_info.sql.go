package scheme
 
//sc_punishment_info 惩罚信息
const sc_punishment_info=`
	DROP TABLE IF EXISTS sc_punishment_info;
	CREATE TABLE IF NOT EXISTS sc_punishment_info (
		pn_id int  not null auto_increment comment '编号' ,
		pn_name varchar(64)  not null  comment '名称' ,
		category int default 0 not null  comment '分类' ,
		score int default 0 not null  comment '分数' ,
		status tinyint default 0 not null  comment '状态' ,
		create_time datetime default current_timestamp not null  comment '创建时间' 
		,primary key (pn_id)
	) ENGINE=InnoDB auto_increment = 2000 DEFAULT CHARSET=utf8mb4 COMMENT='惩罚信息'`