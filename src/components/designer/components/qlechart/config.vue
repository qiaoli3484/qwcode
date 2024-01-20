<template>
    <el-form-item label="数据源">
        <el-select v-model="el.data.temp" placeholder="请选择" size="small" @change="templetChange">
            <el-option v-for="item in templets" :key="item.name" :label="item.name" :value="item.id"></el-option>
        </el-select>
    </el-form-item>

    <el-form-item label="图表名称">
        <el-select v-model="el.name" size="small">
            <el-option v-for="itemA in charts" :key="itemA.id" :label="itemA.name" :value="itemA.name"></el-option>
        </el-select>
    </el-form-item>

    <el-form-item label="数据条件">
        <el-button @click="opentable()">编辑</el-button>
    </el-form-item>

    <el-dialog title="数据条件" v-model="searchVisible" top="5vh" width="650px" :modal="false" :modal-append-to-body="false" draggable>
      <div class="qcard">
        <el-checkbox v-model="el.data.enfilter">过滤器：</el-checkbox>
        <el-select v-model="el.data.filter" placeholder="请选择" class="w-180" size="small">
          <el-option v-for="item in auto_filters" :key="item.id" :label="item.name" :value="item.id"></el-option>
        </el-select>
        <br />
        
        <qlFilter  :fields="fields_C" :tid="el.data.temp" :fieldc="fields" :filter="el.data.filters" />
    
      </div>
      <el-row type="flex" align="middle">
        <el-col :span="2">排序</el-col>
        <el-col :span="10">
          <el-select v-model="el.data.by" size="small">
            <el-option v-for="itemA in fields_C" :key="itemA.id" :label="itemA.name" :value="itemA.alias"></el-option>
            <el-option  label="[创建人]" value="Builder"></el-option>
            <el-option  label="[更新人]" value="LastMender"></el-option>
            <el-option  label="[更新时间]" value="LastModified"></el-option>
            <el-option  label="[保密级别]" value="SecLevel"></el-option>
            <el-option  label="[ID]" value="ID"></el-option>
          </el-select>
        </el-col>
        <el-col :span="5">
          <el-select v-model="el.data.sort" size="small">
            <el-option label="从小到大" value="asc"></el-option>
            <el-option label="从大到小" value="desc"></el-option>
          </el-select>
        </el-col>
      </el-row>
    </el-dialog>
</template>

<script setup>
    import {ElMessage} from 'element-plus'
    import {ref,getCurrentInstance,defineProps,computed,defineEmits,nextTick,inject} from 'vue'
    import qlFilter from '@/components/ql-filter2/filter.vue'
    import {ApiGetFields,ApiGetCharts} from '@/utils/api.js'
    const props=defineProps({
        el:Object,
        templets:Array,
        fields:Array,
    })

    const emit=defineEmits(['update:el'])
    const {proxy}=getCurrentInstance()
    const tid=inject("tid")

    const el=computed({
        get(){
            return props.el
        },
        set(val){
            emit('update:el',val)
        }
    })

    const searchVisible=ref(false)
    const opentable=()=>{
        searchVisible.value = true;
    }

    const fields_C=ref([])
    const templetChange= async (val)=>{
        await getFields(val)
        await getCharts(val)
    }

    const getFields= async (val)=>{
        
        await ApiGetFields(val).then(res=>{
            fields_C.value=res.data
        })
    }

    const charts=ref([])
    const getCharts= async (val)=>{
            await ApiGetCharts(val).then((res) => {
            charts.value=res.data
        });
    }
 
</script>