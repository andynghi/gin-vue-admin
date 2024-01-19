package initialize

import (
	swaggerFiles "github.com/swaggo/files"
	"net/http"

	"github.com/flipped-aurora/gin-vue-admin/server/docs"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/router"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//Initialize the general route

func Routers() *gin.Engine {

	// Set to publish mode
	if global.GVA_CONFIG.System.Env == "public" {
		gin.SetMode(gin.ReleaseMode) //DebugMode ReleaseMode TestMode
	}

	Router := gin.New()
	Router.Use(gin.Recovery())
	if global.GVA_CONFIG.System.Env != "public" {
		Router.Use(gin.Logger())
	}

	InstallPlugin(Router) // Install plug-in
	systemRouter := router.RouterGroupApp.System
	exampleRouter := router.RouterGroupApp.Example
	// If you don’t want to use nginx to proxy the front-end web page, you can modify it under web/.env.production
	// VUE_APP_BASE_API = /
	// VUE_APP_BASE_PATH = http://localhost
	// Then execute the packaging command npm run build. Open the following 3 lines of comments
	//Router.Static("/favicon.ico", "./dist/favicon.ico")
	//Router.Static("/static", "./dist/static") // Static resources in dist
	//Router.StaticFile("/", "./dist/index.html") // Front-end web entrance page

	Router.StaticFS(global.GVA_CONFIG.Local.StorePath, http.Dir(global.GVA_CONFIG.Local.StorePath)) // Provide static addresses for user avatars and files
	// Router.Use(middleware.LoadTls()) // If you need to use https, please open this middleware and go to core/server.go to change the startup mode to Router.RunTLS("port","your cre/pem file","your key file")
	// Cross-domain, if you need to cross-domain, you can open the following comments
	// Router.Use(middleware.Cors()) // Directly allow all cross-domain requests
	// Router.Use(middleware.CorsByRules()) // Release cross-domain requests according to configured rules
	//global.GVA_LOG.Info("use middleware cors")
	docs.SwaggerInfo.BasePath = global.GVA_CONFIG.System.RouterPrefix
	Router.GET(global.GVA_CONFIG.System.RouterPrefix+"/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.GVA_LOG.Info("register swagger handler")
	// Conveniently add routing group prefix for multiple servers to go online.

	PublicGroup := Router.Group(global.GVA_CONFIG.System.RouterPrefix)
	{
		// health monitoring
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}
	{
		systemRouter.InitBaseRouter(PublicGroup) //Register basic function routing without authentication
		systemRouter.InitInitRouter(PublicGroup) // Automatic initialization related
	}
	PrivateGroup := Router.Group(global.GVA_CONFIG.System.RouterPrefix)
	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		systemRouter.InitApiRouter(PrivateGroup, PublicGroup)       //Register function api routing
		systemRouter.InitJwtRouter(PrivateGroup)                    // jwt related routing
		systemRouter.InitUserRouter(PrivateGroup)                   //Register user route
		systemRouter.InitMenuRouter(PrivateGroup)                   //Register menu route
		systemRouter.InitSystemRouter(PrivateGroup)                 // system related routing
		systemRouter.InitCasbinRouter(PrivateGroup)                 // Permission-related routing
		systemRouter.InitAutoCodeRouter(PrivateGroup)               // Create automation code
		systemRouter.InitAuthorityRouter(PrivateGroup)              //Register role route
		systemRouter.InitSysDictionaryRouter(PrivateGroup)          // Dictionary management
		systemRouter.InitAutoCodeHistoryRouter(PrivateGroup)        // Automated code history
		systemRouter.InitSysOperationRecordRouter(PrivateGroup)     // Operation record
		systemRouter.InitSysDictionaryDetailRouter(PrivateGroup)    // Dictionary details management
		systemRouter.InitAuthorityBtnRouterRouter(PrivateGroup)     // Dictionary details management
		systemRouter.InitSysExportTemplateRouter(PrivateGroup)      // Export template
		exampleRouter.InitCustomerRouter(PrivateGroup)              //Customer routing
		exampleRouter.InitFileUploadAndDownloadRouter(PrivateGroup) //File upload and download function routing

	}

	global.GVA_LOG.Info("router register success")
	return Router
}
