// Automatically generate template SysOperationRecord
package system

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// If it contains time.Time, please import the time package yourself
type SysOperationRecord struct {
	global.GVA_MODEL
	Ip           string        `json:"ip" from:"ip" gorm:"column:ip;comment:request ip"`                                     // request ip
	Method       string        `json:"method" form:"method" gorm:"column:method;comment:request method"`                     // Request method
	Path         string        `json:"path" form:"path" gorm:"column:path;comment:request path"`                             // request path
	Status       int           `json:"status" form:"status" gorm:"column:status;comment:request status"`                     // Request status
	Latency      time.Duration `json:"latency" form:"latency" gorm:"column:latency;comment:latency" swaggertype:"string"`    // Latency
	Agent        string        `json:"agent" form:"agent" gorm:"column:agent;comment:agent"`                                 // Agent
	ErrorMessage string        `json:"error_message" form:"error_message" gorm:"column:error_message;comment:error message"` // Error message
	Body         string        `json:"body" form:"body" gorm:"type:text;column:body;comment:Request Body"`                   // Request Body
	Resp         string        `json:"resp" form:"resp" gorm:"type:text;column:resp;comment:Response Body"`                  //Response Body
	UserID       int           `json:"user_id" form:"user_id" gorm:"column:user_id;comment:userid"`                          // user id
	User         SysUser       `json:"user"`
}
