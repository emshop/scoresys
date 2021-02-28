package scheme
 
//sc_reward_info 奖励信息
const sc_reward_info=`
	DROP TABLE IF EXISTS sc_reward_info;
	CREATE TABLE IF NOT EXISTS sc_reward_info (
		rw_id int  not null auto_increment comment '编号' ,
		rw_name varchar(64)  not null  comment '名称' ,
		category int default 0 not null  comment '分类' ,
		score int default 0 not null  comment '分数' ,
		status tinyint default 0 not null  comment '状态' ,
		create_time datetime default current_timestamp   comment '创建时间' 
		,primary key (rw_id)
	) ENGINE=InnoDB auto_increment = 10000 DEFAULT CHARSET=utf8mb4 COMMENT='奖励信息'`