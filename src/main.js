import { createApp } from 'vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import './style.css'
import App from './App.vue'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
import {createPinia} from 'pinia'

import Helper from './Helper.js'
import { post, get, del, put } from '@/util/http.js'

//import qlInput from './components/qlInput.vue'
//import qldialog from './components/qldialog.vue'
//import CustomComponent from './components/CustomComponent.vue'
const app = createApp(App)

import qlcustom from '@/components/customRenderComponent/customRenderComponent.vue'
app.component("qlcustom", qlcustom);


app.config.globalProperties.$get=get;
app.config.globalProperties.$del=del;
app.config.globalProperties.$put=put;
app.config.globalProperties.$post=post;
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
}

//注册饿了么组件
const BoxComponentNames_el = Helper.registerBoxComponentNames(app,"el",import.meta.glob('./components/boxs/base/**/*.vue', {eager: true}))
console.log("饿了么组件", BoxComponentNames_el)


app.provide('BoxComponentNames', BoxComponentNames_el)


//注册组件默认属性
const BoxComponentDefaultValue = {
    'el': Helper.registerBoxComponentDefaultValue(app, 'el', import.meta.glob('./components/boxs/base/**/*.js', {eager: true})),
    //'td': Helper.registerBoxComponentDefaultValue(app, 'td', import.meta.glob('./components/boxs/td/**/*.js', {eager: true}))
}
app.provide('BoxComponentDefaultValue', BoxComponentDefaultValue)


//app.component('qlInput', qlInput)
//app.component('qldialog', qldialog)
//app.component('CustomComponent', CustomComponent)

app.use(ElementPlus, { size: 'small', zIndex: 3000 })
app.use(createPinia())
app.mount('#app')



//注册组件
function registerBoxComponentNames(app,uiName,meta){
    let aa = import.meta.glob('./components/boxs/base/**/*.vue', {eager: true})
    console.log(aa,"222")
}