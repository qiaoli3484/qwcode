
<template>
    <el-button @click="aaaa">啊啊啊</el-button>
    <el-input v-model.trim="data1" v-bind="attr"></el-input>
</template>
<script setup>
import {ref, watch,defineEmits, defineProps,getCurrentInstance,onMounted,onUnmounted} from 'vue';
    const {proxy} =getCurrentInstance();
    const emit=defineEmits(["change"])
    const props=defineProps({
        value:[String, Number],
        placeholder:String,
    })

    const attr=ref({
        size:"small",
        autofocus:true,
    })

    const data = defineModel('data')//新方法

    const data1=ref("测试22")
    const aaaa=()=>{
        proxy.$get("aaa/bbb")
    }

    const outClick=(e)=>{
        /*let dropRef= this.$el
        if(!dropRef.contains(e.target)&&visible.value){
            visible.value=false
        }*/
    }

    onMounted(()=>{
        document.addEventListener('click',outClick)
    })
    onUnmounted(()=>{
        document.removeEventListener('click',outClick)
    })

</script>
