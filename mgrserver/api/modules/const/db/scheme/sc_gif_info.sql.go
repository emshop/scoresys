package scheme
 
//sc_gif_info 礼品信息
const sc_gif_info=`
	DROP TABLE IF EXISTS sc_gif_info;
	CREATE TABLE IF NOT EXISTS sc_gif_info (
		gif_id int  not null auto_increment comment '编号' ,
		gif_name varchar(64)  not null  comment '礼品名称' ,
		score int default 0 not null  comment '分数' ,
		status tinyint default 0 not null  comment '状态' ,
		create_time datetime default current_timestamp not null  comment '创建时间' 
		,primary key (gif_id)
	) ENGINE=InnoDB auto_increment = 1000 DEFAULT CHARSET=utf8mb4 COMMENT='礼品信息'`