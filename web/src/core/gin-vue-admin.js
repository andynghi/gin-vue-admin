/*
 * gin-vue-admin web framework group
 *
 * */
//Load website configuration folder
import { register } from './global'

export default {
  install: (app) => {
    register(app)
    console.log(`
       Welcome to Gin-Vue-Admin
       Current version: v2.5.9
       How to join the group: WeChat: shouzi_1994 QQ group: 622360840
       GVA discussion community: https://support.qq.com/products/371961
       Plug-in market: https://plugin.gin-vue-admin.com
       Default automation document address: http://127.0.0.1:${import.meta.env.VITE_SERVER_PORT}/swagger/index.html
       Default front-end file running address: http://127.0.0.1:${import.meta.env.VITE_CLI_PORT}
       If the project has made you gain, I hope you can treat the team to a cup of Coke: https://www.gin-vue-admin.com/coffee/index.html
    `)
  }
}
