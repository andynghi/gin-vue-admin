import axios from 'axios' //Introduce axios
import { ElMessage, ElMessageBox } from 'element-plus'
import { useUserStore } from '@/pinia/modules/user'
import { emitter } from '@/utils/bus.js'
import router from '@/router/index'

const service = axios.create({
  baseURL: import.meta.env.VITE_BASE_API,
  timeout: 99999
})
let acitveAxios = 0
let timer
const showLoading = () => {
  acitveAxios++
  if (timer) {
    clearTimeout(timer)
  }
  timer = setTimeout(() => {
    if (acitveAxios > 0) {
      emitter.emit('showLoading')
    }
  }, 400)
}

const closeLoading = () => {
  acitveAxios--
  if (acitveAxios <= 0) {
    clearTimeout(timer)
    emitter.emit('closeLoading')
  }
}
// http request interceptor
service.interceptors.request.use(
  config => {
    if (!config.donNotShowLoading) {
      showLoading()
    }
    const userStore = useUserStore()
    config.headers = {
      'Content-Type': 'application/json',
      'x-token': userStore.token,
      'x-user-id': userStore.userInfo.ID,
      ...config.headers
    }
    return config
  },
  error => {
    if (!error.config.donNotShowLoading) {
      closeLoading()
    }
    ElMessage({
      showClose: true,
      message: error,
      type: 'error'
    })
    return error
  }
)

// http response interceptor
service.interceptors.response.use(
  response => {
    const userStore = useUserStore()
    if (!response.config.donNotShowLoading) {
      closeLoading()
    }
    if (response.headers['new-token']) {
      userStore.setToken(response.headers['new-token'])
    }
    if (response.data.code === 0 || response.headers.success === 'true') {
      if (response.headers.msg) {
        response.data.msg = decodeURI(response.headers.msg)
      }
      return response.data
    } else {
      ElMessage({
        showClose: true,
        message: response.data.msg || decodeURI(response.headers.msg),
        type: 'error'
      })
      if (response.data.data && response.data.data.reload) {
        userStore.token = ''
        localStorage.clear()
        router.push({ name: 'Login', replace: true })
      }
      return response.data.msg ? response.data : response
    }
  },
  error => {
    if (!error.config.donNotShowLoading) {
      closeLoading()
    }

    if (!error.response) {
      ElMessageBox.confirm(`
        <p>Request error detected</p>
        <p>${error}</p>
        `, 'Request error', {
        dangerouslyUseHTMLString: true,
        distinguishCancelAndClose: true,
        confirmButtonText: 'Try again later',
        cancelButtonText: 'Cancel'
      })
      return
    }

    switch (error.response.status) {
      case 500:
        ElMessageBox.confirm(`
        <p>Interface error detected ${error}</p>
        <p>Error code<span style="color:red"> 500 </span>: This type of error content is common in background panics. Please check the background log first. If it affects your normal use, you can force logout to clear the cache</p >
        `, 'Interface error', {
          dangerouslyUseHTMLString: true,
          distinguishCancelAndClose: true,
          confirmButtonText: 'Clear cache',
          cancelButtonText: 'Cancel'
        })
          .then(() => {
            const userStore = useUserStore()
            userStore.ClearStorage()
            router.push({ name: 'Login', replace: true })
          })
        break
      case 404:
        ElMessageBox.confirm(`
          <p>Interface error detected ${error}</p>
          <p>Error code<span style="color:red"> 404 </span>: This type of error is mostly caused by the interface not being registered (or not restarted) or the request path (method) does not match the api path (method) - if For automation code please check if there are spaces</p>
          `, 'Interface error', {
          dangerouslyUseHTMLString: true,
          distinguishCancelAndClose: true,
          confirmButtonText: 'I understand',
          cancelButtonText: 'Cancel'
        })
        break
    }

    return error
  }
)
export default service
