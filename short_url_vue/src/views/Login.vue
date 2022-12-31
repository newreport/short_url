<script lang="ts" setup>
import router from '@/router'
import { defineComponent, getCurrentInstance, h, reactive, ref, computed } from 'vue'
import { ElMessage, ElButton, ElCard, ElCol, ElContainer, ElFooter, ElForm, ElFormItem, ElHeader, ElImage, ElInput, ElMain, ElNotification, ElRow } from 'element-plus'
import { UserService } from '@/api/api'
import { useStore } from 'vuex' // 引入useStore 方法

const { appContext } = getCurrentInstance();
const base64 = appContext.config.globalProperties.$base64;
const store = useStore()  // 该方法用于返回store 实例

const name = ref('')
const pwd = ref('')

// console.log('open1 VITE_APP_TITLE:' + import.meta.env.VITE_API_DOMAIN);


const login = () => {
  const login = async () => {
    const loginParams = {
      name: name.value,
      pwd: pwd.value,
    }
    UserService.login(loginParams)//登录
      .then(result => {
        if (result?.status == 200) {
          store.commit('refreshToken', result.data)
          UserService.refreshToken(result.data).then(result => {//刷新refershToken
            if (result?.status == 200) {
              store.commit('accessToken', result.data)
            }
          })
          router.push({ path: '/index' })
        }
      }).catch(err => {
        store.commit('cleanToken')
        ElMessage.error(err)
      })
  }
  login()
}


</script>


<template>
  <div class="login-box">

    <el-container style="height:100%">
      <el-header>
      </el-header>
      <el-main>
        <el-row :gutter="24">
          <el-col class="hidden-xs-only" :xl="{ span: 15, offset: 1 }" :lg="{ span: 14, offset: 1 }" :md="{ span: 14 }"
            :sm="13" :xs="24">
            <el-image src="https://cube.elemecdn.com/6/94/4d3ea53c084bad6931a56d5158a48jpeg.jpeg">
            </el-image>
          </el-col>
          <el-col :xl="{ span: 7 }" :lg="{ span: 8 }" :md="{ span: 10 }" :sm="11" :xs="{ span: 20, offset: 2 }">
            <el-card class="box-card"><template #header>
                <div class="card-header"><span> {{ $t("message.hello") }}</span></div>
              </template>
              <el-form class="demo-ruleForm" ref="ruleFormRef" status-icon>
                <el-form-item label="Name">
                  <el-input v-model="name"></el-input>
                </el-form-item>
                <el-form-item label="Password" prop="pass">
                  <el-input v-model="pwd" type="password" @keyup.enter.native="login" autocomplete="off"></el-input>
                </el-form-item>
                <el-form-item>
                  <el-button type="primary" @click="login">Submit</el-button>
                </el-form-item>
              </el-form>
            </el-card>
          </el-col>
        </el-row>
      </el-main>
      <el-footer></el-footer>
    </el-container>
  </div>
</template>

<style lang="scss" >
.login-box {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
}
</style>

<style lang="scss" scoped>
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

// .box-card {
//   width: 480px;
// }

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
