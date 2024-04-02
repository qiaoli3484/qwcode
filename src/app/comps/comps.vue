<template>
    <div>
        <el-button type="primary" @click="newData">新建</el-button>
    </div>
    <el-input style="width:300px;margin-left:300px" clearable v-model.trim="searchVal" placeholder="输入关键字搜索">
        <template #suffix>
            <el-button type="text" icon="Search" @click="update(0)" />
        </template>
    </el-input>

    <vxe-table size="mini" border :column-config="{resizable: true}" empty-text="没有更多数据了！"
      :row-config="{isHover: true,isCurrent: true,}" :height="tableH"
      :sort-config="{trigger: 'cell'}"
      :data="tableData">
      <vxe-column type="seq" width="60" align="center"></vxe-column>
      <vxe-column width="120" title="操作" align="center">
        <template #default="{ row }">
          <el-button type="text" @click="openDetail(row.id)">编辑</el-button><b />
          <el-button type="text" @click="delData(row.id)">删除</el-button>
        </template>
      </vxe-column>
      
      <vxe-column field="name" title="名称" sortable align="center"></vxe-column>
      <vxe-column field="sex" title="类型" sortable align="center"></vxe-column>
      <vxe-column field="comment" title="备注" sortable align="center"></vxe-column>
    </vxe-table>

    <el-pagination
        v-model:currentPage="page.current"
        v-model:page-size="page.pagesize"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
        :page-sizes="[100, 200, 300, 400]"
        small
        layout="->,total, sizes, prev, pager, next, jumper"
        :total="page.total">
      </el-pagination>

    <el-drawer v-model="drawer" direction="ltr" title="I'm outer Drawer" size="80%">
        <el-button-group>
            <el-button type="primary" @click="newData">新建</el-button>
            <el-button type="primary" @click="saveData">保存</el-button>
            <el-button type="danger" @click="delData(form.id)">删除</el-button>
        </el-button-group>
        <el-tabs v-model="activeName" class="demo-tabs">
            <el-tab-pane label="基础" name="first">
                <el-form :model="form" label-width="auto" style="max-width: 600px">
                    <el-form-item label="名称">
                    <el-input v-model="form.name" />
                    </el-form-item>
                    <el-form-item label="类型">
                        <el-select v-model="form.type" placeholder="请选择组件数据类型">
                            <el-option label="文本型" value="stirng" />
                            <el-option label="数字型" value="number" />
                        </el-select>
                    </el-form-item>
                    <el-form-item label="状态">
                        <el-radio-group v-model="form.state">
                            <el-radio :value="0">启用</el-radio>
                            <el-radio :value="1">停用</el-radio>
                        </el-radio-group>
                    </el-form-item>

                    <el-form-item label="bind参数">
                        <qlparam v-model="form.param.binds" />
                    </el-form-item>

                    <el-form-item label="style参数">
                        <qlparam v-model="form.param.styles" />
                    </el-form-item>

                    <el-form-item label="代码">
                    <el-input v-model="form.code" />
                    </el-form-item>
                    <el-form-item label="备注">
                    <el-input v-model="form.comment" />
                    </el-form-item>
                </el-form>
            </el-tab-pane>
            <el-tab-pane label="代码" name="second">
                <qlcode language="html" ref="ref_code" v-model:code="form.code" />
            </el-tab-pane>
        </el-tabs>
    </el-drawer>
</template>

<script setup>
import { ref,reactive,computed,getCurrentInstance,nextTick} from 'vue'
import qlcode from '@/components/qlcode/qlcode.vue'
import qlparam from './param.vue'

const {proxy} =getCurrentInstance();

const tableData = ref([
  { id: 10001, name: 'Test1', role: 'Develop', sex: 'Man', age: 28, address: 'test abc' },
  { id: 10002, name: 'Test2', role: 'Test', sex: 'Women', age: 22, address: 'Guangzhou' },
  { id: 10003, name: 'Test3', role: 'PM', sex: 'Man', age: 32, address: 'Shanghai' },
  { id: 10004, name: 'Test4', role: 'Designer', sex: 'Women', age: 24, address: 'Shanghai' }
])

const drawer = ref(false)
const ref_code= ref(null)
const openDetail= async (id)=>{
    await getdata(id)
    drawer.value=true
    nextTick(()=>{
        ref_code.value.setcode(form.value.code)
    })
}

const form = ref({
  id:0,
  name: '',
  type:'',
  state:0,
  param:{binds:[],styles:[]},
  code:'',
  comment:''
})
const newData=()=>{
    relaodData()
    drawer.value=true 
}
const saveing=ref(false)
const saveData=async ()=>{
    saveing.value=true
    if (form.value.id==0){
        proxy.$post("api/components/add",JSON.stringify(form.value)).then(res=>{
            form.value.id=res.data
            saveing.value=false
        }).catch(()=>{
            saveing.value=false
        })
    }else{
        proxy.$post("api/components/update",JSON.stringify(form.value)).then(res=>{
            saveing.value=false
        }).catch(()=>{
            saveing.value=false
        }) 
    }
   
}

const relaodData=()=>{
    form.value.id=0
    form.value.name=''
    form.value.type=''
    form.value.state=0
    form.value.param={binds:[],styles:[]}
    form.value.code=
    `     <template>
        输入代码
     </template>
    <script setup>
        import {ref,defineProps,getCurrentInstance} from 'vue';
        const {proxy} =getCurrentInstance();
        const a=ref("测试22")
        const aa=ref("测试22")

        const aaaa=()=>{
            proxy.$get("aaa/bbb")
        }
    <//script>
    `
   
    form.value.comment=''
    if (ref_code.value!=null){
        ref_code.value.setcode(form.value.code)
    }
}

const page=reactive({current: 0,pagesize: 100,total: 0})
const handleSizeChange= (val)=> {
      page.pagesize = val
      update(0)
    }
const handleCurrentChange= (val)=> {
      page.current = val
      update(0)
    }

   const searchVal=ref("")

const update=(n) =>{
      if (n == 1) {
        page.current = 1
      }
      proxy.$post(`api/components/gets?current=${page.current}&pagesize=${page.pagesize}`,JSON.stringify(searchVal.value)).then((res) => {
        tableData.value = res.data
        page.total = res.total
      })
}

const getdata= async (id) =>{
     const {data} =await proxy.$get(`api/components/get?id=${id}`)
        form.value = data
        console.log(form)
}

const tableH=computed(()=>{
    return window.innerHeight - 130;
})

const activeName = ref('first')


const delData=(id) =>{
      proxy.$del(`api/components/del?id=${id}`).then((res) => {
        update(0)
        relaodData()
      })
}
</script>