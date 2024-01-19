import 'element-plus/es/components/message/style/css'
import 'element-plus/es/components/loading/style/css'
import 'element-plus/es/components/notification/style/css'
import 'element-plus/es/components/message-box/style/css'
import './style/element_visiable.scss'
import { createApp } from 'vue'
//Introduce gin-vue-admin front-end initialization related content
import './core/gin-vue-admin'
//Introduce the encapsulated router
import router from '@/router/index'
import '@/permission'
import run from '@/core/gin-vue-admin.js'
import auth from '@/directive/auth'
import { store } from '@/pinia'
import App from './App.vue'
import { initDom } from './utils/positionToCode'

initDom()
/**
 * @description Import the loading progress bar to prevent the first screen from loading too long and the user waiting
 *
 * */
import Nprogress from 'nprogress'
import 'nprogress/nprogress.css'
Nprogress.configure({ showSpinner: false, ease: 'ease', speed: 500 })
Nprogress.start()

/**
 * There is no need to end this block, this block will be ended in the routing middleware.
 * */

const app = createApp(App)
app.config.productionTip = false

app
  .use(run)
  .use(store)
  .use(auth)
  .use(router)
  .mount('#app')

export default app
