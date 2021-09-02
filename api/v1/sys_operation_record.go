package v1

import (
	"github.com/gin-gonic/gin"
	"go-class/global"
	"go-class/model"
	"go-class/model/response"
	"go-class/service"
	"go.uber.org/zap"
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
