<template>
    <el-container>
      <el-aside width="200px">
        <el-scrollbar>
        <qlleft :detail="props.detail" />
        </el-scrollbar>
      </el-aside>

      <el-container>
        <el-header class="toolbar-header">
            <el-button type="primary" icon="Back" link @click="Undo" size="default"></el-button>
            <el-button type="primary" icon="Right" link @click="Redo" :disabled="counter<=-1" size="default"></el-button>
            <el-radio-group v-model="radio2" @change="changesize">
                <el-radio-button label="PC" />
                <el-radio-button label="Pad" />
                <el-radio-button label="H5" />
            </el-radio-group>
            当前版式:
        </el-header>

        <el-main>
            <el-scrollbar>
                <!-- <qlmid v-model:layout="layout" :detail="props.detail"/> -->
            </el-scrollbar>
        </el-main>
      </el-container>
      
       <!-- <el-aside width="250px">
       <qlright v-model:el="currentEl"  v-model:form="layout.form" ref="ref_right"/>
      </el-aside> -->
    </el-container>
</template>


<script setup>
    import {ref,provide,defineProps,computed,defineEmits} from 'vue'
    import qlleft from './ql-left.vue'
    /*import qlmid from './ql-middle.vue'
    import qlright from './ql-right.vue'*/

    const props= defineProps({
        layout:Object,
        detail:Boolean,
    })
    const emit=defineEmits(['update:layout'])
    const layout=computed({
        get:()=>{
            return props.layout
        },
        set:(val)=>{
            emit('update:layout',val)
        }
    })

    const widgets=ref({})
    provide('widgets',widgets)

    const ref_right=ref()
    const nextId=ref("")
    const currentEl=ref({})
    const elSelect=(obj)=>{
        currentEl.value=obj
        if(nextId.value==''||nextId.value==obj.id){
            nextId.value=obj.id
            return
        }
        for (let i=0;i<layout.value.content.length;i++){
            if(layout.value.content[i].id==nextId.value){
                layout.value.content[i].selected=false
                break
            }
            if(fib(layout.value.content[i])){
                break
            }
        }
        nextId.value=obj.id
        if (typeof nextId.value=='undefined'){
            ref_right.value.setactiveName('second')
            return
        }
        ref_right.value.setactiveName('first')
    }

    //递归寻找
    const fib=(obj)=>{
        if(typeof obj.childs=='undefined'){
            return false
        }
        for(let n=0;n<obj.childs.length;n++){
            if(obj.childs[n].id==nextId.value){
                obj.childs[n].selected=false
                return true
            }
            if(fib(obj.childs[n])){
                return true
            }
        }
        return false
    }
    provide('elSelect',elSelect)
    
    const historyData=ref([])
    const hisbefore=()=>{
        //前面插入
        console.log("add his")
        historyData.value.unshift(JSON.stringify(layout.value.content))
        if(historyData.value.length>15){
            historyData.value.pop() //删除最后一个元素
        }
    }
    provide('hisbefore',hisbefore)

    const counter=ref(0)
    //撤销
    const Undo=()=>{
        counter.value++
        if(counter>historyData.value.length){
            counter.value--
            return
        }
        console.log('Undo',counter.value)
        layout.value.content=JSON.parse(historyData.value[counter.value])
    }
    //重做
    const Redo=()=>{
        counter.value--
        console.log('Redo',counter.value)
        if(counter.value<=-1){
            counter.value=-1
            return
        }
        layout.value.content=JSON.parse(historyData.value[counter.value])
    }

    const radio2=ref('PC')
    const formConf=ref({csize:'small',labelPosition:'left',labelWidth:150,wdith:"100px"})
    const changesize=(val)=>{
        layout.value.form.display=val
        if(val=="Pad"){
            layout.value.form.width='960px';
        }else if(val=="H5"){
            layout.value.form.width= '420px'
        }else if(val=="PC"){
            layout.value.form.width= '100%'
        }  
        //960px
        //420px
    }
</script>
<style>
    .el-main {
        --el-main-padding: 0px!important;
    }
    .el-header.toolbar-header{
        font-size: 14px;
        height: 42px !important;
        border-bottom: 1px dotted rgb(204, 204, 204);
    }
</style>
//{ name: 'people', pull: 'clone', put: false }" 表示可拖拽克隆出去，但不接收外面拖拽过来的