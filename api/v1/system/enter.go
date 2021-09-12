package system

import (
	"gin-vue-admin-study/service"
)

type ApiGroup struct {
	AuthorityApi
	JwtApi
	BaseApi
	DBApi
	AuthorityMenuApi
}

var authorityService = service.ServiceGroupApp.SystemServiceGroup.AuthorityService
var jwtService = service.ServiceGroupApp.SystemServiceGroup.JwtService
var operationRecordService = service.ServiceGroupApp.SystemServiceGroup.OperationRecordService
var userService = service.ServiceGroupApp.SystemServiceGroup.UserService
var initDBService = service.ServiceGroupApp.SystemServiceGroup.InitDBService
var menuDbService = service.ServiceGroupApp.SystemServiceGroup.MenuService
