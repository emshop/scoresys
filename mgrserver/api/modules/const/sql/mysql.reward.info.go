package sql
//InsertRewardInfo 添加奖励信息
const InsertRewardInfo = `
insert into sc_reward_info
(
	rw_id,
	rw_name,
	category,
	score,
	status
)
values
(
	@rw_id,
	@rw_name,
	if(isnull(@category)||@category='',0,@category),
	if(isnull(@score)||@score='',0,@score),
	if(isnull(@status)||@status='',0,@status)
)`

//GetRewardInfoByRwID 查询单条数据奖励信息
const GetRewardInfoByRwID = `
select
	t.rw_id,
	t.rw_name,
	t.category,
	t.score,
	t.status,
	t.create_time
from sc_reward_info t
where
	&rw_id`

//GetRewardInfoListCount 获取奖励信息列表条数
const GetRewardInfoListCount = `
select count(1)
from sc_reward_info t
where
	?t.rw_name
	&t.category
	&t.status`

//GetRewardInfoList 查询奖励信息列表数据
const GetRewardInfoList = `
select
	t.rw_id,
	t.rw_name,
	t.category,
	t.score,
	t.status,
	t.create_time 
from sc_reward_info t
where
	?t.rw_name
	&t.category
	&t.status
order by t.rw_id desc
limit @ps offset @offset
`
//UpdateRewardInfoByRwID 更新奖励信息
const UpdateRewardInfoByRwID = `
update sc_reward_info 
set
	rw_name =	@rw_name,
	category =	if(isnull(@category)||@category='',0,@category),
	score =	if(isnull(@score)||@score='',0,@score),
	status =	if(isnull(@status)||@status='',0,@status)
where
	&rw_id`

