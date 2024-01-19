// Automatically generate template {{.StructName}}
package {{.Package}}

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	{{ if .HasTimer }}"time"{{ end }}
	{{ if .NeedJSON }}"gorm.io/datatypes"{{ end }}
)

// {{.Description}} structure {{.StructName}}
type {{.StructName}} struct {
      global.GVA_MODEL {{- range .Fields}}
            {{- if eq .FieldType "enum" }}
      {{.FieldName}}  string `json:"{{.FieldJson}}" form:"{{.FieldJson}}" gorm:"column:{{.ColumnName}};type:enum({{.DataTypeLong}});comment:{{.Comment}};" {{- if .Require -}}binding:"required"{{- end -}}`
            {{- else if eq .FieldType "picture" }}
      {{.FieldName}}  string `json:"{{.FieldJson}}" form:"{{.FieldJson}}" gorm:"column:{{.ColumnName}};comment:{{.Comment}};{{- if .DataTypeLong -}}size:{{.DataTypeLong}};{{- end -}}" {{- if .Require -}}binding:"required"{{- end -}}`
            {{- else if eq .FieldType "video" }}
      {{.FieldName}}  string `json:"{{.FieldJson}}" form:"{{.FieldJson}}" gorm:"column:{{.ColumnName}};comment:{{.Comment}};{{- if .DataTypeLong -}}size:{{.DataTypeLong}};{{- end -}}" {{- if .Require -}}binding:"required"{{- end -}}`
             {{- else if eq .FieldType "file" }}
      {{.FieldName}}  datatypes.JSON `json:"{{.FieldJson}}" form:"{{.FieldJson}}" gorm:"column:{{.ColumnName}};comment:{{.Comment}};{{- if .DataTypeLong -}}size:{{.DataTypeLong}};{{- end -}}" {{- if .Require -}}binding:"required"{{- end -}}`
            {{- else if eq .FieldType "pictures" }}
      {{.FieldName}}  datatypes.JSON `json:"{{.FieldJson}}" form:"{{.FieldJson}}" gorm:"column:{{.ColumnName}};comment:{{.Comment}};{{- if .DataTypeLong -}}size:{{.DataTypeLong}};{{- end -}}" {{- if .Require -}}binding:"required"{{- end -}}`
            {{- else if eq .FieldType "richtext" }}
      {{.FieldName}}  string `json:"{{.FieldJson}}" form:"{{.FieldJson}}" gorm:"column:{{.ColumnName}};comment:{{.Comment}};{{- if .DataTypeLong -}}size:{{.DataTypeLong}};{{- end -}}type:text;" {{- if .Require -}}binding:"required"{{- end -}}`
            {{- else if ne .FieldType "string" }}
      {{.FieldName}}  *{{.FieldType}} `json:"{{.FieldJson}}" form:"{{.FieldJson}}" gorm:"column:{{.ColumnName}};comment:{{.Comment}};{{- if .DataTypeLong -}}size:{{.DataTypeLong}};{{- end -}}" {{- if .Require -}}binding:"required"{{- end -}}`
            {{- else }}
      {{.FieldName}}  {{.FieldType}} `json:"{{.FieldJson}}" form:"{{.FieldJson}}" gorm:"column:{{.ColumnName}};comment:{{.Comment}};{{- if .DataTypeLong -}}size:{{.DataTypeLong}};{{- end -}}" {{- if .Require -}}binding:"required"{{- end -}}`
            {{- end }}  {{ if .FieldDesc }}//{{.FieldDesc}} {{ end }} {{- end }}
      {{- if .AutoCreateResource }}
      CreatedBy uint `gorm:"column:created_by;comment:creator"`
      UpdatedBy uint `gorm:"column:updated_by;comment:updater"`
      DeletedBy uint `gorm:"column:deleted_by;comment:deleted"`
      {{- end}}
}

{{ if .TableName }}
// TableName {{.Description}} {{.StructName}} Custom table name {{.TableName}}
func ({{.StructName}}) TableName() string {
  return "{{.TableName}}"
}
{{ end }}
