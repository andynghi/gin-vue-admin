<template>
  <div class="gva-table-box">
    <warning-bar title="There is some instability when using the GPT-3.5 model, and the success rate is about 50%. Using GPT-4 can greatly improve the success rate, but the cost is higher." />
    <div v-if="!chatToken">
      <el-input
        v-model="skObj.sk"
        class="query-ipt"
        placeholder="Please enter your ChatGpt SK"
        clearable
      />
      <el-button
        type="primary"
        @click="save"
      >Save</el-button>
      <div class="secret">
        <el-empty description="Please go to the gpt website to get your sk: https://platform.openai.com/account/api-keys" />
      </div>
    </div>
    <div v-else>
      <el-form
        :model="form"
        label-width="120px"
      >
        <el-form-item label="Delete current sk:">
          <el-popover
            placement="top"
            width="160"
          >
            <p>Are you sure you want to delete and return? </p>
            <div style="text-align: right; margin-top: 8px;">
              <el-button
                type="primary"
                @click="deleteSK"
              >OK</el-button>
            </div>
            <template #reference>
              <el-button
                type="primary"
                link
                icon="delete"
              >Delete</el-button>
            </template>
          </el-popover>
        </el-form-item>
        <el-form-item label="Query db name:">
          <el-select
            v-model="form.dbname"
            placeholder="Please select a library"
            style="width: 115px"
          >
            <el-option
              v-for="(item, index) in dbArr"
              :key="index"
              :label="item.database"
              :value="item.database"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="Query db description:">
          <el-input
            v-model="form.chat"
            :autosize="{ minRows: 2, maxRows: 4 }"
            type="textarea"
            clearable
            placeholder="Please enter the conversation"
          />
        </el-form-item>
        <el-form-item label="GPT generates SQL:">
          <el-input
            v-model="sql"
            :autosize="{ minRows: 2, maxRows: 4 }"
            type="textarea"
            disabled
            placeholder="Display automatically generated sql here"
          />
        </el-form-item>
        <el-button
          type="primary"
          @click="handleQueryTable"
        >Query</el-button>
      </el-form>
      <div class="tables">
        <el-table
          v-if="tableData.length"
          ref="multipleTable"
          :data="tableData"
          style="width: 100%"
          tooltip-effect="dark"
          height="400px"
        >
          <el-table-column
            v-for="(item, index) in tableData[0]"
            :key="index"
            :prop="index"
            :label="index"
            min-width="200"
            show-overflow-tooltip
          />
        </el-table>
        <p
          v-else
          class="text"
        >Please enter the content you need AI to query for you in the dialog box:)</p>
      </div>
    </div>
  </div>
</template>

<script setup>
import WarningBar from '@/components/warningBar/warningBar.vue'
import { getTableApi,
  createSKApi,
  getSKApi,
  deleteSKApi } from '@/api/chatgpt'
import { getDB as getDBAPI } from '@/api/autoCode'
import { ref, reactive } from 'vue'

const chatToken = ref(null)
const skObj = reactive({
  sk: '',
})
const sql = ref('')
const getSK = async() => {
  const res = await getSKApi()
  chatToken.value = res.data.ok
}

const getDB = async() => {
  const res = await getDBAPI()
  if (res.code === 0) {
    dbArr.value = res.data.dbs
  }
}
getSK()
getDB()
const save = async() => {
  const res = await createSKApi(skObj)
  if (res.code === 0) {
    await getSK()
  }
}

const deleteSK = async() => {
  const res = await deleteSKApi()
  if (res.code === 0) {
    await getSK()
  }
}

const form = ref({
  dbname: '',
  chat: '',
})
const dbArr = ref([])
const tableData = ref([])

const handleQueryTable = async() => {
  const res = await getTableApi(form.value)
  if (res.code === 0) {
    tableData.value = res.data.results || []
  }
  sql.value = res.data.sql
  // Dynamically render the table based on the background return value
}
</script>

<style scoped lang="scss">
.secret{
  padding: 30px;
  margin-top: 20px;
  background: #F5F5F5;
  p {
    line-height: 30px;
  }
}
.query-ipt{
  width: 300px;
  margin-right: 30px;
}
.content{
  p {
    font-size: 16px;
    line-height: 20px;
  }
  padding: 10px;
  width: 100%;
  background: #F5F5F5;
  margin-top: 30px;
}
.tables{
  width: 100%;
  margin-top: 30px;
  .text{
    font-size: 18px;
    color: #6B7687;
    text-align: center;
  }
}
</style>
