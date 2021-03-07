package scheme
 
//sc_score_record 分数记录
const sc_score_record=`
	DROP TABLE IF EXISTS sc_score_record;
	CREATE TABLE IF NOT EXISTS sc_score_record (
		rc_id int  not null auto_increment comment '编号' ,
		uid int  not null  comment '用户' ,
		c_tp tinyint    comment '类型' ,
		rwpu_id bigint    comment '奖励、惩罚编号' ,
		score int  not null  comment '变动分数' ,
		content varchar(64)  not null  comment '内容' ,
		batch_id bigint default 0   comment '加入批次编号' ,
		create_time datetime default current_timestamp not null  comment '创建时间' 
		,primary key (rc_id)
	) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='分数记录'`