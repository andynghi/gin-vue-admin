import config from './config'
import { h } from 'vue'

// Uniformly import the el-icon icon
import * as ElIconModules from '@element-plus/icons-vue'
import svgIcon from '@/components/svgIcon/svgIcon.vue'
//Import the function that converts the icon name

const createIconComponent = (name) => ({
  name: 'SvgIcon',
  render() {
    return h(svgIcon, {
      name: name,
    })
  },
})

const registerIcons = async(app) => {
  const iconModules = import.meta.glob('@/assets/icons/**/*.svg')
  for (const path in iconModules) {
    const iconName = path.split('/').pop().replace(/\.svg$/, '')
    // If iconName contains spaces, it will not be added to the icon library and will prompt that the name is illegal.
    console.log(iconName)
    if (iconName.indexOf(' ') !== -1) {
      console.error(`icon ${iconName}.svg includes whitespace`)
      continue
    }
    const iconComponent = createIconComponent(iconName)
    config.logs.push({
      'key': iconName,
      'label': iconName,
    })
    app.component(iconName, iconComponent)
  }
}

export const register = (app) => {
  // Unified registration of el-icon icon
  for (const iconName in ElIconModules) {
    app.component(iconName, ElIconModules[iconName])
  }
  app.component('SvgIcon', svgIcon)
  registerIcons(app)
  app.config.globalProperties.$GIN_VUE_ADMIN = config
}
