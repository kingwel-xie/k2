import moment from 'moment'
import store from '@/store'
import { LABEL_KEY } from '@/store/modules/dictionary'
import { formatTime, tryParseJson } from '@/utils'

export default {
  install(Vue) {
    // console.log('ISO_8601', moment.ISO_8601)
    // console.log('moment-> ', moment(new Date('0001-01-01T00:00:00Z')).format('YYYY-MM-DD HH:mm Z'))
    /**
     * This directive is used to automatically provide code-to-name formatting functions for dictionary table columns
     *
     * The usage as follows:
     * a.) for bool value
     *  v-fmt.bool OR v-fmt.bool="['Yes', 'No']" default trueOrFalse is ['是', '否']
     *
     * b.) for float value
     *  v-fmt.float OR v-fmt.float="3" default fraction digits '2'
     *
     * c.) for date and/or time value
     *  v-fmt.date OR v-fmt.date="'YYYY.MM.DD'" default format 'YYYY-MM-DD'
     *  v-fmt.time OR v-fmt.time="'HH:mm'" default format 'HH:mm:ss'
     *  v-fmt.datetime OR v-fmt.datetime="'YYYY.MM.DD HH:mm'" default format moment.ISO_8601
     *
     *  Tips: 'format' please refer to https://momentjs.com/
     *
     * d.) for dict table
     *  v-fmt.dict="'TbxChannel'" OR v-fmt.dict="'tbx_container_type'"
     *
     *  *Note*: The dictionary library needs to be prepared in advance
     *
     * *Note*: v-fmt only works for <el-table-column>
     */
    Vue.directive('fmt', {
      bind: function(el, binding, vnode) {
        if (vnode.componentOptions.tag !== 'el-table-column') {
          throw new Error('v-dict only works for <el-table-column>')
        }
        const { modifiers: { debug, bool, float, date, time, datetime, fulltime, dict }, value } = binding
        let formatter = null
        if (bool) {
          // check value is array and length === 2
          const trueOrFalse = value || ['是', '否']
          formatter = (row, column, val) => {
            return trueOrFalse[val ? 0 : 1]
          }
        } else if (float) {
          // value is a number, represents the number of digits after the decimal point
          const fractionDigits = value || 2
          formatter = (row, column, val) => {
            return val ? val.toFixed(fractionDigits) : ''
          }
        } else if (date || time || fulltime) {
          // moment.ISO_8601
          const format = value || (date ? 'YYYY-MM-DD' : time ? 'HH:mm:ss' : 'YYYY-MM-DD HH:mm:ss')
          formatter = (row, column, val) => {
            if (!val || val.indexOf('01-01-01') > -1) {
              return '-'
            }
            return moment(new Date(val)).format(format)
          }
        } else if (datetime) {
          formatter = (row, column, val) => {
            if (!val || val.indexOf('01-01-01') > -1) {
              return '-'
            }
            return formatTime(new Date(val), '{y}-{m}-{d} {h}:{i}')
          }
        } else if (dict) {
          const dicts = store.getters.dictRegistry(value)
          if (debug) {
            console.log(`found dicts for ${value}:`, dicts)
          }
          formatter = (row, column, val) => {
            if (debug) {
              console.log(`formatting ${value} '${val}' with:`, dicts)
            }
            return val ? (dicts[val] || {})[LABEL_KEY] : ''
          }
        }
        vnode.componentInstance.columnConfig.formatter = formatter
      },
      unbind: function(el, binding, vnode) {
        vnode.componentInstance.columnConfig.formatter = undefined
      }
    })
  }
}

