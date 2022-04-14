import Vue from 'vue'

import Cookies from 'js-cookie'

import 'normalize.css/normalize.css' // a modern alternative to CSS resets

import ElementUI from 'element-ui'
import './styles/element-variables.scss'

import '@/styles/index.scss' // global css
import '@/styles/admin.scss'

import App from './App'
import store from './store'
import router from './router'
import permission from './directive/permission'
import clipboard from '@/directive/clipboard'
import ffiov from '@/directive/focus-first-invalid-on-validate'
import formatter from '@/directive/formatter'
import copy from '@/directive/copy'
import dict from '@/filters/dict'

import { getItems } from '@/api/table'
import { getConfigKey } from '@/api/admin/sys-config'
import {
  parseTime,
  resetForm,
  addDateRange,
  parseBoolean
} from '@/utils/custom'

import './icons' // icon
import './permission' // permission control
import './utils/error-log' // error log

import Viser from 'viser-vue'
Vue.use(Viser)

import * as filters from './filters' // global filters

import Pagination from '@/components/Pagination'
import BasicLayout from '@/layout/BasicLayout'
import K2Dialog from '@/components/K2Dialog'
import K2Descriptions from '@/components/K2Descriptions'
import DatetimeRanger from '@/components/DatetimeRanger'
import DictSelect from '@/components/DictSelect'
import DictRadioGroup from '@/components/DictRadioGroup'

// particle effect, see login/index.vue
import VueParticles from 'vue-particles'
Vue.use(VueParticles)

import '@/utils/dialog'

// 全局方法挂载
Vue.prototype.parseBoolean = parseBoolean
Vue.prototype.getItems = getItems
Vue.prototype.getConfigKey = getConfigKey
Vue.prototype.parseTime = parseTime
Vue.prototype.resetForm = resetForm
Vue.prototype.addDateRange = addDateRange

// 全局组件挂载
Vue.component('Pagination', Pagination)
Vue.component('BasicLayout', BasicLayout)
Vue.component('K2Dialog', K2Dialog)
Vue.component('K2Descriptions', K2Descriptions)
Vue.component('DatetimeRanger', DatetimeRanger)
Vue.component('DictSelect', DictSelect)
Vue.component('DictRadioGroup', DictRadioGroup)

Vue.prototype.msgSuccess = function(msg) {
  this.$message({ showClose: true, message: msg, type: 'success' })
}

Vue.prototype.msgError = function(msg) {
  this.$message({ showClose: true, message: msg, type: 'error' })
}

Vue.prototype.msgInfo = function(msg) {
  this.$message.info(msg)
}

Vue.use(permission)
Vue.use(clipboard)
Vue.use(ffiov)
Vue.use(formatter)
Vue.use(copy)
Vue.use(dict)

Vue.use(ElementUI, {
  size: Cookies.get('size') || 'small' // set element-ui default size
})

import 'remixicon/fonts/remixicon.css'

console.info(`欢迎使用 K2`)

// register global utility filters
Object.keys(filters).forEach(key => {
  Vue.filter(key, filters[key])
})

Vue.config.productionTip = false
ElementUI.Dialog.props.closeOnClickModal.default = false
ElementUI.Drawer.props.wrapperClosable.default = false
ElementUI.InputNumber.props.min.default = 0
ElementUI.InputNumber.props.precision.default = 2
ElementUI.InputNumber.props.controlsPosition.default = 'right'

new Vue({
  el: '#app',
  router,
  store,
  render: h => h(App)
})
