<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">
      <el-form-item label="Creation date" prop="createdAt">
      <template #label>
        <span>
          Creation date
          <el-tooltip content="The search range is from start date (inclusive) to end date (exclusive)">
            <el-icon><QuestionFilled /></el-icon>
          </el-tooltip>
        </span>
      </template>
      <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="Start Date" :disabled-date="time=> searchInfo.endCreatedAt ? time.getTime() > searchInfo.endCreatedAt.getTime () : false"></el-date-picker>
       —
      <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="End Date" :disabled-date="time=> searchInfo.startCreatedAt ? time.getTime() < searchInfo.startCreatedAt.getTime () : false"></el-date-picker>
      </el-form-item>
           {{- range .Fields}}  {{- if .FieldSearchType}} {{- if eq .FieldType "bool" }}
            <el-form-item label="{{.FieldDesc}}" prop="{{.FieldJson}}">
            <el-select v-model="searchInfo.{{.FieldJson}}" clearable placeholder="Please select">
                <el-option
                    key="true"
                    label="yes"
                    value="true">
                </el-option>
                <el-option
                    key="false"
                    label="No"
                    value="false">
                </el-option>
            </el-select>
            </el-form-item>
           {{- else if .DictType}}
           <el-form-item label="{{.FieldDesc}}" prop="{{.FieldJson}}">
            <el-select v-model="searchInfo.{{.FieldJson}}" clearable placeholder="Please select" @clear="()=>{searchInfo.{{.FieldJson}}=undefined}">
              <el-option v-for="(item,key) in {{ .DictType }}Options" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
            {{- else}}
        <el-form-item label="{{.FieldDesc}}" prop="{{.FieldJson}}">


        {{- if eq .FieldType "float64" "int"}}
            {{if eq .FieldSearchType "BETWEEN" "NOT BETWEEN"}}
            <el-input v-model.number="searchInfo.start{{.FieldName}}" placeholder="minimum value" />
            —
            <el-input v-model.number="searchInfo.end{{.FieldName}}" placeholder="maximum value" />
           {{- else}}
             {{- if .DictType}}
              <el-select v-model="searchInfo.{{.FieldJson}}" placeholder="Please select" style="width:100%" :clearable="true" >
               <el-option v-for="(item,key) in {{ .DictType }}Options" :key="key" :label="item.label" :value="item.value" />
             </el-select>
                    {{- else}}
             <el-input v-model.number="searchInfo.{{.FieldJson}}" placeholder="Search criteria" />
                    {{- end }}
          {{- end}}
        {{- else if eq .FieldType "time.Time"}}
            {{if eq .FieldSearchType "BETWEEN" "NOT BETWEEN"}}
            <template #label>
            <span>
              {{.FieldDesc}}
              <el-tooltip content="The search range is from start date (inclusive) to end date (exclusive)">
                <el-icon><QuestionFilled /></el-icon>
              </el-tooltip>
            </span>
          </template>
            <el-date-picker v-model="searchInfo.start{{.FieldName}}" type="datetime" placeholder="Start Date" :disabled-date="time=> searchInfo.end{{.FieldName}} ? time.getTime() > searchInfo.end{{.FieldName}}.getTime() : false"></el-date-picker>
            —
            <el-date-picker v-model="searchInfo.end{{.FieldName}}" type="datetime" placeholder="End Date" :disabled-date="time=> searchInfo.start{{.FieldName}} ? time.getTime() < searchInfo.start{{.FieldName}}.getTime() : false"></el-date-picker>
           {{- else}}
           <el-date-picker v-model="searchInfo.{{.FieldJson}}" type="datetime" placeholder="search criteria"></el-date-picker>
          {{- end}}
        {{- else}}
         <el-input v-model="searchInfo.{{.FieldJson}}" placeholder="Search criteria" />
        {{- end}}

        </el-form-item>{{ end }}{{ end }}{{ end }}
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">Query</el-button>
          <el-button icon="refresh" @click="onReset">Reset</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button type="primary" icon="plus" @click="openDialog">New</el-button>
            <el-popover v-model:visible="deleteVisible" :disabled="!multipleSelection.length" placement="top" width="160">
            <p>Are you sure you want to delete it? </p>
            <div style="text-align: right; margin-top: 8px;">
                <el-button type="primary" link @click="deleteVisible = false">Cancel</el-button>
                <el-button type="primary" @click="onDelete">OK</el-button>
            </div>
            <template #reference>
                <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="deleteVisible = true">Delete</el-button>
            </template>
            </el-popover>
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        {{- if .NeedSort}}
        @sort-change="sortChange"
        {{- end}}
        >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="Date" width="180">
            <template #default="scope">{{ "{{ formatDate(scope.row.CreatedAt) }}" }}</template>
        </el-table-column>
        {{- range .Fields}}
        {{- if .DictType}}
        <el-table-column {{- if .Sort}} sortable{{- end}} align="left" label="{{.FieldDesc}}" prop="{{.FieldJson}}" width="120">
            <template #default="scope">
            {{"{{"}} filterDict(scope.row.{{.FieldJson}},{{.DictType}}Options) {{"}}"}}
            </template>
        </el-table-column>
        {{- else if eq .FieldType "bool" }}
        <el-table-column {{- if .Sort}} sortable{{- end}} align="left" label="{{.FieldDesc}}" prop="{{.FieldJson}}" width="120">
            <template #default="scope">{{"{{"}} formatBoolean(scope.row.{{.FieldJson}}) {{"}}"}}</template>
        </el-table-column>
         {{- else if eq .FieldType "time.Time" }}
         <el-table-column {{- if .Sort}} sortable{{- end}} align="left" label="{{.FieldDesc}}" width="180">
            <template #default="scope">{{"{{"}} formatDate(scope.row.{{.FieldJson}}) {{"}}"}}</template>
         </el-table-column>
          {{- else if eq .FieldType "picture" }}
          <el-table-column label="{{.FieldDesc}}" width="200">
              <template #default="scope">
                <el-image style="width: 100px; height: 100px" :src="getUrl(scope.row.{{.FieldJson}})" fit="cover"/>
              </template>
          </el-table-column>
           {{- else if eq .FieldType "pictures" }}
           <el-table-column label="{{.FieldDesc}}" width="200">
              <template #default="scope">
                 <div class="multiple-img-box">
                    <el-image v-for="(item,index) in scope.row.{{.FieldJson}}" style="width: 80px; height: 80px" :src="getUrl(item)" fit="cover"/>
                </div>
              </template>
           </el-table-column>
           {{- else if eq .FieldType "video" }}
           <el-table-column label="{{.FieldDesc}}" width="200">
              <template #default="scope">
               <video
                  style="width: 100px; height: 100px"
                  muted
                  preload="metadata"
                  >
                    <source :src="getUrl(scope.row.{{.FieldJson}}) + '#t=1'">
                  </video>
              </template>
           </el-table-column>
           {{- else if eq .FieldType "richtext" }}
                      <el-table-column label="{{.FieldDesc}}" width="200">
                         <template #default="scope">
                            [Rich text content]
                         </template>
                      </el-table-column>
           {{- else if eq .FieldType "file" }}
                    <el-table-column label="{{.FieldDesc}}" width="200">
                        <template #default="scope">
                             <div class="file-list">
                               <el-tag v-for="file in scope.row.{{.FieldJson}}" :key="file.uid">{{"{{"}}file.name{{"}}"}}</el-tag>
                             </div>
                        </template>
                    </el-table-column>
        {{- else }}
        <el-table-column {{- if .Sort}} sortable{{- end}} align="left" label="{{.FieldDesc}}" prop="{{.FieldJson}}" width="120" />
        {{- end }}
        {{- end }}
        <el-table-column align="left" label="Operation" min-width="120">
            <template #default="scope">
            <el-button type="primary" link class="table-button" @click="getDetails(scope.row)">
                <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>
                check the details
            </el-button>
            <el-button type="primary" link icon="edit" class="table-button" @click="update{{.StructName}}Func(scope.row)">Change</el-button>
            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">Delete</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" :title="type==='create'?'Add':'Modify'" destroy-on-close>
      <el-scrollbar height="500px">
          <el-form :model="formData" label-position="right" ref="elFormRef" :rules="rule" label-width="80px">
        {{- range .Fields}}
            <el-form-item label="{{.FieldDesc}}:"  prop="{{.FieldJson}}" >
          {{- if eq .FieldType "bool" }}
              <el-switch v-model="formData.{{.FieldJson}}" active-color="#13ce66" inactive-color="#ff4949" active-text="Yes" inactive-text="No" clearable > </el-switch>
          {{- end }}
          {{- if eq .FieldType "string" }}
              <el-input v-model="formData.{{.FieldJson}}" :clearable="{{.Clearable}}" placeholder="Please enter {{.FieldDesc}}" />
          {{- end }}
          {{- if eq .FieldType "richtext" }}
              <RichEdit v-model="formData.{{.FieldJson}}"/>
          {{- end }}
          {{- if eq .FieldType "int" }}
          {{- if .DictType}}
              <el-select v-model="formData.{{ .FieldJson }}" placeholder="Please select {{.FieldDesc}}" style="width:100%" :clearable="{{.Clearable}}" >
                <el-option v-for="(item,key) in {{ .DictType }}Options" :key="key" :label="item.label" :value="item.value" />
              </el-select>
          {{- else }}
              <el-input v-model.number="formData.{{ .FieldJson }}" :clearable="{{.Clearable}}" placeholder="Please enter {{.FieldDesc}}" />
          {{- end }}
          {{- end }}
          {{- if eq .FieldType "time.Time" }}
              <el-date-picker v-model="formData.{{ .FieldJson }}" type="date" style="width:100%" placeholder="Select Date" :clearable="{{.Clearable}}" />
          {{- end }}
          {{- if eq .FieldType "float64" }}
              <el-input-number v-model="formData.{{ .FieldJson }}"  style="width:100%" :precision="2" :clearable="{{.Clearable}}"  />
          {{- end }}
          {{- if eq .FieldType "enum" }}
                <el-select v-model="formData.{{ .FieldJson }}" placeholder="Please select {{.FieldDesc}}" style="width:100%" :clearable="{{.Clearable}}" >
                   <el-option v-for="item in [{{.DataTypeLong}}]" :key="item" :label="item" :value="item" />
                </el-select>
          {{- end }}
          {{- if eq .FieldType "picture" }}
                <SelectImage
                 v-model="formData.{{ .FieldJson }}"
                 file-type="image"
                />
          {{- end }}
          {{- if eq .FieldType "pictures" }}
                <SelectImage
                 multiple
                 v-model="formData.{{ .FieldJson }}"
                 file-type="image"
                 />
          {{- end }}
          {{- if eq .FieldType "video" }}
                <SelectImage
                v-model="formData.{{ .FieldJson }}"
                file-type="video"
                />
           {{- end }}
          {{- if eq .FieldType "file" }}
                <SelectFile v-model="formData.{{ .FieldJson }}" />
          {{- end }}
            </el-form-item>
          {{- end }}
          </el-form>
      </el-scrollbar>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">Cancel</el-button>
          <el-button type="primary" @click="enterDialog">Confirm</el-button>
        </div>
      </template>
    </el-dialog>

    <el-dialog v-model="detailShow" style="width: 800px" lock-scroll :before-close="closeDetailShow" title="View details" destroy-on-close>
      <el-scrollbar height="550px">
        <el-descriptions column="1" border>
        {{- range .Fields}}
                <el-descriptions-item label="{{ .FieldDesc }}">
                {{- if .DictType}}
                        {{"{{"}} filterDict(formData.{{.FieldJson}},{{.DictType}}Options) {{"}}"}}
                {{- else if eq .FieldType "picture" }}
                        <el-image style="width: 50px; height: 50px" :preview-src-list="ReturnArrImg(formData.{{ .FieldJson }})" :src="getUrl(formData.{{ .FieldJson }})" fit="cover" />
                {{- else if eq .FieldType "video" }}
                        <video
                              style="width: 50px; height: 50px"
                              muted
                              preload="metadata"
                            >
                            <source :src="getUrl(formData.{{ .FieldJson }}) + '#t=1'">
                        </video>
                {{- else if eq .FieldType "pictures" }}
                        <el-image style="width: 50px; height: 50px; margin-right: 10px" :preview-src-list="ReturnArrImg(formData.{{ .FieldJson }})" :initial-index="index" v-for="(item,index) in formData.{{ .FieldJson }}" :key="index" :src="getUrl(item)" fit="cover" />
                {{- else if eq .FieldType "file" }}
                        <div class="fileBtn" v-for="(item,index) in formData.{{ .FieldJson }}" :key="index">
                          <el-button type="primary" text bg @click="onDownloadFile(item.url)">
                            <el-icon style="margin-right: 5px"><Download /></el-icon>
                            {{"{{"}} item.name {{"}}"}}
                          </el-button>
                        </div>
                  {{- else if eq .FieldType "bool" }}
                    {{"{{"}} formatBoolean(formData.{{.FieldJson}}) {{"}}"}}
                   {{- else if eq .FieldType "time.Time" }}
                      {{"{{"}} formatDate(formData.{{.FieldJson}}) {{"}}"}}
                   {{- else if eq .FieldType "richtext" }}
                        [Rich text content]
                   {{- else}}
                        {{"{{"}} formData.{{.FieldJson}} {{"}}"}}
                   {{- end }}
                </el-descriptions-item>
        {{- end }}
        </el-descriptions>
      </el-scrollbar>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  create{{.StructName}},
  delete{{.StructName}},
  delete{{.StructName}}ByIds,
  update{{.StructName}},
  find{{.StructName}},
  get{{.StructName}}List
} from '@/api/{{.PackageName}}'

{{- if or .HasPic .HasFile}}
import { getUrl } from '@/utils/image'
{{- end }}
{{- if .HasPic }}
//Image selection component
import SelectImage from '@/components/selectImage/selectImage.vue'
{{- end }}

{{- if .HasRichText }}
//Rich text component
import RichEdit from '@/components/richtext/rich-edit.vue'
{{- end }}


{{- if .HasFile }}
//File selection component
import SelectFile from '@/components/selectFile/selectFile.vue'
{{- end }}

// Fully introduced formatting tools, please keep them as needed
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
    name: '{{.StructName}}'
})

//Automatically generated dictionary (may be empty) and fields
    {{- range $index, $element := .DictTypes}}
const {{ $element }}Options = ref([])
    {{- end }}
const formData = ref({
        {{- range .Fields}}
        {{- if eq .FieldType "bool" }}
        {{.FieldJson}}: false,
        {{- end }}
        {{- if eq .FieldType "string" }}
        {{.FieldJson}}: '',
        {{- end }}
        {{- if eq .FieldType "richtext" }}
        {{.FieldJson}}: '',
        {{- end }}
        {{- if eq .FieldType "int" }}
        {{.FieldJson}}: {{- if .DictType }} undefined{{ else }} 0{{- end }},
        {{- end }}
        {{- if eq .FieldType "time.Time" }}
        {{.FieldJson}}: new Date(),
        {{- end }}
        {{- if eq .FieldType "float64" }}
        {{.FieldJson}}: 0,
        {{- end }}
        {{- if eq .FieldType "picture" }}
        {{.FieldJson}}: "",
        {{- end }}
        {{- if eq .FieldType "video" }}
        {{.FieldJson}}: "",
        {{- end }}
        {{- if eq .FieldType "pictures" }}
        {{.FieldJson}}: [],
        {{- end }}
        {{- if eq .FieldType "file" }}
        {{.FieldJson}}: [],
        {{- end }}
        {{- end }}
        })


// Validation rules
const rule = reactive({
    {{- range .Fields }}
            {{- if eq .Require true }}
               {{.FieldJson }} : [{
                   required: true,
                   message: '{{ .ErrorText }}',
                   trigger: ['input','blur'],
               },
               {{- if eq .FieldType "string" }}
               {
                   whitespace: true,
                   message: 'Cannot enter only spaces',
                   trigger: ['input', 'blur'],
              }
              {{- end }}
              ],
            {{- end }}
    {{- end }}
})

const searchRule = reactive({
  createdAt: [
    { validator: (rule, value, callback) => {
      if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
        callback(new Error('Please fill in the end date'))
      } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
        callback(new Error('Please fill in the start date'))
      } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() || searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
        callback(new Error('Start date should be earlier than end date'))
      } else {
        callback()
      }
    }, trigger: 'change' }
  ],
  {{- range .Fields }}
    {{- if .FieldSearchType}}
      {{- if eq .FieldType "time.Time" }}
        {{.FieldJson }} : [{ validator: (rule, value, callback) => {
        if (searchInfo.value.start{{.FieldName}} && !searchInfo.value.end{{.FieldName}}) {
          callback(new Error('Please fill in the end date'))
        } else if (!searchInfo.value.start{{.FieldName}} && searchInfo.value.end{{.FieldName}}) {
          callback(new Error('Please fill in the start date'))
        } else if (searchInfo.value.start{{.FieldName}} && searchInfo.value.end{{.FieldName}} && (searchInfo.value.start{{.FieldName}}.getTime() === searchInfo.value.end{{.FieldName}}.getTime() || searchInfo.value.start{{.FieldName}}.getTime() > searchInfo.value.end{{.FieldName}}.getTime())) {
          callback(new Error('Start date should be earlier than end date'))
        } else {
          callback()
        }
      }, trigger: 'change' }],
      {{- end }}
    {{- end }}
  {{- end }}
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== Table control part ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

{{- if .NeedSort}}
// sort
const sortChange = ({ prop, order }) => {
  searchInfo.value.sort = prop
  searchInfo.value.order = order
  getTableData()
}
{{- end}}

// reset
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// search
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    pageSize.value = 10
    {{- range .Fields}}{{- if eq .FieldType "bool" }}
    if (searchInfo.value.{{.FieldJson}} === ""){
        searchInfo.value.{{.FieldJson}}=null
    }{{ end }}{{ end }}
    getTableData()
  })
}

// paging
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

//Modify page capacity
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// Inquire
const getTableData = async() => {
  const table = await get{{.StructName}}List({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== End of table control part ===============

// Get the required dictionary. It may be empty. Keep it as needed.
const setOptions = async () =>{
{{- range $index, $element := .DictTypes }}
    {{ $element }}Options.value = await getDictFunc('{{$element}}')
{{- end }}
}

// Get the required dictionary. It may be empty. Keep it as needed.
setOptions()


//Multiple selection data
const multipleSelection = ref([])
//Multiple selection
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// delete row
const deleteRow = (row) => {
    ElMessageBox.confirm('Are you sure you want to delete?', 'Prompt', {
        confirmButtonText: 'OK',
        cancelButtonText: 'Cancel',
        type: 'warning'
    }).then(() => {
            delete{{.StructName}}Func(row)
        })
    }


// Batch delete control tags
const deleteVisible = ref(false)

//Multiple selection delete
const onDelete = async() => {
      const ids = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: 'Please select the data to be deleted'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          ids.push(item.ID)
        })
      const res = await delete{{.StructName}}ByIds({ ids })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: 'Delete successfully'
        })
        if (tableData.value.length === ids.length && page.value > 1) {
          page.value--
        }
        deleteVisible.value = false
        getTableData()
      }
    }

// Behavior control tag (does it need to be added or modified inside the pop-up window)
const type = ref('')

// update row
const update{{.StructName}}Func = async(row) => {
    const res = await find{{.StructName}}({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.re{{.Abbreviation}}
        dialogFormVisible.value = true
    }
}


// delete row
const delete{{.StructName}}Func = async (row) => {
    const res = await delete{{.StructName}}({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: 'Delete successfully'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// Pop-up window control tag
const dialogFormVisible = ref(false)


//View details control tag
const detailShow = ref(false)


//Open the details popup window
const openDetailShow = () => {
  detailShow.value = true
}


//Open details
const getDetails = async (row) => {
  //Open the pop-up window
  const res = await find{{.StructName}}({ ID: row.ID })
  if (res.code === 0) {
    formData.value = res.data.re{{.Abbreviation}}
    openDetailShow()
  }
}


// Close the details pop-up window
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
      {{- range .Fields}}
          {{- if eq .FieldType "bool" }}
          {{.FieldJson}}: false,
          {{- end }}
          {{- if eq .FieldType "string" }}
          {{.FieldJson}}: '',
          {{- end }}
          {{- if eq .FieldType "int" }}
          {{.FieldJson}}: {{- if .DictType }} undefined{{ else }} 0{{- end }},
          {{- end }}
          {{- if eq .FieldType "time.Time" }}
          {{.FieldJson}}: new Date(),
          {{- end }}
          {{- if eq .FieldType "float64" }}
          {{.FieldJson}}: 0,
          {{- end }}
          {{- end }}
          }
}


//Open the pop-up window
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

//Close pop-up window
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
    {{- range .Fields}}
        {{- if eq .FieldType "bool" }}
        {{.FieldJson}}: false,
        {{- end }}
        {{- if eq .FieldType "string" }}
        {{.FieldJson}}: '',
        {{- end }}
        {{- if eq .FieldType "int" }}
        {{.FieldJson}}: {{- if .DictType }} undefined{{ else }} 0{{- end }},
        {{- end }}
        {{- if eq .FieldType "time.Time" }}
        {{.FieldJson}}: new Date(),
        {{- end }}
        {{- if eq .FieldType "float64" }}
        {{.FieldJson}}: 0,
        {{- end }}
        {{- end }}
        }
}
// Pop-up window confirmed
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await create{{.StructName}}(formData.value)
                  break
                case 'update':
                  res = await update{{.StructName}}(formData.value)
                  break
                default:
                  res = await create{{.StructName}}(formData.value)
                  break
              }
              if (res.code === 0) {
                ElMessage({
                  type: 'success',
                  message: 'Creation/change successful'
                })
                closeDialog()
                getTableData()
              }
      })
}
{{if .HasFile }}
const downloadFile = (url) => {
    window.open(getUrl(url), '_blank')
}
{{end}}
</script>

<style>
{{if .HasFile }}
.file-list{
  display: flex;
  flex-wrap: wrap;
  gap: 4px;
}

.fileBtn{
  margin-bottom: 10px;
}

.fileBtn:last-child{
  margin-bottom: 0;
}
{{end}}
</style>
