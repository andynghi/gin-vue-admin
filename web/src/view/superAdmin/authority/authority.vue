<template>
  <div class="authority">
    <warning-bar title="Note: Pull down the avatar in the upper right corner to switch roles" />
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button
          type="primary"
          icon="plus"
          @click="addAuthority(0)"
        >Add new role</el-button>
        <el-icon
          class="cursor-pointer"
          @click="toDoc('https://www.bilibili.com/video/BV1kv4y1g7nT?p=8&vd_source=f2640257c21e3b547a790461ed94875e')"
        ><VideoCameraFilled /></el-icon>
      </div>
      <el-table
        :data="tableData"
        :tree-props="{children: 'children', hasChildren: 'hasChildren'}"
        row-key="authorityId"
        style="width: 100%"
      >
        <el-table-column
          label="role ID"
          min-width="180"
          prop="authorityId"
        />
        <el-table-column
          align="left"
          label="Character name"
          min-width="180"
          prop="authorityName"
        />
        <el-table-column
          align="left"
          label="operation"
          width="460"
        >
          <template #default="scope">
            <el-button
              icon="setting"

              type="primary"
              link
              @click="opdendrawer(scope.row)"
            >Set permissions</el-button>
            <el-button
              icon="plus"

              type="primary"
              link
              @click="addAuthority(scope.row.authorityId)"
            >Add sub-role</el-button>
            <el-button
              icon="copy-document"

              type="primary"
              link
              @click="copyAuthorityFunc(scope.row)"
            >Copy</el-button>
            <el-button
              icon="edit"

              type="primary"
              link
              @click="editAuthority(scope.row)"
            >Edit</el-button>
            <el-button
              icon="delete"

              type="primary"
              link
              @click="deleteAuth(scope.row)"
            >Delete</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>
    <!-- Add new character pop-up window -->
    <el-dialog
      v-model="dialogFormVisible"
      :title="dialogTitle"
    >
      <el-form
        ref="authorityForm"
        :model="form"
        :rules="rules"
        label-width="80px"
      >
        <el-form-item
          label="parent role"
          prop="parentId"
        >
          <el-cascader
            v-model="form.parentId"
            style="width:100%"
            :disabled="dialogType==='add'"
            :options="AuthorityOption"
            :props="{ checkStrictly: true,label:'authorityName',value:'authorityId',disabled:'disabled',emitPath:false}"
            :show-all-levels="false"
            filterable
          />
        </el-form-item>
        <el-form-item
          label="role ID"
          prop="authorityId"
        >
          <el-input
            v-model="form.authorityId"
            :disabled="dialogType==='edit'"
            autocomplete="off"
            maxlength="15"
          />
        </el-form-item>
        <el-form-item
          label="Character name"
          prop="authorityName"
        >
          <el-input
            v-model="form.authorityName"
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

    <el-drawer
      v-if="drawer"
      v-model="drawer"
      custom-class="auth-drawer"
      :with-header="false"
      size="40%"
      title="Role configuration"
    >
      <el-tabs
        :before-leave="autoEnter"
        type="border-card"
      >
        <el-tab-pane label="Character menu">
          <Menus
            ref="menus"
            :row="activeRow"
            @changeRow="changeRow"
          />
        </el-tab-pane>
        <el-tab-pane label="role api">
          <Apis
            ref="apis"
            :row="activeRow"
            @changeRow="changeRow"
          />
        </el-tab-pane>
        <el-tab-pane label="Resource permissions">
          <Datas
            ref="datas"
            :authority="tableData"
            :row="activeRow"
            @changeRow="changeRow"
          />
        </el-tab-pane>
      </el-tabs>
    </el-drawer>
  </div>
</template>

<script setup>
import {
  getAuthorityList,
  deleteAuthority,
  createAuthority,
  updateAuthority,
  copyAuthority
} from '@/api/authority'

import Menus from '@/view/superAdmin/authority/components/menus.vue'
import Apis from '@/view/superAdmin/authority/components/apis.vue'
import Datas from '@/view/superAdmin/authority/components/datas.vue'
import WarningBar from '@/components/warningBar/warningBar.vue'

import { ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { toDoc } from '@/utils/doc'
import { VideoCameraFilled } from '@element-plus/icons-vue'

defineOptions({
  name: 'Authority'
})

const mustUint = (rule, value, callback) => {
  if (!/^[0-9]*[1-9][0-9]*$/.test(value)) {
    return callback(new Error('Please enter a positive integer'))
  }
  return callback()
}

const AuthorityOption = ref([
  {
    authorityId: 0,
    authorityName: 'root role'
  }
])
const drawer = ref(false)
const dialogType = ref('add')
const activeRow = ref({})

const dialogTitle = ref('Add new role')
const dialogFormVisible = ref(false)
const apiDialogFlag = ref(false)
const copyForm = ref({})

const form = ref({
  authorityId: 0,
  authorityName: '',
  parentId: 0
})
const rules = ref({
  authorityId: [
    { required: true, message: 'Please enter the role ID', trigger: 'blur' },
    { validator: mustUint, trigger: 'blur', message: 'Must be a positive integer' }
  ],
  authorityName: [
    { required: true, message: 'Please enter the role name', trigger: 'blur' }
  ],
  parentId: [
    { required: true, message: 'Please select a parent role', trigger: 'blur' },
  ]
})

const page = ref(1)
const total = ref(0)
const pageSize = ref(999)
const tableData = ref([])
const searchInfo = ref({})

// Inquire
const getTableData = async() => {
  const table = await getAuthorityList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

const changeRow = (key, value) => {
  activeRow.value[key] = value
}
const menus = ref(null)
const apis = ref(null)
const datas = ref(null)
const autoEnter = (activeName, oldActiveName) => {
  const paneArr = [menus, apis, datas]
  if (oldActiveName) {
    if (paneArr[oldActiveName].value.needConfirm) {
      paneArr[oldActiveName].value.enterAndNext()
      paneArr[oldActiveName].value.needConfirm = false
    }
  }
}
//Copy character
const copyAuthorityFunc = (row) => {
  setOptions()
  dialogTitle.value = 'Copy role'
  dialogType.value = 'copy'
  for (const k in form.value) {
    form.value[k] = row[k]
  }
  copyForm.value = row
  dialogFormVisible.value = true
}
const opdendrawer = (row) => {
  drawer.value = true
  activeRow.value = row
}
// Delete role
const deleteAuth = (row) => {
  ElMessageBox.confirm('This operation will permanently delete the role, do you want to continue?', 'Prompt', {
    confirmButtonText: 'OK',
    cancelButtonText: 'Cancel',
    type: 'warning'
  })
    .then(async() => {
      const res = await deleteAuthority({ authorityId: row.authorityId })
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
    .catch(() => {
      ElMessage({
        type: 'info',
        message: 'Deletion canceled'
      })
    })
}
//Initialize form
const authorityForm = ref(null)
const initForm = () => {
  if (authorityForm.value) {
    authorityForm.value.resetFields()
  }
  form.value = {
    authorityId: 0,
    authorityName: '',
    parentId: 0
  }
}
// close the window
const closeDialog = () => {
  initForm()
  dialogFormVisible.value = false
  apiDialogFlag.value = false
}
// Confirm the pop-up window

const enterDialog = () => {
  authorityForm.value.validate(async valid => {
    if (valid) {
      form.value.authorityId = Number(form.value.authorityId)
      switch (dialogType.value) {
        case 'add':
          {
            const res = await createAuthority(form.value)
            if (res.code === 0) {
              ElMessage({
                type: 'success',
                message: 'Added successfully!'
              })
              getTableData()
              closeDialog()
            }
          }
          break
        case 'edit':
          {
            const res = await updateAuthority(form.value)
            if (res.code === 0) {
              ElMessage({
                type: 'success',
                message: 'Added successfully!'
              })
              getTableData()
              closeDialog()
            }
          }
          break
        case 'copy': {
          const data = {
            authority: {
              authorityId: 0,
              authorityName: '',
              datauthorityId: [],
              parentId: 0
            },
            oldAuthorityId: 0
          }
          data.authority.authorityId = form.value.authorityId
          data.authority.authorityName = form.value.authorityName
          data.authority.parentId = form.value.parentId
          data.authority.dataAuthorityId = copyForm.value.dataAuthorityId
          data.oldAuthorityId = copyForm.value.authorityId
          const res = await copyAuthority(data)
          if (res.code === 0) {
            ElMessage({
              type: 'success',
              message: 'Copying successful! '
            })
            getTableData()
          }
        }
      }

      initForm()
      dialogFormVisible.value = false
    }
  })
}
const setOptions = () => {
  AuthorityOption.value = [
    {
      authorityId: 0,
      authorityName: 'root role'
    }
  ]
  setAuthorityOptions(tableData.value, AuthorityOption.value, false)
}
const setAuthorityOptions = (AuthorityData, optionsData, disabled) => {
  form.value.authorityId = String(form.value.authorityId)
  AuthorityData &&
        AuthorityData.forEach(item => {
          if (item.children && item.children.length) {
            const option = {
              authorityId: item.authorityId,
              authorityName: item.authorityName,
              disabled: disabled || item.authorityId === form.value.authorityId,
              children: []
            }
            setAuthorityOptions(
              item.children,
              option.children,
              disabled || item.authorityId === form.value.authorityId
            )
            optionsData.push(option)
          } else {
            const option = {
              authorityId: item.authorityId,
              authorityName: item.authorityName,
              disabled: disabled || item.authorityId === form.value.authorityId
            }
            optionsData.push(option)
          }
        })
}
//Add role
const addAuthority = (parentId) => {
  initForm()
  dialogTitle.value = 'Add new role'
  dialogType.value = 'add'
  form.value.parentId = parentId
  setOptions()
  dialogFormVisible.value = true
}
//Edit role
const editAuthority = (row) => {
  setOptions()
  dialogTitle.value = 'Edit role'
  dialogType.value = 'edit'
  for (const key in form.value) {
    form.value[key] = row[key]
  }
  setOptions()
  dialogFormVisible.value = true
}

</script>

<style lang="scss">
.authority {
  .el-input-number {
    margin-left: 15px;
    span {
      display: none;
    }
  }
}
.tree-content{
  margin-top: 10px;
  height: calc(100vh - 158px);
  overflow: auto;
}

</style>
