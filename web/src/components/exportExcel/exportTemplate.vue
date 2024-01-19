<template>
  <el-button
    type="primary"
    icon="download"
    @click="exportTemplate"
  >Download template</el-button>
</template>

<script setup>
const props = defineProps({
  templateId: {
    type: String,
    required: true
  }
})

import { ElMessage } from 'element-plus'

const exportTemplate = async() => {
  if (props.templateId === '') {
    ElMessage.error('The component does not set the template ID')
    return
  }
  const baseUrl = import.meta.env.VITE_BASE_API
  const url = `${baseUrl}/sysExportTemplate/exportTemplate?templateID=${props.templateId}`
  window.open(url, '_blank')
}
</script>
