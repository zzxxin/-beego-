package utils

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/beego/beego/v2/server/web"
	"time"
)

// GenerateMD5Hash 生成字符串的MD5哈希值
func GenerateMD5Hash(text string) string {
	hash := md5.New()
	hash.Write([]byte(text))
	return hex.EncodeToString(hash.Sum(nil))
}

// JSONResponse 封装 JSON 返回
func JSONResponse(controller *web.Controller, code int, msg string, data interface{}) {
	response := map[string]interface{}{
		"code": code,
		"msg":  msg,
		"data": data,
	}
	controller.Data["json"] = response
	controller.ServeJSON()
}

func FormatTimestamp(t time.Time) string {
	return t.Format("2006-01-02T15:04:05-07:00")
}
