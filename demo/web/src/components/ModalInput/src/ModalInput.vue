<template>
  <BasicModal v-bind="$attrs" @register="registerModal" @ok="handleSubmit" @cancel="handleCancel">
    <BasicForm ref="formRef" @register="registerForm" />
  </BasicModal>
</template>

<script lang="ts" setup>
  import { BasicModal, useModalInner } from '/@/components/Modal';
  import { BasicForm, FormSchema, useForm } from '/@/components/Form/index';
  import { propTypes } from '/@/utils/propTypes';
  import { ComponentType } from '/@/components/Form/src/types';
  import { InputCallback } from '/@/components/ModalInput';
  import { nextTick, ref, unref } from 'vue';

  const props = defineProps({
    label: propTypes.string.def(''),
    inputType: propTypes.string.def('Input').isRequired,
    inputContent: propTypes.any,
    inputComponentProps: propTypes.object.def({}),
    inputValidator: propTypes.func,
  });

  const formRef = ref();

  defineExpose({ openModal });

  const callbacks: { ok?: InputCallback; cancel?: InputCallback } = {};

  function openModal(onOk: InputCallback, onCancel: InputCallback) {
    callbacks.ok = onOk;
    callbacks.cancel = onCancel;

    setModalProps({
      visible: true,
    });
    nextTick(async () => {
      // clear all fields
      await resetFields();
      if (props.inputContent) {
        await setFieldsValue({
          input: props.inputContent,
        });
      }
      const formEl = unref(formRef);
      const el = (formEl as any)?.$el as HTMLElement;
      if (!formEl || !el) {
        return;
      }
      const inputEl = el.querySelector(
        '.ant-row:first-child input, .ant-row:first-child textarea',
      ) as Nullable<HTMLInputElement>;
      if (!inputEl) return;
      inputEl.focus();
    });
  }

  const formSchema: FormSchema[] = [
    {
      field: 'input',
      label: props.label,
      component: props.inputType as ComponentType,
      componentProps: props.inputComponentProps,
      rules: [
        {
          required: true,
          message: '请输入数据',
        },
        {
          validator(_, value) {
            return new Promise((resolve, reject) => {
              if (!props.inputValidator) return resolve();
              const ret = props.inputValidator(value);
              if (ret === true) {
                return resolve();
              }
              return reject(ret);
            });
          },
        },
      ],
      required: true,
    },
  ];

  const [registerForm, { resetFields, setFieldsValue, validate }] = useForm({
    labelWidth: 100,
    baseColProps: { span: 24 },
    schemas: formSchema,
    autoSubmitOnEnter: true,
    submitFunc: handleSubmit,
    showActionButtonGroup: false,
    actionColOptions: {
      span: 23,
    },
  });

  const [registerModal, { setModalProps, closeModal }] = useModalInner();

  async function handleSubmit() {
    try {
      const values = await validate();
      closeModal();
      callbacks.ok && callbacks.ok(values);
    } finally {
    }
  }

  function handleCancel() {
    callbacks.cancel && callbacks.cancel({});
  }
</script>
