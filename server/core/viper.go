package core

import (
	"flag"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/core/internal"
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	_ "github.com/flipped-aurora/gin-vue-admin/server/packfile"
)

// Viper //
// Priority: Command line > Environment variables > Default value
// Author [SliverHorn](https://github.com/SliverHorn)
func Viper(path ...string) *viper.Viper {
	var config string

	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose config file.")
		flag.Parse()
		if config == "" { // Determine whether the command line parameters are empty
			if configEnv := os.Getenv(internal.ConfigEnv); configEnv == "" { // Determine whether the environment variable stored in the internal.ConfigEnv constant is empty
				switch gin.Mode() {
				case gin.DebugMode:
					config = internal.ConfigDefaultFile
					fmt.Printf("You are using the %s environment name in gin mode, the path to config is %s\n", gin.EnvGinMode, internal.ConfigDefaultFile)
				case gin.ReleaseMode:
					config = internal.ConfigReleaseFile
					fmt.Printf("You are using the %s environment name in gin mode, the path to config is %s\n", gin.EnvGinMode, internal.ConfigReleaseFile)
				case gin.TestMode:
					config = internal.ConfigTestFile
					fmt.Printf("You are using the %s environment name in gin mode, the path to config is %s\n", gin.EnvGinMode, internal.ConfigTestFile)
				}
			} else { // internal.ConfigEnv The environment variable stored in the constant is not empty and assigns the value to config
				config = configEnv
				fmt.Printf("You are using %s environment variable, the path of config is %s\n", internal.ConfigEnv, config)
			}
		} else { // The command line parameters are not empty and assign the value to config
			fmt.Printf("You are using the value passed by the -c parameter on the command line, the path to config is %s\n", config)
		}
	} else { // The first value of the variable parameter passed by the function is assigned to config
		config = path[0]
		fmt.Printf("You are using the value passed by func Viper(), the path of config is %s\n", config)
	}

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&global.GVA_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&global.GVA_CONFIG); err != nil {
		panic(err)
	}

	// Root adaptability Find the corresponding migration location based on the root location to ensure that the root path is valid
	global.GVA_CONFIG.AutoCode.Root, _ = filepath.Abs("..")
	return v
}
