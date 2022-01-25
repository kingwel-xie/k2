<template>
  <div class="slider_small">
    <svg-icon class-name="size-icon" icon-class="size" />
    <el-slider
      v-model="value"
      :min="0"
      :max="3"
      :step="1"
      :marks="marks"
      @change="handleSetSize"
    />
  </div>
</template>

<script>
export default {
  name: 'SizeSelect',
  data() {
    return {
      value: 0,
      marks: {
        0: '最小',
        1: {
          style: {
            color: '#1989FA'
          },
          label: this.$createElement('strong', '较小')
        },
        2: '较大',
        3: '最大'
      },
      sizeOptions: [
        { label: 'Mini', value: 'mini' },
        { label: 'Small', value: 'small' },
        { label: 'Medium', value: 'medium' },
        { label: 'Default', value: 'default' }
      ]
    }
  },
  created() {
    // console.log(this.$store.getters.size)
    this.value = this.sizeOptions.findIndex(x => x.value === this.$store.getters.size)
  },
  methods: {
    handleSetSize(val) {
      const size = this.sizeOptions[val].value
      // console.log('setsize', size)
      this.$ELEMENT.size = size
      this.$store.dispatch('app/setSize', size)
      this.refreshView()
      this.$message({
        message: '切换成功',
        type: 'success'
      })
    },
    refreshView() {
      // In order to make the cached page re-rendered
      this.$store.dispatch('tagsView/delAllCachedViews', this.$route)

      const { fullPath } = this.$route

      this.$nextTick(() => {
        this.$router.replace({
          path: '/redirect' + fullPath
        })
      })
    }
  }

}
</script>

<style scoped>

.slider_small {
  padding: 15px;
  font-size: 16px;
}

/deep/ .el-slider__marks-text  {
  font-size: 12px;
}

</style>
