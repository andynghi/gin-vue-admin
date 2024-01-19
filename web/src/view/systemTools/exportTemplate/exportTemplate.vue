<template>
  <div>
    <WarningBar
      title="This function provides synchronous table export function and asynchronous table export function for large amounts of data. You can choose to click me to customize"
      href="https://flipped-aurora.feishu.cn/docx/KwjxdnvatozgwIxGV0rcpkZSn4d"
    />
    <div class="gva-search-box">
      <el-form
        ref="elSearchFormRef"
        :inline="true"
        :model="searchInfo"
        class="demo-form-inline"
        :rules="searchRule"
        @keyup.enter="onSubmit"
      >
        <el-form-item
          label="Creation date"
          prop="createdAt"
        >
          <template #label>
            <span>
              Creation date
              <el-tooltip content="The search range is from start date (inclusive) to end date (exclusive)">
                <el-icon><QuestionFilled /></el-icon>
              </el-tooltip>
            </span>
          </template>
          <el-date-picker
            v-model="searchInfo.startCreatedAt"
            type="datetime"
            placeholder="start date"
            :disabled-date="time=> searchInfo.endCreatedAt ? time.getTime() > searchInfo.endCreatedAt.getTime() : false"
          />
          —
          <el-date-picker
            v-model="searchInfo.endCreatedAt"
            type="datetime"
            placeholder="end date"
            :disabled-date="time=> searchInfo.startCreatedAt ? time.getTime() < searchInfo.startCreatedAt.getTime() : false"
          />
        </el-form-item>
        <el-form-item
          label="template name"
          prop="name"
        >
          <el-input
            v-model="searchInfo.name"
            placeholder="Search criteria"
          />

        </el-form-item>
        <el-form-item
          label="table name"
          prop="tableName"
        >
          <el-input
            v-model="searchInfo.tableName"
            placeholder="Search criteria"
          />

        </el-form-item>
        <el-form-item
          label="template logo"
          prop="templateID"
        >
          <el-input
            v-model="searchInfo.templateID"
            placeholder="Search criteria"
          />

        </el-form-item>
        <el-form-item>
          <el-button
            type="primary"
            icon="search"
            @click="onSubmit"
          >Query</el-button>
          <el-button
            icon="refresh"
            @click="onReset"
          >Reset</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button
          type="primary"
          icon="plus"
          @click="openDialog"
        >Add</el-button>
        <el-popover
          v-model:visible="deleteVisible"
          :disabled="!multipleSelection.length"
          placement="top"
          width="160"
        >
          <p>Are you sure you want to delete it? </p>
          <div style="text-align: right; margin-top: 8px;">
            <el-button
              type="primary"
              link
              @click="deleteVisible = false"
            >Cancel</el-button>
            <el-button
              type="primary"
              @click="onDelete"
            >OK</el-button>
          </div>
          <template #reference>
            <el-button
              icon="delete"
              style="margin-left: 10px;"
              :disabled="!multipleSelection.length"
              @click="deleteVisible = true"
            >Delete</el-button>
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
      >
        <el-table-column
          type="selection"
          width="55"
        />
        <el-table-column
          align="left"
          label="date"
          width="180"
        >
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column
          align="left"
          label="template name"
          prop="name"
          width="120"
        />
        <el-table-column
          align="left"
          label="table name"
          prop="tableName"
          width="120"
        />
        <el-table-column
          align="left"
          label="template logo"
          prop="templateID"
          width="120"
        />
        <el-table-column
          align="left"
          label="template information"
          prop="templateInfo"
          min-width="120"
        />
        <el-table-column
          align="left"
          label="operation"
          min-width="120"
        >
          <template #default="scope">
            <el-button
              type="primary"
              link
              class="table-button"
              @click="getDetails(scope.row)"
            >
              <el-icon style="margin-right: 5px"><InfoFilled /></el-icon>
              check the details
            </el-button>
            <el-button
              type="primary"
              link
              icon="edit"
              class="table-button"
              @click="updateSysExportTemplateFunc(scope.row)"
            >Change</el-button>
            <el-button
              type="primary"
              link
              icon="delete"
              @click="deleteRow(scope.row)"
            >Delete</el-button>
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
    <el-dialog
      v-model="dialogFormVisible"
      :before-close="closeDialog"
      :title="type==='create'?'Add':'Modify'"
      destroy-on-close
    >
      <el-scrollbar height="500px">
        <el-form
          ref="elFormRef"
          :model="formData"
          label-position="right"
          :rules="rule"
          label-width="100px"
        >
          <el-form-item
            label="template name:"
            prop="name"
          >
            <el-input
              v-model="formData.name"
              :clearable="true"
              placeholder="Please enter the template name"
            />
          </el-form-item>
          <el-form-item
            label="Table name:"
            prop="tableName"
          >
            <el-input
              v-model="formData.tableName"
              :clearable="true"
              placeholder="Please enter the name of the table to be exported"
            />
          </el-form-item>
          <el-form-item
            label="Template ID:"
            prop="templateID"
          >
            <el-input
              v-model="formData.templateID"
              :clearable="true"
              placeholder="The template logo is the logo attribute that the front-end component needs to hang on"
            />
          </el-form-item>
          <el-form-item
            label="Template information:"
            prop="templateInfo"
          >
            <el-input
              v-model="formData.templateInfo"
              type="textarea"
              :rows="16"
              :clearable="true"
              :placeholder="templatePlaceholder"
            />
          </el-form-item>
        </el-form>
      </el-scrollbar>
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

    <el-dialog
      v-model="detailShow"
      style="width: 800px"
      lock-scroll
      :before-close="closeDetailShow"
      title="View details"
      destroy-on-close
    >
      <el-scrollbar height="550px">
        <el-descriptions
          column="1"
          border
        >
          <el-descriptions-item label="template name">
            {{ formData.name }}
          </el-descriptions-item>
          <el-descriptions-item label="table name">
            {{ formData.tableName }}
          </el-descriptions-item>
          <el-descriptions-item label="template label">
            {{ formData.templateID }}
          </el-descriptions-item>
          <el-descriptions-item label="Template information">
            {{ formData.templateInfo }}
          </el-descriptions-item>
        </el-descriptions>
      </el-scrollbar>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  createSysExportTemplate,
  deleteSysExportTemplate,
  deleteSysExportTemplateByIds,
  updateSysExportTemplate,
  findSysExportTemplate,
  getSysExportTemplateList
} from '@/api/exportTemplate.js'

// Fully introduced formatting tools, please keep them as needed
import { formatDate } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import WarningBar from '@/components/warningBar/warningBar.vue'

defineOptions({
  name: 'ExportTemplate'
})

const templatePlaceholder = `Template information format: key identifies the database column name, value identifies the exported excel column name, as follows:
{
  "table_name1":"First column",
  "table_name2":"Second column",
  "table_name3":"Third column",
  "table_name4":"Fourth column",
}`

//Automatically generated dictionary (may be empty) and fields
const formData = ref({
  name: '',
  tableName: '',
  templateID: '',
  templateInfo: '',
})

// Validation rules
const rule = reactive({
  name: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  {
    whitespace: true,
    message: 'Cannot enter only spaces',
    trigger: ['input', 'blur'],
  }
  ],
  tableName: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  {
    whitespace: true,
    message: 'Cannot enter only spaces',
    trigger: ['input', 'blur'],
  }
  ],
  templateID: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  {
    whitespace: true,
    message: 'Cannot enter only spaces',
    trigger: ['input', 'blur'],
  }
  ],
  templateInfo: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  {
    whitespace: true,
    message: 'Cannot enter only spaces',
    trigger: ['input', 'blur'],
  }
  ],
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
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== Table control part ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

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
  const table = await getSysExportTemplateList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
const setOptions = async() => {
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
    deleteSysExportTemplateFunc(row)
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
  const res = await deleteSysExportTemplateByIds({ ids })
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
const updateSysExportTemplateFunc = async(row) => {
  const res = await findSysExportTemplate({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.resysExportTemplate
    dialogFormVisible.value = true
  }
}

// delete row
const deleteSysExportTemplateFunc = async(row) => {
  const res = await deleteSysExportTemplate({ ID: row.ID })
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
const getDetails = async(row) => {
  //Open the pop-up window
  const res = await findSysExportTemplate({ ID: row.ID })
  if (res.code === 0) {
    formData.value = res.data.resysExportTemplate
    openDetailShow()
  }
}

// Close the details pop-up window
const closeDetailShow = () => {
  detailShow.value = false
  formData.value = {
    name: '',
    tableName: '',
    templateID: '',
    templateInfo: '',
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
    name: '',
    tableName: '',
    templateID: '',
    templateInfo: '',
  }
}
// Pop-up window confirmed
const enterDialog = async() => {
  // Determine whether formData.templateInfo is in standard json format. If it is not standard json, assist in adjustment.
  try {
    JSON.parse(formData.value.templateInfo)
  } catch (error) {
    ElMessage({
      type: 'error',
      message: 'The template information format is incorrect, please check'
    })
    return
  }

  elFormRef.value?.validate(async(valid) => {
    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createSysExportTemplate(formData.value)
        break
      case 'update':
        res = await updateSysExportTemplate(formData.value)
        break
      default:
        res = await createSysExportTemplate(formData.value)
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

</script>

<style>

</style>
