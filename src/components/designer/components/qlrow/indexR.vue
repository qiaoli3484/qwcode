<template>
    <handle :el="el" style="outline: 1px dashed #336699;">
        <el-row :align="el.align" :gutter="el.gutter" :justify="el.justify">
            <qlcol v-for="(col,index) in el.childs"  v-model:el="el.childs[index]" v-model:parent="el"  :num="index"/>
        </el-row>
    </handle>
</template>

<script setup>
    import {ref,defineProps,computed,defineEmits} from 'vue'
    import qlcol from '../../components/qlcol/indexR.vue'
    import handle from '../../components/handle.vue'

    const props=defineProps({
        el:Object,
        formData:Object
    })
    const emit=defineEmits(['update:el'])
    
    const el=computed({
        get:()=>{
            return props.el
        },
        set:(val)=>{
            emit('update:el',val)
        }
    })
</script>