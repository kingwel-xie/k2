<template>
  <div class="cell">
    <div class="cell-inner" :class="border ? 'cell-border' : ''">
      <div class="cell-item">
        <div class="cell-item-label">
          <span v-if="label">
            {{ label }}
          </span>
          <span v-else>
            <slot name="label"></slot>
          </span>
          :
        </div>
        <div class="cell-item-value">
          <span v-if="value">
            {{ value }}
          </span>
          <span v-else>
            <slot name="value"></slot>
          </span>
        </div>
      </div>
      <div class="cell-item mt-2">
        <span v-if="extra">{{ extra }}</span>
        <span v-else><slot name="extra"></slot></span>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
  import { propTypes } from '/@/utils/propTypes';

  defineProps({
    label: propTypes.string.def(''),
    value: propTypes.string.def(''),
    extra: propTypes.string.def(''),
    border: propTypes.bool.def(true),
  });
</script>

<style lang="less" scoped>
  .cell-inner {
    box-sizing: border-box;
    width: 100%;
    overflow: hidden;
    color: #323233;
    font-size: 14px;
    line-height: 24px;
    background-color: #fff;
    padding: 8px 0;

    .cell-item {
      display: flex;
      justify-content: space-between;
      align-items: center;
    }
    .cell-item-label {
      span {
        color: #323233;
        font-size: 14px;
        line-height: 24px;
      }
    }
    .cell-item-value {
      color: #969799;
    }
  }

  .cell-border {
    border-bottom: 1px solid #e6ebf5;
  }
</style>
