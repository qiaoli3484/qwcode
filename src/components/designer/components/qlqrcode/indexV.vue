<template>
     <Qrcode :value="text" :level="el.level"  :width="el.width" />
</template>

<script setup>
    import {ref,defineProps,computed,defineEmits,inject} from 'vue'
    import Qrcode from '@/components/Qrcode.vue'
    const props=defineProps({
        el:Object,
        data:String,
    })
    const formData =inject("formData")
    const text=computed(()=>{
        return eval(props.el.text)
    })
</script>