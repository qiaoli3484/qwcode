
import qlrowR from './qlrow/indexR.vue'
import qlrowV from './qlrow/indexV.vue'

import qlformR from './qlform/indexR.vue'
import qlformV from './qlform/indexV.vue'

import qlinputR from './qlinput/indexR.vue'
import qlinputV from './qlinput/indexV.vue'

import qlnumberR from './qlnumber/indexR.vue'
import qlnumberV from './qlnumber/indexV.vue'

import qldateR from './qldate/indexR.vue'
import qldateV from './qldate/indexV.vue'

import qltextareaR from './qltextarea/indexR.vue'
import qltextareaV from './qltextarea/indexV.vue'


import qlsingleR from './qlsingle/indexR.vue'
import qlsingleV from './qlsingle/indexV.vue'

import qlswitchR from './qlswitch/indexR.vue'
import qlswitchV from './qlswitch/indexV.vue'

import qlfileR from './qlfile/indexR.vue'
import qlfileV from './qlfile/indexV.vue'

import qlimageR from './qlimage/indexR.vue'
import qlimageV from './qlimage/indexV.vue'

import qlpictureR from './qlpicture/indexR.vue'
import qlpictureV from './qlpicture/indexV.vue'


import qlselectR from './qlselect/indexR.vue'
import qlselectV from './qlselect/indexV.vue'

import qlbtnR from './qlbtn/indexR.vue'
import qlbtnV from './qlbtn/indexV.vue'

import qltabsR from './qltabs/indexR.vue'
import qltabsV from './qltabs/indexV.vue'

import qlcardR from './qlcard/indexR.vue'
import qlcardV from './qlcard/indexV.vue'

import qltextR from './qltext/indexR.vue'
import qltextV from './qltext/indexV.vue'

import qlhtmlR from './qlhtml/indexR.vue'
import qlhtmlV from './qlhtml/indexV.vue'

import qldividerR from './qldivider/indexR.vue'
import qldividerV from './qldivider/indexV.vue'

import qlbuttonR from './qlbutton/indexR.vue'
import qlbuttonV from './qlbutton/indexV.vue'

import qlchildR from './qlchild/indexR.vue'
import qlchildV from './qlchild/indexV.vue'

import qltableR from './qltable/indexR.vue'
import qltableV from './qltable/indexV.vue'

import qldataTableR from './qldataTable/indexR.vue'
import qldataTableV from './qldataTable/indexV.vue'

import qlqrcodeR from './qlqrcode/indexR.vue'
import qlqrcodeV from './qlqrcode/indexV.vue'

import qlbarcodeR from './qlbarcode/indexR.vue'
import qlbarcodeV from './qlbarcode/indexV.vue'

import qleditR from './qledit/indexR.vue'
import qleditV from './qledit/indexV.vue'

import qlechartR from './qlechart/indexR.vue'
import qlechartV from './qlechart/indexV.vue'

import qlreportR from './qlreport/indexR.vue'
import qlreportV from './qlreport/indexV.vue'

import qleditorR from './qleditor/indexR.vue'
import qleditorV from './qleditor/indexV.vue'

let idd=0
export function getId(){
   return idd++
}



export function getFieldview(cpnt,widgets,isedit){
    if(typeof cpnt.fid=='undefined'){
        switch (cpnt.type) {
          case 'form':
            return qlformV
          case 'table':
            return qltableV
          case 'qlbtn':
            return qlbtnV
          case 'row':
            return qlrowV
          case 'tabs':
            return qltabsV
          case 'card':
            return qlcardV
          case 'text':
            return qltextV 
          case 'html':
            return qlhtmlV 
          case 'divider':
            return qldividerV 
          case 'button':
            return qlbuttonV
          case 'qlchild':
            return qlchildV
          case 'dataTable':
            return qldataTableV
          case 'barcode':
            return qlbarcodeV
          case 'qrcode':
            return qlqrcodeV
          case 'picture':
            return qlpictureV
          case 'echarts':
            return qlechartV
          case 'report':
            return qlreportV  
          case 'qleditor':
            return qleditorV  
              
          // 添加其他组件类型的处理逻辑
        }
    }
    if(isedit){
      return qleditV
    }

    switch (widgets[cpnt.fid].type) {
      case 'select':
        return qlselectV
      case 'textarea':
        return qltextareaV
      case 'input':
        return qlinputV
      case 'input-number':
        return qlnumberV
      case 'date':
        return qldateV
      case 'single':
          return qlsingleV  
      case 'switch':
          return qlswitchV  
      case 'file':
        return qlfileV     
      case 'image':
        return qlimageV
      case 'table':
        return qltableV
      case 'qlbtn':
        return qlbtnV
      case 'tabs':
        return qltabsV
      case 'card':
        return qlcardV
      case 'text':
        return qltextV  
      case 'html':
        return qlhtmlV
      case 'divider':
        return qldividerV
      case 'divider':
        return qldividerV
      case 'button':
        return qlbuttonV
      case 'qlchild':
        return qlchildV
      case 'row':
        return qlrowV
      case 'dataTable':
        return qldataTableV
      case 'barcode':
        return qlbarcodeV
      case 'qrcode':
        return qlqrcodeV
      // 添加其他组件类型的处理逻辑
    }
}


export function getFieldComponent(cpnt,widgets,isedit){
  if(typeof cpnt.fid=='undefined'){
    switch (cpnt.type) {
      case 'table':
        return qltableR
      case 'qlbtn':
        return qlbtnR
      case 'row':
        return qlrowR
      case 'tabs':
        return qltabsR
      case 'card':
        return qlcardR
      case 'text':
        return qltextR 
      case 'html':
        return qlhtmlR 
      case 'divider':
        return qldividerR 
      case 'button':
        return qlbuttonR
      case 'qlchild':
        return qlchildR
      case 'dataTable':
        return qldataTableR
      case 'barcode':
        return qlbarcodeR
      case 'qrcode':
        return qlqrcodeR
      case 'picture':
        return qlpictureR
      case 'echarts':
        return qlechartR
      case 'report':
        return qlreportR 
      case 'qleditor':
        return qleditorR  
      // 添加其他组件类型的处理逻辑
    }
  }
 
  if(isedit){
    return qleditR
  }
  //field.type
  switch (widgets[cpnt.fid].type) {
    case 'form':
      return qlformR
    case 'select':
      return qlselectR
    case 'textarea':
      return qltextareaR
    case 'input':
      return qlinputR
    case 'input-number':
      return qlnumberR
    case 'date':
      return qldateR
    case 'single':
        return qlsingleR  
    case 'switch':
        return qlswitchR  
    case 'file':
      return qlfileR     
    case 'image':
      return qlimageR
    case 'table':
      return qltableR
    case 'qlbtn':
      return qlbtnR
    case 'row':
      return qlrowR
    case 'html':
      return qlhtmlR 
    case 'button':
      return qlbuttonR
    case 'dataTable':
      return qldataTableR
    case 'qlchild':
      return qlchildR
    // 添加其他组件类型的处理逻辑
  }
}

export function cloneEl(val,detail){
  switch (val.type) {
    case 'row':
      return {id:uuid(),type:"row",selected:false,...rowConf(),childs:[{...colConf()}]}
    case 'table':
      return {id:uuid(),label:"表格",type:"table",width:"100%",trs:1,tds:1,selected:false,childs:[{...trConf(tdConf(0,0))}]}
    case 'tabs':
      return {id:uuid(),label:"标签页",type:"tabs",selected:false,childs:[{...tabPane()}]}
    case 'text':
      return {id:uuid(),label:"静态文本",type:"text",text:'静态文本',textalign:'left',fontsize:'13px',fontColor:'',fontWeight:'normal',selected:false}
    case 'picture':
      return {id:uuid(),label:"静态图片",type:"picture",url:'静态文本',width:'100px',height:'100px',fit:'scale-down',selected:false}
    case 'html':
      return {id:uuid(),label:"html",type:"html",html:'<b>html</b>',selected:false}
    case 'html':
      return {id:uuid(),label:"富文本",type:"qleditor",html:'<b>html</b>',selected:false}
    case 'divider':
      return {id:uuid(),label:"分割线",type:"divider",text:'分割线',textalign:'center',selected:false}
    case 'button':
      return {id:uuid(),label:"按钮",type:"button",csize:'small',ctype:'primary',onMounted:'',onClick:'',selected:false}
    case 'dataTable':
      return {id:uuid(),label:"数据表格",type:"dataTable",selected:false,...dataTableConf()}
    case 'barcode':
      return {id:uuid(),label:"条形码",type:"barcode",selected:false,text:'123',format:'CODE39',displayValue:true,height:100}  
    case 'qrcode':
      return {id:uuid(),label:"二维码",type:"qrcode",selected:false,text:'二维码',Level:'',width:0}   
    case 'echarts':
      return {id:uuid(),label:"图表",type:"echarts",selected:false,width:'100px',height:'100px',...echartsConf()}
    case 'report':
      return {id:uuid(),label:"报表",type:"report",selected:false,width:'100px',height:'100px',...echartsConf()}
    case 'qlchild'://子表
      return {id:uuid(),selected:false,fid:'c'+val.id,hidelabel: false}//子表参数重新规划
    default:
      if(detail){
          return {selected:false,textalign:'left',fontsize:'13px',fontColor:'',fontWeight:'normal',label:val.label,fid:val.id,id:uuid(),alias:val.alias,type:val.type,showfmt:'defaule'}
      }
      return {id:uuid(),selected:false,fid:val.id,hidelabel: false}
  }
}

export function tabPane(){
  return {label:"111",disabled:false,name:'start',lazy:'false',childs:[]}
}

export function rowConf(){
  return {label:"栅格",gutter:0,justify:'start',align:'middle'}
}

export function colConf(){
  return {id:uuid(),type:'col',span:4,offset:0,push:0,pull:0,selected:false,bgColor:'',childs:[]}
}

export function trConf(j){
  if(j){
  return {id:uuid(),type:'tr',selected:false,childs:[{...j}]}
  }
  return {id:uuid(),type:'tr',selected:false,childs:[]}
}

export function tdConf(r,d){
  return {id:uuid(),type:'td',selected:false,colspan:1,rowspan:1,trn:r,tdn:d,width:'',bgColor:'',childs:[]}
}


export function swapArray(arr, index1, index2) {
 
  arr[index1] = arr.splice(index2, 1, arr[index1])[0];
  //const ab=JSON.stringify(arr[index1])
  //arr.splice(index1, 1)
  //arr.splice(index2, 0,JSON.parse(ab))
   return arr;
}
export const isUndefined=(val)=>{
  const aaa="row|table|tabs|text|picture|html|qleditor|divider|button|dataTable|barcode|qrcode|echarts|report|"
  return aaa.indexOf(val.type)>-1
}


function uuid() {
  let uuid = new Date().getTime() + Math.random().toString(36).substr(2);
  uuid+=Date.now().toString(36)+ Math.random().toString(36).substr(2);
  return uuid;
}

function dataTableConf(){
  return{
   width:'100%',
   height:'300px',
   stripe:true,
   border:true,
   csize:'small',
   index:false,
   rowKey:"ID",
   showSummary:false,
   data:{temp:0,by:'',sort:'',filter:'',enfilter:false,filters:[]},
   rows:[]
  }
}
function echartsConf(){
  return{
    name:'',
    data:{temp:0,by:'',sort:'',filter:'',enfilter:false,filters:[]},
  }
}
