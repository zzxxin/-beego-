package filters

import (
	"github.com/beego/beego/v2/server/web/context"
	"net/http"
)

// SignMiddleware 签权校验中间件
func SignMiddleware(ctx *context.Context) {
	// 示例签权校验逻辑
	signature := ctx.Input.Query("signature")
	if signature == "" || !isValidSignature(signature) {
		// 如果签名无效，返回 403
		ctx.ResponseWriter.WriteHeader(http.StatusForbidden)
		ctx.ResponseWriter.Write([]byte("签名无效"))
		return
	}
}

// 模拟签名校验逻辑
func isValidSignature(signature string) bool {
	// 这里可以实现实际的签权逻辑，比如对请求参数进行签名校验
	// 此处为简化示例，假设只要 signature 不为空则校验通过
	return signature == "valid_signature"
}
