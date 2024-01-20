<template>
    <div>
        <el-progress v-if="widget.ddlSpecial=='Progress'" :percentage="props.data*100" :stroke-width="20" :color="[{color: '#fc0000', percentage: 20},{color: '#ff5151', percentage: 40},{color: '#f56c6c', percentage: 60},{color: '#e6a23c', percentage: 80},{color: '#5cb87a', percentage: 100}]" :text-inside="true"></el-progress>
        
        <p v-else :style="{textAlign:el.textalign,fontSize:el.fontsize,color:el.fontColor,fontWeight:el.fontWeight,minHeight:data!=''?'unset':'22px'}">
            {{widget.Currency+data+widget.Unit}}
        </p>
    </div>
</template>

<script setup>
    import {ref,defineProps,computed,inject} from 'vue'
    const props=defineProps({
        data:Number,
        el:Object,
        widget:Object,
    })
    //widget.extend.ddlSpecial: "Progress"

    const tid=inject("tid")
    

    const data=computed(()=>{
        //前缀  后缀
        
        if (props.widget.ddlSpecial=='Percentage'){//百分比(%)
            return percentage(props.data)
        }else if (props.widget.ddlSpecial=='permillage'){//千分比(‰)
            return permillage(props.data)
        }else if (props.widget.ddlSpecial=='ThousandthSplit'){//千位分隔符
            return formatterNum(num2str(props.data))
        }

        return num2str(props.data)
        //<el-option label="百分比(%)" value="Percentage"></el-option>
        //<el-option label="百分比(进度)" value="Progress"></el-option>
        //<el-option label="千分比(‰)" value="Permillage"></el-option>
        //<el-option label="千位分隔符" value="ThousandthSplit"></el-option>
    })

    const num2str=(val)=>{
        const num=val.toString()
        if(props.widget.bits==0){
            return num
        }

        if(props.widget.bits>0&&num.indexOf(".")==-1){
            return (val+".") +bzero(props.widget.bits)
        }

        let arr=num.split(".")
        return arr[0]+"."+arr[1]+bzero(props.widget.bits-arr[1].length)
    }

    const bzero=(n)=>{
        let s1=""
        for(let i=0;i<n;i++){
            s1=s1+"0"
        }
        return s1
    }

     // 阿拉伯数字转换大写
     const numberParseChina=(money)=>{
        //汉字的数字
        var cnNums = ['零', '壹', '贰', '叁', '肆', '伍', '陆', '柒', '捌', '玖'];
        //基本单位
        var cnIntRadice = ['', '拾', '佰', '仟'];
        //对应整数部分扩展单位
        var cnIntUnits = ['', '万', '亿', '兆'];
        //对应小数部分单位
        var cnDecUnits = ['角', '分', '毫', '厘'];
        //整数金额时后面跟的字符
        var cnInteger = '整';
        //整型完以后的单位
        var cnIntLast = '圆';
        //最大处理的数字
        var maxNum = 999999999999999.9999;
        //金额整数部分
        var integerNum;
        //金额小数部分
        var decimalNum;
        //输出的中文金额字符串
        var chineseStr = '';
        //分离金额后用的数组，预定义
        var parts;
        if (money == '') {
            return '';
        }
        money = parseFloat(money);
        if (money >= maxNum) {
            //超出最大处理数字
            return '';
        }
        if (money == 0) {
            chineseStr = cnNums[0] + cnIntLast + cnInteger;
            return chineseStr;
        }
        //转换为字符串
        money = money.toString();
        if (money.indexOf('.') == -1) {
            integerNum = money;
            decimalNum = '';
        } else {
            parts = money.split('.');
            integerNum = parts[0];
            decimalNum = parts[1].substr(0, 4);
        }
        //获取整型部分转换
        if (parseInt(integerNum, 10) > 0) {
            var zeroCount = 0;
            var IntLen = integerNum.length;
            for (var i = 0; i < IntLen; i++) {
                var n = integerNum.substr(i, 1);
                var p = IntLen - i - 1;
                var q = p / 4;
                var m = p % 4;
                if (n == '0') {
                    zeroCount++;
                } else {
                    if (zeroCount > 0) {
                        chineseStr += cnNums[0];
                    }
                    //归零
                    zeroCount = 0;
                    chineseStr += cnNums[parseInt(n)] + cnIntRadice[m];
                }
                if (m == 0 && zeroCount < 4) {
                    chineseStr += cnIntUnits[q];
                }
            }
            chineseStr += cnIntLast;
        }
        //小数部分
        if (decimalNum != '') {
            var decLen = decimalNum.length;
            for (var i = 0; i < decLen; i++) {
                var n = decimalNum.substr(i, 1);
                if (n != '0') {
                    chineseStr += cnNums[Number(n)] + cnDecUnits[i];
                }
            }
        }
        if (chineseStr == '') {
            chineseStr += cnNums[0] + cnIntLast + cnInteger;
        } else if (decimalNum == '') {
            chineseStr += cnInteger;
        }
        return chineseStr;
    }
    const formatter=(cellValue,prefix,suffix)=> {
        return prefix+cellValue+suffix//XEUtils.commafy(XEUtils.toNumber(cellValue), { digits: 2 })+suffix
    }
    const formatterNum=(cellValue,decimals)=> {
        return  Number(cellValue).toLocaleString('en-US', {maximumFractionDigits: decimals}) //XEUtils.commafy(XEUtils.toNumber(cellValue), { digits: 2 })+suffix
    }
    const percentage=(cellValue)=> {
        return Math.round(cellValue*100) +'%'
    }
    const permillage=(cellValue)=> {
        return Math.round(cellValue*1000) +'‰'
    }
</script>