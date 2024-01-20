<template>
        <el-col :span="el.span" class="container" @click.stop="onselect" :class="{dragselect:el.selected}" :offset="el.offset" :push="el.push" :pull="el.pull">
            <div v-show="el.selected" class="drag-handler">栅格列</div>
            <div v-show="el.selected" class="drag-action"> 
                <el-icon><Back /></el-icon>
                <el-icon><Top /></el-icon>
                <el-icon><Right /></el-icon>
                <el-icon><CopyDocument /></el-icon>
                <el-icon><Delete /></el-icon>
            </div>
            <vuedraggable  v-model="el.childs"  @start="drag = true" @end="drag = false"  group="people" item-key="id" handle=".handle" class="row-bg">
                <template #item="{element,index}">
                        <component 
                            :is="getFieldComponent(element,detail)"
                            :placeholder="element.placeholder"
                            v-model:parent="el"
                            v-model:el="el.childs[index]"
                            >
                        </component>
                </template>
            </vuedraggable>
        </el-col>
</template>

<!-- <div @click.stop="(e)=>{onselect(e,element.id)}" class="row-bg aaa">
    <div class="t handle"><el-icon><Rank /></el-icon>{{element.name}}</div>
    <div class="c">{{element.name}}</div> -->

<script setup>
    import {ref,defineProps,computed,defineEmits,inject} from 'vue'
    import vuedraggable from 'vuedraggable'
    import {getFieldComponent} from '../../components/common.js'

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
    const elSelect=inject("elSelect")
    const onselect=()=>{
        el.value.selected=true
        elSelect(el.value)
    }
    const detail=inject("detail")
    
</script>


<style scoped>
.row-bg{
    min-height:30px;
}

</style>