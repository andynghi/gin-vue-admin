<template>
  <div>
    <warning-bar title="id, created_at, updated_at, deleted_at will be automatically generated and please do not create it again. When searching, if the condition is LIKE, only strings are supported" />
    <el-form
      ref="fieldDialogFrom"
      :model="middleDate"
      label-width="120px"
      label-position="right"
      :rules="rules"
      class="grid grid-cols-2"
    >
      <el-form-item
        label="field name"
        prop="fieldName"
      >
        <el-input
          v-model="middleDate.fieldName"
          autocomplete="off"
          style="width:80%"
        />
        <el-button
          style="width:18%;margin-left:2%"
          @click="autoFill"
        >
          <span style="font-size: 12px">Autofill</span>
        </el-button>
      </el-form-item>
      <el-form-item
        label="Field Chinese name"
        prop="fieldDesc"
      >
        <el-input
          v-model="middleDate.fieldDesc"
          autocomplete="off"
        />
      </el-form-item>
      <el-form-item
        label="Field JSON"
        prop="fieldJson"
      >
        <el-input
          v-model="middleDate.fieldJson"
          autocomplete="off"
        />
      </el-form-item>
      <el-form-item
        label="database field name"
        prop="columnName"
      >
        <el-input
          v-model="middleDate.columnName"
          autocomplete="off"
        />
      </el-form-item>
      <el-form-item
        label="Database field description"
        prop="comment"
      >
        <el-input
          v-model="middleDate.comment"
          autocomplete="off"
        />
      </el-form-item>
      <el-form-item
        label="field type"
        prop="fieldType"
      >
        <el-select
          v-model="middleDate.fieldType"
          style="width:100%"
          placeholder="Please select field type"
          clearable
          @change="clearOther"
        >
          <el-option
            v-for="item in typeOptions"
            :key="item.value"
            :label="item.label"
            :value="item.value"
            :disabled="item.disabled"
          />
        </el-select>
      </el-form-item>
      <el-form-item
        :label="middleDate.fieldType === 'enum' ? 'enumeration value' : 'type length'"
        prop="dataTypeLong"
      >
        <el-input
          v-model="middleDate.dataTypeLong"
          :placeholder="middleDate.fieldType === 'enum'?`Example:'Beijing','Tianjin'`:'Database type length'"
        />
      </el-form-item>
      <el-form-item
        label="field query conditions"
        prop="fieldSearchType"
      >
        <el-select
          v-model="middleDate.fieldSearchType"
          style="width:100%"
          placeholder="Please select field query conditions"
          clearable
        >
          <el-option
            v-for="item in typeSearchOptions"
            :key="item.value"
            :label="item.label"
            :value="item.value"
            :disabled="canSelect(item.value)"
          />
        </el-select>
      </el-form-item>
      <el-form-item
        label="associated dictionary"
        prop="dictType"
      >
        <el-select
          v-model="middleDate.dictType"
          style="width:100%"
          :disabled="middleDate.fieldType!=='int'"
          placeholder="Please select a dictionary"
          clearable
        >
          <el-option
            v-for="item in dictOptions"
            :key="item.type"
            :label="`${item.type}(${item.name})`"
            :value="item.type"
          />
        </el-select>
      </el-form-item>
      <el-form-item label="Whether to sort">
        <el-switch v-model="middleDate.sort" />
      </el-form-item>
      <el-form-item label="Is it required">
        <el-switch v-model="middleDate.require" />
      </el-form-item>
      <el-form-item label="Can it be cleared">
        <el-switch v-model="middleDate.clearable" />
      </el-form-item>
      <el-form-item label="Verification failure copy">
        <el-input v-model="middleDate.errorText" />
      </el-form-item>

    </el-form>
  </div>
</template>

<script setup>
import { toLowerCase, toSQLLine } from '@/utils/stringFun'
import { getSysDictionaryList } from '@/api/sysDictionary'
import WarningBar from '@/components/warningBar/warningBar.vue'
import { ref } from 'vue'

defineOptions({
  name: 'FieldDialog'
})

const props = defineProps({
  dialogMiddle: {
    type: Object,
    default: function() {
      return {}
    }
  },
  typeOptions: {
    type: Array,
    default: function() {
      return []
    }
  },
  typeSearchOptions: {
    type: Array,
    default: function() {
      return []
    }
  },
})

const middleDate = ref({})
const dictOptions = ref([])

const rules = ref({
  fieldName: [
    { required: true, message: 'Please enter the English name of the field', trigger: 'blur' }
  ],
  fieldDesc: [
    { required: true, message: 'Please enter the Chinese name of the field', trigger: 'blur' }
  ],
  fieldJson: [
    { required: true, message: 'Please enter the field to format json', trigger: 'blur' }
  ],
  columnName: [
    { required: true, message: 'Please enter database fields', trigger: 'blur' }
  ],
  fieldType: [
    { required: true, message: 'Please select field type', trigger: 'blur' }
  ]
})

const init = async() => {
  middleDate.value = props.dialogMiddle
  const dictRes = await getSysDictionaryList({
    page: 1,
    pageSize: 999999
  })

  dictOptions.value = dictRes.data
}
init()

const autoFill = () => {
  middleDate.value.fieldJson = toLowerCase(middleDate.value.fieldName)
  middleDate.value.columnName = toSQLLine(middleDate.value.fieldJson)
}

const canSelect = (item) => {
  const fieldType = middleDate.value.fieldType
  if (fieldType !== 'string' && item === 'LIKE') {
    return true
  }

  if ((fieldType !== 'int' && fieldType !== 'time.Time' && fieldType !== 'float64') && (item === 'BETWEEN' || item === 'NOT BETWEEN')) {
    return true
  }
  return false
}

const clearOther = () => {
  middleDate.value.fieldSearchType = ''
  middleDate.value.dictType = ''
}

const fieldDialogFrom = ref(null)
defineExpose({ fieldDialogFrom })
</script>
