<template>
  <div>
    <warning-bar title="You need to configure the email configuration file in advance. To prevent unnecessary spam, the online experience function does not allow this feature to be experienced." />
    <div class="gva-form-box">
      <el-form
        ref="emailForm"
        label-position="right"
        label-width="80px"
        :model="form"
      >
        <el-form-item label="target email">
          <el-input v-model="form.to" />
        </el-form-item>
        <el-form-item label="mail">
          <el-input v-model="form.subject" />
        </el-form-item>
        <el-form-item label="Email content">
          <el-input
            v-model="form.body"
            type="textarea"
          />
        </el-form-item>
        <el-form-item>
          <el-button @click="sendTestEmail">Send test email</el-button>
          <el-button @click="sendEmail">Send email</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>

</template>

<script setup>
import WarningBar from '@/components/warningBar/warningBar.vue'
import { emailTest } from '@/plugin/email/api/email.js'
import { ElMessage } from 'element-plus'
import { reactive, ref } from 'vue'

defineOptions({
  name: 'Email',
})

const emailForm = ref(null)
const form = reactive({
  to: '',
  subject: '',
  body: '',
})
const sendTestEmail = async() => {
  const res = await emailTest()
  if (res.code === 0) {
    ElMessage.success('sent successfully')
  }
}

const sendEmail = async() => {
  const res = await emailTest()
  if (res.code === 0) {
    ElMessage.success('Sent successfully, please check')
  }
}
</script>

