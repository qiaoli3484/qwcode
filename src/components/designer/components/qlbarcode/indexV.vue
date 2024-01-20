<template>
    <Barcode :value="text"  :height="el.height"  :displayValue="el.displayValue"   />
</template>

<script setup>
    import {ref,defineProps,computed,defineEmits,inject} from 'vue'
    const props=defineProps({
        el:Object,
        data:String,
    })
    const formData =inject("formData")
    const text=computed(()=>{
        return eval(props.el.text)
    })
</script>