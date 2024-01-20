<template>
    <div @click.stop="onselect" class="container" :class="{dragselect:el.selected}"><!-- style="position:relative;" -->
        <div v-show="el.selected" class="drag-handler handle"><el-icon><Rank /></el-icon>{{el.label?el.label:widget.label}}</div>
        <div v-show="el.selected" class="drag-action"> 
            <el-icon :size="14" color="#fff" class="action-icon" @click.stop="BackRow" title="选中父组件"><Back /></el-icon>
            <el-icon :size="14" color="#fff" class="action-icon" @click.stop="TopCol"  title="上移组件" v-show="parent.length>1"><Top /></el-icon>
            <el-icon :size="14" color="#fff" class="action-icon" @click.stop="BackCol" title="下移组件" v-show="parent.length>1"><Bottom /></el-icon>
            <slot name="header"></slot>
            <el-icon :size="14" color="#fff" class="action-icon" @click.stop="ondelel"><Delete /></el-icon>
        </div>
        <el-form-item :label="el.hidelabel?undefined:widget.label" :label-width="el.hidelabel?'0px':layout.form.labelWidth" :required="widget.required" :prop="widget.alias" v-if="formitem" :rules="rules">
            <slot></slot>
            <template #error='{ error }'>
                <el-tooltip class='item error-tip' effect='dark' :content='error' placement='top' >
                    <el-icon v-if="error!=''" color="#f56c6c" style="position: absolute;right: 5px;top: 4px;"><Warning /></el-icon>
                </el-tooltip>
            </template>
        </el-form-item>
        <template v-else>
            <slot></slot>
        </template>
    </div>
</template>

<script setup>
    import {ref,defineProps,computed,defineEmits,inject,watch} from 'vue'
    //import {swapArray} from './common.js'
    import {ElMessage} from 'element-plus'
    const props=defineProps({
        el:Object,
        formData:Object,
        num:Number,
        parent:[Array,Object],
        formitem:Boolean,
        widget:Object,
    })
    const emit=defineEmits(['update:el','update:parent'])
    
    const el=computed({
        get:()=>{
            return props.el
        },
        set:(val)=>{
            emit('update:el',val)
        }
    })
    const parent=computed({
        get:()=>{
            return props.parent
        },
        set:(val)=>{
            emit('update:parent',val)
        }
    })
    const layout=inject("layout")

    const elSelect=inject("elSelect")
    const onselect=()=>{
        el.value.selected=true
        elSelect(el.value)
    }
    const ondelel=()=>{
        if(typeof parent.value.childs=='undefined'){
            parent.value.splice(props.num,1)
            return
        }
        parent.value.childs.splice(props.num,1)
    }

    const BackRow=()=>{
        parent.value.selected=true
        elSelect(parent.value)
    }

    const BackCol=()=>{
        if(typeof parent.value.childs=='undefined'){
            if(props.num+1 != parent.value.length){
                swapArray(parent.value,props.num,props.num+1)
            }else{
                ElMessage({showClose:true,message:"已经处于置底，无法下移",type:'info'})
            }
            return
        }
        if(props.num+1 != parent.value.childs.length){
            swapArray(parent.value.childs,props.num,props.num+1)
        }else{
            ElMessage({showClose:true,message:"已经处于置底，无法下移",type:'info'})
        }
    }

    const TopCol=()=>{
        if(props.num != 0){
            if(typeof parent.value.childs=='undefined'){
                swapArray(parent.value,props.num,props.num-1)
            }else{
                swapArray(parent.value.childs,props.num,props.num-1)
            }
        }else{
            ElMessage({showClose:true,message:"已经处于置顶，无法上移",type:'info'})
        }
    }

    const rules=computed(()=>{
        let arr=[]
        if(props.widget.required){
            arr.push({required: true,message: `请输入${props.widget.label}`,trigger: 'blur'})
        }
        if(props.widget.rule!=''){
            arr.push({pattern: new RegExp(props.widget.rule),message: `请输入正确的${props.widget.label}`,trigger: ['blur', 'change']})
        }
        return arr.length==0?undefined:arr
        //{ required: true,pattern: /(^\\d{15}$)|(^\\d{18}$)|(^\\d{17}(\\d|X|x)$)/, message: '请输入正确的身份证号', trigger: 'blur' } 
    })
</script>