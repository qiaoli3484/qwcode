<template>
    <el-image :style="{width:el.width,height:el.height}" :src="el.url" :fit="el.fit" />
</template>

<script setup>
    import {ref,defineProps,computed,defineEmits,inject,getCurrentInstance} from 'vue'
    import {ElMessage} from 'element-plus'
    const {proxy}=getCurrentInstance()
    const props=defineProps({
        el:Object,
        data:String,
    })
    const emit=defineEmits(['update:data'])
</script>