package system

type ServiceGroup struct {
	JwtService
	AuthorityService
	OperationRecordService
	UserService
	InitDBService
	MenuService
}
