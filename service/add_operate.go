package services

import (
	"beegoweb/models"
	"encoding/json"
	"strconv"
	"time"

	"github.com/beego/beego/v2/server/web/context"
)

// AddOperateLog 添加操作日志到数据库，自动处理请求、响应、路由和方法
func AddOperateLog(ctx *context.Context, desc string, responseData interface{}, logType int) error {

	if logType == 0 {
		logType = 1
	}
	// 获取请求体，使用 RequestBody 确保请求体可以重复读取
	// 解析表单数据
	err := ctx.Request.ParseForm()
	if err != nil {
		// 解析失败返回错误
		return err
	}
	formData := ctx.Request.Form

	// 将表单数据转换为 JSON 字符串
	formJson, err := json.Marshal(formData)
	if err != nil {
		formJson = []byte("无法序列化的表单数据")
	}

	// 获取响应数据
	respJson, err := json.Marshal(responseData)
	if err != nil {
		// 如果输出的响应无法序列化为 JSON，记录错误信息，但不阻塞日志插入
		respJson = []byte("无法序列化的响应")
	}

	// 获取当前路由和方法名
	router := ctx.Input.URL()
	method := ctx.Input.Method()

	var operator uint
	if op, ok := ctx.Input.GetData("userId").(uint); ok {
		operator = op
	}
	// 构造日志记录
	var operateLog = &models.OperatorLog{
		Desc:     desc,
		Operator: strconv.FormatUint(uint64(operator), 10),
		// Operator:   string(operator),
		Request:    string(formJson), // 记录请求参数
		Response:   string(respJson), // 记录响应内容
		Router:     router,
		Method:     method,
		Type:       logType,
		CreateTime: time.Now().Unix(),
	} // 使用 GORM 插入到数据库
	if err := models.CreateOperatorLog(operateLog); err != nil {
		return err
	}

	return nil
}
