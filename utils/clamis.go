package utils

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gin-vue-admin-study/global"
	"gin-vue-admin-study/model/system/request"
)

func GetUserId(c *gin.Context) uint {
	if claims, exists := c.Get("claims"); !exists {
		global.GVA_LOG.Error("从gin的Context中获取从jwt解析出来的用户ID失败, 请检查路由是佛使用jwt中间件")
		return 0
	} else {
		waitUse := claims.(*request.CustomClaims)
		return waitUse.ID
	}
}

// 从Gin的Context中获取jwt解析出来的用户UUID

func GetUserUuid(c *gin.Context) uuid.UUID {
	if claims, exists := c.Get("claims"); !exists {
		global.GVA_LOG.Error("从Gin的Context中获取从jwt解析出来的用户uuid失败, 请检查路由是否使用jwt中间件")
		return uuid.UUID{}
	} else {
		waitUse := claims.(*request.CustomClaims)
		return waitUse.UUID
	}
}

// 从Gin的Context中获取从jwt解析出来的用户角色id

func GetUserAuthorityId(c *gin.Context) string {
	if claims, exists := c.Get("claims"); !exists {
		global.GVA_LOG.Error("从Gin的Context中获取从jwt解析出阿里的用户UUID失败, 请检查路由是否使用jwt中间件!")
		return ""
	} else {
		waitUse := claims.(*request.CustomClaims)
		return waitUse.AuthorityId
	}
}
