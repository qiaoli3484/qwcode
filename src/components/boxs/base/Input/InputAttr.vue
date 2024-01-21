<template>
    <el-form label-width="100px" label-position="left" size="small">
        <el-form-item label="唯一ID">
            <el-input size="small" v-model.trim="el.id" readonly></el-input>
        </el-form-item>
        <el-form-item label="标签">
            <el-input size="small" v-model.trim="el.label"></el-input>
        </el-form-item>
        <el-form-item label="隐藏字段标签"  v-if="!isUndefined(el.hidelabel)">
            <el-switch v-model="el.hidelabel" />
        </el-form-item>
        
        <el-form-item label="只读" v-if="!isUndefined(el.readonly)">
            <el-switch v-model="el.readonly" />
        </el-form-item>
        <el-form-item label="占位文本" v-if="!isUndefined(el.placeholder)">
            <el-input  v-model.trim="el.placeholder"></el-input>
        </el-form-item>
        <el-form-item label="最大输入长度" v-if="!isUndefined(el.maxlength)">
            <el-input-number v-model="el.maxlength" :min="1" :max="100" />
        </el-form-item>

        <el-form-item label="type">
            <el-select v-model="props.item.type" style="width: 100%">
                <el-option v-for="item in type选项" :key="item" :label="item.label" :value="item.value"/>
            </el-select>
        </el-form-item>
    </el-form>
</template>

<script setup>
import {defineProps} from 'vue'
const props=defineProps({
    item:Object
})

let size选项 = ref([
  {"label": "自定义", "value": "custom"},
  {"label": "大号", "value": "large"},
  {"label": "小号", "value": "small"},
]);

let type选项 = ref([
  {"label": "单行文本框", "value": "text"},
  {"label": "多行文本框", "value": "textarea"},
  {"label": "密码输入框", "value": "password"},
  {"label": "数字输入框", "value": "number"},
]);

let eventName = ref([
  {"label": "被单击", "value": "click"},
  {"label": "鼠标左键被按下", "value": "mousedown"},
  {"label": "鼠标左键被放开", "value": "mouseup"},
  {"label": "被双击", "value": "dblclick"},
  {"label": "鼠标右键被按下", "value": "contextmenu"},
  {"label": "鼠标位置被移动", "value": "mousemove"},
  {"label": "获得焦点", "value": "focus"},
  {"label": "失去焦点", "value": "blur"},
  {"label": "按下某键", "value": "keydown"},
  {"label": "放开某键", "value": "keyup"},
  {"label": "滚轮被滚动", "value": "mousewheel"}
])
</script>