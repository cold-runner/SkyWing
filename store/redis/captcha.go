package redis

import (
	"go.uber.org/zap"
	"time"
)

func (r *Rdb) Set(id string, value string) error {
	if err := r.Client.Set(id, value, time.Minute*3).Err(); err != nil {
		zap.L().Error("存储验证码失败", zap.Error(err))
		return err
	}
	return nil
}
func (r *Rdb) Get(id string, clear bool) string {
	val, err := r.Client.Get(id).Result()
	if err != nil {
		zap.L().Error("获取验证码失败或已过期", zap.Error(err))
		return ""
	}
	if clear {
		err = r.Client.Del(id).Err()
		if err != nil {
			zap.L().Error("值已过期", zap.Error(err))
			return ""
		}
	}
	return val
}
func (r *Rdb) Verify(id, answer string, clear bool) bool {
	storeData := r.Get(id, clear)
	if answer != storeData {
		return false
	}
	return true
}
