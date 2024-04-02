<template>
  <div class="common-layout">
    <el-container>
      <el-aside :width="width + 'px'">
        <qlMenu :Collapse="Collapse" :menuactivename="menuactivename" @select="menuSelect"></qlMenu>
      </el-aside>
      <el-container>
        <el-header>
          <qlHead @collapse="(v,w)=>{Collapse=v;width=w;}" @change="menuSelect" :activeIndex="menuactivename"></qlHead>
          </el-header>
        <el-main>
          <qlMain :navarray="navarray" @select="menuSelect" @romove="delnav"></qlMain>
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script setup>
import {ref,reactive} from "vue"
import qlMenu from "./components/ql-menu.vue"
import qlMain from "./components/ql-main.vue"
import qlHead from "./components/ql-head.vue"
const menuactivename=ref('首页|/home')
const Collapse=ref(false)
const width=ref(150)
const navarray=reactive([
        {
          name: '首页',
          path: '/home',
          showClose: false,
          check: true,
        },
      ])
const menuSelect=(index, indexPath)=>{
      if (index == '事务中心') {
        addnav('事务中心', '/flows')
      } else if (index == '消息中心') {
        addnav('消息中心', '/msgs')
      } else
        addnav(index,indexPath)
    }
const addnav=(name, path)=> {
      const exist = navarray.some((val) => {
        return val.name == name
      })
      if (!exist) {
        for (let i = navarray.length - 1; i >= 0; i--) {
          navarray[i].check = false
        }
        menuactivename.value = name+'|'+path
        navarray.push({
          name: name,
          path: path,
          showClose: true,
          check: true,
        })
      } else {
        checknav(name)
      }
    }
const checknav=(name)=>{
      let index = -1
      for (let i = navarray.length-1; i >= 0; i--) {
        if (navarray[i].name == name){index = i}
        navarray[i].check = false
      }
      menuactivename.value = navarray[index].name+'|'+navarray[index].path
      navarray[index].check=true
}
const delnav=(name)=>{
      let index = -1
      let check=false
      for (let i = navarray.length - 1; i >= 0; i--) {
        if (navarray[i].name == name){
          check=navarray[i].check
          index = i
          break
        }
      }
      navarray.splice(index, 1)
      if(!check){//关闭没选中的
       
      }else if(navarray.length > index){
        navarray[index].check = true
        menuactivename.value = navarray[index].name+'|'+navarray[index].path
      }else{
        navarray[index - 1].check = true
        menuactivename.value = navarray[index - 1].name+'|'+navarray[index - 1].path
      }
    }
</script>

<style>
html,
body {
  width: 100%;
  height: 100%;
}

* {
  padding: 0px 0px;
  margin: 0px 0px;
}

.el-main {
  padding: 0px!important;
  height: calc(100vh - 60px);
}

.el-notification__title {
  padding: 3.5px 0px;
}

.common-layout .el-aside {
    /*background-color: var(--el-color-primary-light-8);
    color: var(--el-text-color-primary);
    text-align: center;*/
    /*overflow: auto;*/
    overflow-y:auto;
    background-color: rgb(84, 92, 100);
}
.common-layout .el-header {
     line-height:60px
}


.content-box {
  position: absolute;
  top: 0;
  bottom: 0;
  left: 0;
  right: 0;
  overflow: auto;
  z-index: 1;
}

.head_box {
  height: 45px;
  background: #fff;
  box-shadow: 0 1px 4px rgba(0, 21, 41, 0.08);
}

.navbar-fixed-top {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  z-index: 1030;
}


.el-menu {
  border-right-width: 0;
}

.el-menu-item-group__title {
  padding: 0;
}

.el-menu-item,
.el-submenu__title {
  height: 30px!important;
  line-height: 30px!important;
  min-width: 64px!important; 
}
.el-sub-menu__title{
  height: 40px!important;
  line-height: 40px!important;
}

.el-submenu .el-menu-item {
  height: 30px;
  line-height: 30px;
  padding: 0 15px;
  min-width: 150px;
}

.el-menu--horizontal>.el-menu-item {
  height: 43px !important;
  line-height: 43px !important;
}
</style>
