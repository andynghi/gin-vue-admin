<template>
  <div>
    <warning-bar
      href="https://www.bilibili.com/video/BV1kv4y1g7nT?p=3"
      title="This function is for development environment use and is not recommended for release to production. Please see the video for specific usage effects https://www.bilibili.com/video/BV1kv4y1g7nT?p=3"
    />
    <!-- Get fields directly from the database -->
    <div class="gva-search-box">
      <el-collapse
        v-model="activeNames"
        style="margin-bottom:12px"
      >
        <el-collapse-item name="1">
          <template #title>
            <div :style="{fontSize:'16px',paddingLeft:'20px'}">
              Click here to create code from an existing database
              <el-icon class="header-icon ">
                <pointer />
              </el-icon>
            </div>
          </template>
          <el-form
            ref="getTableForm"
            style="margin-top:24px"
            :inline="true"
            :model="dbform"
            label-width="120px"
          >
            <el-form-item
              label="Business Library"
              prop="selectDBtype"
            >
              <template #label>
                <el-tooltip
                  content="Note: You need to go to db-list in advance to configure multiple databases by yourself. If not configured, you need to configure and restart the service before using it. (You can select the corresponding database table here, which can be understood as which database to select the table from)"
                  placement="bottom"
                  effect="light"
                >
                  <div> Business Library <el-icon><QuestionFilled /></el-icon> </div>
                </el-tooltip>
              </template>
              <el-select
                v-model="dbform.businessDB"
                clearable
                style="width:194px"
                placeholder="Select business library"
                @change="getDbFunc"
              >
                <el-option
                  v-for="item in dbList"
                  :key="item.aliasName"
                  :value="item.aliasName"
                  :label="item.aliasName"
                  :disabled="item.disable"
                >
                  <div>
                    <span>{{ item.aliasName }}</span>
                    <span style="float:right;color:#8492a6;font-size:13px">{{ item.dbName }}</span>
                  </div>
                </el-option>
              </el-select>
            </el-form-item>
            <el-form-item
              label="database name"
              prop="structName"
            >
              <el-select
                v-model="dbform.dbName"
                clearable
                filterable
                placeholder="Please select a database"
                @change="getTableFunc"
              >
                <el-option
                  v-for="item in dbOptions"
                  :key="item.database"
                  :label="item.database"
                  :value="item.database"
                />
              </el-select>
            </el-form-item>
            <el-form-item
              label="table name"
              prop="structName"
            >
              <el-select
                v-model="dbform.tableName"
                :disabled="!dbform.dbName"
                filterable
                placeholder="Please select a table"
              >
                <el-option
                  v-for="item in tableOptions"
                  :key="item.tableName"
                  :label="item.tableName"
                  :value="item.tableName"
                />
              </el-select>
            </el-form-item>
            <el-form-item>
              <el-button
                type="primary"
                @click="getColumnFunc"
              >Create using this table</el-button>
            </el-form-item>
          </el-form>
        </el-collapse-item>
      </el-collapse>
    </div>
    <div class="gva-search-box">
      <!-- Initial version of automated code tool -->
      <el-form
        ref="autoCodeForm"
        :rules="rules"
        :model="form"
        label-width="120px"
        :inline="true"
      >
        <el-form-item
          label="Struct name"
          prop="structName"
        >
          <el-input
            v-model="form.structName"
            placeholder="The first letter is automatically converted to uppercase"
          />
        </el-form-item>
        <el-form-item
          label="TableName"
          prop="tableName"
        >
          <el-input
            v-model="form.tableName"
            placeholder="Specify table name (not required)"
          />
        </el-form-item>
        <el-form-item
          prop="abbreviation"
        >
          <template #label>
            <el-tooltip
              content="The abbreviation will be used as the input parameter object name and routing group"
              placement="bottom"
              effect="light"
            >
              <div> Struct abbreviation <el-icon><QuestionFilled /></el-icon> </div>
            </el-tooltip>
          </template>
          <el-input
            v-model="form.abbreviation"
            placeholder="Please enter the Struct abbreviation"
          />
        </el-form-item>
        <el-form-item
          label="Struct Chinese name"
          prop="description"
        >
          <el-input
            v-model="form.description"
            placeholder="Chinese description as automatic api description"
          />
        </el-form-item>
        <el-form-item
          prop="packageName"
        >
          <template #label>
            <el-tooltip
              content="Default name of the generated file (recommended to be in camel case format, with the first letter lowercase, such as sysXxxXxxx)"
              placement="bottom"
              effect="light"
            >
              <div> File name <el-icon><QuestionFilled /></el-icon> </div>
            </el-tooltip>
          </template>
          <el-input
            v-model="form.packageName"
            placeholder="Please enter file name"
            @blur="toLowerCaseFunc(form,'packageName')"
          />
        </el-form-item>
        <el-form-item
          label="Package"
          prop="package"
        >
          <el-select
            v-model="form.package"
            style="width:194px"
          >
            <el-option
              v-for="item in pkgs"
              :key="item.ID"
              :value="item.packageName"
              :label="item.packageName"
            />
          </el-select>
          <el-icon
            class="cursor-pointer ml-2 text-gray-600"
            @click="getPkgs"
          ><refresh /></el-icon>
          <el-icon
            class="cursor-pointer ml-2 text-gray-600"
            @click="goPkgs"
          ><document-add /></el-icon>
        </el-form-item>
        <el-form-item
          label="Business Library"
          prop="businessDB"
        >
          <template #label>
            <el-tooltip
              content="Note: You need to go to db-list in advance to configure multiple databases by yourself. If this item is empty, the gva library will be used to create the automation code (global.GVA_DB). After filling in, the code of the specified library will be created (global.MustGetGlobalDBByDBName(dbname ))"
              placement="bottom"
              effect="light"
            >
              <div> Business Library <el-icon><QuestionFilled /></el-icon> </div>
            </el-tooltip>
          </template>
          <el-select
            v-model="form.businessDB"
            style="width:194px"
            placeholder="Select business library"
          >
            <el-option
              v-for="item in dbList"
              :key="item.aliasName"
              :value="item.aliasName"
              :label="item.aliasName"
              :disabled="item.disable"
            >
              <div>
                <span>{{ item.aliasName }}</span>
                <span style="float:right;color:#8492a6;font-size:13px">{{ item.dbName }}</span>
              </div>
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item>
          <template #label>
            <el-tooltip
              content="Note: created_by updated_by deleted_by will be automatically added to the structure to facilitate users to control resource permissions"
              placement="bottom"
              effect="light"
            >
              <div> Create resource identifier <el-icon><QuestionFilled /></el-icon> </div>
            </el-tooltip>
          </template>
          <el-checkbox v-model="form.autoCreateResource" />
        </el-form-item>
        <el-form-item>
          <template #label>
            <el-tooltip
              content="Note: Register the automatically generated API into the database"
              placement="bottom"
              effect="light"
            >
              <div> Automatically create API </div>
            </el-tooltip>
          </template>
          <el-checkbox v-model="form.autoCreateApiToSql" />
        </el-form-item>
        <el-form-item>
          <template #label>
            <el-tooltip
              content="Note: Automatically migrate the generated files to the corresponding location of the yaml configuration"
              placement="bottom"
              effect="light"
            >
              <div> Automatically move files </div>
            </el-tooltip>
          </template>
          <el-checkbox v-model="form.autoMoveFile" />
        </el-form-item>
      </el-form>
    </div>
    <!-- Component list -->
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button
          type="primary"
          @click="editAndAddField()"
        >Add new field</el-button>
      </div>
      <el-table :data="form.fields">
        <el-table-column
          align="left"
          type="index"
          label="sequence"
          width="60"
        />
        <el-table-column
          align="left"
          prop="fieldName"
          label="field name"
          width="160"
        >
          <template #default="{row}">
            <el-input v-model="row.fieldName" />
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          prop="fieldDesc"
          label="Chinese name"
          width="160"
        >
          <template #default="{row}">
            <el-input v-model="row.fieldDesc" />
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          prop="require"
          label="required"
        >
          <template #default="{row}"> <el-checkbox v-model="row.require" /></template>
        </el-table-column>
        <el-table-column
          align="left"
          prop="sort"
          label="Sort"
        >
          <template #default="{row}"> <el-checkbox v-model="row.sort" /> </template>
        </el-table-column>
        <el-table-column
          align="left"
          prop="fieldJson"
          width="160px"
          label="Field Json"
        >
          <template #default="{row}">
            <el-input v-model="row.fieldJson" />
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          prop="fieldType"
          label="field type"
          width="160"
        >
          <template #default="{row}">
            <el-select
              v-model="row.fieldType"
              style="width:100%"
              placeholder="Please select field type"
              clearable
            >
              <el-option
                v-for="item in typeOptions"
                :key="item.value"
                :label="item.label"
                :value="item.value"
              />
            </el-select>
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          prop="dataTypeLong"
          label="Database field length"
          width="160"
        >
          <template #default="{row}">
            <el-input v-model="row.dataTypeLong" />
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          prop="columnName"
          label="database field"
          width="160"
        >
          <template #default="{row}">
            <el-input v-model="row.columnName" />
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          prop="comment"
          label="Database field description"
          width="160"
        >
          <template #default="{row}">
            <el-input v-model="row.comment" />
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          prop="fieldSearchType"
          label="Search criteria"
          width="130"
        >
          <template #default="{row}">
            <el-select
              v-model="row.fieldSearchType"
              style="width:100%"
              placeholder="Please select field query conditions"
              clearable
            >
              <el-option
                v-for="item in typeSearchOptions"
                :key="item.value"
                :label="item.label"
                :value="item.value"
                :disabled="
                  (row.fieldType!=='string'&&item.value==='LIKE')||
                    ((row.fieldType!=='int'&&row.fieldType!=='time.Time'&&row.fieldType!=='float64')&&(item.value==='BETWEEN' || item.value==='NOT BETWEEN'))
                "
              />
            </el-select>
          </template>

        </el-table-column>
        <el-table-column
          align="left"
          label="operation"
          width="300"
          fixed="right"
        >
          <template #default="scope">
            <el-button

              type="primary"
              link
              icon="edit"
              @click="editAndAddField(scope.row)"
            >Advanced Editor</el-button>
            <el-button

              type="primary"
              link
              :disabled="scope.$index === 0"
              @click="moveUpField(scope.$index)"
            >Move up</el-button>
            <el-button

              type="primary"
              link
              :disabled="(scope.$index + 1) === form.fields.length"
              @click="moveDownField(scope.$index)"
            >Move down</el-button>
            <el-popover
              v-model="scope.row.visible"
              placement="top"
            >
              <p>Are you sure to delete? </p>
              <div style="text-align: right; margin-top: 8px;">
                <el-button
                  type="primary"
                  link
                  @click="scope.row.visible = false"
                >Cancel</el-button>
                <el-button
                  type="primary"
                  @click="deleteField(scope.$index)"
                >OK</el-button>
              </div>
              <template #reference>
                <el-button
                  type="primary"
                  link
                  icon="delete"
                  @click="scope.row.visible = true"
                >Delete</el-button>
              </template>
            </el-popover>
          </template>
        </el-table-column>
      </el-table>
      <!-- Component list -->
      <div class="gva-btn-list justify-end mt-4">
        <el-button
          type="primary"
          @click="enterForm(true)"
        >Preview code</el-button>
        <el-button
          type="primary"
          @click="enterForm(false)"
        >Generate code</el-button>
      </div>
    </div>
    <!-- Component pop-up window -->
    <el-dialog
      v-model="dialogFlag"
      width="70%"
      title="Component content"
    >
      <FieldDialog
        v-if="dialogFlag"
        ref="fieldDialogNode"
        :dialog-middle="dialogMiddle"
        :type-options="typeOptions"
        :type-search-options="typeSearchOptions"
      />
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">Cancel</el-button>
          <el-button
            type="primary"
            @click="enterDialog"
          >Confirm</el-button>
        </div>
      </template>
    </el-dialog>

    <el-dialog v-model="previewFlag">
      <template #header>
        <div class="flex items-center py-1.5">
          <p>Action bar:</p>
          <el-button
            type="primary"
            @click="selectText"
          >Select all</el-button>
          <el-button
            type="primary"
            @click="copy"
          >Copy</el-button>
        </div>
      </template>
      <PreviewCodeDialog
        v-if="previewFlag"
        ref="previewNode"
        :preview-code="preViewCode"
      />
      <template #footer>
        <div
          class="dialog-footer"
          style="padding-top:14px;padding-right:14px"
        >
          <el-button
            type="primary"
            @click="previewFlag = false"
          >Confirm</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>

import FieldDialog from '@/view/systemTools/autoCode/component/fieldDialog.vue'
import PreviewCodeDialog from '@/view/systemTools/autoCode/component/previewCodeDialg.vue'
import { toUpperCase, toHump, toSQLLine, toLowerCase } from '@/utils/stringFun'
import { createTemp, getDB, getTable, getColumn, preview, getMeta, getPackageApi } from '@/api/autoCode'
import { getDict } from '@/utils/dictionary'
import { ref, getCurrentInstance, reactive, watch, toRaw } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import WarningBar from '@/components/warningBar/warningBar.vue'

defineOptions({
  name: 'AutoCode'
})

const typeOptions = ref([
  {
    label: 'string',
    value: 'string'
  },
  {
    label: 'rich text',
    value: 'richtext'
  },
  {
    label: 'integer',
    value: 'int'
  },
  {
    label: 'boolean',
    value: 'bool'
  },
  {
    label: 'floating point',
    value: 'float64'
  },
  {
    label: 'time',
    value: 'time.Time'
  },
  {
    label: 'enumeration',
    value: 'enum'
  },
  {
    label: 'Single picture (string)',
    value: 'picture',
  },
  {
    label: 'Multiple pictures (json string)',
    value: 'pictures',
  },
  {
    label: 'Video (string)',
    value: 'video',
  },
  {
    label: 'File (json string)',
    value: 'file',
  }
])

const typeSearchOptions = ref([
  {
    label: '=',
    value: '='
  },
  {
    label: '<>',
    value: '<>'
  },
  {
    label: '>',
    value: '>'
  },
  {
    label: '<',
    value: '<'
  },
  {
    label: 'LIKE',
    value: 'LIKE'
  },
  {
    label: 'BETWEEN',
    value: 'BETWEEN'
  },
  {
    label: 'NOT BETWEEN',
    value: 'NOT BETWEEN'
  }
])

const fieldTemplate = {
  fieldName: '',
  fieldDesc: '',
  fieldType: '',
  dataType: '',
  fieldJson: '',
  columnName: '',
  dataTypeLong: '',
  comment: '',
  require: false,
  sort: false,
  errorText: '',
  clearable: true,
  fieldSearchType: '',
  dictType: ''
}
const route = useRoute()
const router = useRouter()
const activeNames = reactive([])
const preViewCode = ref({})
const dbform = ref({
  businessDB: '',
  dbName: '',
  tableName: ''
})
const tableOptions = ref([])
const addFlag = ref('')
const fdMap = ref({})
const form = ref({
  structName: '',
  tableName: '',
  packageName: '',
  package: '',
  abbreviation: '',
  description: '',
  businessDB: '',
  autoCreateApiToSql: true,
  autoMoveFile: true,
  autoCreateResource: false,
  fields: []
})
const rules = ref({
  structName: [
    { required: true, message: 'Please enter the structure name', trigger: 'blur' }
  ],
  abbreviation: [
    { required: true, message: 'Please enter the structure abbreviation', trigger: 'blur' }
  ],
  description: [
    { required: true, message: 'Please enter a structure description', trigger: 'blur' }
  ],
  packageName: [
    {
      required: true,
      message: 'File name: sysXxxxXxxx',
      trigger: 'blur'
    }
  ],
  package: [
    { required: true, message: 'Please select package', trigger: 'blur' }
  ]
})
const dialogMiddle = ref({})
const bk = ref({})
const dialogFlag = ref(false)
const previewFlag = ref(false)

const toLowerCaseFunc = (form, key) => {
  form[key] = toLowerCase(form[key])
}
const previewNode = ref(null)
const selectText = () => {
  previewNode.value.selectText()
}
const copy = () => {
  previewNode.value.copy()
}
const editAndAddField = (item) => {
  dialogFlag.value = true
  if (item) {
    addFlag.value = 'edit'
    bk.value = JSON.parse(JSON.stringify(item))
    dialogMiddle.value = item
  } else {
    addFlag.value = 'add'
    dialogMiddle.value = JSON.parse(JSON.stringify(fieldTemplate))
  }
}
const moveUpField = (index) => {
  if (index === 0) {
    return
  }
  const oldUpField = form.value.fields[index - 1]
  form.value.fields.splice(index - 1, 1)
  form.value.fields.splice(index, 0, oldUpField)
}
const moveDownField = (index) => {
  const fCount = form.value.fields.length
  if (index === fCount - 1) {
    return
  }
  const oldDownField = form.value.fields[index + 1]
  form.value.fields.splice(index + 1, 1)
  form.value.fields.splice(index, 0, oldDownField)
}

const currentInstance = getCurrentInstance()
const enterDialog = () => {
  currentInstance.refs.fieldDialogNode.fieldDialogFrom.validate(valid => {
    if (valid) {
      dialogMiddle.value.fieldName = toUpperCase(
        dialogMiddle.value.fieldName
      )
      if (addFlag.value === 'add') {
        form.value.fields.push(dialogMiddle.value)
      }
      dialogFlag.value = false
    } else {
      return false
    }
  })
}
const closeDialog = () => {
  if (addFlag.value === 'edit') {
    dialogMiddle.value = bk.value
  }
  dialogFlag.value = false
}
const deleteField = (index) => {
  form.value.fields.splice(index, 1)
}
const autoCodeForm = ref(null)
const enterForm = async(isPreview) => {
  if (form.value.fields.length <= 0) {
    ElMessage({
      type: 'error',
      message: 'Please fill in at least one field'
    })
    return false
  }
  if (
    form.value.fields.some(item => item.fieldName === form.value.structName)
  ) {
    ElMessage({
      type: 'error',
      message: 'There is a field with the same name as the structure'
    })
    return false
  }

  if (form.value.package === form.value.abbreviation) {
    ElMessage({
      type: 'error',
      message: 'Package and structure abbreviations cannot have the same name'
    })
    return false
  }

  autoCodeForm.value.validate(async valid => {
    if (valid) {
      for (const key in form.value) {
        if (typeof form.value[key] === 'string') {
          form.value[key] = form.value[key].trim()
        }
      }
      form.value.structName = toUpperCase(form.value.structName)
      form.value.tableName = form.value.tableName.replace(' ', '')
      if (!form.value.tableName) {
        form.value.tableName = toSQLLine(toLowerCase(form.value.structName))
      }
      if (form.value.structName === form.value.abbreviation) {
        ElMessage({
          type: 'error',
          message: 'structName and struct abbreviation cannot be the same'
        })
        return false
      }
      form.value.humpPackageName = toSQLLine(form.value.packageName)
      if (isPreview) {
        const data = await preview(form.value)
        preViewCode.value = data.data.autoCode
        previewFlag.value = true
      } else {
        const data = await createTemp(form.value)
        if (data.headers?.success === 'false') {
          return
        }
        if (form.value.autoMoveFile) {
          ElMessage({
            type: 'success',
            message: 'Automation code created successfully, automatic move successful'
          })
          return
        }
        ElMessage({
          type: 'success',
          message: 'Automation code created successfully and downloading'
        })
        const blob = new Blob([data])
        const fileName = 'ginvueadmin.zip'
        if ('download' in document.createElement('a')) {
          // Not IE browser
          const url = window.URL.createObjectURL(blob)
          const link = document.createElement('a')
          link.style.display = 'none'
          link.href = url
          link.setAttribute('download', fileName)
          document.body.appendChild(link)
          link.click()
          document.body.removeChild(link) // Remove element after download is complete
          window.URL.revokeObjectURL(url) // Release the blob object
        } else {
          // IE 10+
          window.navigator.msSaveBlob(blob, fileName)
        }
      }
    } else {
      return false
    }
  })
}

const dbList = ref([])
const dbOptions = ref([])

const getDbFunc = async() => {
  dbform.value.dbName = ''
  dbform.value.tableName = ''
  const res = await getDB({ businessDB: dbform.value.businessDB })
  if (res.code === 0) {
    dbOptions.value = res.data.dbs
    dbList.value = res.data.dbList
  }
}
const getTableFunc = async() => {
  const res = await getTable({ businessDB: dbform.value.businessDB, dbName: dbform.value.dbName })
  if (res.code === 0) {
    tableOptions.value = res.data.tables
  }
  dbform.value.tableName = ''
}

const getColumnFunc = async() => {
  const gormModelList = ['id', 'created_at', 'updated_at', 'deleted_at']
  const res = await getColumn(dbform.value)
  if (res.code === 0) {
    let dbtype = ''
    if (dbform.value.businessDB !== '') {
      const dbtmp = dbList.value.find(item => item.aliasName === dbform.value.businessDB)
      const dbraw = toRaw(dbtmp)
      dbtype = dbraw.dbtype
    }
    const tbHump = toHump(dbform.value.tableName)
    form.value.structName = toUpperCase(tbHump)
    form.value.tableName = dbform.value.tableName
    form.value.packageName = tbHump
    form.value.abbreviation = tbHump
    form.value.description = tbHump + 'table'
    form.value.autoCreateApiToSql = true
    form.value.autoMoveFile = true
    form.value.fields = []
    res.data.columns &&
          res.data.columns.forEach(item => {
            if (!gormModelList.some(gormfd => gormfd === item.columnName)) {
              const fbHump = toHump(item.columnName)
              form.value.fields.push({
                fieldName: toUpperCase(fbHump),
                fieldDesc: item.columnComment || fbHump + 'field',
                fieldType: fdMap.value[item.dataType],
                dataType: item.dataType,
                fieldJson: fbHump,
                dataTypeLong: item.dataTypeLong && item.dataTypeLong.split(',')[0],
                columnName: dbtype === 'oracle' ? item.columnName.toUpperCase() : item.columnName,
                comment: item.columnComment,
                require: false,
                errorText: '',
                clearable: true,
                fieldSearchType: '',
                dictType: ''
              })
            }
          })
  }
}
const setFdMap = async() => {
  const fdTypes = ['string', 'int', 'bool', 'float64', 'time.Time']
  fdTypes.forEach(async fdtype => {
    const res = await getDict(fdtype)
    res && res.forEach(item => {
      fdMap.value[item.label] = fdtype
    })
  })
}
const getAutoCodeJson = async(id) => {
  const res = await getMeta({ id: Number(id) })
  if (res.code === 0) {
    form.value = JSON.parse(res.data.meta)
  }
}

const pkgs = ref([])
const getPkgs = async() => {
  const res = await getPackageApi()
  if (res.code === 0) {
    pkgs.value = res.data.pkgs
  }
}

const goPkgs = () => {
  router.push({ name: 'autoPkg' })
}

const init = () => {
  getDbFunc()
  setFdMap()
  getPkgs()
  const id = route.params.id
  if (id) {
    getAutoCodeJson(id)
  }
}
init()

watch(() => route.params.id, () => {
  if (route.name === 'autoCodeEdit') {
    init()
  }
})

</script>
