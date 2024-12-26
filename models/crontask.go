package models

import (
	"beegoweb/pkg/db"
	"errors"
	"time"

	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
)

type CronTask struct {
	ID             int          `gorm:"primaryKey;autoIncrement;comment:主键ID"`
	TaskName       string       `gorm:"type:varchar(255);not null;comment:任务名称"`
	CronExpression string       `gorm:"type:varchar(255);not null;comment:Cron 表达式"`
	TaskStatus     int          `gorm:"type:tinyint(1);default:1;comment:任务状态：1=启用, 0=禁用"`
	TaskDesc       string       `gorm:"type:varchar(255);default:'';comment:任务描述"`
	CreatedAt      time.Time    `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdatedAt      time.Time    `gorm:"type:timestamp;default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP;comment:更新时间"`
	EntryID        cron.EntryID `json:"-"` // 用于存储 Cron 调度器中的 EntryID
}

// 定义表名
func (CronTask) TableName() string {
	return "bk_cron_tasks"
}

func GetCronByQueryPage(page int, pageSize int, where string, args ...interface{}) ([]CronTask, int64, error) {
	var cron []CronTask
	var total int64

	// 计算偏移量
	offset := (page - 1) * pageSize

	// 查询符合条件的总记录数
	err := db.DB.Model(&CronTask{}).Where(where, args...).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 获取分页数据
	result := db.DB.Where(where, args...).Order("id desc").Limit(pageSize).Offset(offset).Find(&cron)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	return cron, total, nil
}

// 添加cron
func CreateCron(cron *CronTask) error {
	return db.DB.Create(cron).Error
}

// GetCronByQuery 根据条件查询用户
func GetCronByQuery(where string, args ...interface{}) (*CronTask, error) {
	var cron CronTask
	result := db.DB.Where(where, args...).First(&cron)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &cron, result.Error
}

// 修改cron
func UpdateCron(fields map[string]interface{}, where string, args ...interface{}) error {
	// 使用 GORM 的 Updates 方法，只更新传递的字段
	return db.DB.Model(&CronTask{}).Where(where, args...).Updates(fields).Error
}

func UpdateCronStatus(id int64, fieldName string, fieldValue interface{}) error {
	// 调用模型层更新字段
	return db.DB.Model(&CronTask{}).Where("id = ?", id).UpdateColumn(fieldName, fieldValue).Error
}
