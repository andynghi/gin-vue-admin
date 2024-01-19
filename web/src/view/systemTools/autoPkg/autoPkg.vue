<template>
  <div>
    <warning-bar
      href="https://www.bilibili.com/video/BV1kv4y1g7nT?p=3"
      title="This function is for development environment use and is not recommended for release to production. Please see the video for specific usage effects https://www.bilibili.com/video/BV1kv4y1g7nT?p=3"
    />
    <div class="gva-table-box">
      <div class="gva-btn-list gap-3 flex items-center">
        <el-button
          type="primary"
          icon="plus"
          @click="openDialog('addApi')"
        >Add</el-button>
        <el-icon
          class="cursor-pointer"
          @click="toDoc('https://www.bilibili.com/video/BV1kv4y1g7nT?p=3&vd_source=f2640257c21e3b547a790461ed94875e')"
        ><VideoCameraFilled /></el-icon>
      </div>
      <el-table :data="tableData">
        <el-table-column
          align="left"
          label="id"
          width="60"
          prop="ID"
        />
        <el-table-column
          align="left"
          label="package name"
          width="150"
          prop="packageName"
        />
        <el-table-column
          align="left"
          label="display name"
          width="150"
          prop="label"
        />
        <el-table-column
          align="left"
          label="description"
          min-width="150"
          prop="desc"
        />

        <el-table-column
          align="left"
          label="operation"
          width="200"
        >
          <template #default="scope">
            <el-button
              icon="delete"

              type="primary"
              link
              @click="deleteApiFunc(scope.row)"
            >Delete</el-button>
          </template>
        </el-table-column>
      </el-table>

    </div>

    <el-dialog
      v-model="dialogFormVisible"
      :before-close="closeDialog"
      title="Create Package"
    >
      <warning-bar title="New Pkg for automated code use" />
      <el-form
        ref="pkgForm"
        :model="form"
        :rules="rules"
        label-width="80px"
      >
        <el-form-item
          label="package name"
          prop="packageName"
        >
          <el-input
            v-model="form.packageName"
            autocomplete="off"
          />
        </el-form-item>
        <el-form-item
          label="display name"
          prop="label"
        >
          <el-input
            v-model="form.label"
            autocomplete="off"
          />
        </el-form-item>
        <el-form-item
          label="description"
          prop="desc"
        >
          <el-input
            v-model="form.desc"
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
  createPackageApi,
  getPackageApi,
  deletePackageApi,
} from '@/api/autoCode'
import { ref } from 'vue'
import WarningBar from '@/components/warningBar/warningBar.vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { toDoc } from '@/utils/doc'
import { VideoCameraFilled } from '@element-plus/icons-vue'

defineOptions({
  name: 'AutoPkg',
})

const form = ref({
  packageName: '',
  label: '',
  desc: '',
})

const validateNum = (rule, value, callback) => {
  if ((/^\d+$/.test(value[0]))) {
    callback(new Error('Cannot start with a number'))
  } else {
    callback()
  }
}

const rules = ref({
  packageName: [
    { required: true, message: 'Please enter the package name', trigger: 'blur' },
    { validator: validateNum, trigger: 'blur' }
  ],
})

const dialogFormVisible = ref(false)
const openDialog = () => {
  dialogFormVisible.value = true
}

const closeDialog = () => {
  dialogFormVisible.value = false
  form.value = {
    packageName: '',
    label: '',
    desc: '',
  }
}

const pkgForm = ref(null)
const enterDialog = async() => {
  pkgForm.value.validate(async valid => {
    if (valid) {
      const res = await createPackageApi(form.value)
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
  })
}

const tableData = ref([])
const getTableData = async() => {
  const table = await getPackageApi()
  if (table.code === 0) {
    tableData.value = table.data.pkgs
  }
}

const deleteApiFunc = async(row) => {
  ElMessageBox.confirm('This operation only deletes the pkg storage in the database. Please delete the corresponding directory structure in the backend to keep it consistent with the database!', 'Prompt', {
    confirmButtonText: 'OK',
    cancelButtonText: 'Cancel',
    type: 'warning'
  })
    .then(async() => {
      const res = await deletePackageApi(row)
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: 'Delete successfully!'
        })
        getTableData()
      }
    })
}

getTableData()
</script>
