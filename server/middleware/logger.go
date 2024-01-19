package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// LogLayout log layout
type LogLayout struct {
	Time      time.Time
	Metadata  map[string]interface{} // Store custom original data
	Path      string                 // access path
	Query     string                 // carries query
	Body      string                 // carries body data
	IP        string                 //ip address
	UserAgent string                 //Agent
	Error     string                 // error
	Cost      time.Duration          // spend time
	Source    string                 // source
}

type Logger struct {
	// Filter user-defined filter
	Filter func(c *gin.Context) bool
	// FilterKeyword keyword filter (key)
	FilterKeyword func(layout *LogLayout) bool
	// AuthProcess authentication processing
	AuthProcess func(c *gin.Context, layout *LogLayout)
	//Log processing
	Print func(LogLayout)
	// Source service unique identifier
	Source string
}

func (l Logger) SetLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		var body []byte
		if l.Filter != nil && !l.Filter(c) {
			body, _ = c.GetRawData()
			// Stuff the original body back
			c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
		}
		c.Next()
		cost := time.Since(start)
		layout := LogLayout{
			Time:      time.Now(),
			Path:      path,
			Query:     query,
			IP:        c.ClientIP(),
			UserAgent: c.Request.UserAgent(),
			Error:     strings.TrimRight(c.Errors.ByType(gin.ErrorTypePrivate).String(), "\n"),
			Cost:      cost,
			Source:    l.Source,
		}
		if l.Filter != nil && !l.Filter(c) {
			layout.Body = string(body)
		}
		if l.AuthProcess != nil {
			// Process the information required for authentication
			l.AuthProcess(c, &layout)
		}
		if l.FilterKeyword != nil {
			// Determine key/value by yourself, desensitize etc.
			l.FilterKeyword(&layout)
		}
		// Process logs by yourself
		l.Print(layout)
	}
}

func DefaultLogger() gin.HandlerFunc {
	return Logger{
		Print: func(layout LogLayout) {
			//Standard output, k8s does collection
			v, _ := json.Marshal(layout)
			fmt.Println(string(v))
		},
		Source: "GVA",
	}.SetLoggerMiddleware()
}
