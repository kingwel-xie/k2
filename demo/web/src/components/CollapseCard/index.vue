<template>
  <el-card class="box-card" body-style="padding: 0 20px !important;">
    <div slot="header" class="clearfix header" role="button" @click="handleCollapseExpand">
      <slot v-if="$slots.header || header" name="header">
        <el-icon :name="collapse ? 'circle-plus-outline' : 'remove-outline'" style="padding-right: 4px" />
        <span>{{ header }}</span>
      </slot>
      <el-icon style="margin: 0 8px 0 auto" :name="collapse ? 'arrow-right' : 'arrow-down'" />
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
    header: {
      type: String,
      default: ''
    },
    collapseByDefault: {
      type: [Boolean, String],
      default: false
    },
    bodyStyle: {
      type: String,
      default: ''
    }
  },
  data() {
    return {
      collapse: this.collapseByDefault === '' ? true : this.collapseByDefault
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
  padding: 8px 0 6px 0;
}

.header {
  display: flex;
  align-items: center;
  cursor: pointer;
}

</style>
