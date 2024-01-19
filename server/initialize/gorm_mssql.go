/*
 * @Author: Flying against the light 191180776@qq.com
 * @Date: 2022-12-08 17:25:49
 * @LastEditors: Flying against the light 191180776@qq.com
 * @LastEditTime: 2022-12-08 18:00:00
 * @FilePath: \server\initialize\gorm_mssql.go
 * @Description: This is the default setting, please set `customMade`, open koroFileHeader to view the configuration and set it: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/initialize/internal"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

// GormMssql initializes Mssql database
// Author [LouisZhang](191180776@qq.com)
func GormMssql() *gorm.DB {
	m := global.GVA_CONFIG.Mssql
	if m.Dbname == "" {
		return nil
	}
	mssqlConfig := sqlserver.Config{
		DSN:               m.Dsn(), // DSN data source name
		DefaultStringSize: 191,     //Default length of string type field
	}
	if db, err := gorm.Open(sqlserver.New(mssqlConfig), internal.Gorm.Config(m.Prefix, m.Singular)); err != nil {
		return nil
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE="+m.Engine)
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

// GormMssqlByConfig initializes the Mysql database using the incoming configuration
func GormMssqlByConfig(m config.Mssql) *gorm.DB {
	if m.Dbname == "" {
		return nil
	}
	mssqlConfig := sqlserver.Config{
		DSN:               m.Dsn(), // DSN data source name
		DefaultStringSize: 191,     //Default length of string type field
	}
	if db, err := gorm.Open(sqlserver.New(mssqlConfig), internal.Gorm.Config(m.Prefix, m.Singular)); err != nil {
		panic(err)
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE=InnoDB")
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}
