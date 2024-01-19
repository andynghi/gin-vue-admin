<template>
  <vue-office-docx :src="docx" @rendered="rendered">

  </vue-office-docx>
</template>
<script>
export default {
  name: "Docx"
}
</script>
<script setup>
import {ref, watch} from 'vue'

//Introduce VueOfficeDocx component
import VueOfficeDocx from '@vue-office/docx'
//Introduce related styles
import '@vue-office/docx/lib/index.css'

const props = defineProps({
  modelValue: {
    type: String,
    default: () => ""
  }
})
const docx = ref(null)
watch(
    () => props.modelValue,
    value => docx.value = value,
    {immediate: true}
)
const rendered = () => {
}
</script>
<style lang="scss" scoped>

</style>