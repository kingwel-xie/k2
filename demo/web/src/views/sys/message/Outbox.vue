<template>
  <CollapseContainer :canExpan="false" class="container-setting">
    <template #title>
      <span class="text-md mr-3">{{ t('common.sentMessage') }}({{ total }})</span>
      <a-button
        type="dashed"
        shape="circle"
        preIcon="ant-design:reload-outlined"
        size="small"
        @click="handleReloadMessage"
        class="mr-2"
      />
    </template>
    <template #action>
      <a-button
        type="danger"
        size="small"
        :disabled="messageList.length === 0"
        @click="handleDeleteAll"
      >
        全部删除
      </a-button>
    </template>
    <List :loading="initLoading" :dataSource="messageList">
      <template v-if="messageList.length < total" #loadMore>
        <div
          v-if="!initLoading && !loading"
          :style="{ textAlign: 'center', marginTop: '12px', height: '32px', lineHeight: '32px' }"
        >
          <a-button @click="onLoadMore">{{ t('common.loadMore') }}</a-button>
        </div>
      </template>
      <template #renderItem="{ item }">
        <ListItem>
          <template #actions>
            <Popover :content="t('common.delText')">
              <a-button
                type="dashed"
                shape="circle"
                size="small"
                preIcon="ant-design:delete-outlined"
                @click="handleDeleteOne(item)"
              />
            </Popover>
          </template>
          <Skeleton avatar :title="false" :loading="!!item.loading" active>
            <ListItemMeta>
              <template #avatar>
                <Icon icon="ic:outline-email" color="success" :size="24" />
              </template>
              <template #title>
                <div class="flex justify-between">
                  <TTitle type="secondary" :level="5" :content="item.title" />
                  <div class="text-secondary">
                    发送至
                    <Receivers :receivers="item.receivers" />
                    <span>
                      {{ formatToDateTime(item.createdAt) }}
                    </span>
                  </div>
                </div>
              </template>
              <template #description>
                <ShowContent :value="item.content" />
                <div v-if="false" style="white-space: pre-wrap">
                  <TParagraph
                    type="secondary"
                    :content="item.content"
                    :ellipsis="{ rows: 2, expandable: true, symbol: 'more' }"
                  />
                </div>
              </template>
            </ListItemMeta>
          </Skeleton>
        </ListItem>
      </template>
    </List>
  </CollapseContainer>
</template>
<script lang="ts" setup>
  import { onMounted, ref, nextTick, unref } from 'vue';
  import {
    Popover,
    List,
    ListItem,
    ListItemMeta,
    Skeleton,
    // TypographyText as TText,
    TypographyTitle as TTitle,
    TypographyParagraph as TParagraph,
  } from 'ant-design-vue';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { CollapseContainer } from '/@/components/Container';
  import { Icon } from '/@/components/Icon';
  import { formatToDateTime } from '/@/utils/dateUtil';
  import {
    deleteSysOutboxEntry,
    deleteSysOutboxMany,
    getSysOutboxList,
  } from '/@/api/admin/sys-outbox';
  import Receivers from '/@/views/sys/message/Receivers.vue';
  import ShowContent from '/@/views/sys/message/ShowContent.vue';

  const { t } = useI18n();

  onMounted(async () => {
    await reload();
  });

  async function reload() {
    const params = { pageIndex: unref(page), pageSize: 10 };
    const res = await getSysOutboxList(params);
    initLoading.value = false;
    messageList.value = res.list;
    total.value = res['count'];
  }

  async function onLoadMore() {
    loading.value = true;
    const original = messageList.value;
    // for skeleton
    messageList.value = messageList.value.concat(
      [...new Array(3)].map(() => ({ loading: true, title: '', content: '' })),
    );

    page.value += 1;
    const params = { pageIndex: unref(page), pageSize: 10 };
    const res = await getSysOutboxList(params);
    messageList.value = original.concat(res.list);
    total.value = res['count'];
    loading.value = false;
    await nextTick(() => {
      // Resetting window's offsetTop to display react-virtualized demo underfloor.
      // In real scene, you can use public method of react-virtualized:
      // https://stackoverflow.com/questions/46700726/how-to-use-public-method-updateposition-of-react-virtualized
      window.dispatchEvent(new Event('resize'));
    });
  }

  const page = ref(1);
  const total = ref(0);
  const initLoading = ref(true);
  const loading = ref(false);
  const messageList = ref<Recordable[]>([]);

  async function handleDeleteOne(item: Recordable) {
    const foundIndex = unref(messageList).findIndex((x) => x.id === item.id);
    if (foundIndex !== -1) {
      await deleteSysOutboxEntry(item);
      messageList.value.splice(foundIndex, 1);
      total.value -= 1;
    }
  }

  async function handleDeleteAll() {
    const messages = unref(messageList);
    if (!messages.length) return;
    await deleteSysOutboxMany(messages.map((item) => item.id));
    messageList.value = [];
    total.value = 0;
  }

  function handleReloadMessage() {
    reload();
  }
</script>

<style lang="less" scoped>
  .container-setting {
    margin: 12px;
  }
</style>
