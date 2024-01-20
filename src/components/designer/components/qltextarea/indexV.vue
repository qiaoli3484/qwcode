<template>
    <formitem :el="el" :widget="widget">
    <el-input 
    v-model.trim="data" 
    :placeholder="widget.placeholder" 
    :readonly="widget.readonly" 
    type="textarea"  
    :rows="widget.rows"
    :show-word-limit="widget.showWordLimit" />
    </formitem>
</template>

<!-- :formatter='el.formatter'
:parser='el.parser -->'

<script setup>
    import {ref,defineProps,computed,defineEmits,inject} from 'vue'
    import formitem from '../formitem.vue'
    const props=defineProps({
        el:Object,
        data:{type:String,default:''},
    })
    const emit=defineEmits(['update:data'])
    const widgets=inject("widgets")
    const widget=computed(()=>{
       return widgets.value[props.el.fid]
    })
    const data=computed({
        get:()=>{
            return props.data
        },
        set:(val)=>{
            emit('update:data',val)
        }
    })
</script>