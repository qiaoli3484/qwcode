app.component('ql-checkbox', {
    emits: ['update:data'],
    props:{
        data:String,
        options:Array,
        size:String,
        disabled:{default:false,type:Boolean}
    },
    template: ` 
        <el-checkbox-group v-model="value" :size="size">
            <el-checkbox v-for="key in options" :label="key" :key="key" />
        </el-checkbox-group>
                `,
    computed:{
        value:{
            get(){
                if (this.data=="")return [];
                return this.data.split(",")
            },
            set(val){
                this.$emit("update:data",val.join(","))
            }
        }
    }
});