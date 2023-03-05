<template>
  <div :class="prefixCls">
    <Popover title="" trigger="click" :overlayClassName="`${prefixCls}__overlay`">
      <Badge :count="count" dot :numberStyle="numberStyle">
        <BellOutlined />
      </Badge>
      <template #content>
        <Tabs>
          <template v-for="item in listData" :key="item.key">
            <TabPane>
              <template #tab>
                {{ item.name }}
                <span v-if="item.list.length !== 0">({{ item.list.length }})</span>
              </template>
              <!-- 绑定title-click事件的通知列表中标题是“可点击”的-->
              <NoticeList :list="item.list" v-if="item.key === '1'" @title-click="onNoticeClick" />
              <NoticeList :list="item.list" v-else />
            </TabPane>
          </template>
        </Tabs>
      </template>
    </Popover>
  </div>
</template>
<script lang="ts">
  import { computed, defineComponent, onMounted, ref, unref } from 'vue';
  import { Popover, Tabs, Badge } from 'ant-design-vue';
  import { BellOutlined } from '@ant-design/icons-vue';
  import { ListItem, TabItem } from './data';
  import NoticeList from './NoticeList.vue';
  import { useDesign } from '/@/hooks/web/useDesign';
  import { useMessage } from '/@/hooks/web/useMessage';
  import { formatToDateTime } from '/@/utils/dateUtil';
  import { useIntervalFn } from '@vueuse/core';
  import { getMessageUnreadApi, readMessageApi } from '/@/api/sys/user';

  export default defineComponent({
    components: { Popover, BellOutlined, Tabs, TabPane: Tabs.TabPane, Badge, NoticeList },
    setup() {
      const ONE_SECONDS = 1000;
      const { prefixCls } = useDesign('header-notify');
      const { createMessage } = useMessage();

      const reload = async () => {
        const unreadData = await getMessageUnreadApi();
        listData.value = [
          {
            key: '1',
            name: '通知',
            list: unreadData.noticeList.map((x) => {
              return {
                id: String(x.id),
                avatar: 'ant-design:form-outline',
                title: x.title,
                description: x.content,
                datetime: formatToDateTime(x.createdAt),
                type: '1',
              };
            }),
          },
          {
            key: '2',
            name: '消息',
            list: unreadData.messageList.map((x) => {
              return {
                id: String(x.id),
                avatar: 'ant-design:form-outline',
                title: x.title,
                description: x.content,
                datetime: formatToDateTime(x.createdAt),
                type: '1',
              };
            }),
          },
        ];
      };

      onMounted(async () => {
        await reload();
      });

      useIntervalFn(reload, 300 * ONE_SECONDS);

      const listData = ref<TabItem[]>([]);

      const count = computed(() => {
        let count = 0;
        for (let i = 0; i < unref(listData).length; i++) {
          count += unref(listData)[i].list.length;
        }
        return count;
      });

      function onNoticeClick(record: ListItem) {
        createMessage.success('你点击了通知，ID=' + record.id);
        readMessageApi([Number(record.id)], !record.titleDelete);
        // 可以直接将其标记为已读（为标题添加删除线）,此处演示的代码会切换删除线状态
        record.titleDelete = !record.titleDelete;
      }

      return {
        prefixCls,
        listData,
        count,
        onNoticeClick,
        numberStyle: {},
      };
    },
  });
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
