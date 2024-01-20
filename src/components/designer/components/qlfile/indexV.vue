<template>
     <formitem :el="el" :widget="widget">
        <span v-for="(item,index) in fileList">
            <img :src="'/static/images/fileico/'+gettype(item.name)+'.gif'" align="absBottom" onerror="this.src='/static/images/fileico/unknown.gif'" class="fileico">
            <el-button type="primary" link @click.stop="handlePreview(item)" size="small">{{item.name}}
            <el-icon style="padding-left:10px" @click.stop="handleRemove(item,index)" @mouseleave.stop="e=>{e.target.style.color='black'}" @mouseenter.stop="e=>{e.target.style.color='red'}"><Close /></el-icon>
            </el-button>
        </span>
            <el-upload
                class="upload-demo"
                :action="'/api/upload/file?tid='+tid+'&alias='+widget.alias"
                multiple
                :accept="widget.accept"
                :limit="widget.limit"
                :on-progress="handleprogress"
                :on-success="handleSuccess"
                :before-upload="beforeAvatarUpload"
                :on-error="handleerror" 
                :show-file-list="false"
                ><el-button type="primary" size="small" icon="UploadFilled" link></el-button></el-upload>
                <el-progress v-if="percentage>0&&percentage<99" :percentage="percentage" />
    </formitem>
</template>

<!-- :formatter='el.formatter'
:parser='el.parser -->'

<script setup>
    import {ref,defineProps,computed,defineEmits,inject,getCurrentInstance} from 'vue'
    import {ElMessage} from 'element-plus'
    import formitem from '../formitem.vue'
    import {ApiAddPhoto,ApiDelFile} from '@/utils/api'
    const {proxy}=getCurrentInstance()
    const props=defineProps({
        el:Object,
        data:String,
    })
    const emit=defineEmits(['update:data'])
    
    const widgets=inject("widgets")
    const widget=computed(()=>{
       return widgets.value[props.el.fid]
    })

    const tid=inject("tid")
    const rid=inject("rid")

    const loading=ref(false)

    const handleerror=(err, file, fileList)=>{
        ElMessage({showClose:true,message:err,type:'error'})
    }
    const gettype=(val)=>{
        return val.match(/\.([^.]+)$/)[1]
    }

    const beforeAvatarUpload = (rawFile) => {
        if (widget.value.accept.indexOf(".*")==-1&&widget.value.accept.indexOf(rawFile.name.match(/\.([^.]+)$/)[1])==-1) {//!== 'image/jpeg'
            ElMessage.error(`上传文件的格式错误!`)
            return false
        } 

        if (rawFile.size / 1024 / 1024 > widget.value.size) {
            ElMessage.error(`只能上传${widget.value.size}M大小文件!`)
            return false
        }
        return true
    }
   
    const handleRemove=(file,index)=>{
       
        ApiDelFile(tid.value,widget.value.alias,file.new).then((res)=>{
            ElMessage({showClose:true,message:res.msg,type:'success'})
            let ar=[]
            fileList.value.forEach((item,i)=>{
                if (i!=index){
                    ar.push(`${item.name}:${item.new}`)
                }
            })
            emit('update:data',ar.join("|"))
        })
        //this.T4N7AL1.FJw1vc9=fileList
        console.log(file);
    }
    const handlePreview=(file)=>{
        window.open(file.url,file.name)
    }
    const  handleSuccess=(response, file)=>{
        if(typeof(response)!='undefined'){
            if(response.code!=200){
                ElMessage({showClose:true,message:response.msg,type:'error'})
            }else{
                let ar=[]
                    fileList.value.forEach(item=>{
                        ar.push(`${item.name}:${item.new}`)
                    })
                if(response.data){
                    ar.push(response.data.name+":"+response.data.new)
                }
                emit('update:data',ar.join("|"))
            }
        }
    }
    const percentage=ref(0)
    const handleprogress=(evt,UploadFile,UploadFiles)=>{
        percentage.value=evt.percentage
    }

    const fileList=computed(()=>{
        let ar=[]
        if(typeof props.data!='undefined'&&props.data!=''){
            const aa=props.data.split("|")
            aa.forEach(item=>{
                const cc=item.split(":")
                ar.push({name:cc[0],new:cc[1],url:`/api/uploaded/${tid.value}/${widget.value.alias}/${cc[1]}`})
            })
        }
        return ar
    })

</script>