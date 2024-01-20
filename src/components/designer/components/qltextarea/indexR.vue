<template>
    <qlhandle :el="el" widget="widget" formitem>
        <el-input 
        v-model.trim="data" 
        :placeholder="widget.placeholder" 
        :readonly="widget.readonly" 
        type="textarea" 
        :rows="widget.rows"
        :show-word-limit="widget.showWordLimit"
        />
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