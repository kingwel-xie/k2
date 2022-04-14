<template>
  <div>
    <el-descriptions
      :title="title"
      v-bind="$attrs"
    >
      <template #title>
        <span v-if="title">
          <span style="display:inline-block;background-color: #81a3e5;width:2px;height:16px;margin-right:5px; float: left;margin-top:2px" />
          {{ title }}
        </span>
      </template>

      <slot></slot>

    </el-descriptions>
  </div>
</template>

<script>
export default {
  name: 'K2Descriptions',
  props: {
    title: {
      type: String,
      default: undefined
    }
  },
  data() {
    return {
    }
  },
  methods: {
  }
}
</script>

<style lang="scss" scoped>

</style>
