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
  import { formSchema } from './typeData';
  import { BasicModal, useModalInner } from '/@/components/Modal';
  import { useI18n } from '/@/hooks/web/useI18n';
  import {
    getSysDictTypeByKey,
    addSysDictTypeEntry,
    updateSysDictTypeEntry,
  } from '/@/api/admin/sys-dict';

  const { t } = useI18n();
  const emit = defineEmits(['register', 'success']);

  const isUpdate = ref(true);
  const title = ref('');

  const [registerForm, { resetFields, setFieldsValue, validate }] = useForm({
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
      data.record = await getSysDictTypeByKey(data.record.id);
      data.record.status = String(data.record.status);
      await setFieldsValue({
        ...data.record,
      });
      setModalProps({ loading: false });
      title.value = t('common.modalEditText') + data.record.dictName;
    } else {
      title.value = t('common.modalAddText');
    }
  });

  async function handleSubmit() {
    try {
      const values = await validate();
      values.status = parseInt(values.status);

      setModalProps({ confirmLoading: true });
      if (!unref(isUpdate)) {
        await addSysDictTypeEntry(values);
      } else {
        await updateSysDictTypeEntry(values);
      }

      closeModal();
      emit('success', unref(isUpdate), values);
    } finally {
      setModalProps({ confirmLoading: false });
    }
  }
</script>
