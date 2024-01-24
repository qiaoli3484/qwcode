<template>
    <div ref="editor"></div>
</template>
<script setup>
import { ref,onMounted,defineProps,defineExpose,defineEmits} from 'vue'
import { EditorState,Compartment } from "@codemirror/state";
import { basicSetup } from "codemirror";
import { EditorView,keymap } from "@codemirror/view";
import { javascript } from "@codemirror/lang-javascript";
import {html} from "@codemirror/lang-html"
//import { dracula } from '@uiw/codemirror-theme-dracula';//主题
import { indentWithTab } from "@codemirror/commands";//字符缩进

    const props= defineProps({
        width:{type:String,default:'300px'},
        language:{type:String,default:'javascript'},
        code:String
    })

    const emit=defineEmits(['update:code'])

    const languageMap={
        js:javascript(),
        html:html()
    }


    const editor=ref(null)
    const languageConf = new Compartment()
    let state = EditorState.create({
        extensions: [
        basicSetup,
        languageConf.of(languageMap[props.language]),
        keymap.of([indentWithTab]),
        EditorView.updateListener.of((v) => {
            emit('update:code',v.state.doc.toString())
        }) //监测得到的最新代码
        ],
        // 编辑器中的内容
        doc: props.code,
    });
    const view=ref(null)
    
    const init=()=>{
        view.value = new EditorView({
        state,
        // 编辑器 挂载的dom
        parent: editor.value,
        });
        /*view.value.on('change', () => {
            emit('update:code', view.value.state.doc.toString())
        })*/
    }
    const setcode=(val)=>{
        view.value.dispatch({
          changes: { from: 0,to:view.value.state.doc.length, insert:val },
        });
    }

    const getcode=()=>{
        return view.value.state.doc.toString();
    }

    defineExpose({
        setcode,
        getcode
    })
    //view.state.doc.toString();
    //npm install --save codemirror @codemirror/state @codemirror/view
    //npm install @codemirror/lang-javascript
    onMounted(()=>{
        init()
    })
    
</script>

<style>
.cm-editor { 
     border: 1px solid #ddd;
     height:v-bind(width);
    }
.cm-scroller { overflow: auto; }
</style>