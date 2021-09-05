package system

import (
	"go-class/global"
	"go-class/model/system"
)

type OperationRecordService struct {
}

// 创建记录

func (operationRecordService *OperationRecordService) CreateSysOperationRecord(sysOperationRecord system.SysOperationRecord) (err error) {
	err = global.GVA_DB.Create(&sysOperationRecord).Error
	return err
}
