<template>
    <qlhandle :el="el" :widget="widget" formitem>
        <el-select filterable v-if="el.remote" 
        @click="remoteMethod('')"  
        v-model="data" 
        :placeholder="widget.placeholder" 
        :multiple="widget.multiple" 
        collapse-tags collapse-tags-tooltip 
        :loading="loading" :remote="widget.remote" 
        :remote-method="remoteMethod" @change="handleSelect" 
        size="small" clearable style="width:100%;">
            <el-option v-for="key in options" :key="key.value" :label="key.label" :value="key.value">
                <span v-html="hhh(key)" v-if="el.formatter!=''"></span>
            </el-option>
        </el-select>

        <el-select-v2 filterable v-else 
        :placeholder="widget.placeholder" 
        v-model="data"  
        :options="options"
        :multiple="widget.multiple" 
        collapse-tags collapse-tags-tooltip 
        @change="handleSelect" 
        style="width:100%;"
        clearable />
    </qlhandle>
</template>

<script setup>
    import {ref,defineProps,computed,defineEmits,inject,onMounted} from 'vue'
    import qlhandle from '../handle.vue'
    const data=ref()
    const props=defineProps({
        el:Object,
        formData:Object
    })
    const emit=defineEmits(['update:el'])
    
    const el=computed({
        get:()=>{
            return props.el
        },
        set:(val)=>{
            emit('update:el',val)
        }
    })
    const widgets=inject("widgets")
    const widget=computed(()=>{
       return widgets.value[el.value.fid]
    })

    const tid=inject("tid")
    const rid=inject("rid")

    const form=inject("formData")

    const loading=ref(false)
    const options=ref([])
    const remoteMethod=(query)=>{
        loading.value = true;
        let val={}
        if( typeof form!='undefined'){
            val=form.value
        }
        proxy.$post("/api/auto/"+widget.value.alias+"?value="+query+"&tid="+widget.value.tid,JSON.stringify(val)).then(res=>{
            options.value=res.data
            loading.value = false;
        }).catch(()=>{
            loading.value = false;
            options.value = [];
        })
    }

    const handleSelect=(val)=>{  
        const aa=options.value.find(item =>{ return item.value==val});
        if( typeof form!='undefined'){
            val=form.value
            const str= widget.value.onChange.replace(/this.form/g,'form.value')
            console.log(str)
            eval(str)
        }else if(typeof row!='undefined'){
           const str= widget.value.onChange.replace(/this.form/g,'row.value')
           console.log(str)
           eval(str)
        }
        
        /*if(this.mapping!=""& typeof(aa)!='undefined'){
            eval(this.mapping)
        }*/
    }

    const hhh=(key)=>{
        if (this.mapping!=''){
            const s1=this.mapping.replace(/this.form.(\w+)= aa./gi,"key.").replace(/;/g,"+ ' '+")
            return key.value+ eval(s1.substring(0,s1.length-6))
        }
        return key.value
    }

    onMounted(()=>{
        if(!widget.value.remote){
            options.value=widget.value.options.split("|").map((item)=>({
                value:item,label:item
            }))
        }
    })
</script>