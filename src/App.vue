<script setup>
import {ref} from 'vue'
//import HelloWorld from './components/designer/ql-view.vue'

import HelloWorld from './components/customBox/index.vue'
import comps from './app/comps.vue'


// let aa = import.meta.glob('./components/boxs/base/**/*.vue', {eager: true})
//     console.log(aa,"222")
//     const componentsContext=aa
//     let ComponentNames = []
//     const keys = Object.keys(componentsContext).map((path) => {
//         const name = path.split('/').pop().replace(/\.\w+$/, '');
//         const module = componentsContext[path];
//         //如果名称后面是 属性 或者是 Attr 就不要加入
//         if (!(name.endsWith("属性") || name.endsWith("Attr"))) {
//             ComponentNames.push(name)
//         }
//         console.log("注册组件", "uiName", name)
//         //app.component(uiName + name, module.default);
//         // app.component(name, module.default);
//         return name;
//     });
//     console.log(aa,"222",ComponentNames)

const layout=ref([])
</script>

<template>
  <HelloWorld msg="Vite + Vue" :layout="layout"/>
  <comps></comps>
</template>

<style scoped>

</style>
