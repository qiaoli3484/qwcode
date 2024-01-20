<template>
    <formitem :el="el" :widget="widget" :edit="edit" :data="data" v-slot="scope" :onfocus="onfocus">
        <el-input-number 
        ref="numRef"
        v-model.trim="data" 
        :placeholder="widget.placeholder" 
        :precision="widget.bits" 
        :controls="false" 
        :readonly="widget.readonly"
        @blur="onExitEditMode(scope)"
        />
    </formitem>
</template>

<!-- :formatter='el.formatter'
:parser='el.parser -->'

<script setup>
    import {ref,defineProps,computed,defineEmits,nextTick,inject} from 'vue'
    import formitem from '../formitem.vue'
    const props=defineProps({
        el:Object,
        data:{type:Number,default:0},
        edit:Boolean
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

    const numRef=ref(null)
    const onfocus=()=>{
        nextTick(()=>{
            console.log(numRef.value)
            numRef.value.focus()
        })
    }
    const onExitEditMode=(scope)=>{
      if(props.edit){
        scope.row.editing=false
      }
    }
</script>