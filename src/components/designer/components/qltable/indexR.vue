<template>
    <qlhandle :el="el">
        <template v-slot:header>
            <el-button  link size="small" @click="addtr">行</el-button>
            <el-button  link size="small" @click="addtd">列</el-button>
        </template>
        
        <table class="qtable" style="outline: 1px dashed #336699;" :style="{width:el.width,backgroundColor:el.bgColor}">
            <qltr v-for="(col,index) in el.childs"  v-model:el="el.childs[index]" :rnum="index"/>
        </table>
    </qlhandle>
</template>

<script setup>
    import {ElMessage} from 'element-plus'
    import {ref,defineProps,computed,defineEmits,inject,provide} from 'vue'
    import qltr from './qltr/indexR.vue'
    import qlhandle from '../handle.vue'
    import {trConf,tdConf} from '../../components/common.js'

    const props=defineProps({
        el:Object,
        formData:Object
    })
    const emit=defineEmits(['update:el'])
    
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
    provide('backtable',onselect)
    
    const hisbefore=inject("hisbefore")
    const trs= computed(()=>{
        return props.el.trs
    })
    const tds= computed(()=>{
        return props.el.tds
    })
    provide('trs',trs)
    provide('tds',tds)
    const addtr=()=>{
        hisbefore()
        let tr =trConf()
        for(let i=0;i<tds.value;i++){
            tr.childs.push(tdConf(trs.value,i))
        }
        el.value.childs.push(tr)
        props.el.trs++
    }

    const addtd=()=>{
        hisbefore()
        console.log(trs.value)
        for(let i=0;i<trs.value;i++){
            el.value.childs[i].childs.push(tdConf(i,tds.value))
        }
        props.el.tds++
    }
    
    
    const addCol=(index)=>{
        hisbefore()
        if(index<0)index=0;
        updatetdn(index,(v)=>{if(v.tdn>=index)v.tdn++})
        for(let i=0;i<trs.value;i++){
            el.value.childs[i].childs.splice(index,0,tdConf(i,index))
        }
        props.el.tds++
    }
    provide('addCol',addCol)


    const addRow=(index)=>{
        hisbefore()
        let tr =trConf()
        for(let i=0;i<tds.value;i++){
            tr.childs.push(tdConf(index,i))
        }
        if(index<0)index=0;
        updatetdn(index,(v)=>{if(v.trn>=index)v.trn++})
        el.value.childs.splice(index,0,tr)
        props.el.trs++
       
    }
    provide('addRow',addRow)
     
    const mergeCol=(h,l,val)=>{
        let span=el.value.childs[h].childs[l]
        if(val=='left'){
            if(el.value.childs[h].childs.length==props.el.tds||el.value.childs[h].childs.reduce((t,v)=>{return t+v.colspan},0)==props.el.tds){//当前行没有合并过
                let obj=el.value.childs[h].childs[l-1]
                if(span.rowspan!=obj.rowspan){
                    ElMessage({ showClose: true, message: "无法合并", type: "error" });
                    return
                }
                hisbefore()
                obj.colspan=span.colspan+obj.colspan
                obj.childs.push(...span.childs)
                el.value.childs[h].childs.splice(l,1)
                return
            }

            //复杂情况
            let obj=el.value.childs[h].childs[l-1]
            console.log(span,obj,h,l)
            //需要检测一下 是否能合并
            if(obj.tdn+obj.colspan!=span.tdn){
                ElMessage({ showClose: true, message: "无法合并", type: "error" });
                return
            }
            if(obj.rowspan!=span.rowspan){
                  return
            }
            hisbefore()
            obj.colspan=span.colspan+obj.colspan
            obj.childs.push(...span.childs)
            el.value.childs[h].childs.splice(l,1)
        }
        if(val=='right'){
            let obj=el.value.childs[h].childs[l+1]
             //需要检测一下 是否能合并
            if(span.tdn+span.colspan!=obj.tdn){
                  return
            }

            if(obj.rowspan!=span.rowspan){
                  return
            }
            hisbefore()
            span.colspan=span.colspan+obj.colspan
            span.rowspan=obj.rowspan
            span.childs.push(...obj.childs)
            el.value.childs[h].childs.splice(l+1,1)
        }
    }
    provide('mergeCol',mergeCol)

    const mergeRow=(h,l,val)=>{
        let span=el.value.childs[h].childs[l]
        if(val=='up'){
            //el.trn}}--{{el.tdn
            foo:for(let i=h-1;i>=0;i--){
                
                for(let n=0;n<el.value.childs[i].childs.length;n++){
                    console.log(i,n,el.value.childs[i].childs.length)
                    if (el.value.childs[i].childs[n].tdn==span.tdn){
                        l=n
                        h=i
                        break foo;
                    }
                }
            }

            let obj=el.value.childs[h].childs[l]
            if(obj.colspan!=span.colspan){
                return
            }
            hisbefore()
            obj.rowspan=span.rowspan+obj.rowspan
           
            obj.childs.push(...span.childs)
            delspan(span.id)
        }
        if(val=='down'){
            let aa=0
            for(let n=0;n<el.value.childs[h+span.rowspan].childs.length;n++){
                if (el.value.childs[h+span.rowspan].childs[n].tdn==span.tdn){
                    aa=n
                    break 
                }
            }
            let obj=el.value.childs[h+span.rowspan].childs[aa]
            if(obj.colspan!=span.colspan){
                console.log(obj.colspan,span.colspan,h+span.rowspan,aa)
                ElMessage({ showClose: true, message: "无法合并", type: "error" });
                return
            }
            hisbefore()
            span.rowspan=span.rowspan+obj.rowspan
            span.childs.push(...obj.childs)
            //rowspan 已经改变
            delspan(obj.id)
        }
    }
    provide('mergeRow',mergeRow)

    //拆分列
    const splitCol=(h,l)=>{
        hisbefore()
        let span=el.value.childs[h].childs[l]
        for(let i=1;i<span.colspan;i++){
            let td=tdConf(h,span.tdn+i)
            td.rowspan=span.rowspan
            td.colspan=1
            el.value.childs[h].childs.splice(l+i,0,td)
        }
        span.colspan=1
    }
    provide('splitCol',splitCol)

    //拆分行
    const splitRow=(h,l)=>{
        hisbefore()
        let span=el.value.childs[h].childs[l]
        for(let i=1;i<span.rowspan;i++){
            let td=tdConf(h+i,span.tdn)
            td.rowspan=1
            td.colspan=span.colspan
            el.value.childs[h+i].childs.splice(l,0,td)
        }
        span.rowspan=1
    }
    provide('splitRow',splitRow)

    const delspan=(id)=>{
        for(let i=0;i<el.value.childs.length;i++){
            for(let n=0;n<el.value.childs[i].childs.length;n++){
                const obj=el.value.childs[i].childs[n]
                if(obj.id==id){
                    el.value.childs[i].childs.splice(n,1)
                    return
                }
            }
        }
    }
    const mergeAllCol=(h)=>{
        //检测
       let colspan=0
       let childs=[]
       let num=0
        for(let n=1;n<el.value.childs[h].childs.length;n++){
            const obj=el.value.childs[h].childs[n]
            if(obj.rowspan!=el.value.childs[h].childs[n-1].rowspan){
                return
            }
            colspan=colspan+obj.colspan
            childs.push(...obj.childs)
            num++
        }
        let obj=el.value.childs[h].childs[0]
        if(obj.colspan+colspan!=tds.value){
            console.log(obj.colspan+colspan,tds.value,"数量不符合,可能有合并列")
            ElMessage({ showClose: true, message: "存在列宽不一致的单元格, 不可合并整行", type: "error" });
            return
        }
        hisbefore()
        el.value.childs[h].childs.splice(1,num)
        obj.colspan=obj.colspan+colspan
        obj.childs.push(...childs)
    }
    provide('mergeAllCol',mergeAllCol)

    //合并列
    const mergeAllRow=(l,tdn)=>{
          //检测
        let rowspan=0
        let childs=[]
        let num=[]
        let nextobj=null
        for(let n=0;n<el.value.childs.length;n++){
            for(let i=0;i<el.value.childs[n].childs.length;i++){
                if(el.value.childs[n].childs[i].tdn==tdn){
                    if(nextobj==null){
                        nextobj=el.value.childs[n].childs[i]
                        rowspan=rowspan+nextobj.rowspan
                        num.push(i)
                    }else{
                        const obj=el.value.childs[n].childs[i]
                        if(obj.colspan!=nextobj.colspan&&obj.tdn!=nextobj.tdn){
                            return
                        }
                        rowspan=rowspan+obj.rowspan
                        childs.push(...obj.childs)
                        num.push(i)
                        nextobj=obj
                    }
                }
            }
        }

        let obj=el.value.childs[0].childs[num[0]]
        
        if(rowspan!=el.value.childs.length){
            ElMessage({ showClose: true, message: "存在列宽不一致的单元格, 不可合并整列", type: "error" });
            return
        }
        hisbefore()
        for(let n=1;n<el.value.childs.length;n++){
            el.value.childs[n].childs.splice(num[n],1)
        }
        console.log(rowspan)
        obj.rowspan=rowspan
        obj.childs.push(...childs)
    }
    provide('mergeAllRow',mergeAllRow)

    //删除列
    const delCols=(r,c,tdn)=>{
        if(el.value.childs[0].childs[c].rowspan==props.el.trs){//整列已合并可以删除
            el.value.childs[0].childs.splice(c,1)
            props.el.tds--
            updatetdn(tdn,(v)=>{if(v.tdn>tdn)v.tdn--})
            return
        }

        //检测
        let ok=false
        for(let n=0;n<el.value.childs.length;n++){
            for(let i=0;i<el.value.childs[n].childs.length;i++){
                if(el.value.childs[n].childs[i].tdn==tdn&&el.value.childs[n].childs[i].colspan==el.value.childs[r].childs[c].colspan){//判断一致
                    ok=true
                    let m=el.value.childs[n].childs[i].rowspan
                    if(m>1){//跳过合并行
                        n=n+m-1
                    }
                }
            }
            if(!ok){
                ElMessage({ showClose: true, message: "存在列宽不一致的单元格, 不可删除整列", type: "error" });
                return
            }
            ok=false
        }

        //删除
        for(let n=0;n<el.value.childs.length;n++){
            for(let i=0;i<el.value.childs[n].childs.length;i++){
                if(el.value.childs[n].childs[i].tdn==tdn){
                    let w=el.value.childs[n].childs[i].rowspan
                    el.value.childs[n].childs.splice(i,1)
                    if(w>1){//跳过合并行
                        n=n+w-1
                    }
                }
            }
        }
        props.el.tds--
        updatetdn(tdn,(v)=>{if(v.tdn>tdn)v.tdn--})
    }
    provide('delCols',delCols)

    //删除行
    const delRows=(l,trn)=>{
        if(el.value.childs[l].childs[0].colspan==props.el.tds){//整行合并
            el.value.childs.splice(l,1)
            props.el.trs--
            updatetdn(trn,(v)=>{if(v.trn>trn)v.trn--})
            return
        }

        if(el.value.childs[l].childs.reduce((t,v)=>{return t+v.colspan},0)!=props.el.tds){//整行合并
            ElMessage({ showClose: true, message: "存在行高不一致的单元格, 不可删除整行", type: "error" });
            return
        }

        //检测
        for(let n=0;n<el.value.childs[l].childs.length;n++){
            if(el.value.childs[l].childs[n].rowspan>2){
                ElMessage({ showClose: true, message: "存在行高不一致的单元格, 不可删除整行", type: "error" });
            }
        }

        el.value.childs.splice(l,1)
        props.el.trs--
        updatetdn(trn,(v)=>{if(v.trn>trn)v.trn--})

    }
    provide('delRows',delRows)

    const updatetdn=(tdn,fn)=>{
        //更新tdn
        for(let n=0;n<el.value.childs.length;n++){
            for(let i=0;i<el.value.childs[n].childs.length;i++){
                //if(el.value.childs[n].childs[i].tdn>tdn){
                    fn(el.value.childs[n].childs[i])
                    //el.value.childs[n].childs[i].tdn--
                //}
            }
        }
    }
    //合并 会不会导致 行数 列数变少 ???
    //provide('trs',trs)
    //provide('tds',tds)
    
    //更新宽度
    const updateW=(l,tdn,w)=>{
        for(let n=0;n<el.value.childs.length;n++){
            for(let i=0;i<el.value.childs[n].childs.length;i++){
                if(el.value.childs[n].childs[i].tdn==tdn&&el.value.childs[n].childs[i].colspan==1){
                    el.value.childs[n].childs[i].width=w
                    //el.value.childs[n].childs[i].tdn--
                }
            }
        }
    }
    provide('updateW',updateW)

    //设置表格宽度
    const setTbw=(val)=>{
        el.value.width=val
        //auto
    }
    provide('setTbw',setTbw)
</script>

<style>
.qtable {
    table-layout: fixed;
}
.qtable .el-form-item {
    margin-bottom: 0px!important;
}

</style>


<!-- for (let j = 0; j < table.rows[0].cells.length; j++) {
    
    table.rows[0].cells[j].onmousedown = function () {
       console.log(table.rows[0].cells[j].width)

        //记录单元格
        tTD = this;
        if (event.offsetX > tTD.offsetWidth - 10) {
            tTD.mouseDown = true;
            tTD.oldX = event.x;
            tTD.oldWidth = tTD.offsetWidth;
        }
        //记录Table宽度
        //table = tTD; while (table.tagName != ‘TABLE') table = table.parentElement;
        //tTD.tableWidth = table.offsetWidth;
    };
    table.rows[0].cells[j].onmouseup = function () {
        //结束宽度调整
        if (tTD === undefined) tTD = this;
        tTD.mouseDown = false;
        tTD.style.cursor = 'default';
    };
    table.rows[0].cells[j].onmousemove = function () {
        //更改鼠标样式
        if (event.offsetX > this.offsetWidth - 10)
            this.style.cursor = 'col-resize';
        else
            this.style.cursor = 'default';
        //取出暂存的Table Cell
        if (tTD === undefined) tTD = this;
        //调整宽度
        if (tTD.mouseDown != null && tTD.mouseDown === true) {
            tTD.style.cursor = 'default';
            if (tTD.oldWidth + (event.x - tTD.oldX) > 0)
                tTD.width = tTD.oldWidth + (event.x - tTD.oldX);
            //调整列宽
            tTD.style.width = tTD.width;
            tTD.style.cursor = 'col-resize';
            //调整该列中的每个Cell
            table = tTD;
            //  console.log(tTD.width)
            while (table.tagName !== 'TABLE') table = table.parentElement;
                      //  console.log()
                  table.rows[0].cells[j].style.width = tTD.width + "px"
                  // border = "1px solid #000"
                  // console.log(table.rows[0].cells[j].style.width)
                  // console.log()
            // for (let j = 0; j < table.rows.length; j++) {
            //     //  console.log(tTD.width)
            
            // }
            //调整整个表
            //table.width = tTD.tableWidth + (tTD.offsetWidth – tTD.oldWidth);
            //table.style.width = table.width;
        }
    };
} -->
