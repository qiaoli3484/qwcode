
<template>
  <span style="display: none;"></span>
  <!-- <s3Layer v-model="layer.visible" :content="layer.content" :title="layer.title" :type="2" :maxmin="true" :area="layer.area"></s3Layer> -->
</template>

<script setup>
  import {ref,defineEmits,reactive} from 'vue'
  // import { s3Layer } from 'vue3-layer';
  // import 'vue3-layer/dist/s3Layer.css';

  const emit=defineEmits(["update:name","update:online"])
  const wshost=ref('')
  const websock=ref(null)

  const websock_重连=ref(0)
  const agentData=ref('')
  const noRead=ref(0)

  const init_ws=()=>{
      websock.value = new WebSocket('ws://' + wshost.value + '/ws')
      websock.value.onmessage = ws_onmessage
      websock.value.onclose = ws_onclose
      websock.value.onopen = ws_onopen
    }
  const layer=reactive({visible:false,content:'',title:'',w:'',area:['auto','auto']})
  const layer_show=(n,v)=>{
      layer.title=n
      layer.content=v
      layer.visible=true
      layer.area[0]=window.innerWidth * 0.9 +'px';
      layer.area[1]=window.innerHeight - 50  +'px';
  }   

const ws_onopen=()=>{
  if (websock.value.readyState == 0) {
      console.log('连接尚未建立')
      ws_reconnet()
  } else if (websock.value.readyState == 1) {
      const d = new Date()
      console.log('连接已建立', +d.getMinutes() + ':' + d.getSeconds())
      clearInterval(websock_重连.value)
      console.log('close reconnet id', websock_重连.value)
      websock_重连.value = 0
  } else if (websock.value.readyState == 3) {
      alert('发送失败，请重新登录')
      ws_reconnet()
  }
}
  const ws_onmessage=(e)=>{
    let res = JSON.parse(e.data)
    if (res.type == 'heartbeat') {
      emit("update:name",res.data.name)
      return
    } else if (res.type == 'msg') {
      noRead.value = res.noRead
      layer_show(res.title, res.content)
      //this.layer_show(res.title, res.content)
    }else if (res.type == 'online') {
      let arr= []
      for(let x in res.data){
        arr.push({name:res.data[x],id:x})
      }
      emit("update:online",arr)
      return
      //this.layer_show(res.title, res.content)
    }
  }
  
  const ws_onclose=(e)=>{
    const d = new Date()
    console.log("connection closed ('连接关闭')" + d.getMinutes() + ':' + d.getSeconds())
    ws_reconnet()
  }

  const ws_reconnet=()=>{
    if (websock_重连.value == 0) {
      websock_重连.value = setInterval(init_ws, 10000)
      console.log('开始从新连接', websock_重连.value)
    }
  }

  const ws_send=(Data)=>{
      if (websock.value.readyState == 1) {
        websock.value.send(Data)
      } else if (websock.value.readyState == 0) {
        init_ws()
        setTimeout(()=> {
          websock.value.send(agentData)
        }, 500)
      }
    }

    const openxx=(res)=>{
      this.$notify({
        title: res.title,
        message: res.message,
        dangerouslyUseHTMLString: res.dangerouslyUseHTMLString,
        duration: res.duration,
        type: res.type,
      })
    }

    wshost.value = window.location.host
    init_ws()
</script>