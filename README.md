当前实现了 用户的登录、以及基础的权限管理功能，通过给角色分配权限  用户分配角色的方式来限制用户的一些权限操作内容
权限校验为了方便处理  如果配置了需要校验权限的内容则处理权限校验 如果没有在系统中添加权限路由信息则不会对路由做权限校验


项目中不包含业务逻辑

## **项目目录说明如下**

1. 数据库配置内容再conf中
2. filters 为中间件 在router中注册 beego.InsertFilter("/*", beego.BeforeRouter, filters.AuthMiddleware)
3. models 为数据库查询
4. pkg 目录为 工具类 数据库初始化 以及基础注册函数
5. service  如果需要做分离  可以把数据处理放到service中 如果不需要的话 直接在controller中处理也可
6. static 为bootstarp引用的css 和js
7. unils 文件夹中包含的是一些通用的方法  例如导出 gwt 等等
8. views 为前端的页面 后缀tpl
9. cron 操作 需要在页面中添加执行的周期和方法  在main.go中的注册cron任务 把cron的实际执行全部都放到jobs目录下 方便后续管理