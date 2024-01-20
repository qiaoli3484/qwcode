<template>
    <qlhandle :el="el" formitem>
        <qlel v-model:data="data" :el="el"/>
    </qlhandle>
</template>

<!-- :formatter='el.formatter'
:parser='el.parser -->'

<script setup>
    import {ref,defineProps,computed,defineEmits,inject} from 'vue'
    import qlhandle from '../handle.vue'
    import qlel from './indexV.vue'
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