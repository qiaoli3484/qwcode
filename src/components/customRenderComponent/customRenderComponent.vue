<template>
    <el-button @click="load()" v-if="!props.auto">更新</el-button>

    <component v-if="!props.auto" :is="remote" :layer="props.layer" />
    <component  v-else  :is="remote" :layer="props.layer"  v-model:data='props.vdata[props.layer.alias]' />
</template>

<script setup>
import {ref, watch,shallowRef,getCurrentInstance,defineProps,defineAsyncComponent} from 'vue';
import * as Vue from "vue";
import {loadModule} from 'vue3-sfc-loader';

const {proxy} =getCurrentInstance();
//
const props=defineProps({
    content:String,
    url:String,
    layer:[Object,Array],
    vdata:[Object,Array],
    auto:{type:Boolean,default:true}
})
//v-model:data='props.vdata[props.layer.alias]' 第一次没有
console.log(props.content,props.layer,props.vdata,"2222111")

const remote = shallowRef();

const load=()=>{
  loadComponent()
}

const loadComponent =()=>{
    console.log(props.content)
    try {
        const options = {
          moduleCache: {
            vue: Vue
          },
          async getFile(url) {
            const {data} = await proxy.$get("/api/components/getcode/"+url)
            console.log(url,222)
            /*const res = await fetch(url);
            if ( !res.ok )
              throw Object.assign(new Error(res.statusText + ' ' + url), { res });
            return {
              getContentData: asBinary => asBinary ? res.arrayBuffer() : res.text(),
            }*/
            return data
          }, 
          addStyle(textContent) {
            const style = Object.assign(document.createElement("style"), {
              textContent
            });
            const ref = document.head.getElementsByTagName("style")[0] || null;
            document.head.insertBefore(style, ref);
          }
        }
      const comp =defineAsyncComponent(() =>loadModule(props.content,options))

      remote.value=comp
  } catch (err) {
    console.error(err);
  }
}


if (props.auto){
  loadComponent()
}


//创建监听 item.HTML 重新渲染
/*watch(() => props.content, () => {
  console.log("重新渲染")
  loadComponent();
})*/
</script>
/*
compiledCache: {
  get: key => window.localStorage.getItem(key),
  set: (key, value) => window.localStorage.setItem(key, value),
},*/