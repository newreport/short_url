<script lang="ts" setup>
import { reactive, ref } from 'vue'
const input = ref('')
const labelPosition = ref('top')


</script>


<template>
  <div class="common-layout">
    <el-container>
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
              <el-form 
              ref="ruleFormRef"  status-icon  label-width="120px"
                class="demo-ruleForm">
                <el-form-item label="Name">
                  <el-input  />
                </el-form-item>
                <el-form-item label="Password" prop="pass">
                  <el-input  type="password" autocomplete="off" />
                </el-form-item>
                <el-form-item>
                  <el-button type="primary">Submit</el-button>
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
