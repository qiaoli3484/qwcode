<template>
    <formitem :el="el" :widget="widget" :edit="edit" :data="data" v-slot="scope" :onfocus="onfocus">
        <el-input 
        ref="inputRef"
        v-model.trim="data" 
        :placeholder="widget.placeholder" 
        :readonly="widget.readonly" 
        :show-word-limit="widget.showWordLimit" 
        @blur="onExitEditMode(scope)"
        @dblclick="ondblclick"
        size="small" />
    </formitem>
</template>

<script setup>
    import {ref,defineProps,computed,defineEmits,onMounted,nextTick,inject} from 'vue'
    import formitem from '../formitem.vue'
    const props=defineProps({
        el:Object,
        data:{type:String,default:""},
        edit:Boolean
    })
    const emit=defineEmits(['update:data'])
    
    const data=computed({
        get:()=>{
            return props.data
        },
        set:(val)=>{
            emit('update:data',val)
        }
    })
    const widgets=inject("widgets")
    const widget=computed(()=>{
       return widgets.value[props.el.fid]
    })

    const inputRef=ref(null)
    const onfocus=()=>{
        nextTick(()=>{
            console.log(inputRef.value)
            inputRef.value.focus()
        })
    }
    const onExitEditMode=(scope)=>{
      if(props.edit){
        scope.row.editing=false
      }
    }
    const ondblclick=(e)=>{
          
    }
    onMounted(()=>{
        if(data.value==''){
            data.value=widget.value.default
        }
    })
</script>