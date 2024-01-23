import {ref} from 'vue'
import {defineStore} from 'pinia'

export const useAppStore = defineStore('AppStore', {
        state: () => {
            return {
                当前拖拽组件数据:ref(null)
            }
        },
        actions:{

        }
    }
)