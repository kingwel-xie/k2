<template>
  <BasicModal
    v-bind="$attrs"
    @register="registerModal"
    :title="title"
    :width="720"
    @ok="handleSubmit"
  >
    <BasicForm @register="registerForm">
      <template #tokenEnabled="{ model, field }">
        <Checkbox
          v-model:checked="model[field]"
          @change="(e) => handleTokenEnabledChange(e.target.checked, model, field)"
          >启用 API Token
        </Checkbox>
      </template>
    </BasicForm>
    <Row v-if="isUpdate" :gutter="10">
      <Col :offset="4" :span="1.5">
        <a-button
          :disabled="resetButtonDisabled"
          preIcon="ant-design:alert-outlined"
          color="warning"
          @click="handleResetToken"
        >
          重置Token
        </a-button>
      </Col>
    </Row>
  </BasicModal>
</template>
<script lang="ts" setup>
  import { h, ref, unref } from 'vue';
  import { Row, Col, Checkbox } from 'ant-design-vue';
  import { BasicModal, useModalInner } from '/@/components/Modal';
  import { BasicForm, useForm } from '/@/components/Form';
  import { accountFormSchema } from './data';
  import {
    addAccountEntry,
    updateAccountEntry,
    isAccountExist,
    resetUserToken,
  } from '/@/api/admin/system';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { useMessage } from '/@/hooks/web/useMessage';

  const { t } = useI18n();
  const emit = defineEmits(['success', 'register']);

  const isUpdate = ref(true);
  const title = ref('');
  const resetButtonDisabled = ref(false);

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

    resetButtonDisabled.value = !data.record?.token;
    if (unref(isUpdate)) {
      data.record.tokenEnabled = !!data.record.token;
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

  function handleTokenEnabledChange(val, model, _field) {
    if (!val) {
      model['token'] = undefined;
      resetButtonDisabled.value = true;
    } else {
      resetButtonDisabled.value = false;
    }
  }

  function handleResetToken() {
    const { createConfirm, createMessage } = useMessage();
    createConfirm({
      iconType: 'warning',
      title: () => h('span', '警告'),
      content: () => h('span', '重新生成 API Token 可能会影响正在使用的，是否继续？'),
      onOk: async () => {
        const values = await validate();
        const newToken = await resetUserToken(values.userId);
        await setFieldsValue({
          token: newToken,
        });
        createMessage.success('Token重置成功');
      },
    });
  }
</script>
