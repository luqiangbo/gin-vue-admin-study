package v1

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"goClass/global"
	"goClass/model"
	"goClass/model/response"
	"goClass/service"
)

func CreateSysOperationRecord(c *gin.Context) {
	var sysOperationRecord model.SysOperationRecord
	_ = c.ShouldBindJSON(&sysOperationRecord)
	if err := service.CreateSysOperationRecord(sysOperationRecord); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}
