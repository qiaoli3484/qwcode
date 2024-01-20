<template>
        <el-col :span="el.span" class="container row-bg" @click.stop="onselect" :class="{dragselect:el.selected}" :offset="el.offset" :push="el.push" :pull="el.pull" style="outline: 1px dashed #336699;">
            <div v-show="el.selected" class="drag-handler">栅格列</div>
            <div v-show="el.selected" class="drag-action"> 
                <el-icon :size="14" color="#fff" class="action-icon" @click.stop="BackRow" title="选中父组件"><Back /></el-icon>
                <el-icon :size="14" color="#fff" class="action-icon" @click.stop="TopCol"  title="上移组件" v-show="parent.childs.length>1"><Top /></el-icon>
                <el-icon :size="14" color="#fff" class="action-icon" @click.stop="BackCol" title="下移组件" v-show="parent.childs.length>1"><Bottom /></el-icon>
                <el-icon :size="14" color="#fff" class="action-icon" @click.stop="copycol"><CopyDocument /></el-icon>
                <el-icon :size="14" color="#fff" class="action-icon" @click.stop="delcol"><Delete /></el-icon>
            </div>
            <vuedraggable  v-model="el.childs"  @start="drag = true" @end="drag = false"  group="people" item-key="id" handle=".handle" style="height:100%;width:100%">
                <template #item="{element,index}">
                        <component 
                            :is="getFieldComponent(element,widgets,detail)"
                            v-model:el="el.childs[index]"
                            v-model:parent="el"
                            :num="index">
                           
                        </component>
                </template>
            </vuedraggable>
        </el-col>
</template>

<script setup>
    import {ref,defineProps,computed,defineEmits,inject} from 'vue'
    import vuedraggable from 'vuedraggable'
    import {ElMessage} from 'element-plus'
    import {getFieldComponent,colConf,swapArray} from '../../components/common.js'

    const props=defineProps({
        el:Object,
        formData:Object,
        parent:Object,
        num:Number,
    })
    const emit=defineEmits(['update:el','update:parent'])
    
    const el=computed({
        get:()=>{
            return props.el
        },
        set:(val)=>{
            emit('update:el',val)
        }
    })
    const widgets=inject("widgets")
  
    const parent=computed({
        get:()=>{
            return props.parent
        },
        set:(val)=>{
            emit('update:parent',val)
        }
    })
    
    const detail=inject("detail")
    const elSelect=inject("elSelect")
    const onselect=()=>{
        el.value.selected=true
        elSelect(el.value)
    }
    const copycol=()=>{
        let c=colConf()
        c.span=el.value.span
        parent.value.childs.splice(props.num+1,0,c)
    }
    const delcol=()=>{
        if(parent.value.childs.length<=1){
            el.childs=[]
            return
        }
        parent.value.childs.splice(props.num,1,)
    }

    const BackRow=()=>{
        parent.value.selected=true
        elSelect(parent.value)
    }

    const BackCol=()=>{
        if(props.num+1 != parent.value.childs.length){
            swapArray(parent.value.childs,props.num,props.num+1)
        }else{
            ElMessage({showClose:true,message:"已经处于置底，无法下移",type:'info'})
        }
    }

    const TopCol=()=>{
        if(props.num != 0){
            swapArray(parent.value.childs,props.num,props.num-1)
        }else{
            ElMessage({showClose:true,message:"已经处于置顶，无法上移",type:'info'})
        }
    }
</script>


<style scoped>
.row-bg{
    min-height:30px;
}

</style>