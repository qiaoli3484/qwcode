app.component('ql-select' ,{ 
    template:`
    <div style="display: inline-block;">
        <el-select v-model="value" v-if="!auto" filterable remote :placeholder="placeholder" :remote-method="remoteMethod" :loading="loading"  
            @change="selectchange1" @clear="cleardata" :multiple="multiple" :style="{width:width}" clearable :size="size">

            <el-option v-for="item in options" :key="item.value" :label="item.value" :value="item.value">
            <span style="float: left">{{ item.value }}</span>
            <slot v-bind:item="item"></slot>
            </el-option>

        </el-select>
        <el-autocomplete v-model="value" v-if="auto" :fetch-suggestions="querySearchAsync" :placeholder="placeholder" @select="selectchange2"
            @change="selectchange3" @clear="cleardata" :style="{width:width}" clearable :size="size">
        </el-autocomplete>
    </div>`,
    model:{
        prop:'value',
        event:'change'
    },
    props:{
        value:[String, Array],
        onchange:Function,
        alias:String,
        size:String,
        data:Object,
        width:{type:String,default:"100%"},
        auto:{type:Boolean,default:false},
        placeholder:{type:String,default:""},
        multiple:{type:Boolean,default:false},
    },
    data:function(){
        return{
            loading: false,
            options:[],
        }
    },
    methods:{
        remoteMethod(query) {
            this.loading = true;
            this.callget(`/api/auto/${this.alias}`,{value:query,...this.data}).then(res=>{
                    this.options = res.data
                    this.loading = false;
            }).catch(()=>{this.loading = false})
        },
        selectchange1(val){
            this.$emit('change',val)
            if (typeof(this.onchange)!='undefined'){
                const aa=this.options.find(item =>{ return item.value==val});
                this.onchange(aa)
            }
            this.options=[]
        },
        selectchange2(val){
            this.$emit('change',val.value)
            if (typeof(this.onchange)!='undefined'){
                this.onchange(val)
            }
        },
        selectchange3(val){
            this.$emit('change',val)
        },
        cleardata(){
            this.$emit('change','')
        },
        querySearchAsync(query, cb) {
            var results=[{}];
            this.callget(`/api/auto/${this.alias}`,{value:query,...this.data}).then(res=>{
                results=res.data
                cb(results);
            })
        },
    }
    })