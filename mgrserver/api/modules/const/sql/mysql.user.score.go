package sql

//score 更新用户分数
const UpdateUserScoreByUid = `
update sc_user_info 
set
score =	score + (select sum(t.score) from sc_score_record t where t.batch_id = @batch_id and t.uid= @uid)	
where
	&uid`
