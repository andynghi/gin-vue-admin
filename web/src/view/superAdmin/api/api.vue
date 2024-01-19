<template>
  <div>
    <div class="gva-search-box">
      <el-form
        ref="searchForm"
        :inline="true"
        :model="searchInfo"
      >
        <el-form-item label="path">
          <el-input
            v-model="searchInfo.path"
            placeholder="path"
          />
        </el-form-item>
        <el-form-item label="description">
          <el-input
            v-model="searchInfo.description"
            placeholder="description"
          />
        </el-form-item>
        <el-form-item label="API group">
          <el-input
            v-model="searchInfo.apiGroup"
            placeholder="api group"
          />
        </el-form-item>
        <el-form-item label="Request">
          <el-select
            v-model="searchInfo.method"
            clearable
            placeholder="Please select"
          >
            <el-option
              v-for="item in methodOptions"
              :key="item.value"
              :label="`${item.label}(${item.value})`"
              :value="item.value"
            />
          </el-select>
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
          @click="openDialog('addApi')"
        >Add</el-button>
        <el-icon
          class="cursor-pointer"
          @click="toDoc('https://www.bilibili.com/video/BV1kv4y1g7nT?p=7&vd_source=f2640257c21e3b547a790461ed94875e')"
        ><VideoCameraFilled /></el-icon>
        <el-popover
          v-model="deleteVisible"
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
              :disabled="!apis.length"
              @click="deleteVisible = true"
            >Delete</el-button>
          </template>
        </el-popover>
        <el-popover
          v-model="freshVisible"
          placement="top"
          width="160"
        >
          <p>Are you sure you want to refresh the Casbin cache? </p>
          <div style="text-align: right; margin-top: 8px;">
            <el-button
              type="primary"
              link
              @click="freshVisible = false"
            >Cancel</el-button>
            <el-button
              type="primary"
              @click="onFresh"
            >OK</el-button>
          </div>
          <template #reference>
            <el-button
              icon="Refresh"
              @click="freshVisible = true"
            >Refresh cache</el-button>
          </template>
        </el-popover>
      </div>
      <el-table
        :data="tableData"
        @sort-change="sortChange"
        @selection-change="handleSelectionChange"
      >
        <el-table-column
          type="selection"
          width="55"
        />
        <el-table-column
          align="left"
          label="id"
          min-width="60"
          prop="ID"
          sortable="custom"
        />
        <el-table-column
          align="left"
          label="API path"
          min-width="150"
          prop="path"
          sortable="custom"
        />
        <el-table-column
          align="left"
          label="API group"
          min-width="150"
          prop="apiGroup"
          sortable="custom"
        />
        <el-table-column
          align="left"
          label="API introduction"
          min-width="150"
          prop="description"
          sortable="custom"
        />
        <el-table-column
          align="left"
          label="Request"
          min-width="150"
          prop="method"
          sortable="custom"
        >
          <template #default="scope">
            <div>
              {{ scope.row.method }} / {{ methodFilter(scope.row.method) }}
            </div>
          </template>
        </el-table-column>

        <el-table-column
          align="left"
          fixed="right"
          label="operation"
          width="200"
        >
          <template #default="scope">
            <el-button
              icon="edit"

              type="primary"
              link
              @click="editApiFunc(scope.row)"
            >Edit</el-button>
            <el-button
              icon="delete"

              type="primary"
              link
              @click="deleteApiFunc(scope.row)"
            >Delete</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          layout="total, sizes, prev, pager, next, jumper"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>

    </div>

    <el-dialog
      v-model="dialogFormVisible"
      :before-close="closeDialog"
      :title="dialogTitle"
    >
      <warning-bar title="New API, you need to configure permissions in role management before you can use it" />
      <el-form
        ref="apiForm"
        :model="form"
        :rules="rules"
        label-width="80px"
      >
        <el-form-item
          label="path"
          prop="path"
        >
          <el-input
            v-model="form.path"
            autocomplete="off"
          />
        </el-form-item>
        <el-form-item
          label="Request"
          prop="method"
        >
          <el-select
            v-model="form.method"
            placeholder="Please select"
            style="width:100%"
          >
            <el-option
              v-for="item in methodOptions"
              :key="item.value"
              :label="`${item.label}(${item.value})`"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item
          label="api group"
          prop="apiGroup"
        >
          <el-input
            v-model="form.apiGroup"
            autocomplete="off"
          />
        </el-form-item>
        <el-form-item
          label="api introduction"
          prop="description"
        >
          <el-input
            v-model="form.description"
            autocomplete="off"
          />
        </el-form-item>
      </el-form>
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
  </div>
</template>

<script setup>
import {
  getApiById,
  getApiList,
  createApi,
  updateApi,
  deleteApi,
  deleteApisByIds,
  freshCasbin
} from '@/api/api'
import { toSQLLine } from '@/utils/stringFun'
import WarningBar from '@/components/warningBar/warningBar.vue'
import { ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { VideoCameraFilled } from '@element-plus/icons-vue'
import { toDoc } from '@/utils/doc'

defineOptions({
  name: 'Api',
})

const methodFilter = (value) => {
  const target = methodOptions.value.filter(item => item.value === value)[0]
  return target && `${target.label}`
}

const apis = ref([])
const form = ref({
  path: '',
  apiGroup: '',
  method: '',
  description: ''
})
const methodOptions = ref([
  {
    value: 'POST',
    label: 'Create',
    type: 'success'
  },
  {
    value: 'GET',
    label: 'View',
    type: ''
  },
  {
    value: 'PUT',
    label: 'update',
    type: 'warning'
  },
  {
    value: 'DELETE',
    label: 'delete',
    type: 'danger'
  }
])

const type = ref('')
const rules = ref({
  path: [{ required: true, message: 'Please enter the api path', trigger: 'blur' }],
  apiGroup: [
    { required: true, message: 'Please enter the group name', trigger: 'blur' }
  ],
  method: [
    { required: true, message: 'Please select the request method', trigger: 'blur' }
  ],
  description: [
    { required: true, message: 'Please enter api introduction', trigger: 'blur' }
  ]
})

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

const onReset = () => {
  searchInfo.value = {}
}
// search

const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  getTableData()
}

// paging
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// sort
const sortChange = ({ prop, order }) => {
  if (prop) {
    if (prop === 'ID') {
      prop = 'id'
    }
    searchInfo.value.orderKey = toSQLLine(prop)
    searchInfo.value.desc = order === 'descending'
  }
  getTableData()
}

// Inquire
const getTableData = async() => {
  const table = await getApiList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// batch operation
const handleSelectionChange = (val) => {
  apis.value = val
}

const deleteVisible = ref(false)
const onDelete = async() => {
  const ids = apis.value.map(item => item.ID)
  const res = await deleteApisByIds({ ids })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: res.msg
    })
    if (tableData.value.length === ids.length && page.value > 1) {
      page.value--
    }
    deleteVisible.value = false
    getTableData()
  }
}
const freshVisible = ref(false)
const onFresh = async() => {
  const res = await freshCasbin()
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: res.msg
    })
  }
  freshVisible.value = false
}

// Pop-up window related
const apiForm = ref(null)
const initForm = () => {
  apiForm.value.resetFields()
  form.value = {
    path: '',
    apiGroup: '',
    method: '',
    description: ''
  }
}

const dialogTitle = ref('New Api')
const dialogFormVisible = ref(false)
const openDialog = (key) => {
  switch (key) {
    case 'addApi':
      dialogTitle.value = 'New Api'
      break
    case 'edit':
      dialogTitle.value = 'EditApi'
      break
    default:
      break
  }
  type.value = key
  dialogFormVisible.value = true
}
const closeDialog = () => {
  initForm()
  dialogFormVisible.value = false
}

const editApiFunc = async(row) => {
  const res = await getApiById({ id: row.ID })
  form.value = res.data.api
  openDialog('edit')
}

const enterDialog = async() => {
  apiForm.value.validate(async valid => {
    if (valid) {
      switch (type.value) {
        case 'addApi':
          {
            const res = await createApi(form.value)
            if (res.code === 0) {
              ElMessage({
                type: 'success',
                message: 'Added successfully',
                showClose: true
              })
            }
            getTableData()
            closeDialog()
          }

          break
        case 'edit':
          {
            const res = await updateApi(form.value)
            if (res.code === 0) {
              ElMessage({
                type: 'success',
                message: 'Edit successfully',
                showClose: true
              })
            }
            getTableData()
            closeDialog()
          }
          break
        default:
          // eslint-disable-next-line no-lone-blocks
          {
            ElMessage({
              type: 'error',
              message: 'Unknown operation',
              showClose: true
            })
          }
          break
      }
    }
  })
}

const deleteApiFunc = async(row) => {
  ElMessageBox.confirm('This operation will permanently delete this API under all roles, do you want to continue?', 'Prompt', {
    confirmButtonText: 'OK',
    cancelButtonText: 'Cancel',
    type: 'warning'
  })
    .then(async() => {
      const res = await deleteApi(row)
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: 'Delete successfully!'
        })
        if (tableData.value.length === 1 && page.value > 1) {
          page.value--
        }
        getTableData()
      }
    })
}

</script>

<style scoped lang="scss">
.warning {
  color: #dc143c;
}
</style>
