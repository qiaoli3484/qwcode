<template>
    <td class="container" @click.stop="onselect" :class="{dragselect:el.selected}" 
       style="height: 24px;outline: 1px dashed #336699;padding: 3px 3px 2px 3px;cursor:default;" 
       :colspan="el.colspan" 
       :rowspan="el.rowspan"
       :style="{width:el.width,backgroundColor:el.bgColor}"
       ref="refTd"
       @mousemove="onmousemove"
       @mousedown="onmousedown"
       >
        <div v-show="el.selected" class="drag-handler">单元格</div>
        <div v-show="el.selected" class="drag-action"> 
            <el-icon :size="18" color="#fff" class="action-icon" @click.stop="backtable"><Back /></el-icon>
            <el-dropdown  trigger="click" size="small">
                <el-icon :size="18" color="#fff" class="action-icon"><Menu /></el-icon>
                <template #dropdown>
                <el-dropdown-menu>
                    <el-dropdown-item @click.stop="addCol(props.cnum)" >插入左侧列</el-dropdown-item>
                    <el-dropdown-item @click.stop="addCol(props.cnum+1)">插入右侧列</el-dropdown-item>
                    <el-dropdown-item @click.stop="addRow(props.rnum)">插入上方行</el-dropdown-item>
                    <el-dropdown-item @click.stop="addRow(props.rnum+1)" >插入下方行</el-dropdown-item>
                    <el-dropdown-item @click.stop="mergeCol(props.rnum,props.cnum,'left')" divided :disabled="props.cnum==0">合并左侧单元格</el-dropdown-item>
                    <el-dropdown-item @click.stop="mergeCol(props.rnum,props.cnum,'right')" :disabled="props.cnum==tds-1">合并右侧单元格</el-dropdown-item>
                    <el-dropdown-item @click.stop="mergeAllCol(props.rnum)">合并整行</el-dropdown-item>
                    <el-dropdown-item @click.stop="mergeRow(props.rnum,props.cnum,'up')" divided :disabled="props.rnum==0">合并上方单元格</el-dropdown-item>
                    <el-dropdown-item @click.stop="mergeRow(props.rnum,props.cnum,'down')" :disabled="props.rnum+el.rowspan==trs">合并下方单元格</el-dropdown-item>
                    <el-dropdown-item @click.stop="mergeAllRow(props.cnum,el.tdn)">合并整列</el-dropdown-item>
                    <el-dropdown-item @click.stop="splitRow(props.rnum,props.cnum)" :disabled="el.rowspan==1" divided>撤销列合并</el-dropdown-item>
                    <el-dropdown-item @click.stop="splitCol(props.rnum,props.cnum)" :disabled="el.colspan==1" >撤销行合并</el-dropdown-item>
                    <el-dropdown-item @click.stop="delRows(props.rnum,el.trn)" divided>删除整行</el-dropdown-item>
                    <el-dropdown-item @click.stop="delCols(props.rnum,props.cnum,el.tdn)" >删除整列</el-dropdown-item>
                </el-dropdown-menu>
                </template>
            </el-dropdown>
        </div>
        <!-- {{el.trn}}:{{el.tdn}} -->
        <vuedraggable  v-model="el.childs"  @start="drag = true" @end="drag = false"  group="people" item-key="id" handle=".handle" style="height:100%;width:100%;">
            <template #item="{element,index}">
                    <component 
                        :is="getFieldComponent(element,widgets,detail)"
                        v-model:el="el.childs[index]"
                        v-model:parent="el"
                        :num="index">
                    </component>
            </template>
        </vuedraggable>
    </td>
</template>

<script setup>
    import {ref,defineProps,computed,defineEmits,inject} from 'vue'
    import vuedraggable from 'vuedraggable'
    import {getFieldComponent} from '../../../components/common.js'

    const props=defineProps({
        el:Object,
        formData:Object,
        cnum:Number,
        rnum:Number,
    })
    const emit=defineEmits(['update:el'])
    const widgets=inject("widgets")

    const el=computed({
        get:()=>{
            return props.el
        },
        set:(val)=>{
            emit('update:el',val)
        }
    })
    const elSelect=inject("elSelect")
    const onselect=()=>{
        el.value.selected=true
        elSelect(el.value)
    }

    const trs=inject('trs')
    const tds=inject('tds')
    const backtable=inject("backtable")
    const addCol=inject("addCol")
    const addRow=inject("addRow")
    const mergeCol=inject("mergeCol")
    const mergeRow=inject("mergeRow")
    const splitCol=inject("splitCol")
    const splitRow=inject("splitRow")
    const mergeAllCol=inject("mergeAllCol")
    const mergeAllRow=inject("mergeAllRow")
    const delCols=inject("delCols")
    const delRows=inject("delRows")
    const detail=inject("detail")
    const updateW=inject("updateW")
    const setTbw=inject("setTbw")

    const refTd=ref()
    const onrize=ref(false)
    const onmousemove=(e)=>{
        if (e.offsetX > refTd.value.offsetWidth - 10) {
            refTd.value.style.cursor = 'col-resize';
            onrize.value=true
        } else {
            onrize.value=false
            refTd.value.style.cursor = 'default';
        }
        //console.log(e.offsetX,refTd.value.offsetWidth,e,refTd.value)
    }

    const onmousedown=(e)=>{
        if(!onrize.value){
            return
        }

        // 获取按下鼠标时 盒子与页面的距离
        var originBoxX = e.offsetX;
        var originBoxY = e.offsetY;
         
        console.log("按下",originBoxX,originBoxY,e)
        // 获取按下鼠标时 鼠标与页面的距离
        var mouseX = e.pageX;
        var mouseY = e.pageY;
        
        const nextTr=refTd.value.nextElementSibling  //下一个节点
        if(typeof nextTr=='undefined'){
            setTbw("auto") //拖动最右边宽度
        }
        //nextTr.offsetWidth //素的宽度

        const oldwidth=refTd.value.offsetWidth

        const pp=refTd.value.parentNode.parentNode
        pp.style.cursor = 'col-resize';
        pp.style.userSelect='none';//禁止选择元素

        console.log("按下",pp,pp.clientX,pp.clientY)

        let posYold=refTd.value.offsetLeft+e.offsetX

        const div=document.createElement('div')
        div.style.position="absolute"
        div.style.left= posYold+"px"
        div.style.top="0"
        div.style.bottom="0"
        div.style.width="1px"
        div.style.backgroundColor="#CDD0D6"
        pp.appendChild(div)

        // 在页面上移动
        let distanceX=0;
        document.onmousemove=(event)=>{
            // 鼠标滑动的距离 = 鼠标移动后的位置 - 按下鼠标时的位置
            distanceX = event.pageX - mouseX;
            var distanceY = event.pageY - mouseY;

            if(nextTr&&distanceX+10>nextTr.offsetWidth){//不能超过右边宽度
                return
            }
            if( distanceX < 10-oldwidth){//不能超过左边宽度
                return
            }

            // 给元素重新赋值 上左定位的位置
            //box.style.left = originBoxX + distanceX + "px";
            //box.style.top = originBoxY + distanceY + "px";
            div.style.left= posYold + distanceX +"px"
            //console.log(distanceX,distanceY,originBoxX + distanceX)
        }
        document.onmouseup=()=>{
            //this.ismove=false
            //el.value.width=oldwidth+distanceX-7+"px"
            pp.style.userSelect='auto';//解除禁止选择元素
            pp.style.cursor = 'default';
            pp.removeChild(div)
            document.onmousemove=null
            document.onmouseup=null

            updateW(props.cnum,el.value.tdn,oldwidth+distanceX-7+"px")
        }
    }
</script>