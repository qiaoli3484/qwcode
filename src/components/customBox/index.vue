<template>

<qlcode language="html" ref="ref_code" v-model:code="content" />
<el-button @click="reload">重载</el-button>
<qlcustom :content="content" v-if="reloading"  :auto="false"  />

</template>
<script setup>
import {ref,nextTick} from 'vue'
import qlcode from '@/components/qlcode/qlcode.vue'
//import qlcustom from '@/components/customRenderComponent/customRenderComponent.vue'

const content=ref("myComponent.vue")
const ref_code=ref(null)
const reloading=ref(true)

const reload=()=>{
    try {
    reloading.value=false
    console.log(111)
    /*nextTick(()=>{*/
        console.log(222)
        setTimeout(()=>{
            reloading.value=true
            console.log(333)
       },2000)
    /*})*/
    /**/
    }catch(err){
        console.error(err);
    }
}

nextTick(()=>{
    ref_code.value.setcode(content.value)
})


</script>