<template>
  <el-form inline="true" label-position="left" label-width="68px" :model="form">
    <el-form-item label="源链接">
      <el-input v-model="form.sourceURL" />
    </el-form-item>
    <el-form-item label="短链接">
      <el-input v-model="form.shortURL" />
    </el-form-item>
    <el-form-item label="创建时间">
      <el-date-picker v-model="form.createdAt" type="daterange" unlink-panels range-separator="To"
        start-placeholder="Start date" end-placeholder="End date" />
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="onSubmit">Query</el-button>
    </el-form-item>
  </el-form>


  <el-table :data="tableData" border style="width: 100%" :default-sort="{ prop: 'date', order: 'descending' }"
    @selection-change="handleSelectionChange">
    <el-table-column type="selection" width="55" />
    <el-table-column prop="date" sortable label="源链接" />
    <el-table-column prop="name" sortable label="短链接" />
    <el-table-column prop="name" sortable label="创建时间" />
    <el-table-column prop="name" sortable label="修改时间" />
    <el-table-column prop="name" sortable label="分组" />
    <el-table-column label="过期时间"></el-table-column>
    <el-table-column label="启用">
      <el-switch inline-prompt active-text="是" inactive-text="否" />
    </el-table-column>

    <el-table-column prop="name" label="备注" />
    <el-table-column fixed="right" width="193">
      <template #header>
        <el-button size="small" @click="dialogVisible = true">新增</el-button>
        <el-button size="small">导入</el-button>
        <el-button size="small">删除</el-button>
      </template>
      <template #default>
        <el-button link type="primary" size="small">Detail</el-button>
        <el-button link type="primary" size="small">Edit</el-button>
        <el-button link type="primary" size="small">Delete</el-button>
      </template>
    </el-table-column>
  </el-table>


  <el-pagination background layout="total, sizes, prev, pager, next, jumper" :total="1000"
    :page-sizes="[100, 200, 300, 400]" v-model:currentPage="currentPage4" v-model:page-size="pageSize4"
    @size-change="handleSizeChange" @current-change="handleCurrentChange" hide-on-single-page="true" />


  <el-dialog v-model="dialogVisible" title="新增链接" :before-close="handleClose">
    <el-form label-position="left" label-width="79px">
      <el-form-item label="启用">
        <el-switch v-model="formAdd.isEnable" class="mt-2" style="margin-left: 24px" inline-prompt :active-icon="Check"
          :inactive-icon="Close" />
      </el-form-item>
      <el-form-item label="自动生成">
        <el-switch v-model="formAdd.isAutoGenerate" class="mt-2" style="margin-left: 24px" inline-prompt
          :active-icon="Check" :inactive-icon="Close" @click="testClick()" />
      </el-form-item>
      <el-form-item label="源链接">
        <el-input v-model="formAdd.sourceURL" />
      </el-form-item>
      <el-form-item label="短链接" v-if="!formAdd.isAutoGenerate">
        <el-input v-model="formAdd.sourceURL" />
      </el-form-item>
      <el-form-item label="URL长度" v-if="formAdd.isAutoGenerate">
        <el-input-number v-model="formAdd.urlLength" :min="4" :max="16" />
      </el-form-item>
      <el-form-item label="过期时间">
        <el-date-picker v-model="formAdd.expAt" type="datetime" placeholder="Select date and time" />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="addURL()">Submit</el-button>
        <el-button>Reset</el-button>
      </el-form-item>
    </el-form>
  </el-dialog>
</template>
<script lang="ts" setup>
import { ElContainer, ElHeader, ElRow, ElCol, ElMain } from 'element-plus';
import { defineComponent, ref, reactive } from 'vue'
import { Check, Close } from '@element-plus/icons-vue'

const addURL = () => {
  console.log(formAdd)
}

const testClick = () => {
  console.log(formAdd.isAutoGenerate)
}

const handleSelectionChange = (val) => {
  console.log(val)
}

const handleClose = (done: () => void) => {
  ElMessageBox.confirm('Are you sure to close this dialog?')
    .then(() => {
      done()
    })
    .catch(() => {
      // catch error
    })
}
const handleSizeChange = (val: number) => {
  console.log(`${val} items per page`)
}
const handleCurrentChange = (val: number) => {
  console.log(`current page: ${val}`)
}
const dialogVisible = ref(true)
const currentPage4 = ref(4)
const pageSize4 = ref(100)
const tableData = [
  {
    date: '2016-05-03',
    name: 'Tom',
    address: 'No. 189, Grove St, Los Angeles',
  },
  {
    date: '2016-05-02',
    name: 'Tom',
    address: 'No. 189, Grove St, Los Angeles',
  },
  {
    date: '2016-05-04',
    name: 'Tom',
    address: 'No. 189, Grove St, Los Angeles',
  },
  {
    date: '2016-05-01',
    name: 'Tom',
    address: 'No. 189, Grove St, Los Angeles',
  },
]
const input = ref('')
const onSubmit = () => {
  console.log('submit!')
}

const formAdd = reactive({
  sourceURL: '',
  isAutoGenerate: true,
  targetURL: '',
  remarks: '',
  urlLength: 6,
  isEnable: true,
  expAt: ''
})

const form = reactive({
  sourceURL: '',
  shortURL: '',
  createdAt: '',
  group: '',
  isEnable: true,
  isExp: false
})
</script>

<style  lang="scss" scoped>
.el-main {
  background-color: #1359a0;
}
</style>