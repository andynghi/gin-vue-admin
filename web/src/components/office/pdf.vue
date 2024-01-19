<template>
  <vue-office-pdf

      :src="pdf"
      @rendered="renderedHandler"
      @error="errorHandler"
  />
</template>
<script>
export default {
  name: "Pdf"
}
</script>
<script setup>
import {ref, watch} from "vue"

//Introduce VueOfficeDocx component
import VueOfficePdf from "@vue-office/pdf";
//Introduce related styles
import '@vue-office/docx/lib/index.css'
console.log("pdf===>")
const props = defineProps({
  modelValue: {
    type: String,
    default: () => ""
  }
})
const pdf = ref(null)
watch(() => props.modelValue, val => pdf.value = val, {immediate: true})
const renderedHandler = () => {
  console.log("pdf loaded successfully")
}
const errorHandler = () => {
  console.log("pdf error")
}
</script>