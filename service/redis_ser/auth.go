package redis_ser

import (
	"FastFiber_v2/global"
	"FastFiber_v2/utils/jwts"
	"context"
	"fmt"
	"go.uber.org/zap"
	"time"
)

func Logout(accessToken, refreshToken string) {
	atoken, err := jwts.CheckAccessToken(accessToken)
	rtoken, err := jwts.CheckRefreshToken(refreshToken)

	if err != nil {
		global.Log.Error("token验证失败", zap.Error(err))
		return
	}

	// 存储access_token黑名单
	accessKey := fmt.Sprintf("logout_access_%s", accessToken)
	// 存储refresh_token黑名单
	refreshKey := fmt.Sprintf("logout_refresh_%s", refreshToken)

	global.Redis.Set(context.Background(), accessKey, "", atoken.ExpiresAt.Sub(time.Now()))
	global.Redis.Set(context.Background(), refreshKey, "", rtoken.ExpiresAt.Sub(time.Now()))

}

func HashLogout(accessToken string) (ok bool) {
	accessKey := fmt.Sprintf("logout_access_%s", accessToken)

	_, err := global.Redis.Get(context.Background(), accessKey).Result()

	if err != nil {
		return false
	}
	return true
}
