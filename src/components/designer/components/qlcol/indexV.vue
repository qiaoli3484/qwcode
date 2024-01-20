<template>
    <el-col :span="el.span" :offset="el.offset" :push="el.push" :pull="el.pull" v-for="item in el.childs" >
        <component v-if="isUndefined(item)"
                :is="getFieldview(item,widgets,false)"
                :el="item"
                v-model:data="data"
                v-model:row="data"
                >
        </component>
        <component v-else
            :is="getFieldview(item,widgets,detail)"
            :el="item"
            v-model:data="data[widgets[item.fid].alias]"
            v-model:row="data">
        </component> 
    </el-col>
</template>

<script setup>
import {ref,defineProps,computed,defineEmits,inject} from 'vue'
import {getFieldview,isUndefined} from '../../components/common.js'

const props=defineProps({
    el:Object,
    data:Object,
})
const emit=defineEmits(['update:data'])

const data=computed({
    get:()=>{
        return props.data
    },
    set:(val)=>{
        emit('update:el',val)
    }
})
const widgets=inject("widgets")
const detail=inject("detail")
</script>


<style scoped>
.row-bg{
min-height:30px;
}

</style>