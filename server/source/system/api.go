package system

import (
	"context"
	sysModel "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type initApi struct{}

const initOrderApi = system.InitOrderSystem + 1

// auto run
func init() {
	system.RegisterInit(initOrderApi, &initApi{})
}

func (i initApi) InitializerName() string {
	return sysModel.SysApi{}.TableName()
}

func (i *initApi) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&sysModel.SysApi{})
}

func (i *initApi) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&sysModel.SysApi{})
}

func (i *initApi) InitializeData(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	entities := []sysModel.SysApi{
		{ApiGroup: "jwt", Method: "POST", Path: "/jwt/jsonInBlacklist", Description: "jwt added to blacklist (exit, required)"},

		{ApiGroup: "System User", Method: "DELETE", Path: "/user/deleteUser", Description: "Delete User"},
		{ApiGroup: "System User", Method: "POST", Path: "/user/admin_register", Description: "User Registration"},
		{ApiGroup: "System User", Method: "POST", Path: "/user/getUserList", Description: "Get user list"},
		{ApiGroup: "System User", Method: "PUT", Path: "/user/setUserInfo", Description: "Set User Information"},
		{ApiGroup: "System User", Method: "PUT", Path: "/user/setSelfInfo", Description: "Set self-information (required)"},
		{ApiGroup: "System User", Method: "GET", Path: "/user/getUserInfo", Description: "Get your own information (required)"},
		{ApiGroup: "System User", Method: "POST", Path: "/user/setUserAuthorities", Description: "Set Authorization Group"},
		{ApiGroup: "System User", Method: "POST", Path: "/user/changePassword", Description: "Change Password (Recommended)"},
		{ApiGroup: "System User", Method: "POST", Path: "/user/setUserAuthority", Description: "Modify user role (required)"},
		{ApiGroup: "System User", Method: "POST", Path: "/user/resetPassword", Description: "Reset user password"},

		{ApiGroup: "api", Method: "POST", Path: "/api/createApi", Description: "Create api"},
		{ApiGroup: "api", Method: "POST", Path: "/api/deleteApi", Description: "DeleteApi"},
		{ApiGroup: "api", Method: "POST", Path: "/api/updateApi", Description: "UpdateApi"},
		{ApiGroup: "api", Method: "POST", Path: "/api/getApiList", Description: "Get api list"},
		{ApiGroup: "api", Method: "POST", Path: "/api/getAllApis", Description: "Get all apis"},
		{ApiGroup: "api", Method: "POST", Path: "/api/getApiById", Description: "Get api details"},
		{ApiGroup: "api", Method: "DELETE", Path: "/api/deleteApisByIds", Description: "Batch delete api"},

		{ApiGroup: "Role", Method: "POST", Path: "/authority/copyAuthority", Description: "Copy Role"},
		{ApiGroup: "Role", Method: "POST", Path: "/authority/createAuthority", Description: "Create Role"},
		{ApiGroup: "Role", Method: "POST", Path: "/authority/deleteAuthority", Description: "Delete Role"},
		{ApiGroup: "Role", Method: "PUT", Path: "/authority/updateAuthority", Description: "Update role information"},
		{ApiGroup: "Role", Method: "POST", Path: "/authority/getAuthorityList", Description: "Get role list"},
		{ApiGroup: "Role", Method: "POST", Path: "/authority/setDataAuthority", Description: "Set role resource permissions"},

		{ApiGroup: "casbin", Method: "POST", Path: "/casbin/updateCasbin", Description: "Change role api permissions"},
		{ApiGroup: "casbin", Method: "POST", Path: "/casbin/getPolicyPathByAuthorityId", Description: "Get permission list"},

		{ApiGroup: "Menu", Method: "POST", Path: "/menu/addBaseMenu", Description: "Add Menu"},
		{ApiGroup: "Menu", Method: "POST", Path: "/menu/getMenu", Description: "Get menu tree (required)"},
		{ApiGroup: "Menu", Method: "POST", Path: "/menu/deleteBaseMenu", Description: "Delete Menu"},
		{ApiGroup: "Menu", Method: "POST", Path: "/menu/updateBaseMenu", Description: "Update Menu"},
		{ApiGroup: "Menu", Method: "POST", Path: "/menu/getBaseMenuById", Description: "Get menu based on id"},
		{ApiGroup: "Menu", Method: "POST", Path: "/menu/getMenuList", Description: "Get the basic menu list in pages"},
		{ApiGroup: "Menu", Method: "POST", Path: "/menu/getBaseMenuTree", Description: "Get user dynamic routing"},
		{ApiGroup: "Menu", Method: "POST", Path: "/menu/getMenuAuthority", Description: "Get the specified role menu"},
		{ApiGroup: "Menu", Method: "POST", Path: "/menu/addMenuAuthority", Description: "Add menu and role association"},

		{ApiGroup: "Multiple upload", Method: "GET", Path: "/fileUploadAndDownload/findFile", Description: "Find the target file (secondary transfer)"},
		{ApiGroup: "Multiple Upload", Method: "POST", Path: "/fileUploadAndDownload/breakpointContinue", Description: "Breakpoint Resume"},
		{ApiGroup: "Multiple upload", Method: "POST", Path: "/fileUploadAndDownload/breakpointContinueFinish", Description: "Breakpoint resume completion"},
		{ApiGroup: "Multiple upload", Method: "POST", Path: "/fileUploadAndDownload/removeChunk", Description: "Remove file after upload is complete"},

		{ApiGroup: "File Upload and Download", Method: "POST", Path: "/fileUploadAndDownload/upload", Description: "File Upload Example"},
		{ApiGroup: "File Upload and Download", Method: "POST", Path: "/fileUploadAndDownload/deleteFile", Description: "Delete File"},
		{ApiGroup: "File upload and download", Method: "POST", Path: "/fileUploadAndDownload/editFileName", Description: "File name or note editing"},
		{ApiGroup: "File Upload and Download", Method: "POST", Path: "/fileUploadAndDownload/getFileList", Description: "Get the uploaded file list"},

		{ApiGroup: "System Service", Method: "POST", Path: "/system/getServerInfo", Description: "Get server information"},
		{ApiGroup: "System Service", Method: "POST", Path: "/system/getSystemConfig", Description: "Get the configuration file content"},
		{ApiGroup: "System Service", Method: "POST", Path: "/system/setSystemConfig", Description: "Set configuration file content"},

		{ApiGroup: "Customer", Method: "PUT", Path: "/customer/customer", Description: "Update Customer"},
		{ApiGroup: "Customer", Method: "POST", Path: "/customer/customer", Description: "Create Customer"},
		{ApiGroup: "Customer", Method: "DELETE", Path: "/customer/customer", Description: "Delete Customer"},
		{ApiGroup: "Customer", Method: "GET", Path: "/customer/customer", Description: "Get a single customer"},
		{ApiGroup: "Customer", Method: "GET", Path: "/customer/customerList", Description: "Get customer list"},

		{ApiGroup: "Code Generator", Method: "GET", Path: "/autoCode/getDB", Description: "Get all databases"},
		{ApiGroup: "Code Generator", Method: "GET", Path: "/autoCode/getTables", Description: "Get database tables"},
		{ApiGroup: "Code Generator", Method: "POST", Path: "/autoCode/createTemp", Description: "Automation Code"},
		{ApiGroup: "Code Generator", Method: "POST", Path: "/autoCode/preview", Description: "Preview Automation Code"},
		{ApiGroup: "Code Generator", Method: "GET", Path: "/autoCode/getColumn", Description: "Get all fields of the selected table"},
		{ApiGroup: "Code Generator", Method: "POST", Path: "/autoCode/createPlug", Description: "Automatically create plug-in packages"},
		{ApiGroup: "Code Generator", Method: "POST", Path: "/autoCode/installPlugin", Description: "Install Plugin"},
		{ApiGroup: "Code Generator", Method: "POST", Path: "/autoCode/pubPlug", Description: "Packaging Plug-in"},

		{ApiGroup: "Package (pkg) generator", Method: "POST", Path: "/autoCode/createPackage", Description: "Generate package (package)"},
		{ApiGroup: "Package (pkg) Generator", Method: "POST", Path: "/autoCode/getPackage", Description: "Get all packages (package)"},
		{ApiGroup: "Package (pkg) Generator", Method: "POST", Path: "/autoCode/delPackage", Description: "Delete Package (package)"},

		{ApiGroup: "Code Generator History", Method: "POST", Path: "/autoCode/getMeta", Description: "Get meta information"},
		{ApiGroup: "Code Generator History", Method: "POST", Path: "/autoCode/rollback", Description: "Rollback of automatically generated code"},
		{ApiGroup: "Code Generator History", Method: "POST", Path: "/autoCode/getSysHistory", Description: "Query rollback records"},
		{ApiGroup: "Code Generator History", Method: "POST", Path: "/autoCode/delSysHistory", Description: "Delete rollback record"},

		{ApiGroup: "System Dictionary Details", Method: "PUT", Path: "/sysDictionaryDetail/updateSysDictionaryDetail", Description: "Update dictionary content"},
		{ApiGroup: "System Dictionary Details", Method: "POST", Path: "/sysDictionaryDetail/createSysDictionaryDetail", Description: "New dictionary content"},
		{ApiGroup: "System Dictionary Details", Method: "DELETE", Path: "/sysDictionaryDetail/deleteSysDictionaryDetail", Description: "Delete dictionary content"},
		{ApiGroup: "System Dictionary Details", Method: "GET", Path: "/sysDictionaryDetail/findSysDictionaryDetail", Description: "Get dictionary content based on ID"},
		{ApiGroup: "System dictionary details", Method: "GET", Path: "/sysDictionaryDetail/getSysDictionaryDetailList", Description: "Get dictionary content list"},

		{ApiGroup: "System Dictionary", Method: "POST", Path: "/sysDictionary/createSysDictionary", Description: "New Dictionary"},
		{ApiGroup: "System Dictionary", Method: "DELETE", Path: "/sysDictionary/deleteSysDictionary", Description: "Delete Dictionary"},
		{ApiGroup: "System Dictionary", Method: "PUT", Path: "/sysDictionary/updateSysDictionary", Description: "Update Dictionary"},
		{ApiGroup: "System Dictionary", Method: "GET", Path: "/sysDictionary/findSysDictionary", Description: "Get dictionary based on ID"},
		{ApiGroup: "System Dictionary", Method: "GET", Path: "/sysDictionary/getSysDictionaryList", Description: "Get dictionary list"},

		{ApiGroup: "Operation Record", Method: "POST", Path: "/sysOperationRecord/createSysOperationRecord", Description: "New Operation Record"},
		{ApiGroup: "Operation Record", Method: "GET", Path: "/sysOperationRecord/findSysOperationRecord", Description: "Get operation record based on ID"},
		{ApiGroup: "Operation Record", Method: "GET", Path: "/sysOperationRecord/getSysOperationRecordList", Description: "Get the operation record list"},
		{ApiGroup: "Operation Record", Method: "DELETE", Path: "/sysOperationRecord/deleteSysOperationRecord", Description: "Delete Operation Record"},
		{ApiGroup: "Operation Record", Method: "DELETE", Path: "/sysOperationRecord/deleteSysOperationRecordByIds", Description: "Batch delete operation history"},

		{ApiGroup: "Resumable upload (plug-in version)", Method: "POST", Path: "/simpleUploader/upload", Description: "Plug-in version multi-part upload"},
		{ApiGroup: "Resumable upload (plug-in version)", Method: "GET", Path: "/simpleUploader/checkFileMd5", Description: "File integrity verification"},
		{ApiGroup: "Resumable upload (plug-in version)", Method: "GET", Path: "/simpleUploader/mergeFileMd5", Description: "Upload completed merged file"},

		{ApiGroup: "email", Method: "POST", Path: "/email/emailTest", Description: "Send test email"},
		{ApiGroup: "email", Method: "POST", Path: "/email/emailSend", Description: "Send email example"},

		{ApiGroup: "Button authority", Method: "POST", Path: "/authorityBtn/setAuthorityBtn", Description: "Set button authority"},
		{ApiGroup: "Button authority", Method: "POST", Path: "/authorityBtn/getAuthorityBtn", Description: "Get existing button authority"},
		{ApiGroup: "Button Authority", Method: "POST", Path: "/authorityBtn/canRemoveAuthorityBtn", Description: "Delete Button"},

		{ApiGroup: "Table Template", Method: "POST", Path: "/sysExportTemplate/createSysExportTemplate", Description: "New export template"},
		{ApiGroup: "Table Template", Method: "DELETE", Path: "/sysExportTemplate/deleteSysExportTemplate", Description: "Delete Export Template"},
		{ApiGroup: "Table Template", Method: "DELETE", Path: "/sysExportTemplate/deleteSysExportTemplateByIds", Description: "Batch delete export templates"},
		{ApiGroup: "Table Template", Method: "PUT", Path: "/sysExportTemplate/updateSysExportTemplate", Description: "Update Export Template"},
		{ApiGroup: "Table Template", Method: "GET", Path: "/sysExportTemplate/findSysExportTemplate", Description: "Get the export template based on ID"},
		{ApiGroup: "Table Template", Method: "GET", Path: "/sysExportTemplate/getSysExportTemplateList", Description: "Get export template list"},
		{ApiGroup: "Table Template", Method: "GET", Path: "/sysExportTemplate/exportExcel", Description: "Export Excel"},
		{ApiGroup: "Form Template", Method: "GET", Path: "/sysExportTemplate/exportTemplate", Description: "Download Template"},
		{ApiGroup: "Table Template", Method: "POST", Path: "/sysExportTemplate/importExcel", Description: "Import Excel"},
	}
	if err := db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, sysModel.SysApi{}.TableName()+"Table data initialization failed!")
	}
	next := context.WithValue(ctx, i.InitializerName(), entities)
	return next, nil
}

func (i *initApi) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.Where("path = ? AND method = ?", "/authorityBtn/canRemoveAuthorityBtn", "POST").
		First(&sysModel.SysApi{}).Error, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}
