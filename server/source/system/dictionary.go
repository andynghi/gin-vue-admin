package system

import (
	"context"
	sysModel "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

const initOrderDict = initOrderCasbin + 1

type initDict struct{}

// auto run
func init() {
	system.RegisterInit(initOrderDict, &initDict{})
}

func (i *initDict) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&sysModel.SysDictionary{})
}

func (i *initDict) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&sysModel.SysDictionary{})
}

func (i initDict) InitializerName() string {
	return sysModel.SysDictionary{}.TableName()
}

func (i *initDict) InitializeData(ctx context.Context) (next context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	True := true
	entities := []sysModel.SysDictionary{
		{Name: "gender", Type: "gender", Status: &True, Desc: "gender dictionary"},
		{Name: "Database int type", Type: "int", Status: &True, Desc: "Database type corresponding to int type"},
		{Name: "Database time and date type", Type: "time.Time", Status: &True, Desc: "Database time and date type"},
		{Name: "Database Float Type", Type: "float64", Status: &True, Desc: "Database Float Type"},
		{Name: "Database String", Type: "string", Status: &True, Desc: "Database String"},
		{Name: "Database bool type", Type: "bool", Status: &True, Desc: "Database bool type"},
	}

	if err = db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, sysModel.SysDictionary{}.TableName()+"Table data initialization failed!")
	}
	next = context.WithValue(ctx, i.InitializerName(), entities)
	return next, nil
}

func (i *initDict) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.Where("type = ?", "bool").First(&sysModel.SysDictionary{}).Error, gorm.ErrRecordNotFound) { // Determine whether data exists
		return false
	}
	return true
}
