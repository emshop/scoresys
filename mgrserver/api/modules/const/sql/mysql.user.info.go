package sql
//InsertUserInfo 添加用户信息
const InsertUserInfo = `
insert into sc_user_info
(
	uid,
	name,
	url,
	status
)
values
(
	@uid,
	@name,
	@url,
	if(isnull(@status)||@status='',0,@status)
)`

//GetUserInfoByUid 查询单条数据用户信息
const GetUserInfoByUid = `
select
	t.uid,
	t.name,
	t.url,
	t.score,
	t.status,
	t.create_time
from sc_user_info t
where
	&uid`

//GetUserInfoListCount 获取用户信息列表条数
const GetUserInfoListCount = `
select count(1)
from sc_user_info t
where
	?t.name
	&t.status`

//GetUserInfoList 查询用户信息列表数据
const GetUserInfoList = `
select
	t.uid,
	t.name,
	t.url,
	t.score,
	t.status,
	t.create_time 
from sc_user_info t
where
	?t.name
	&t.status
order by t.uid desc
limit @ps offset @offset
`
//UpdateUserInfoByUid 更新用户信息
const UpdateUserInfoByUid = `
update sc_user_info 
set
	name =	@name,
	url =	@url,
	status =	if(isnull(@status)||@status='',0,@status)
where
	&uid`

