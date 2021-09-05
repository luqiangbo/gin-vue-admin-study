package system

import (
	"go-class/service"
)

type ApiGroup struct {
	AuthorityApi
	JwtApi
	BaseApi
}

var authorityService = service.ServiceGroupApp.SystemServiceGroup.AuthorityService
var jwtService = service.ServiceGroupApp.SystemServiceGroup.JwtService
var operationRecordService = service.ServiceGroupApp.SystemServiceGroup.OperationRecordService
var userService = service.ServiceGroupApp.SystemServiceGroup.UserService
