<template>
    <el-row class="aaa">
        <el-col :span="col.span" v-for="col in props.el.cols" class="row-bg aaa" >
            <vuedraggable  v-model="col.widgetList"  @start="drag = true" @end="drag = false" @change="dragchange" group="people" item-key="id" handle=".handle">
                <template #item="{element}">
                    <div @click="(e)=>{onselect(e,element.id)}" class="row-bg aaa">
                    <div class="t handle"><el-icon><Rank /></el-icon>{{element.name}}</div>
                    <div class="c">{{element.name}}</div>
                        <component 
                            :key="element.name"
                            :is="getFieldComponent(element)"
                            :label="element.label"
                            :type="element.type"
                            :placeholder="element.placeholder"
                            v-model="formData[element.name]">
                        </component>
                    </div>
                </template>
            </vuedraggable>
        </el-col>
    </el-row>
</template>
<script setup>
    import {ref,defineProps,computed,defineEmits} from 'vue'
    import vuedraggable from 'vuedraggable'
    //import {getFieldComponent} from '@/components/common.js'

    const props=defineProps({
        el:Object,
        formData:Object
    })
    const emit=defineEmits(['update:el'])
    
    const linkTo=computed({
        get:()=>{
            return props.el
        },
        set:(val)=>{
            emit('update:el',val)
        }
    })
    const onselect=(e,id)=>{
        console.log(e)
        console.log("选中的")
    }
</script>


<style scoped>
  .aaa{
    outline: 1px dashed #336699;
  }

.el-row {
  margin-bottom: 20px;
}
.el-row:last-child {
  margin-bottom: 0;
}
.el-col {
  border-radius: 4px;
}
.row-bg{
    min-height:30px;
}
.grid-content {
  border-radius: 4px;
  min-height: 36px;
}
</style>