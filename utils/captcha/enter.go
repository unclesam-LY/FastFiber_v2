package captcha

import (
	"FastFiber_v2/global"
	"context"
	"time"
)

var CaptChaStore = NewRedisStore()

type RedisStore struct {
	ctx context.Context
}

func NewRedisStore() *RedisStore {
	return &RedisStore{ctx: context.Background()}
}

func (s *RedisStore) Set(id string, value string) error {
	return global.Redis.Set(s.ctx,
		"captcha:"+id,
		value,
		time.Minute*5, // 验证码5分钟有效
	).Err()
}

func (s *RedisStore) Get(id string, clear bool) string {
	key := "captcha:" + id
	val, _ := global.Redis.Get(s.ctx, key).Result()
	if clear {
		global.Redis.Del(s.ctx, key)
	}
	return val
}

func (s *RedisStore) Verify(id, answer string, clear bool) bool {
	return answer == s.Get(id, clear)
}
