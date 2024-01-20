<template>
    <qlhandle :el="el">
        <Qrcode :value="el.text" :level="el.level"  :width="el.width" />
    </qlhandle>
</template>

<script setup>
    import {ref,defineProps,computed,defineEmits,inject} from 'vue'
    import qlhandle from '../handle.vue'
    import Qrcode from '@/components/Qrcode.vue'
    
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