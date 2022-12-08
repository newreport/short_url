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
    <el-table-column prop="role" sortable label="权限" />
    <el-table-column prop="remarks" label="备注" />
    <el-table-column fixed="right" width="193">
        <template #header>
          <el-button size="small">新增</el-button>
        </template>
        <template #default="scope">

          <el-button link type="primary" size="small">Detail</el-button>

          <el-button link type="primary" size="small" >Edit</el-button>
          <el-popconfirm title="确定删除吗?" >
            <template #reference>
              <el-button link type="primary" size="small">Delete</el-button>
            </template>
          </el-popconfirm>
        </template>
      </el-table-column>
  </el-table>
</div>
</template>
  
<script lang="ts" setup>
import { ref, reactive } from 'vue'
import { UserService } from '@/api/api'

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
  const getUsersPage= async ()=>{
    const getUsersPageParams ={
      offset: page.currentPage,
      limit: page.pageSize,
      sort: "",
      name:form.name.length==0?"":"lk"+form.name,
      nickname:form.nickname.length==0?"":"lk"+form.nickname,
      phone:form.phone.length==0?"":"lk"+form.phone
    }
    UserService.getUsersPage(getUsersPageParams).then(result=>{
      if(result?.status==200){
        page.count = result.data.count
        tableData.value = result.data.data
      }
    })
  }
  getUsersPage()
}

</script>

<style lang="scss" scoped>

</style>