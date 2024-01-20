<template>
    <qlhandle :el="el">
        <el-image :style="{width:el.width,height:el.height}" :src="el.url" :fit="el.fit" />
    </qlhandle>
</template>

<!-- :formatter='el.formatter'
:parser='el.parser -->'

<script setup>
    import {ref,defineProps,computed,defineEmits,inject} from 'vue'
    import qlhandle from '../handle.vue'
    const fits = ['fill', 'contain', 'cover', 'none', 'scale-down']
    const data=ref('')
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