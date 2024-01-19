<template>
  <div>
    <div class="grid grid-cols-12 w-full gap-2">
      <div class="col-span-3 h-full">
        <div class="w-full h-full bg-white px-4 py-8 rounded-lg shadow-lg box-border">
          <div class="user-card px-6 text-center bg-white shrink-0">
            <div class="flex justify-center">
              <SelectImage
                v-model="userStore.userInfo.headerImg"
                file-type="image"
              />
            </div>
            <div class="py-6 text-center">
              <p
                v-if="!editFlag"
                class="text-3xl flex justify-center items-center gap-4"
              >
                {{ userStore.userInfo.nickName }}
                <el-icon
                  class="cursor-pointer text-sm"
                  color="#66b1ff"
                  @click="openEdit"
                >
                  <edit />
                </el-icon>
              </p>
              <p
                v-if="editFlag"
                class="flex justify-center items-center gap-4"
              >
                <el-input v-model="nickName" />
                <el-icon
                  class="cursor-pointer"
                  color="#67c23a"
                  @click="enterEdit"
                >
                  <check />
                </el-icon>
                <el-icon
                  class="cursor-pointer"
                  color="#f23c3c"
                  @click="closeEdit"
                >
                  <close />
                </el-icon>
              </p>
              <p class="text-gray-500 mt-2 text-md">This guy is lazy and has nothing left</p>
            </div>
            <div class="w-full h-full text-left">
              <ul class="inline-block h-full w-full">
                <li class="info-list">
                  <el-icon>
                    <user />
                  </el-icon>
                  {{ userStore.userInfo.nickName }}
                </li>
                <el-tooltip
                  class="item"
                  effect="light"
                  content="Beijing Reversal Aurora Technology Co., Ltd.-Technical Department-Front-End Business Group"
                  placement="top"
                >
                  <li class="info-list">
                    <el-icon>
                      <data-analysis />
                    </el-icon>
                    Beijing Reversal Aurora Technology Co., Ltd.-Technical Department-Front-End Business Group
                  </li>
                </el-tooltip>
                <li class="info-list">
                  <el-icon>
                    <video-camera />
                  </el-icon>
                  Chaoyang District, Beijing, China
                </li>
                <el-tooltip
                  class="item"
                  effect="light"
                  content="GoLang/JavaScript/Vue/Gorm"
                  placement="top"
                >
                  <li class="info-list">
                    <el-icon>
                      <medal />
                    </el-icon>
                    GoLang/JavaScript/Vue/Gorm
                  </li>
                </el-tooltip>
              </ul>
            </div>
          </div>
        </div>
      </div>
      <div class="col-span-9 ">
        <div class="bg-white h-full px-4 py-8 rounded-lg shadow-lg box-border">
          <el-tabs
            v-model="activeName"
            @tab-click="handleClick"
          >
            <el-tab-pane
              label="Account binding"
              name="second"
            >
              <ul>
                <li class="borderd">
                  <p class="pb-2.5 text-xl text-gray-600">Secret mobile phone</p>
                  <p class="pb-2.5 text-lg text-gray-400">
                    Bound mobile phone: {{ userStore.userInfo.phone }}
                    <a
                      href="javascript:void(0)"
                      class="float-right text-blue-400"
                      @click="changePhoneFlag = true"
                    >Modify now</a>
                  </p>
                </li>
                <li class="borderd pt-2.5">
                  <p class="pb-2.5 text-xl text-gray-600">Secret email</p>
                  <p class="pb-2.5 text-lg text-gray-400">
                    Bound email address: {{ userStore.userInfo.email }}
                    <a
                      href="javascript:void(0)"
                      class="float-right text-blue-400"
                      @click="changeEmailFlag = true"
                    >Modify now</a>
                  </p>
                </li>
                <li class="borderd pt-2.5">
                  <p class="pb-2.5 text-xl text-gray-600">Security issue</p>
                  <p class="pb-2.5 text-lg text-gray-400">
                    No security question set
                    <a
                      href="javascript:void(0)"
                      class="float-right text-blue-400"
                    >Go to Settings</a>
                  </p>
                </li>
                <li class="borderd pt-2.5">
                  <p class="pb-2.5 text-xl text-gray-600">Change password</p>
                  <p class="pb-2.5 text-lg text-gray-400">
                    Change personal password
                    <a
                      href="javascript:void(0)"
                      class="float-right text-blue-400"
                      @click="showPassword = true"
                    >Change password</a>
                  </p>
                </li>
              </ul>
            </el-tab-pane>
          </el-tabs>
        </div>
      </div>
    </div>

    <el-dialog
      v-model="showPassword"
      title="Change Password"
      width="360px"
      @close="clearPassword"
    >
      <el-form
        ref="modifyPwdForm"
        :model="pwdModify"
        :rules="rules"
        label-width="80px"
      >
        <el-form-item
          :minlength="6"
          label="Original password"
          prop="password"
        >
          <el-input
            v-model="pwdModify.password"
            show-password
          />
        </el-form-item>
        <el-form-item
          :minlength="6"
          label="new password"
          prop="newPassword"
        >
          <el-input
            v-model="pwdModify.newPassword"
            show-password
          />
        </el-form-item>
        <el-form-item
          :minlength="6"
          label="Confirm password"
          prop="confirmPassword"
        >
          <el-input
            v-model="pwdModify.confirmPassword"
            show-password
          />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button

            @click="showPassword = false"
          >Cancel</el-button>
          <el-button

            type="primary"
            @click="savePassword"
          >Confirm</el-button>
        </div>
      </template>
    </el-dialog>

    <el-dialog
      v-model="changePhoneFlag"
      title="Bind mobile phone"
      width="600px"
    >
      <el-form :model="phoneForm">
        <el-form-item
          label="mobile phone number"
          label-width="120px"
        >
          <el-input
            v-model="phoneForm.phone"
            placeholder="Please enter your mobile phone number"
            autocomplete="off"
          />
        </el-form-item>
        <el-form-item
          label="Verification code"
          label-width="120px"
        >
          <div class="flex w-full gap-4">
            <el-input
              v-model="phoneForm.code"
              class="flex-1"
              autocomplete="off"
              placeholder="Please design your own SMS service, just write it here for simulation"
              style="width:300px"
            />
            <el-button
              type="primary"
              :disabled="time>0"
              @click="getCode"
            >{{ time>0?`(${time}s) and then re-obtain`:'Get verification code' }}</el-button>
          </div>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button

            @click="closeChangePhone"
          >Cancel</el-button>
          <el-button
            type="primary"

            @click="changePhone"
          >Change</el-button>
        </span>
      </template>
    </el-dialog>

    <el-dialog
      v-model="changeEmailFlag"
      title="Bind email"
      width="600px"
    >
      <el-form :model="emailForm">
        <el-form-item
          label="email"
          label-width="120px"
        >
          <el-input
            v-model="emailForm.email"
            placeholder="Please enter your email address"
            autocomplete="off"
          />
        </el-form-item>
        <el-form-item
          label="Verification code"
          label-width="120px"
        >
          <div class="flex w-full gap-4">
            <el-input
              v-model="emailForm.code"
              class="flex-1"
              placeholder="Please design your own email service, just write it here for simulation"
              autocomplete="off"
              style="width:300px"
            />
            <el-button
              type="primary"
              :disabled="emailTime>0"
              @click="getEmailCode"
            >{{ emailTime>0?`(${emailTime}s) and then re-obtain`:'Get verification code' }}</el-button>
          </div>
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button

            @click="closeChangeEmail"
          >Cancel</el-button>
          <el-button
            type="primary"

            @click="changeEmail"
          >Change</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { setSelfInfo, changePassword } from '@/api/user.js'
import { reactive, ref, watch } from 'vue'
import { ElMessage } from 'element-plus'
import { useUserStore } from '@/pinia/modules/user'
import SelectImage from '@/components/selectImage/selectImage.vue'

defineOptions({
  name: 'Person',
})

const activeName = ref('second')
const rules = reactive({
  password: [
    { required: true, message: 'Please enter password', trigger: 'blur' },
    { min: 6, message: 'minimum 6 characters', trigger: 'blur' },
  ],
  newPassword: [
    { required: true, message: 'Please enter a new password', trigger: 'blur' },
    { min: 6, message: 'minimum 6 characters', trigger: 'blur' },
  ],
  confirmPassword: [
    { required: true, message: 'Please enter the confirmation password', trigger: 'blur' },
    { min: 6, message: 'minimum 6 characters', trigger: 'blur' },
    {
      validator: (rule, value, callback) => {
        if (value !== pwdModify.value.newPassword) {
          callback(new Error('Two passwords are inconsistent'))
        } else {
          callback()
        }
      },
      trigger: 'blur',
    },
  ],
})

const userStore = useUserStore()
const modifyPwdForm = ref(null)
const showPassword = ref(false)
const pwdModify = ref({})
const nickName = ref('')
const editFlag = ref(false)
const savePassword = async() => {
  modifyPwdForm.value.validate((valid) => {
    if (valid) {
      changePassword({
        password: pwdModify.value.password,
        newPassword: pwdModify.value.newPassword,
      }).then((res) => {
        if (res.code === 0) {
          ElMessage.success('Password changed successfully!')
        }
        showPassword.value = false
      })
    } else {
      return false
    }
  })
}

const clearPassword = () => {
  pwdModify.value = {
    password: '',
    newPassword: '',
    confirmPassword: '',
  }
  modifyPwdForm.value.clearValidate()
}

watch(() => userStore.userInfo.headerImg, async(val) => {
  const res = await setSelfInfo({ headerImg: val })
  if (res.code === 0) {
    userStore.ResetUserInfo({ headerImg: val })
    ElMessage({
      type: 'success',
      message: 'Setup successful',
    })
  }
})

const openEdit = () => {
  nickName.value = userStore.userInfo.nickName
  editFlag.value = true
}

const closeEdit = () => {
  nickName.value = ''
  editFlag.value = false
}

const enterEdit = async() => {
  const res = await setSelfInfo({
    nickName: nickName.value
  })
  if (res.code === 0) {
    userStore.ResetUserInfo({ nickName: nickName.value })
    ElMessage({
      type: 'success',
      message: 'Setup successful',
    })
  }
  nickName.value = ''
  editFlag.value = false
}

const handleClick = (tab, event) => {
  console.log(tab, event)
}

const changePhoneFlag = ref(false)
const time = ref(0)
const phoneForm = reactive({
  phone: '',
  code: ''
})

const getCode = async() => {
  time.value = 60
  let timer = setInterval(() => {
    time.value--
    if (time.value <= 0) {
      clearInterval(timer)
      timer = null
    }
  }, 1000)
}

const closeChangePhone = () => {
  changePhoneFlag.value = false
  phoneForm.phone = ''
  phoneForm.code = ''
}

const changePhone = async() => {
  const res = await setSelfInfo({ phone: phoneForm.phone })
  if (res.code === 0) {
    ElMessage.success('Modification successful')
    userStore.ResetUserInfo({ phone: phoneForm.phone })
    closeChangePhone()
  }
}

const changeEmailFlag = ref(false)
const emailTime = ref(0)
const emailForm = reactive({
  email: '',
  code: ''
})

const getEmailCode = async() => {
  emailTime.value = 60
  let timer = setInterval(() => {
    emailTime.value--
    if (emailTime.value <= 0) {
      clearInterval(timer)
      timer = null
    }
  }, 1000)
}

const closeChangeEmail = () => {
  changeEmailFlag.value = false
  emailForm.email = ''
  emailForm.code = ''
}

const changeEmail = async() => {
  const res = await setSelfInfo({ email: emailForm.email })
  if (res.code === 0) {
    ElMessage.success('Modification successful')
    userStore.ResetUserInfo({ email: emailForm.email })
    closeChangeEmail()
  }
}

</script>

<style lang="scss">
.borderd {
  @apply border-b-2 border-solid border-gray-100 border-t-0 border-r-0 border-l-0;
    &:last-child{
      @apply border-b-0;
    }
 }

.info-list{
  @apply w-full whitespace-nowrap overflow-hidden text-ellipsis py-3 text-lg text-gray-700
}

</style>
