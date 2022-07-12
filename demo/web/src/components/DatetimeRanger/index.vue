<template>
  <div class="datetime-ranger">
    <el-date-picker
      v-model="value0"
      v-bind="$attrs"
      type="daterange"
      :picker-options="datetimePickerOptions()"
      range-separator="至"
      start-placeholder="开始日期"
      end-placeholder="结束日期"
      align="right"
    />
  </div>
</template>

<script>
export default {
  name: 'DatetimeRanger',
  props: {
    value: {
      type: Array,
      default: () => []
    }
  },
  data() {
    return {
      value0: this.value
    }
  },
  watch: {
    value(val) {
      // console.log('yyy', val)
      this.value0 = val
    },
    value0(val) {
      // console.log('zzz', val)
      if (val[0]) {
        val[0].setHours(0, 0, 0)
      }
      if (val[1]) {
        val[1].setHours(23, 59, 59, 999)
      }
      this.$emit('input', val)
    }
  },
  methods: {
    datetimePickerOptions() {
      return {
        shortcuts: [{
          text: '最近一周',
          onClick(picker) {
            const end = new Date()
            const start = new Date()
            start.setTime(start.getTime() - 3600 * 1000 * 24 * 7)
            picker.$emit('pick', [start, end])
          }
        }, {
          text: '最近一个月',
          onClick(picker) {
            const end = new Date()
            const start = new Date()
            start.setTime(start.getTime() - 3600 * 1000 * 24 * 30)
            picker.$emit('pick', [start, end])
          }
        }, {
          text: '最近三个月',
          onClick(picker) {
            const end = new Date()
            const start = new Date()
            start.setTime(start.getTime() - 3600 * 1000 * 24 * 90)
            picker.$emit('pick', [start, end])
          }
        }, {
          text: '最近半年',
          onClick(picker) {
            const end = new Date()
            const start = new Date()
            start.setTime(start.getTime() - 3600 * 1000 * 24 * 183)
            picker.$emit('pick', [start, end])
          }
        }, {
          text: '本月',
          onClick(picker) {
            const end = new Date()
            const start = new Date()
            start.setDate(1)
            picker.$emit('pick', [start, end])
          }
        }, {
          text: '本季度',
          onClick(picker) {
            const end = new Date()
            const start = new Date()
            const month = start.getMonth()
            start.setMonth(month / 3 * 3, 1)
            picker.$emit('pick', [start, end])
          }
        }, {
          text: '本年',
          onClick(picker) {
            const end = new Date()
            const start = new Date()
            start.setMonth(0, 1)
            picker.$emit('pick', [start, end])
          }
        }]
      }
    }
  }
}
</script>
