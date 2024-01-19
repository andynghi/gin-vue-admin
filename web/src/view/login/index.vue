<template>
  <div
    id="userLayout"
    class="w-full h-full relative"
  >
    <div
      class="rounded-lg flex items-center justify-evenly w-full h-full bg-white md:w-screen md:h-screen md:bg-[#194bfb]"
    >
      <div class="md:w-3/5 w-10/12 h-full flex items-center justify-evenly">
        <div class="oblique h-[130%] w-3/5 bg-white transform -rotate-12 absolute -ml-52" />
        <!-- Split inclined block -->
        <div class="z-[999] pt-12 pb-10 md:w-96 w-full  rounded-lg flex flex-col justify-between box-border">
          <div>
            <div class="flex items-center justify-center">

              <img
                class="w-24"
                :src="$GIN_VUE_ADMIN.appLogo"
                alt
              >
            </div>
            <div class="mb-9">
              <p class="text-center text-4xl font-bold">{{ $GIN_VUE_ADMIN.appName }}</p>
              <p class="text-center text-sm font-normal text-gray-500 mt-2.5">A management platform using Golang and Vue
              </p>
            </div>
            <el-form
              ref="loginForm"
              :model="loginFormData"
              :rules="rules"
              :validate-on-rule-change="false"
              @keyup.enter="submitForm"
            >
              <el-form-item
                prop="username"
                class="mb-6"
              >
                <el-input
                  v-model="loginFormData.username"
                  size="large"
                  placeholder="Please enter username"
                  suffix-icon="user"
                />
              </el-form-item>
              <el-form-item
                prop="password"
                class="mb-6"
              >
                <el-input
                  v-model="loginFormData.password"
                  show-password
                  size="large"
                  type="password"
                  placeholder="Please enter password"
                />
              </el-form-item>
              <el-form-item
                v-if="loginFormData.openCaptcha"
                prop="captcha"
                class="mb-6"
              >
                <div class="flex w-full justify-between">
                  <el-input
                    v-model="loginFormData.captcha"
                    placeholder="Please enter the verification code"
                    size="large"
                    class="flex-1 mr-5"
                  />
                  <div class="w-1/3 h-11 bg-[#c3d4f2] rounded">
                    <img
                      v-if="picPath"
                      class="w-full h-full"
                      :src="picPath"
                      alt="Please enter the verification code"
                      @click="loginVerify()"
                    >
                  </div>
                </div>
              </el-form-item>
              <el-form-item class="mb-6">
                <el-button
                  class="shadow shadow-blue-600 h-11 w-full"
                  type="primary"
                  size="large"
                  @click="submitForm"
                >Login</el-button>
              </el-form-item>
              <el-form-item class="mb-6">
                <el-button
                  class="shadow shadow-blue-600 h-11 w-full"
                  type="primary"
                  size="large"
                  @click="checkInit"
                >Go to initialization</el-button>

              </el-form-item>
            </el-form>
          </div>
        </div>
      </div>
      <div class="hidden md:block w-1/2 h-full float-right bg-[#194bfb]"><img
        class="h-full"
        src="@/assets/login_right_banner.jpg"
        alt="banner"
      ></div>
    </div>

    <BottomInfo class="left-0 right-0 absolute bottom-3 mx-auto  w-full z-20">
      <div class="links items-center justify-center gap-2 hidden md:flex">
        <a
          href="http://doc.henrongyi.top/"
          target="_blank"
        >
          <img
            src="@/assets/docs.png"
            class="w-8 h-8"
            alt="documentation"
          >
        </a>
        <a
          href="https://support.qq.com/product/371961"
          target="_blank"
        >
          <img
            src="@/assets/kefu.png"
            class="w-8 h-8"
            alt="Customer Service"
          >
        </a>
        <a
          href="https://github.com/flipped-aurora/gin-vue-admin"
          target="_blank"
        >
          <img
            src="@/assets/github.png"
            class="w-8 h-8"
            alt="github"
          >
        </a>
        <a
          href="https://space.bilibili.com/322210472"
          target="_blank"
        >
          <img
            src="@/assets/video.png"
            class="w-8 h-8"
            alt="Video site"
          >
        </a>
      </div>
    </BottomInfo>
  </div>
</template>

<script setup>
import { captcha } from '@/api/user'
import { checkDB } from '@/api/initdb'
import BottomInfo from '@/view/layout/bottomInfo/bottomInfo.vue'
import { reactive, ref } from 'vue'
import { ElMessage } from 'element-plus'
import { useRouter } from 'vue-router'
import { useUserStore } from '@/pinia/modules/user'

defineOptions({
  name: 'Login',
})

const router = useRouter()
//verification function
const checkUsername = (rule, value, callback) => {
  if (value.length < 5) {
    return callback(new Error('Please enter the correct user name'))
  } else {
    callback()
  }
}
const checkPassword = (rule, value, callback) => {
  if (value.length < 6) {
    return callback(new Error('Please enter the correct password'))
  } else {
    callback()
  }
}

// get verification code
const loginVerify = () => {
  captcha({}).then(async(ele) => {
    rules.captcha.push({
      max: ele.data.captchaLength,
      min: ele.data.captchaLength,
      message: `Please enter the ${ele.data.captchaLength} digit verification code`,
      trigger: 'blur',
    })
    picPath.value = ele.data.picPath
    loginFormData.captchaId = ele.data.captchaId
    loginFormData.openCaptcha = ele.data.openCaptcha
  })
}
loginVerify()

// Login related operations
const loginForm = ref(null)
const picPath = ref('')
const loginFormData = reactive({
  username: 'admin',
  password: '123456',
  captcha: '',
  captchaId: '',
  openCaptcha: false,
})
const rules = reactive({
  username: [{ validator: checkUsername, trigger: 'blur' }],
  password: [{ validator: checkPassword, trigger: 'blur' }],
  captcha: [
    {
      message: 'Verification code format is incorrect',
      trigger: 'blur',
    },
  ],
})

const userStore = useUserStore()
const login = async() => {
  return await userStore.LoginIn(loginFormData)
}
const submitForm = () => {
  loginForm.value.validate(async(v) => {
    if (v) {
      const flag = await login()
      if (!flag) {
        loginVerify()
      }
    } else {
      ElMessage({
        type: 'error',
        message: 'Please fill in the login information correctly',
        showClose: true,
      })
      loginVerify()
      return false
    }
  })
}

//Jump initialization
const checkInit = async() => {
  const res = await checkDB()
  if (res.code === 0) {
    if (res.data?.needInit) {
      userStore.NeedInit()
      router.push({ name: 'Init' })
    } else {
      ElMessage({
        type: 'info',
        message: 'Database information has been configured and cannot be initialized',
      })
    }
  }
}

</script>
