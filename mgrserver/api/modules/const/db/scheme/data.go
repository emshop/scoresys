package scheme

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/micro-plat/hydra"
)

func init() {
	//注册服务包
	hydra.OnReady(func() error {
		hydra.Installer.DB.AddSQL(
			"INSERT INTO sc_dict_info(id,name,value,type,status,sort_no) VALUES(1,'启用','0','status',0,0),(2,'生活习惯','0','category',0,0),(3,'学校学习','1','category',0,0),(4,'家庭学习','2','category',0,1),(5,'奖励','reward_info','c_tp',0,0),(6,'兑换','gif_info','c_tp',0,0),(7,'惩罚','punishment_info','c_tp',0,0);",
			"INSERT INTO sc_reward_info(rw_id,rw_name,category,score,status,create_time) VALUES(1,'7点前起床',0,5,0,'2021-02-28 12:42:34'),(2,'7点30前吃完早餐',0,5,0,'2021-02-28 13:13:58'),(3,'一日三餐光盘',0,10,0,'2021-02-28 13:15:05'),(4,'一天喝完1杯+水',0,10,0,'2021-02-28 13:15:28'),(5,'自已洗漱、洗澡',0,10,0,'2021-02-28 13:16:00'),(7,'主动收拾书桌、准备文具',0,5,0,'2021-02-28 13:21:50'),(8,'爱干净、讲卫生',0,10,0,'2021-02-28 13:22:34'),(9,'听父母、老师、长辈的话',0,10,0,'2021-02-28 13:23:25'),(10,'主动收拾碗筷、洗碗、扫地或收拾',0,5,0,'2021-02-28 13:23:48'),(11,'不生气、不发火、不打闹',0,10,0,'2021-02-28 13:24:29'),(12,'主动给老师、同学打招呼',1,5,0,'2021-02-28 13:27:42'),(13,'认真听老师讲课(手放好，耳朵、大脑都在听)',1,20,0,'2021-02-28 13:36:51'),(14,'在学校完成作业',1,20,0,'2021-02-28 13:37:09'),(15,'得到老师表扬',1,50,0,'2021-02-28 13:37:30'),(16,'自主完成叫叫阅读',2,10,0,'2021-02-28 13:39:09'),(19,'跳绳600个',0,15,0,'2021-02-28 20:31:08');",
			"INSERT INTO sc_gif_info(gif_id,gif_name,score,status,create_time) VALUES(1,'看电视30分钟',-100,0,'2021-02-28 13:46:22');",
			"INSERT INTO sc_user_info(uid,name,url,birthday,score,status,create_time) VALUES(1,'杨展硕','/yzs.jpg','2014-08-06 00:00:00',200,0,'2021-03-04 10:37:55'),(2,'杨铠硕','/yks.jpg','2016-05-19 00:00:00',200,0,'2021-03-04 10:38:18');",
		)
		return nil
	})
}
