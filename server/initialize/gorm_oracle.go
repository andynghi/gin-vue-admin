package initialize

import (
	//"github.com/dzwvip/oracle"
	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/initialize/internal"

	//_ "github.com/godror/godror"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// GormOracle initializes the oracle database
// If you need the Oracle library, uncomment the import and change mysql.Config below to oracle.Config; mysql.New to oracle.New
func GormOracle() *gorm.DB {
	m := global.GVA_CONFIG.Oracle
	if m.Dbname == "" {
		return nil
	}
	oracleConfig := mysql.Config{
		DSN:               m.Dsn(), // DSN data source name
		DefaultStringSize: 191,     //Default length of string type field
	}
	if db, err := gorm.Open(mysql.New(oracleConfig), internal.Gorm.Config(m.Prefix, m.Singular)); err != nil {
		panic(err)
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

// GormOracleByConfig initializes the Oracle database using the incoming configuration
func GormOracleByConfig(m config.Oracle) *gorm.DB {
	if m.Dbname == "" {
		return nil
	}
	oracleConfig := mysql.Config{
		DSN:               m.Dsn(), // DSN data source name
		DefaultStringSize: 191,     //Default length of string type field
	}
	if db, err := gorm.Open(mysql.New(oracleConfig), internal.Gorm.Config(m.Prefix, m.Singular)); err != nil {
		panic(err)
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}
