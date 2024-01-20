<template>
    <el-dialog v-model="dialogVisible" title="Tips" width="30%" draggable>
      <template v-for="item in props.el.childs">
            <component v-if="isUndefined(item)"
                    :is="getFieldview(item,widgets,false)"
                    :el="item"
                    v-model:data="data"
                    v-model:row="data"
                    >
            </component>
            <component v-else
                :is="getFieldview(item,widgets,detail)"
                :el="item"
                v-model:data="data[widgets[item.fid].alias]"
                v-model:row="data"
                >
            </component> 
        </template>
  </el-dialog>
</template>

<script setup>
    import {ref,defineProps,computed,defineEmits,provide,defineExpose,onMounted,inject} from 'vue'
    import {ElMessage} from 'element-plus'
    import {getFieldview,isUndefined} from '../../components/common.js'
    import {ApiGetRenderForm,ApiGetForm} from '@/utils/api.js'
    const props= defineProps({
        el:Object,
        data:{
            type:Object,
            default(rawProps) {
                return {}
                }
            },
    })
    const emit= defineEmits(['update:data'])
    const data=computed({
        get:()=>{
            return props.data
        },
        set:(val)=>{
            emit('update:data',val)
        }
    })
    
    provide('labelWidth',props.el.labelWidth)
    provide('formData',data)
    const formRef=ref(null)

    const detail=inject("detail")
    const widgets=inject("widgets")

    const submitForm=()=>{
        return new Promise((resolve, reject) => {
            formRef.value.validate((valid) => {
                if (valid) {
                    resolve(true)
                } else {
                    resolve(false)
                    ElMessage({showClose:true,message:"提交失败!",type:'error'})
                    return false
                }
            })
        })
    }

    const clearData = () => {
        if (!formRef.value) return;
        formRef.value.resetFields()
    }

    const tid=inject("tid")
    const rid=inject("rid")
    const initdata= async()=>{
        if(rid.value==0)return;
        const res= await ApiGetForm(tid.value,rid.value)
        data.value={...data.value,...res.data[props.el.alias]}
    }

    onMounted(()=>{
        initdata()
        /*props.el.childs.forEach(item=>{
            if(item.type=='input-number'){
                data.value[item.alias]=0
            }else{
                data.value[item.alias]=""
            }
        })*/
    })



    defineExpose({
        submitForm,
        clearData
    })
</script>