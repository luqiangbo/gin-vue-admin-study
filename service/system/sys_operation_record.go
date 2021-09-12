package system

import (
	"gin-vue-admin-study/global"
	"gin-vue-admin-study/model/system/tables"
)

type OperationRecordService struct {
}

// 创建记录

func (operationRecordService *OperationRecordService) CreateSysOperationRecord(sysOperationRecord tables.SysOperationRecord) (err error) {
	err = global.GVA_DB.Create(&sysOperationRecord).Error
	return err
}
