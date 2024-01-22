<template>
    <div style="background-color:#E6E8EB;height:calc(100vh - 80px);padding:5px;">
        <el-form ref="formRef" :model="data"
            :label-position="layout.form.labelPosition" 
            :size="layout.form.csize" 
            style="overflow:auto;background: rgb(255, 255, 255);height:100%;" :style="{width:layout.form.width}">
            <vuedraggable  v-model="layout.content"  @start="drag = true" @end="drag = false" @change="dragchange" style="height:100%;margin:2px;" group="people" item-key="id" handle=".handle">
                <template #item="{element,index}">
                    <!-- <component 
                        :is="getFieldComponent(element,widgets,props.detail)"
                        v-model:el="layout.content[index]"
                        v-model:parent="layout.content"
                        :num="index"
                        v-model:data="data[element.alias]"
                        >
                    </component> --> 
                </template>
            </vuedraggable>
        </el-form>
    </div>
</template>
<script setup>
    import {ref,defineProps,defineEmits,computed,provide,inject} from 'vue'
    import vuedraggable from 'vuedraggable'
    //import {getFieldComponent} from './components/common.js'

    const props=defineProps({
        layout:Object,
        detail:Boolean
    })
    const emit= defineEmits(['update:layout'])
    const layout=computed({
        get:()=>{
            return props.layout
        },
        set:(val)=>{
            emit('update:layout',val)
        }
    })
    const drag=ref(true)
    const data=ref({})
    
    const dragchange=()=>{

    }
    //layout.form.labelWidth
    provide("layout",layout)
    provide("detail",props.detail)

    const widgets=inject("widgets")
</script>
<style>
.container1{
    padding: 10px;
    background: #f1f2f3;
    overflow-x: hidden;
    overflow-y: auto;
}

.row-bg{
    min-height: 24px;
}
.container{
    /*min-height: 24px!important;
    padding: 3px;*/
    /*outline: 1px dashed #336699;
    border: 1px dashed #336699;*/
    position: relative;
}
.drag-handler{
    color: #fff;
    background-color: #409EFF;
    position:absolute;
    top: 0px;
    left:0px;
    z-index: 1000;
    cursor: move;
    font-size: 12px;
    font-style: normal;
}

.drag-action{
    color: #fff;
    background-color: #409EFF;
    position:absolute;
    right: 0px;
    bottom: 0px;
    z-index: 1000;
}

.dragselect{
    outline: 2px solid #409EFF!important;
    /*box-sizing: content-box;*/
  }
.action-icon{
    cursor: pointer;
    margin:2px;
}
</style>