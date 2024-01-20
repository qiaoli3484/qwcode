<template>
    <qlhandle :el="el" :widget="widget" formitem>
        <el-switch v-model="data" inline-prompt active-text="是" inactive-text="否" :readonly="widget.readonly" style="--el-switch-on-color: #13ce66; --el-switch-off-color: #ff4949" />
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
</script>