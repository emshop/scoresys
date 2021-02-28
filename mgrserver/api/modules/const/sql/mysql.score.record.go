package sql
//InsertScoreRecord 添加兑换记录
const InsertScoreRecord = `
insert into sc_score_record
(
	
	c_tp,
	score,
	remain
)
values
(
	
	if(isnull(@c_tp)||@c_tp='',0,@c_tp),
	if(isnull(@score)||@score='',0,@score),
	if(isnull(@remain)||@remain='',0,@remain)
)`

//GetScoreRecordByRcID 查询单条数据兑换记录
const GetScoreRecordByRcID = `
select
	t.c_tp,
	t.score,
	t.remain,
	t.create_time
from sc_score_record t
where
	&rc_id`

//GetScoreRecordListCount 获取兑换记录列表条数
const GetScoreRecordListCount = `
select count(1)
from sc_score_record t
where
	&t.uid
	&t.c_tp`

//GetScoreRecordList 查询兑换记录列表数据
const GetScoreRecordList = `
select
	t.c_tp,
	t.score,
	t.create_time 
from sc_score_record t
where
	&t.uid
	&t.c_tp
order by t.rc_id desc
limit @ps offset @offset
`
//UpdateScoreRecordByRcID 更新兑换记录
const UpdateScoreRecordByRcID = `
update sc_score_record 
set
	c_tp =	if(isnull(@c_tp)||@c_tp='',0,@c_tp),
	score =	if(isnull(@score)||@score='',0,@score),
	remain =	if(isnull(@remain)||@remain='',0,@remain)
where
	&rc_id`

