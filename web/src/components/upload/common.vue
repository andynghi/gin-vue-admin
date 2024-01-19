<template>
  <div>
    <el-upload
      :action="`${path}/fileUploadAndDownload/upload`"
      :before-upload="checkFile"
      :on-error="uploadError"
      :on-success="uploadSuccess"
      :show-file-list="false"
      class="upload-btn"
    >
      <el-button type="primary">Normal upload</el-button>
    </el-upload>
  </div>
</template>

<script setup>

import { ref } from 'vue'
import { ElMessage } from 'element-plus'
import { useUserStore } from '@/pinia/modules/user'
import { isVideoMime, isImageMime } from '@/utils/image'

defineOptions({
  name: 'UploadCommon',
})

const emit = defineEmits(['on-success'])
const path = ref(import.meta.env.VITE_BASE_API)

const userStore = useUserStore()
const fullscreenLoading = ref(false)

const checkFile = (file) => {
  fullscreenLoading.value = true
  const isLt500K = file.size / 1024 / 1024 < 0.5 // 500K, @todo should support setting in the project
  const isLt5M = file.size / 1024 / 1024 < 5 // 5MB, @todo should support the setting in the project
  const isVideo = isVideoMime(file.type)
  const isImage = isImageMime(file.type)
  let pass = true
  if (!isVideo && !isImage) {
    ElMessage.error('Uploaded pictures can only be in jpg, png, svg, webp format, uploaded videos can only be in mp4, webm format!')
    fullscreenLoading.value = false
    pass = false
  }
  if (!isLt5M && isVideo) {
    ElMessage.error('The uploaded video size cannot exceed 5MB')
    fullscreenLoading.value = false
    pass = false
  }
  if (!isLt500K && isImage) {
    ElMessage.error('The size of uncompressed uploaded images cannot exceed 500KB, please use compression to upload')
    fullscreenLoading.value = false
    pass = false
  }

  console.log('upload file check result: ', pass)

  return pass
}

const uploadSuccess = (res) => {
  const { data } = res
  if (data.file) {
    emit('on-success', data.file.url)
  }
}

const uploadError = () => {
  ElMessage({
    type: 'error',
    message: 'Upload failed'
  })
  fullscreenLoading.value = false
}

</script>

