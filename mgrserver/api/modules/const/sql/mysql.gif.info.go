package sql
//InsertGifInfo 添加礼品信息
const InsertGifInfo = `
insert into sc_gif_info
(
	gif_id,
	gif_name,
	score,
	status
)
values
(
	@gif_id,
	@gif_name,
	if(isnull(@score)||@score='',0,@score),
	if(isnull(@status)||@status='',0,@status)
)`

//GetGifInfoByGifID 查询单条数据礼品信息
const GetGifInfoByGifID = `
select
	t.gif_id,
	t.gif_name,
	t.score,
	t.status,
	t.create_time
from sc_gif_info t
where
	&gif_id`

//GetGifInfoListCount 获取礼品信息列表条数
const GetGifInfoListCount = `
select count(1)
from sc_gif_info t
where
	?t.gif_name
	&t.status`

//GetGifInfoList 查询礼品信息列表数据
const GetGifInfoList = `
select
	t.gif_id,
	t.gif_name,
	t.score,
	t.status,
	t.create_time 
from sc_gif_info t
where
	?t.gif_name
	&t.status
order by t.gif_id desc
limit @ps offset @offset
`
//UpdateGifInfoByGifID 更新礼品信息
const UpdateGifInfoByGifID = `
update sc_gif_info 
set
	gif_name =	@gif_name,
	score =	if(isnull(@score)||@score='',0,@score),
	status =	if(isnull(@status)||@status='',0,@status)
where
	&gif_id`

