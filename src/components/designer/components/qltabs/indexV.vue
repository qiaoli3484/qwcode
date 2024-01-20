<template>
    <el-tabs v-model="activeName" @tab-click="handleClick">
        <el-tab-pane :label="itemA.label" :name="itemA.label" v-for="(itemA,indexA) in el.childs"  :key="itemA.label">
            <template v-for="(item,index) in itemA.childs" :key="index">
                <component v-if="isUndefined(item)"
                    :is="getFieldview(item,widgets,detail)"
                    :el="item"
                    v-model:data="data">
                </component>
                <component v-else
                    :is="getFieldview(item,widgets,detail)"
                    :el="item"
                    v-model:data="data[widgets[item.fid].alias]"
                    v-model:row="data">
                </component> 
            </template>
        </el-tab-pane>
    </el-tabs>
</template>

<script setup>
    import {ref,defineProps,computed,defineEmits,onMounted,inject} from 'vue'
    import {getFieldview,isUndefined} from '../../components/common.js'

    const props=defineProps({
        el:Object,
        data:Object,
    })
    const emit=defineEmits(['update:data'])
    const widgets=inject("widgets")
    const detail=inject("detail")

    const data=computed({
        get:()=>{
            return props.data
        },
        set:(val)=>{
            emit('update:data',val)
        }
    })
    const activeName = ref(props.el.childs[0].label)

    
    const handleClick = (tab, event) => {
    console.log(tab, event)
    }

    onMounted(()=>{
        /*for(let n=0;n<props.el.childs.length;n++){
            for(let i=0;i<props.el.childs[n].childs.length;i++){
               const s1= props.el.childs[n].childs[i]
               if(s1.type=='qlchild'){
                    data.value[s1.alias]=[]
               }else{
                    data.value[item.alias]={}
               }
            }
        }*/
    })
</script>