<template>
  <div>
    <div class="gva-table-box">
      <el-form
        label-width="140px"
        class="w-[680px]"
      >
        <el-form-item label="plugin name">
          <el-input
            v-model="form.plugName"
            placeholder="required (starting with an English capital letter)"
            @blur="titleCase"
          />
        </el-form-item>
        <el-form-item label="routing group">
          <el-input
            v-model="form.routerGroup"
            placeholder="will be used as a plugin routing group"
          />
        </el-form-item>
        <el-form-item label="Use global attributes">
          <el-checkbox v-model="form.hasGlobal" />
        </el-form-item>
        <el-form-item
          v-if="form.hasGlobal"
          label="global properties"
        >
          <div
            v-for="(i,k) in form.global"
            :key="k"
            class="plug-row"
          >
            <span>
              <el-input
                v-model="i.key"
                placeholder="key required"
              />
            </span>
            <span>
              <el-select
                v-model="i.type"
                placeholder="type required"
              >
                <el-option
                  label="string"
                  value="string"
                />
                <el-option
                  label="int"
                  value="int"
                />
                <el-option
                  label="float32"
                  value="float32"
                />
                <el-option
                  label="float64"
                  value="float64"
                />
                <el-option
                  label="bool"
                  value="bool"
                />
              </el-select>
            </span>
            <span>
              <el-input
                v-model="i.desc"
                placeholder="Remarks required"
              />
            </span>
            <span>
              <el-button
                :icon="Plus"
                circle
                @click="addkv(form.global)"
              />
            </span>
            <span>
              <el-button
                :icon="Minus"
                circle
                @click="minkv(form.global,k)"
              />
            </span>
          </div>
        </el-form-item>
        <el-form-item label="Use Request">
          <el-checkbox v-model="form.hasRequest" />
        </el-form-item>
        <el-form-item
          v-if="form.hasRequest"
          label="Request"
        >
          <div
            v-for="(i,k) in form.request"
            :key="k"
            class="plug-row"
          >
            <span>
              <el-input
                v-model="i.key"
                placeholder="key required"
              />
            </span>
            <span>
              <el-select
                v-model="i.type"
                placeholder="type required"
              >
                <el-option
                  label="string"
                  value="string"
                />
                <el-option
                  label="int"
                  value="int"
                />
                <el-option
                  label="float32"
                  value="float32"
                />
                <el-option
                  label="float64"
                  value="float64"
                />
                <el-option
                  label="bool"
                  value="bool"
                />
              </el-select>
            </span>
            <span>
              <el-input
                v-model="i.desc"
                placeholder="Remarks required"
              />
            </span>
            <span>
              <el-button
                :icon="Plus"
                circle
                @click="addkv(form.request)"
              />
            </span>
            <span>
              <el-button
                :icon="Minus"
                circle
                @click="minkv(form.request,k)"
              />
            </span>
          </div>
        </el-form-item>
        <el-form-item label="Use Response">
          <el-checkbox v-model="form.hasResponse" />
        </el-form-item>
        <el-form-item
          v-if="form.hasResponse"
          label="Response"
        >
          <div
            v-for="(i,k) in form.response"
            :key="k"
            class="plug-row"
          >
            <span>
              <el-input
                v-model="i.key"
                placeholder="key required"
              />
            </span>
            <span>
              <el-select
                v-model="i.type"
                placeholder="type required"
              >
                <el-option
                  label="string"
                  value="string"
                />
                <el-option
                  label="int"
                  value="int"
                />
                <el-option
                  label="float32"
                  value="float32"
                />
                <el-option
                  label="float64"
                  value="float64"
                />
                <el-option
                  label="bool"
                  value="bool"
                />
              </el-select>
            </span>
            <span>
              <el-input
                v-model="i.desc"
                placeholder="Remarks required"
              />
            </span>
            <span>
              <el-button
                :icon="Plus"
                circle
                @click="addkv(form.response)"
              />
            </span>
            <span>
              <el-button
                :icon="Minus"
                circle
                @click="minkv(form.response,k)"
              />
            </span>
          </div>
        </el-form-item>
        <el-form-item>
          <el-button
            type="primary"
            @click="createPlug"
          >Create</el-button>
          <el-icon
            class="cursor-pointer ml-3"
            @click="toDoc('https://www.bilibili.com/video/BV1kv4y1g7nT?p=13&vd_source=f2640257c21e3b547a790461ed94875e')"
          ><VideoCameraFilled /></el-icon>
        </el-form-item>
      </el-form>

    </div>
  </div>
</template>

<script setup>
import { toUpperCase } from '@/utils/stringFun'

import {
  Plus,
  Minus, VideoCameraFilled
} from '@element-plus/icons-vue'
import { ElMessage, ElMessageBox } from 'element-plus'

import { createPlugApi } from '@/api/autoCode.js'

import { reactive } from 'vue'
import { toDoc } from '@/utils/doc'

const form = reactive({
  plugName: '',
  routerGroup: '',
  hasGlobal: true,
  hasRequest: true,
  hasResponse: true,
  global: [{
    key: '',
    type: '',
    desc: '',
  }],
  request: [{
    key: '',
    type: '',
    desc: '',
  }],
  response: [{
    key: '',
    type: '',
    desc: '',
  }]
})

const titleCase = () => {
  form.plugName = toUpperCase(form.plugName)
}

const createPlug = async() => {
  if (!form.plugName || !form.routerGroup) {
    ElMessage.error('Plug-in name and plug-in routing group are required')
    return
  }
  if (form.hasGlobal) {
    const intercept = form.global.some(i => {
      if (!i.key || !i.type) {
        return true
      }
    })
    if (intercept) {
      ElMessage.error('The key and type of global attributes are required')
      return
    }
  }
  if (form.hasRequest) {
    const intercept = form.request.some(i => {
      if (!i.key || !i.type) {
        return true
      }
    })
    if (intercept) {
      ElMessage.error('The key and type of the request attribute are required')
      return
    }
  }
  if (form.hasResponse) {
    const intercept = form.response.some(i => {
      if (!i.key || !i.type) {
        return true
      }
    })
    if (intercept) {
      ElMessage.error('The key and type of the response attribute are required')
      return
    }
  }
  const res = await createPlugApi(form)
  if (res.code === 0) {
    ElMessageBox('Created successfully, the plug-in has been automatically written to the back-end plugin directory, please create it according to your own logic')
  }
}

const addkv = (arr) => {
  arr.push({
    key: '',
    value: '',
  })
}

const minkv = (arr, key) => {
  if (arr.length === 1) {
    ElMessage.warning('At least one global property')
    return
  }
  arr.splice(key, 1)
}

</script>

<style lang="scss" scoped>
    .plug-row{
      @apply flex items-center w-full;
        &+&{
          @apply mt-3;
        }
        &>span{
          @apply ml-2;
        }
    }
</style>
