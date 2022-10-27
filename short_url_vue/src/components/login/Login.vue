<script lang="ts" setup>
import { h, reactive, ref } from 'vue'
import { ElNotification } from 'element-plus'
import {UserService} from '../../api/api'
import axios from 'axios';
const name = ref('')
const pwd = ref('')


const open1 = () => {
  console.log('open1 VITE_APP_TITLE:' + import.meta.env.VITE_API_DOMAIN);

  axios.get('http://127.0.0.1:7777/v1/user', {
  })
  .then(function (response) {
    console.log(response.data);
    let ress= UserService.getall()
  console.log(ress)
  })
  .catch(function (error) {
    console.log(error);
  });
  // ElNotification({
  //   title: 'Title',
  //   message: h('i', { style: 'color: teal' }, name.value + pwd.value),
  // })

  // const login1 = async () => {
  // let loginParams = {
  //   name: 'admin',
  //   pwd: 'admin',
  // }
  // console.log("调用了api")
  // let res = await UserService.login1(loginParams)
  // console.log(res)
// }
// login1()
}
</script>


<template>
  <div class="login-box">
    <el-container style="height:100%">
      <el-header>url link</el-header>
      <el-main>
        <el-row>
          <el-col :span="1"></el-col>
          <el-col :xs="24" :sm="24" :md="14">
            <el-image src="https://cube.elemecdn.com/6/94/4d3ea53c084bad6931a56d5158a48jpeg.jpeg">
              <template #placeholder>
                <div class="image-slot">Loading<span class="dot">...</span></div>
              </template>
            </el-image>
          </el-col>
          <el-col :span="1"></el-col>

          <el-col :span="6">

            <el-card class="box-card">
              <template #header>
                <div class="card-header">
                  <span>Login</span>
                </div>
              </template>
              <el-form ref="ruleFormRef" status-icon label-width="120px" class="demo-ruleForm">
                <el-form-item label="Name">
                  <el-input v-model="name" />
                </el-form-item>
                <el-form-item label="Password" prop="pass">
                  <el-input v-model="pwd" type="password" autocomplete="off" />
                </el-form-item>
                <el-form-item>
                  <el-button type="primary" @click="open1">Submit</el-button>
                </el-form-item>
              </el-form>

            </el-card>
          </el-col>

          <el-col :span="2"></el-col>

        </el-row>
      </el-main>
      <el-footer>
      </el-footer>
    </el-container>
  </div>
</template>



<style lang="scss" scoped>
.login-box {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.text {
  font-size: 14px;
}

.item {
  margin-bottom: 18px;
}

.box-card {
  width: 480px;
}

.el-footer {
  background-color: #1359a0;
}

.el-header {
  background-color: #1359a0;
}

.el-row {
  margin-bottom: 20px;
}

.el-row:last-child {
  margin-bottom: 0;
}

.el-col {
  border-radius: 4px;
}

.grid-content {
  border-radius: 4px;
  min-height: 36px;
}
</style>
