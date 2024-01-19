<template>
  <div>
    <div class="gva-table-box">
      <div class="gva-btn-list justify-between">
        <span class="text font-bold">Dictionary details</span>
        <el-button
          type="primary"
          icon="plus"
          @click="openDialog"
        >Add dictionary item</el-button>
      </div>
      <el-table
        ref="multipleTable"
        :data="tableData"
        style="width: 100%"
        tooltip-effect="dark"
        row-key="ID"
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
          label="display value"
          prop="label"
        />

        <el-table-column
          align="left"
          label="dictionary value"
          prop="value"
        />

        <el-table-column
          align="left"
          label="extended value"
          prop="extend"
        />

        <el-table-column
          align="left"
          label="enabled status"
          prop="status"
          width="120"
        >
          <template #default="scope">{{ formatBoolean(scope.row.status) }}</template>
        </el-table-column>

        <el-table-column
          align="left"
          label="sort mark"
          prop="sort"
          width="120"
        />

        <el-table-column
          align="left"
          label="operation"
          width="180"
        >
          <template #default="scope">
            <el-button
              type="primary"
              link
              icon="edit"
              @click="updateSysDictionaryDetailFunc(scope.row)"
            >Change</el-button>
            <el-popover
              v-model:visible="scope.row.visible"
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
                  @click="deleteSysDictionaryDetailFunc(scope.row)"
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
      :title="type==='create'?'Add dictionary item':'Modify dictionary item'"
    >
      <el-form
        ref="dialogForm"
        :model="formData"
        :rules="rules"
        label-width="110px"
      >
        <el-form-item
          label="display value"
          prop="label"
        >
          <el-input
            v-model="formData.label"
            placeholder="Please enter display value"
            clearable
            :style="{width: '100%'}"
          />
        </el-form-item>
        <el-form-item
          label="dictionary value"
          prop="value"
        >
          <el-input-number
            v-model.number="formData.value"
            step-strictly
            :step="1"
            placeholder="Please enter dictionary value"
            clearable
            :style="{width: '100%'}"
            min="-2147483648"
            max="2147483647"
          />
        </el-form-item>
        <el-form-item
          label="extended value"
          prop="extend"
        >
          <el-input
            v-model="formData.extend"
            placeholder="Please enter extended value"
            clearable
            :style="{width: '100%'}"
          />
        </el-form-item>
        <el-form-item
          label="enabled status"
          prop="status"
          required
        >
          <el-switch
            v-model="formData.status"
            active-text="Open"
            inactive-text="Inactive"
          />
        </el-form-item>
        <el-form-item
          label="sort mark"
          prop="sort"
        >
          <el-input-number
            v-model.number="formData.sort"
            placeholder="sort mark"
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
  createSysDictionaryDetail,
  deleteSysDictionaryDetail,
  updateSysDictionaryDetail,
  findSysDictionaryDetail,
  getSysDictionaryDetailList
} from '@/api/sysDictionaryDetail' // Please replace the address here.
import { ref, watch } from 'vue'
import { ElMessage } from 'element-plus'
import { formatBoolean, formatDate } from '@/utils/format'

defineOptions({
  name: 'SysDictionaryDetail'
})

const props = defineProps({
  sysDictionaryID: {
    type: Number,
    default: 0
  }
})

const formData = ref({
  label: null,
  value: null,
  status: true,
  sort: null
})
const rules = ref({
  label: [
    {
      required: true,
      message: 'Please enter the display value',
      trigger: 'blur'
    }
  ],
  value: [
    {
      required: true,
      message: 'Please enter the dictionary value',
      trigger: 'blur'
    }
  ],
  sort: [
    {
      required: true,
      message: 'Sort mark',
      trigger: 'blur'
    }
  ]
})

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])

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
  const table = await getSysDictionaryDetailList({
    page: page.value,
    pageSize: pageSize.value,
    sysDictionaryID: props.sysDictionaryID
  })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

const type = ref('')
const dialogFormVisible = ref(false)
const updateSysDictionaryDetailFunc = async(row) => {
  const res = await findSysDictionaryDetail({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.reSysDictionaryDetail
    dialogFormVisible.value = true
  }
}

const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    label: null,
    value: null,
    status: true,
    sort: null,
    sysDictionaryID: props.sysDictionaryID
  }
}
const deleteSysDictionaryDetailFunc = async(row) => {
  row.visible = false
  const res = await deleteSysDictionaryDetail({ ID: row.ID })
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

const dialogForm = ref(null)
const enterDialog = async() => {
  dialogForm.value.validate(async valid => {
    formData.value.sysDictionaryID = props.sysDictionaryID
    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createSysDictionaryDetail(formData.value)
        break
      case 'update':
        res = await updateSysDictionaryDetail(formData.value)
        break
      default:
        res = await createSysDictionaryDetail(formData.value)
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
const openDialog = () => {
  type.value = 'create'
  dialogFormVisible.value = true
}

watch(() => props.sysDictionaryID, () => {
  getTableData()
})

</script>

<style>
</style>
