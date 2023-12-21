import { createApp } from 'vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import './style.css'
import App from './App.vue'

import qlInput from './components/qlInput.vue'
import qldialog from './components/qldialog.vue'
const app = createApp(App)

app.component('qlInput', qlInput)
app.component('qldialog', qldialog)

app.use(ElementPlus, { size: 'small', zIndex: 3000 })
app.mount('#app')
