<template>
    <qlhandle :el="el">
        <el-button :type="el.ctype" :size="el.csize" @click="onClick">{{el.label}}</el-button>
    </qlhandle>
</template>

<script setup>
    import {ref,defineProps,computed,defineEmits,inject} from 'vue'
    import qlhandle from '../handle.vue'
    const props=defineProps({
        el:Object,
        formData:Object
    })
    const emit=defineEmits(['update:el'])
    const onClick=()=>{
        eval(el.value.onClick)
    }
    
    const el=computed({
        get:()=>{
            return props.el
        },
        set:(val)=>{
            emit('update:el',val)
        }
    })
    const evil=(fn)=>{
            var Fn = Function;  //一个变量指向Function，防止有些前端编译工具报错
            return new Fn('return ' + fn)();
        }
</script>