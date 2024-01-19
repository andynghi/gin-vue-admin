package plugin

import (
	"github.com/gin-gonic/gin"
)

const (
	OnlyFuncName = "Plugin"
)

// Plugin plug-in mode interface
type Plugin interface {
	// Register register route
	Register(group *gin.RouterGroup)

	// RouterPath user returns the registered route
	RouterPath() string
}
