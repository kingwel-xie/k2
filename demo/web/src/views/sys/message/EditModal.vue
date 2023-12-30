<template>
  <BasicModal
    v-bind="$attrs"
    @register="registerModal"
    :title="title"
    width="980px"
    @ok="handleSubmit"
  >
    <BasicForm @register="registerForm">
      <!--        <Select v-model:value="model[field]" :options="sysUsers" mode="multiple" />-->
      <!--      </template>-->
      <template #content="{ model, field }">
        <Tinymce v-model:modelValue="model[field]" />
      </template>
    </BasicForm>
  </BasicModal>
</template>
<script lang="ts" setup>
  import { ref } from 'vue';
  import { BasicForm, useForm } from '/@/components/Form';
  import { Tinymce } from '/@/components/Tinymce';
  import { formSchema } from './data';
  import { BasicModal, useModalInner } from '/@/components/Modal';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { addSysInboxEntry } from '/@/api/admin/sys-inbox';

  const { t } = useI18n();
  const emit = defineEmits(['register', 'success']);
  const title = ref('');

  const [registerForm, { resetFields, validate, setFieldsValue }] = useForm({
    labelWidth: 90,
    baseColProps: { span: 24 },
    schemas: formSchema,
    showActionButtonGroup: false,
  });

  const [registerModal, { setModalProps, closeModal }] = useModalInner(async (data) => {
    await resetFields();
    if (data.message) {
      await setFieldsValue({
        receiver: [data.message.sender],
        title: 'Re: ' + data.message.title,
      });
    }
    title.value = t('common.newMessage');
  });

  async function handleSubmit() {
    try {
      const values = await validate();
      setModalProps({ confirmLoading: true });
      await addSysInboxEntry(values);

      closeModal();
      emit('success', false, values);
    } finally {
      setModalProps({ confirmLoading: false });
    }
  }
</script>
