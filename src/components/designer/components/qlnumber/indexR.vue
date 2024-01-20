<template>
     <qlhandle :el="el" :widget="widget" formitem>
        <el-input-number v-model.trim="data" :placeholder="widget.placeholder" :precision="widget.bits" :controls="false" :readonly="widget.readonly"/>
    </qlhandle>
</template>

<!-- :formatter='el.formatter'
:parser='el.parser -->'

<script setup>
    import {ref,defineProps,computed,defineEmits,inject} from 'vue'
    import qlhandle from '../handle.vue'

    const data=ref('')
    const props=defineProps({
        el:Object,
        formData:Object
    })
    const emit=defineEmits(['update:el'])
    const widgets=inject("widgets")
    const widget=computed(()=>{
       return widgets.value[props.el.fid]
    })
    const el=computed({
        get:()=>{
            return props.el
        },
        set:(val)=>{
            emit('update:el',val)
        }
    })
    const elSelect=inject("elSelect")
    const onselect=()=>{
        el.value.selected=true
        elSelect(el.value)
    }
</script>