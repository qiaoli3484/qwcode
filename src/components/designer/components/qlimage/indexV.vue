<template>
    <formitem  :el="el" :widget="widget">
        <el-upload class="qlpic"
            multiple
            @contextmenu.prevent="()=>{}" 
            style="display:unset"
            :action="'/api/upload/photo?tid='+tid+'&alias='+widget.alias"
            :accept="widget.accept"
            list-type="picture-card" 
            :file-list="fileList"
            :limit="widget.limit"
            title="左键添加或右键点击后按Ctrl+V粘贴截图/图片" 
            :on-preview="handlePictureCardPreview" 
            :on-error="handleerror"
            :before-upload="beforeAvatarUpload"
            :on-success="handleSuccess"
            :on-remove="handleRemove" >
            <div @paste="pasteIntercept" style="font-size:30px">
            +
            </div>
        </el-upload>
        <el-dialog v-model="dialogVisible" width="80%" draggable>
            <el-scrollbar>
                <img w-full :src="dialogImageUrl" alt="Preview Image" />
            </el-scrollbar>
        </el-dialog>
    </formitem>
</template>

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
    const handleRemove=(file, fileList)=>{
        // `/api/upload/file?tid=${this.tid}&alias=${this.field.alias}&name=${file.new}`
        ApiDelFile(tid.value,widget.value.alias,file.new).then((res)=>{
            ElMessage({showClose:true,message:res.msg,type:'success'})
            update(null)
        })
    }
    
    const handleSuccess=(response, file, fileList)=>{
        if(typeof(response)!='undefined'){
            if(response.code!=200){
                ElMessage({showClose:true,message:response.msg,type:'error'})
            }else{
                update(response.data)
            }
        }
    }
    const beforeAvatarUpload = (rawFile) => {
        if ( widget.value.accept.indexOf(rawFile.type)==-1) {//!== 'image/jpeg'
            ElMessage.error(`上传图片的格式错误!${rawFile.type}`)
            return false
        }
        
        if (rawFile.size / 1024 / 1024 > widget.value.size) {
            ElMessage.error(`只能上传${widget.value.size}M大小图片!`)
            return false
        }
        return true
    }

    const pasteIntercept=(event)=>{
        const items = (event.clipboardData || window.clipboardData).items;
        let file = null;

        if (!items || items.length === 0) {
            ElMessage.error("当前浏览器不支持本地");
            return;
        }
        
        // 搜索剪切板items
        let type=""
        for (let i = 0; i < items.length; i++) {
            if (widget.value.accept.indexOf(items[i].type)>-1) {//items[i].type == "image/png"
                type=items[i].type
                console.log(items[i].getAsFile())
                file = items[i].getAsFile();
                break;
            }
        }
        if (!file) {
            ElMessage.error("粘贴内容非图片或不支持格式");
            return;
        }
        console.log(file.size/ 1024 / 1024 , widget.value.size)
        if (file.size / 1024 / 1024 > widget.value.size) {
            ElMessage.error(`只能上传${widget.value.size}M大小图片!`)
            return ;
        }
        if (fileList.value.length==widget.value.limit){
            ElMessage.error(`只能上传${widget.value.limit}张图片!`)
            return ;
        }
        // 此时file就是我们的剪切板中的图片对象
        // 如果需要预览，可以执行下面代码
        // const reader = new FileReader();
        //reader.onload = event => {
        //    preview.innerHTML = "<img src="+event.target.result+">";
        //};
        //reader.readAsDataURL(file);
        //this.file = file;
        uploadPlans(file,type)
    }
    //上传文件成功后回调
    const uploadPlans=(file,type)=>{
    //let file = this.file;
        if (!file) {
            ElMessage.error("请粘贴图片后上传");
            return;
        }
        loading.value = true;
        let form = new FormData();
        form.append("file", file);
        form.append("type", type);//"image/png"
        //uploadCertificate是封装的axios请求，自己根据需求传参
        ApiAddPhoto(tid.value,widget.value.alias,form).then((res)=>{
            update(res.data)
        })
    }

    const dialogVisible=ref(false)
    const dialogImageUrl=ref('')
    const  handlePictureCardPreview=(file)=>{
        dialogImageUrl.value = file.url;
        dialogVisible.value = true;
       // window.open(file.url,file.name)
        //openwk(file.url,file.name)
    }
    const update=(data)=>{
        let ar=[]
        fileList.value.forEach(item=>{
            ar.push(`${item.name}:${item.new}`)
        })
        if(data){
            ar.push(`${data.name}:${data.new}`)
        }
        emit('update:data',ar.join("|"))
    }

    const fileList=computed(()=>{
        let ar=[]
        if (typeof props.data != 'undefined'&&props.data!=''){
            const aa=props.data.split("|")
            aa.forEach(item=>{
                const cc=item.split(":")
                ar.push({name:cc[0],new:cc[1],url:`/api/uploaded/${tid.value}/${widget.value.alias}/${cc[1]}`})
            })
        }
        return ar
    })

</script>

<style>
.el-upload--picture-card {width: 100px!important; height: 100px!important;line-height: 100px!important;}
.el-upload-list--picture-card .el-upload-list__item{width: 100px!important;height: 100px!important;line-height: 100px!important;}
</style>