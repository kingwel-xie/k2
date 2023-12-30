<template>
  <div :class="prefixCls">
    <Popover title="" trigger="hover">
      <Badge :count="count" dot :numberStyle="numberStyle" @click="handleGoMessage">
        <BellOutlined />
      </Badge>
      <template #content>
        <span class="text-md text-secondary"> 您有 {{ count }} 条新消息。 </span>
      </template>
    </Popover>
    <NotifyModal @register="registerModal" />
  </div>
</template>
<script lang="ts" setup>
  import { computed, onMounted, unref } from 'vue';
  import { Badge, Popover } from 'ant-design-vue';
  import { BellOutlined } from '@ant-design/icons-vue';
  import { useDesign } from '/@/hooks/web/useDesign';
  import { useGo } from '/@/hooks/web/usePage';
  import { useUserStoreWithOut } from '/@/store/modules/user';
  import { useModal } from '/@/components/Modal';
  import NotifyModal from '/@/layouts/default/header/components/notify/NotifyModal.vue';

  onMounted(() => {
    const noticeList = userStore.getUnreadMessage.noticeList?.filter((x) => x.importance === 'Y');
    if (noticeList?.length > 0) {
      openModal(true, noticeList);
    }
  });

  const { prefixCls } = useDesign('header-notify');
  const go = useGo();
  const numberStyle = {};
  const userStore = useUserStoreWithOut();

  const count = computed(() => {
    return userStore.getUnreadMessage.numMessages + userStore.getUnreadMessage.numNotices;
  });

  function handleGoMessage() {
    go('/dashboard/message');
  }

  const [registerModal, { openModal }] = useModal();
</script>
<style lang="less">
  @prefix-cls: ~'@{namespace}-header-notify';

  .@{prefix-cls} {
    padding-top: 2px;

    &__overlay {
      width: 320px;
    }

    .ant-tabs-content {
      width: 300px;
    }

    .ant-badge {
      font-size: 18px;

      .ant-badge-multiple-words {
        padding: 0 4px;
      }

      svg {
        width: 0.9em;
      }
    }
  }
</style>
