package system

import (
	"github.com/gin-gonic/gin"
	"gin-vue-admin-study/global"
	modelCommonRes "gin-vue-admin-study/model/common/response"
	modelSystemReq "gin-vue-admin-study/model/system/request"
	"go.uber.org/zap"
)

type DBApi struct {
}

// 初始化数据库

func (s *DBApi) InitDB(c *gin.Context) {
	if global.GVA_DB != nil {
		global.GVA_LOG.Error("已存在数据库配置")
		modelCommonRes.FailWithMessage("已存在数据库配置", c)
		return
	}
	var req modelSystemReq.InitDB

	if err := c.ShouldBindJSON(&req); err != nil {
		global.GVA_LOG.Error("参数校验不通过", zap.Any("err", err))
		modelCommonRes.FailWithMessage("1自动创建数据库失败, 请查看后台日志, 检查后进行初始化", c)
		return
	}
	if err := initDBService.InitDB(req); err != nil {
		global.GVA_LOG.Error("自动创建数据库失败!", zap.Any("err", err))
		modelCommonRes.FailWithMessage("0自动创建数据库失败, 请查看后台日志, 检查后在进行初始化", c)
		return
	}
	modelCommonRes.OkWithData("自动创建数据成功", c)
}
