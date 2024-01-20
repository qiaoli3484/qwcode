<template>
    <template v-if="edit">
        <div class="table-v2-inline-editing-trigger" v-if="!row.editing" @click="onEnterEditMode">{{data}}</div>
        <slot v-else :row="row"></slot>
    </template>
    
    <el-form-item  v-else
    :label="el.hidelabel?undefined:widget.label" 
    :label-width="el.hidelabel?'0px':labelWidth" 
    :rules="rules"
    :required="widget.required"
    :prop="widget.alias">
        <template #label='{ label }'>
            <div style="display: flex;align-items: center;">
                {{label}}
                <el-tooltip class='item error-tip' trigger='click' :content='widget.placeholder' placement='top' >
                    <el-icon v-if="widget.placeholder!=''" size="12px" style="cursor:pointer;" ><InfoFilled /></el-icon>
                </el-tooltip>
            </div>
        </template>
        <slot></slot>
        <template #error='{ error }'>
            <el-tooltip class='item error-tip' effect='dark' :content='error' placement='top' >
                <el-icon v-if="error!=''" color="#f56c6c" style="position: absolute;right: 5px;top: 4px;"><Warning /></el-icon>
            </el-tooltip>
        </template>
    </el-form-item>

</template>

<script setup>
    import {ref,defineProps,computed,reactive,inject} from 'vue'
    const props=defineProps({
        el:Object,
        edit:{type:Boolean,default:false},
        data:[String,Number],
        onfocus:Function,
        widget:Object,
    })
    const rules=computed(()=>{
        let arr=[]
        if(props.widget.required){
            arr.push({required: true,message: `请输入${props.widget.label}`,trigger: 'blur'})
        }
        if(props.widget.rule!=''){
            arr.push({pattern: new RegExp(props.widget.rule),message: `请输入正确的${props.widget.label}`,trigger: ['blur', 'change']})
        }
        if(props.widget.validate!=''){
            arr.push({pattern: new RegExp(props.widget.validate),message: `${props.widget.explain}`,trigger: ['blur', 'change']})
        }
        //Validate      string `json:"validate"` //验证
	    //Explain       string `json:"explain"`  //验证提示
        return arr.length==0?undefined:arr
        //{ required: true,pattern: /(^\\d{15}$)|(^\\d{18}$)|(^\\d{17}(\\d|X|x)$)/, message: '请输入正确的身份证号', trigger: 'blur' } 
    })

    const row=reactive({editing:false})
    const onEnterEditMode=()=>{
        row.editing = true
        props.onfocus()
    }
    const labelWidth=inject("labelWidth")
</script>

<style>
.table-v2-inline-editing-trigger {
  border: 1px transparent dotted;
  min-height: 22px;
}

.table-v2-inline-editing-trigger:hover {
  border-color: var(--el-color-primary);
}
</style>