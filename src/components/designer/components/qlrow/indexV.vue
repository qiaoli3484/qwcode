<template>
        <el-row :align="el.align" :gutter="el.gutter" :justify="el.justify">
            <qlcol v-for="col in el.childs"  :el="col"  v-model:data="data"/>
        </el-row>
</template>

<script setup>
    import {ref,defineProps,computed,defineEmits,inject} from 'vue'
    import qlcol from '../../components/qlcol/indexV.vue'

    const props=defineProps({
        el:Object,
        data:Object
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