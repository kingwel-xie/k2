<template>
  <BasicModal
    v-bind="$attrs"
    @register="registerModal"
    :title="title"
    width="550px"
    @ok="handleSubmit"
  >
    <BasicForm @register="registerForm" />
  </BasicModal>
</template>
<script lang="ts" setup>
  import { ref, unref } from 'vue';
  import { BasicForm, useForm } from '/@/components/Form/index';
  import { formSchema } from './data';
  import { BasicModal, useModalInner } from '/@/components/Modal';
  import { useI18n } from '/@/hooks/web/useI18n';
  import {
    getTbxCountryByKey,
    addTbxCountryEntry,
    updateTbxCountryEntry,
  } from '/@/api/kobh/tbx-country';

  const { t } = useI18n();
  const emit = defineEmits(['register', 'success']);

  const isUpdate = ref(true);
  const title = ref('');

  const [registerForm, { resetFields, setFieldsValue, validate, updateSchema }] = useForm({
    labelWidth: 90,
    baseColProps: { span: 24 },
    schemas: formSchema,
    showActionButtonGroup: false,
  });

  const [registerModal, { setModalProps, closeModal }] = useModalInner(async (data) => {
    await resetFields();
    isUpdate.value = !!data?.isUpdate;

    if (unref(isUpdate)) {
      setModalProps({ loading: true });
      data.record = await getTbxCountryByKey(data.record.code);
      await setFieldsValue({
        ...data.record,
      });
      setModalProps({ loading: false });
      title.value = t('common.modalEditText') + data.record.code;
    } else {
      title.value = t('common.modalAddText');
    }

    await updateSchema([
      {
        field: 'code',
        dynamicDisabled: unref(isUpdate),
      },
    ]);
  });

  async function handleSubmit() {
    try {
      const values = await validate();

      setModalProps({ confirmLoading: true });
      if (!unref(isUpdate)) {
        await addTbxCountryEntry(values);
      } else {
        await updateTbxCountryEntry(values);
      }

      closeModal();
      emit('success', unref(isUpdate), values);
    } finally {
      setModalProps({ confirmLoading: false });
    }
  }
</script>
