<template>
  <div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button
          type="primary"
          icon="plus"
          @click="goAutoCode(null)"
        >Add</el-button>
      </div>
      <el-table :data="tableData">
        <el-table-column
          type="selection"
          width="55"
        />
        <el-table-column
          align="left"
          label="id"
          width="60"
          prop="ID"
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
          label="structure name"
          min-width="150"
          prop="structName"
        />
        <el-table-column
          align="left"
          label="structure description"
          min-width="150"
          prop="structCNName"
        />
        <el-table-column
          align="left"
          label="table name"
          min-width="150"
          prop="tableName"
        />
        <el-table-column
          align="left"
          label="rollback mark"
          min-width="150"
          prop="flag"
        >
          <template #default="scope">
            <el-tag
              v-if="scope.row.flag"
              type="danger"

              effect="dark"
            >
              Rolled back
            </el-tag>
            <el-tag
              v-else

              type="success"
              effect="dark"
            >
              Not rolled back
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="operation"
          min-width="240"
        >
          <template #default="scope">
            <div>
              <el-button
                type="primary"
                link
                :disabled="scope.row.flag === 1"
                @click="rollbackFunc(scope.row,true)"
              >Rollback (delete table)</el-button>
              <el-button
                type="primary"
                link
                :disabled="scope.row.flag === 1"
                @click="rollbackFunc(scope.row,false)"
              >Rollback (without deleting the table)</el-button>
              <el-button
                type="primary"
                link
                @click="goAutoCode(scope.row)"
              >Reuse</el-button>
              <el-button
                type="primary"
                link
                @click="deleteRow(scope.row)"
              >Delete</el-button>
            </div>
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
import { getSysHistory, rollback, delSysHistory } from '@/api/autoCode.js'
import { useRouter } from 'vue-router'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref } from 'vue'
import { formatDate } from '@/utils/format'

defineOptions({
  name: 'AutoCodeAdmin'
})

const router = useRouter()

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
  const table = await getSysHistory({
    page: page.value,
    pageSize: pageSize.value
  })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

const deleteRow = async(row) => {
  ElMessageBox.confirm('This operation will delete this history, do you want to continue?', 'Prompt', {
    confirmButtonText: 'OK',
    cancelButtonText: 'Cancel',
    type: 'warning'
  }).then(async() => {
    const res = await delSysHistory({ id: Number(row.ID) })
    if (res.code === 0) {
      ElMessage.success('Deletion successful')
      getTableData()
    }
  })
}
const rollbackFunc = async(row, flag) => {
  if (flag) {
    ElMessageBox.confirm(`This operation will delete the automatically created files and api (the table will be deleted!!!), do you want to continue?`, 'Prompt', {
      confirmButtonText: 'OK',
      cancelButtonText: 'Cancel',
      type: 'warning'
    }).then(async() => {
      ElMessageBox.confirm(`This operation will delete automatically created files and APIs (the table will be deleted!!!), please continue to confirm!!!`, 'The table will be deleted', {
        confirmButtonText: 'OK',
        cancelButtonText: 'Cancel',
        type: 'warning'
      }).then(async() => {
        ElMessageBox.confirm(`This operation will delete automatically created files and APIs (the table will be deleted!!!), please continue to confirm!!!`, 'The table will be deleted', {
          confirmButtonText: 'OK',
          cancelButtonText: 'Cancel',
          type: 'warning'
        }).then(async() => {
          const res = await rollback({ id: Number(row.ID), deleteTable: flag })
          if (res.code === 0) {
            ElMessage.success('Rollback successful')
            getTableData()
          }
        })
      })
    })
  } else {
    ElMessageBox.confirm(`This operation will delete automatically created files and APIs, do you want to continue?`, 'Prompt', {
      confirmButtonText: 'OK',
      cancelButtonText: 'Cancel',
      type: 'warning'
    }).then(async() => {
      const res = await rollback({ id: Number(row.ID), deleteTable: flag })
      if (res.code === 0) {
        ElMessage.success('Rollback successful')
        getTableData()
      }
    })
  }
}
const goAutoCode = (row) => {
  if (row) {
    router.push({ name: 'autoCodeEdit', params: {
      id: row.ID
    }})
  } else {
    router.push({ name: 'autoCode' })
  }
}

</script>
