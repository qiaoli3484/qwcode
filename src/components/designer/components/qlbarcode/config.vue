<template>
    <el-form-item label="条码格式" v-if="el.type=='barcode'">
        <el-select v-model="el.format" size="small" >
            <el-option label="CODE 39" value="CODE39"></el-option>
            <el-option label="CODE 39 Extended" value="CODE39Ext"></el-option>
            <el-option label="CODE 128" value="CODE128"></el-option>
            <el-option label="EAN-13" value="EAN13"></el-option>
            <el-option label="EAN-8" value="EAN8"></el-option>
            <el-option label="Codabar" value="Codabar"></el-option>
            <el-option label="PostNet" value="PostNet"></el-option>
            <el-option label="Standard 2 of 5" value="Sta2of5"></el-option>
            <el-option label="Industrial 2 of 5" value="Ind2of5"></el-option>
            <el-option label="Interleaved 2 of 5" value="Int2of5"></el-option>
            <el-option label="CODE 128-A" value="CODE128A"></el-option>
            <el-option label="CODE 128-B" value="CODE128B"></el-option>
            <el-option label="CODE 128-C" value="CODE128C"></el-option>
        </el-select>
    </el-form-item>
</template>

<script setup>
    import {defineProps,computed,defineEmits} from 'vue'
    const props=defineProps({
        el:Object,
    })

    const emit=defineEmits(['update:el'])
    const el=computed({
        get(){
            return props.el
        },
        set(val){
            emit('update:el',val)
        }
    })
</script>