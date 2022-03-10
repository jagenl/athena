// Package database 数据库操作
package database

import (
	"athena/pkg/config"
	"database/sql"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// DB 对象
var DB *gorm.DB
var SQLDB *sql.DB

// Connect 连接数据库
func Connect(dbConfig gorm.Dialector ) {

	// 使用 gorm.Open 连接数据库
	var err error
	DB, err = gorm.Open(dbConfig, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	})
	// 处理错误
	if err != nil {
		fmt.Println(err.Error())
	}

	// 获取底层的 sqlDB
	SQLDB, err = DB.DB()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func CurrentDatabase() (dbname string) {
	dbname = DB.Migrator().CurrentDatabase()
	return
}

func DeleteAllTables() error {

	var err error
	switch config.GetString("database.connection") {
	case "mysql":
		err = deleteMysqlDatabase()
	case "sqlite":
		deleteAllSqliteTables()
	default:
		panic(errors.New("database connection not supported"))
	}

	return err
}

func deleteAllSqliteTables() error {
	tables := []string{}
	DB.Select(&tables, "SELECT name FROM sqlite_master WHERE type='table'")
	for _, table := range tables {
		DB.Migrator().DropTable(table)
	}
	return nil
}

func deleteMysqlDatabase() error {
	dbname := CurrentDatabase()
	sql := fmt.Sprintf("DROP DATABASE %s;", dbname)
	if err := DB.Exec(sql).Error; err != nil {
		return err
	}
	sql = fmt.Sprintf("CREATE DATABASE %s;", dbname)
	if err := DB.Exec(sql).Error; err != nil {
		return err
	}
	sql = fmt.Sprintf("USE %s;", dbname)
	if err := DB.Exec(sql).Error; err != nil {
		return err
	}
	return nil
}

func TableName(obj interface{}) string {
	stmt := &gorm.Statement{DB: DB}
	stmt.Parse(obj)
	return stmt.Schema.Table
}
