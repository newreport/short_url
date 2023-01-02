<template>
  <div>
    <el-form :inline="true" label-position="left" label-width="68px" :model="form">
      <el-form-item label="用户名">
        <el-input v-model="form.name" />
      </el-form-item>
      <el-form-item label="昵称">
        <el-input v-model="form.nickname" />
      </el-form-item>
      <el-form-item label="手机号">
        <el-input v-model="form.phone" />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="getUsersPage()">Query</el-button>
      </el-form-item>
    </el-form>

    <el-table :data="tableData" border style="width: 100%">
      <el-table-column type="selection" width="55" />
      <el-table-column prop="group" sortable label="分组" />
      <el-table-column prop="name" sortable label="用户名" />
      <el-table-column prop="domain" sortable label="域名" />
      <el-table-column prop="nickname" sortable label="昵称" />
      <el-table-column prop="urlLength" sortable label="默认长度" />
      <el-table-column prop="phone" sortable label="电话" />
      <el-table-column label="权限">
        <template #default="scope">
          {{ scope.row.role == 1 ? "管理员" : "普通用户" }}
        </template>
      </el-table-column>
      <el-table-column prop="remarks" label="备注" />
      <el-table-column fixed="right" width="193">
        <template #header>
          <el-button size="small" @click="dialogVisible = true; cleanAddUser();">新增</el-button>
        </template>
        <template #default="scope">


          <el-button link type="primary" size="small" @click="
  formAddEdit.id = scope.row.id;
formAddEdit.author = scope.row.author;
formAddEdit.autoInsertSpace = scope.row.autoInsertSpace;
formAddEdit.domain = scope.row.domain;
formAddEdit.group = scope.row.group;
formAddEdit.i18n = scope.row.i18n;
formAddEdit.id = scope.row.id;
formAddEdit.name = scope.row.name;
formAddEdit.nickname = scope.row.nickname;
formAddEdit.phone = scope.row.phone;
formAddEdit.pwd = scope.row.pwd;
formAddEdit.remarks = scope.row.remarks;
formAddEdit.role = scope.row.role;
formAddEdit.urlLength = scope.row.urlLength;
dialogVisible = true;
          ">Edit</el-button>
          <el-popconfirm title="确定删除吗?">
            <template #reference>
              <el-button link type="primary" size="small">Delete</el-button>
            </template>
          </el-popconfirm>
        </template>
      </el-table-column>
    </el-table>


    <el-dialog v-model="dialogVisible" :title="formAddEdit.id > 0 ? '修改用户' : '新增用户'">
      <el-form label-position="left" label-width="79px">


        <el-form-item label="用户名">
          <el-input v-model="formAddEdit.name" placeholder="请输入用户名" />
        </el-form-item>


        <el-form-item label="密码">
          <el-input v-model="formAddEdit.pwd" type="password" placeholder="请输入密码" />
        </el-form-item>

        <el-form-item label="域名">
          <el-input v-model="formAddEdit.domain" placeholder="请输入域名" />
        </el-form-item>


        <el-form-item label="昵称">
          <el-input v-model="formAddEdit.nickname" />
        </el-form-item>


        <el-form-item label="权限">
          <el-input v-model="formAddEdit.role" />
        </el-form-item>

        <el-form-item label="默认长度">
          <el-input-number v-model="formAddEdit.urlLength" :min="4" :max="16"></el-input-number>
        </el-form-item>


        <el-form-item label="分组">
          <el-input v-model="formAddEdit.group" />
        </el-form-item>


        <el-form-item label="联系方式">
          <el-input v-model="formAddEdit.phone" />
        </el-form-item>

        <el-form-item label="留白">
          <el-switch v-model="formAddEdit.autoInsertSpace" class="mt-2" style="margin-left: 24px" inline-prompt
            :active-icon="Check" :inactive-icon="Close" />
        </el-form-item>

        <el-form-item label="国际化">
          <el-select v-model="formAddEdit.i18n" placeholder="请选择">
            <el-option value="中文"></el-option>
            <el-option value="Eenlish"></el-option>
            <el-option value="日本語"></el-option>
          </el-select>

        </el-form-item>

        <el-form-item label="备注">
          <el-input v-model="formAddEdit.remarks" />
        </el-form-item>



        <el-form-item>
          <el-button type="primary" @click="formAddEdit.id > 0 ? updateUser() : addUser()"> Submit</el-button>
          <el-button @click="cleanAddUser()">Reset</el-button>
        </el-form-item>
      </el-form>
    </el-dialog>
  </div>
</template>
  
<script lang="ts" setup>
import { ref, reactive } from 'vue'
import { UserService } from '@/api/api'
import { Check, Close } from '@element-plus/icons-vue'
import { el } from 'element-plus/es/locale';


const form = reactive({
  name: '',
  nickname: '',
  phone: '',
})


const tableData = ref()

const page = reactive({
  currentPage: 0,
  pageSize: 100,
  count: 10
})



const getUsersPage = () => {
  const getUsersPage = async () => {
    const getUsersPageParams = {
      offset: page.currentPage,
      limit: page.pageSize,
      sort: "",
      name: form.name.length == 0 ? "" : "lk" + form.name,
      nickname: form.nickname.length == 0 ? "" : "lk" + form.nickname,
      phone: form.phone.length == 0 ? "" : "lk" + form.phone
    }
    UserService.getUsersPage(getUsersPageParams).then(result => {
      if (result?.status == 200) {
        page.count = result.data.count
        tableData.value = result.data.data
      }
    })
  }
  getUsersPage()
}

const dialogVisible = ref(false)

const formAddEdit = reactive({
  author: '',//头像
  autoInsertSpace: false,
  crt: '',//创建时间
  domain: '',//域名
  group: '',
  i18n: '',
  id: 0,
  name: '',
  nickname: '',
  phone: '',
  pwd: '',
  remarks: '',
  role: 0,
  upt: '',
  urlLength: 6
})

const updateUser = () => {
  if (!formAddEdit.domain || !formAddEdit.name || !formAddEdit.pwd) {
    ElMessage.warning("请输入完整")
    return
  }
  dialogVisible.value = false

  const updateUser = async () => {
    const updateUserParams = {
      id:formAddEdit.id,
      autoInsertSpace: formAddEdit.autoInsertSpace,
      domain: formAddEdit.domain,
      group: formAddEdit.group,
      i18n: formAddEdit.i18n,
      name: formAddEdit.name,
      nickname: formAddEdit.nickname,
      phone: formAddEdit.phone,
      pwd: formAddEdit.pwd,
      remarks: formAddEdit.remarks,
      role: formAddEdit.role,
      urlLength: formAddEdit.urlLength
    }
    UserService.updateUser(updateUserParams).
      then(result => {
        if (result?.status == 200) {
          ElMessage.success(result.data)
          cleanAddUser()
          getUsersPage()
        } else {
          ElMessage.success(result.data)
          cleanAddUser()
          getUsersPage()
        }
      })
  }
  updateUser()

}

const addUser = () => {
  if (!formAddEdit.domain || !formAddEdit.name || !formAddEdit.pwd) {
    ElMessage.warning("请输入完整")
    return
  }
  dialogVisible.value = false

  const addUser = async () => {
    const addUserParams = {
      autoInsertSpace: formAddEdit.autoInsertSpace,
      domain: formAddEdit.domain,
      group: formAddEdit.domain,
      i18n: formAddEdit.i18n,
      name: formAddEdit.name,
      nickname: formAddEdit.nickname,
      phone: formAddEdit.phone,
      pwd: formAddEdit.pwd,
      remarks: formAddEdit.remarks,
      role: formAddEdit.role,
      urlLength: formAddEdit.urlLength
    }
    UserService.addUser(addUserParams).
      then(result => {
        if (result?.status == 200) {
          ElMessage.success(result.data)
          cleanAddUser()
          getUsersPage()
        }
      })
  }
  addUser()
}
const cleanAddUser = () => {
  formAddEdit.author = ''
  formAddEdit.autoInsertSpace = false
  formAddEdit.crt = ''
  formAddEdit.domain = ''
  formAddEdit.group = ''
  formAddEdit.i18n = 'English'
  formAddEdit.id = 0
  formAddEdit.name = ''
  formAddEdit.nickname = ''
  formAddEdit.phone = ''
  formAddEdit.pwd = ''
  formAddEdit.remarks = ''
  formAddEdit.role = 0
  formAddEdit.upt = ''
  formAddEdit.urlLength = 6
}
</script>

<style lang="scss" scoped>

</style>