<template>
    <qlhandle :el="el">
        <el-table :data="data" 
              :size="el.csize" 
              :height="el.height"
              :style="{width: el.width}"
              ref="ref_Table"  
              :border="el.border" 
              :row-key="el.rowKey" 
              :stripe="el.stripe" 
              highlight-current-row
              >

            <el-table-column v-if="el.index" label="序号" type="index" width="50" align="center"></el-table-column>
            <el-table-column align="center" label="操作" width="72" resizable>
                <template #default="{row,$index}">
                    <el-button size="small" type="primary" link @click="add_($index,row)" icon="Plus"></el-button>
                    <el-button size="small" type="danger" link @click="del_($index,row)" icon="Minus"></el-button>
                </template>
            </el-table-column>

            <template v-for="(th,index) in el.rows" :key="th.alias">
                <el-table-column :label="th.label" 
                                :prop="th.alias" 
                                :width="th.width" 
                                :fixed="th.fixed==''?undefined:th.fixed" 
                                :align="th.align" 
                                show-overflow-tooltip 
                                :resizable="th.resizable">
                    <!-- <template #default="scope">
                        <component
                            :is="getFieldComponent(th.component)"
                            :el="th.component"
                            v-model:data="scope.row[th.prop]"
                            v-model:row="scope.row"
                            edit
                            >
                        </component>
                    </template> -->
                </el-table-column>
            </template>
        </el-table>
    </qlhandle>
</template>

<script setup>
    import {ref,defineProps,computed,defineEmits,inject} from 'vue'
    import qlhandle from '../handle.vue'
    import {getFieldComponent} from '../../components/common.js'
    const data=ref([])
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
</script>