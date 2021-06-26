package service

import (
	"goClass/global"
	"goClass/model"
)

func CreateSysOperationRecord(sysOperationRecord model.SysOperationRecord) (err error) {
	err = global.GVA_DB.Create(&sysOperationRecord).Error
	return err
}
