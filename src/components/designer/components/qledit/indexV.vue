<template>
        <template v-if="!el.readonly">
            <el-tooltip placement="top-start" effect="light"  trigger="hover" popper-class="mypopper" :offset="-5">
                <template #content>
                <el-button type="primary" icon="Edit" size="small" @click="init"/>
                </template>
                <qimage  v-if="el.type=='image'" :el="el" :data="data" :widget="widget"></qimage>
                <qnumber v-else-if="el.type=='input-number'" :el="el" :data="data" :widget="widget"></qnumber>
                <qyesno  v-else-if="el.type=='switch'" :el="el" :data="data" :widget="widget" />
                <qradio v-else-if="el.type=='single'" :el="el" :data="data" :widget="widget" />
                <p v-else :style="{textAlign:el.textalign,fontSize:el.fontsize,color:el.fontColor,fontWeight:el.fontWeight,minHeight:data!=''?'unset':'22px'}">{{data}}</p>
            </el-tooltip>
       </template>
       <template v-else>
            <p :style="{textAlign:el.textalign,fontSize:el.fontsize,color:el.fontColor,fontWeight:el.fontWeight,minHeight:data!=''?'unset':'22px'}">{{data}}</p>
       </template>


        <el-dialog v-model="Visible" :title="widget.label" :width="450" draggable :modal="false" :close-on-click-modal="false">
            <el-form ref="formRef" :model="formData" size="small">
                 <component
                    :is="getFieldview(el,widgets,false)"
                    :el="el"
                    v-model:data="data">
                </component> 
            </el-form>
            <!--
            <el-alert v-if="explain!=''" :title="explain" type="warning" show-icon :closable="false" />
            <el-row v-if="istable">
                <el-button @click="Visible = false" size="small">添加明细</el-button>
                <el-button @click="Visible = false" size="small">删除一行</el-button>
            </el-row>
            -->
            <template #footer>
                <span class="dialog-footer">
                    <el-button @click="Visible = false">取消</el-button>
                    <el-button type="primary" @click="editfield" :loading="load_">保存</el-button>
                </span>
            </template>
        </el-dialog>
</template>

<script setup>
    import {ref,defineProps,computed,defineEmits,inject} from 'vue'
    import {getFieldview,isUndefined} from '../../components/common.js'
    import {ElMessage} from 'element-plus'
    import {ApiEditFieldVal} from '@/utils/api'
    import qimage from './image.vue'
    import qnumber from './qnumber.vue'
    import qyesno from './qyesno.vue'
    import qradio from './qradio.vue'

    const props=defineProps({
        el:Object,
        data:String,
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

    const formData =inject("formData")

    const widgets=inject("widgets")
    const widget=computed(()=>{
       return widgets.value[props.el.fid]
    })
    /*const text=computed(()=>{
        return eval(props.el.text)
    })*/

    const Visible=ref(false)
    const load_=ref(false)
    const init= async ()=>{
        /*let url=""
        if(this.mode=="flow"){
            url='/api/field/flow?tid='+this.vtid+'&rid='+this.vrid+'&id='+this.id
        }else{
            url='/api/field?tid='+this.vtid+'&rid='+this.vrid+'&id='+this.id
        }
        const res =  await this.$get(url)
        this.form.data=res.data.value
        this.field=res.data.field
        this.explain=this.field.explain//提示*/
        Visible.value=true
    }

    const tid=inject('tid')
    const rid=inject('rid')
    const formRef=ref()
    const editfield= async ()=>{
        await formRef.value.validate((valid) => {
            if (valid) {
                load_.value=true
                ApiEditFieldVal(tid.value,rid.value,JSON.stringify({alias:widget.value.alias,value:data.value})).then((res)=>{
                    load_.value=false
                    ElMessage({showClose:true,message:res.msg,type:'success'})
                }).catch(()=>{load_.value=false})
            } else {
                console.log('error submit!')
                return false
            }
        })
    }
</script>
<style>
.mypopper{
    padding:0!important ;
}
</style>