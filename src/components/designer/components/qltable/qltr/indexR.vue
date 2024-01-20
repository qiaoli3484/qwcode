<template>
    <tr>
        <qltd v-for="(col,index) in el.childs"  v-model:el="el.childs[index]" :cnum="index" :rnum="props.rnum"/>
    </tr>
</template>

<script setup>
    import {ref,defineProps,computed,defineEmits} from 'vue'
    import qltd from '../qltd/indexR.vue'

    const props=defineProps({
        el:Object,
        formData:Object,
        rnum:Number
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