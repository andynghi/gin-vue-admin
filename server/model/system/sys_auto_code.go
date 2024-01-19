package system

import (
	"errors"
	"go/token"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// AutoCodeStruct initial version automation code tool
type AutoCodeStruct struct {
	StructName         string   `json:"structName"`         // Struct name
	TableName          string   `json:"tableName"`          // table name
	PackageName        string   `json:"packageName"`        // File name
	HumpPackageName    string   `json:"humpPackageName"`    // go file name
	Abbreviation       string   `json:"abbreviation"`       // Struct abbreviation
	Description        string   `json:"description"`        // Struct Chinese name
	AutoCreateApiToSql bool     `json:"autoCreateApiToSql"` // Whether to automatically create api
	AutoCreateResource bool     `json:"autoCreateResource"` // Whether to automatically create resource identifiers
	AutoMoveFile       bool     `json:"autoMoveFile"`       // Whether to automatically move files
	BusinessDB         string   `json:"businessDB"`         // Business database
	Fields             []*Field `json:"fields,omitempty"`
	HasTimer           bool
	DictTypes          []string `json:"-"`
	Package            string   `json:"package"`
	PackageT           string   `json:"-"`
	NeedSort           bool     `json:"-"`
	HasPic             bool     `json:"-"`
	HasRichText        bool     `json:"-"`
	HasFile            bool     `json:"-"`
	NeedJSON           bool     `json:"-"`
}

func (a *AutoCodeStruct) Pretreatment() {
	a.KeyWord()
	a.SuffixTest()
}

// KeyWord is the processing of go keywords plus _ to prevent compilation errors.
// Author [SliverHorn](https://github.com/SliverHorn)
func (a *AutoCodeStruct) KeyWord() {
	if token.IsKeyword(a.Abbreviation) {
		a.Abbreviation = a.Abbreviation + "_"
	}
}

// SuffixTest handles the _test suffix
// Author [SliverHorn](https://github.com/SliverHorn)
func (a *AutoCodeStruct) SuffixTest() {
	if strings.HasSuffix(a.HumpPackageName, "test") {
		a.HumpPackageName = a.HumpPackageName + "_"
	}
}

type Field struct {
	FieldName       string `json:"fieldName"`       // Field name
	FieldDesc       string `json:"fieldDesc"`       // Chinese name
	FieldType       string `json:"fieldType"`       // Field data type
	FieldJson       string `json:"fieldJson"`       // FieldJson
	DataTypeLong    string `json:"dataTypeLong"`    // Database field length
	Comment         string `json:"comment"`         // Database field description
	ColumnName      string `json:"columnName"`      // Database field
	FieldSearchType string `json:"fieldSearchType"` // Search conditions
	DictType        string `json:"dictType"`        // Dictionary
	Require         bool   `json:"require"`         // Is it required?
	ErrorText       string `json:"errorText"`       // Verification failure text
	Clearable       bool   `json:"clearable"`       // Whether it can be cleared
	Sort            bool   `json:"sort"`            // Whether to increase sorting
}

var ErrAutoMove error = errors.New("The code was created successfully and the file was moved successfully")

type SysAutoCode struct {
	global.GVA_MODEL
	PackageName string `json:"packageName" gorm:"comment:package name"`
	Label       string `json:"label" gorm:"comment:display name"`
	Desc        string `json:"desc" gorm:"comment:Description"`
}

type AutoPlugReq struct {
	PlugName    string         `json:"plugName"` // Must start with capitals
	Snake       string         `json:"snake"`    // The backend automatically converts to snake
	RouterGroup string         `json:"routerGroup"`
	HasGlobal   bool           `json:"hasGlobal"`
	HasRequest  bool           `json:"hasRequest"`
	HasResponse bool           `json:"hasResponse"`
	NeedModel   bool           `json:"needModel"`
	Global      []AutoPlugInfo `json:"global,omitempty"`
	Request     []AutoPlugInfo `json:"request,omitempty"`
	Response    []AutoPlugInfo `json:"response,omitempty"`
}

func (a *AutoPlugReq) CheckList() {
	a.Global = bind(a.Global)
	a.Request = bind(a.Request)
	a.Response = bind(a.Response)

}
func bind(req []AutoPlugInfo) []AutoPlugInfo {
	var r []AutoPlugInfo
	for _, info := range req {
		if info.Effective() {
			r = append(r, info)
		}
	}
	return r
}

type AutoPlugInfo struct {
	Key  string `json:"key"`
	Type string `json:"type"`
	Desc string `json:"desc"`
}

func (a AutoPlugInfo) Effective() bool {
	return a.Key != "" && a.Type != "" && a.Desc != ""
}
