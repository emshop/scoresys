package scheme
 
//sc_dict_info 字典配置
const sc_dict_info=`
	DROP TABLE IF EXISTS sc_dict_info;
	CREATE TABLE IF NOT EXISTS sc_dict_info (
		id int  not null auto_increment comment '编号' ,
		name varchar(64)  not null  comment '名称' ,
		value varchar(32)  not null  comment '值' ,
		type varchar(32)  not null  comment '类型' ,
		status tinyint default 0 not null  comment '状态' ,
		sort_no tinyint default 0 not null  comment '排序值' 
		,primary key (id)
	) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='字典配置'`