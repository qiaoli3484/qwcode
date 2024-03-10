app.component('ql-autocomplete', {
    emits: ['change'],
    props:{
        value:[String, Number],
        placeholder:String,
    },
    data: function () {
        return {
            visible:false
        }
    },
    template: ` <el-autocomplete v-else v-model="data" size='small'  autofocus></el-autocomplete>`,
    methods:{
        isShow(){
            this.visible=true
        },
        outClick(e){
            let dropRef= this.$el
            if(!dropRef.contains(e.target)&&this.visible){
                this.visible=false
            }
        }
    },
    mounted(){
        document.addEventListener('click',this.outClick)
    },
    destroyed(){
        document.removeEventListener('click',this.outClick)
    },
    computed:{
        data:{
            get(){
                return this.value
            },
            set(val){
                this.$emit("change",val)
            }
        }
    }
})