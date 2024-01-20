<template>
    <table class="qtable" :style="{width:el.width,backgroundColor:el.bgColor}">
        <tbody>
            <tr v-for="itemA in el.childs">
                <td :colspan="itemB.colspan" :rowspan="itemB.rowspan" v-for="itemB in itemA.childs" :style="{width:itemB.width,backgroundColor:itemB.bgColor}">
                    <template v-for="item in itemB.childs">
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
                </td>
            </tr>
        </tbody>
    </table>
</template>
<!-- style="height: 24px;outline: 1px dashed #336699;" style="outline: 1px dashed #336699;" -->
<script setup>
    import {ref,defineProps,computed,defineEmits,inject,provide} from 'vue'
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
            emit('update:data',val)
        }
    })
    const widgets=inject("widgets")
    const detail=inject("detail")
</script>
<style scoped>
.qtable .el-form-item {
    margin-bottom: 0px!important;
}
.qtable {
    width: 100%;
    border-spacing: 0;
    border-collapse: collapse;
    border: none;
}
.qtable > tbody > tr > * {
    border-color: #000;
}
.qtable > tbody > tr > * {
    border: 1px #666 solid;
    padding: 3px 2px 2px 3px;
}
.qtable > tbody > tr > td {
    height: 24px;
}
</style>