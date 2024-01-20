<template>
    <qlhandle :el="el">
        <div :style="{ height: el.height,width:el.width}" style="border:1px solid #1b1b1b;"></div>
    </qlhandle>
</template>

<script setup>
    import {ref,defineProps,computed,defineEmits,inject} from 'vue'
    import qlhandle from '../handle.vue'

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