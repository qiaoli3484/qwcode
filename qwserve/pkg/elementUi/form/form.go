package form

import (
	"fmt"
	"simpleTool/pkg/comm"
	"simpleTool/pkg/vue/elementUi/base"
	"strconv"
	"strings"
)

type Elform struct {
	base.Base
	view     strings.Builder
	rules    strings.Builder
	data     strings.Builder
	methods  strings.Builder
	mounted  strings.Builder
	computed strings.Builder
	use      map[string]string
}

func InitForm(key string) *Elform {
	b := &Elform{
		use: make(map[string]string),
	}
	b.Setkey(key)

	return b
}

func (b *Elform) Default() *Elform {
	b.data.WriteString(b.Getkey() + ":{ID:0")
	b.rules.WriteString(b.Getkey() + "Rules:{")
	b.methods.WriteString(fmt.Sprintf(`
	submitForm() {
        return this.$refs['ref_%[1]s'].validate()
      },
      resetForm() {
        this.$refs['ref_%[1]s'].resetFields();
      }`, b.Getkey()))

	b.view.WriteString(fmt.Sprintf(`<el-form ref="ref_%[1]s" :model="%[1]s"  :rules="%[1]sRules" size="small" >`, b.Getkey()))
	return b
}

func (e *Elform) AddEl(el comm.FEler, require bool, vif string) *Elform {
	e.view.WriteString("<el-form-item label=\"" + el.GetName())

	if len(vif) != 0 {
		e.view.WriteString("\" v-if=\"" + strings.ReplaceAll(vif, "this.", e.Getkey()+"."))
	}

	if require {
		e.view.WriteString("\" prop=\"" + el.Getkey())
	}
	e.view.WriteString("\">")
	e.view.WriteString(el.View(e.Getkey()))
	e.view.WriteString("</el-form-item>")

	//rules: {name: []}
	if require {
		e.rules.WriteString(el.GetRule())
	}
	/*name: [
		{ required: true, message: 'Please input Activity name', trigger: 'blur' },
		{ min: 3, max: 5, message: 'Length should be 3 to 5', trigger: 'blur' },
	  ],*/

	//data

	e.data.WriteString("," + el.Data())

	e.methods.WriteString(el.Method())
	e.mounted.WriteString(el.Init())
	e.computed.WriteString(el.Computed())

	for kk, vv := range el.Component() {
		e.use[kk] = vv
	}

	return e
}

func (e *Elform) View(alias string) string {
	/*if e.nview {
		return ""
	}
	var view strings.Builder
	//:inline=\"true\" label-width=\"100px\"
	view.WriteString(fmt.Sprintf("<el-form :model=\"%[1]s\" inline-message  ref=\"ref_form\" >", e.alias))

	view.WriteString("<el-row align='middle'>\n")
	for _, v := range e.Items {
		if v.hide {
			view.WriteString(fmt.Sprintf("<ql-form-item %s %s v-show='false'>\n", v.el.This(), v.att))
		} else {
			view.WriteString(fmt.Sprintf("<el-col %s class='t-c'>%s</el-col><ql-form-item prop=\"%s\" %s :rules=\"%s\" >\n", v.att, v.label, v.alias, v.att, v.el.Rule()))
		}
		//view.WriteString("<template #error='{ error }'><ql-err-f v-model='error'></ql-err-f></template>")
		view.WriteString(v.el.View())
		//view.WriteString("<template #error='{ error }'><ql-err-f v-model='error'></ql-err-f></template>")
		view.WriteString("</ql-form-item>\n")

	}*/
	return e.view.String() + "</el-form>"

}

func arradd(arr ...string) int {
	sum := 0
	for _, v := range arr {
		n, _ := strconv.Atoi(v)
		sum += n
	}
	return sum
}

func (e *Elform) Data() string {

	/*var data, data2 string
	for _, v := range e.Items {
		data += v.el.Data()
		data2 += v.el.Data2()
	}*/
	//fmt.Sprintf("%s:{%sID:0},saveload:false,%s", e.alias, data, data2)

	return e.data.String() + "}," + e.rules.String() + "c:[]}"
}

func (e *Elform) Method() string {
	return e.methods.String()
}

func (e *Elform) Component() map[string]string {
	return e.use
}

func (e *Elform) Computed() string {
	return ""
}
func (e *Elform) Init() string {
	return ""
}

// func (e *Elform) Methods() string {
// 	var m strings.Builder
// 	for _, v := range e.Items {
// 		m.WriteString(v.el.Methods())
// 	}
// 	var fun string

// 	/*

// 			editdata() {
// 		    if(this.saveload)return;
// 		    this.$refs['ref_%[1]s'].validate( async (valid) => {
// 		      if (valid) {
// 		        this.saveload=true
// 		        if(this.%[1]s.ID==0){
// 		            const res= await this.$post('/api/edit?tid='+this.tid,JSON.stringify({%[1]s:this.%[1]s %[2]s})).catch(()=>{this.saveload=false})
// 		              this.%[1]s=res.data.%[1]s
// 		              %[3]s
// 		              this.$message({showClose:true,message:res.msg,type:'success'})
// 		              this.saveload=false
// 		        } else {
// 		            const res= await this.$put('/api/edit?tid='+this.tid,JSON.stringify({%[1]s:this.%[1]s %[2]s})).catch(()=>{this.saveload=false})
// 		              this.%[1]s=res.data.%[1]s
// 		              %[3]s
// 		              this.$message({showClose:true,message:res.msg,type:'success'})
// 		              this.saveload=false
// 		        }
// 		      } else {
// 		        return false;
// 		      }
// 		    });
// 		  },

// 	*/

// 	fun = `
//   editfield(val) {
//     if(this.saveload)return;
//     this.$refs['ref_%[1]s'].validate( async (valid) => {
//       if (valid) {
//         this.saveload=true
//         if(this.%[1]s.ID==0){
//               this.saveload=false
//         } else {
//             const res= await this.$put('/api/edit?tid='+this.tid,JSON.stringify({%[1]s:this.%[1]s %[2]s})).catch(()=>{this.saveload=false})
//               this.%[1]s=res.data.%[1]s
//               console.log(window.opener)
//               window.parent.document.getElementById(this.getQuery("id")).innerHTML=res.data.%[1]s[val];
//               this.saveload=false
//               this.closedown()
//         }
//       } else {
//         return false;
//       }
//     });
//   },
//   initdata(){
//     if(this.rid=='')return;
//       this.$get('/api/edit?tid='+this.tid+'&rid='+this.rid).then(res=>{
//         this.%[1]s=res.data.%[1]s
//         %[3]s
//       })
//   },
//   cleardata(){
//     this.%[1]s={%[4]sID:0}
//   },
//   export_excel(){

//   },
//   upload_excel(){

//   },
//   deldata(){
//     this.del_sure('当前数据',()=>{
//       this.saveload=true
//       this.$del('/api/edit?tid='+this.tid+'&rid='+this.%[1]s.ID).then(res=>{
//         this.$message({showClose:true,message:res.msg,type:'success'})
//         this.cleardata()
//         this.saveload=false
//       }).catch(()=>{
//         this.saveload=false
//       })
//     })
//   },
//   custom_(name,val){
//       this.$post("/api/"+val+"/"+name+"?tid="+this.tid,JSON.stringify({rid:+this.%[1]s.ID})).then(res=>{
//         this.$message({showClose:true,message:res.msg,type:'success'})
//       })
//   },
//   async flow_submit(val,m,t){
//     if(this.%[1]s.ID==0){
//       this.editdata()
//     }
//     if(t&&this.NeedPlanTime==""){
//       this.$message({showClose:true,message:"必须填入计划完成日期",type:'error'})
//       return
//     }
//     const msg = await this.menta().catch()
//     if(m&&!msg.value){
//       this.$message({showClose:true,message:"必须填入意见",type:'error'})
//       return
//     }
//     this.$post("/api/flowwork?tid="+this.tid+"&rid="+this.%[1]s.ID+"&param="+val,JSON.stringify({msg:msg.value,NeedPlanTime:this.NeedPlanTime})).then(res=>{
//       this.$message({showClose:true,message:res.msg,type:'success'})
//       this.closedown()
//     })
//   },
//   menta(){
//     return this.$prompt('请输入意见', '提示', {
//         confirmButtonText: '确定',
//         cancelButtonText: '取消',
//       })
//   },
//   `
// 	var str1, str2 string
// 	for i, _ := range e.vals {
// 		str1 += fmt.Sprintf(",%[1]s:this.$refs.ref_%[1]s.data", e.vals[i])
// 		str2 += fmt.Sprintf("this.$refs.ref_%[1]s.data=res.data.%[1]s\n", e.vals[i])
// 	}

// 	var data string
// 	for _, v := range e.Items {
// 		data += v.el.Data()
// 	}

// 	m.WriteString(fmt.Sprintf(fun, e.alias, str1, str2, data))
// 	return m.String()
// }

func (e *Elform) Mounted() string {
	return ""
}

func (e *Elform) Rule() string {
	return ""
}

// 嵌入时用
func (e *Elform) Model(talias, val string) {

}
