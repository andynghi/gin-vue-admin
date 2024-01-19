package core

import (
	"fmt"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/initialize"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	if global.GVA_CONFIG.System.UseMultipoint || global.GVA_CONFIG.System.UseRedis {
		//Initialize redis service
		initialize.Redis()
	}
	if global.GVA_CONFIG.System.UseMongo {
		err := initialize.Mongo.Initialization()
		if err != nil {
			zap.L().Error(fmt.Sprintf("%+v", err))
		}
	}
	//Load jwt data from db
	if global.GVA_DB != nil {
		system.LoadAll()
	}

	Router := initialize.Routers()
	Router.Static("/form-generator", "./resource/page")

	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	s := initServer(address, Router)
	// Ensure text is output in order
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.GVA_LOG.Info("server run success on ", zap.String("address", address))

	fmt.Printf(`
	Welcome to gin-vue-admin
	Current version: v2.5.9
    How to join the group: WeChat ID: shouzi_1994 QQ group: 622360840
	Plug-in market: https://plugin.gin-vue-admin.com
	GVA discussion community: https://support.qq.com/products/371961
	Default automation document address: http://127.0.0.1%s/swagger/index.html
	Default front-end file running address: http://127.0.0.1:8080
	If the project has made you gain, I hope you can treat the team to a cup of Coke: https://www.gin-vue-admin.com/coffee/index.html
`, address)
	global.GVA_LOG.Error(s.ListenAndServe().Error())
}
