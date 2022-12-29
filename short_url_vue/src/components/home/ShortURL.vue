<template>
  <div>
    <el-form :inline="true" label-position="left" label-width="68px" :model="form">
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
        <el-button type="primary" @click="getShortsPage">Query</el-button>
      </el-form-item>
    </el-form>


    <el-table :data="tableData" border style="width: 100%" :default-sort="{ prop: 'sourceURL', order: 'descending' }"
      @selection-change="handleSelectionChange" row-key="id">
      <el-table-column type="selection" width="55" />
      <el-table-column prop="sourceURL" sortable label="源链接" />
      <el-table-column prop="targetURL" sortable label="短链接" />
      <el-table-column sortable label="创建时间">
        <template #default="scope">
          {{ scope.row.crt == "0001-01-01T00:00:00Z" ? "" : moment(scope.row.crt).format("YYYY-MM-DD HH:mm:ss") }}
        </template>
      </el-table-column>
      <el-table-column prop="upt" sortable label="修改时间">
        <template #default="scope">
          {{ scope.row.upt == "0001-01-01T00:00:00Z" ? "" : moment(scope.row.upt).format("YYYY-MM-DD HH:mm:ss") }}
        </template>
      </el-table-column>
      <el-table-column prop="shortGroup" sortable label="分组" />
      <el-table-column label="过期时间">
        <template #default="scope">
          {{ scope.row.exp == "0001-01-01T00:00:00Z" ? "" : moment(scope.row.exp).format("YYYY-MM-DD HH:mm:ss") }}
        </template>

      </el-table-column>
      <el-table-column label="启用">
        <template #default="scope">
          <el-switch v-model="scope.row.isEnable" inline-prompt active-text="是" inactive-text="否" />
        </template>
      </el-table-column>

      <el-table-column prop="remarks" label="备注" />
      <el-table-column fixed="right" width="193">
        <template #header>
          <el-button size="small" @click="dialogVisible = true; cleanAddShort();">新增</el-button>
          <el-button size="small">导入</el-button>
          <el-button size="small">删除</el-button>
        </template>
        <template #default="scope">

          <el-button link type="primary" size="small">Detail</el-button>

          <el-button link type="primary" size="small" @click="dialogVisible = true;
formAddEdit.id = scope.row.id;
formAddEdit.sourceURL = scope.row.sourceURL;
formAddEdit.targetURL = scope.row.targetURL;
formAddEdit.urlLength = scope.row.targetURL.length;
formAddEdit.isEnable = scope.row.isEnable;
formAddEdit.exp = scope.row.exp == '0001-01-01T00:00:00Z' ? null : scope.row.exp;
formAddEdit.shortGroup = scope.row.shortGroup;
          ">Edit</el-button>
          <el-popconfirm title="确定删除吗?" @confirm="deleteShort(scope.row.id)">
            <template #reference>
              <el-button link type="primary" size="small">Delete</el-button>
            </template>
          </el-popconfirm>
        </template>
      </el-table-column>
    </el-table>


    <el-pagination background layout="total, sizes, prev, pager, next, jumper" :total="page.count"
      :page-sizes="[100, 200, 300, 400]" v-model:currentPage="page.currentPage" v-model:page-size="page.pageSize"
      @size-change="handleSizeChange" @current-change="handleCurrentChange" :hide-on-single-page="true" />


    <el-dialog v-model="dialogVisible" :title="formAddEdit.id ? '修改链接' : '新增链接'" :before-close="handleClose">
      <el-form label-position="left" label-width="79px">
        <el-form-item label="启用">
          <el-switch v-model="formAddEdit.isEnable" class="mt-2" style="margin-left: 24px" inline-prompt
            :active-icon="Check" :inactive-icon="Close" />
        </el-form-item>
        <el-form-item label="自动生成">
          <el-switch v-model="formAddEdit.isAutoGenerate" class="mt-2" style="margin-left: 24px" inline-prompt
            :active-icon="Check" :inactive-icon="Close" />
        </el-form-item>
        <el-form-item label="源链接">
          <el-input v-model="formAddEdit.sourceURL" />
        </el-form-item>
        <el-form-item label="分组">
          <el-input v-model="formAddEdit.shortGroup"/>
        </el-form-item>
        <el-form-item label="短链接" v-if="!formAddEdit.isAutoGenerate">
          <el-input v-model="formAddEdit.targetURL" />
        </el-form-item>
        <el-form-item label="URL长度" v-if="formAddEdit.isAutoGenerate">
          <el-input-number v-model="formAddEdit.urlLength" :min="4" :max="16" />
        </el-form-item>
        <el-form-item label="过期时间">
          <el-date-picker v-model="formAddEdit.exp" type="datetime" placeholder="Select date and time" />
        </el-form-item>
        <el-form-item label="备注">
          <el-input v-model="formAddEdit.remarks" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="formAddEdit.id ? updateShort() : addShort();"> Submit</el-button>
          <el-button @click="cleanAddShort()">Reset</el-button>
        </el-form-item>
      </el-form>
    </el-dialog>
  </div>
</template>
<script lang="ts" setup>
import { ElContainer, ElHeader, ElRow, ElCol, ElMain, ElMessageBox, ElMessage } from 'element-plus';
import { getCurrentInstance, defineComponent, ref, reactive } from 'vue'
import { Check, Close } from '@element-plus/icons-vue'
import { ShortService } from '@/api/api'

const { appContext } = getCurrentInstance();
const moment = appContext.config.globalProperties.$moment;

const form = reactive({
  sourceURL: '',
  shortURL: '',
  createdAt: '',
  shortGroup: '',
  isEnable: true,
  isExp: false
})

const page = reactive({
  currentPage: 0,
  pageSize: 100,
  count: 10
})

let tableData = ref()

const getShortsPage = () => {
  const getShortsPage = async () => {
    const getShortsPageParams = {
      offset: page.currentPage,
      limit: page.pageSize,
      sort: "",
      source_url: form.sourceURL.length == 0 ? "" : "lk" + form.sourceURL,
      target_url: form.shortURL.length == 0 ? "" : "lk" + form.shortURL,
    }
    // ShortService.getShortsPage(getShortsPageParams).then(result => {
    //   if (result?.status == 200) {
    //     page.count = result.data.count
    //     tableData.value = result.data.data
    //   }
    // })
  }
  getShortsPage()
}
getShortsPage()

const dialogVisible = ref(false)

const formAddEdit = reactive({
  id: '',
  sourceURL: '',
  isAutoGenerate: true,
  targetURL: '',
  remarks: '',
  urlLength: 6,
  isEnable: true,
  exp: '',
  shortGroup: ''
})

const addShort = () => {
  dialogVisible.value = false
  const addShort = async () => {
    const addShortParams = {
      sourceURL: formAddEdit.sourceURL,
      automactic:formAddEdit.isAutoGenerate,
      length: formAddEdit.urlLength,
      shortGroup: formAddEdit.shortGroup,
      targetURL:  formAddEdit.targetURL,
      isEnable: formAddEdit.isEnable,
      remarks: formAddEdit.remarks,
      exp:formAddEdit.exp
    }
    ShortService.createShort(addShortParams).
      then(result => {
        if (result?.status == 200) {
          ElMessage.success(result.data)
          cleanAddShort()
          getShortsPage()
        }
      })
  }
  addShort()
}

const updateShort = () => {
  dialogVisible.value = false
  const updateShort = async () => {
    let updateShortParams=tableData.value.find(x=>x.id==formAddEdit.id)
    updateShortParams.automactic=formAddEdit.isAutoGenerate
    updateShortParams.length=formAddEdit.urlLength
    updateShortParams.targetURL=formAddEdit.targetURL
    updateShortParams.shortGroup=formAddEdit.shortGroup
    updateShortParams.isEnable=formAddEdit.isEnable
    updateShortParams.remarks=formAddEdit.remarks
    updateShortParams.exp=formAddEdit.exp
    ShortService.updateShort(updateShortParams)
      .then(result => {
        if (result?.status == 200) {
          ElMessage.success(result.data)
          cleanAddShort()
          getShortsPage()
        }
      }).catch(err => {

      })
  }
  updateShort()
}

const deleteShort = (id: string) => {
  const deleteShort = async () => {
    const deleteShortParams = {
      id: id
    }
    console.log(deleteShortParams)
    ShortService.deleteShort(deleteShortParams).then(result => {
      if (result?.status == 200) {
        ElMessage.success(result.data)
        getShortsPage()
      }
    })
  }
  deleteShort()
}

const cleanAddShort = () => {
  formAddEdit.id = '',
  formAddEdit.sourceURL = ''
  formAddEdit.isAutoGenerate = true
  formAddEdit.targetURL = ''
  formAddEdit.remarks = ''
  formAddEdit.urlLength = 6
  formAddEdit.isEnable = true
  formAddEdit.exp = ''
  formAddEdit.shortGroup = ''
}

const handleSelectionChange = (val) => {
  console.log(val)
}

const handleClose = (done: () => void) => {
  // ElMessageBox.confirm('Are you sure to close this dialog?')
  //   .then(() => {
  done()
  //   })
  //   .catch(() => {
  //     // catch error
  //   })
}

const handleSizeChange = (val: number) => {
  console.log(`${val} items per page`)
}
const handleCurrentChange = (val: number) => {
  console.log(`current page: ${val}`)
}






</script>

<style  lang="scss" scoped>
.el-main {
  background-color: #1359a0;
}
</style>