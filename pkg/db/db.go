package db

import (
	"fmt"
	"log"
	"time"

	"github.com/beego/beego/v2/server/web"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Init() {
	user := web.AppConfig.DefaultString("user", "root")
	password := web.AppConfig.DefaultString("password", "123456")
	host := web.AppConfig.DefaultString("host", "127.0.0.1")
	port := web.AppConfig.DefaultString("port", "3306")
	dbname := web.AppConfig.DefaultString("dbname", "beegoweb")
	charset := web.AppConfig.DefaultString("charset", "utf8mb4")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		user, password, host, port, dbname, charset)

	var err error
	// 从 app.conf 中读取是否启用 SQL 日志配置
	slowThreshold := web.AppConfig.DefaultInt("slow_threshold", 1) // 默认慢查询阈值 1 秒

	// 配置 GORM 的日志记录器
	newLogger := logger.New(
		log.New(log.Writer(), "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Duration(slowThreshold) * time.Second,
			LogLevel:      logger.Error, // 使用更详细的日志级别
			Colorful:      true,
		},
	)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger, // 设置自定义的日志记录器
	})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	fmt.Printf("Connected to Mysql \n")

}
