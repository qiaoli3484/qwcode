<template>
    <div v-if="el.showfmt=='控件'">
        <span style="pointer-events: none" >
            <el-checkbox-group v-model="data"  v-if="widget.multiple">
                <el-checkbox v-for="key in options" :label="key" :key="key" />
            </el-checkbox-group>
            <el-radio-group v-model="data" v-else>
                <el-radio v-for="key in options" :label="key" :key="key">{{key}}</el-radio>
            </el-radio-group>
        </span>
    </div>
    <p v-else :style="{textAlign:el.textalign,fontSize:el.fontsize,color:el.fontColor,fontWeight:el.fontWeight,minHeight:data!=''?'unset':'22px'}">{{props.data}}</p>
</template>


<script setup>
    import {ref,defineProps,computed,defineEmits,onMounted} from 'vue'
    const props=defineProps({
        el:Object,
        data:String,
        widget:Object,
    })
    
    const emit=defineEmits(['update:data'])
    
    const data=computed(()=>{
        if(props.widget.multiple){
            return props.data==''?[]:props.data.split(",")
        }
        return props.data
    })

    const options=computed(()=>{
       return props.widget.list.split("|")
    })
</script>