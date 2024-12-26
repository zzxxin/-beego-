package filters

import (
	"beegoweb/models"
	"fmt"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

// 全局异常处理中间件
var ErrorHandlingMiddleware = func(ctx *context.Context) {
	defer func() {
		if err := recover(); err != nil {
			// 捕获系统级错误，获取相关的请求信息
			request := ctx.Request
			msg := fmt.Sprintf("System error: %v", err)
			uri := request.RequestURI                 // 当前请求的 URI
			env, _ := web.AppConfig.String("runmode") // 可以通过配置文件或环境变量获取
			params := fmt.Sprintf("%v", request.Form) // 获取请求的参数
			method := request.Method                  // 请求的 HTTP 方法

			// 调用 LogSystemError 保存错误日志到数据库
			if logErr := models.LogError(msg, uri, env, params, method); logErr != nil {
				// 如果日志记录失败，输出到控制台
				fmt.Printf("Failed to log system error: %v\n", logErr)
			}
			// 返回友好的错误提示，可以自定义一个统一的错误页面
			ctx.ResponseWriter.WriteHeader(500)
			ctx.ResponseWriter.Write([]byte("Internal Server Error. Please try again later."))
			// 也可以根据需求记录到文件或者发送报警通知
		}
	}()
}
