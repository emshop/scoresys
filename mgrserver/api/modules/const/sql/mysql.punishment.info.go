package sql
//InsertPunishmentInfo 添加惩罚信息
const InsertPunishmentInfo = `
insert into sc_punishment_info
(
	pn_id,
	pn_id,
	pn_name,
	category,
	score,
	status
)
values
(
	@pn_id,
	if(isnull(@pn_id)||@pn_id='',0,@pn_id),
	@pn_name,
	if(isnull(@category)||@category='',0,@category),
	if(isnull(@score)||@score='',0,@score),
	if(isnull(@status)||@status='',0,@status)
)`

//GetPunishmentInfoByPnID 查询单条数据惩罚信息
const GetPunishmentInfoByPnID = `
select
	t.pn_id,
	t.pn_name,
	t.category,
	t.score,
	t.status,
	t.create_time
from sc_punishment_info t
where
	&pn_id`

//GetPunishmentInfoListCount 获取惩罚信息列表条数
const GetPunishmentInfoListCount = `
select count(1)
from sc_punishment_info t
where
	?t.pn_name
	&t.status`

//GetPunishmentInfoList 查询惩罚信息列表数据
const GetPunishmentInfoList = `
select
	t.pn_id,
	t.pn_name,
	t.category,
	t.score,
	t.status,
	t.create_time 
from sc_punishment_info t
where
	?t.pn_name
	&t.status
order by t.pn_id desc
limit @ps offset @offset
`
//UpdatePunishmentInfoByPnID 更新惩罚信息
const UpdatePunishmentInfoByPnID = `
update sc_punishment_info 
set
	pn_name =	@pn_name,
	category =	if(isnull(@category)||@category='',0,@category),
	score =	if(isnull(@score)||@score='',0,@score),
	status =	if(isnull(@status)||@status='',0,@status)
where
	&pn_id`

