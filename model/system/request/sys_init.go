package request

type InitDB struct {
	Host     string `json:"host"`                         // 服务器地址
	Port     string `json:"port"`                         // 数据库链接端口
	UserName string `json:"user_name" binding:"required"` // 数据库用户名
	Password string `json:"password"`                     // 数据库密码
	DBName   string `json:"db_name" binding:"required"`   // 数据名
}
