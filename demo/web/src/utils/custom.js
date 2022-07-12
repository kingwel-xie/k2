
// float格式化
export function toFloat(val, d) {
  return val !== undefined ? val.toFixed(d || 2) : '-'
}

// 日期格式化
export function parseBoolean(val) {
  return val ? '是' : '否'
}

export function parseTimeCSharp(time) {
  return new Date(parseInt(time.substr(6, 13))).toLocaleString()
}

// 表单重置
export function resetForm(refName) {
  if (this.$refs[refName]) {
    this.$refs[refName].resetFields()
  }
}

export function dateRangeByDays(days) {
  const end = new Date()
  const start = new Date()
  start.setTime(start.getTime() - 3600 * 1000 * 24 * days)
  start.setHours(0, 0, 0)
  end.setHours(23, 59, 59, 999)
  return [start, end]
}

// 添加日期范围
export function addDateRange(params, dateRange) {
  var search = params
  search.beginTime = ''
  search.endTime = ''
  if (dateRange != null && dateRange !== '') {
    search.beginTime = this.dateRange[0]
    search.endTime = this.dateRange[1]
  }
  return search
}

// 字符串格式化(%s )
export function sprintf(str) {
  var args = arguments; var flag = true; var i = 1
  str = str.replace(/%s/g, function() {
    var arg = args[i++]
    if (typeof arg === 'undefined') {
      flag = false
      return ''
    }
    return arg
  })
  return flag ? str : ''
}

// 转换字符串，undefined,null等转化为""
export function praseStrEmpty(str) {
  if (!str || str === 'undefined' || str === 'null') {
    return ''
  }
  return str
}

// a trick to generate UUID
export function genUuid() {
  const temp_url = URL.createObjectURL(new Blob())
  const uuid = temp_url.toString() // blob:https://xxx.com/b250d159-e1b6-4a87-9002-885d90033be3
  URL.revokeObjectURL(temp_url)
  return uuid.substr(uuid.lastIndexOf('/') + 1)
}

// split a string with [, or \n] to a string array
// used to split delimited serial numbers
export function stringToArray(str, separator = /[,;\n]/) {
  return str ? str.split(separator).filter(x => x !== '').map(x => x.trim()) : undefined
}

export function validateRange(dateRange, maxDays) {
  if (!dateRange || dateRange.length !== 2) {
    return '必须指定查询时间间隔'
  }

  const diff = Date.parse(dateRange[1]) - Date.parse(dateRange[0])
  const days = Math.floor(diff / (24 * 3600 * 1000))
  // console.log('days', days)
  if (days > maxDays) {
    return '查询时间天数间隔为 ' + days + ', 最大允许为 ' + maxDays
  }
}

// validator, number > 0
export function validateGt0(rule, value, callback) {
  if (Number(value) > 0) {
    callback()
  } else {
    callback(new Error('请输入大于 0 的数'))
  }
}
