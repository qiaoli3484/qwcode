<template>
    <div>
        <ul class="el-upload-list el-upload-list--picture-card">
            <li class="el-upload-list__item is-success" v-for="item in data" :title="item.name" style="cursor:pointer;">
                <img class="el-upload-list__item-thumbnail" :src="item.url2" :alt="item.name" @click="openurl(item)">
            </li>
        </ul>

        <el-dialog v-model="dialogVisible" width="80%">
            <el-scrollbar>
                <img w-full :src="dialogImageUrl" alt="Preview Image" />
            </el-scrollbar>
        </el-dialog>
    </div>
</template>

<script setup>
    import {ref,defineProps,computed,inject} from 'vue'
    const props=defineProps({
        data:String,
        el:Object
    })

    const tid=inject("tid")

    const data=computed(()=>{
        let ar=[]
        if(typeof props.data=='undefined'||props.data==""){
            return ar
        }
        const aa=props.data.split("|")
        aa.forEach(item=>{
            const cc=item.split(":")
            ar.push({name:cc[0],new:cc[1],url:`/api/uploaded/${tid.value}/${props.el.alias}/${cc[1]}`,url2:`/api/uploaded/${tid.value}/${props.el.alias}/Thumb${cc[1]}`})
        })
        return ar
    })
    //黄春花发起了盐城中海-华樾项目【项目合同管理】 (1).jpg:KJy9oq8cOE.jpg|陈惠发起了金华中天美好-江璟园二期酒店项目【项目合同管理】.jpg:qOMkBJiwOW.jpg|system.png:DWREeZNkot.png
    const dialogVisible=ref(false)
    const dialogImageUrl=ref()
    const openurl=(item)=>{
        dialogImageUrl.value=item.url
        dialogVisible.value=true
    }
</script>