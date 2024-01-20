<template>
     <p :style="{textAlign:el.textalign,fontSize:el.fontsize,color:el.fontColor,fontWeight:el.fontWeight}">{{text}}</p>
</template>

<script setup>
    import {ref,defineProps,computed,defineEmits,inject} from 'vue'
    const props=defineProps({
        el:Object,
        data:String,
    })
    const formData =inject("formData")
    const text=computed(()=>{
        //eval(props.el.text)
        return  props.el.text
    })
</script>