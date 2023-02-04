<template>
  <BasicModal v-bind="$attrs" @register="registerModal" :title="title" @ok="handleSubmit">
    <BasicForm @register="registerForm" />
  </BasicModal>
</template>
<script lang="ts" setup>
  import { ref, unref } from 'vue';
  import { BasicModal, useModalInner } from '/@/components/Modal';
  import { BasicForm, useForm } from '/@/components/Form';
  import { formSchema } from './data';
  import { addDeptEntry, getDeptTree, updateDeptEntry } from '/@/api/admin/system';
  import { useI18n } from '/@/hooks/web/useI18n';

  const { t } = useI18n();
  const emit = defineEmits(['success', 'register']);

  const isUpdate = ref(true);
  const title = ref('');

  const [registerForm, { resetFields, setFieldsValue, updateSchema, validate }] = useForm({
    labelWidth: 100,
    baseColProps: { span: 24 },
    schemas: formSchema,
    showActionButtonGroup: false,
  });

  const [registerModal, { setModalProps, closeModal }] = useModalInner(async (data) => {
    const treeData = await getDeptTree();
    // fix sys-dept tree data, assign id 0 a label
    const root = { id: 0, label: '-' };
    treeData.unshift(root);

    await resetFields();
    isUpdate.value = !!data?.isUpdate;

    if (unref(isUpdate)) {
      // kingwel, unfortunately, status in 'sys-dept' is an integer
      data.record.status = String(data.record.status);
      await setFieldsValue({
        ...data.record,
      });
      title.value = t('common.modalEditText') + data.record.deptName;
    } else {
      title.value = t('common.modalAddText');
    }

    await updateSchema({
      field: 'parentId',
      componentProps: { treeData },
      dynamicDisabled: unref(isUpdate),
    });
  });

  async function handleSubmit() {
    try {
      const values = await validate();
      setModalProps({ confirmLoading: true });
      // note: status, to integer
      values.status = parseInt(values.status);
      if (!unref(isUpdate)) {
        await addDeptEntry(values);
      } else {
        await updateDeptEntry(values);
      }

      closeModal();
      emit('success', unref(isUpdate), values);
    } finally {
      setModalProps({ confirmLoading: false });
    }
  }
</script>
