package services

import (
	"beegoweb/models"
	"time"
)

type SystemSetService struct{}

func (s SystemSetService) HandleSystemSet(type_name string, status int) (string, error) {

	// 查询是否已经添加过 没有的话添加 添加过的话修改
	setting, _ := models.GetSettingByQuery("type = ?", type_name)
	if setting != nil {
		fields := map[string]interface{}{
			"status": status,
		}
		err := models.UpdateSetting(fields, "type = ?", type_name)
		if err != nil {
			return "修改失败", err
		}
	} else {
		addData := &models.SystemSetting{
			Type:       type_name,
			Status:     status,
			CreateTime: int(time.Now().Unix()),
		}
		if err := models.CreateSetting(addData); err != nil {
			return "添加失败", err
		}
	}
	return "", nil
}
