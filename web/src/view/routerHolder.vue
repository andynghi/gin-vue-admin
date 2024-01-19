<!-- This route can be used as a general routing page for parent routing. If you want to customize the parent routing page, please refer to @/view/superAdmin/index.vue -->
<template>
  <div>
    <router-view v-slot="{ Component }">
      <transition
        mode="out-in"
        name="el-fade-in-linear"
      >
        <keep-alive :include="routerStore.keepAliveRouters">
          <div>
            <component :is="Component" />
          </div>
        </keep-alive>
      </transition>
    </router-view>
  </div>
</template>

<script setup>
defineOptions({
  name: 'RouterHolder'
})
import { useRouterStore } from '@/pinia/modules/router'
const routerStore = useRouterStore()
</script>
