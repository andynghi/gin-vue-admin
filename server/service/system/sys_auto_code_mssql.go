package system

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
)

var AutoCodeMssql = new(autoCodeMssql)

type autoCodeMssql struct{}

// GetDB gets all database names of the database
// Author [piexlmax](https://github.com/piexlmax)
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *autoCodeMssql) GetDB(businessDB string) (data []response.Db, err error) {
	var entities []response.Db
	sql := "select name AS 'database' from sysdatabases;"
	if businessDB == "" {
		err = global.GVA_DB.Raw(sql).Scan(&entities).Error
	} else {
		err = global.GVA_DBList[businessDB].Raw(sql).Scan(&entities).Error
	}
	return entities, err
}

// GetTables gets all table names in the database
// Author [piexlmax](https://github.com/piexlmax)
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *autoCodeMssql) GetTables(businessDB string, dbName string) (data []response.Table, err error) {
	var entities []response.Table

	sql := fmt.Sprintf(`select name as 'table_name' from %s.DBO.sysobjects where xtype='U'`, dbName)
	if businessDB == "" {
		err = global.GVA_DB.Raw(sql).Scan(&entities).Error
	} else {
		err = global.GVA_DBList[businessDB].Raw(sql).Scan(&entities).Error
	}

	return entities, err
}

// GetColumn gets all field names, type values, etc. of the specified database and specified data table
// Author [piexlmax](https://github.com/piexlmax)
// Author [SliverHorn](https://github.com/SliverHorn)
func (s *autoCodeMssql) GetColumn(businessDB string, tableName string, dbName string) (data []response.Column, err error) {
	var entities []response.Column
	sql := fmt.Sprintf(`select sc.name as column_name,st.name as data_type, sc.length as data_type_long
	 from %s.DBO.syscolumns sc,systypes st where sc.xtype=st.xtype and st.usertype=0 and sc.id in (select id from %s.DBO.sysobjects where xtype='U' and name='%s');`, dbName, dbName, tableName)

	if businessDB == "" {
		err = global.GVA_DB.Raw(sql).Scan(&entities).Error
	} else {
		err = global.GVA_DBList[businessDB].Raw(sql).Scan(&entities).Error
	}

	return entities, err
}
