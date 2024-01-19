package ws

import (
	"github.com/flipped-aurora/ws/core/biz"
	"github.com/flipped-aurora/ws/core/data"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"nhooyr.io/websocket"
)

type wsPlugin struct {
	logger               *zap.Logger                       // Log output object
	manageBuf            int64                             // buffer
	registeredMsgHandler map[int32]func(biz.IMessage) bool // Message processing
	checkMap             map[string]biz.CheckFunc          // User verification

	admin     biz.IManage
	adminCase *biz.AdminCase
}

func DefaultRegisteredMsgHandler(admin biz.IManage, logger *zap.Logger) map[int32]func(biz.IMessage) bool {
	return map[int32]func(msg biz.IMessage) bool{
		1: func(msg biz.IMessage) bool {
			// Find the method to register the client in w.admin
			client, ok := admin.FindClient(msg.GetTo())
			if !ok {
				logger.Info("The user was not found")
				return false
			}
			return client.SendMes(msg)
		},
	}
}

func DefaultCheckMap() map[string]biz.CheckFunc {
	return map[string]biz.CheckFunc{
		"gva_ws": func(c interface{}) (string, bool) {
			// First assert that it is gin.content
			cc, ok := c.(*gin.Context)
			if !ok {
				return "", false
			}
			token := cc.Query("jwt")
			// can carry jwt
			if len(token) == 0 {
				return "", false
			}
			// Parse jwt...

			return token, true
		},
	}
}

func (w *wsPlugin) Register(g *gin.RouterGroup) {
	// gva_ws is the identity verification function
	g.GET("/ws", w.adminCase.HandlerWS("gva_ws", &websocket.AcceptOptions{
		InsecureSkipVerify: true,
	}))
	g.POST("/sendMsg", w.adminCase.SendMsg("gva_ws"))
}

func (w *wsPlugin) RouterPath() string {
	return "gva_ws"
}

func GenerateWs(logger *zap.Logger, manageBuf int64, checkMap map[string]biz.CheckFunc) *wsPlugin {
	m := data.NewManage(manageBuf)
	t := data.NewTopic()
	h := data.NewHandle()
	admin := data.NewAdmin(m, t, h, logger)
	for s, checkFunc := range checkMap {
		admin.AddCheckFunc(s, checkFunc)
	}
	registeredMsgHandler := DefaultRegisteredMsgHandler(admin, logger)

	for key, handler := range registeredMsgHandler {
		admin.RegisteredMsgHandler(key, handler)
	}
	return &wsPlugin{
		logger: logger, manageBuf: manageBuf,
		registeredMsgHandler: registeredMsgHandler, checkMap: checkMap, admin: admin, adminCase: biz.NewAdmin(admin),
	}
}
