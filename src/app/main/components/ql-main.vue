<script setup>
import {defineProps,defineEmits} from "vue"
const prpps=defineProps({
      navarray:Array
})
const emit= defineEmits(['romove','select'])
const checknav=(name,path) =>{
      emit('select',name,path)
    }
const delnav=(name) =>{
     emit('romove',name)
    }
</script>
<template>
  <div class="tabNav">
    <div class="tabNav-wp">
      <div class="navbtn" v-for="(item, index) in prpps.navarray" :key="index" :class="{ nav_check: item.check }"
        @click="checknav(item.name,item.path)">
        <el-icon><opportunity /></el-icon>
        <span>{{ item.name }}</span>
        <el-icon v-show="item.showClose" style="margin-left: 10px" @click.stop="delnav(item.name)" ><close /></el-icon>
      </div>
    </div>
  </div>
  <div class="Gcb-article">
    <iframe scrolling="yes" frameborder="0" :src="item.path" v-for="item in prpps.navarray" :key="item.name" v-show="item.check"></iframe>
  </div>
</template>

<style scoped>

.tabNav {
  height: 35px;
  position: relative;
  background: #fff;
  border-bottom: 1px solid #d8dce5;
  box-shadow: 0 1px 3px 0 rgba(0, 0, 0, 0.12), 0 0 3px 0 rgba(0, 0, 0, 0.04);
}

.tabNav-wp {
  position: absolute;
  left: 0px;
  width: 100%;
  height: 34px;
  word-wrap: break-word;
  white-space: nowrap;
  overflow: hidden;
}

.navbtn {
  display: inline-block;
  position: relative;
  cursor: pointer;
  height: 26px;
  line-height: 26px;
  border: 1px solid #d8dce5;
  color: #495060;
  background: #fff;
  padding: 0 8px;
  font-size: 12px;
  margin-left: 5px;
  margin-top: 4px;
  border-radius: 2px;
  -webkit-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
  user-select: none;
}

.nav_check {
  background-color: #42b983;
  color: #fff;
  border-color: #42b983;
}


.Gcb-article {
  height: calc(100vh - 96px);
  width: 100%;
  overflow: auto;
  z-index: 1;
}

.Gcb-article iframe {
  height: calc(100vh - 100px);
  width: 100%;
  /*background-image: url(/static/images/[name].png);*/
}

.el-menu {
  border-right-width: 0;
}

.el-menu-item-group__title {
  padding: 0;
}

.el-menu-item,
.el-submenu__title {
  height: 40px;
  line-height: 40px;
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
