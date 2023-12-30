<template>
  <div class="cursor-pointer underline" @click="showMessage">点击查看</div>
  <BasicModal
    @register="registerModal"
    :title="title"
    :width="800"
    :cancelText="t('common.closeText')"
    :showOkBtn="false"
    :useWrapper="false"
  >
    <div v-html="value"></div>
  </BasicModal>
</template>
<script lang="ts" setup>
  import { BasicModal, useModalInner } from '/@/components/Modal';
  import { useI18n } from '/@/hooks/web/useI18n';

  const { t } = useI18n();

  defineProps({
    title: {
      type: String as PropType<string>,
      default: '',
    },
    value: {
      type: String as PropType<string>,
      default: '',
    },
  });

  const [registerModal, { setModalProps }] = useModalInner();

  function showMessage() {
    // open modal
    setModalProps({
      visible: true,
    });
  }
</script>

<style lang="less" scoped>
  .ellipsis-container-multiline {
    display: -webkit-box;
    -webkit-box-orient: vertical;
    overflow: hidden;
    /* 设置行数限制，例如： */
    -webkit-line-clamp: 3;
  }
  .ellipsis-container-multiline.expanded {
    -webkit-line-clamp: unset;
  }
</style>
