package request

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"os"
)

type InitDB struct {
	DBType   string `json:"dbType"`                    // Database type
	Host     string `json:"host"`                      // Server address
	Port     string `json:"port"`                      // Database connection port
	UserName string `json:"userName"`                  // Database user name
	Password string `json:"password"`                  // Database password
	DBName   string `json:"dbName" binding:"required"` // Database name
	DBPath   string `json:"dbPath"`                    // sqlite database file path
}

// MysqlEmptyDsn msyql empty database database creation link
// Author SliverHorn
func (i *InitDB) MysqlEmptyDsn() string {
	if i.Host == "" {
		i.Host = "127.0.0.1"
	}
	if i.Port == "" {
		i.Port = "3306"
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/", i.UserName, i.Password, i.Host, i.Port)
}

// PgsqlEmptyDsn pgsql empty database database creation link
// Author SliverHorn
func (i *InitDB) PgsqlEmptyDsn() string {
	if i.Host == "" {
		i.Host = "127.0.0.1"
	}
	if i.Port == "" {
		i.Port = "5432"
	}
	return "host=" + i.Host + " user=" + i.UserName + " password=" + i.Password + " port=" + i.Port + " dbname=" + "postgres" + " " + "sslmode=disable TimeZone=Asia/Shanghai"
}

// SqliteEmptyDsn sqlite empty database database creation link
// Author Kafumio
func (i *InitDB) SqliteEmptyDsn() string {
	separator := string(os.PathSeparator)
	return i.DBPath + separator + i.DBName + ".db"
}

// ToMysqlConfig converts config.Mysql
// Author [SliverHorn](https://github.com/SliverHorn)
func (i *InitDB) ToMysqlConfig() config.Mysql {
	return config.Mysql{
		GeneralDB: config.GeneralDB{
			Path:         i.Host,
			Port:         i.Port,
			Dbname:       i.DBName,
			Username:     i.UserName,
			Password:     i.Password,
			MaxIdleConns: 10,
			MaxOpenConns: 100,
			LogMode:      "error",
			Config:       "charset=utf8mb4&parseTime=True&loc=Local",
		},
	}
}

// ToPgsqlConfig converts config.Pgsql
// Author [SliverHorn](https://github.com/SliverHorn)
func (i *InitDB) ToPgsqlConfig() config.Pgsql {
	return config.Pgsql{
		GeneralDB: config.GeneralDB{
			Path:         i.Host,
			Port:         i.Port,
			Dbname:       i.DBName,
			Username:     i.UserName,
			Password:     i.Password,
			MaxIdleConns: 10,
			MaxOpenConns: 100,
			LogMode:      "error",
			Config:       "sslmode=disable TimeZone=Asia/Shanghai",
		},
	}
}

// ToSqliteConfig conversion config.Sqlite
// Author [Kafumio](https://github.com/Kafumio)
func (i *InitDB) ToSqliteConfig() config.Sqlite {
	return config.Sqlite{
		GeneralDB: config.GeneralDB{
			Path:         i.DBPath,
			Port:         i.Port,
			Dbname:       i.DBName,
			Username:     i.UserName,
			Password:     i.Password,
			MaxIdleConns: 10,
			MaxOpenConns: 100,
			LogMode:      "error",
			Config:       "",
		},
	}
}
