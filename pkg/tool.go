package pkg

import (
	"beegoweb/utils"
	"html/template"
	"strconv"
	"time"

	beego "github.com/beego/beego/v2/server/web"
)

// Init 工具类初始化 例如分页等
func Init() {
	// 注册全局模板函数，确保所有模板都可以使用
	beego.AddFuncMap("add", utils.Add)
	beego.AddFuncMap("sub", utils.Sub)
	beego.AddFuncMap("inArray", utils.InArray)
	beego.AddFuncMap("itoa", strconv.Itoa)
	// 注册 str 函数
	beego.AddFuncMap("str", utils.Str)
	// 注册 PageLinks 函数
	beego.AddFuncMap("PageLinks", func(pagination *utils.Pagination) template.HTML {
		return pagination.PageLinks() // 直接返回 template.HTML
	})
	// 注册格式化时间的函数
	beego.AddFuncMap("FormatTime", func(t time.Time) string {
		return t.Format("2006-01-02 15:04:05")
	})
	// 注册时间戳格式化的函数
	beego.AddFuncMap("FormatTimestamp", func(timestamp int64) string {
		t := time.Unix(timestamp, 0)
		return t.Format("2006-01-02 15:04:05")
	})

	beego.AddFuncMap("PageLinks", PageLinks)

}

// PageLinks 全局函数，用于获取分页链接
func PageLinks(pagination *utils.Pagination) template.HTML {
	return pagination.PageLinks()
}
