package routers

import (
	"beegoweb/controllers"
	"beegoweb/controllers/api"
	"beegoweb/controllers/role"
	"beegoweb/controllers/socket"
	"beegoweb/controllers/user"
	"beegoweb/controllers/web"
	"beegoweb/filters"
	"strings"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

// 判断是否是API请求
func isAPIRequest(ctx *context.Context) bool {
	// 使用 ctx.Input.URL() 来获取请求的路径
	path := ctx.Input.URL()
	return strings.HasPrefix(path, "/api")
}

func init() {
	// 注册全局中间件，应用于所有路由
	beego.InsertFilter("/*", beego.BeforeRouter, func(ctx *context.Context) {
		// 如果是API请求，则使用 SignMiddleware（签权中间件）
		if isAPIRequest(ctx) {
			filters.SignMiddleware(ctx)
		} else {
			// 否则使用 AuthMiddleware（登录校验中间件）
			filters.AuthMiddleware(ctx)
		}
	})

	// 错误处理中间件，应该最后执行，捕获所有未处理的错误
	beego.InsertFilter("/*", beego.BeforeRouter, filters.ErrorHandlingMiddleware)

	// 为 api 路由使用 SignMiddleware（签权校验）
	beego.InsertFilter("/api/*", beego.BeforeRouter, filters.SignMiddleware)

	// 默认路径
	beego.Router("/", &web.IndexController{}, "*:Index")
	// 首页
	beego.Router("/web", &web.IndexController{}, "*:Webmain")

	// 登录页
	beego.Router("/user_login", &controllers.LoginController{}, "*:UserLogin")
	beego.Router("/do_login", &controllers.LoginController{}, "*:DoLogin")
	beego.Router("/login_out", &controllers.LoginController{}, "*:LoginOut")

	/******************** 用户操作 start **************************/
	beego.Router("/user_add", &user.UserController{}, "*:AddUser")
	beego.Router("/user_list", &user.UserController{}, "*:UserList")
	beego.Router("/user_edit", &user.UserController{}, "*:UserEdit")
	// 测试导出功能
	beego.Router("/user_export", &user.UserController{}, "*:UserExport")
	// 分配角色
	beego.Router("/allot_role", &user.UserController{}, "*:AllotRole")
	beego.Router("/user_status", &user.UserController{}, "*:UserStatus")

	/******************** 用户操作 end **************************/

	/******************** 角色操作 start **************************/
	beego.Router("/add_role", &role.RoleController{}, "*:AddRole")
	beego.Router("/role_list", &role.RoleController{}, "*:RoleList")
	beego.Router("/up_role", &role.RoleController{}, "*:UpRole")
	beego.Router("/role_bind_right", &role.RoleController{}, "*:RoleBindRight")
	/******************** 角色操作 end **************************/

	/******************** 权限操作 start **************************/
	beego.Router("/right_list", &role.RightController{}, "*:RightList")
	beego.Router("/add_right", &role.RightController{}, "*:AddRight")
	beego.Router("/up_right", &role.RightController{}, "*:UpRight")

	/******************** 权限操作 end **************************/

	/******************** 计划任务 start **************************/
	beego.Router("/cron_list", &controllers.CronController{}, "*:CronList")
	beego.Router("/add_cron", &controllers.CronController{}, "*:AddCron")
	beego.Router("/up_cron", &controllers.CronController{}, "*:UpCron")

	/******************** 计划任务 end **************************/

	/******************** api结构 start **************************/
	// 走路由签权
	beego.Router("/api/test_api", &api.TestController{}, "*:TestApi")
	beego.Router("/api/test_api_mq", &api.TestController{}, "*:QueueMq")

	/******************** api结构 end **************************/

	/******************** 聊天室 start **************************/
	// 做websocket 链接
	beego.Router("/ws", &socket.WebSocketController{})
	// 用户打开聊天室  可走私聊 、 群聊模式
	beego.Router("/chatroom", &socket.ChatController{}, "*:ChatRoom")
	beego.Router("/chat/get_channel", &socket.ChatController{}, "*:GetChannel")
	beego.Router("/chat/get_messages", &socket.ChatController{}, "*:GetMessages")

	/******************** 聊天室 end **************************/

	/** 系统的一些配置开关 start **/
	beego.Router("/cron_set", &controllers.SystemSetController{}, "*:CronSet")

	/** 系统的一些配置开关 end **/
}
