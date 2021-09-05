package system

import (
	"errors"
	"go-class/global"
	"go-class/model/system"
	"gorm.io/gorm"
	"time"
)

type JwtService struct {
}

func (JwtService *JwtService) JsonInBlacklist(jwtList system.JwtBlacklist) (err error) {
	err = global.GVA_DB.Create(&jwtList).Error
	return
}

func (JwtService *JwtService) IsBlacklist(jwt string) bool {
	err := global.GVA_DB.Where("jwt = ?", jwt).First(&system.JwtBlacklist{}).Error
	isNotFound := errors.Is(err, gorm.ErrRecordNotFound)
	return !isNotFound

}

func (JwtService *JwtService) GetRedisJWT(username string) (err error, redisJWT string) {
	redisJWT, err = global.GVA_REDIS.Get(username).Result()
	return err, redisJWT
}

func (JwtService *JwtService) SetRedisJWT(jwt string, userName string) (err error) {
	// 此处过期时间等于jwt过期时间
	timer := time.Duration(global.GVA_CONFIG.JWT.ExpiresTime) * time.Second
	err = global.GVA_REDIS.Set(userName, jwt, timer).Err()
	return err
}
