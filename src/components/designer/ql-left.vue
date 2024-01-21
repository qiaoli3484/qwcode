<template>
     <el-tabs class="tabs-left" v-model="tabsName" type="border-card">
        <el-tab-pane label="组件设置" name="first">
            <el-collapse v-model="activeName">
                <el-collapse-item :title="item.label" :name="item.label" v-for="item in data" :key="item.label">
                    <vuedraggable  v-model="item.list"  
                    @start="drag = true" 
                    @end="drag = false"
                    @change="dragchange" 
                    :group="{ name: 'people', pull: 'clone', put: false }" 
                    :clone="(val)=>cloneEl(val,detail)"
                    item-key="label"
                    >
                        <template #item="{element}">
                            <div class="aaa" >
                                <span> 
                                    <el-icon>
                                        <Edit />
                                    </el-icon>
                                    {{element.label}}
                                </span>
                            </div>
                        </template>
                    </vuedraggable>
                </el-collapse-item>
            </el-collapse>
        </el-tab-pane>
        <el-tab-pane label="表单字段" name="second">
            <el-collapse v-model="activeName">
                <el-collapse-item title="字段" name="字段">
                    <vuedraggable  v-model="formlist"  
                    @start="drag = true" 
                    @end="drag = false"
                    @change="dragchange" 
                    :group="{ name: 'people', pull: 'clone', put: false }" 
                    :clone="(val)=>cloneEl(val,detail)"
                    item-key="label">
                        <template #item="{element}">
                            <div class="aaa" >
                                <span> 
                                    <el-icon>
                                        <Edit />
                                    </el-icon>
                                    {{element.label}}
                                </span>
                            </div>
                        </template>
                    </vuedraggable>
                </el-collapse-item>
            <el-collapse-item title="子表" name="子表">
                    <vuedraggable  v-model="childlist"  
                    @start="drag = true" 
                    @end="drag = false"
                    @change="dragchange" 
                    :group="{ name: 'people', pull: 'clone', put: false }" 
                    :clone="(val)=>cloneEl(val,detail)"
                    item-key="label">
                        <template #item="{element}">
                            <div class="aaa" >
                                <span> 
                                    <el-icon>
                                        <Edit />
                                    </el-icon>
                                    {{element.label}}
                                </span>
                            </div>
                        </template>
                    </vuedraggable>
                </el-collapse-item>

                <el-collapse-item :title="item.name" :name="item.name" v-for="item in tables" :key="item.name">
                    <vuedraggable  v-model="item.list"  
                    @start="drag = true" 
                    @end="drag = false"
                    @change="dragchange" 
                    :group="{ name: 'people', pull: 'clone', put: false }" 
                    :clone="(val)=>cloneEl(val,detail)"
                    item-key="label">
                        <template #item="{element}">
                            <div class="aaa" >
                                <span> 
                                    <el-icon>
                                        <Edit />
                                    </el-icon>
                                    {{element.label}}
                                </span>
                            </div>
                        </template>
                    </vuedraggable>
                </el-collapse-item>

            </el-collapse>
            <el-button @click="addtable">添加引用字段</el-button>
        </el-tab-pane>
    </el-tabs>
   
    <el-dialog title="数据条件" v-model="tableVisible" top="5vh" width="650px" :modal="false" :modal-append-to-body="false" draggable>
        <search :tables="tables" :tid="tid" :templets="templets" :fields="formlist" :table="table"/>
    </el-dialog>
        
</template>
<script setup>
    import { ref,onMounted,getCurrentInstance,inject,defineProps} from 'vue'
    import vuedraggable from 'vuedraggable'
    //import {cloneEl} from './components/common.js'
    import search from './components/defineSearch.vue'
    //import {ApiGetTemplets,ApiGetWidget,ApiGetWidgetMap} from '@/utils/api.js'
    const {proxy}=getCurrentInstance()
    const props= defineProps({
        detail:{type:Boolean,default:false},
    })
    const activeName = ref(["1"])
    const tabsName= ref("first")
    const data=ref([
            {label:'容器',list:[{icon:'',label:'栅格',type:'row'},{icon:'',label:'表格',type:'table'},{icon:'',label:'标签页',type:'tabs'},{icon:'',label:'卡片',type:'card'}]},
            {label:'基础组件',list:[{icon:'',label:'静态文字',type:'text'},{icon:'',label:'静态图片',type:'picture'},{icon:'',label:'HTML',type:'html'},{icon:'',label:'条形码',type:'barcode'},{icon:'',label:'二维码',type:'qrcode'},{icon:'',label:'按钮',type:'button'},{icon:'',label:'分隔线',type:'divider'},{icon:'',label:'富文本',type:'qleditor'}]},
            {label:'高级组件',list:[{icon:'',label:'数据表格',type:'dataTable'},{icon:'',label:'图表',type:'echarts'},{icon:'',label:'报表',type:'report'}]}
        ])
    const dragchange=()=>{
    }

    const tid=inject("tid")
    const rid=inject("rid")
  
    const tableVisible=ref(false)
    const tables=ref([])
    const table=ref({})
    const addtable=()=>{
        table.value={
            name: "",
            label: "",
            temp: "",
            by: "",
            sort: "",
            filter: "",
            enfilter: false,
            filters: [{ alias: "", value: "", radio: "and", compare: "等于" }], //输入提示和查询
            list: [],
            report:"",
            chart:""
        }
        tableVisible.value=true
    }

    const formlist=ref([])
    const childlist=ref([])
    const templets=ref([])
    const getForm=()=>{
        /*ApiGetWidget(tid.value).then(res => {
            formlist.value=res.data.fields
            childlist.value=res.data.childs
        })
        ApiGetTemplets().then(res => {
            templets.value=res.data
        })*/
    }

    const widgets=inject("widgets")
    //从这里获取组件信息
    const getwidgets=()=>{
        /*ApiGetWidgetMap(tid.value).then(res => {
            widgets.value=res.data
        })*/
    }


    onMounted(()=>{
        getForm()
        getwidgets()
    })
</script>

<style >
    .tabs-left>.el-tabs__content{
        padding: 0 5px 5px 5px!important;
    }
  .aaa{
    display: inline-block;
    height: 28px;
    line-height: 28px;
    width: 90px;
    float: left;
    color: #606266;
    margin: 2px 6px 6px 2px;
    cursor: move;
    white-space: nowrap;
    text-overflow: ellipsis;
    overflow: hidden;
    border: 1px solid #D4D7DE;
    border-radius: 3px;
    font-size: 12px;
  }

  .aaa:hover{
    background: #EBEDF0;
    border: 1px solid #409EFF;
  }

  .tabs-left .el-tabs__header.is-top {
    margin: 0px!important;
}
</style>