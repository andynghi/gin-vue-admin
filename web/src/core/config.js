/**
 * Website configuration file
 */

const config = {
  appName: 'Gin-Vue-Admin',
  appLogo: 'https://www.gin-vue-admin.com/img/logo.png',
  showViteLogo: true,
  logs: [],
}

export const viteLogo = (env) => {
  if (config.showViteLogo) {
    const chalk = require('chalk')
    console.log(
      chalk.green(
        `> Welcome to use Gin-Vue-Admin, open source address: https://github.com/flipped-aurora/gin-vue-admin`
      )
    )
    console.log(
      chalk.green(
        `> Current version: v2.5.9`
      )
    )
    console.log(
      chalk.green(
        `> How to join the group: WeChat: shouzi_1994 QQ group: 622360840`
      )
    )
    console.log(
      chalk.green(
        `>GVA discussion community: https://support.qq.com/products/371961`
      )
    )
    console.log(
      chalk.green(
        `>Plug-in Market: https://plugin.gin-vue-admin.com`
      )
    )
    console.log(
      chalk.green(
        `>Default automation document address: http://127.0.0.1:${env.VITE_SERVER_PORT}/swagger/index.html`
      )
    )
    console.log(
      chalk.green(
        `>Default front-end file running address: http://127.0.0.1:${env.VITE_CLI_PORT}`
      )
    )
    console.log(
      chalk.green(
        `> If the project has made you gain, I hope you can treat the team to a cup of Coke: https://www.gin-vue-admin.com/coffee/index.html`
      )
    )
    console.log('\n')
  }
}

export default config
