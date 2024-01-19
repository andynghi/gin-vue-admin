<template>
  <div id="app">
    <el-config-provider :locale="zhCn">
      <router-view />
    </el-config-provider>
  </div>
</template>

<script setup>
import zhCn from 'element-plus/dist/locale/zh-cn.mjs'
// Use the following statement before element 2.3.8
// import zhCn from 'element-plus/lib/locale/lang/zh-cn'

defineOptions({
  name: 'App'
})

</script>
<style lang="scss">
@tailwind base;
@tailwind components;
@tailwind utilities;
//Introduce initialization style
#app {
  background: #eee;
  height: 100vh;
  overflow: hidden;
  font-weight: 400 !important;
}
.el-button{
  font-weight: 400 !important;
}
</style>
