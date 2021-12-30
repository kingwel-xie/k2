<template>
  <el-card class="box-card" body-style="padding: 0 20px !important;">
    <div slot="header" class="clearfix">
      <slot v-if="$slots.header || header" name="header"><span>{{ header }}</span></slot>
      <el-button
        style="float: right; padding: 3px 0"
        type="text"
        :icon="collapse ? 'el-icon-arrow-right' : 'el-icon-arrow-down'"
        @click="handleCollapseExpand"
      />
    </div>
    <div v-show="!collapse" class="card-content" :style="bodyStyle">
      <slot />
    </div>
  </el-card>
</template>

<script>
export default {
  name: 'CollapseCard',
  props: {
    header: {},
    bodyStyle: {}
  },
  data() {
    return {
      collapse: false
    }
  },
  methods: {
    handleCollapseExpand() {
      this.collapse = !this.collapse
    }
  }
}
</script>

<style scoped>
.card-content {
  padding: 20px 0 10px 0;
}
</style>
