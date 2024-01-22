import { createApp } from 'vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import './style.css'
import App from './App.vue'
//import * as ElementPlusIconsVue from '@element-plus/icons-vue'


//import qlInput from './components/qlInput.vue'
//import qldialog from './components/qldialog.vue'
//import CustomComponent from './components/CustomComponent.vue'
const app = createApp(App)

/*for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
}*/

registerBoxComponentNames(app,"","")

//app.component('qlInput', qlInput)
//app.component('qldialog', qldialog)
//app.component('CustomComponent', CustomComponent)

app.use(ElementPlus, { size: 'small', zIndex: 3000 })
app.mount('#app')



//注册组件
function registerBoxComponentNames(app,uiName,meta){
    let aa = import.meta.glob('./components/boxs/base/**/*.vue', {eager: true})
    console.log(aa,"222")
}