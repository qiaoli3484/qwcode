<template>
<qlcode language="html" ref="ref_code" v-model:code="content" />
<qlcustom :content="content" />

</template>
<script setup>
import {ref,nextTick} from 'vue'
import qlcode from '@/components/qlcode/qlcode.vue'
import qlcustom from '@/components/customRenderComponent/customRenderComponent.vue'

const content=ref("<template><h2>aaa</h2></template>")
const ref_code=ref(null)


nextTick(()=>{
    ref_code.value.setcode(content.value)
})


</script>