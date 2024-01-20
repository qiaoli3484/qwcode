<template>
    <formitem :el="el" :widget="widget">
    <el-switch v-model="data" inline-prompt active-text="是" inactive-text="否" :readonly="widget.readonly" style="--el-switch-on-color: #13ce66; --el-switch-off-color: #ff4949" />
    </formitem>
</template>

<!-- :formatter='el.formatter'
:parser='el.parser -->'

<script setup>
    import {ref,defineProps,computed,defineEmits,inject,onMounted} from 'vue'
    import formitem from '../formitem.vue'
    const props=defineProps({
        el:Object,
        data:String,
    })
    const emit=defineEmits(['update:data'])
    const widgets=inject("widgets")
    const widget=computed(()=>{
       return widgets.value[props.el.fid]
    })
    const data=computed({
        get:()=>{
            if(props.data=='0'){
                return false
            }
            return true
        },
        set:(val)=>{
            emit('update:data',val==true?'1':'0')
        }
    })

    onMounted(()=>{
        data.value=widget.value.default=='是'?'1':'0'
    })
</script>