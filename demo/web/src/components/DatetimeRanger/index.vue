<template>
  <div class="datetime-ranger">
    <el-date-picker
      v-model="value0"
      type="datetimerange"
      :picker-options="datetimePickerOptions()"
      range-separator="至"
      start-placeholder="开始日期"
      end-placeholder="结束日期"
      align="right"
      value-format="yyyy-MM-dd HH:mm:ss"
    />
  </div>
</template>

<script>
export default {
  name: 'DatetimeRanger',
  props: {
    value: {
      type: Array,
      default: () => {}
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
            start.setTime(start.getTime() - 3600 * 1000 * 24 * 180)
            picker.$emit('pick', [start, end])
          }
        }, {
          text: '本月',
          onClick(picker) {
            const end = new Date()
            const start = new Date()
            start.setDate(1)
            start.setHours(0, 0, 0)
            picker.$emit('pick', [start, end])
          }
        }, {
          text: '本季度',
          onClick(picker) {
            const end = new Date()
            const start = new Date()
            const month = start.getMonth()
            start.setMonth(month / 3 * 3, 1)
            start.setHours(0, 0, 0)
            picker.$emit('pick', [start, end])
          }
        }, {
          text: '本年',
          onClick(picker) {
            const end = new Date()
            const start = new Date()
            start.setMonth(0, 1)
            start.setHours(0, 0, 0)
            picker.$emit('pick', [start, end])
          }
        }]
      }
    }
  }
}
</script>
