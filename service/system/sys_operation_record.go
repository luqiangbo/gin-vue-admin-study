package system

import (
	"go-class/global"
	"go-class/model/system/tables"
)

type OperationRecordService struct {
}

// 创建记录

func (operationRecordService *OperationRecordService) CreateSysOperationRecord(sysOperationRecord tables.SysOperationRecord) (err error) {
	err = global.GVA_DB.Create(&sysOperationRecord).Error
	return err
}
