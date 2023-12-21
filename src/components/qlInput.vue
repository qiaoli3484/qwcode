<template>
    <el-input v-model="item.data"></el-input>
</template>

<script setup>
import {defineEmits, defineProps} from "vue";
const { item } = defineProps(['item'])

const emits = defineEmits(["CustomEvent"]);
console.log("自定义组件事件", emits);
function onSendEvent(name,data) {
  console.log("自定义组件",name,"发送数据", item.data);
  emits("CustomEvent",name,data);
}
</script>