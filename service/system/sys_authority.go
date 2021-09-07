package system

import (
	"errors"
	"go-class/global"
	"go-class/model/system/tables"
	"gorm.io/gorm"
)

type AuthorityService struct {
}

var AuthorityServiceApp = new(AuthorityService)

func (authorityService *AuthorityService) CreateAuthority(props tables.SysAuthority) (err error, authority tables.SysAuthority) {
	var m tables.SysAuthority
	if !errors.Is(global.GVA_DB.Where("authority_id = ?", props.AuthorityId).First(&m).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同角色id"), props
	}
	err = global.GVA_DB.Create(&props).Error
	return err, props
}
