<template>
  <BasicModal
    v-bind="$attrs"
    @register="registerModal"
    :title="title"
    :showCancelBtn="false"
    width="780px"
    @ok="handleSubmit"
  >
    <div class="flex justify-center">
      <TTitle :level="3" :content="currentNotice.title" />
    </div>
    <div class="flex justify-end">
      <div class="text-secondary">
        {{ t('common.inbox.from') }}
        <span class="mx-1 cursor-pointer underline">
          {{ currentNotice.sender }}
        </span>
        <span>
          {{ formatToDateTime(currentNotice.createdAt) }}
        </span>
      </div>
    </div>
    <div v-html="currentNotice.content"></div>
    <template #insertFooter>
      <Checkbox v-model:checked="readAll">{{ t('common.inbox.markAll') }}</Checkbox>
      <a-button
        danger
        preIcon="ant-design:left-square-outlined"
        @click="handleReadPrev"
        :disabled="noticeData.index === 0"
      >
        {{ t('common.inbox.prev') }}
      </a-button>
      <a-button
        danger
        preIcon="ant-design:right-square-outlined"
        @click="handleReadNext"
        :disabled="noticeData.index + 1 === noticeData.list.length"
      >
        {{ t('common.inbox.next') }}
      </a-button>
    </template>
  </BasicModal>
</template>
<script lang="ts" setup>
  import { computed, reactive, ref, unref } from 'vue';
  import { Checkbox, TypographyTitle as TTitle } from 'ant-design-vue';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { BasicModal, useModalInner } from '/@/components/Modal';
  import { readMessageApi } from '/@/api/sys/user';
  import { formatToDateTime } from '/@/utils/dateUtil';

  const emit = defineEmits(['register', 'success']);

  const { t } = useI18n();
  const title = ref();
  const readAll = ref();

  const noticeData = reactive({
    list: [] as Recordable[],
    index: 0,
  });

  const currentNotice = computed(() => {
    const notice = noticeData.list[noticeData.index] || {};
    return notice;
  });

  const [registerModal, { setModalProps, closeModal }] = useModalInner(async (data) => {
    noticeData.list = data;
    noticeData.index = 0;
    title.value = `系统：${data.length} 条重要公告`;
    // read the first one
    if (data.length > 0) {
      await readMessageApi([data[0].id], true);
      data[0].read = true;
    }
  });

  async function handleSubmit() {
    try {
      // mark read all
      if (unref(readAll)) {
        const list = noticeData.list.filter((item) => !item.read).map((item) => item.id);
        if (list.length > 0) {
          await readMessageApi(list, true);
        }
      }

      closeModal();
      emit('success', true, {});
    } finally {
      setModalProps({ confirmLoading: false });
    }
  }

  async function handleReadPrev() {
    if (noticeData.index > 0) {
      noticeData.index -= 1;
    }
  }

  async function handleReadNext() {
    if (noticeData.index < noticeData.list.length - 1) {
      noticeData.index += 1;
      // mark this one as read
      const thisNotice = noticeData.list[noticeData.index];
      if (!thisNotice.read) {
        await readMessageApi([thisNotice.id], true);
        thisNotice.read = true;
      }
    }
  }
</script>
