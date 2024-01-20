<template>
   <div :style="{ height: el.height,width:el.width}" v-html="data"></div>
</template>

<script setup>
    import {ref,defineProps,onMounted,nextTick,getCurrentInstance,inject} from 'vue'
    import {ApiGetReportData} from '@/utils/api.js'
    const {proxy}=getCurrentInstance()

    const props=defineProps({
        el:Object,
        edit:Boolean
    })
    const pid=inject("tid")
    const rid=inject("rid")

    const data= ref('')
    const generateReport=()=>{
        ApiGetReportData(props.el.data.temp,props.el.name,pid,rid,JSON.stringify(props.el.data)).then(res=>{
            data.value=res.data
            //console.log(res)
            //res.data.xAxis.axisLabel.formatter=formatter
            //nextTick(() => {
            //    Chart.setOption(res.data)
            //    //echarts图表自适应实现
            //    Chart.resize()
            //    window.addEventListener('resize', function () {
            //    Chart.resize()
            //    })
            //})
        })
    }
   
    onMounted(()=>{
        generateReport()
    })
</script>