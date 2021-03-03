
PK:主键
SEQ:SEQ，mysql自增，oracle序列
C: 创建数据时的字段
R: 单条数据读取时的字段 
U: 修改数据时需要的字段
D: 删除，默认为更新字段状态值为1，D[(更新状态值)]
Q: 查询条件的字段
L：(前端页面)列表里列出的字段
OB：查询时的order by字段；默认降序； OB[(顺序)]，越小越先排序
DI: 字典编号，数据表作为字典数据时的id字段
DN: 字典名称，数据表作为字典数据时的name字段
SL: "select"      //表单下拉框,默认使用dds字典表枚举,指定表名的SL[(字典表名)]
CB: "checkbox"    //表单复选框,默认使用dds字典表枚举,指定表名的CB[(字典表名)]
RD: "radio"       //表单单选框,默认使用dds字典表枚举,指定表名的RB[(字典表名)]
TA: "textarea"    //表单文本域
CC: "color-class"  //状态颜色过滤器
DATE: "date-picker" //表单日期选择器
DTIME: "datetime-picker" //表单日期时间选择器,
FIXED: 列表表单固定列
SORT: 列表表单排序列
列表自定义索引， 约定给当前表添加一行，字段名为_el_table_index,约束为索引大小

//C,R,U,Q,L子约束
f:前端过滤器，L(f:过滤器参数)
e:枚举参数

//枚举级联  #字段名，组件约束和子约束都可以使用

//排除表 ^表名


## 一、基础信息表

###  1. 奖励信息[sc_reward_info]

| 字段名      | 类型         | 默认值  | 为空  |       约束       | 描述     |
| ----------- | ------------ | :-----: | :---: | :--------------: | :------- |
| rw_id       | number(10)   |  10000  |  否   |    PK,SEQ,l,r    | 编号     |
| rw_name     | varchar2(64) |         |  否   | DN,u,l,r,l,r,c,q | 名称     |
| category    | number(2)    |    0    |  否   |   q,u,l,r,c,sl   | 分类     |
| score       | number(10)   |    0    |  否   |    DI,u,l,r,c    | 分数     |
| status      | number(1)    |    0    |  否   | u,l,r,c,q,sl,cc  | 状态     |
| create_time | date         | sysdate |  是   |       l,r        | 创建时间 |


###  2. 惩罚信息[sc_punishment_info]

| 字段名      | 类型         | 默认值  | 为空  |      约束       | 描述     |
| ----------- | ------------ | :-----: | :---: | :-------------: | :------- |
| pn_id       | number(10)   |  2000   |  否   |   PK,SEQ,l,r    | 编号     |
| pn_name     | varchar2(64) |         |  否   | DN,c,u,r,l,r,q  | 名称     |
| category    | number(10)   |    0    |  否   |   u,l,r,c,sl    | 分类     |
| score       | number(10)   |    0    |  否   |   DI,c,u,l,r    | 分数     |
| status      | number(1)    |    0    |  否   | u,l,r,c,q,sl,cc | 状态     |
| create_time | date         | sysdate |  否   |       l,r       | 创建时间 |



###  3. 礼品信息[sc_gif_info]

| 字段名      | 类型         | 默认值  | 为空  |      约束       | 描述     |
| ----------- | ------------ | :-----: | :---: | :-------------: | :------- |
| gif_id      | number(10)   |  1000   |  否   |   SEQ,PK,l,r    | 编号     |
| gif_name    | varchar2(64) |         |  否   |  DN,c,u,r,l,q   | 礼品名称 |
| score       | number(10)   |    0    |  否   |   DI,c,u,r,l    | 分数     |
| status      | number(1)    |    0    |  否   | c,u,r,l,q,sl,cc | 状态     |
| create_time | date         | sysdate |  否   |       l,r       | 创建时间 |


###  4. 用户信息[sc_user_info]

| 字段名      | 类型          | 默认值  | 为空  |      约束      | 描述     |
| ----------- | ------------- | :-----: | :---: | :------------: | :------- |
| uid         | number(10)    |   100   |  否   | SEQ,PK,DI,l,r  | 用户编号 |
| name        | varchar2(64)  |         |  否   | DN,c,u,r,l,r,q | 姓名     |
| url         | varchar2(128) |         |  否   |    c,u,l,r     | 头像     |
| birthday    | date          |         |  否   |  DATE,c,u,l,r  | 出生日期 |
| score       | number(10)    |    0    |  否   |      r,l       | 分数     |
| status      | number(1)     |    0    |  否   |  c,u,r,l,q,sl  | 状态     |
| create_time | date          | sysdate |  否   |      l,r       | 创建时间 |


###  5. 分数记录[sc_score_record]

| 字段名      | 类型       | 默认值  | 为空  |          约束           | 描述     |
| ----------- | ---------- | :-----: | :---: | :---------------------: | :------- |
| rc_id       | number(10) |         |  否   |        PK,DI,SEQ        | 编号     |
| uid         | number(10) |         |  否   | DN,sl(sc_user_info),q,c | 用户     |
| c_tp        | number(1)  |         |  是   |     l,u,c,r,q,sl,c      | 类型     |
| score       | number(10) |         |  否   |   l,u,c,r,c,sl(#c_tp)   | 变动分数 |
| remain      | number(20) |    0    |  是   |            r            | 剩余     |
| create_time | date       | sysdate |  否   |           r,l           | 创建时间 |

### 6. 字典配置[sc_dict_info]

| 字段名  | 类型         | 默认值 | 为空  |       约束       | 描述   |
| ------- | ------------ | :----: | :---: | :--------------: | :----- |
| id      | number(10)   |        |  否   | PK,IS,SEQ,l,r,DI | 编号   |
| name    | varchar2(64) |        |  否   |   q,c,u,l,r,DN   | 名称   |
| value   | varchar2(32) |        |  否   |     c,u,l,r      | 值     |
| type    | varchar2(32) |        |  否   |    q,c,u,l,r     | 类型   |
| status  | number(1)    |   0    |  否   | q,c,u,l,r,cc,sl  | 状态   |
| sort_no | number(2)    |   0    |  否   |     c,u,l,r      | 排序值 |