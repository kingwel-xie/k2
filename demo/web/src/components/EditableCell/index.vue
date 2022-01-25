<template>
  <div class="edit-cell" @click="onFieldClick">
    <el-tooltip
      v-if="!editMode && !showInput"
      :placement="toolTipPlacement"
      :open-delay="toolTipDelay"
      :content="toolTipContent"
      :disabled="!canEdit"
    >
      <div
        tabindex="0"
        class="cell-content"
        :class="computedClasses"
        @keyup.enter="onFieldClick"
      >
        <slot name="content" />
      </div>
    </el-tooltip>
    <component
      :is="editableComponent"
      v-if="editMode || showInput"
      ref="input"
      v-model="model"
      v-bind="$attrs"
      @focus="onFieldClick"
      @keyup.enter.native="onInputExit"
      v-on="listeners"
    >
      <slot name="edit-component-slot" />
    </component>
  </div>
</template>
<script>
export default {
  name: 'EditableCell',
  inheritAttrs: false,
  props: {
    value: {
      type: [String, Number, Object],
      default: ''
    },
    initValue: {
      type: Number,
      default: 0
    },
    toolTipContent: {
      type: String,
      default: '点击进入编辑'
    },
    toolTipDelay: {
      type: Number,
      default: 500
    },
    toolTipPlacement: {
      type: String,
      default: 'right-start'
    },
    showInput: {
      type: Boolean,
      default: false
    },
    editableComponent: {
      type: String,
      default: 'el-input'
    },
    closeEvent: {
      type: String,
      default: 'blur'
    },
    canEdit: {
      type: Boolean,
      default: false
    }
  },
  data() {
    return {
      editMode: false
    }
  },
  computed: {
    computedClasses() {
      return {
        'edit-enabled-cell': this.canEdit,
        'edit-enabled-cell-less': this.initValue !== 0 && this.initValue > this.value,
        'edit-enabled-cell-greater': this.initValue !== 0 && this.initValue < this.value
      }
    },
    model: {
      get() {
        return this.value
      },
      set(val) {
        this.$emit('input', val)
      }
    },
    listeners() {
      return {
        [this.closeEvent]: this.onInputExit,
        ...this.$listeners
      }
    }
  },
  methods: {
    onFieldClick() {
      if (this.canEdit) {
        this.editMode = true
        this.$nextTick(() => {
          const inputRef = this.$refs.input
          if (inputRef && inputRef.focus) {
            inputRef.focus()
          }
        })
      }
    },
    onInputExit() {
      this.editMode = false
    },
    onInputChange(val) {
      this.$emit('input', val)
    }
  }
}
</script>

<style>
.cell-content {
  min-height: 30px;
  padding-left: 5px;
  padding-top: 5px;
  border: 1px solid transparent;
}
.edit-enabled-cell {
  border: 1px dashed #409eff;
}
.edit-enabled-cell-less {
  border: 1px dashed #00ff00;
}
.edit-enabled-cell-greater {
  border: 1px dashed #ff0000;
}
</style>

<style scoped>

.class /deep/  .a input::-webkit-outer-spin-button,
.class /deep/  .a input::-webkit-inner-spin-button {
  -webkit-appearance: none;
}
.class /deep/  .a input[type="number"]{
  -moz-appearance: textfield;
}
</style>
