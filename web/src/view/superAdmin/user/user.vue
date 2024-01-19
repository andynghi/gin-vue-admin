<template>
  <div>
    <warning-bar title="Note: Pull down the avatar in the upper right corner to switch roles" />
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button
          type="primary"
          icon="plus"
          @click="addUser"
        >Add user</el-button>
      </div>
      <el-table
        :data="tableData"
        row-key="ID"
      >
        <el-table-column
          align="left"
          label="avatar"
          min-width="75"
        >
          <template #default="scope">
            <CustomPic
              style="margin-top:8px"
              :pic-src="scope.row.headerImg"
            />
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="ID"
          min-width="50"
          prop="ID"
        />
        <el-table-column
          align="left"
          label="username"
          min-width="150"
          prop="userName"
        />
        <el-table-column
          align="left"
          label="nickname"
          min-width="150"
          prop="nickName"
        />
        <el-table-column
          align="left"
          label="mobile phone number"
          min-width="180"
          prop="phone"
        />
        <el-table-column
          align="left"
          label="email"
          min-width="180"
          prop="email"
        />
        <el-table-column
          align="left"
          label="User role"
          min-width="200"
        >
          <template #default="scope">
            <el-cascader
              v-model="scope.row.authorityIds"
              :options="authOptions"
              :show-all-levels="false"
              collapse-tags
              :props="{ multiple:true,checkStrictly: true,label:'authorityName',value:'authorityId',disabled:'disabled',emitPath:false}"
              :clearable="false"
              @visible-change="(flag)=>{changeAuthority(scope.row,flag,0)}"
              @remove-tag="(removeAuth)=>{changeAuthority(scope.row,false,removeAuth)}"
            />
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="enable"
          min-width="150"
        >
          <template #default="scope">
            <el-switch
              v-model="scope.row.enable"
              inline-prompt
              :active-value="1"
              :inactive-value="2"
              @change="()=>{switchEnable(scope.row)}"
            />
          </template>
        </el-table-column>

        <el-table-column
          label="operation"
          min-width="250"
          fixed="right"
        >
          <template #default="scope">
            <el-popover
              v-model:visible="scope.row.visible"
              placement="top"
              width="160"
            >
              <p>Are you sure you want to delete this user</p>
              <div style="text-align: right; margin-top: 8px;">
                <el-button
                  type="primary"
                  link
                  @click="scope.row.visible = false"
                >Cancel</el-button>
                <el-button
                  type="primary"
                  @click="deleteUserFunc(scope.row)"
                >OK</el-button>
              </div>
              <template #reference>
                <el-button
                  type="primary"
                  link
                  icon="delete"
                >Delete</el-button>
              </template>
            </el-popover>
            <el-button
              type="primary"
              link
              icon="edit"
              @click="openEdit(scope.row)"
            >Edit</el-button>
            <el-button
              type="primary"
              link
              icon="magic-stick"
              @click="resetPasswordFunc(scope.row)"
            >Reset password</el-button>
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
      v-model="addUserDialog"
      title="user"
      :show-close="false"
      :close-on-press-escape="false"
      :close-on-click-modal="false"
    >
      <div style="height:60vh;overflow:auto;padding:0 12px;">
        <el-form
          ref="userForm"
          :rules="rules"
          :model="userInfo"
          label-width="80px"
        >
          <el-form-item
            v-if="dialogFlag === 'add'"
            label="username"
            prop="userName"
          >
            <el-input v-model="userInfo.userName" />
          </el-form-item>
          <el-form-item
            v-if="dialogFlag === 'add'"
            label="password"
            prop="password"
          >
            <el-input v-model="userInfo.password" />
          </el-form-item>
          <el-form-item
            label="nickname"
            prop="nickName"
          >
            <el-input v-model="userInfo.nickName" />
          </el-form-item>
          <el-form-item
            label="mobile phone number"
            prop="phone"
          >
            <el-input v-model="userInfo.phone" />
          </el-form-item>
          <el-form-item
            label="email"
            prop="email"
          >
            <el-input v-model="userInfo.email" />
          </el-form-item>
          <el-form-item
            label="User role"
            prop="authorityId"
          >
            <el-cascader
              v-model="userInfo.authorityIds"
              style="width:100%"
              :options="authOptions"
              :show-all-levels="false"
              :props="{ multiple:true,checkStrictly: true,label:'authorityName',value:'authorityId',disabled:'disabled',emitPath:false}"
              :clearable="false"
            />
          </el-form-item>
          <el-form-item
            label="enable"
            prop="disabled"
          >
            <el-switch
              v-model="userInfo.enable"
              inline-prompt
              :active-value="1"
              :inactive-value="2"
            />
          </el-form-item>
          <el-form-item
            label="avatar"
            label-width="80px"
          >
            <div
              style="display:inline-block"
              @click="openHeaderChange"
            >
              <img
                v-if="userInfo.headerImg"
                alt="avatar"
                class="header-img-box"
                :src="(userInfo.headerImg && userInfo.headerImg.slice(0, 4) !== 'http')?path+userInfo.headerImg:userInfo.headerImg"
              >
              <div
                v-else
                class="header-img-box"
              >Select from media library</div>
            </div>
          </el-form-item>

        </el-form>

      </div>

      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeAddUserDialog">Cancel</el-button>
          <el-button
            type="primary"
            @click="enterAddUserDialog"
          >Confirm</el-button>
        </div>
      </template>
    </el-dialog>
    <ChooseImg
      ref="chooseImg"
      :target="userInfo"
      :target-key="`headerImg`"
    />
  </div>
</template>


<script setup>

import {
  getUserList,
  setUserAuthorities,
  register,
  deleteUser
} from '@/api/user'

import { getAuthorityList } from '@/api/authority'
import CustomPic from '@/components/customPic/index.vue'
import ChooseImg from '@/components/chooseImg/index.vue'
import WarningBar from '@/components/warningBar/warningBar.vue'
import { setUserInfo, resetPassword } from '@/api/user.js'

import { nextTick, ref, watch } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'

defineOptions({
  name: 'User',
})

const path = ref(import.meta.env.VITE_BASE_API + '/')
// Initialization related
const setAuthorityOptions = (AuthorityData, optionsData) => {
  AuthorityData &&
        AuthorityData.forEach(item => {
          if (item.children && item.children.length) {
            const option = {
              authorityId: item.authorityId,
              authorityName: item.authorityName,
              children: []
            }
            setAuthorityOptions(item.children, option.children)
            optionsData.push(option)
          } else {
            const option = {
              authorityId: item.authorityId,
              authorityName: item.authorityName
            }
            optionsData.push(option)
          }
        })
}

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
  const table = await getUserList({ page: page.value, pageSize: pageSize.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

watch(() => tableData.value, () => {
  setAuthorityIds()
})

const initPage = async() => {
  getTableData()
  const res = await getAuthorityList({ page: 1, pageSize: 999 })
  setOptions(res.data.list)
}

initPage()

const resetPasswordFunc = (row) => {
  ElMessageBox.confirm(
    'Reset this user's password to 123456?',
    'warn',
    {
      confirmButtonText: 'OK',
      cancelButtonText: 'Cancel',
      type: 'warning',
    }
  ).then(async() => {
    const res = await resetPassword({
      ID: row.ID,
    })
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: res.msg,
      })
    } else {
      ElMessage({
        type: 'error',
        message: res.msg,
      })
    }
  })
}
const setAuthorityIds = () => {
  tableData.value && tableData.value.forEach((user) => {
    user.authorityIds = user.authorities && user.authorities.map(i => {
      return i.authorityId
    })
  })
}

const chooseImg = ref(null)
const openHeaderChange = () => {
  chooseImg.value.open()
}

const authOptions = ref([])
const setOptions = (authData) => {
  authOptions.value = []
  setAuthorityOptions(authData, authOptions.value)
}

const deleteUserFunc = async(row) => {
  const res = await deleteUser({ id: row.ID })
  if (res.code === 0) {
    ElMessage.success('Deletion successful')
    row.visible = false
    await getTableData()
  }
}

// Pop-up window related
const userInfo = ref({
  username: '',
  password: '',
  nickName: '',
  headerImg: '',
  authorityId: '',
  authorityIds: [],
  enable: 1,
})

const rules = ref({
  userName: [
    { required: true, message: 'Please enter username', trigger: 'blur' },
    { min: 5, message: 'minimum 5 characters', trigger: 'blur' }
  ],
  password: [
    { required: true, message: 'Please enter user password', trigger: 'blur' },
    { min: 6, message: 'minimum 6 characters', trigger: 'blur' }
  ],
  nickName: [
    { required: true, message: 'Please enter user nickname', trigger: 'blur' }
  ],
  phone: [
    { pattern: /^1([38][0-9]|4[014-9]|[59][0-35-9]|6[2567]|7[0-8])\d{8 }$/, message: 'Please enter a legal mobile phone number', trigger: 'blur' },
  ],
  email: [
    { pattern: /^([0-9A-Za-z\-_.]+)@([0-9a-z]+\.[a-z]{2,3}(\.[a-z]{2} )?)$/g, message: 'Please enter the correct email', trigger: 'blur' },
  ],
  authorityId: [
    { required: true, message: 'Please select a user role', trigger: 'blur' }
  ]
})
const userForm = ref(null)
const enterAddUserDialog = async() => {
  userInfo.value.authorityId = userInfo.value.authorityIds[0]
  userForm.value.validate(async valid => {
    if (valid) {
      const req = {
        ...userInfo.value
      }
      if (dialogFlag.value === 'add') {
        const res = await register(req)
        if (res.code === 0) {
          ElMessage({ type: 'success', message: 'Creation successful' })
          await getTableData()
          closeAddUserDialog()
        }
      }
      if (dialogFlag.value === 'edit') {
        const res = await setUserInfo(req)
        if (res.code === 0) {
          ElMessage({ type: 'success', message: 'Editing successful' })
          await getTableData()
          closeAddUserDialog()
        }
      }
    }
  })
}

const addUserDialog = ref(false)
const closeAddUserDialog = () => {
  userForm.value.resetFields()
  userInfo.value.headerImg = ''
  userInfo.value.authorityIds = []
  addUserDialog.value = false
}

const dialogFlag = ref('add')

const addUser = () => {
  dialogFlag.value = 'add'
  addUserDialog.value = true
}

const tempAuth = {}
const changeAuthority = async(row, flag, removeAuth) => {
  if (flag) {
    if (!removeAuth) {
      tempAuth[row.ID] = [...row.authorityIds]
    }
    return
  }
  await nextTick()
  const res = await setUserAuthorities({
    ID: row.ID,
    authorityIds: row.authorityIds
  })
  if (res.code === 0) {
    ElMessage({ type: 'success', message: 'Role setting successful' })
  } else {
    if (!removeAuth) {
      row.authorityIds = [...tempAuth[row.ID]]
      delete tempAuth[row.ID]
    } else {
      row.authorityIds = [removeAuth, ...row.authorityIds]
    }
  }
}

const openEdit = (row) => {
  dialogFlag.value = 'edit'
  userInfo.value = JSON.parse(JSON.stringify(row))
  addUserDialog.value = true
}

const switchEnable = async(row) => {
  userInfo.value = JSON.parse(JSON.stringify(row))
  await nextTick()
  const req = {
    ...userInfo.value
  }
  const res = await setUserInfo(req)
  if (res.code === 0) {
    ElMessage({ type: 'success', message: `${req.enable === 2 ? 'Disable' : 'Enable'}Success` })
    await getTableData()
    userInfo.value.headerImg = ''
    userInfo.value.authorityIds = []
  }
}

</script>

<style lang="scss">
  .header-img-box {
    @apply w-52 h-52 border border-solid border-gray-300 rounded-xl flex justify-center items-center cursor-pointer;
 }
</style>
