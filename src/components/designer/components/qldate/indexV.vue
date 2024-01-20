<template>
    <formitem :el="el" :widget="widget">
        <el-date-picker v-if="dateType!='time'" v-model="data" :type="dateType" :value-format="getformat()" clearable :readonly="widget.readonly" :placeholder="widget.placeholder"></el-date-picker>
        <el-time-picker v-else  v-model="data" :format="getformat()"  :readonly="widget.readonly" :placeholder="widget.placeholder" clearable></el-time-picker>
    </formitem>
</template>

<!-- :formatter='el.formatter'
:parser='el.parser -->'

<script setup>
    import {ref,defineProps,computed,defineEmits,inject,onMounted} from 'vue'
    import formitem from '../formitem.vue'
    const props=defineProps({
        el:Object,
        data:[String,Number],
    })
    const emit=defineEmits(['update:data'])
    const widgets=inject("widgets")
    const widget=computed(()=>{
       return widgets.value[props.el.fid]
    })

    const data=computed({
        get:()=>{
            if(widget.value.format=='time'&&props.data!=''){
                const _arr = props.data.split(":");
                return new Date(2010, 10, 10, Number(_arr[0]), Number(_arr[1]), Number(_arr[2]))
            }
            return props.data
        },
        set:(val)=>{
            if(widget.value.format=='time'&&val!=''){
                emit('update:data', val.getHours()+":"+val.getMinutes()+":"+val.getSeconds())
            }else{
                emit('update:data',val)
            }
        }
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

    onMounted(()=>{
        if(data.value==''){
            data.value=widget.value.default;
        }
    })

    const disabledDate = (time) => {
        return time.getTime() > Date.now()
    }
</script>