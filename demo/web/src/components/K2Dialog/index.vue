<template>
  <div>
    <el-dialog
      :close-on-click-modal="false"
      :title="title"
      :width="width"
      :visible.sync="subVisible"
      :fullscreen="fullscreen"
      @close="handleClosed"
      v-bind="$attrs"
    >
      <template #title>
        <span class="el-dialog__title">
          <span style="display:inline-block;background-color: #3478f5;width:3px;height:20px;margin-right:5px; float: left;margin-top:2px" />
          {{ title }}
        </span>
        <div v-if="showFullScreen" class="__dialog_fullscreen_button" @click="fullscreen = !fullscreen">
          <i class="el-icon-full-screen" />
        </div>
      </template>

      <slot></slot>

      <template #footer>
        <slot name="footer"></slot>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'K2Dialog',
  props: {
    visible: {
      type: Boolean,
      default: false
    },
    showFullScreen: {
      type: Boolean,
      default: false
    },
    width: {
      type: String,
      default: ''
    },
    title: {
      type: String,
      default: ''
    }
  },
  data() {
    return {
      subVisible: this.visible,
      fullscreen: false
    }
  },
  watch: {
    visible(val) {
      this.subVisible = val
    },
    subVisible(val) {
      this.$emit('update:visible', val)
    }
  },
  methods: {
    handleClosed() {
      this.$emit('closed')
    }
  }
}
</script>

<style lang="scss" scoped>

.__dialog_fullscreen_button i {
  color: #909399;
  position: absolute;
  top: 20px;
  right: 40px;
  padding: 0;
  background: transparent;
  border: none;
  outline: none;
  cursor: pointer;
  font-size: 16px
}
.el-icon-full-screen{
  cursor: pointer;
}
.el-icon-full-screen:before {
  content: "\e719";
}

</style>
