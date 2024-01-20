<template>
    <el-form-item label="数据源">
        <el-select v-model="el.data.temp" placeholder="请选择" size="small" @change="templetChange">
            <el-option v-for="item in templets" :key="item.name" :label="item.name" :value="item.id"></el-option>
        </el-select>
    </el-form-item>

    <el-form-item label="表格列">
        <el-button @click="showtablelist()">编辑</el-button>
    </el-form-item>

    <el-form-item label="数据条件">
        <el-button @click="opentable()">编辑</el-button>
    </el-form-item>

    <el-dialog title="数据条件" v-model="searchVisible" top="5vh" width="650px" :modal="false" :modal-append-to-body="false" draggable>
      <div class="qcard">
        <el-checkbox v-model="el.data.enfilter" @change="changefilter">过滤器：</el-checkbox>
        <el-select v-model="el.data.filter" placeholder="请选择" class="w-180" size="small">
          <el-option v-for="item in auto_filters" :key="item.id" :label="item.name" :value="item.name"></el-option>
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


    <!-- 设置表格 -->
    <el-dialog title="表格列编辑" v-model="tableVisible" width="750px" :modal="false"
      :modal-append-to-body="false" :close-on-click-modal="false" append-to-body draggable >
        <el-table :data="el.rows" size="small" height="300" border ref="ref_Table" highlight-current-row row-key="alias">
            <el-table-column width="60" align="center">
                <template v-slot>
                    <span class="drag-btn" style="cursor:move">
                        <el-icon><rank /></el-icon>
                    </span>
                </template>
            </el-table-column>
            <el-table-column label="序号" type="index" width="50" align="center"></el-table-column>
            <el-table-column align="center" label="名称" prop="alias" width="72">
                <template #default="scope">
                    <el-select v-model="scope.row.alias" placeholder="请选择" size="small" @change="(val)=>fieldChange(val,scope.row)">
                        <el-option v-for="itemA in fields_C" :key="itemA.id" :label="itemA.name" :value="itemA.alias"></el-option>
                    </el-select>
                </template>
            </el-table-column>
            <el-table-column align="center" label="宽度" width="72">
                <template #default="scope">
                   <el-input v-model.trim="scope.row.width" size="small"></el-input>
                </template>
            </el-table-column>
            <el-table-column align="center" label="排序" width="72">
                <template #default="scope">
                    <el-switch v-model="scope.row.resizable" size="small" />
                </template>
            </el-table-column>
            <el-table-column align="center" label="固定" width="72">
                <template #default="scope">
                    <el-select v-model="scope.row.fixed" placeholder="请选择" size="small">
                        <el-option  label="不固定" value=""></el-option>
                        <el-option  label="靠左" value="left"></el-option>
                        <el-option  label="靠右" value="right"></el-option>
                    </el-select>
                </template>
            </el-table-column>
            <el-table-column align="center" label="对齐" width="72">
                <template #default="scope">
                    <el-select v-model="scope.row.align" placeholder="请选择" size="small">
                        <el-option  label="靠左" value="left"></el-option>
                        <el-option  label="居中" value="center"></el-option>
                        <el-option  label="靠右" value="right"></el-option>
                    </el-select>
                </template>
            </el-table-column>
            <el-table-column align="center" label="格式化" width="72">
                <template #default="scope">
                    <el-input v-model.trim="scope.row.formatter" size="small"></el-input>
                </template>
            </el-table-column>
            <el-table-column align="center" label="操作" width="72" resizable>
                <template #default="{row,$index}">
                    <el-button size="small" type="primary" link @click="add_($index,row)" icon="Plus"></el-button>
                    <el-button size="small" type="danger" link @click="del_($index,row)" icon="Minus"></el-button>
                </template>
            </el-table-column>
        </el-table>
    </el-dialog>
</template>

<script setup>
    import {ElMessage} from 'element-plus'
    import {ref,getCurrentInstance,defineProps,computed,defineEmits,nextTick,inject,onMounted} from 'vue'
    import Sortable from 'sortablejs'
    import {ApiGetFields,ApiGetFilters} from '@/utils/api.js'
    import qlFilter from '@/components/ql-filter2/filter.vue'
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
    const opentable=async()=>{
        if(fields_C.value.length==0){
            await getFields(el.value.data.temp)
        }
        searchVisible.value = true;
    }

    const fields_C=ref([])
    const templetChange= async (val)=>{
        await getFields(val)
        el.value.rows=[]
        fields_C.value.forEach((item)=>{
            el.value.rows.push({alias:item.alias,label:item.name,width:"100",resizable:false,fixed:'',align:'center',formatter:''})
        })
    }

    const getFields= async (val)=>{
        
        await ApiGetFields(val).then(res=>{
            fields_C.value=res.data
        })
    }

    const tableVisible=ref(false)
    const showtablelist= async ()=>{
        if(fields_C.value.length==0){
            await getFields(el.value.data.temp)
        }
        tableVisible.value = true;
        nextTick(() => {
            rowDrop();
        });
    }
    let sortable=null
    const ref_Table=ref(null)
    const rowDrop=()=>{
            const tbody = ref_Table.value.$el.querySelector('.el-table__body-wrapper tbody')
            sortable = Sortable.create(tbody, {
            handle:'.drag-btn',
            onEnd({ newIndex, oldIndex }) {
                    const currRow = el.value.rows.splice(oldIndex, 1)[0]
                    el.value.rows.splice(newIndex, 0, currRow)
                    /*proxy.$post('/field/sortfield?tid='+tid.value,JSON.stringify(fields.value.map((item) => item.id))).then((res) => {
                        console.log(res)
                    })*/
                    ref_Table.value.doLayout()
                }
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
    const add_=(index)=>{
        el.value.rows.splice(index+1,0,{alias:"",label:"",width:"100",resizable:false,fixed:'',align:'center',formatter:''})
    }
    const del_=(index)=>{
        if(el.value.rows.length<=1){
            return
        }
        el.value.rows.splice(index, 1);
    }
    const fieldChange=(val,row)=>{
        fields_C.value.forEach((item)=>{
            if(item.alias==val){
                row.label=item.name
                return
            }
        })
    }


    const auto_filters=ref([])
    const getfilters=(val)=>{
        ApiGetFilters(val).then((res) => {
            auto_filters.value=res.data
        })
    }
    const changefilter=(val)=>{
        if(val){
            getfilters(el.value.data.temp)
            return
        }
        auto_filters.value=[]
        el.value.data.filter=""
    }
    onMounted(()=>{
        if(el.value.data.enfilter){
            getfilters(el.value.data.temp)
        }
    })
</script>