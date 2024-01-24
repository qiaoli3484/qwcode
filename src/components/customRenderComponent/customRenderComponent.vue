<template>
    <el-button @click="loadComponent()">更新</el-button>
    <component :is="remote" />
</template>

<script setup>
import {ref, watch,markRaw, defineProps} from 'vue';
import {loadModule} from 'vue3-sfc-loader';

const props=defineProps({
    content:String
})

const remote = ref(null);

const loadComponent = async()=>{
    /*let style = item.HTML.match(/<style scoped>([\s\S]*)<\/style>/)?.[1] || "";
    if (style === "") {
        style = item.HTML.match(/<style>([\s\S]*)<\/style>/)?.[1] || "";
    }
    const 自定义css = style;
    //如果 item.HTML 没有 <template> 标签，就添加一个
    if (item.HTML.indexOf("<template>") === -1) {
        //将 <template> 欻人到开头 将 </template>插入到 <style>前面
        item.HTML = item.HTML.replace(/([\s\S]*)<style>/, "<template>\n$1\n</template>\n\n<style>");
    }*/
    const Vue = await import('vue')
    const options = {
      moduleCache: {
        vue: Vue
      },
      getFile(url) {
        console.log(url,"1111")
        /*const res = await fetch(url);
        if ( !res.ok )
          throw Object.assign(new Error(res.statusText + ' ' + url), { res });
        return {
          getContentData: asBinary => asBinary ? res.arrayBuffer() : res.text(),
        }*/
        return Promise.resolve(props.content)
      }, 
      addStyle() {
       
      }
    }
   console.log(options)
   remote.value=Vue.defineAsyncComponent(() =>loadModule("/hah.vue",options))
}



loadComponent()

//创建监听 item.HTML 重新渲染
watch(() => props.content, () => {
  console.log("重新渲染")
  loadComponent();
})
</script>
/*
compiledCache: {
  get: key => window.localStorage.getItem(key),
  set: (key, value) => window.localStorage.setItem(key, value),
},*/