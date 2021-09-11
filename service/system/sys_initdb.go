package system

import (
	"database/sql"
	"fmt"
	"github.com/spf13/viper"
	"go-class/config"
	"go-class/global"
	modelSystemReq "go-class/model/system/request"
	"go-class/model/system/tables"
	"go-class/source"
	"go-class/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"path/filepath"
)

type InitDBService struct {
}

// 创建数据库

func (initDBService *InitDBService) createTable(dsn string, driver string, createSql string) error {

	db, err := sql.Open(driver, dsn)
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
		}
	}(db)

	if err = db.Ping(); err != nil {
		return err
	}
	_, err = db.Exec(createSql)
	return err
}

func (initDBService *InitDBService) writeConfig(viper *viper.Viper, mysql config.Mysql) error {
	global.GVA_CONFIG.Mysql = mysql
	cs := utils.StructToMap(global.GVA_CONFIG)
	for k, v := range cs {
		viper.Set(k, v)
	}
	return viper.WriteConfig()
}

func (initDBService *InitDBService) initDB(InitDBFunctions ...tables.InitDbFunc) (err error) {
	for _, v := range InitDBFunctions {
		err = v.Init()
		if err != nil {
			return err
		}
	}
	return nil
}

// 创建数据库和表

func (initDBService *InitDBService) InitDB(conf modelSystemReq.InitDB) error {
	if conf.Host == "" {
		conf.Host = "127.0.0.1"
	}
	if conf.Port == "" {
		conf.Port = "3306"
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/", conf.UserName, conf.Password, conf.Host, conf.Port)
	createSql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;", conf.DBName)
	if err := initDBService.createTable(dsn, "mysql", createSql); err != nil {
		return err
	}

	MysqlConfig := config.Mysql{
		Path:     fmt.Sprintf("%s:%s", conf.Host, conf.Port),
		Dbname:   conf.DBName,
		Username: conf.UserName,
		Password: conf.Password,
		Config:   "charset=utf8mb4&parseTime=True&loc=Asia%2FShanghai",
	}

	if MysqlConfig.Dbname == "" {
		return nil
	}
	linkDns := MysqlConfig.Username + ":" + MysqlConfig.Password + "@tcp(" + MysqlConfig.Path + ")/" + MysqlConfig.Dbname + "?" + MysqlConfig.Config
	mysqlConfig := mysql.Config{
		DSN:                       linkDns, // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		DisableDatetimePrecision:  true,    // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,    // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,    // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true, NamingStrategy: schema.NamingStrategy{
		//TablePrefix:   "", // 表名前缀，`User`表为`t_users`
		SingularTable: true, // 使用单数表名，启用该选项后，`User` 表将是`user`
	}}); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(MysqlConfig.MaxIdleConns)
		sqlDB.SetMaxOpenConns(MysqlConfig.MaxOpenConns)
		global.GVA_DB = db
	}

	err := global.GVA_DB.AutoMigrate(
		tables.SysUser{},
		tables.SysAuthority{},
		tables.SysBaseMenu{},
		tables.SysBaseMenuParameter{},
		tables.JwtBlacklist{},
		tables.SysOperationRecord{},
		tables.SysUseAuthority{},
	)

	if err != nil {
		global.GVA_DB = nil
		return err
	}

	err = initDBService.initDB(
		source.Admin,
		source.AuthorityMenu,
		source.Authority,
		source.AuthoritiesMenus,
		source.AuthorityIdList,
		source.BaseMenu,
		source.UserAuthority,
	)

	if err != nil {
		global.GVA_DB = nil
		return err
	}
	if err = initDBService.writeConfig(global.GVA_VP, MysqlConfig); err != nil {
		return err
	}
	global.GVA_CONFIG.AutoCode.Root, _ = filepath.Abs("..")
	return nil
}
