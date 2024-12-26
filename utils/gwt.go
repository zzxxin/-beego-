package utils

import (
	beego "github.com/beego/beego/v2/server/web"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte("your_secret_key")

type Claims struct {
	UserID   uint   `json:"user_id"`
	UserName string `json:"user_name"`
	jwt.RegisteredClaims
}

// GenerateToken 生成JWT
func GenerateToken(userID uint, userName string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	jwt_user, _ := beego.AppConfig.String("login_jwt_token")
	claims := Claims{
		UserID:   userID,
		UserName: userName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			IssuedAt:  jwt.NewNumericDate(nowTime),
			Issuer:    jwt_user,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

// ParseToken 解析JWT
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
