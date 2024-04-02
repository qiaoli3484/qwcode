<template>
    <div style="float: left;" @click="icochange">
      <div style="cursor: pointer; font-size: 20px">
        <el-icon>
            <component :is="elicon"></component>
          <Fold />
        </el-icon>
      </div>
    </div>
   <div style="float: left">
      <el-menu :default-active="prop.activeIndex" class="el-menu-demo" mode="horizontal" @select="menuSelect">
        <el-menu-item index="事务中心">事务中心</el-menu-item>
        <el-menu-item index="消息中心">消息中心</el-menu-item>
        <el-menu-item index="3" disabled>待添加</el-menu-item>
      </el-menu>
    </div>

    <!-- <el-badge v-show="noRead > 0" :model-value="noRead" class="mark" style="cursor: pointer"></el-badge> -->

    <div style="float: right;width:150px;height: 60px; cursor: pointer;display: flex;flex-direction: row;align-items: center;justify-content: space-around;">
          <el-tooltip class="item" effect="dark" :content="tiptext" placement="bottom">
            <el-icon @click="changescreen"><FullScreen /></el-icon>
          </el-tooltip>

          <el-dropdown size="small" @command="userCommand">
            <span class="el-dropdown-link">
              {{name}}
              <el-icon class="el-icon--right">
                <arrow-down />
              </el-icon>
            </span>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item command="a">个人信息</el-dropdown-item>
                <el-dropdown-item command="signOut" divided>退出登录</el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>

            <el-popover placement="bottom" :width="200" trigger="hover">
              <template #reference>
                  <el-button link>在线</el-button>
              </template>
              <div style="height: 80px; width: 150px; overflow: hidden">
                <ul style="overflow: auto; height: 80px; width: 190px">
                  <li v-for="item in online" style="display: inline-block; width: 75px" :key="item.id">
                    <el-icon style="color: chocolate"><User /></el-icon>{{ item.name }}
                  </li>
                </ul>
              </div>
            </el-popover>
    </div>
  <qlwsmsg v-model:name="name" v-model:online="online" />
</template>

<script setup>
import qlwsmsg from "./ql-wsmsg.vue"
import {ref,onMounted,defineEmits,defineProps} from "vue" 
const prop=defineProps({
      noRead:Number,
      activeIndex:String
})
const emit= defineEmits(['collapse','change'])

const name=ref('')
const online=ref([])
const menuSelect=(key, keyPath)=>{
      emit('change',key, keyPath)
}

const tiptext=ref('显示全屏')
const fullscreen=ref(false)
const elicon=ref('fold')
const changescreen=()=>{
      fullscreen.value = !fullscreen.value
      if (fullscreen.value) {
        fullScreen()
      } else {
        exitScreen()
      }
    }
const fullScreen=()=>{
      const el = document.documentElement;
      const rfs =el.requestFullScreen ||el.webkitRequestFullScreen ||el.mozRequestFullScreen ||el.msRequestFullscreen;
      if (typeof rfs != 'undefined' && rfs) {
        rfs.call(el)
      }
      return
    }
 const exitScreen=()=>{
      if (document.exitFullscreen) {
        document.exitFullscreen()
      } else if (document.mozCancelFullScreen) {
        document.mozCancelFullScreen()
      } else if (document.webkitCancelFullScreen) {
        document.webkitCancelFullScreen()
      } else if (document.msExitFullscreen) {
        document.msExitFullscreen()
      }
    }

const icochange=()=>{
    if(elicon.value=='fold'){
       elicon.value='expand'
       emit('collapse',true,64)
    }else{
       elicon.value='fold'
       emit('collapse',false,150)
    }
}
const userCommand=(command)=>{
      if (command == 'signOut') {
        document.cookie = '' ////清除cookie
        window.location.href = '/'
        window.close()
        return
      }
    }
onMounted(()=>{
    document.addEventListener('fullscreenchange', () => {
      if (tiptext.value == '退出全屏') {
        tiptext.value = '显示全屏'
        fullscreen.value = false
        return
      }
      tiptext.value = '退出全屏'
    })
    document.addEventListener('webkitfullscreenchange', () => {
      if (tiptext.value == '退出全屏') {
        tiptext.value = '显示全屏'
        fullscreen.value = false
        return
      }
      tiptext.value = '退出全屏'
    })
    document.addEventListener('mozfullscreenchange', () => {
      if (tiptext.value == '退出全屏') {
        tiptext.value = '显示全屏'
        fullscreen.value = false
        return
      }
      tiptext.value = '退出全屏'
    })
    document.addEventListener('MSFullscreenChange', () => {
      if (tiptext.value == '退出全屏') {
        tiptext.value = '显示全屏'
        fullscreen.value = false
        return
      }
      tiptext.value = '退出全屏'
    })
})
</script>