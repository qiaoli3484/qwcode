<template>
    <qlhandle :el="el">
        <Barcode :value="el.text"  :height="el.height"  :displayValue="el.displayValue"   />
    </qlhandle>
</template>

<script setup>
    import {ref,defineProps,computed,defineEmits,inject} from 'vue'
    import qlhandle from '../handle.vue'
    import Barcode from '@/components/Barcode.vue'
    
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