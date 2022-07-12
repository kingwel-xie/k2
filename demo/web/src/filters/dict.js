import store from '@/store'
import { LABEL_KEY, VALUE_KEY } from '@/store/modules/dictionary'

export default {
  install(Vue) {
    Vue.filter('dict', (val, kind) => {
      return val ? (store.getters.dictRegistry(kind)[val] || {})[LABEL_KEY] : val
    })
    Vue.filter('rdict', (val, kind) => {
      const list = store.getters.listRegistry(kind) || []
      const n = list.find(x => x[LABEL_KEY] === val)
      return n && n[VALUE_KEY]
    })
  }
}
