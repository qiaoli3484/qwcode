<template>
    <qlhandle :el="el" :widget="widget" formitem>
        <el-checkbox-group v-model="data"  v-if="widget.multiple">
            <el-checkbox v-for="key in options" :label="key" :key="key" />
        </el-checkbox-group>
        <el-radio-group v-model="data" v-else>
            <el-radio v-for="key in options" :label="key" :key="key">{{key}}</el-radio>
        </el-radio-group>
    </qlhandle>
</template>

<!-- :formatter='el.formatter'
:parser='el.parser -->'

<script setup>
    import {ref,defineProps,computed,defineEmits,inject} from 'vue'
    import qlhandle from '../handle.vue'
   
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

    const data=ref(widget.value.multiple?[]:'')

    const options=computed(()=>{
       return widget.value.options.split("|")
    })
</script>