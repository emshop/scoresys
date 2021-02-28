package sql
//InsertDictInfo 添加字典配置
const InsertDictInfo = `
insert into sc_dict_info
(
	
	name,
	value,
	type,
	status,
	sort_no
)
values
(
	
	@name,
	@value,
	@type,
	if(isnull(@status)||@status='',0,@status),
	if(isnull(@sort_no)||@sort_no='',0,@sort_no)
)`

//GetDictInfoByID 查询单条数据字典配置
const GetDictInfoByID = `
select
	t.id,
	t.name,
	t.value,
	t.type,
	t.status,
	t.sort_no
from sc_dict_info t
where
	&id`

//GetDictInfoListCount 获取字典配置列表条数
const GetDictInfoListCount = `
select count(1)
from sc_dict_info t
where
	?t.name
	&t.type
	&t.status`

//GetDictInfoList 查询字典配置列表数据
const GetDictInfoList = `
select
	t.id,
	t.name,
	t.value,
	t.type,
	t.status,
	t.sort_no 
from sc_dict_info t
where
	?t.name
	&t.type
	&t.status
order by t.id desc
limit @ps offset @offset
`
//UpdateDictInfoByID 更新字典配置
const UpdateDictInfoByID = `
update sc_dict_info 
set
	name =	@name,
	value =	@value,
	type =	@type,
	status =	if(isnull(@status)||@status='',0,@status),
	sort_no =	if(isnull(@sort_no)||@sort_no='',0,@sort_no)
where
	&id`

