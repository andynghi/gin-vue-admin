<template>
  <div>
    <div class="w-full h-screen bg-gray-50 flex items-center justify-center">
      <div class="flex flex-col items-center text-2xl gap-4">
        <img src="../../assets/notFound.png">
        <p>The page was sucked away by a mysterious force (if you are using the open source version, please contact us to fix it)</p>
        <p style="font-size:18px;line-height:40px;">The common problem is that there is no current route for this role. If you are sure you want to use this route, please go to role management to assign it</p>
        <p>↓</p>
        <img
          src="../../assets/qm.png"
          class="w-16 h-16 mt-20"
        >
        <el-button @click="toDashboard">Return to home page</el-button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { useUserStore } from '@/pinia/modules/user'
import { useRouter } from 'vue-router'

defineOptions({
  name: 'Error'
})

const userStore = useUserStore()
const router = useRouter()
const toDashboard = () => {
  router.push({ name: userStore.userInfo.authority.defaultRouter })
}
</script>
