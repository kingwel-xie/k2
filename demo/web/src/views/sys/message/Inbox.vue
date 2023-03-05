<template>
  <CollapseContainer :canExpan="false" class="inbox-setting">
    <template #title>
      <span class="text-md mr-3">
        {{ t('common.newMessage') }}({{ unreadMessages }}/{{ totalMessages }})
      </span>
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
      <span class="mr-20">
        <span>过滤条件：</span>
        <DictSelect
          size="small"
          v-model:value="model.type"
          dictName="sys_notice_type"
          style="min-width: 100px"
          allowClear
          :showSearch="false"
          @change="page = 1"
        />
      </span>
      <a-button type="primary" size="small" @click="handleNewMessage" class="mr-2">
        写新消息
      </a-button>
      <a-button
        type="warning"
        size="small"
        :disabled="filteredList.filter((item) => !item.read).length === 0"
        @click="handleReadAll"
      >
        全部已读
      </a-button>
    </template>
    <List :loading="initLoading" :dataSource="filteredList">
      <template v-if="messageList.length < totalMessages" #loadMore>
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
            <Switch
              v-model:checked="item.read"
              checked-children="已读"
              un-checked-children="未读"
              @change="(val) => handleReadOne(item, val)"
            />
            <Popover :content="t('common.replyMessage')">
              <a-button
                type="dashed"
                shape="circle"
                size="small"
                preIcon="ant-design:rollback-outlined"
                @click="handleReply(item)"
              />
            </Popover>
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
                <Icon :icon="itemIcon(item)" :color="itemColor(item)" :size="24" />
              </template>
              <template #title>
                <div class="flex justify-between">
                  <div class="flex" style="align-items: flex-start">
                    <Tag v-if="item.type === 'notification'" :color="itemColor(item)">通知</Tag>
                    <TTitle
                      v-if="item.read"
                      type="secondary"
                      :level="5"
                      :delete="item.read"
                      :content="item.title"
                    />
                    <TTitle v-else :level="5" :delete="item.read" :content="item.title" />
                  </div>
                  <div class="text-secondary">
                    来自于
                    <span class="mx-1 cursor-pointer underline" @click="handleReply(item)">
                      {{ item.sender }}
                    </span>
                    <span>
                      {{ formatToDateTime(item.createdAt) }}
                    </span>
                  </div>
                </div>
              </template>
              <template #description>
                <div style="white-space: pre-wrap">
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
    <EditModal @register="registerModal" @success="handleSuccess" />
  </CollapseContainer>
</template>
<script lang="ts" setup>
  import { onMounted, ref, nextTick, unref, reactive, computed } from 'vue';
  import {
    Popover,
    List,
    ListItem,
    ListItemMeta,
    Skeleton,
    Switch,
    Tag,
    // TypographyText as TText,
    TypographyTitle as TTitle,
    TypographyParagraph as TParagraph,
  } from 'ant-design-vue';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { CollapseContainer } from '/@/components/Container';
  import DictSelect from '/@/components/Form/src/components/DictSelect.vue';
  import { Icon } from '/@/components/Icon';
  import { deleteSysInboxEntry, getSysInboxList } from '/@/api/admin/sys-inbox';
  import { useModal } from '/@/components/Modal';
  import { formatToDateTime } from '/@/utils/dateUtil';
  import EditModal from './EditModal.vue';
  import { useUserStoreWithOut } from '/@/store/modules/user';
  import { getMessageUnreadApi, readMessageApi } from '/@/api/sys/user';

  const { t } = useI18n();

  onMounted(async () => {
    await reload();
  });

  const filteredList = computed(() => {
    return model.type
      ? unref(messageList).filter((x) => x.type === model.type)
      : unref(messageList);
  });

  async function reload() {
    const params = { pageIndex: unref(page), pageSize: 10, type: model.type };
    const res = await getSysInboxList(params);
    initLoading.value = false;
    messageList.value = res.list;
    totalMessages.value = res['count'];
    const unread = await getMessageUnreadApi();
    unreadMessages.value = unread.numMessages + unread.numNotices;
    // update userStore, so that the notify-header gets updated instantly
    const userStore = useUserStoreWithOut();
    userStore.setUnreadMessage(unread);
  }

  async function onLoadMore() {
    loading.value = true;
    const original = messageList.value;
    // for skeleton
    messageList.value = messageList.value.concat(
      [...new Array(3)].map(() => ({ loading: true, title: '', content: '' })),
    );

    page.value += 1;
    const params = { pageIndex: unref(page), pageSize: 10, type: model.type };
    const res = await getSysInboxList(params);
    messageList.value = original.concat(res.list);
    totalMessages.value = res['count'];
    loading.value = false;
    await nextTick(() => {
      // Resetting window's offsetTop to display react-virtualized demo underfloor.
      // In real scene, you can use public method of react-virtualized:
      // https://stackoverflow.com/questions/46700726/how-to-use-public-method-updateposition-of-react-virtualized
      window.dispatchEvent(new Event('resize'));
    });
  }

  const page = ref(1);
  const totalMessages = ref(0);
  const unreadMessages = ref(0);
  const initLoading = ref(true);
  const loading = ref(false);
  const messageList = ref<Recordable[]>([]);
  const model = reactive({
    type: undefined,
  });

  const [registerModal, { openModal }] = useModal();

  function handleReadOne(item: Recordable, val: boolean) {
    readMessageApi([item.id], val);
    unreadMessages.value += val ? -1 : 1;
  }

  async function handleReadAll() {
    const messages = unref(messageList).filter((item) => !item.read);
    if (!messages.length) return;
    await readMessageApi(
      messages.map((item) => item.id),
      true,
    );
    messages.forEach((item) => (item.read = true));
    unreadMessages.value = 0;
  }

  async function handleDeleteOne(item: Recordable) {
    const foundIndex = unref(messageList).findIndex((x) => x.id === item.id);
    if (foundIndex !== -1) {
      await deleteSysInboxEntry(item);
      messageList.value.splice(foundIndex, 1);
      totalMessages.value -= 1;
      if (!item.read) {
        unreadMessages.value -= 1;
      }
    }
  }

  function handleReloadMessage() {
    reload();
  }

  function handleNewMessage() {
    openModal(true, {});
  }

  function handleReply(item: Recordable) {
    openModal(true, {
      message: item,
    });
  }

  function handleSuccess() {}

  function itemIcon(item: any) {
    if (item.type === 'notification') {
      return item.read ? 'ic:outline-notifications-off' : 'ic:outline-notifications';
    } else {
      return item.read ? 'ic:outline-mark-email-read' : 'ic:outline-email';
    }
  }

  function itemColor(item: any) {
    if (item.type === 'notification') {
      return item.read ? 'lightgray' : 'red';
    } else {
      return item.read ? 'lightgray' : 'blue';
    }
  }
</script>

<style lang="less" scoped>
  .inbox-setting {
    margin: 12px;
    background-color: @component-background;
  }
  .extra {
    float: right;
    margin-top: 10px;
    margin-right: 20px;
  }
</style>
