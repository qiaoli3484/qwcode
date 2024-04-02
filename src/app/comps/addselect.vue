<template>
    <el-select v-model="data.default" placeholder="选择默认值" style="width: 240px">
      <el-option
        v-for="(item,index) in data.params"
        :key="item.value"
        :label="item.label"
        :value="item.value"
      >
            <span style="float: left">{{ item.label }}</span>
            <span
            style="
            float: right;
            color: var(--el-text-color-secondary);
            font-size: 13px;
            "
        >{{ item.value }}</span>
        <el-button type="danger" text icon="CircleClose" @click.stop="deldata(index)"></el-button>
      </el-option>
      <template #footer>
        <el-button v-if="!isAdding" text bg size="small" @click="onAddOption">
          添加参数
        </el-button>
        <template v-else>
          <el-input
            v-model="optionName"
            class="option-input"
            placeholder="标签"
            size="small"
          />
          <el-input
            v-model="optionval"
            class="option-input"
            placeholder="参数"
            size="small"
          />
          <el-button type="primary" size="small" @click="onConfirm">
            确认
          </el-button>
          <el-button size="small" @click="clear">取消</el-button>
        </template>
      </template>
    </el-select>
  </template>
  
  <script  setup>
  import { ref,computed,defineProps,defineEmits } from 'vue'

    const prop=defineProps({
        modelValue:Object
    })
   const emit = defineEmits(['update:modelValue'])

   const data=computed({
        get: () => {
            if (typeof prop.modelValue=='undefined'){
                return new Object
            }
            return prop.modelValue
        },
        set: (val) => {
            emit('update:modelValue',val)
        }
    })

  const isAdding = ref(false)
  const optionName = ref('')
  const optionval = ref('')

  
  const onAddOption = () => {
    isAdding.value = true
  }
  
  const onConfirm = () => {
    if (optionName.value) {
        data.value.params.push({
        label: optionName.value,
        value: optionval.value,
      })
      clear()
    }
  }
  
  const clear = () => {
    optionName.value = ''
    optionval.value= ''
    isAdding.value = false
  }

  const deldata=(index)=>{
    data.value.params.splice(index,1)
  }
  </script>
  
  <style lang="scss" scoped>
  .option-input {
    width: 100%;
    margin-bottom: 8px;
  }
  </style>