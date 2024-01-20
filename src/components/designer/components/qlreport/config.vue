<template>
    <el-form-item label="数据源">
        <el-select v-model="el.data.temp" placeholder="请选择" size="small" @change="templetChange">
            <el-option v-for="item in templets" :key="item.name" :label="item.name" :value="item.id"></el-option>
        </el-select>
    </el-form-item>

    <el-form-item label="报表名称">
        <el-select v-model="el.name" size="small">
            <el-option v-for="itemA in reports" :key="itemA.id" :label="itemA.name" :value="itemA.name"></el-option>
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
        <el-button v-if="el.data.filters.length==0" size="small" @click="addautoComplete(0)">添加条件</el-button>
        <div class="t-c" v-for="(item, index) in el.data.filters">
          <el-radio-group v-model="item.radio">
            <el-radio label="and">并且</el-radio>
            <el-radio label="or">或者</el-radio>
          </el-radio-group><br />
          <el-select v-model="item.alias" placeholder="请选择" class="w-180" size="small">
            <el-option v-for="itemA in fields_C" :key="itemA.id" :label="itemA.name" :value="itemA.alias">
            </el-option>
          </el-select>
          <el-select v-model="item.compare" placeholder="请选择" class="w-80" size="small">
            <el-option label="等于" value="等于"></el-option>
            <el-option label="不等于" value="不等于"></el-option>
            <el-option label="大于" value="大于"></el-option>
            <el-option label="小于" value="小于"></el-option>
            <el-option label="大于等于" value="大于等于"></el-option>
            <el-option label="小于等于" value="小于等于"></el-option>
            <el-option label="包含" value="包含"></el-option>
            <el-option label="不包含" value="不包含"></el-option>
            <el-option label="包含于" value="包含于"></el-option>
            <el-option label="不包含于" value="不包含于"></el-option>
            <el-option label="始于" value="始于"></el-option>
            <el-option label="止于" value="止于"></el-option>
          </el-select>
          当前
          <el-select v-model="item.value" placeholder="请选择" class="w-180" size="small">
            <el-option v-for="itemA in fields" :key="itemA.id" :label="itemA.name" :value="itemA.alias">
            </el-option>
          </el-select>
          <el-button type="text" @click="addautoComplete(index)">增加</el-button>
          <el-button type="text" @click="delautoComplete(index)">删除</el-button>
        </div>
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
        await getReport(val)
    }

    const getFields= async (val)=>{
        await proxy.$get("/fields/get", { tid: val }).then(res=>{
            fields_C.value=res.data
        })
    }

    const addautoComplete=(index)=>{
        el.value.data.filters.splice(index + 1, 0, {
        alias: "",
        value: "",
        radio: "and",
        compare: "等于",
      });
    }
    const delautoComplete=(index)=>{
      el.value.data.filters.splice(index, 1);
    }
  
    const reports=ref([])
    const getReport=(val)=>{
        proxy.$get("/report/all", { tid: val }).then((res) => {
            reports.value=res.data
        });
    }
 
</script>