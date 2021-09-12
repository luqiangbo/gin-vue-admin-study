package system

import (
	"github.com/gin-gonic/gin"
	"gin-vue-admin-study/global"
	"gin-vue-admin-study/model/common/response"
	"gin-vue-admin-study/model/system/tables"
	"go.uber.org/zap"
)

type OperationRecordApi struct {
}

// 创建记录

func (s *OperationRecordApi) CreateSysOperationRecord(c *gin.Context) {
	var sysOperationRecord tables.SysOperationRecord
	_ = c.ShouldBindJSON(&sysOperationRecord)
	if err := operationRecordService.CreateSysOperationRecord(sysOperationRecord); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}
