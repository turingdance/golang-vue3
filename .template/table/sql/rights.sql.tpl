

-- ----------------------------
-- 获得最大sortIndexMax
-- ----------------------------
SELECT @sortIndexMax := (select max(sort_index)+0.01 from sys_rights where type = 'group' and sort_index < 1000);
-- ----------------------------
-- {{.App.Title}}({{.App.Name}})不存在则创建
-- ----------------------------
INSERT INTO sys_rights(type,uri,title,biz,pid,redirect,component,sort_index) 
SELECT 'group','/{{.App.Name|lcfirst}}','{{.App.Title}}','{{.App.Name|lcfirst}}',0, '/{{.App.Name|lcfirst}}/{{.Module|lcfirst}}','Layout',@sortIndexMax
from dual 
WHERE NOT EXISTS(SELECT type,title FROM sys_rights where type='group' and biz = '{{.App.Name|lcfirst}}' );
-- ----------------------------
-- 获取应用ID
-- ----------------------------
SELECT @appId := (select id from sys_rights where biz = '{{.App.Name|lcfirst}}');

-- ----------------------------
-- {{.Title}}({{.Module}})不存在则创建
-- ----------------------------
insert into sys_rights (type,component,uri,icon,title,pid,biz,redirect,sort_index,hidden)
select 'view',  '{{.App.Name|lcfirst}}/{{.Module|lcfirst}}/index', '/{{.App.Name|lcfirst}}/{{.Module|lcfirst}}', 'cascader', '{{.Title}}',  @appId, '{{.App.Name|lcfirst}}:{{.Module|lcfirst}}','', 1, 0
from dual 
WHERE 
NOT EXISTS(SELECT type,title FROM sys_rights where type='view' and biz = '{{.App.Name|lcfirst}}:{{.Module|lcfirst}}' );
-- ----------------------------
-- 获取模块ID
-- ----------------------------
SELECT @parentId := (select id from sys_rights where type='view' and biz = '{{.App.Name|lcfirst}}:{{.Module|lcfirst}}');

-- ----------------------------
-- 搜索
-- ----------------------------
insert into sys_rights (type,title,uri, icon , pid,   biz,redirect,sort_index,hidden)
select 'widget','{{.Title}}搜索','/{{.App.Name|lcfirst}}/{{.Module|lcfirst}}/search',  'search', @parentId,'{{.App.Name|lcfirst}}:{{.Module|lcfirst}}:search','',1,0
from dual 
WHERE NOT EXISTS(select * from sys_rights where biz='{{.App.Name|lcfirst}}:{{.Module|lcfirst}}:search' and pid=@parentId);

-- ----------------------------
-- 详情
-- ----------------------------
insert into sys_rights (type,title,uri, icon, pid,   biz,redirect,sort_index,hidden)
select 'widget','{{.Title}}详情','get:/{{.App.Name|lcfirst}}/{{.Module|lcfirst}}/{pkId}',  'tickets',@parentId,'{{.App.Name|lcfirst}}:{{.Module|lcfirst}}:getOne','',2,0 
from dual  
WHERE NOT EXISTS(select * from sys_rights where biz='{{.App.Name|lcfirst}}:{{.Module|lcfirst}}:getOne' and pid=@parentId);

-- ----------------------------
-- 增加
-- ----------------------------
insert into sys_rights (type,title,uri, icon, pid,   biz,redirect,sort_index,hidden)
select 'widget','{{.Title}}新增','post:/{{.App.Name|lcfirst}}/{{.Module|lcfirst}}',  'plus', @parentId,'{{.App.Name|lcfirst}}:{{.Module|lcfirst}}:create','',3,0 
from dual 
WHERE NOT EXISTS(select * from sys_rights where biz='{{.App.Name|lcfirst}}:{{.Module|lcfirst}}:create' and pid=@parentId);

-- ----------------------------
-- 修改
-- ----------------------------
insert into sys_rights (type,title,uri, icon, pid,   biz,redirect,sort_index,hidden)
select 'widget','{{.Title}}修改','put:/{{.App.Name|lcfirst}}/{{.Module|lcfirst}}',  'edit', @parentId,'{{.App.Name|lcfirst}}:{{.Module|lcfirst}}:update','',4,0
from dual 
WHERE NOT EXISTS(select * from sys_rights where biz='{{.App.Name|lcfirst}}:{{.Module|lcfirst}}:update' and pid=@parentId);

-- ----------------------------
-- 删除
-- ----------------------------
insert into sys_rights (type,title,uri, icon , pid,   biz,redirect,sort_index,hidden)
select 'widget','{{.Title}}删除单条','delete:/{{.App.Name|lcfirst}}/{{.Module|lcfirst}}/{pkId}',  'delete', @parentId,'{{.App.Name|lcfirst}}:{{.Module|lcfirst}}:delete','',5,0
from dual 
WHERE NOT EXISTS(select * from sys_rights where biz='{{.App.Name|lcfirst}}:{{.Module|lcfirst}}:delete' and pid=@parentId);

-- ----------------------------
-- 批量删除
-- ----------------------------
insert into sys_rights (type,title,uri, icon, pid,   biz,redirect,sort_index,hidden)
select 'widget','{{.Title}}批量删除','delete:/{{.App.Name|lcfirst}}/{{.Module|lcfirst}}',  'delete', @parentId,'{{.App.Name|lcfirst}}:{{.Module|lcfirst}}:deleteIts','',5,0
from dual 
WHERE NOT EXISTS(select * from sys_rights where biz='{{.App.Name|lcfirst}}:{{.Module|lcfirst}}:delete' and pid=@parentId);

-- ----------------------------
-- 导出
-- ----------------------------
insert into sys_rights (type,title,uri, icon , pid,   biz,redirect,sort_index,hidden)
select 'widget','{{.Title}}导出','get,post:/{{.App.Name|lcfirst}}/{{.Module|lcfirst}}/export',  'export',@parentId,'{{.App.Name|lcfirst}}:{{.Module|lcfirst}}:export','',7,0
from dual 
WHERE NOT EXISTS(select * from sys_rights where biz='{{.App.Name|lcfirst}}:{{.Module|lcfirst}}:export' and pid=@parentId);

