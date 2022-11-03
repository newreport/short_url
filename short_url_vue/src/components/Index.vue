<template>
  <div class="common-layout login-box">
    <el-container style="height:100%">
      <el-header>
        <el-radio-group v-model="isCollapse" style="margin-bottom: 20px">
          <el-radio-button :label="false" @click="menuOpen">expand</el-radio-button>
          <el-radio-button :label="true" @click="menuClose">collapse</el-radio-button>
        </el-radio-group>
        <select v-model="$i18n.locale">
          <option v-for="locale in $i18n.availableLocales" :key="`locale-${locale}`" :value="locale">
            {{ locale }}
          </option>
        </select>
        {{ $t("message.hello") }}
      </el-header>
      <el-container>
        <el-aside ref="hello" :style="{width:menuWidth}">
          <el-menu :collapse-transition="false" default-active="1-1" :collapse="isCollapse"
            @open="handleOpen" @close="handleClose">
            <el-sub-menu index="1">
              <template #title>
                <el-icon>
                  <House />
                </el-icon>
                <span>{{ $t("menu.menu1") }}</span>
              </template>
              <el-menu-item index="1-1" @click="changePage(shortURLPage)">
                <el-icon>
                  <Link />
                </el-icon>
                {{ $t("menu.menu1_1") }}
              </el-menu-item>
            </el-sub-menu>
            <el-sub-menu index="2">
              <template #title>
                <el-icon>
                  <Setting />
                </el-icon>
                <span>{{ $t("menu.menu2") }}</span>
              </template>
              <el-menu-item index="2-1" @click="changePage(userPage)">
                <el-icon><User /></el-icon>{{ $t("menu.menu2_1") }}</el-menu-item>
            </el-sub-menu>
          </el-menu>
        </el-aside>
        <el-main>
          <component :is="currentTab" class="tab"></component>
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

  
<script lang="ts" setup>
import shortURLPage from "@components/home/ShortURL.vue";
import userPage from "@components/setting/User.vue";
import { getCurrentInstance, ref, onMounted,shallowRef } from 'vue'
import {
  Document,
  Menu as IconMenu,
  Location,
  Setting,
} from '@element-plus/icons-vue'
let menuWidth = ref('200px')
const hello = ref<any>(null);

let tabs={
  shortURLPage,
  userPage
}
let currentTab=shallowRef(shortURLPage)


const isCollapse = ref(false)
const handleOpen = (key: string, keyPath: string[]) => {
  // console.log(key, keyPath)
  // console.log(hello.value)
  console.log("打开")
  // hello.value.clientWidth=6000
  // menuWidth.value='500px'
}
const handleClose = (key: string, keyPath: string[]) => {
  // console.log(key, keyPath)
  console.log("关闭")
  // console.log(hello.value)
  // hello.value.width='200px'
}

const changePage=(page)=>{
  currentTab.value=page
}

const menuOpen = () => {
  menuWidth.value = '200px'
}
const menuClose = () => {
  menuWidth.value = '63px'
}

const { proxy } = getCurrentInstance()
let c_width = ref("")
onMounted(() => {
  // c_width.value=proxy.$el.

})

// 组件中


</script>
  
<style lang="scss" scoped>
.el-menu-vertical-demo:not(.el-menu--collapse) {
  // width: 200px;
  min-height: 400px;
}
.tab {
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  border: 1px solid #ccc;
  padding: 10px;
}
.el-aside {
  background-color: #e0f194;
}

.el-header {
  background-color: #1359a0;
}
</style>