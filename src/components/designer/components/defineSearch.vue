<template>
      <el-row type="flex" align="middle">
        <el-col :span="2">名称:</el-col>
        <el-col :span="10">
          <el-input size="small" v-model="table.name"></el-input>
        </el-col>
        <el-col :span="4">，引用标记：</el-col>
        <el-col :span="5">
          <el-select v-model="table.label" size="small">
            <el-option v-for="item in 30" :key="item" :label="'$D' + item" :value="'$D' + item"></el-option>
          </el-select>
        </el-col>
      </el-row>
      <div class="qcard">
        <div class="qcard_title">
          查询:
          <el-select v-model="table.temp" placeholder="请选择" size="small" @change="templetChange">
            <el-option v-for="item in templets" :key="item.name" :label="item.name" :value="item.id"></el-option>
          </el-select>
        </div>
        <el-checkbox v-model="table.enfilter">过滤器：</el-checkbox>
        <el-select v-model="table.filter" placeholder="请选择" class="w-180" size="small">
          <el-option v-for="item in auto_filters" :key="item.id" :label="item.name" :value="item.id"></el-option>
        </el-select>
        <br />
        <div class="t-c" v-for="(item, index) in table.filters">
          <el-radio-group v-model="item.radio" style="margin-left: 140px">
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
            <el-option v-for="itemA in fields" :key="itemA.id" :label="itemA.label" :value="itemA.alias">
            </el-option>
          </el-select>
          <el-button type="text" @click="addautoComplete(index)">增加</el-button>
          <el-button type="text" @click="delautoComplete(index)">删除</el-button>
        </div>
      </div>
      <el-row type="flex" align="middle">
        <el-col :span="2">排序</el-col>
        <el-col :span="10">
          <el-select v-model="table.by" size="small">
            <el-option v-for="itemA in fields_C" :key="itemA.id" :label="itemA.name" :value="itemA.alias"></el-option>
            <el-option  label="[创建人]" value="Builder"></el-option>
            <el-option  label="[更新人]" value="LastMender"></el-option>
            <el-option  label="[更新时间]" value="LastModified"></el-option>
            <el-option  label="[保密级别]" value="SecLevel"></el-option>
            <el-option  label="[ID]" value="ID"></el-option>
          </el-select>
        </el-col>
        <el-col :span="5">
          <el-select v-model="table.sort" size="small">
            <el-option label="从小到大" value="asc"></el-option>
            <el-option label="从大到小" value="desc"></el-option>
          </el-select>
        </el-col>
      </el-row>

      <div class="qcard">
        <div class="qcard_title">引用代码:</div>
        <el-checkbox-group v-model="table.list">
            <el-checkbox v-for="itemA in fields_C" :key="itemA.id" :label="itemA.name" :value="itemA.alias"></el-checkbox>
        </el-checkbox-group>
      </div>
      <div style="text-align: center">
        <el-button type="primary" size="small" @click="addtable()">确定</el-button>
        <el-button type="primary" size="small" @click="searchVisible = false">取消</el-button>
      </div>
    <!--</el-dialog>-->
</template>

<script setup>
    import {ref,getCurrentInstance,defineProps,computed,defineEmits,nextTick} from 'vue'
    //import {ApiGetFields,ApiGetFilters} from '@/utils/api.js'
    const props=defineProps({
        tid:Number,
        tables:Array,
        templets:Array,
        fields:Array,
        table:Object,
    })
    const emit=defineEmits(['update:tables','update:table'])

    const {proxy}=getCurrentInstance()
    const searchVisible=ref(false)

    const table=computed({
        get:()=>{
            return props.table
        },
        set:(val)=>{
            emit('update:table',val)
        }
    })

    const current=ref(-1)
    const tables=computed({
        get(){
          return props.tables
        },
        set(val){
          emit('update:tables',val)
        }
    })

    const opentable=()=>{
        cleartable();
        if (current.value == -1) {
        table.value.label = "$D" + (tables.value.length + 1);
        }
        searchVisible.value = true;
    }

    const auto_filters=ref([])
    const templetChange=(val) => {
      getFields(val, 'fields_C')
      ApiGetFilters(val).then(res=>{
         auto_filters.value=res.data
      })
    }

    const fields_C=ref([])
    const edittable=(index)=>{
          current.value = index;
          table.value = tables.value[index];
          templetChange(table.value.temp)
          searchVisible.value = true;
        }
    const deltable=(index)=>{
          tables.value.splice(index, 1);
          savetable();
        }
    const getFields=(val, key)=>{
          ApiGetFields(val).then((res) => {
            if (key=='fields_C'){
                fields_C.value=res.data
            }
          });
        }

    const addautoComplete=(index)=>{
          table.value.filters.splice(index + 1, 0, {
            alias: "",
            value: "",
            radio: "and",
            compare: "等于",
          });
        }
    const delautoComplete=(index)=>{
          if (table.value.filters.length == 1) {
            table.value.filters = [
              { alias: "", value: "", radio: "and", compare: "等于" },
            ];
            return;
          }
          table.value.filters.splice(index, 1);
        }
    const addtable=()=>{
          if (table.value.name == "") {
            ElMessage({
              showClose: true,
              message: "名称不能为空",
              type: "error",
            });
            return;
          }
          if (table.value.label == "") {
            ElMessage({
              showClose: true,
              message: "引用标记不能为空",
              type: "error",
            });
            return;
          }

          if (current.value == -1) {
            tables.value.push(table.value);
          } else {
            tables.value[current.value] = table.value;
          }
          searchVisible.value = false;
          cleartable();
          //savetable();
        }
  
    const cleartable=()=>{
          table.value = {
            ID: 0,
            name: "",
            label: "",
            temp: "",
            by: "",
            sort: "",
            filter: "",
            enfilter: false,
            filters: [{ alias: "", value: "", radio: "and", compare: "等于" }], //输入提示和查询
            list: [],
          };
        }
    const savetable=()=>{
          current.value = -1;
          proxy.$post("/layout/tables?tid=" + props.tid,JSON.stringify(tables.value)).then((res) => {
            ElMessage({ showClose: true, message: res.msg, type: "success" });
          });
        }
</script>