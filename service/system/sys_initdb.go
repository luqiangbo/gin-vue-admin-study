package system

import (
	"database/sql"
	"fmt"
	modelSystemReq "go-class/model/system/request"
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

// 创建数据库

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
	return nil
}
