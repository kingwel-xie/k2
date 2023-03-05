<template>
  <ScrollContainer>
    <div ref="wrapperRef" class="xl:flex">
      <Tabs
        tab-position="left"
        :tabBarStyle="tabBarStyle"
        class="xl:w-3/4 w-full"
        :class="prefixCls"
      >
        <template v-for="item in settingList" :key="item.key">
          <TabPane :tab="item.name">
            <component :is="item.component" />
          </TabPane>
        </template>
      </Tabs>
    </div>
  </ScrollContainer>
</template>

<script lang="ts">
  import { defineComponent } from 'vue';
  import { Tabs, TabPane } from 'ant-design-vue';
  import { ScrollContainer } from '/@/components/Container/index';
  import Inbox from './Inbox.vue';
  import Outbox from './Outbox.vue';

  export default defineComponent({
    components: {
      ScrollContainer,
      Tabs,
      TabPane,
      Inbox,
      Outbox,
    },

    setup() {
      // tab的list
      const settingList = [
        {
          key: '1',
          name: '新消息',
          component: 'Inbox',
        },
        {
          key: '2',
          name: '已发送',
          component: 'Outbox',
        },
        // {
        //   key: '3',
        //   name: '草稿',
        //   component: 'Draft',
        // },
      ];

      return {
        prefixCls: 'account-setting',
        settingList,
        tabBarStyle: {
          width: '180px',
        },
      };
    },
  });
</script>
<style lang="less">
  .account-setting {
    margin: 12px;
    background-color: @component-background;

    .base-title {
      padding-left: 0;
    }

    .ant-tabs-tab-active {
      background-color: @item-active-bg;
    }
  }
</style>
