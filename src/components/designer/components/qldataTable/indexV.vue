<template>
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
                    :is="getFieldview(th.component)"
                    :el="th.component"
                    v-model:data="scope.row[th.alias]"
                    v-model:row="scope.row"
                    edit
                    >
                </component>
            </template> -->
        </el-table-column>
        </template>
    </el-table>
</template>

<script setup>
    import {ref,defineProps,computed,defineEmits,getCurrentInstance,onMounted,inject} from 'vue'
    import {getFieldview} from '../../components/common.js'
    const {proxy}=getCurrentInstance()
    import {ElMessage,ElMessageBox} from 'element-plus'
    const props=defineProps({
        el:Object,
        data:Array,
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
        proxy.$post("/api/get/dataTable?tid="+props.el.data.temp+"&rid="+rid.value+"&pid="+pid.value,JSON.stringify(props.el.data)).then(res=>{
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
        console.log(props.el.childs)
        for(let n=0;n<props.el.childs.length;n++){
            str=str+`,"${props.el.childs[n].prop}":""` 
        }
        str=str+"}"
        data.value.splice(index+1,0,JSON.parse(str))
    }

    const del_=async(index,row)=>{
        if(row.ID=='0'){
            data.value.splice(index,1);
        }else{
            const type = await ElMessageBox.confirm('您确定要删除该数据?','Warning',{
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'warning',
            draggable: true,
            })
            if (type === 'confirm') {
                const res =  await proxy.$del('/api/edit?tid='+this.cid+'&rid='+row.ID)
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