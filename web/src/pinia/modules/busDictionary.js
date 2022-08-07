import { findBusDictionary } from '@/api/busDictionary'

import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useBusDictionaryStore = defineStore('busDictionary', () => {
  const dictionaryMap = ref({})

  const setDictionaryMap = (dictionaryRes) => {
    // 合并多个对象
    dictionaryMap.value = { ...dictionaryMap.value, ...dictionaryRes }
  }

  const getDictionary = async(type) => {
    if (dictionaryMap.value[type] && dictionaryMap.value[type].length) {
      return dictionaryMap.value[type]
    } else {
      const res = await findBusDictionary({ "DictName":type })
      if (res.code === 0) {
        const dictionaryRes = {}
        const dict = []
        res.data.rebusDict && res.data.rebusDict.forEach(item => {
          dict.push({
            label: item.name,
            value: item.ID
          })
        })
        dictionaryRes[type] = dict
        setDictionaryMap(dictionaryRes)
        return dictionaryMap.value[type]
      }
    }
  }

  return {
    dictionaryMap,
    setDictionaryMap,
    getDictionary
  }
})
