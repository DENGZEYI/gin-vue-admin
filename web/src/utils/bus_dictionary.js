// 用于获取应用层小组、商品、供应商字典
import { useBusDictionaryStore } from '@/pinia/modules/busDictionary'

export const getBusDict = async(type) => {
  const dictionaryStore = useBusDictionaryStore()
  await dictionaryStore.getDictionary(type)
  return dictionaryStore.dictionaryMap[type]
}
