package jwts

import (
	"FastFiber_v2/global"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Claims struct {
	UserID uint `json:"user_id"`
	RoleID uint `json:"role_id"`
}

type RefreshClaims struct {
	UserID uint `json:"user_id"`
}

type AccessPayload struct {
	Claims
	jwt.RegisteredClaims
}

type RefreshPayload struct {
	RefreshClaims
	jwt.RegisteredClaims
}

// GenAccessToken 生成Access Token
func GenAccessToken(data Claims) (string, error) {
	claims := AccessPayload{
		Claims: data,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(
				time.Duration(global.Config.Jwt.AccessExpire) * time.Hour,
			)),
			Issuer: global.Config.Jwt.Issuer,
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).
		SignedString([]byte(global.Config.Jwt.AccessSecret))
}

// GenRefreshToken 生成Refresh Token
func GenRefreshToken(data RefreshClaims) (string, error) {
	claims := RefreshPayload{
		RefreshClaims: data,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(
				time.Duration(global.Config.Jwt.RefreshExpire) * time.Hour,
			)),
			Issuer: global.Config.Jwt.Issuer + "_refresh",
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).
		SignedString([]byte(global.Config.Jwt.RefreshSecret))
}

// CheckAccessToken 验证访问令牌
func CheckAccessToken(token string) (*AccessPayload, error) {
	return parseAccessToken(token, global.Config.Jwt.AccessSecret)
}

// CheckRefreshToken 验证刷新令牌
func CheckRefreshToken(token string) (*RefreshPayload, error) {
	return parseRefreshToken(token, global.Config.Jwt.RefreshSecret)
}

// parseToken 验证token
func parseAccessToken(token, secret string) (*AccessPayload, error) {
	tokenObj, err := jwt.ParseWithClaims(token, &AccessPayload{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := tokenObj.Claims.(*AccessPayload); ok && tokenObj.Valid {
		return claims, nil
	} else {
		return nil, errors.New("token无效")
	}

}

func parseRefreshToken(token, secret string) (*RefreshPayload, error) {
	tokenObj, err := jwt.ParseWithClaims(token, &RefreshPayload{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := tokenObj.Claims.(*RefreshPayload); ok && tokenObj.Valid {
		return claims, nil
	} else {
		return nil, errors.New("token无效")
	}

}
