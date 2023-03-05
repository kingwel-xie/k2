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
  </div>
</template>
<script lang="ts" setup>
  import { computed } from 'vue';
  import { Badge, Popover } from 'ant-design-vue';
  import { BellOutlined } from '@ant-design/icons-vue';
  import { useDesign } from '/@/hooks/web/useDesign';
  import { useGo } from '/@/hooks/web/usePage';
  import { useUserStoreWithOut } from '/@/store/modules/user';

  const { prefixCls } = useDesign('header-notify');
  const go = useGo();
  const numberStyle = {};

  const count = computed(() => {
    const userStore = useUserStoreWithOut();
    return userStore.getUnreadMessage.numMessages + userStore.getUnreadMessage.numNotices;
  });

  function handleGoMessage() {
    go('/dashboard/message');
  }
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
