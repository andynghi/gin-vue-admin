<template>
  <div class="system">
    <el-form
      ref="form"
      :model="config"
      label-width="240px"
    >
      <!--  System start  -->
      <el-collapse v-model="activeNames">
        <el-collapse-item
          title="System Configuration"
          name="1"
        >
          <el-form-item label="environment value">
            <!-- <el-input v-model="config.system.env" />-->
            <el-select
              v-model="config.system.env"
              style="width:100%"
            >
              <el-option value="public" />
              <el-option value="develop" />
            </el-select>
          </el-form-item>
          <el-form-item label="port value">
            <el-input v-model.number="config.system.addr" />
          </el-form-item>
          <el-form-item label="Database type">
            <el-select
              v-model="config.system['db-type']"
              style="width:100%"
            >
              <el-option value="mysql" />
              <el-option value="pgsql" />
            </el-select>
          </el-form-item>
          <el-form-item label="Oss type">
            <el-select
              v-model="config.system['oss-type']"
              style="width:100%"
            >
              <el-option value="local" />
              <el-option value="qiniu" />
              <el-option value="tencent-cos" />
              <el-option value="aliyun-oss" />
              <el-option value="huawei-obs" />
            </el-select>
          </el-form-item>
          <el-form-item label="Multiple login interception">
            <el-checkbox v-model="config.system['use-multipoint']">Enable</el-checkbox>
          </el-form-item>
          <el-form-item label="Enable redis">
            <el-checkbox v-model="config.system['use-redis']">Enable</el-checkbox>
          </el-form-item>
          <el-form-item label="Number of current limits">
            <el-input-number v-model.number="config.system['iplimit-count']" />
          </el-form-item>
          <el-form-item label="current limit time">
            <el-input-number v-model.number="config.system['iplimit-time']" />
          </el-form-item>
          <el-tooltip
            content="After the modification is completed, please pay attention to modifying VITE_BASE_PATH in the front-end env environment"
            placement="top-start"
          >
            <el-form-item label="Global routing prefix">
              <el-input v-model="config.system['router-prefix']" />
            </el-form-item>
          </el-tooltip>
        </el-collapse-item>
        <el-collapse-item
          title="jwt signature"
          name="2"
        >
          <el-form-item label="jwt signature">
            <el-input v-model="config.jwt['signing-key']" />
          </el-form-item>
          <el-form-item label="Validity period">
            <el-input v-model="config.jwt['expires-time']" />
          </el-form-item>
          <el-form-item label="Buffer period">
            <el-input v-model="config.jwt['buffer-time']" />
          </el-form-item>
          <el-form-item label="Issued by">
            <el-input v-model="config.jwt.issuer" />
          </el-form-item>
        </el-collapse-item>
        <el-collapse-item
          title="Zap log configuration"
          name="3"
        >
          <el-form-item label="level">
            <el-input v-model.number="config.zap.level" />
          </el-form-item>
          <el-form-item label="output">
            <el-input v-model="config.zap.format" />
          </el-form-item>
          <el-form-item label="Log prefix">
            <el-input v-model="config.zap.prefix" />
          </el-form-item>
          <el-form-item label="Log folder">
            <el-input v-model="config.zap.director" />
          </el-form-item>
          <el-form-item label="encoding level">
            <el-input v-model="config.zap['encode-level']" />
          </el-form-item>
          <el-form-item label="stack name">
            <el-input v-model="config.zap['stacktrace-key']" />
          </el-form-item>
          <el-form-item label="Log retention time (default is in days)">
            <el-input v-model.number="config.zap['max-age']" />
          </el-form-item>
          <el-form-item label="show row">
            <el-checkbox v-model="config.zap['show-line']" />
          </el-form-item>
          <el-form-item label="Output console">
            <el-checkbox v-model="config.zap['log-in-console']" />
          </el-form-item>
        </el-collapse-item>
        <el-collapse-item
          title="Redis admin database configuration"
          name="4"
        >
          <el-form-item label="Library">
            <el-input v-model.number="config.redis.db" />
          </el-form-item>
          <el-form-item label="address">
            <el-input v-model="config.redis.addr" />
          </el-form-item>
          <el-form-item label="password">
            <el-input v-model="config.redis.password" />
          </el-form-item>
        </el-collapse-item>

        <el-collapse-item
          title="Mongo database configuration"
          name="14"
        >
          <el-form-item label="collection name (table name, generally not written)">
            <el-input v-model="config.mongo.coll" />
          </el-form-item>
          <el-form-item label="mongodb options">
            <el-input v-model="config.mongo.options" />
          </el-form-item>
          <el-form-item label="database name(database name)">
            <el-input v-model="config.mongo.database" />
          </el-form-item>
          <el-form-item label="username">
            <el-input v-model="config.mongo.username" />
          </el-form-item>
          <el-form-item label="password">
            <el-input v-model="config.mongo.password" />
          </el-form-item>
          <el-form-item label="Minimum connection pool">
            <el-input v-model="config.mongo['min-pool-size']" />
          </el-form-item>
          <el-form-item label="Maximum connection pool">
            <el-input v-model="config.mongo['max-pool-size']" />
          </el-form-item>
          <el-form-item label="socket timeout">
            <el-input v-model="config.mongo['socket-timeout-ms']" />
          </el-form-item>
          <el-form-item label="Connection timeout">
            <el-input v-model="config.mongo['socket-timeout-ms']" />
          </el-form-item>
          <el-form-item label="Whether to enable zap log">
            <el-checkbox v-model="config.mongo['is-zap']" />
          </el-form-item>
          <el-form-item label="hosts">
            <template v-for="(item,k) in config.mongo.hosts">
            <div
              v-for="(_,k2) in item"
              :key="k2"
            >
              <el-form-item
                :key="k+k2"
                :label="k2"
              >
                <el-input v-model="item[k2]" />
              </el-form-item>
            </div>
          </template>
          </el-form-item>
        </el-collapse-item>

        <el-collapse-item
          title="Email Configuration"
          name="5"
        >
          <el-form-item label="recipient email">
            <el-input
              v-model="config.email.to"
              placeholder="can be multiple, separated by commas"
            />
          </el-form-item>
          <el-form-item label="port">
            <el-input v-model.number="config.email.port" />
          </el-form-item>
          <el-form-item label="Sender's Email">
            <el-input v-model="config.email.from" />
          </el-form-item>
          <el-form-item label="host">
            <el-input v-model="config.email.host" />
          </el-form-item>
          <el-form-item label="Is it ssl">
            <el-checkbox v-model="config.email['is-ssl']" />
          </el-form-item>
          <el-form-item label="secret">
            <el-input v-model="config.email.secret" />
          </el-form-item>
          <el-form-item label="Test email">
            <el-button @click="email">Test email</el-button>
          </el-form-item>
        </el-collapse-item>
        <el-collapse-item
          title="Verification code configuration"
          name="7"
        >
          <el-form-item label="Character length">
            <el-input v-model.number="config.captcha['key-long']" />
          </el-form-item>
          <el-form-item label="image width">
            <el-input v-model.number="config.captcha['img-width']" />
          </el-form-item>
          <el-form-item label="image height">
            <el-input v-model.number="config.captcha['img-height']" />
          </el-form-item>
        </el-collapse-item>
        <el-collapse-item
          title="Database Configuration"
          name="9"
        >
          <template v-if="config.system['db-type'] === 'mysql'">
            <el-form-item label="username">
              <el-input v-model="config.mysql.username" />
            </el-form-item>
            <el-form-item label="password">
              <el-input v-model="config.mysql.password" />
            </el-form-item>
            <el-form-item label="address">
              <el-input v-model="config.mysql.path" />
            </el-form-item>
            <el-form-item label="database">
              <el-input v-model="config.mysql['db-name']" />
            </el-form-item>
            <el-form-item label="prefix">
              <el-input v-model="config.mysql['refix']" />
            </el-form-item>
            <el-form-item label="Plural table">
              <el-switch v-model="config.mysql['singular']" />
            </el-form-item>
            <el-form-item label="Engine">
              <el-input v-model="config.mysql['engine']" />
            </el-form-item>
            <el-form-item label="maxIdleConns">
              <el-input v-model.number="config.mysql['max-idle-conns']" />
            </el-form-item>
            <el-form-item label="maxOpenConns">
              <el-input v-model.number="config.mysql['max-open-conns']" />
            </el-form-item>
            <el-form-item label="Write log">
              <el-checkbox v-model="config.mysql['log-zap']" />
            </el-form-item>
            <el-form-item label="Log mode">
              <el-input v-model="config.mysql['log-mode']" />
            </el-form-item>
          </template>
          <template v-if="config.system['db-type'] === 'pgsql'">
            <el-form-item label="username">
              <el-input v-model="config.pgsql.username" />
            </el-form-item>
            <el-form-item label="password">
              <el-input v-model="config.pgsql.password" />
            </el-form-item>
            <el-form-item label="address">
              <el-input v-model="config.pgsql.path" />
            </el-form-item>
            <el-form-item label="database">
              <el-input v-model="config.pgsql.dbname" />
            </el-form-item>
            <el-form-item label="prefix">
              <el-input v-model="config.pgsql['refix']" />
            </el-form-item>
            <el-form-item label="Plural table">
              <el-switch v-model="config.pgsql['singular']" />
            </el-form-item>
            <el-form-item label="Engine">
              <el-input v-model="config.pgsql['engine']" />
            </el-form-item>
            <el-form-item label="maxIdleConns">
              <el-input v-model.number="config.pgsql['max-idle-conns']" />
            </el-form-item>
            <el-form-item label="maxOpenConns">
              <el-input v-model.number="config.pgsql['max-open-conns']" />
            </el-form-item>
            <el-form-item label="Write log">
              <el-checkbox v-model="config.pgsql['log-zap']" />
            </el-form-item>
            <el-form-item label="Log mode">
              <el-input v-model="config.pgsql['log-mode']" />
            </el-form-item>
          </template>
        </el-collapse-item>

        <el-collapse-item
          title="oss configuration"
          name="10"
        >
          <template v-if="config.system['oss-type'] === 'local'">
            <h2>Local file configuration</h2>
            <el-form-item label="Local file access path">
              <el-input v-model="config.local.path" />
            </el-form-item>
            <el-form-item label="Local file storage path">
              <el-input v-model="config.local['store-path']" />
            </el-form-item>
          </template>
          <template v-if="config.system['oss-type'] === 'qiniu'">
            <h2>qiniu upload configuration</h2>
            <el-form-item label="storage area">
              <el-input v-model="config.qiniu.zone" />
            </el-form-item>
            <el-form-item label="Space name">
              <el-input v-model="config.qiniu.bucket" />
            </el-form-item>
            <el-form-item label="CDN accelerated domain name">
              <el-input v-model="config.qiniu['img-path']" />
            </el-form-item>
            <el-form-item label="Whether to use https">
              <el-checkbox v-model="config.qiniu['use-https']">Enable</el-checkbox>
            </el-form-item>
            <el-form-item label="accessKey">
              <el-input v-model="config.qiniu['access-key']" />
            </el-form-item>
            <el-form-item label="secretKey">
              <el-input v-model="config.qiniu['secret-key']" />
            </el-form-item>
            <el-form-item label="Whether the upload uses CDN upload acceleration">
              <el-checkbox v-model="config.qiniu['use-cdn-domains']">Enable</el-checkbox>
            </el-form-item>
          </template>
          <template v-if="config.system['oss-type'] === 'tencent-cos'">
            <h2>Tencent Cloud COS upload configuration</h2>
            <el-form-item label="bucket name">
              <el-input v-model="config['tencent-cos']['bucket']" />
            </el-form-item>
            <el-form-item label="Region">
              <el-input v-model="config['tencent-cos'].region" />
            </el-form-item>
            <el-form-item label="secretID">
              <el-input v-model="config['tencent-cos']['secret-id']" />
            </el-form-item>
            <el-form-item label="secretKey">
              <el-input v-model="config['tencent-cos']['secret-key']" />
            </el-form-item>
            <el-form-item label="path prefix">
              <el-input v-model="config['tencent-cos']['path-prefix']" />
            </el-form-item>
            <el-form-item label="Access domain name">
              <el-input v-model="config['tencent-cos']['base-url']" />
            </el-form-item>
          </template>
          <template v-if="config.system['oss-type'] === 'aliyun-oss'">
            <h2>Alibaba Cloud OSS upload configuration</h2>
            <el-form-item label="area">
              <el-input v-model="config['aliyun-oss'].endpoint" />
            </el-form-item>
            <el-form-item label="accessKeyId">
              <el-input v-model="config['aliyun-oss']['access-key-id']" />
            </el-form-item>
            <el-form-item label="accessKeySecret">
              <el-input v-model="config['aliyun-oss']['access-key-secret']" />
            </el-form-item>
            <el-form-item label="bucket name">
              <el-input v-model="config['aliyun-oss']['bucket-name']" />
            </el-form-item>
            <el-form-item label="Access domain name">
              <el-input v-model="config['aliyun-oss']['bucket-url']" />
            </el-form-item>
          </template>
          <template v-if="config.system['oss-type'] === 'huawei-obs'">
            <h2>Huawei Cloud Obs upload configuration</h2>
            <el-form-item label="path">
              <el-input v-model="config['hua-wei-obs'].path" />
            </el-form-item>
            <el-form-item label="bucket name">
              <el-input v-model="config['hua-wei-obs'].bucket" />
            </el-form-item>
            <el-form-item label="area">
              <el-input v-model="config['hua-wei-obs'].endpoint" />
            </el-form-item>
            <el-form-item label="accessKey">
              <el-input v-model="config['hua-wei-obs']['access-key']" />
            </el-form-item>
            <el-form-item label="secretKey">
              <el-input v-model="config['hua-wei-obs']['secret-key']" />
            </el-form-item>
          </template>

        </el-collapse-item>

        <el-collapse-item
          title="Excel upload configuration"
          name="11"
        >
          <el-form-item label="Synthetic target address">
            <el-input v-model="config.excel.dir" />
          </el-form-item>
        </el-collapse-item>

        <el-collapse-item
          title="Automated code configuration"
          name="12"
        >
          <el-form-item label="Whether to automatically restart (linux)">
            <el-checkbox v-model="config.autocode['transfer-restart']" />
          </el-form-item>
          <el-form-item label="root(project root path)">
            <el-input
              v-model="config.autocode.root"
              disabled
            />
          </el-form-item>
          <el-form-item label="Server(backend code address)">
            <el-input v-model="config.autocode['transfer-restart']" />
          </el-form-item>
          <el-form-item label="SApi (backend api folder address)">
            <el-input v-model="config.autocode['server-api']" />
          </el-form-item>
          <el-form-item label="SInitialize(backend Initialize folder)">
            <el-input v-model="config.autocode['server-initialize']" />
          </el-form-item>
          <el-form-item label="SModel(backend Model file address)">
            <el-input v-model="config.autocode['server-model']" />
          </el-form-item>
          <el-form-item label="SRequest(backend Request folder address)">
            <el-input v-model="config.autocode['server-request']" />
          </el-form-item>
          <el-form-item label="SRouter (backend Router folder address)">
            <el-input v-model="config.autocode['server-router']" />
          </el-form-item>
          <el-form-item label="SService (backend Service folder address)">
            <el-input v-model="config.autocode['server-service']" />
          </el-form-item>
          <el-form-item label="Web(front-end folder address)">
            <el-input v-model="config.autocode.web" />
          </el-form-item>
          <el-form-item label="WApi (backend WApi folder address)">
            <el-input v-model="config.autocode['web-api']" />
          </el-form-item>
          <el-form-item label="WForm(backend WForm folder address)">
            <el-input v-model="config.autocode['web-form']" />
          </el-form-item>
          <el-form-item label="WTable(backend WTable folder address)">
            <el-input v-model="config.autocode['web-table']" />
          </el-form-item>
        </el-collapse-item>

        <el-collapse-item
          title="Timer(Timer Task)"
          name="13"
        >
          <el-form-item label="Start (whether to enable)">
            <el-checkbox v-model="config.timer['start']" />
          </el-form-item>
          <el-form-item label="Spec(CRON expression)">
            <el-input v-model="config.timer.spec" />
          </el-form-item>
          <template v-for="(item,k) in config.timer.detail">
            <div
              v-for="(_,k2) in item"
              :key="k2"
            >
              <el-form-item
                :key="k+k2"
                :label="k2"
              >
                <el-input v-model="item[k2]" />
              </el-form-item>
            </div>
          </template>
        </el-collapse-item>
      </el-collapse>
    </el-form>
    <div class="mt-4">
      <el-button
        type="primary"
        @click="update"
      >Update now</el-button>
      <el-button
        type="primary"
        @click="reload"
      >Restart service (under development)</el-button>
    </div>
  </div>
</template>

<script setup>
import { getSystemConfig, setSystemConfig } from '@/api/system'
import { emailTest } from '@/api/email'
import { ref, reactive } from 'vue'
import { ElMessage } from 'element-plus'

defineOptions({
  name: 'Config'
})

const activeNames = reactive([])
const config = ref({
  system: {
    'iplimit-count': 0,
    'iplimit-time': 0
  },
  jwt: {},
  mysql: {},
  pgsql: {},
  excel: {},
  autocode: {},
  redis: {},
  mongo: {
    coll: '',
    options: '',
    database: '',
    username: '',
    password: '',
    'min-pool-size': '',
    'max-pool-size': '',
    'socket-timeout-ms': '',
    'connect-timeout-ms': '',
    'is-zap': '',
    hosts: [
      {
        host: '',
        port: ''
      }
    ]
  },
  qiniu: {},
  'tencent-cos': {},
  'aliyun-oss': {},
  'hua-wei-obs': {},
  captcha: {},
  zap: {},
  local: {},
  email: {},
  timer: {
    detail: {}
  }
})

const initForm = async() => {
  const res = await getSystemConfig()
  if (res.code === 0) {
    config.value = res.data.config
  }
}
initForm()
const reload = () => {}
const update = async() => {
  const res = await setSystemConfig({ config: config.value })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: 'Configuration file set successfully'
    })
    await initForm()
  }
}
const email = async() => {
  const res = await emailTest()
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: 'Mail sent successfully'
    })
    await initForm()
  } else {
    ElMessage({
      type: 'error',
      message: 'Failed to send email'
    })
  }
}

</script>

<style lang="scss">
.system {
  @apply bg-white p-9 rounded;
  h2 {
    @apply p-2.5 my-2.5 text-lg shadow;
  }
}
</style>
