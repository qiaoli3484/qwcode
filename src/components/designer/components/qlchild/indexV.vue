<template>
    <!-- <el-input v-model.trim="data" :placeholder="el.placeholder" :readonly="el.readonly"/> -->

    <div>
    <el-button type="text" size="small" @click="add_(data.length)">添加</el-button>
    <el-button type="text" @click="openurl('/api/export?mode=3&cid='+widget.id+'&rid='+rid+'&tid='+pid)" size="small">下载</el-button>
    <el-upload style="display:inline-block;margin-left:10px" accept=".xlsx" :action="'/api/upload/rows?tid='+widget.id" :on-change="uploadchange" :on-success="uploadsuccess" :show-file-list="false"><el-button size="small" type="text" :loading="loading_">导入</el-button></el-upload>
    <el-button type="text" size="small" v-for="item in Generate"  @click="layer_show(item.name,'/select?tid='+item.from+'&to='+item.to)">来源{{item.name}}</el-button>
    </div>

    <el-table :data="data" size="small" height="500"  ref="ref_Table"  border highlight-current-row  style="width: 100%">
        <el-table-column label="序号" type="index" width="50" align="center"></el-table-column>
        <el-table-column align="center" label="操作" width="72" resizable>
            <template #default="{row,$index}">
                <el-button size="small" type="primary" link @click="add_($index,row)" icon="Plus"></el-button>
                <el-button size="small" type="danger" link @click="del_($index,row)" icon="Minus"></el-button>
            </template>
        </el-table-column>

        <template v-for="(th,index) in widget.childs" :key="th.alias">
        <!-- <el-table-column align="center" type="checkbox" width="50" v-if="th.type=='selection'" :key="th.type" resizable></el-table-column>
        <el-table-column align="center" label="下阶" width="50" v-if="th.type=='treenode'" tree-node resizable></el-table-column> -->
        <el-table-column :label="th.label" :prop="th.prop" :width="th.width" :fixed="th.fixed==''?undefined:th.fixed" align="center" show-overflow-tooltip resizable>
            <template #default="scope">

                <component
                :is="getFieldview(th,widgets,false)"
                :el="th"
                v-model:data="scope.row[th.prop]"
                v-model:row="scope.row"
                edit
                >
            </component>
                <!-- <ql-slot v-if="th.render!=''" :row="scope.row" :th="th" :index="scope.$index" :render="renderhtml"></ql-slot>
                <span v-else>{{evaljs(scope.row[th.prop],th.formatter)}}</span> -->
            </template>
        </el-table-column>
        </template>
    </el-table>

</template>

<script setup>
    import {ref,defineProps,computed,defineEmits,getCurrentInstance,onMounted,inject} from 'vue'
    import {getFieldview} from '../../components/common.js'
    const {proxy}=getCurrentInstance()
    import {ApiGetFormChild,ApiDelForm} from '@/utils/api.js'

    import {ElMessage,ElMessageBox} from 'element-plus'
    const props=defineProps({
        el:Object,
        data:{type:Array,default(rawProps){
            return []
        }},
    })

    const emit=defineEmits(['update:data'])
    
    const data=computed({
        get:()=>{
            console.log(props.data,"get")
            return props.data
        },
        set:(val)=>{
            console.log(val,"set")
            emit('update:data',val)
        }
    })

    const widgets=inject("widgets")
    const widget=computed(()=>{
       return widgets.value[props.el.fid]
    })

    const autoHeight=()=>{
        //可能有父级
        this.$nextTick(()=>{
            this.tableH=window.innerHeight-this.$refs.ref_Table.$el.offsetTop-this.$refs.ref_Table.$el.offsetParent.offsetTop-50
            this.tableW=window.innerWidth-10
            //console.log(window.innerHeight,this.$refs.multipleTable,this.tableH)
        })
    }
    const pid=inject("tid")
    const rid=inject("rid")

    const load_=ref(false)
    const getList=()=>{
        load_.value=true
        ApiGetFormChild(widget.value.id,pid.value,rid.value).then(res=>{
                data.value=res.data
                load_.value=false
                //this.currentdata=[]
                //this.currentsorce=[]
                //this.pagetotal=res.total - 0
                //this.footerData=res.sum[0]
                //合计数
        }).catch(()=>{
                load_.value=false
        })
    }

    const add_=(index,row)=>{
        let str='{"ID":0,"editing":false'
        console.log(str)
        console.log(widget.value.childs)
        for(let n=0;n<widget.value.childs.length;n++){
           const f= widgets.value[widget.value.childs[n].fid]
            if(f.type=="input-number"){
                str=str+`,"${widget.value.childs[n].prop}":${Number(f.default)}` 
            }else if(f.type=="date"){
                str=str+`,"${widget.value.childs[n].prop}":"${f.default.indeOf("Now")==-1?f.default:Date.now().toString()}"` 
                //default: "Now()"
            }else{
                str=str+`,"${widget.value.childs[n].prop}":"${f.default}"` 
            }
        }
        str=str+"}"
        console.log(data)
        data.value.splice(index+1,0,JSON.parse(str))
    }

    const del_=async(index,row)=>{
        if(row.ID=='0'){
            data.value.splice(index,1);
        }else{
            const type = await ElMessageBox.confirm('您确定要删除该数据?','删除',{
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning',
                draggable: true,
            })
            if (type === 'confirm') {
                
                const res =  await ApiDelForm(widget.value.id,row.ID)
                data.value.splice(index,1);
                ElMessage({showClose:true,message:res.msg,type:'success'})
            }
        }
    }

    const loading_=ref(false)
    const uploadchange=(file, fileList)=>{
        loading_.value=true
        if (file.status=='success'){
            data.value=file.response.data
            loading_.value=false
        }
    }
    const uploadsuccess=(response, uploadFile, uploadFiles)=>{
        if(response.code!=200){
            ElMessage({showClose:true,message:response.msg,type:'error'})
        }
    }
    const openurl=(url,title)=>{
        window.location.href=url
        //let index = parent.layer.getFrameIndex(window.name);
        //parent.layer.title(title, index)
    }
    onMounted(()=>{
        getList()
    })
    /*watch(()=>props.el,()=>{
        for(let n=0;n<props.el.childs.length;n++){
            for(let i=0;i<props.el.childs[n].childs[i].length;i++){
               const s1= props.el.childs[n].childs[i]
               if(s1.type=='qlchild'){
                    data.value[s1.alias]=[]
               }else{
                    data.value[item.alias]={}
               }
            }
        }
    })*/
</script>