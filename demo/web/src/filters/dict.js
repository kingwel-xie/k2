import store from '@/store'
import { LABEL_KEY } from '@/store/modules/dictionary'

export default {
  install(Vue) {
    Vue.filter('dict', (val, kind) => {
      return val ? (store.getters.dictRegistry(kind)[val] || {})[LABEL_KEY] : val
    })
  }
}
