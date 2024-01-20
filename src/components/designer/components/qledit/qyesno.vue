<template>
    <span style="font-size:16px;" :style="{color:props.data==1?'#1b5':'#e62'}">{{props.data==1?'√':'Х'}}</span>
</template>

<script setup>
    import {ref,defineProps,computed,inject} from 'vue'
    const props=defineProps({
        data:Number,
        el:Object,
        widget:Object,
    })
</script>