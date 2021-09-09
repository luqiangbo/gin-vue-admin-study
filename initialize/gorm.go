package initialize

import (
	"go-class/global"
	"go-class/initialize/internal"
	"go-class/model/system/tables"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
)

// 初始化数据库并产生数据库变量

func Gorm() *gorm.DB {
	switch global.GVA_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	default:
		return GormMysql()
	}
}

// 注册数据库表专用

func MysqlTables(dbs *gorm.DB) {
	err := dbs.AutoMigrate(
		tables.SysUser{},
		tables.JwtBlacklist{},
		tables.SysOperationRecord{},

		//tables.SysBaseMenu{},
		//tables.SysAuthority{},
		//tables.SysBaseMenuParameter{},

	)
	if err != nil {
		global.GVA_LOG.Error("register table failed", zap.Any("err", err))
		os.Exit(0)
	}
	global.GVA_LOG.Info("register table success")
}

func GormMysql() *gorm.DB {
	m := global.GVA_CONFIG.Mysql
	if m.Dbname == "" {
		return nil
	}
	dsn := m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.Dbname + "?" + m.Config
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置

	}
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             1000, // 慢 SQL 阈值
			IgnoreRecordNotFoundError: true, // 忽略ErrRecordNotFound（记录未找到）错误
		},
	)
	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix:   "", // 表名前缀，`User`表为`t_users`
			SingularTable: true, // 使用单数表名，启用该选项后，`User` 表将是`user`
		},
	}); err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

func gormConfig(mod bool) *gorm.Config {
	var config = &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
	switch global.GVA_CONFIG.Mysql.LogZap {
	case "silent", "Silent":
		config.Logger = internal.Default.LogMode(logger.Silent)
	case "error", "Error":
		config.Logger = internal.Default.LogMode(logger.Error)
	case "warn", "Warn":
		config.Logger = internal.Default.LogMode(logger.Warn)
	case "info", "Info":
		config.Logger = internal.Default.LogMode(logger.Info)
	case "zap", "Zap":
		config.Logger = internal.Default.LogMode(logger.Info)
	default:
		if mod {
			config.Logger = internal.Default.LogMode(logger.Info)
			break
		}
		config.Logger = internal.Default.LogMode(logger.Silent)
	}
	return config
}
