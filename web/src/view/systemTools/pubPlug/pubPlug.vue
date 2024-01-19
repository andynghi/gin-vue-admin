<template>
  <div class="p-5 bg-white">
    <WarningBar title="Currently only supports standard plug-ins (standard directory plug-ins generated through plug-in templates). Please package non-standard plug-ins by yourself" />
    <div class="flex items-center gap-3">
      <el-input
        v-model="plugName"
        placeholder="[Plug-in name] filled in the plug-in template"
      />
      <el-button
        type="primary"
        @click="pubPlugin"
      >Packaging plug-in</el-button>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import WarningBar from '@/components/warningBar/warningBar.vue'
import { pubPlug } from '@/api/autoCode.js'
import { ElMessage } from 'element-plus'
const plugName = ref('')

const pubPlugin = async() => {
  const res = await pubPlug({ plugName: plugName.value })
  if (res.code === 0) {
    ElMessage.success(res.msg)
  }
}

</script>

