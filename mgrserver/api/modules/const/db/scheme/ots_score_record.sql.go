package scheme
 
//ots_score_record 兑换表
const ots_score_record=`
	DROP TABLE IF EXISTS ots_score_record;
	CREATE TABLE IF NOT EXISTS ots_score_record (
		rc_id int  not null  comment '编号' ,
		uid int  not null  comment '用户' ,
		c_tp tinyint    comment '类型' ,
		score int  not null  comment '变动分数' ,
		remain bigint default 0   comment '剩余' ,
		create_time datetime default current_timestamp not null  comment '创建时间' 
		,primary key (rc_id)
	) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='兑换表'`