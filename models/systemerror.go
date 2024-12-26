package models

import (
	"github.com/beego/beego/v2/client/orm"
	"time"
)

// 定义 bk_system_error 表的结构体
type SystemError struct {
	Id         int    `orm:"auto"`
	Msg        string `orm:"type(text)"`   // 错误信息
	Uri        string `orm:"size(200)"`    // 路由
	Env        string `orm:"size(20)"`     // 环境变量
	Params     string `orm:"type(text)"`   // 请求参数
	Method     string `orm:"size(30)"`     // 请求方式
	CreateTime int    `orm:"auto_now_add"` // 创建时间
	UpdateTime int    `orm:"auto_now"`     // 修改时间
}

// 初始化数据库表的映射
func init() {
	orm.RegisterModel(new(SystemError))
}

// 保存错误日志到数据库
func LogError(msg, uri, env, params, method string) error {
	o := orm.NewOrm()
	errorLog := SystemError{
		Msg:        msg,
		Uri:        uri,
		Env:        env,
		Params:     params,
		Method:     method,
		CreateTime: int(time.Now().Unix()), // 创建时间为当前时间戳
		UpdateTime: int(time.Now().Unix()), // 修改时间为当前时间戳
	}
	_, err := o.Insert(&errorLog)
	return err
}
