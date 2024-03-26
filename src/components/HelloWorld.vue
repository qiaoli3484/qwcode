<script setup>
import { ref ,reactive,computed} from 'vue'
import { useEventListener } from './event'


defineProps({
  msg: String,
})

const count = ref(0)

const vFocus = {
  mounted: (el,binding) => {
    el.addEventListener("click",()=>{
      console.log(count.value,binding.value,dome)
      count.value++
    })
    //el.focus()
  }
}

const dome = reactive([
  {
    id:1,
    type:"elButton"
  },
  {
    id:2,
    type:"qlInput",
    data:"",
    child:[{
      id:3,
      type:"qldialog",
      data:false,
    }]
  },
  {
    id:4,
    type:"CustomComponent"
  },
])

</script>

<template>
  <h1>{{ msg }}</h1>

  <div class="card">
    <el-button v-focus type="button" >count is {{ count }}</el-button>
    <p>
      Edit
      <code>components/HelloWorld.vue</code> to test HMR
    </p>
  </div>
  <component v-focus="item" :item="item" v-for="item in dome" :is="item.type" :key="item.id"></component>
  
  <p class="read-the-docs">Click on the Vite and Vue logos to learn more</p>
</template>

<style scoped>
.read-the-docs {
  color: #888;
}
</style>
