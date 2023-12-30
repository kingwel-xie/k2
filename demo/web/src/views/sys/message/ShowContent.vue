<template>
  <div
    ref="text"
    class="ellipsis-container-multiline"
    v-bind:class="{ expanded: showMore }"
    v-html="value"
  ></div>
  <a-button
    v-if="shouldShowButton"
    type="link"
    :preIcon="showMore ? 'ant-design:caret-down-outlined' : 'ant-design:caret-right-outlined'"
    @click="showMore = !showMore"
    >{{ showMore ? '隐藏' : '全部...' }}
  </a-button>
</template>
<script lang="ts" setup>
  import { computed, ref, unref } from 'vue';

  defineProps({
    value: {
      type: String as PropType<string>,
      default: '',
    },
  });

  const text = ref<HTMLElement>();
  const showMore = ref(false);
  const shouldShowButton = computed(() => {
    const textElem = unref(text);
    if (!textElem) return false;
    return textElem.scrollHeight > textElem.clientHeight;
  });
</script>

<style lang="less" scoped>
  .ellipsis-container-multiline {
    display: -webkit-box;
    -webkit-box-orient: vertical;
    overflow: hidden;
    /* 设置行数限制，例如： */
    -webkit-line-clamp: 3;
    /* 设置宽度限制，例如： */
    //max-width: 800px;
  }
  .ellipsis-container-multiline.expanded {
    -webkit-line-clamp: unset;
  }
</style>
