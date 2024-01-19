
<template>
  <div>
    <el-upload
      :action="`${path}/fileUploadAndDownload/upload`"
      :show-file-list="false"
      :on-success="handleImageSuccess"
      :before-upload="beforeImageUpload"
      :multiple="false"
    >
      <el-button type="primary">Compressed upload</el-button>
    </el-upload>
  </div>
</template>

<script setup>
import ImageCompress from '@/utils/image'
import { ref } from 'vue'
import { ElMessage } from 'element-plus'
import { useUserStore } from '@/pinia/modules/user'

defineOptions({
  name: 'UploadImage',
})

const emit = defineEmits(['on-success'])
const props = defineProps({
  imageUrl: {
    type: String,
    default: ''
  },
  fileSize: {
    type: Number,
    default: 2048 // Perform compression after exceeding 2M
  },
  maxWH: {
    type: Number,
    default: 1920 //Maximum image length and width
  }
})

const path = ref(import.meta.env.VITE_BASE_API)

const userStore = useUserStore()

const beforeImageUpload = (file) => {
  const isJPG = file.type === 'image/jpeg'
  const isPng = file.type === 'image/png'
  if (!isJPG && !isPng) {
    ElMessage.error('Uploaded avatar pictures can only be in jpg or png format!')
    return false
  }

  const isRightSize = file.size / 1024 < props.fileSize
  if (!isRightSize) {
    // Compression
    const compress = new ImageCompress(file, props.fileSize, props.maxWH)
    return compress.compress()
  }
  return isRightSize
}

const handleImageSuccess = (res) => {
  const { data } = res
  if (data.file) {
    emit('on-success', data.file.url)
  }
}

</script>

<style lang="scss" scoped>
.image-uploader {
  border: 1px dashed #d9d9d9;
  width: 180px;
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
}
.image-uploader {
  border-color: #409eff;
}
.image-uploader-icon {
  font-size: 28px;
  color: #8c939d;
  width: 178px;
  height: 178px;
  line-height: 178px;
  text-align: center;
}
.image {
  width: 178px;
  height: 178px;
  display: block;
}
</style>
