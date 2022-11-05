<template>
  <el-form inline="true" label-position="left" label-width="68px" :model="form1">
    <el-form-item label="用户名">
      <el-input v-model="form1.name" />
    </el-form-item>
    <el-form-item label="昵称">
      <el-input v-model="form1.nickname" />
    </el-form-item>
    <el-form-item label="手机号">
      <el-input v-model="form1.phone" />
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="query()">Query</el-button>
    </el-form-item>
  </el-form>
</template>
  
<script lang="ts" setup>
import { ref, reactive } from 'vue'

const form1 = reactive({
  name: '',
  nickname: '',
  phone: ''
})

const query = () => {
}

</script>

<style lang="scss" scoped>

</style>