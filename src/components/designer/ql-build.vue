<template>
   <template v-for="(item, index) in props.jsonData.childs" :key="index" >
        <component v-if="item.type=='form'" ref="itemRefs"
            :is="getFieldview(item,widgets,false)"
            :el="item"
            v-model:data="data[item.alias]">
        </component>
        <component v-else
            :is="getFieldview(item,widgets,false)"
            :el="item"
            v-model:data="data">
        </component>
    </template>
</template>

<script setup>
import {ref,defineProps,computed,defineEmits,onMounted,watch,provide,inject} from 'vue'
//import {getFieldview} from './components/common.js'

    const props= defineProps({
        jsonData:Object,
        data:Object,
        detail:{type:Boolean,default:false}
    })

    const emit= defineEmits(['update:data'])
    const data=computed({
        get:()=>{
            return props.data
        },
        set:(val)=>{
            emit('update:data',val)
        }
    })
    const widgets=inject("widgets")
   
    const itemRefs = ref([])
    provide('itemRefs',itemRefs)
    provide('detail',props.detail)
    const initdata=()=>{
        props.jsonData.childs.forEach(item=>{
            if(item.type=="form"){
                data.value[item.alias]={}
                walkdata(item,data.value[item.alias])
            }else{
                walkdata(item,data.value)
            }
        })
    }

    const walkdata=(obj,data)=>{
        if(typeof obj.childs=='undefined')return;
        for(let i=0;i<obj.childs.length;i++){
            const item=obj.childs[i]
            if(typeof item.type =='qlchild'){//||item.type=="tabs"||item.type==""||item.type=="row"||item.type=="col"
                //data[item.alias]=[]
                //walkdata(item,data[item.alias])
            }else if(typeof item.hidelabel =='undefined'){//||item.type=="tabs"||item.type==""||item.type=="row"||item.type=="col"
                walkdata(item,data)
            }else if(typeof item.alias=='undefined'){
                
            }else if(item.type=="input-number"){
                data[item.alias]=0
            }else if(item.type=="single"){
            
            }else{
                data[item.alias]=''
            }

        }
    }

    //provide('initdata',initdata)
    watch(()=>props.jsonData,()=>{
        initdata()
    })




    onMounted(()=>{
       
    })

</script>

<style>
/*.el-form-item {
    margin-bottom:0px!important;
}*/
</style>