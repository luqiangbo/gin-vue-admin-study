package service

import (
	"goClass/global"
	"goClass/model"
	"time"
)

func JsonInBlacklist(jwtList model.JwtBlacklist) (err error) {
	err = global.GVA_DB.Create(&jwtList).Error
	return
}

func GetRedisJWT(username string) (err error, redisJWT string) {
	redisJWT, err = global.GVA_REDIS.Get(username).Result()
	return err, redisJWT
}

func SetRedisJWT(jwt string, userName string) (err error) {
	// 此处过期时间等于jwt过期时间
	timer := time.Duration(global.GVA_CONFIG.JWT.ExpiresTime) * time.Second
	err = global.GVA_REDIS.Set(userName, jwt, timer).Err()
	return err
}
