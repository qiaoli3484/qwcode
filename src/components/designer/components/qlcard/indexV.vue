<template>
    <el-input v-model.trim="data" :placeholder="el.placeholder" :readonly="el.readonly"/>
</template>

<!-- :formatter='el.formatter'
:parser='el.parser -->'

<script setup>
    import {ref,defineProps,computed,defineEmits,inject} from 'vue'
    const props=defineProps({
        el:Object,
        data:{type:String,default:''},
    })
    const emit=defineEmits(['update:data'])
    
    const data=computed({
        get:()=>{
            return props.data
        },
        set:(val)=>{
            emit('update:data',val)
        }
    })
</script>