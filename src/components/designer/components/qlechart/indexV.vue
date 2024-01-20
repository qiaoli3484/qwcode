<template>
   <div ref="myechar" :style="{ height: el.height,width:el.width}"></div>
</template>

<script setup>
    import {ref,defineProps,onMounted,nextTick,getCurrentInstance,inject} from 'vue'
    import * as echarts from 'echarts'
    import {ApiGetChartData} from '@/utils/api.js'
    const {proxy}=getCurrentInstance()

    
    const props=defineProps({
        el:Object,
        edit:Boolean
    })
    const myechar=ref(null)
    let Chart=null

    const pid=inject("tid")
    const rid=inject("rid")

    const generateChart=()=>{
        
        ApiGetChartData(props.el.data.temp,props.el.name,pid.value,rid.value,JSON.stringify(props.el.data)).then(res=>{
            //console.log(res)
            //res.data.xAxis.axisLabel.formatter=formatter
            nextTick(() => {
                Chart.setOption(res.data)
                //echarts图表自适应实现
                Chart.resize()
                window.addEventListener('resize', function () {
                Chart.resize()
                })
            })

            //let myChart = echarts.init(document.getElementById('wItVyUbIkZOc'), "white")
            //nextTick(()=>{
           //     myChart.setOption(res.data);
           // })
            //myChart.on('click', function (params) {
            //        window.open('https://www.baidu.com/s?wd=' + encodeURIComponent(params.name));
            //    });
        })
    }
   
    onMounted(()=>{
        Chart = echarts.init(myechar.value)
        generateChart()
    })
</script>