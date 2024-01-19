<template>
  <div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button
          type="primary"
          icon="plus"
          @click="addMenu('0')"
        >Add root menu</el-button>
        <el-icon
          class="cursor-pointer"
          @click="toDoc('https://www.bilibili.com/video/BV1kv4y1g7nT/?p=4&vd_source=f2640257c21e3b547a790461ed94875e')"
        ><VideoCameraFilled /></el-icon>
      </div>

      <!-- Since the menu here corresponds one-to-one with the list on the left, there is no need for paging. The pageSize defaults to 999 -->
      <el-table
        :data="tableData"
        row-key="ID"
      >
        <el-table-column
          align="left"
          label="ID"
          min-width="100"
          prop="ID"
        />
        <el-table-column
          align="left"
          label="display name"
          min-width="120"
          prop="authorityName"
        >
          <template #default="scope">
            <span>{{ scope.row.meta.title }}</span>
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="icon"
          min-width="140"
          prop="authorityName"
        >
          <template #default="scope">
            <div
              v-if="scope.row.meta.icon"
              class="icon-column"
            >
              <el-icon>
                <component :is="scope.row.meta.icon" />
              </el-icon>
              <span>{{ scope.row.meta.icon }}</span>
            </div>
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="RouteName"
          show-overflow-tooltip
          min-width="160"
          prop="name"
        />
        <el-table-column
          align="left"
          label="Routing Path"
          show-overflow-tooltip
          min-width="160"
          prop="path"
        />
        <el-table-column
          align="left"
          label="whether to hide"
          min-width="100"
          prop="hidden"
        >
          <template #default="scope">
            <span>{{ scope.row.hidden?"Hide":"Show" }}</span>
          </template>
        </el-table-column>
        <el-table-column
          align="left"
          label="parent node"
          min-width="90"
          prop="parentId"
        />
        <el-table-column
          align="left"
          label="Sort"
          min-width="70"
          prop="sort"
        />
        <el-table-column
          align="left"
          label="file path"
          min-width="360"
          prop="component"
        />
        <el-table-column
          align="left"
          fixed="right"
          label="operation"
          width="300"
        >
          <template #default="scope">
            <el-button
              type="primary"
              link
              icon="plus"
              @click="addMenu(scope.row.ID)"
            >Add submenu</el-button>
            <el-button
              type="primary"
              link
              icon="edit"
              @click="editMenu(scope.row.ID)"
            >Edit</el-button>
            <el-button

              type="primary"
              link
              icon="delete"
              @click="deleteMenu(scope.row.ID)"
            >Delete</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>
    <el-dialog
      v-model="dialogFormVisible"
      :before-close="handleClose"
      :title="dialogTitle"
    >
      <warning-bar title="New menu needs to be configured with permissions in role management before it can be used" />
      <el-form
        v-if="dialogFormVisible"
        ref="menuForm"
        :inline="true"
        :model="form"
        :rules="rules"
        label-position="top"
        label-width="85px"
      >
        <el-form-item
          label="RouteName"
          prop="path"
          style="width:30%"
        >
          <el-input
            v-model="form.name"
            autocomplete="off"
            placeholder="unique English string"
            @change="changeName"
          />
        </el-form-item>
        <el-form-item
          prop="path"
          style="width:30%"
        >
          <template #label>
            <span style="display: inline-flex;align-items: center;">
              <span>Routing Path</span>
              <el-checkbox
                v-model="checkFlag"
                style="margin-left:12px;height: auto"
              >Add parameters</el-checkbox>
            </span>
          </template>

          <el-input
            v-model="form.path"
            :disabled="!checkFlag"
            autocomplete="off"
            placeholder="It is recommended to only splice parameters at the rear"
          />
        </el-form-item>
        <el-form-item
          label="whether to hide"
          style="width:30%"
        >
          <el-select
            v-model="form.hidden"
            placeholder="Whether to hide in the list"
          >
            <el-option
              :value="false"
              label="No"
            />
            <el-option
              :value="true"
              label="yes"
            />
          </el-select>
        </el-form-item>
        <el-form-item
          label="parent node ID"
          style="width:30%"
        >
          <el-cascader
            v-model="form.parentId"
            style="width:100%"
            :disabled="!isEdit"
            :options="menuOption"
            :props="{ checkStrictly: true,label:'title',value:'ID',disabled:'disabled',emitPath:false}"
            :show-all-levels="false"
            filterable
          />
        </el-form-item>
        <el-form-item
          label="file path"
          prop="component"
          style="width:60%"
        >
          <el-input
            v-model="form.component"
            autocomplete="off"
            placeholder="Page:view/xxx/xx.vue Plug-in:plugin/xx/xx.vue"
            @blur="fmtComponent"
          />
          <span style="font-size:12px;margin-right:12px;">If the menu contains submenus, please create a router-view secondary routing page or</span><el-button
            style="margin-top:4px"
            @click="form.component = 'view/routerHolder.vue'"
          >Click my settings</el-button>
        </el-form-item>
        <el-form-item
          label="display name"
          prop="meta.title"
          style="width:30%"
        >
          <el-input
            v-model="form.meta.title"
            autocomplete="off"
          />
        </el-form-item>
        <el-form-item
          label="icon"
          prop="meta.icon"
          style="width:30%"
        >
          <icon
            :meta="form.meta"
            style="width:100%"
          />
        </el-form-item>
        <el-form-item
          label="sort mark"
          prop="sort"
          style="width:30%"
        >
          <el-input
            v-model.number="form.sort"
            autocomplete="off"
          />
        </el-form-item>
        <el-form-item
          prop="meta.activeName"
          style="width:30%"
        >
          <template #label>
            <div>
              <span> Highlight menu </span>
              <el-tooltip
                content="Note: When this route is reached, the name specified in the left menu will be active (lit). It can be empty. If it is empty, it will be the name of this route."
                placement="top"
                effect="light"
              >
                <el-icon><QuestionFilled /></el-icon>
              </el-tooltip>
            </div>
          </template>
          <el-input
            v-model="form.meta.activeName"
            :placeholder="form.name"
            autocomplete="off"
          />
        </el-form-item>
        <el-form-item
          label="KeepAlive"
          prop="meta.keepAlive"
          style="width:30%"
        >
          <el-select
            v-model="form.meta.keepAlive"
            style="width:100%"
            placeholder="Whether to keepAlive cached pages"
          >
            <el-option
              :value="false"
              label="No"
            />
            <el-option
              :value="true"
              label="yes"
            />
          </el-select>
        </el-form-item>
        <el-form-item
          label="CloseTab"
          prop="meta.closeTab"
          style="width:30%"
        >
          <el-select
            v-model="form.meta.closeTab"
            style="width:100%"
            placeholder="whether to automatically close the tab"
          >
            <el-option
              :value="false"
              label="No"
            />
            <el-option
              :value="true"
              label="yes"
            />
          </el-select>
        </el-form-item>
        <el-form-item style="width:30%">
          <template #label>
            <div>
              <span> Whether it is a basic page </span>
              <el-tooltip
                content="If you select Yes, the left menu and top information will not be displayed."
                placement="top"
                effect="light"
              >
                <el-icon><QuestionFilled /></el-icon>
              </el-tooltip>
            </div>
          </template>

          <el-select
            v-model="form.meta.defaultMenu"
            style="width:100%"
            placeholder="whether it is a basic page"
          >
            <el-option
              :value="false"
              label="No"
            />
            <el-option
              :value="true"
              label="yes"
            />
          </el-select>
        </el-form-item>
      </el-form>
      <div>
        <div class="flex items-center gap-2">
          <el-button
            type="primary"
            icon="edit"
            @click="addParameter(form)"
          >New menu parameter</el-button>
          <el-icon
            class="cursor-pointer"
            @click="toDoc('https://www.bilibili.com/video/BV1kv4y1g7nT?p=9&vd_source=f2640257c21e3b547a790461ed94875e')"
          ><VideoCameraFilled /></el-icon>
        </div>
        <el-table
          :data="form.parameters"
          style="width: 100%;margin-top: 12px;"
        >
          <el-table-column
            align="left"
            prop="type"
            label="parameter type"
            width="180"
          >
            <template #default="scope">
              <el-select
                v-model="scope.row.type"
                placeholder="Please select"
              >
                <el-option
                  key="query"
                  value="query"
                  label="query"
                />
                <el-option
                  key="params"
                  value="params"
                  label="params"
                />
              </el-select>
            </template>
          </el-table-column>
          <el-table-column
            align="left"
            prop="key"
            label="parameter key"
            width="180"
          >
            <template #default="scope">
              <div>
                <el-input v-model="scope.row.key" />
              </div>
            </template>
          </el-table-column>
          <el-table-column
            align="left"
            prop="value"
            label="parameter value"
          >
            <template #default="scope">
              <div>
                <el-input v-model="scope.row.value" />
              </div>
            </template>
          </el-table-column>
          <el-table-column align="left">
            <template #default="scope">
              <div>
                <el-button
                  type="danger"

                  icon="delete"
                  @click="deleteParameter(form.parameters,scope.$index)"
                >Delete</el-button>
              </div>
            </template>
          </el-table-column>
        </el-table>

        <div class="flex items-center gap-2 mt-3">
          <el-button
            type="primary"
            icon="edit"
            @click="addBtn(form)"
          >New controllable buttons
          </el-button>
          <el-icon
            class="cursor-pointer"
            @click="toDoc('https://www.gin-vue-admin.com/guide/web/button-auth.html')"
          ><QuestionFilled /></el-icon>
          <el-icon
            class="cursor-pointer"
            @click="toDoc('https://www.bilibili.com/video/BV1kv4y1g7nT?p=11&vd_source=f2640257c21e3b547a790461ed94875e')"
          ><VideoCameraFilled /></el-icon>
        </div>

        <el-table
          :data="form.menuBtn"
          style="width: 100%;margin-top: 12px;"
        >
          <el-table-column
            align="left"
            prop="name"
            label="Button name"
            width="180"
          >
            <template #default="scope">
              <div>
                <el-input v-model="scope.row.name" />
              </div>
            </template>
          </el-table-column>
          <el-table-column
            align="left"
            prop="name"
            label="Remarks"
            width="180"
          >
            <template #default="scope">
              <div>
                <el-input v-model="scope.row.desc" />
              </div>
            </template>
          </el-table-column>
          <el-table-column align="left">
            <template #default="scope">
              <div>
                <el-button
                  type="danger"

                  icon="delete"
                  @click="deleteBtn(form.menuBtn,scope.$index)"
                >Delete</el-button>
              </div>
            </template>
          </el-table-column>
        </el-table>
      </div>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">Cancel</el-button>
          <el-button
            type="primary"
            @click="enterDialog"
          >Confirm</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  updateBaseMenu,
  getMenuList,
  addBaseMenu,
  deleteBaseMenu,
  getBaseMenuById
} from '@/api/menu'
import icon from '@/view/superAdmin/menu/icon.vue'
import WarningBar from '@/components/warningBar/warningBar.vue'
import { canRemoveAuthorityBtnApi } from '@/api/authorityBtn'
import { reactive, ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { QuestionFilled, VideoCameraFilled } from '@element-plus/icons-vue'

import { toDoc } from '@/utils/doc'

defineOptions({
  name: 'Menus',
})

const rules = reactive({
  path: [{ required: true, message: 'Please enter the menu name', trigger: 'blur' }],
  component: [
    { required: true, message: 'Please enter the file path', trigger: 'blur' }
  ],
  'meta.title': [
    { required: true, message: 'Please enter the menu display name', trigger: 'blur' }
  ]
})

const page = ref(1)
const total = ref(0)
const pageSize = ref(999)
const tableData = ref([])
const searchInfo = ref({})
// Inquire
const getTableData = async() => {
  const table = await getMenuList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

//Add new parameters
const addParameter = (form) => {
  if (!form.parameters) {
    form.parameters = []
  }
  form.parameters.push({
    type: 'query',
    key: '',
    value: ''
  })
}

const fmtComponent = () => {
  form.value.component = form.value.component.replace(/\\/g, '/')
}

// Delete parameters
const deleteParameter = (parameters, index) => {
  parameters.splice(index, 1)
}

//Add controllable button
const addBtn = (form) => {
  if (!form.menuBtn) {
    form.menuBtn = []
  }
  form.menuBtn.push({
    name: '',
    desc: '',
  })
}
//Delete controllable button
const deleteBtn = async(btns, index) => {
  const btn = btns[index]
  if (btn.ID === 0) {
    btns.splice(index, 1)
    return
  }
  const res = await canRemoveAuthorityBtnApi({ id: btn.ID })
  if (res.code === 0) {
    btns.splice(index, 1)
  }
}

const form = ref({
  ID: 0,
  path: '',
  name: '',
  hidden: false,
  parentId: '',
  component: '',
  meta: {
    activeName: '',
    title: '',
    icon: '',
    defaultMenu: false,
    closeTab: false,
    keepAlive: false
  },
  parameters: [],
  menuBtn: []
})
const changeName = () => {
  form.value.path = form.value.name
}

const handleClose = (done) => {
  initForm()
  done()
}
//delete menu
const deleteMenu = (ID) => {
  ElMessageBox.confirm('This operation will permanently delete this menu under all roles, do you want to continue?', 'Prompt', {
    confirmButtonText: 'OK',
    cancelButtonText: 'Cancel',
    type: 'warning'
  })
    .then(async() => {
      const res = await deleteBaseMenu({ ID })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: 'Delete successfully!'
        })
        if (tableData.value.length === 1 && page.value > 1) {
          page.value--
        }
        getTableData()
      }
    })
    .catch(() => {
      ElMessage({
        type: 'info',
        message: 'Deletion canceled'
      })
    })
}
// Initialize the table method in the pop-up window
const menuForm = ref(null)
const checkFlag = ref(false)
const initForm = () => {
  checkFlag.value = false
  menuForm.value.resetFields()
  form.value = {
    ID: 0,
    path: '',
    name: '',
    hidden: false,
    parentId: '',
    component: '',
    meta: {
      title: '',
      icon: '',
      defaultMenu: false,
      closeTab: false,
      keepAlive: false
    }
  }
}
//Close pop-up window

const dialogFormVisible = ref(false)
const closeDialog = () => {
  initForm()
  dialogFormVisible.value = false
}
//Add menu
const enterDialog = async() => {
  menuForm.value.validate(async valid => {
    if (valid) {
      let res
      if (isEdit.value) {
        res = await updateBaseMenu(form.value)
      } else {
        res = await addBaseMenu(form.value)
      }
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: isEdit.value ? 'Edited successfully' : 'Added successfully!'
        })
        getTableData()
      }
      initForm()
      dialogFormVisible.value = false
    }
  })
}

const menuOption = ref([
  {
    ID: '0',
    title: 'Root Menu'
  }
])
const setOptions = () => {
  menuOption.value = [
    {
      ID: '0',
      title: 'root directory'
    }
  ]
  setMenuOptions(tableData.value, menuOption.value, false)
}
const setMenuOptions = (menuData, optionsData, disabled) => {
  menuData &&
        menuData.forEach(item => {
          if (item.children && item.children.length) {
            const option = {
              title: item.meta.title,
              ID: String(item.ID),
              disabled: disabled || item.ID === form.value.ID,
              children: []
            }
            setMenuOptions(
              item.children,
              option.children,
              disabled || item.ID === form.value.ID
            )
            optionsData.push(option)
          } else {
            const option = {
              title: item.meta.title,
              ID: String(item.ID),
              disabled: disabled || item.ID === form.value.ID
            }
            optionsData.push(option)
          }
        })
}

//Add menu method, if id is 0, add root menu
const isEdit = ref(false)
const dialogTitle = ref('Add menu')
const addMenu = (id) => {
  dialogTitle.value = 'New menu'
  form.value.parentId = String(id)
  isEdit.value = false
  setOptions()
  dialogFormVisible.value = true
}
//Modify menu method
const editMenu = async(id) => {
  dialogTitle.value = 'Edit menu'
  const res = await getBaseMenuById({ id })
  form.value = res.data.menu
  isEdit.value = true
  setOptions()
  dialogFormVisible.value = true
}

</script>

<style scoped lang="scss">
.warning {
  color: #dc143c;
}
.icon-column{
  display: flex;
  align-items: center;
  .el-icon{
    margin-right: 8px;
  }
}
</style>
