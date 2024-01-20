<template>
    <el-tabs v-model="activeName" class="demo-tabs" @tab-click="handleClick">
        <el-tab-pane label="组件设置" name="first">
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

                <el-form-item label="显示统计字数" v-if="!isUndefined(el.showWordLimit)">
                    <el-switch v-model="el.showWordLimit" />
                </el-form-item>
                
                <el-form-item label="输入框行数" v-if="el.type=='textarea'&&!isUndefined(el.rows)">
                    <el-input-number v-model="el.rows" :min="1" :max="100" />
                </el-form-item>
                  
                <el-form-item label="数值精度" v-if="!isUndefined(el.bits)">
                    <el-input-number v-model="el.bits" :min="0" :max="20" />
                </el-form-item>
                

                <el-form-item label="切换密码" v-if="!isUndefined(el.showPassword)">
                    <el-switch v-model="el.showPassword" />
                </el-form-item>
                <el-form-item label="禁用" v-if="!isUndefined(el.disabled)">
                    <el-switch v-model="el.disabled" />
                </el-form-item>


                <el-form-item label="宽度" v-if="!isUndefined(el.width)">
                    <el-input  v-model.trim="el.width"></el-input>
                </el-form-item>
                <el-form-item label="高度" v-if="!isUndefined(el.height)">
                    <el-input  v-model.trim="el.height"></el-input>
                </el-form-item>

                <el-form-item label="背景颜色" v-if="!isUndefined(el.bgColor)">
                    <el-color-picker v-model="el.bgColor" :predefine="predefineColors2" />
                </el-form-item>
                
                <el-form-item label="显示类型" v-if="!isUndefined(el.ctype)">
                    <el-select size="small" v-model="el.ctype">
                        <el-option label="默认" value="" />
                        <el-option label="基础" value="primary" />
                        <el-option label="成功" value="success" />
                        <el-option label="信息" value="info" />
                        <el-option label="错误" value="warning" />
                        <el-option label="危险" value="danger" />
                    </el-select>
                </el-form-item>
                <el-form-item label="尺寸" v-if="!isUndefined(el.csize)">
                    <el-radio-group v-model="el.csize">
                        <el-radio-button label="large">最大</el-radio-button>
                        <el-radio-button label="default">适中</el-radio-button>
                        <el-radio-button label="small">最小</el-radio-button>
                    </el-radio-group>
                </el-form-item>
               
                <template v-if="!isUndefined(el.text)">
                    <el-form-item label="文本">
                    </el-form-item>
                    <el-form-item label-width="0px">
                        <el-input
                            v-model="el.text"
                            :autosize="{ minRows: 2, maxRows: 4 }"
                            type="textarea"
                        />
                    </el-form-item>
                </template>

                <el-form-item label="文字位置" v-if="!isUndefined(el.textalign)">
                    <el-radio-group v-model="el.textalign">
                        <el-radio-button label="left">居左</el-radio-button>
                        <el-radio-button label="center">居中</el-radio-button>
                        <el-radio-button label="right">居右</el-radio-button>
                    </el-radio-group>
                </el-form-item>
                <el-form-item label="文字大小" v-if="!isUndefined(el.fontsize)">
                    <el-input size="small" v-model.trim="el.fontsize"></el-input>
                </el-form-item>

                <el-form-item label="文字颜色" v-if="!isUndefined(el.fontColor)">
                    <el-color-picker v-model="el.fontColor" show-alpha :predefine="predefineColors" />
                </el-form-item>
                <el-form-item label="文字加粗" v-if="!isUndefined(el.fontWeight)">
                    <el-radio-group v-model="el.fontWeight">
                        <el-radio-button label="lighter">加细</el-radio-button>
                        <el-radio-button label="normal">默认</el-radio-button>
                        <el-radio-button label="bold">加粗</el-radio-button>
                    </el-radio-group>
                </el-form-item>

                <el-form-item label="显示格式" v-if="!isUndefined(el.fontWeight)">
                    <el-radio-group v-model="el.showfmt">
                        <el-radio-button label="default">默认</el-radio-button>
                        <el-radio-button label="控件">控件</el-radio-button>
                    </el-radio-group>
                </el-form-item>
                
                <template v-if="!isUndefined(el.html)">
                    <el-form-item label="HTML">
                    </el-form-item>
                    <el-form-item label-width="0px">
                        <el-input
                            v-model="el.html"
                            :autosize="{ minRows: 2, maxRows: 4 }"
                            type="textarea"
                        />
                    </el-form-item>
                </template>

                
                <template v-if="el.type=='tabs'">
                    <el-form-item label="选项卡设置">
                    </el-form-item>
                    <el-form-item label-width="0px">
                    <vuedraggable  v-model="el.childs"  @start="drag = true" @end="drag = false"  group="people" item-key="id" handle=".handle" class="row-bg">
                        <template #item="{element,index}">
                            <el-row>
                                <el-col :span="3"><el-icon class="handle action-icon"><Rank /></el-icon></el-col>
                                <el-col :span="18"><el-input size="small" v-model.trim="element.label"></el-input></el-col>
                                <el-col :span="3"><el-button type="danger" size="small" icon="delete" link></el-button></el-col>
                            </el-row>
                        </template>
                    </vuedraggable>
                    <el-row>
                        <el-button type="primary" size="small" link @click="addtab">增加选项卡页</el-button>
                    </el-row>
                    </el-form-item>
               </template>

               <template v-if="el.type=='col'">
                    <el-form-item label="栅格属性设置">
                    </el-form-item>
                    <el-form-item label="栅格宽度" v-if="el.span">
                        <el-input-number v-model="el.span" :min="1" :max="24" />
                    </el-form-item>
                    <el-form-item label="左侧间隔格数" v-if="el.offset">
                        <el-input-number v-model="el.offset" :min="0" :max="24" />
                    </el-form-item>
                    <el-form-item label="右移动格数" v-if="el.push">
                        <el-input-number v-model="el.push" :min="0" :max="24" />
                    </el-form-item>
                    <el-form-item label="左移动格数" v-if="el.pull">
                        <el-input-number v-model="el.pull" :min="0" :max="24" />
                    </el-form-item>
               </template>

               <template v-if="el.type=='row'">
                    <el-form-item label="栅格属性设置">
                    </el-form-item>
                    <el-form-item label="栅格间隔" v-if="el.gutter">
                        <el-input-number v-model="el.gutter" :min="0" :max="12" />
                    </el-form-item>

                    <el-form-item label="水平排列方式" v-if="el.justify">
                        <el-select size="small" v-model="el.justify">
                            <el-option label="start" value="start" />
                            <el-option label="end" value="end" />
                            <el-option label="center" value="center" />
                            <el-option label="space-around" value="space-around" />
                            <el-option label="space-between" value="space-between" />
                            <el-option label="space-evenly" value="space-evenly" />
                        </el-select>
                    </el-form-item>
                    <el-form-item label="垂直排列方式" v-if="el.align">
                        <el-select size="small" v-model="el.align">
                            <el-option label="start" value="start" />
                            <el-option label="end" value="end" />
                            <el-option label="center" value="center" />
                            <el-option label="space-around" value="space-around" />
                            <el-option label="space-between" value="space-between" />
                            <el-option label="space-evenly" value="space-evenly" />
                        </el-select>
                    </el-form-item>
                    <el-form-item label="当前栅格列:">
                        <el-button type="primary" link size="small" @click="addcol">增加列</el-button>
                    </el-form-item>
                    <el-form-item label-width="0px">
                        <el-row v-for="(item,index) in el.childs">
                            <el-col :span="12">栅格{{index}}宽度</el-col>
                            <el-col :span="12">
                                <el-input-number size="small" v-model="item.span" :min="1" :max="12" />
                            </el-col>
                        </el-row>
                    </el-form-item>
               </template>
                

                <dataTableConfig v-if="el.type=='dataTable'" v-model:el="el" :templets="templets" :fields="fields"/>
                <echartConfig v-if="el.type=='echarts'" v-model:el="el" :templets="templets" :fields="fields" />
                <reportConfig v-if="el.type=='report'" v-model:el="el" :templets="templets" :fields="fields" />

                
                <el-form-item label="onMounted" v-if="!isUndefined(el.onMounted)">
                   <el-button  icon="Edit" @click="opencode('onMounted')">编写代码</el-button>
                </el-form-item>
                <el-form-item label="onClick" v-if="!isUndefined(el.onClick)">
                    <el-button icon="Edit" @click="opencode('onClick')">编写代码</el-button>
                </el-form-item>
                <el-form-item label="onChange" v-if="!isUndefined(el.onChange)">
                    <el-button icon="Edit" @click="opencode('onChange')">编写代码</el-button>
                </el-form-item>

            </el-form>
        </el-tab-pane>
        <el-tab-pane label="表单设置" name="second">
            <el-form label-width="100px" label-position="left" size="small">
                <el-form-item label="表单宽度" >
                    <el-input v-model="form.width" />
                </el-form-item>
                <el-form-item label="全局组件大小">
                    <el-radio-group v-model="form.csize">
                        <el-radio-button label="large">最大</el-radio-button>
                        <el-radio-button label="default">适中</el-radio-button>
                        <el-radio-button label="small">最小</el-radio-button>
                    </el-radio-group>
                </el-form-item>
                <el-form-item label="字段标签位置" >
                    <el-radio-group v-model="form.labelPosition">
                        <el-radio-button label="top">居上</el-radio-button>
                        <el-radio-button label="left">居左</el-radio-button>
                        <el-radio-button label="right">居右</el-radio-button>
                    </el-radio-group>
                </el-form-item>
                <el-form-item label="字段标签宽度" >
                    <el-input-number v-model="form.labelWidth" :min="1" :max="200" />
                </el-form-item>
                <el-form-item label="保存后" >
                    <el-radio-group v-model="form.saveMode">
                        <el-radio :label="0">无</el-radio>
                        <el-radio :label="1">新建</el-radio>
                        <el-radio :label="2">返回</el-radio>
                    </el-radio-group>
                </el-form-item>

                <el-form-item label="表单全局函数" >
                    <el-button icon="Edit" @click="opencode('onChange')">编写代码</el-button>
                </el-form-item>
                <el-form-item label="onFormMounted" >
                    <el-button icon="Edit" @click="opencode('onChange')">编写代码</el-button>
                </el-form-item>
                <el-form-item label="onFormDataChange" >
                    <el-button icon="Edit" @click="opencode('onChange')">编写代码</el-button>
                </el-form-item>
            </el-form>
        </el-tab-pane>

        
    </el-tabs>
    <el-dialog v-model="codeVisible" title="组件事件代码" width="60%" :before-close="closecode" draggable>
    {{el.label}}.{{codename}} {
        <qlcode  ref="ref_code"></qlcode>
    }
    </el-dialog>
</template>

<script setup>
    import {ref,defineEmits,defineProps,computed,nextTick,onMounted,inject,defineExpose} from 'vue'
    import {tabPane,colConf} from './components/common.js'
    import vuedraggable from 'vuedraggable'
    import qlcode from './components/qlcode.vue'
    import dataTableConfig from './components/qldataTable/config.vue'
    import echartConfig from './components/qlechart/config.vue'
    import reportConfig from './components/qlreport/config.vue'
    import {ApiGetTemplets,ApiGetFields} from '@/utils/api.js'

    const props=defineProps({
        el:Object,
        form:Object
    })

    const emit= defineEmits(['update:el','update:form'])
    const el=computed({
        get:()=>{
            return props.el
        },
        set:(val)=>{
            emit('update:el',val)
        }
    })
    const form=computed({
        get:()=>{
            return props.form
        },
        set:(val)=>{
            emit('update:form',val)
        }
    })
    const activeName=ref('first')
    const handleClick = (tab, event) => {
    console.log(tab, event)
    }
    const setactiveName=(val)=>{
        activeName.value=val
    }

    const addcol=()=>{
        el.value.childs.push(colConf())
    }

    const addtab=()=>{
        props.el.childs.push({...tabPane()})
    }
    const isUndefined=(val)=>{
        return typeof val =='undefined'
    }
    const ref_code=ref(null)
    const codeVisible=ref(false)
    const codedata=ref('')
    const codename=ref('')
    const opencode=(val)=>{
        codename.value=val
        //codedata.value=
        codeVisible.value=true
        nextTick(()=>{
            ref_code.value.setcode(el.value[val])
        })
    }
    const closecode=(done)=>{
       
        el.value[codename.value]=ref_code.value.getcode()
        codename.value=''
        //codedata.value=''
        done()
    }

    const predefineColors = ref([
        '#ff4500',
        '#ff8c00',
        '#ffd700',
        '#90ee90',
        '#00ced1',
        '#1e90ff',
        '#c71585',
        'rgba(255, 69, 0, 0.68)',
        'rgb(255, 120, 0)',
        'hsv(51, 100, 98)',
        'hsva(120, 40, 94, 0.5)',
        'hsl(181, 100%, 37%)',
        'hsla(209, 100%, 56%, 0.73)',
        '#c7158577',
        ])
    const predefineColors2= ref([
        '#ffffff',
        '#F2F3F5',
        '#dddddd',
        '#ffd700',
        '#90ee90',
        '#00ced1',
        '#1e90ff',
        '#c71585',
        'rgb(255, 69, 0)',
        'rgb(255, 120, 0)',
        ])

    const templets=ref([])
    const getTemplets=()=>{
            ApiGetTemplets().then((res) => {
                templets.value = res.data;
            });
    }

    const tid=inject("tid")

    const fields=ref([])
    const getFields=()=>{
        ApiGetFields(tid.value).then((res) => {
            fields.value= res.data
        });
    }
    const radio=ref(0)
    defineExpose({
        setactiveName
    })
    onMounted(()=>{
        getTemplets()
        getFields()
    })
</script>
