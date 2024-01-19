package system

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"gorm.io/gorm"
	"sort"
)

const (
	Mysql           = "mysql"
	Pgsql           = "pgsql"
	Sqlite          = "sqlite"
	InitSuccess     = "\n[%v] --> Initial data successful!\n"
	InitDataExist   = "\n[%v] --> The initial data of %v already exists!\n"
	InitDataFailed  = "\n[%v] --> %v initial data failed! \nerr: %+v\n"
	InitDataSuccess = "\n[%v] --> %v initial data successful!\n"
)

const (
	InitOrderSystem   = 10
	InitOrderInternal = 1000
	InitOrderExternal = 100000
)

var (
	ErrMissingDBContext        = errors.New("missing db in context")
	ErrMissingDependentContext = errors.New("missing dependent value in context")
	ErrDBTypeMismatch          = errors.New("db type mismatch")
)

// SubInitializer provides the interface used by source/*/init(). Each initializer completes an initialization process.
type SubInitializer interface {
	InitializerName() string // does not necessarily represent a single table, so it is changed to a broader semantics
	MigrateTable(ctx context.Context) (next context.Context, err error)
	InitializeData(ctx context.Context) (next context.Context, err error)
	TableCreated(ctx context.Context) bool
	DataInserted(ctx context.Context) bool
}

// TypedDBInitHandler executes the passed in initializer
type TypedDBInitHandler interface {
	EnsureDB(ctx context.Context, conf *request.InitDB) (context.Context, error) //Build the database, failure is a fatal error, so let it panic
	WriteConfig(ctx context.Context) error                                       // Write back configuration
	InitTables(ctx context.Context, inits initSlice) error                       // Create table handler
	InitData(ctx context.Context, inits initSlice) error                         // Create data handler
}

// orderedInitializer combines an ordered field for sorting
type orderedInitializer struct {
	order int
	SubInitializer
}

// initSlice is used by the initializer when sorting dependencies
type initSlice []*orderedInitializer

var (
	initializers initSlice
	cache        map[string]*orderedInitializer
)

// RegisterInit registers the initialization process to be executed, which will be called during InitDB()
func RegisterInit(order int, i SubInitializer) {
	if initializers == nil {
		initializers = initSlice{}
	}
	if cache == nil {
		cache = map[string]*orderedInitializer{}
	}
	name := i.InitializerName()
	if _, existed := cache[name]; existed {
		panic(fmt.Sprintf("Name conflict on %s", name))
	}
	ni := orderedInitializer{order, i}
	initializers = append(initializers, &ni)
	cache[name] = &ni
}

/* ---- * service * ---- */

type InitDBService struct{}

// InitDB creates the database and initializes the main entrance
func (initDBService *InitDBService) InitDB(conf request.InitDB) (err error) {
	ctx := context.TODO()
	if len(initializers) == 0 {
		return errors.New("No initialization process available, please check whether the initialization has been completed")
	}
	sort.Sort(&initializers) // Ensure that dependent initializers are executed later
	// Note: If the initializer has only a single dependency, it can be written as B=A+1, C=A+1; since there is no dependency between BC, who comes first does not affect the initialization
	// If there are multiple dependencies, it can be written as C=A+B, D=A+B+C, E=A+1;
	// C must be >A|B, so it is executed after AB, D must be >A|B|C, so it is executed after ABC, and E only depends on A. The order has nothing to do with CD, so it does not matter which one of E or CD is executed first. Influence
	var initHandler TypedDBInitHandler
	switch conf.DBType {
	case "mysql":
		initHandler = NewMysqlInitHandler()
		ctx = context.WithValue(ctx, "dbtype", "mysql")
	case "pgsql":
		initHandler = NewPgsqlInitHandler()
		ctx = context.WithValue(ctx, "dbtype", "pgsql")
	case "sqlite":
		initHandler = NewSqliteInitHandler()
		ctx = context.WithValue(ctx, "dbtype", "sqlite")
	default:
		initHandler = NewMysqlInitHandler()
		ctx = context.WithValue(ctx, "dbtype", "mysql")
	}
	ctx, err = initHandler.EnsureDB(ctx, &conf)
	if err != nil {
		return err
	}

	db := ctx.Value("db").(*gorm.DB)
	global.GVA_DB = db

	if err = initHandler.InitTables(ctx, initializers); err != nil {
		return err
	}
	if err = initHandler.InitData(ctx, initializers); err != nil {
		return err
	}

	if err = initHandler.WriteConfig(ctx); err != nil {
		return err
	}
	initializers = initSlice{}
	cache = map[string]*orderedInitializer{}
	return nil
}

// createDatabase creates a database (called in EnsureDB())
func createDatabase(dsn string, driver string, createSql string) error {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(db)
	if err = db.Ping(); err != nil {
		return err
	}
	_, err = db.Exec(createSql)
	return err
}

// createTables creates tables (default dbInitHandler.initTables behavior)
func createTables(ctx context.Context, inits initSlice) error {
	next, cancel := context.WithCancel(ctx)
	defer func(c func()) { c() }(cancel)
	for _, init := range inits {
		if init.TableCreated(next) {
			continue
		}
		if n, err := init.MigrateTable(next); err != nil {
			return err
		} else {
			next = n
		}
	}
	return nil
}

/* -- sortable interface -- */

func (a initSlice) Len() int {
	return len(a)
}

func (a initSlice) Less(i, j int) bool {
	return a[i].order < a[j].order
}

func (a initSlice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
