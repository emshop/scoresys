package sql
//InsertScoreRecord 添加分数记录
const InsertScoreRecord = `
insert into sc_score_record
(
	
	uid,
	c_tp,
	rwpu_id,
	score,
	content,
	batch_id
)
values
(
	
	if(isnull(@uid)||@uid='',0,@uid),
	if(isnull(@c_tp)||@c_tp='',0,@c_tp),
	if(isnull(@rwpu_id)||@rwpu_id='',0,@rwpu_id),
	if(isnull(@score)||@score='',0,@score),
	@content,
	if(isnull(@batch_id)||@batch_id='',0,@batch_id)
)`

//GetScoreRecordByRcID 查询单条数据分数记录
const GetScoreRecordByRcID = `
select
	t.uid,
	t.c_tp,
	t.rwpu_id,
	t.score,
	t.content,
	t.batch_id,
	t.create_time
from sc_score_record t
where
	&rc_id`

//GetScoreRecordListCount 获取分数记录列表条数
const GetScoreRecordListCount = `
select count(1)
from sc_score_record t
where
	&t.uid
	and t.create_time >= @create_time 
	and t.create_time < date_add(@create_time, interval 1 day)`

//GetScoreRecordList 查询分数记录列表数据
const GetScoreRecordList = `
select
	t.uid,
	t.c_tp,
	t.rwpu_id,
	t.score,
	t.content,
	t.create_time 
from sc_score_record t
where
	&t.uid
	and t.create_time >= @create_time 
	and t.create_time < date_add(@create_time, interval 1 day)
order by t.rc_id desc
limit @ps offset @offset
`
//UpdateScoreRecordByRcID 更新分数记录
const UpdateScoreRecordByRcID = `
update sc_score_record 
set
	c_tp =	if(isnull(@c_tp)||@c_tp='',0,@c_tp),
	rwpu_id =	if(isnull(@rwpu_id)||@rwpu_id='',0,@rwpu_id),
	score =	if(isnull(@score)||@score='',0,@score),
	content =	@content
where
	&rc_id`

