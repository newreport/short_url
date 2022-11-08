<template>
  <div>
  <el-form inline="true" label-position="left" label-width="68px" :model="form">
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
      <el-button type="primary" @click="query()">Query</el-button>
    </el-form-item>
  </el-form>

  <el-table :data="table" border style="width: 100%">
    <el-table-column type="selection" width="55" />
    <el-table-column prop="name" sortable label="用户名" />
    <el-table-column prop="group" sortable label="分组" />
    <el-table-column prop="nickname" sortable label="昵称" />
    <el-table-column prop="urlLength" sortable label="默认长度" />
    <el-table-column prop="phone" sortable label="电话" />
    <el-table-column prop="remarks" sortable label="备注" />
  </el-table>
</div>
</template>
  
<script lang="ts" setup>
import { ref, reactive } from 'vue'

const form = reactive({
  name: '',
  nickname: '',
  phone: '',
})

const table = [{
  name: 'admin',
  nickname: 'admin',
  urlLength: 6,
  phone: '1234567890',
  remarks: '管理员'
}]

const query = () => {
}

</script>

<style lang="scss" scoped>

</style>