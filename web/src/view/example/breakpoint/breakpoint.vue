<template>
  <div class="break-point">
    <div class="gva-table-box">
      <el-divider content-position="left">Large file upload</el-divider>
      <form
        id="fromCont"
        method="post"
      >
        <div
          class="fileUpload"
          @click="inputChange"
        >
          Select a document
          <input
            v-show="false"
            id="file"
            ref="FileInput"
            multiple="multiple"
            type="file"
            @change="choseFile"
          >
        </div>
      </form>
      <el-button
        :disabled="limitFileSize"
        type="primary"
        class="uploadBtn"
        @click="getFile"
      >Upload file</el-button>
      <div class="el-upload__tip">Please upload files no larger than 5MB</div>
      <div class="list">
        <transition
          name="list"
          tag="p"
        >
          <div
            v-if="file"
            class="list-item"
          >
            <el-icon>
              <document />
            </el-icon>
            <span>{{ file.name }}</span>
            <span class="percentage">{{ percentage }}%</span>
            <el-progress
              :show-text="false"
              :text-inside="false"
              :stroke-width="2"
              :percentage="percentage"
            />
          </div>
        </transition>
      </div>
      <div class="tips">This version is a trial version of the trial function. Style beautification and performance optimization are in progress. Upload the sliced files and the synthesized complete files to the breakpointDir folder and fileDir folder of the QMPlusserver directory respectively</div>
    </div>
  </div>

</template>

<script setup>
import SparkMD5 from 'spark-md5'
import {
  findFile,
  breakpointContinueFinish,
  removeChunk,
  breakpointContinue
} from '@/api/breakpoint'
import { ref, watch } from 'vue'
import { ElMessage } from 'element-plus'

defineOptions({
  name: 'BreakPoint'
})

const file = ref(null)
const fileMd5 = ref('')
const formDataList = ref([])
const waitUpLoad = ref([])
const waitNum = ref(NaN)
const limitFileSize = ref(false)
const percentage = ref(0)
const percentageFlage = ref(true)

//Function to select files
const choseFile = async(e) => {
  const fileR = new FileReader() // Create a reader to read the file stream
  const fileInput = e.target.files[0] // Get the current file
  const maxSize = 5 * 1024 * 1024
  file.value = fileInput // file is lost globally for later use. It can be improved to func parameter passing form.
  percentage.value = 0
  if (file.value.size < maxSize) {
    fileR.readAsArrayBuffer(file.value) //Read the file into ArrayBuffer mainly to maintain consistency with the back-end stream
    fileR.onload = async e => {
      // The callback e read as arrayBuffer is the method's own parameter, which is equivalent to the e stream of dom stored in e.target.result
      const blob = e.target.result
      const spark = new SparkMD5.ArrayBuffer() // Create md5 manufacturing tool (md5 is used to detect file consistency. If you don’t understand here, please call me and ask me)
      spark.append(blob) //File stream thrown into tool
      fileMd5.value = spark.end() // The tool ends and generates a md5 of the total file.
      const FileSliceCap = 1 * 1024 * 1024 //Number of slice bytes
      let start = 0 // Define where the sharding starts
      let end = 0 // The end of each slice is cut a
      let i = 0 // which slice
      formDataList.value = [] // A pool of sharded storage, losing the global
      while (end < file.value.size) {
        // When the ending number is greater than the total size of the file, end slicing
        start = i * FileSliceCap // Calculate the starting position of each slice
        end = (i + 1) * FileSliceCap // Calculate the end position of each slice
        var fileSlice = file.value.slice(start, end) // Start slicing file.slice is the h5 method to slice the file. The parameters are the starting and ending bytes.
        const formData = new window.FormData() // Create FormData to store information passed to the backend
        formData.append('fileMd5', fileMd5.value) // Store the Md5 of the total file to let the backend know whose slice it is.
        formData.append('file', fileSlice) //Current slice
        formData.append('chunkNumber', i) // Which piece is the current one?
        formData.append('fileName', file.value.name) // The file name of the current file. The naming used for back-end file slices. formData.appen is a method for adding parameters to the formData object.
        formDataList.value.push({ key: i, formData }) // Store the current slice information into the pool we just prepared
        i++
      }
      const params = {
        fileName: file.value.name,
        fileMd5: fileMd5.value,
        chunkTotal: formDataList.value.length
      }
      const res = await findFile(params)
      // After all slicing is completed, send a request to the backend to pull the slicing information stored in the background of the current file to detect how many successfully uploaded slices there are.
      const finishList = res.data.file.ExaFileChunk // Successfully uploaded slices
      const IsFinish = res.data.file.IsFinish // Whether the same file has different names (if the file md5 is the same and the file name is different, the default is the same file but different file names. At this time, the backend database only needs to copy the database file without uploading the file. That is, the instant transmission function)
      if (!IsFinish) {
        // When resuming the upload from a breakpoint
        waitUpLoad.value = formDataList.value.filter(all => {
          return !(
            finishList &&
              finishList.some(fi => fi.FileChunkNumber === all.key)
          ) // Find the slices that need to be uploaded
        })
      } else {
        waitUpLoad.value = [] //If the second upload occurs, there are no slices to be uploaded.
        ElMessage.success('The file has been transferred in seconds')
      }
      waitNum.value = waitUpLoad.value.length // Record length is used for percentage display
    }
  } else {
    limitFileSize.value = true
    ElMessage('Please upload files smaller than 5M')
  }
}

const getFile = () => {
  // Confirm button
  if (file.value === null) {
    ElMessage('Please upload the file first')
    return
  }
  if (percentage.value === 100) {
    percentageFlage.value = false
  }
  sliceFile() // Upload slice
}

const sliceFile = () => {
  waitUpLoad.value &&
        waitUpLoad.value.forEach(item => {
          // Slices to be uploaded
          item.formData.append('chunkTotal', formDataList.value.length) // The total number of slices is carried to the background. It is always useful.
          const fileR = new FileReader() // Same function as above
          const fileF = item.formData.get('file')
          fileR.readAsArrayBuffer(fileF)
          fileR.onload = e => {
            const spark = new SparkMD5.ArrayBuffer()
            spark.append(e.target.result)
            item.formData.append('chunkMd5', spark.end()) // Get the current slice md5 backend to verify the slice integrity
            upLoadFileSlice(item)
          }
        })
}

watch(() => waitNum.value, () => { percentage.value = Math.floor(((formDataList.value.length - waitNum.value) / formDataList.value.length) * 100) })

const upLoadFileSlice = async(item) => {
  //Slice upload
  const fileRe = await breakpointContinue(item.formData)
  if (fileRe.code !== 0) {
    return
  }
  waitNum.value-- // Percentage increase
  if (waitNum.value === 0) {
    // After the slice is transmitted, the file is synthesized
    const params = {
      fileName: file.value.name,
      fileMd5: fileMd5.value
    }
    const res = await breakpointContinueFinish(params)
    if (res.code === 0) {
      // Delete cache tiles after compositing files
      const params = {
        fileName: file.value.name,
        fileMd5: fileMd5.value,
        filePath: res.data.filePath,
      }
      ElMessage.success('Upload successful')
      await removeChunk(params)
    }
  }
}

const FileInput = ref(null)
const inputChange = () => {
  FileInput.value.dispatchEvent(new MouseEvent('click'))
}
</script>

<style lang='scss' scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
#fromCont{
  display: inline-block;
}
.fileUpload{
    padding: 3px 10px;
    font-size: 12px;
    height: 20px;
    line-height: 20px;
    position: relative;
    cursor: pointer;
    color: #000;
    border: 1px solid #c1c1c1;
    border-radius: 4px;
    overflow: hidden;
    display: inline-block;
    input{
      position: absolute;
      font-size: 100px;
      right: 0;
      top: 0;
      opacity: 0;
      cursor: pointer;
    }
}
 .fileName{
    display: inline-block;
    vertical-align: top;
    margin: 6px 15px 0 15px;
  }
  .uploadBtn{
    position: relative;
    top: -10px;
    margin-left: 15px;
  }
  .tips{
    margin-top: 30px;
    font-size: 14px;
    font-weight: 400;
    color: #606266;
  }
  .el-divider{
    margin: 0 0 30px 0;
  }

 .list{
   margin-top:15px;
 }
 .list-item {
  display: block;
  margin-right: 10px;
  color: #606266;
  line-height: 25px;
  margin-bottom: 5px;
  width: 40%;
   .percentage{
          float: right;
        }
}
.list-enter-active, .list-leave-active {
  transition: all 1s;
}
.list-enter, .list-leave-to
/* .list-leave-active for below version 2.1.8 */ {
  opacity: 0;
  transform: translateY(-30px);
}
</style>
