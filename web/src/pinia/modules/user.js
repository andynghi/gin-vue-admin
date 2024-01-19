import { login, getUserInfo, setSelfInfo } from '@/api/user'
import { jsonInBlacklist } from '@/api/jwt'
import router from '@/router/index'
import { ElLoading, ElMessage } from 'element-plus'
import { defineStore } from 'pinia'
import { ref, computed, watch } from 'vue'
import { useRouterStore } from './router'
import cookie from 'js-cookie'

export const useUserStore = defineStore('user', () => {
  const loadingInstance = ref(null)

  const userInfo = ref({
    uuid: '',
    nickName: '',
    headerImg: '',
    authority: {},
    sideMode: 'dark',
    activeColor: 'var(--el-color-primary)',
    baseColor: '#fff'
  })
  const token = ref(window.localStorage.getItem('token') || cookie.get('x-token') || '')
  const setUserInfo = (val) => {
    userInfo.value = val
  }

  const setToken = (val) => {
    token.value = val
  }

  const NeedInit = () => {
    token.value = ''
    window.localStorage.removeItem('token')
    localStorage.clear()
    router.push({ name: 'Init', replace: true })
  }

  const ResetUserInfo = (value = {}) => {
    userInfo.value = {
      ...userInfo.value,
      ...value
    }
  }
  /* Get user information */
  const GetUserInfo = async() => {
    const res = await getUserInfo()
    if (res.code === 0) {
      setUserInfo(res.data.userInfo)
    }
    return res
  }
  /* Log in*/
  const LoginIn = async(loginInfo) => {
    loadingInstance.value = ElLoading.service({
      fullscreen: true,
      text: 'Logging in, please wait...',
    })
    try {
      const res = await login(loginInfo)
      if (res.code === 0) {
        setUserInfo(res.data.user)
        setToken(res.data.token)
        const routerStore = useRouterStore()
        await routerStore.SetAsyncRouter()
        const asyncRouters = routerStore.asyncRouters
        asyncRouters.forEach(asyncRouter => {
          router.addRoute(asyncRouter)
        })

        if (!router.hasRoute(userInfo.value.authority.defaultRouter)) {
          ElMessage.error('Please contact the administrator for authorization')
        } else {
          await router.replace({ name: userInfo.value.authority.defaultRouter })
        }

        loadingInstance.value.close()

        const isWin = ref(/windows/i.test(navigator.userAgent))
        if (isWin.value) {
          window.localStorage.setItem('osType', 'WIN')
        } else {
          window.localStorage.setItem('osType', 'MAC')
        }
        return true
      }
    } catch (e) {
      loadingInstance.value.close()
    }
    loadingInstance.value.close()
  }
  /* Sign out*/
  const LoginOut = async() => {
    const res = await jsonInBlacklist()
    if (res.code === 0) {
      await ClearStorage()
      router.push({ name: 'Login', replace: true })
      window.location.reload()
    }
  }
  /* Clean data */
  const ClearStorage = async() => {
    token.value = ''
    sessionStorage.clear()
    localStorage.clear()
    cookie.remove('x-token')
  }
  /*Set sidebar mode*/
  const changeSideMode = async(data) => {
    const res = await setSelfInfo({ sideMode: data })
    if (res.code === 0) {
      userInfo.value.sideMode = data
      ElMessage({
        type: 'success',
        message: 'Set successfully'
      })
    }
  }

  const mode = computed(() => userInfo.value.sideMode)
  const sideMode = computed(() => {
    if (userInfo.value.sideMode === 'dark') {
      return '#191a23'
    } else if (userInfo.value.sideMode === 'light') {
      return '#fff'
    } else {
      return userInfo.value.sideMode
    }
  })
  const baseColor = computed(() => {
    if (userInfo.value.sideMode === 'dark') {
      return '#fff'
    } else if (userInfo.value.sideMode === 'light') {
      return '#191a23'
    } else {
      return userInfo.value.baseColor
    }
  })
  const activeColor = computed(() => {
    return 'var(--el-color-primary)'
  })

  watch(() => token.value, () => {
    window.localStorage.setItem('token', token.value)
  })

  return {
    userInfo,
    token,
    NeedInit,
    ResetUserInfo,
    GetUserInfo,
    LoginIn,
    LoginOut,
    changeSideMode,
    mode,
    sideMode,
    setToken,
    baseColor,
    activeColor,
    loadingInstance,
    ClearStorage
  }
})
