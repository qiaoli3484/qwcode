<template>
    <div v-show="data.length==0"><el-button type="primary" icon="Plus" @click="add(0)">添加参数</el-button></div>
    <el-row v-for="(item,index) in data" :key="index">
        <el-col span="6">
            <el-button-group size="small">
                <el-button type="primary" icon="Plus" @click="add(index)"></el-button>
                <el-button type="danger" icon="Minus" @click="del(index)"></el-button>
            </el-button-group>
        </el-col>
        <el-col span="6">
            <el-input v-model.trim="item.title"></el-input>
        </el-col>
        <el-col span="6">
            <qladdsel v-model="item.param" />
        </el-col>
    </el-row>
</template>
<script setup>
   import { ref ,computed,defineEmits,defineProps} from 'vue'
   import qladdsel  from './addselect.vue'

   const prop=defineProps({
        modelValue:Array
   })
   const emit = defineEmits(['update:modelValue'])

   const data=computed({
        get: () => {
            if (typeof prop.modelValue=='undefined'){
                return new Array
            }
            return prop.modelValue
        },
        set: (val) => {
            emit('update:modelValue',val)
        }
    })

   const add=(index)=>{
        data.value.splice(index+1,0,{title:'',param:{default:'',params:[]}})
        //data.value.push({title:'',param:{default:'',title:'',params:[]}})
        
   }

   const del=(index)=>{
        data.value.splice(index,1)
        //data.value.push({title:'',param:{default:'',title:'',params:[]}})
   }
</script>