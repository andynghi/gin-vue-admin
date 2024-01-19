<template>
  <div>
    <div class="gva-search-box">
      <el-form
        :inline="true"
        :model="searchInfo"
      >
        <el-form-item label="Request method">
          <el-input
            v-model="searchInfo.method"
            placeholder="Search criteria"
          />
        </el-form-item>
        <el-form-item label="Request path">
          <el-input
            v-model="searchInfo.path"
            placeholder="Search criteria"
          />
        </el-form-item>
        <el-form-item label="Result status code">
          <el-input
            v-model="searchInfo.status"
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
              style="margin-left: 10px;"
              :disabled="!multipleSelection.length"
              @click="deleteVisible = true"
            >Delete</el-button>
          </template>
        </el-popover>
      </div>
      <el-table
        ref="multipleTable"
        :data="tableData"
        style="width: 100%"
        tooltip-effect="dark"
        row-key="ID"
        @selection-change="handleSelectionChange"
      >
        <el-table-column
          align="left"
          type="selection"
          width="55"
        />
        <el-table-column
          align="left"
          label="operator"
          width="140"
        >
          <template #default="scope">
            <div>{{ scope.row.user.userName }}({{ scope.row.user.nickName }})</div>
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="date"
          width="180"
        >
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column
          align="left"
          label="status code"
          prop="status"
          width="120"
        >
          <template #default="scope">
            <div>
              <el-tag type="success">{{ scope.row.status }}</el-tag>
            </div>
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="Request IP"
          prop="ip"
          width="120"
        />
        <el-table-column
          align="left"
          label="Request method"
          prop="method"
          width="120"
        />
        <el-table-column
          align="left"
          label="request path"
          prop="path"
          width="240"
        />
        <el-table-column
          align="left"
          label="Request"
          prop="path"
          width="80"
        >
          <template #default="scope">
            <div>
              <el-popover
                v-if="scope.row.body"
                placement="left-start"
                trigger="click"
              >
                <div class="popover-box">
                  <pre>{{ fmtBody(scope.row.body) }}</pre>
                </div>
                <template #reference>
                  <el-icon style="cursor: pointer;"><warning /></el-icon>
                </template>
              </el-popover>

              <span v-else>None</span>
            </div>
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="response"
          prop="path"
          width="80"
        >
          <template #default="scope">
            <div>
              <el-popover
                v-if="scope.row.resp"
                placement="left-start"
                trigger="click"
              >
                <div class="popover-box">
                  <pre>{{ fmtBody(scope.row.resp) }}</pre>
                </div>
                <template #reference>
                  <el-icon style="cursor: pointer;"><warning /></el-icon>
                </template>
              </el-popover>
              <span v-else>None</span>
            </div>
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="operation"
        >
          <template #default="scope">
            <el-popover
              v-model="scope.row.visible"
              placement="top"
              width="160"
            >
              <p>Are you sure you want to delete it? </p>
              <div style="text-align: right; margin-top: 8px;">
                <el-button
                  type="primary"
                  link
                  @click="scope.row.visible = false"
                >Cancel</el-button>
                <el-button
                  type="primary"
                  @click="deleteSysOperationRecordFunc(scope.row)"
                >OK</el-button>
              </div>
              <template #reference>
                <el-button
                  icon="delete"
                  type="primary"
                  link
                  @click="scope.row.visible = true"
                >Delete</el-button>
              </template>
            </el-popover>
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
  </div>
</template>

<script setup>
import {
  deleteSysOperationRecord,
  getSysOperationRecordList,
  deleteSysOperationRecordByIds
} from '@/api/sysOperationRecord' // Please replace the address here.
import { formatDate } from '@/utils/format'
import { ref } from 'vue'
import { ElMessage } from 'element-plus'

defineOptions({
  name: 'SysOperationRecord'
})

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
const onReset = () => {
  searchInfo.value = {}
}
// See this method for conditional search front-end
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  if (searchInfo.value.status === '') {
    searchInfo.value.status = null
  }
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

// Inquire
const getTableData = async() => {
  const table = await getSysOperationRecordList({
    page: page.value,
    pageSize: pageSize.value,
    ...searchInfo.value,
  })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

const deleteVisible = ref(false)
const multipleSelection = ref([])
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}
const onDelete = async() => {
  const ids = []
  multipleSelection.value &&
        multipleSelection.value.forEach(item => {
          ids.push(item.ID)
        })
  const res = await deleteSysOperationRecordByIds({ ids })
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
const deleteSysOperationRecordFunc = async(row) => {
  row.visible = false
  const res = await deleteSysOperationRecord({ ID: row.ID })
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
const fmtBody = (value) => {
  try {
    return JSON.parse(value)
  } catch (err) {
    return value
  }
}

</script>

<style lang="scss">
.table-expand {
  padding-left: 60px;
  font-size: 0;
  label {
    width: 90px;
    color: #99a9bf;
    .el-form-item {
      margin-right: 0;
      margin-bottom: 0;
      width: 50%;
    }
  }
}
.popover-box {
  background: #112435;
  color: #f08047;
  height: 600px;
  width: 420px;
  overflow: auto;
}
.popover-box::-webkit-scrollbar {
  display: none; /* Chrome Safari */
}
</style>
