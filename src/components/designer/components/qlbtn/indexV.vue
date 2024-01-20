<template>

    <el-button-group style="margin-bottom:10px">
        <el-button v-if="el.btns.indexOf('new')>-1" @click="cleardata()" type="primary" size="small">新建</el-button>

        <el-button v-if="el.btns.indexOf('edit')>-1" @click="openurl('/form?tid='+tid+'&rid='+rid)" type="primary" size="small">编辑</el-button>

        <el-button v-if="el.btns.indexOf('layout')>-1" @click="openurl('/form/layout/html?tid='+tid+'&rid='+rid)" type="primary" size="small">版式</el-button>
        <el-button v-if="el.btns.indexOf('save')>-1" @click="editdata()" :loading="saveload" type="primary" size="small">保存</el-button>
            
        <template v-if="el.btns.indexOf('flow')>-1">
        <el-button  @click="flow_submit('start')" :loading="saveload"  type="primary" size="small">提交</el-button>
        <el-button  @click="editdata()" v-if="rid==0" :loading="saveload" type="primary" size="small">保存草稿</el-button>
        </template>
        <el-button  v-for="item in flow" @click="flow_submit(item.n)" :loading="saveload" type="primary" size="small">{{item.s}}</el-button>

        <el-button v-if="el.btns.indexOf('del')>-1" @click="deldata()" icon="Delete" :loading="saveload" type="danger" size="small">删除</el-button>
        <el-button v-if="el.btns.indexOf('print')>-1" @click="printer()" type="primary" size="small">打印</el-button>
        <el-button  @click="closedown()" icon="SwitchButton" type="info" size="small">关闭</el-button>
    </el-button-group>

    <el-dialog v-model="dialogVisible" title="Tips" width="30%" draggable>
        <el-date-picker type="datetime" :disabled-date="disabledDate" size="small" value-format="YYYY-MM-DD HH:mm:ss" clearable  v-model="PlanTime" placeholder="填写事务计划完成时间"></el-date-picker>
        <el-input v-model="msg" placeholder="请输入意见" show-word-limit type="textarea" />
        <template #footer>
            <span class="dialog-footer">
                <el-button @click="dialogVisible = false">取消</el-button>
                <el-button type="primary" @click="flow_submit2">确认</el-button>
            </span>
        </template>
    </el-dialog>
</template>

<script setup>
    import {ref,defineProps,inject} from 'vue'
    import {ElMessage,ElMessageBox} from 'element-plus'
    import {ApiAddFormData,ApiDelForm,ApiRunCustom,ApiRunFlow,ApiEditFormData} from '@/utils/api.js'

    const props=defineProps({
        el:Object,
        data:String,
    })

    const itemRefs=inject('itemRefs')
    const tid=inject("tid")
    const rid=inject("rid")
    const data=inject("data")

    const saveload=ref(false)
    const  editdata= async()=>{
            if(saveload.value)return;
            saveload.value=true
            let arrs =[]
            for(let i=0;i<itemRefs.value.length;i++){
                if(itemRefs.value[i].submitForm){
                  const res = await itemRefs.value[i].submitForm()
                  if(!res){
                    saveload.value=false
                    return
                  }
                }
            }

            if(rid.value==0){
                const res= await ApiAddFormData(tid.value,JSON.stringify(data.value)).catch(()=>{saveload.value=false;return})
                    //that[props.el.alias]=res.data[props.el.alias]
                    data.value[props.el.alias]=res.data[props.el.alias]
                    rid.value=res.data[props.el.alias].ID
                    //this.TBfViSa=res.data.TBfViSa
                    /*arrs.forEach(item=>{
                        that[item].data=res.data[item]
                    })*/
                    ElMessage({showClose:true,message:res.msg,type:'success'})
                    saveload.value=false
            } else {
                const res= await ApiEditFormData(tid.value,rid.value,JSON.stringify(data.value)).catch(()=>{saveload.value=false;return})
                    //this.TBfViSa=res.data.TBfViSa
                    //that[props.el.alias]=res.data[props.el.alias]
                    //this.$refs.ref_TrGuYDi.data=res.data.TrGuYDi
                    /*arrs.forEach(item=>{
                        that[item].data=res.data[item]
                    })*/
                    ElMessage({showClose:true,message:res.msg,type:'success'})
                    saveload.value=false
            }
        }
    const del_sure = function(content) {
        return ElMessageBox.confirm('删除 '+content, '提示', {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'warning'
        })
    }    
    const deldata=()=>{
        del_sure('当前数据').then(()=>{
            saveload.value=true
            ApiDelForm(tid.value,data.value[props.el.alias].ID).then(res=>{
                ElMessage({showClose:true,message:res.msg,type:'success'})
                cleardata()
                saveload.value=false
            }).catch(()=>{
                saveload.value=false
            })
        })
    }
    const custom_=(val)=>{
        ApiRunCustom(tid.value,data.value[props.el.alias].ID,val).then(res=>{
            ElMessage({showClose:true,message:res.msg,type:'success'})
        })
    }

    const disabledDate=(time)=>{
        return time.getTime() < Date.now()
    }
    const flow_submit=async(val)=>{
        let rid=0
        this.flag=val
        if (rid.value){
            rid=rid.value
        }else{
            rid=data.value[props.el.alias].ID
        }
        if(rid==0){
            await this.editdata()
            rid=data.value[props.el.alias].ID
        }
        if(this.needplantime||this.opnrequired){
            this.dialogVisible=true
            return
        }
        ApiRunFlow(tid.value,rid,val,JSON.stringify({msg:this.msg,NeedPlanTime:this.PlanTime})).then(res=>{
            ElMessage({showClose:true,message:res.msg,type:'success'})
            closedown()
        })
    }

    const flow_submit2= async()=>{
        let rid=0
        if (rid.value){
            rid=rid.value
        }else{
            rid=data.value[props.el.alias].ID
        }
        if(this.needplantime&&this.PlanTime==""){
            ElMessage({showClose:true,message:"必须填入计划完成日期",type:'error'})
            return
        }
        if(this.opnrequired&&this.msg==""){
            ElMessage({showClose:true,message:"必须填入意见",type:'error'})
            return 
        }
        
        ApiRunFlow(tid.value,rid,this.flag,JSON.stringify({msg:this.msg,NeedPlanTime:this.PlanTime})).then(res=>{
            ElMessage({showClose:true,message:res.msg,type:'success'})
            closedown()
        })
    }
    const printer=()=>{
        print('#myPrintArea')
    }

    const openurl=(url)=>{
        window.open(url,'_self')
    }


    const initdata=inject("initdata")
    const cleardata=()=>{
        rid.value=0
        initdata()
        /*itemRefs.value.forEach(item=>{
            if(item.clearData){
                item.clearData()
            }
        })*/
    }
    const closedown=inject('closedown')
    /*const closedown =()=>{
        let index = parent.layer.getFrameIndex(window.name);
        parent.layer.close(index);
    }*/
</script>