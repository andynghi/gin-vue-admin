<template>
  <div>
    <warning-bar
      title="This function is only used to create roles and many2many relationship tables. For specific use, you must combine the tables to implement business. For details, please refer to the sample code (customer example). This function is not recommended. It is recommended to use the plug-in market [Organization Management Function (Click to go)] to manage resource permissions."
      href="https://plugin.gin-vue-admin.com/#/layout/newPluginInfo?id=36"
    />
    <div class="sticky top-0.5 z-10 bg-white my-4">
      <el-button
        class="float-left"
        type="primary"
        @click="all"
      >Select all</el-button>
      <el-button
        class="float-left"
        type="primary"
        @click="self"
      >This role</el-button>
      <el-button
        class="float-left"
        type="primary"
        @click="selfAndChildren"
      >This role and sub-roles</el-button>
      <el-button
        class="float-right"
        type="primary"
        @click="authDataEnter"
      >Confirm</el-button>
    </div>
    <div class="clear-both pt-4">
      <el-checkbox-group
        v-model="dataAuthorityId"
        @change="selectAuthority"
      >
        <el-checkbox
          v-for="(item,key) in authoritys"
          :key="key"
          :label="item"
        >{{ item.authorityName }}</el-checkbox>
      </el-checkbox-group>
    </div>
  </div>
</template>

<script setup>
import { setDataAuthority } from '@/api/authority'
import WarningBar from '@/components/warningBar/warningBar.vue'
import { ref } from 'vue'
import { ElMessage } from 'element-plus'

defineOptions({
  name: 'Datas'
})

const props = defineProps({
  row: {
    default: function() {
      return {}
    },
    type: Object
  },
  authority: {
    default: function() {
      return []
    },
    type: Array
  }
})

const authoritys = ref([])
const needConfirm = ref(false)
//Tile the character
const roundAuthority = (authoritysData) => {
  authoritysData && authoritysData.forEach(item => {
    const obj = {}
    obj.authorityId = item.authorityId
    obj.authorityName = item.authorityName
    authoritys.value.push(obj)
    if (item.children && item.children.length) {
      roundAuthority(item.children)
    }
  })
}

const dataAuthorityId = ref([])
const init = () => {
  roundAuthority(props.authority)
  props.row.dataAuthorityId && props.row.dataAuthorityId.forEach(item => {
    const obj = authoritys.value && authoritys.value.filter(au => au.authorityId === item.authorityId) && authoritys.value.filter(au => au.authorityId === item.authorityId)[0]
    dataAuthorityId.value.push(obj)
  })
}

init()

// Unified switching interception method exposed to the outer layer
const enterAndNext = () => {
  authDataEnter()
}

const emit = defineEmits(['changeRow'])
const all = () => {
  dataAuthorityId.value = [...authoritys.value]
  emit('changeRow', 'dataAuthorityId', dataAuthorityId.value)
  needConfirm.value = true
}
const self = () => {
  dataAuthorityId.value = authoritys.value.filter(item => item.authorityId === props.row.authorityId)
  emit('changeRow', 'dataAuthorityId', dataAuthorityId.value)
  needConfirm.value = true
}
const selfAndChildren = () => {
  const arrBox = []
  getChildrenId(props.row, arrBox)
  dataAuthorityId.value = authoritys.value.filter(item => arrBox.indexOf(item.authorityId) > -1)
  emit('changeRow', 'dataAuthorityId', dataAuthorityId.value)
  needConfirm.value = true
}
const getChildrenId = (row, arrBox) => {
  arrBox.push(row.authorityId)
  row.children && row.children.forEach(item => {
    getChildrenId(item, arrBox)
  })
}
// submit
const authDataEnter = async() => {
  const res = await setDataAuthority(props.row)
  if (res.code === 0) {
    ElMessage({ type: 'success', message: 'Resource setting successful' })
  }
}

//   choose
const selectAuthority = () => {
  emit('changeRow', 'dataAuthorityId', dataAuthorityId.value)
  needConfirm.value = true
}

defineExpose({
  enterAndNext,
  needConfirm
})

</script>
