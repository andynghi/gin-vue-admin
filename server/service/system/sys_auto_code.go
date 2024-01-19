package system

import (
	"archive/zip"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"

	ast2 "github.com/flipped-aurora/gin-vue-admin/server/utils/ast"

	"github.com/flipped-aurora/gin-vue-admin/server/resource/autocode_template/subcontract"
	cp "github.com/otiai10/copy"
	"go.uber.org/zap"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"

	"gorm.io/gorm"
)

const (
	autoPath           = "autocode_template/"
	autocodePath       = "resource/autocode_template"
	plugPath           = "resource/plug_template"
	packageService     = "service/%s/enter.go"
	packageServiceName = "service"
	packageRouter      = "router/%s/enter.go"
	packageRouterName  = "router"
	packageAPI         = "api/v1/%s/enter.go"
	packageAPIName     = "api/v1"
)

type autoPackage struct {
	path string
	temp string
	name string
}

var (
	packageInjectionMap map[string]astInjectionMeta
	injectionPaths      []injectionMeta
)

func Init(Package string) {
	injectionPaths = []injectionMeta{
		{
			path: filepath.Join(global.GVA_CONFIG.AutoCode.Root,
				global.GVA_CONFIG.AutoCode.Server, fmt.Sprintf(global.GVA_CONFIG.AutoCode.SApi, Package), "enter.go"),
			funcName:    "ApiGroup",
			structNameF: "%sApi",
		},
		{
			path: filepath.Join(global.GVA_CONFIG.AutoCode.Root,
				global.GVA_CONFIG.AutoCode.Server, fmt.Sprintf(global.GVA_CONFIG.AutoCode.SRouter, Package), "enter.go"),
			funcName:    "RouterGroup",
			structNameF: "%sRouter",
		},
		{
			path: filepath.Join(global.GVA_CONFIG.AutoCode.Root,
				global.GVA_CONFIG.AutoCode.Server, fmt.Sprintf(global.GVA_CONFIG.AutoCode.SService, Package), "enter.go"),
			funcName:    "ServiceGroup",
			structNameF: "%sService",
		},
	}

	packageInjectionMap = map[string]astInjectionMeta{
		packageServiceName: {
			path: filepath.Join(global.GVA_CONFIG.AutoCode.Root,
				global.GVA_CONFIG.AutoCode.Server, "service", "enter.go"),
			importCodeF:  "github.com/flipped-aurora/gin-vue-admin/server/%s/%s",
			packageNameF: "%s",
			groupName:    "ServiceGroup",
			structNameF:  "%sServiceGroup",
		},
		packageRouterName: {
			path: filepath.Join(global.GVA_CONFIG.AutoCode.Root,
				global.GVA_CONFIG.AutoCode.Server, "router", "enter.go"),
			importCodeF:  "github.com/flipped-aurora/gin-vue-admin/server/%s/%s",
			packageNameF: "%s",
			groupName:    "RouterGroup",
			structNameF:  "%s",
		},
		packageAPIName: {
			path: filepath.Join(global.GVA_CONFIG.AutoCode.Root,
				global.GVA_CONFIG.AutoCode.Server, "api/v1", "enter.go"),
			importCodeF:  "github.com/flipped-aurora/gin-vue-admin/server/%s/%s",
			packageNameF: "%s",
			groupName:    "ApiGroup",
			structNameF:  "%sApiGroup",
		},
	}
}

type injectionMeta struct {
	path        string
	funcName    string
	structNameF string // formatted
}

type astInjectionMeta struct {
	path         string
	importCodeF  string
	structNameF  string
	packageNameF string
	groupName    string
}

type tplData struct {
	template         *template.Template
	autoPackage      string
	locationPath     string
	autoCodePath     string
	autoMoveFilePath string
}

type AutoCodeService struct{}

var AutoCodeServiceApp = new(AutoCodeService)

// @author: [songzhibin97](https://github.com/songzhibin97)
// @function: PreviewTemp
// @description: Preview creation code
// @param: model.AutoCodeStruct
// @return: map[string]string, error

func (autoCodeService *AutoCodeService) PreviewTemp(autoCode system.AutoCodeStruct) (map[string]string, error) {
	makeDictTypes(&autoCode)
	for i := range autoCode.Fields {
		if autoCode.Fields[i].FieldType == "time.Time" {
			autoCode.HasTimer = true
		}
		if autoCode.Fields[i].Sort {
			autoCode.NeedSort = true
		}
		if autoCode.Fields[i].FieldType == "picture" {
			autoCode.HasPic = true
		}
		if autoCode.Fields[i].FieldType == "video" {
			autoCode.HasPic = true
		}
		if autoCode.Fields[i].FieldType == "richtext" {
			autoCode.HasRichText = true
		}
		if autoCode.Fields[i].FieldType == "pictures" {
			autoCode.HasPic = true
			autoCode.NeedJSON = true
		}
		if autoCode.Fields[i].FieldType == "file" {
			autoCode.HasFile = true
			autoCode.NeedJSON = true
		}
	}
	dataList, _, needMkdir, err := autoCodeService.getNeedList(&autoCode)
	if err != nil {
		return nil, err
	}

	// Before writing the file, create the folder first
	if err = utils.CreateDir(needMkdir...); err != nil {
		return nil, err
	}

	// Create map
	ret := make(map[string]string)

	// Generate map
	for _, value := range dataList {
		ext := ""
		if ext = filepath.Ext(value.autoCodePath); ext == ".txt" {
			continue
		}
		f, err := os.OpenFile(value.autoCodePath, os.O_CREATE|os.O_WRONLY, 0o755)
		if err != nil {
			return nil, err
		}
		if err = value.template.Execute(f, autoCode); err != nil {
			return nil, err
		}
		_ = f.Close()
		f, err = os.OpenFile(value.autoCodePath, os.O_CREATE|os.O_RDONLY, 0o755)
		if err != nil {
			return nil, err
		}
		builder := strings.Builder{}
		builder.WriteString("```")

		if ext != "" && strings.Contains(ext, ".") {
			builder.WriteString(strings.Replace(ext, ".", "", -1))
		}
		builder.WriteString("\n\n")
		data, err := io.ReadAll(f)
		if err != nil {
			return nil, err
		}
		builder.Write(data)
		builder.WriteString("\n\n```")

		pathArr := strings.Split(value.autoCodePath, string(os.PathSeparator))
		ret[pathArr[1]+"-"+pathArr[3]] = builder.String()
		_ = f.Close()

	}
	defer func() { // Remove intermediate files
		if err := os.RemoveAll(autoPath); err != nil {
			return
		}
	}()
	return ret, nil
}

func makeDictTypes(autoCode *system.AutoCodeStruct) {
	DictTypeM := make(map[string]string)
	for _, v := range autoCode.Fields {
		if v.DictType != "" {
			DictTypeM[v.DictType] = ""
		}
	}

	for k := range DictTypeM {
		autoCode.DictTypes = append(autoCode.DictTypes, k)
	}
}

// @author: [piexlmax](https://github.com/piexlmax)
// @function: CreateTemp
// @description: Create code
// @param: model.AutoCodeStruct
// @return: err error

func (autoCodeService *AutoCodeService) CreateTemp(autoCode system.AutoCodeStruct, ids ...uint) (err error) {
	makeDictTypes(&autoCode)
	for i := range autoCode.Fields {
		if autoCode.Fields[i].FieldType == "time.Time" {
			autoCode.HasTimer = true
		}
		if autoCode.Fields[i].Sort {
			autoCode.NeedSort = true
		}
		if autoCode.Fields[i].FieldType == "picture" {
			autoCode.HasPic = true
		}
		if autoCode.Fields[i].FieldType == "video" {
			autoCode.HasPic = true
		}
		if autoCode.Fields[i].FieldType == "richtext" {
			autoCode.HasRichText = true
		}
		if autoCode.Fields[i].FieldType == "pictures" {
			autoCode.NeedJSON = true
			autoCode.HasPic = true
		}
		if autoCode.Fields[i].FieldType == "file" {
			autoCode.NeedJSON = true
			autoCode.HasFile = true
		}
	}
	// Add judgment: Create struct repeatedly
	if autoCode.AutoMoveFile && AutoCodeHistoryServiceApp.Repeat(autoCode.BusinessDB, autoCode.StructName, autoCode.Package) {
		return RepeatErr
	}
	dataList, fileList, needMkdir, err := autoCodeService.getNeedList(&autoCode)
	if err != nil {
		return err
	}
	meta, _ := json.Marshal(autoCode)

	// Add judgment: Package is not empty
	if autoCode.Package == "" {
		return errors.New("Package is empty\n")
	}

	// Before writing the file, create the folder first
	if err = utils.CreateDir(needMkdir...); err != nil {
		return err
	}

	// Generate file
	for _, value := range dataList {
		f, err := os.OpenFile(value.autoCodePath, os.O_CREATE|os.O_WRONLY, 0o755)
		if err != nil {
			return err
		}
		if err = value.template.Execute(f, autoCode); err != nil {
			return err
		}
		_ = f.Close()
	}

	defer func() { // Remove intermediate files
		if err := os.RemoveAll(autoPath); err != nil {
			return
		}
	}()
	bf := strings.Builder{}
	idBf := strings.Builder{}
	injectionCodeMeta := strings.Builder{}
	for _, id := range ids {
		idBf.WriteString(strconv.Itoa(int(id)))
		idBf.WriteString(";")
	}
	if autoCode.AutoMoveFile { // Determine whether automatic transfer is required
		Init(autoCode.Package)
		for index := range dataList {
			autoCodeService.addAutoMoveFile(&dataList[index])
		}
		// Determine whether all target files can be moved
		for _, value := range dataList {
			if utils.FileExist(value.autoMoveFilePath) {
				return errors.New(fmt.Sprintf("The target file already exists: %s\n", value.autoMoveFilePath))
			}
		}
		for _, value := range dataList { //Move files
			if err := utils.FileMove(value.autoCodePath, value.autoMoveFilePath); err != nil {
				return err
			}
		}

		{
			//Inject automatic migration in gorm.go
			path := filepath.Join(global.GVA_CONFIG.AutoCode.Root,
				global.GVA_CONFIG.AutoCode.Server, global.GVA_CONFIG.AutoCode.SInitialize, "gorm.go")
			varDB := utils.MaheHump(autoCode.BusinessDB)
			ast2.AddRegisterTablesAst(path, "RegisterTables", autoCode.Package, varDB, autoCode.BusinessDB, autoCode.StructName)
		}

		{
			// router.go injection automatic migration
			path := filepath.Join(global.GVA_CONFIG.AutoCode.Root,
				global.GVA_CONFIG.AutoCode.Server, global.GVA_CONFIG.AutoCode.SInitialize, "router.go")
			ast2.AddRouterCode(path, "Routers", autoCode.Package, autoCode.StructName)
		}
		//Inject into each enter
		err = injectionCode(autoCode.StructName, &injectionCodeMeta)
		if err != nil {
			return
		}
		//Save generated information
		for _, data := range dataList {
			if len(data.autoMoveFilePath) != 0 {
				bf.WriteString(data.autoMoveFilePath)
				bf.WriteString(";")
			}
		}
	} else { // package
		if err = utils.ZipFiles("./ginvueadmin.zip", fileList, ".", "."); err != nil {
			return err
		}
	}
	if autoCode.AutoMoveFile || autoCode.AutoCreateApiToSql {
		if autoCode.TableName != "" {
			err = AutoCodeHistoryServiceApp.CreateAutoCodeHistory(
				string(meta),
				autoCode.StructName,
				autoCode.Description,
				bf.String(),
				injectionCodeMeta.String(),
				autoCode.TableName,
				idBf.String(),
				autoCode.Package,
				autoCode.BusinessDB,
			)
		} else {
			err = AutoCodeHistoryServiceApp.CreateAutoCodeHistory(
				string(meta),
				autoCode.StructName,
				autoCode.Description,
				bf.String(),
				injectionCodeMeta.String(),
				autoCode.StructName,
				idBf.String(),
				autoCode.Package,
				autoCode.BusinessDB,
			)
		}
	}
	if err != nil {
		return err
	}
	if autoCode.AutoMoveFile {
		return system.ErrAutoMove
	}
	return nil
}

// @author: [piexlmax](https://github.com/piexlmax)
// @function: GetAllTplFile
// @description: Get all tpl files in the pathName folder
// @param: pathName string, fileList []string
// @return: []string, error

func (autoCodeService *AutoCodeService) GetAllTplFile(pathName string, fileList []string) ([]string, error) {
	files, err := os.ReadDir(pathName)
	for _, fi := range files {
		if fi.IsDir() {
			fileList, err = autoCodeService.GetAllTplFile(pathName+"/"+fi.Name(), fileList)
			if err != nil {
				return nil, err
			}
		} else {
			if strings.HasSuffix(fi.Name(), ".tpl") {
				fileList = append(fileList, pathName+"/"+fi.Name())
			}
		}
	}
	return fileList, err
}

// @author: [piexlmax](https://github.com/piexlmax)
// @function: GetDB
// @description: Get all field names, type values, etc. of the specified database and specified data table
// @param: tableName string, dbName string
// @return: err error, Columns []request.ColumnReq

func (autoCodeService *AutoCodeService) DropTable(BusinessDb, tableName string) error {
	if BusinessDb != "" {
		return global.MustGetGlobalDBByDBName(BusinessDb).Exec("DROP TABLE " + tableName).Error
	} else {
		return global.GVA_DB.Exec("DROP TABLE " + tableName).Error
	}
}

// @author: [SliverHorn](https://github.com/SliverHorn)
// @author: [songzhibin97](https://github.com/songzhibin97)
// @function: addAutoMoveFile
// @description: Generate the corresponding migration file path
// @param: *tplData
// @return: null

func (autoCodeService *AutoCodeService) addAutoMoveFile(data *tplData) {
	base := filepath.Base(data.autoCodePath)
	fileSlice := strings.Split(data.autoCodePath, string(os.PathSeparator))
	n := len(fileSlice)
	if n <= 2 {
		return
	}
	if strings.Contains(fileSlice[1], "server") {
		if strings.Contains(fileSlice[n-2], "router") {
			data.autoMoveFilePath = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server,
				fmt.Sprintf(global.GVA_CONFIG.AutoCode.SRouter, data.autoPackage), base)
		} else if strings.Contains(fileSlice[n-2], "api") {
			data.autoMoveFilePath = filepath.Join(global.GVA_CONFIG.AutoCode.Root,
				global.GVA_CONFIG.AutoCode.Server, fmt.Sprintf(global.GVA_CONFIG.AutoCode.SApi, data.autoPackage), base)
		} else if strings.Contains(fileSlice[n-2], "service") {
			data.autoMoveFilePath = filepath.Join(global.GVA_CONFIG.AutoCode.Root,
				global.GVA_CONFIG.AutoCode.Server, fmt.Sprintf(global.GVA_CONFIG.AutoCode.SService, data.autoPackage), base)
		} else if strings.Contains(fileSlice[n-2], "model") {
			data.autoMoveFilePath = filepath.Join(global.GVA_CONFIG.AutoCode.Root,
				global.GVA_CONFIG.AutoCode.Server, fmt.Sprintf(global.GVA_CONFIG.AutoCode.SModel, data.autoPackage), base)
		} else if strings.Contains(fileSlice[n-2], "request") {
			data.autoMoveFilePath = filepath.Join(global.GVA_CONFIG.AutoCode.Root,
				global.GVA_CONFIG.AutoCode.Server, fmt.Sprintf(global.GVA_CONFIG.AutoCode.SRequest, data.autoPackage), base)
		}
	} else if strings.Contains(fileSlice[1], "web") {
		if strings.Contains(fileSlice[n-1], "js") {
			data.autoMoveFilePath = filepath.Join(global.GVA_CONFIG.AutoCode.Root,
				global.GVA_CONFIG.AutoCode.Web, global.GVA_CONFIG.AutoCode.WApi, base)
		} else if strings.Contains(fileSlice[n-2], "form") {
			data.autoMoveFilePath = filepath.Join(global.GVA_CONFIG.AutoCode.Root,
				global.GVA_CONFIG.AutoCode.Web, global.GVA_CONFIG.AutoCode.WForm, filepath.Base(filepath.Dir(filepath.Dir(data.autoCodePath))), strings.TrimSuffix(base, filepath.Ext(base))+"Form.vue")
		} else if strings.Contains(fileSlice[n-2], "table") {
			data.autoMoveFilePath = filepath.Join(global.GVA_CONFIG.AutoCode.Root,
				global.GVA_CONFIG.AutoCode.Web, global.GVA_CONFIG.AutoCode.WTable, filepath.Base(filepath.Dir(filepath.Dir(data.autoCodePath))), base)
		}
	}
}

// @author: [piexlmax](https://github.com/piexlmax)
// @author: [SliverHorn](https://github.com/SliverHorn)
// @function: CreateApi
// @description: Automatically create api data,
// @param: a *model.AutoCodeStruct
// @return: err error

func (autoCodeService *AutoCodeService) AutoCreateApi(a *system.AutoCodeStruct) (ids []uint, err error) {
	apiList := []system.SysApi{
		{
			Path:        "/" + a.Abbreviation + "/" + "create" + a.StructName,
			Description: "New" + a.Description,
			ApiGroup:    a.Description,
			Method:      "POST",
		},
		{
			Path:        "/" + a.Abbreviation + "/" + "delete" + a.StructName,
			Description: "Delete" + a.Description,
			ApiGroup:    a.Description,
			Method:      "DELETE",
		},
		{
			Path:        "/" + a.Abbreviation + "/" + "delete" + a.StructName + "ByIds",
			Description: "Batch delete" + a.Description,
			ApiGroup:    a.Description,
			Method:      "DELETE",
		},
		{
			Path:        "/" + a.Abbreviation + "/" + "update" + a.StructName,
			Description: "Update" + a.Description,
			ApiGroup:    a.Description,
			Method:      "PUT",
		},
		{
			Path:        "/" + a.Abbreviation + "/" + "find" + a.StructName,
			Description: "Get based on ID" + a.Description,
			ApiGroup:    a.Description,
			Method:      "GET",
		},
		{
			Path:        "/" + a.Abbreviation + "/" + "get" + a.StructName + "List",
			Description: "Get" + a.Description + "List",
			ApiGroup:    a.Description,
			Method:      "GET",
		},
	}
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		for _, v := range apiList {
			var api system.SysApi
			if errors.Is(tx.Where("path = ? AND method = ?", v.Path, v.Method).First(&api).Error, gorm.ErrRecordNotFound) {
				if err = tx.Create(&v).Error; err != nil { // Roll back the transaction when an error is encountered
					return err
				} else {
					ids = append(ids, v.ID)
				}
			}
		}
		return nil
	})
	return ids, err
}

func (autoCodeService *AutoCodeService) getNeedList(autoCode *system.AutoCodeStruct) (dataList []tplData, fileList []string, needMkdir []string, err error) {
	// Remove all spaces
	utils.TrimSpace(autoCode)
	for _, field := range autoCode.Fields {
		utils.TrimSpace(field)
	}
	// Get all tpl files in the basePath folder
	tplFileList, err := autoCodeService.GetAllTplFile(autocodePath, nil)
	if err != nil {
		return nil, nil, nil, err
	}
	dataList = make([]tplData, 0, len(tplFileList))
	fileList = make([]string, 0, len(tplFileList))
	needMkdir = make([]string, 0, len(tplFileList)) // When there are multiple tpl files in the folder, it is more reasonable to change to map
	//Generate the tplData structure based on the file path, and the data needs to be filled in
	for _, value := range tplFileList {
		dataList = append(dataList, tplData{locationPath: value, autoPackage: autoCode.Package})
	}
	// Generate *Template, fill in the template field
	for index, value := range dataList {
		dataList[index].template, err = template.ParseFiles(value.locationPath)
		if err != nil {
			return nil, nil, nil, err
		}
	}
	// Generate the file path and fill in the autoCodePath field. readme.txt.tpl does not comply with the rules and requires special processing.
	// resource/template/web/api.js.tpl -> autoCode/web/autoCode.PackageName/api/autoCode.PackageName.js
	// resource/template/readme.txt.tpl -> autoCode/readme.txt
	for index, value := range dataList {
		trimBase := strings.TrimPrefix(value.locationPath, autocodePath+"/")
		if trimBase == "readme.txt.tpl" {
			dataList[index].autoCodePath = autoPath + "readme.txt"
			continue
		}

		if lastSeparator := strings.LastIndex(trimBase, "/"); lastSeparator != -1 {
			origFileName := strings.TrimSuffix(trimBase[lastSeparator+1:], ".tpl")
			firstDot := strings.Index(origFileName, ".")
			if firstDot != -1 {
				var fileName string
				if origFileName[firstDot:] != ".go" {
					fileName = autoCode.PackageName + origFileName[firstDot:]
				} else {
					fileName = autoCode.HumpPackageName + origFileName[firstDot:]
				}

				dataList[index].autoCodePath = filepath.Join(autoPath, trimBase[:lastSeparator], autoCode.PackageName,
					origFileName[:firstDot], fileName)
			}
		}

		if lastSeparator := strings.LastIndex(dataList[index].autoCodePath, string(os.PathSeparator)); lastSeparator != -1 {
			needMkdir = append(needMkdir, dataList[index].autoCodePath[:lastSeparator])
		}
	}
	for _, value := range dataList {
		fileList = append(fileList, value.autoCodePath)
	}
	return dataList, fileList, needMkdir, err
}

// injectionCode encapsulates code injection
func injectionCode(structName string, bf *strings.Builder) error {
	for _, meta := range injectionPaths {
		code := fmt.Sprintf(meta.structNameF, structName)
		ast2.ImportForAutoEnter(meta.path, meta.funcName, code)
		bf.WriteString(fmt.Sprintf("%s@%s@%s;", meta.path, meta.funcName, code))
	}
	return nil
}

func (autoCodeService *AutoCodeService) CreateAutoCode(s *system.SysAutoCode) error {
	if s.PackageName == "autocode" || s.PackageName == "system" || s.PackageName == "example" || s.PackageName == "" {
		return errors.New("The reserved package name cannot be used")
	}
	if !errors.Is(global.GVA_DB.Where("package_name = ?", s.PackageName).First(&system.SysAutoCode{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("The same PackageName exists")
	}
	if e := autoCodeService.CreatePackageTemp(s.PackageName); e != nil {
		return e
	}
	return global.GVA_DB.Create(&s).Error
}

func (autoCodeService *AutoCodeService) GetPackage() (pkgList []system.SysAutoCode, err error) {
	err = global.GVA_DB.Find(&pkgList).Error
	return pkgList, err
}

func (autoCodeService *AutoCodeService) DelPackage(a system.SysAutoCode) error {
	return global.GVA_DB.Delete(&a).Error
}

func (autoCodeService *AutoCodeService) CreatePackageTemp(packageName string) error {
	Init(packageName)
	pendingTemp := []autoPackage{{
		path: packageService,
		name: packageServiceName,
		temp: string(subcontract.Server),
	}, {
		path: packageRouter,
		name: packageRouterName,
		temp: string(subcontract.Router),
	}, {
		path: packageAPI,
		name: packageAPIName,
		temp: string(subcontract.API),
	}}
	for i, s := range pendingTemp {
		pendingTemp[i].path = filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, filepath.Clean(fmt.Sprintf(s.path, packageName)))
	}
	//Select template
	for _, s := range pendingTemp {
		err := os.MkdirAll(filepath.Dir(s.path), 0755)
		if err != nil {
			return err
		}

		f, err := os.Create(s.path)
		if err != nil {
			return err
		}

		defer f.Close()

		temp, err := template.New("").Parse(s.temp)
		if err != nil {
			return err
		}
		err = temp.Execute(f, struct {
			PackageName string `json:"package_name"`
		}{packageName})
		if err != nil {
			return err
		}
	}
	// After creation, insert the structure code at the corresponding location
	for _, v := range pendingTemp {
		meta := packageInjectionMap[v.name]
		if err := ast2.ImportReference(meta.path, fmt.Sprintf(meta.importCodeF, v.name, packageName), fmt.Sprintf(meta.structNameF, utils.FirstUpper(packageName)), fmt.Sprintf(meta.packageNameF, packageName), meta.groupName); err != nil {
			return err
		}
	}
	return nil
}

// CreatePlug automatically creates plug-in templates
func (autoCodeService *AutoCodeService) CreatePlug(plug system.AutoPlugReq) error {
	// Check whether the list parameters are valid
	plug.CheckList()
	tplFileList, _ := autoCodeService.GetAllTplFile(plugPath, nil)
	for _, tpl := range tplFileList {
		temp, err := template.ParseFiles(tpl)
		if err != nil {
			zap.L().Error("parse err", zap.String("tpl", tpl), zap.Error(err))
			return err
		}
		pathArr := strings.SplitAfter(tpl, "/")
		if strings.Index(pathArr[2], "tpl") < 0 {
			dirPath := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, fmt.Sprintf(global.GVA_CONFIG.AutoCode.SPlug, plug.Snake+"/"+pathArr[2]))
			os.MkdirAll(dirPath, 0755)
		}
		file := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, fmt.Sprintf(global.GVA_CONFIG.AutoCode.SPlug, plug.Snake+"/"+tpl[len(plugPath):len(tpl)-4]))
		f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			zap.L().Error("open file", zap.String("tpl", tpl), zap.Error(err), zap.Any("plug", plug))
			return err
		}
		defer f.Close()

		err = temp.Execute(f, plug)
		if err != nil {
			zap.L().Error("exec err", zap.String("tpl", tpl), zap.Error(err), zap.Any("plug", plug))
			return err
		}
	}
	return nil
}

func (autoCodeService *AutoCodeService) InstallPlugin(file *multipart.FileHeader) (web, server int, err error) {
	const GVAPLUGPINATH = "./gva-plug-temp/"
	defer os.RemoveAll(GVAPLUGPINATH)
	_, err = os.Stat(GVAPLUGPINATH)
	if os.IsNotExist(err) {
		os.Mkdir(GVAPLUGPINATH, os.ModePerm)
	}

	src, err := file.Open()
	if err != nil {
		return -1, -1, err
	}
	defer src.Close()

	out, err := os.Create(GVAPLUGPINATH + file.Filename)
	if err != nil {
		return -1, -1, err
	}
	defer out.Close()

	_, err = io.Copy(out, src)

	paths, err := utils.Unzip(GVAPLUGPINATH+file.Filename, GVAPLUGPINATH)
	paths = filterFile(paths)
	var webIndex = -1
	var serverIndex = -1
	for i := range paths {
		paths[i] = filepath.ToSlash(paths[i])
		pathArr := strings.Split(paths[i], "/")
		ln := len(pathArr)
		if ln < 2 {
			continue
		}
		if pathArr[ln-2] == "server" && pathArr[ln-1] == "plugin" {
			serverIndex = i
		}
		if pathArr[ln-2] == "web" && pathArr[ln-1] == "plugin" {
			webIndex = i
		}
	}
	if webIndex == -1 && serverIndex == -1 {
		zap.L().Error("Non-standard plug-in, please follow the documentation for automatic migration")
		return webIndex, serverIndex, errors.New("Non-standard plug-ins, please follow the documentation for automatic migration")
	}

	if webIndex != -1 {
		err = installation(paths[webIndex], global.GVA_CONFIG.AutoCode.Server, global.GVA_CONFIG.AutoCode.Web)
		if err != nil {
			return webIndex, serverIndex, err
		}
	}

	if serverIndex != -1 {
		err = installation(paths[serverIndex], global.GVA_CONFIG.AutoCode.Server, global.GVA_CONFIG.AutoCode.Server)
	}
	return webIndex, serverIndex, err
}

func installation(path string, formPath string, toPath string) error {
	arr := strings.Split(filepath.ToSlash(path), "/")
	ln := len(arr)
	if ln < 3 {
		return errors.New("arr")
	}
	name := arr[ln-3]

	var form = filepath.ToSlash(global.GVA_CONFIG.AutoCode.Root + formPath + "/" + path)
	var to = filepath.ToSlash(global.GVA_CONFIG.AutoCode.Root + toPath + "/plugin/")
	_, err := os.Stat(to + name)
	if err == nil {
		zap.L().Error("autoPath already has a plug-in with the same name, please install it manually", zap.String("to", to))
		return errors.New(toPath + "A plug-in with the same name already exists, please install it manually")
	}
	return cp.Copy(form, to, cp.Options{Skip: skipMacSpecialDocument})
}

func filterFile(paths []string) []string {
	np := make([]string, 0, len(paths))
	for _, path := range paths {
		if ok, _ := skipMacSpecialDocument(path); ok {
			continue
		}
		np = append(np, path)
	}
	return np
}

func skipMacSpecialDocument(src string) (bool, error) {
	if strings.Contains(src, ".DS_Store") || strings.Contains(src, "__MACOSX") {
		return true, nil
	}
	return false, nil
}

func (autoCodeService *AutoCodeService) PubPlug(plugName string) (zipPath string, err error) {
	if plugName == "" {
		return "", errors.New("Plug-in name cannot be empty")
	}

	// Prevent path crossing
	plugName = filepath.Clean(plugName)

	webPath := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Web, "plugin", plugName)
	serverPath := filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, "plugin", plugName)
	//Create a new zip file

	// Determine whether the directory exists
	_, err = os.Stat(webPath)
	if err != nil {
		return "", errors.New("web path does not exist")
	}
	_, err = os.Stat(serverPath)
	if err != nil {
		return "", errors.New("server path does not exist")
	}

	fileName := plugName + ".zip"
	//Create a new zip file
	zipFile, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer zipFile.Close()

	//Create a zip writer
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	webHeaderName := filepath.Join(plugName, "web", "plugin", plugName)
	err = autoCodeService.doZip(zipWriter, webPath, webHeaderName)
	if err != nil {
		return
	}
	serverHeaderName := filepath.Join(plugName, "server", "plugin", plugName)
	err = autoCodeService.doZip(zipWriter, serverPath, serverHeaderName)
	if err != nil {
		return
	}
	return filepath.Join(global.GVA_CONFIG.AutoCode.Root, global.GVA_CONFIG.AutoCode.Server, fileName), nil
}

/*
*

	zipWriter zip writer
	serverPath stored path
	headerName writes the path of zip

*
*/
func (autoCodeService *AutoCodeService) doZip(zipWriter *zip.Writer, serverPath, headerName string) (err error) {
	// Traverse the serverPath directory and add all non-hidden files to the zip archive
	err = filepath.Walk(serverPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip hidden files and directories
		if strings.HasPrefix(info.Name(), ".") {
			return nil
		}

		//Create a new file header
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		// Set the name of the file header to the relative path of the file
		rel, _ := filepath.Rel(serverPath, path)
		header.Name = filepath.Join(headerName, rel)
		// The directory needs to be spelled with a "/", otherwise a file identical to the directory will appear in the compressed package
		if info.IsDir() {
			header.Name += "/"
		}
		// Add files to zip archive
		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		// Open the file and copy its contents into the zip archive
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = io.Copy(writer, file)
		if err != nil {
			return err
		}

		return nil
	})
	return err
}
