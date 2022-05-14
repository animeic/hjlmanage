<template>
  <el-container class="layout-container-demo" style="height: 500px">
    <el-aside width="200px">
       <el-menu
        :default-active="activeIndex"
        class="aside-menu"
        @select="handleSelect"
      >
        <el-menu-item index="1"><el-icon><Grid /></el-icon>数据看板</el-menu-item>
        <el-menu-item index="2"><el-icon><Goods /></el-icon>品牌</el-menu-item>
        <el-menu-item index="3"><el-icon><Management /></el-icon>商品管理</el-menu-item>
        <el-menu-item index="4"><el-icon><Bottom /></el-icon>商品入库明细</el-menu-item>
        <el-menu-item index="5"><el-icon><Top /></el-icon>商品出口明细</el-menu-item>
        <el-sub-menu index="6">
          <template #title><el-icon><Folder /></el-icon>中间表</template>
          <el-menu-item index="6-1"><el-icon><ArrowDown /></el-icon>入库</el-menu-item>
          <el-menu-item index="6-2"><el-icon><ArrowUp /></el-icon>出库</el-menu-item>
        </el-sub-menu>
      </el-menu>
    </el-aside>

    <el-container>
      <el-header style="text-align: right; font-size: 12px">
      </el-header>

      <el-main>
      </el-main>
    </el-container>

  </el-container>
</template>

<script lang="ts" setup>
import { ref,onMounted } from 'vue'
import {
  Grid,
  Goods,
  Management,
  Bottom,
  Top,
  Folder,
  ArrowDown,
  ArrowUp,
  } from '@element-plus/icons-vue'

const activeIndex = ref('1')
const handleSelect = (key: string, keyPath: string[]) => {
  console.log(key, keyPath)
  alert(key)
}
const user = useVisitor()
const name = ref('');
const type = ref(1);

onBeforeMount(() => {
  if(user.value == ""){ 
    toLogin()
  }else{
    toHome()
  }
})


const toLogin = () => {
  navigateTo({
    path: '/login',
    query: {
      name: name.value,
      type: type.value
    }
  })
}
const toHome = () => {
  navigateTo({
    path: '/home',
  })
}

// function navigate(){
//   return navigateTo({
//     path: '/login',
//     query: {
//       name: name.value,
//       type: type.value
//     }
//   })
// }


</script>

<style scoped>
.layout-container-demo .el-header {
  position: relative;
  background-color: var(--el-color-primary-light-7);
  color: var(--el-text-color-primary);
}
.layout-container-demo .el-aside {
  color: var(--el-text-color-primary);
  background: var(--el-color-primary-light-8);
}
.layout-container-demo .el-menu {
  border-right: none;
}
.layout-container-demo .el-main {
  padding: 0;
}
.layout-container-demo .toolbar {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  height: 100%;
  right: 20px;
}
</style>
