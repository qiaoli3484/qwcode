<template>

<qlcode language="html" ref="ref_code" v-model:code="content" />
<el-button @click="reload">重载</el-button>
<el-button @click="reloading=false">关闭</el-button>
<el-button @click="reloading=true">打开</el-button>
<qlcustom :content="content" v-if="reloading"/>

</template>
<script setup>
import {ref,nextTick} from 'vue'
import qlcode from '@/components/qlcode/qlcode.vue'
import qlcustom from '@/components/customRenderComponent/customRenderComponent.vue'

const content=ref("<template><h2>aaa</h2></template>")
const ref_code=ref(null)
const reloading=ref(true)

const reload=()=>{
    reloading.value=false
    console.log(111)
    nextTick(()=>{
        setTimeout(()=>{
            reloading.value=true
            console.log(222)
       },2000)
    })
    /**/
}

nextTick(()=>{
    ref_code.value.setcode(content.value)
})


</script>