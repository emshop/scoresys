package sql

//score 更新用户分数
const UpdateUserScoreByUid = `
update sc_user_info 
set
score =	score + @score	
where
	&uid`
