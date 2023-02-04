<template>
  <BasicModal
    v-bind="$attrs"
    @register="registerModal"
    :title="title"
    :width="720"
    @ok="handleSubmit"
  >
    <BasicForm @register="registerForm" />
  </BasicModal>
</template>
<script lang="ts" setup>
  import { ref, unref } from 'vue';
  import { BasicModal, useModalInner } from '/@/components/Modal';
  import { BasicForm, useForm } from '/@/components/Form';
  import { accountFormSchema } from './data';
  import {
    addAccountEntry,
    updateAccountEntry,
    getDeptTree,
    isAccountExist,
  } from '/@/api/admin/system';
  import { useI18n } from '/@/hooks/web/useI18n';

  const { t } = useI18n();
  const emit = defineEmits(['success', 'register']);

  const isUpdate = ref(true);
  const title = ref('');

  const [registerForm, { setFieldsValue, updateSchema, resetFields, validate }] = useForm({
    labelWidth: 100,
    baseColProps: { span: 12 },
    schemas: accountFormSchema,
    showActionButtonGroup: false,
    actionColOptions: {
      span: 23,
    },
  });

  const [registerModal, { setModalProps, closeModal }] = useModalInner(async (data) => {
    await resetFields();
    isUpdate.value = !!data?.isUpdate;

    if (unref(isUpdate)) {
      await setFieldsValue({
        ...data.record,
      });
      title.value = t('common.modalEditText') + data.record.username;
    } else {
      title.value = t('common.modalAddText');
    }

    const validator = async (_, value) => {
      return new Promise<void>((resolve, reject) => {
        if (unref(isUpdate)) return resolve();
        isAccountExist(value)
          .then(({ code, msg }) => {
            switch (code) {
              case 551:
                break;
              case 200:
                return Promise.reject({ message: '占用' });
              default:
                return Promise.reject({ message: msg });
            }
            return resolve();
          })
          .catch((err) => {
            reject(err.message || '验证失败');
          });
      });
    };

    const treeData = await getDeptTree();
    await updateSchema([
      {
        field: 'username',
        dynamicDisabled: unref(isUpdate),
        rules: [
          {
            required: !unref(isUpdate),
            message: '请输入用户名',
            validator,
          },
        ],
      },
      {
        field: 'password',
        show: !unref(isUpdate),
        required: !unref(isUpdate),
      },
      {
        field: 'deptId',
        componentProps: { treeData },
      },
    ]);
  });

  async function handleSubmit() {
    try {
      const values = await validate();
      setModalProps({ confirmLoading: true });
      if (!unref(isUpdate)) {
        await addAccountEntry(values);
      } else {
        console.log('sysuser', values);
        await updateAccountEntry(values);
      }

      closeModal();
      emit('success', unref(isUpdate), values);
    } finally {
      setModalProps({ confirmLoading: false });
    }
  }
</script>
