<template>
    <el-menu :collapse="props.Collapse" :default-active="props.menuactivename" class="el-menu-vertical-demo" @select="menuSelect"
        background-color="#545c64" text-color="#fff" unique-opened active-text-color="#ffd04b">
         <el-menu-item index="首页|/home">
            <el-icon><home-filled /></el-icon>
            <template #title>
                <span>首页</span>
            </template>
        </el-menu-item>
        <el-sub-menu :index="items.name" v-for="items in pagegroup" :key="items.name">
            <template #title>
                <el-icon>
                    <component :is="items.icon"></component>
                </el-icon>
                <span>{{ items.name }}</span>
            </template>
             
            <template v-for="item in items.children" :key="item.name">
                <el-sub-menu :index="item.name" v-if="item.type==0">
                    <template #title><span>{{ item.name }}</span></template>
                    <el-menu-item :index="ii.name+'|'+ii.path" v-for="ii in item.children" :key="ii.name" >{{ ii.name }}</el-menu-item>
                </el-sub-menu>
                <el-menu-item v-else :index="item.name+'|'+item.path"> {{ item.name }}</el-menu-item>
            </template>
           
            <!-- <el-menu-item-group>
                <el-menu-item :index="item.name" v-for="item in items.children" :key="item.name">
                    {{ item.name }}
                </el-menu-item>
            </el-menu-item-group> -->

        </el-sub-menu>
       
    </el-menu>
</template>

<script setup>
import {defineEmits,ref,onMounted,getCurrentInstance,defineProps} from "vue"
const {proxy} =getCurrentInstance();
const props=defineProps({
    Collapse:{
        type:Boolean,
        default:false
    },
    menuactivename:{
        type:String,
        default:"首页"
    }
})
const pagegroup=ref([])
const emit = defineEmits(['select'])
const menuSelect=(index, indexPath)=>{
    const arr= index.split('|')
    emit('select',arr[0],arr[1])
    //emit('select',pagegroup.value[indexPath[0]-1].children[index].name,pagegroup.value[indexPath[0]-1].children[index].path)
}
const hasPages=() =>{
      pagegroup.value.push(
        {
          name: '系统管理',
          icon: 'icon-ql-project-o',
          children: [
            { path: 'roles.html', name: '角色管理' },
            { path: 'users.html', name: '用户管理' },
            { path: 'config.html', name: '系统配置' },
            { path: 'folder.html', name: '页面配置' },
          ],
        })
      proxy.$get('index/folders').then((res) => {
        pagegroup.value = res.data
      })
    }
onMounted(()=>{
    hasPages()
})
</script>
<style scope>
.el-menu {
    z-index: 1;
}
</style>