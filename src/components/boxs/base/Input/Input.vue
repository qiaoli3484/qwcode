<template>
    <el-input 
    ref="inputRef"
    :type="item.type"
    v-model.trim="item.data" 
    :placeholder="item.placeholder" 
    :readonly="item.readonly" 
    :show-word-limit="item.showWordLimit" 
    @blur="onExitEditMode(scope)"
    @dblclick="ondblclick"
    size="small" />
</template>
<script setup>
import {defineProps} from 'vue'
const props=defineProps({
    item:Object
})

</script>