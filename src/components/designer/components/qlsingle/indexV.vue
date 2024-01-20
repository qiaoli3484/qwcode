<template>
    <formitem :el="el" :widget="widget">
        <el-checkbox-group v-model="data"  v-if="widget.multiple">
            <el-checkbox v-for="key in options" :label="key" :key="key" />
        </el-checkbox-group>
        <el-radio-group v-model="data" v-else>
            <el-radio v-for="key in options" :label="key" :key="key">{{key}}</el-radio>
        </el-radio-group>
    </formitem>
    
    <!-- <el-select :placeholder="el.placeholder" v-model="data"  :multiple="el.multiple" collapse-tags collapse-tags-tooltip @change="handleSelect" size="small" clearable style="width:100%">
        <el-option v-for="key in options" :key="key" :label="key" :value="key"></el-option>
    </el-select> -->
</template>


<script setup>
    import {ref,defineProps,computed,defineEmits,onMounted,inject} from 'vue'
    import formitem from '../formitem.vue'
    const props=defineProps({
        el:Object,
        data:String,
    })
    
    const emit=defineEmits(['update:data'])

    const widgets=inject("widgets")
    const widget=computed(()=>{
       return widgets.value[props.el.fid]
    })

    const data=computed({
        get:()=>{
            if(widget.value.multiple){
                if(typeof props.data=="undefined"){
                    return []
                }
                return props.data==''?[]:props.data.split(",")
            }
            return props.data
        },
        set:(val)=>{
            if(widget.value.multiple){
                emit('update:data',val.join(","))
            }else{
                emit('update:data',val)
            }
        }
    })

    const options=computed(()=>{
       return widget.value.options.split("|")
    })

    onMounted(()=>{
        if(widget.value.multiple&&typeof data.value!="undefined"){
            data.value= props.el.default==''?[]:widget.value.default.split(",")

        }else if(data.value==''){
            data.value=widget.value.default
        }
    })
</script>