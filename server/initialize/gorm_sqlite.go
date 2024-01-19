package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/initialize/internal"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

// GormSqlite initializes Sqlite database
func GormSqlite() *gorm.DB {
	s := global.GVA_CONFIG.Sqlite
	if s.Dbname == "" {
		return nil
	}

	if db, err := gorm.Open(sqlite.Open(s.Dsn()), internal.Gorm.Config(s.Prefix, s.Singular)); err != nil {
		panic(err)
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(s.MaxIdleConns)
		sqlDB.SetMaxOpenConns(s.MaxOpenConns)
		return db
	}
}

// GormSqliteByConfig initializes the Sqlite database using the incoming configuration
func GormSqliteByConfig(s config.Sqlite) *gorm.DB {
	if s.Dbname == "" {
		return nil
	}

	if db, err := gorm.Open(sqlite.Open(s.Dsn()), internal.Gorm.Config(s.Prefix, s.Singular)); err != nil {
		panic(err)
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(s.MaxIdleConns)
		sqlDB.SetMaxOpenConns(s.MaxOpenConns)
		return db
	}
}
