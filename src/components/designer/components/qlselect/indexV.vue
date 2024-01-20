<template>
    <formitem :el="el" :widget="widget" :edit="edit" :data="data" v-slot="scope" :onfocus="onfocus">
        <el-select filterable ref="selectRef" v-if="widget.remote"  
        @click="remoteMethod('')" @blur="onExitEditMode(scope)" 
        v-model="data" :placeholder="widget.placeholder" 
        :multiple="widget.multiple" collapse-tags collapse-tags-tooltip 
        :loading="loading" :remote="widget.remote" 
        :remote-method="remoteMethod" @change="handleSelect"  
        @clear="cleardata"
        style="width:100%;" clearable>
            <el-option v-for="key in options" :key="key.value" :label="key.label" :value="key.value">
                <span v-html="hhh(key)" v-if="widget.formatter!=''"></span>
            </el-option>
        </el-select>
        <el-select-v2 filterable ref="selectRef" v-else
            :placeholder="widget.placeholder"
            v-model="data"  
            :multiple="widget.multiple" 
            collapse-tags collapse-tags-tooltip
            :options="options"
            @change="handleSelect"
            @blur="onExitEditMode(scope)"  
            style="width:100%;"
            clearable />
    </formitem>
</template>

<script setup>
    import {ref,defineProps,computed,defineEmits,inject,getCurrentInstance,onMounted,nextTick} from 'vue'
    import formitem from '../formitem.vue'
    const {proxy}=getCurrentInstance()
    const props=defineProps({
        el:Object,
        data:String,
        row:Object,
        edit:Boolean,
    })
    const emit=defineEmits(['update:data','update:row'])
    const widgets=inject("widgets")
    const widget=computed(()=>{
       return widgets.value[props.el.fid]
    })

    const data=computed({
        get:()=>{
            if(widget.value.multiple){
                if(typeof props.data=='undefined'){
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

    const row=computed({
        get:()=>{
            return props.row
        },
        set:(val)=>{
            emit('update:row',val)
        }
    })

    const tid=inject("tid")
    const rid=inject("rid")

    const form=inject("formData")

    const loading=ref(false)
    const options=ref([])
    const remoteMethod=(query)=>{
        loading.value = true;
        /*let val={}
        if( typeof form!='undefined'){
            val=form.value
        }*/
        //console.log(props.row,form)

        proxy.$post("/api/auto/"+widget.value.alias+"?value="+query+"&tid="+widget.value.tid,JSON.stringify(row.value)).then(res=>{
            options.value=res.data
            loading.value = false;
        }).catch(()=>{
            loading.value = false;
            options.value = [];
        })
    }

    const handleSelect=(val)=>{  
        const aa=options.value.find(item =>{ return item.value==val});
        const str= widget.value.onChange.replace(/this.form/g,'row.value')
            console.log(str)
            eval(str)

        /*if( typeof row !='undefined'){//优先
            val=form.value
            const str= props.el.onChange.replace(/this.form/g,'row.value')
            console.log(str)
            eval(str)
        }else if(typeof form !='undefined'){
           const str= props.el.onChange.replace(/this.form/g,'form.value')
           console.log(str)
           eval(str)
        }*/
        console.log(aa,typeof form,typeof row)
        /*if(this.mapping!=""& typeof(aa)!='undefined'){
            eval(this.mapping)
        }*/
    }
    const cleardata=()=>{
        //优先
    }

    const hhh=(key)=>{
        if (this.mapping!=''){
            const s1=this.mapping.replace(/this.form.(\w+)= aa./gi,"key.").replace(/;/g,"+ ' '+")
            return key.value+ eval(s1.substring(0,s1.length-6))
        }
        return key.value
    }


    const selectRef=ref(null)
    const onfocus=()=>{
        nextTick(()=>{
            console.log(selectRef.value)
            selectRef.value.focus()
        })
    }

    const onExitEditMode=(scope)=>{
      if(props.edit){
        setTimeout(()=>{//延时等待数据选择
            scope.row.editing=false
        },500)
      }
    }

    onMounted(()=>{
        if(!widget.value.remote){
            options.value=widget.value.options.split("|").map((item)=>({
                value:item,label:item
            }))
        }
    })
</script>