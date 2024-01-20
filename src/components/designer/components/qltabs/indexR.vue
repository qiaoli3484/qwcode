
<template>
    <qlhandle :el="el">
        <el-tabs v-model="activeName" @tab-click="handleClick">
            <el-tab-pane :label="item.label" :name="item.label" v-for="(item,indexA) in el.childs" style="padding: 5px;">
                <vuedraggable  v-model="el.childs[indexA].childs"  @start="drag = true" @end="drag = false"  group="people" item-key="id" handle=".handle" class="row-bg">
                    <template #item="{element,index}">
                        <component 
                            :is="getFieldComponent(element,widgets,detail)"
                            v-model:el="el.childs[indexA].childs[index]"
                            v-model:parent="el.childs[indexA]"
                            :num="index">
                        </component>
                    </template>
                </vuedraggable>
            </el-tab-pane>
        </el-tabs>
    </qlhandle>
</template>

<script setup>
    import {ref,defineProps,computed,defineEmits,inject} from 'vue'
    import vuedraggable from 'vuedraggable'
    import qlhandle from '../handle.vue'
    import {getFieldComponent} from '../../components/common.js'

    const props=defineProps({
        el:Object,
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

    const widgets=inject("widgets")
    const detail=inject("detail")

    const activeName = ref('first')

    const handleClick = (tab, event) => {
    console.log(tab, event)
    }
</script>

