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
  import { addAccountEntry, updateAccountEntry, isAccountExist } from '/@/api/admin/system';
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

    const validatorUsername = async (_, value) => {
      if (unref(isUpdate)) return Promise.resolve();
      if (!value) return Promise.reject('请输入用户名');
      try {
        const { code, msg } = await isAccountExist(value);
        switch (code) {
          case 551:
            break;
          case 200:
            return Promise.reject('已被占用');
          default:
            return Promise.reject(msg);
        }
        return Promise.resolve();
      } catch (err: any) {
        return Promise.reject(err.message || '验证失败');
      }
    };

    await updateSchema([
      {
        field: 'username',
        dynamicDisabled: unref(isUpdate),
        rules: [{ required: true, validator: validatorUsername, trigger: 'blur' }],
      },
      {
        field: 'password',
        show: !unref(isUpdate),
        required: !unref(isUpdate),
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
        await updateAccountEntry(values);
      }

      closeModal();
      emit('success', unref(isUpdate), values);
    } finally {
      setModalProps({ confirmLoading: false });
    }
  }
</script>
