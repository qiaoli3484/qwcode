<template>
    <textarea ref="upEditor" name="description" v-model="data" rows="10" cols="80" data-sample-short></textarea>
    <editor />
</template>

<!-- :formatter='el.formatter'
:parser='el.parser -->'

<script setup>
    import editor from '@/components/vue-tinymce.vue'
    import {ref,defineProps,computed,defineEmits,inject,nextTick,onMounted} from 'vue'

    const props=defineProps({
        el:Object,
        data:{type:String,default:''},
    })
    const emit=defineEmits(['update:data'])
    
    const data=computed({
        get:()=>{
            return props.data
        },
        set:(val)=>{
            emit('update:data',val)
        }
    })

    const getContent=()=>{
      return edita.value.document.getBody().getHtml(); //获取textarea的值
    }

    const setContent=(content)=>{
      edita.value.document.getBody().setHtml(content);
    }

    const upEditor=ref(null)
    const edita=ref(null)
    onMounted(()=>{
        nextTick(() => {
            edita.value = CKEDITOR.replace(
            upEditor.value, {height: 480,}
            );
            edita.value.on("change",()=>{

                  console.log("更改")
            })
            /*document.addEventListener("click", (event)=>{
            if(event.target.id=='cke_57_label'){
                bbb.value=!bbb.value
                if(bbb.value){
                    setTimeout(()=>{
                    console.log(document.getElementsByTagName('iframe')[0].contentWindow.document.getElementById("app"))
                    createApp({data() {return {message: 'Hello Vue!!'}}}).mount(document.getElementsByTagName('iframe')[0].contentWindow.document.getElementById("app"))
                    },500)
                }
            }
            })*/
        })
    })
</script>