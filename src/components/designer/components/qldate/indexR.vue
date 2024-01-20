<template>
     <qlhandle :el="el" :widget="widget" formitem>
        <el-date-picker v-if="dateType!='time'" v-model.trim="data" :type="dateType" size="small" :value-format="getformat()" clearable :readonly="widget.readonly" :placeholder="widget.placeholder"></el-date-picker>
        <el-time-picker v-else  v-model.trim="data" :readonly="widget.readonly" :placeholder="widget.placeholder"></el-time-picker>
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
    const widgets=inject("widgets")
    const widget=computed(()=>{
       return widgets.value[el.value.fid]
    })

    const getformat=()=>{
            if(widget.value.format=='datetime'){
                return 'YYYY-MM-DD HH:mm:ss'
            }else if(widget.value.format=='month'){
                return 'YYYY-MM'
            }else if(widget.value.format=='year'){
                return 'YYYY'
            }else if(widget.value.format=='date'){
                return 'YYYY-MM-DD'
            }else if(widget.value.format=='time'){
                return 'HH:mm:ss'
            }  
        }

    const dateType=computed(()=>{
        if(widget.value.format){
            return widget.value.format
        }
        return "datetime"
    })
</script>