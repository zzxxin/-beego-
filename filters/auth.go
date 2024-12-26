package filters

import (
	"beegoweb/models"
	"beegoweb/utils"
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"net/http"
	"strconv"
	"strings"
)

// AuthMiddleware 中间件函数：检查用户是否登录并进行重定向
func AuthMiddleware(ctx *context.Context) {
	// 检查请求路径，如果是登录页面则不进行拦截
	if ctx.Request.RequestURI == "/user_login" || ctx.Request.RequestURI == "/do_login" {
		return
	}
	user_login, _ := web.AppConfig.String("login_identification")
	// 从 Cookie 中获取 Token
	tokenString := ctx.GetCookie(user_login)

	// 如果 Token 不存在
	if tokenString == "" {
		ctx.Redirect(302, "/user_login")
		return
	}

	// 解析和验证 Token
	claims, err := utils.ParseToken(tokenString)
	if err != nil {
		ctx.Redirect(302, "/user_login")
		return
	}
	// 获取用户权限
	GetUserRight(ctx, claims.UserID)
	ctx.Input.SetData("userId", claims.UserID)

}

// GetUserRight 获取用户权限数据并存储到 ctx 中
func GetUserRight(ctx *context.Context, userId uint) error {
	// 获取当前路由
	currentRoute := ctx.Input.URL()
	trimmedRoute := strings.TrimPrefix(currentRoute, "/")
	// 获取用户信息
	userInfo, err := models.GetUserByID(userId)
	if err != nil {
		return err
	}

	var rightData []models.AdminRights

	// 获取全部权限
	allRight, _ := models.GetAllRights("1=1")

	// 定义一个数组用来存储 RightLogo 字段的值
	var rightLogos []string

	// 遍历 allRight，提取 RightLogo 字段
	for _, right := range allRight {
		// 假设 RightLogo 是一个字符串字段
		rightLogos = append(rightLogos, right.RightLogo)
	}

	// roleListData, _ := json.MarshalIndent(rightLogos, "", "   ")
	// fmt.Printf("全部的权限信息为：%s\n", roleListData)

	if userInfo.IsSuper == "Y" {
		// 超级管理员，获取全部权限
		rightData, err = models.GetAllRights("is_menu = ?", 1)
		if err != nil {
			return err
		}
	} else {
		// 普通用户，根据角色获取权限
		roleIds := strings.Split(userInfo.RoleIDs, ",")
		var roleList []models.AdminRole
		// 根据角色ID获取角色信息
		err := models.GetRolesByIDs(roleIds, &roleList)
		if err != nil {
			return err
		}

		var rightIds []uint
		for _, role := range roleList {
			var roleRights []string // 先将 rights 解析为字符串数组

			// 解析 rights 为 []string
			if err := json.Unmarshal([]byte(role.Rights), &roleRights); err != nil {
				fmt.Printf("解析 rights 字段时出错: %v\n", err)
				return err
			}

			// 将字符串数组转换为 uint 数组
			for _, right := range roleRights {
				rightID, _ := strconv.ParseUint(right, 10, 64) // 将 string 转换为 uint
				rightIds = append(rightIds, uint(rightID))
			}
		}

		rightData, err = models.GetRightsByIDs(rightIds)
		if err != nil {
			return err
		}

		// 如果路由配置了 并且在普通用户的权限内 则有权限否则 返回异常
		var userRightLogos []string

		// 遍历 allRight，提取 RightLogo 字段
		for _, right := range rightData {
			// 假设 RightLogo 是一个字符串字段
			userRightLogos = append(userRightLogos, right.RightLogo)
		}

		// right, _ := json.MarshalIndent(userRightLogos, "", "   ")
		// fmt.Printf("普通用户的权限数据为：%s\n", right)
		// 当前权限在配置校验权限路由中
		if utils.InArray(trimmedRoute, rightLogos) == true {
			if utils.InArray(trimmedRoute, userRightLogos) != true {
				ctx.ResponseWriter.WriteHeader(http.StatusForbidden)
				ctx.ResponseWriter.Write([]byte(`
        <html>
            <head>
                <script type="text/javascript">
                    alert("您没有权限访问该页面！");
                    window.history.back();
                </script>
            </head>
            <body></body>
        </html>
    `))
				return nil
			}
		}
	}

	// 构建用户权限映射
	userRight := make(map[string][]models.AdminRights)
	for _, right := range rightData {
		userRight[right.GroupName] = append(userRight[right.GroupName], right)
	}

	// 将用户权限数据存储到上下文中
	ctx.Input.SetData("userRight", map[string]interface{}{
		"right_list": userRight,
	})

	return nil
}
