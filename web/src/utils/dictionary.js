import { useDictionaryStore } from '@/pinia/modules/dictionary'
// Get dictionary method usage example getDict('sex').then(res) or async function const res = await getDict('sex')
export const getDict = async(type) => {
  const dictionaryStore = useDictionaryStore()
  await dictionaryStore.getDictionary(type)
  return dictionaryStore.dictionaryMap[type]
}

//Dictionary text display method
export const showDictLabel = (dict, code) => {
  if (!dict) {
    return ''
  }
  const dictMap = {}
  dict.forEach(item => {
    dictMap[item.value] = item.label
  })
  return dictMap[code]
}
