<template>
  <div v-loading.fullscreen.lock="fullscreenLoading">
    <div class="gva-table-box">
      <warning-bar
        title="Click "File Name/Remarks" to edit the file name or comment content."
      />
      <div class="gva-btn-list">
        <upload-common
          v-model:imageCommon="imageCommon"
          @on-success="getTableData"
        />
        <upload-image
          v-model:imageUrl="imageUrl"
          :file-size="512"
          :max-w-h="1080"
          @on-success="getTableData"
        />
        <el-input
          v-model="search.keyword"
          class="keyword"
          placeholder="Please enter file name or remarks"
        />
        <el-button
          type="primary"
          icon="search"
          @click="getTableData"
        >Query</el-button>
      </div>

      <el-table :data="tableData">
        <el-table-column
          align="left"
          label="preview"
          width="100"
        >
          <template #default="scope">
            <CustomPic
              pic-type="file"
              :pic-src="scope.row.url"
              preview
            />
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="date"
          prop="UpdatedAt"
          width="180"
        >
          <template #default="scope">
            <div>{{ formatDate(scope.row.UpdatedAt) }}</div>
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="File name/Remarks"
          prop="name"
          width="180"
        >
          <template #default="scope">
            <div
              class="name"
              @click="editFileNameFunc(scope.row)"
            >{{ scope.row.name }}</div>
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="link"
          prop="url"
          min-width="300"
        />
        <el-table-column
          align="left"
          label="label"
          prop="tag"
          width="100"
        >
          <template #default="scope">
            <el-tag
              :type="scope.row.tag === 'jpg' ? 'info' : 'success'"
              disable-transitions
            >{{ scope.row.tag }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="operation"
          width="160"
        >
          <template #default="scope">
            <el-button
              icon="download"
              type="primary"
              link
              @click="downloadFile(scope.row)"
            >Download</el-button>
            <el-button
              icon="delete"
              type="primary"
              link
              @click="deleteFileFunc(scope.row)"
            >Delete</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :style="{ float: 'right', padding: '20px' }"
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
import { getFileList, deleteFile, editFileName } from '@/api/fileUploadAndDownload'
import { downloadImage } from '@/utils/downloadImg'
import CustomPic from '@/components/customPic/index.vue'
import UploadImage from '@/components/upload/image.vue'
import UploadCommon from '@/components/upload/common.vue'
import { formatDate } from '@/utils/format'
import WarningBar from '@/components/warningBar/warningBar.vue'

import { ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'

defineOptions({
  name: 'Upload',
})

const path = ref(import.meta.env.VITE_BASE_API)

const imageUrl = ref('')
const imageCommon = ref('')

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const search = ref({})
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
  const table = await getFileList({ page: page.value, pageSize: pageSize.value, ...search.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}
getTableData()

const deleteFileFunc = async(row) => {
  ElMessageBox.confirm('This operation will permanently delete the file, do you want to continue?', 'Prompt', {
    confirmButtonText: 'OK',
    cancelButtonText: 'Cancel',
    type: 'warning',
  })
    .then(async() => {
      const res = await deleteFile(row)
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: 'Delete successfully!',
        })
        if (tableData.value.length === 1 && page.value > 1) {
          page.value--
        }
        getTableData()
      }
    })
    .catch(() => {
      ElMessage({
        type: 'info',
        message: 'Deletion canceled',
      })
    })
}

const downloadFile = (row) => {
  if (row.url.indexOf('http://') > -1 || row.url.indexOf('https://') > -1) {
    downloadImage(row.url, row.name)
  } else {
    debugger
    downloadImage(path.value + '/' + row.url, row.name)
  }
}

/**
 * Edit file name or remarks
 * @param row
 * @returns {Promise<void>}
 */
const editFileNameFunc = async(row) => {
  ElMessageBox.prompt('Please enter the file name or remarks', 'Edit', {
    confirmButtonText: 'OK',
    cancelButtonText: 'Cancel',
    inputPattern: /\S/,
    inputErrorMessage: 'Cannot be empty',
    inputValue: row.name
  }).then(async({ value }) => {
    row.name = value
    // console.log(row)
    const res = await editFileName(row)
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: 'Editing successful!',
      })
      getTableData()
    }
  }).catch(() => {
    ElMessage({
      type: 'info',
      message: 'Cancel modification'
    })
  })
}
</script>

<style scoped>
.name {
  cursor: pointer;
}

</style>
