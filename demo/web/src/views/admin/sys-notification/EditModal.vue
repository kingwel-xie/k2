<template>
  <BasicModal
    v-bind="$attrs"
    @register="registerModal"
    :title="title"
    width="550px"
    @ok="handleSubmit"
  >
    <BasicForm @register="registerForm">
      <template #targets="{ model, field }">
        <ApiTreeSelect
          v-if="model['targetType'] === 'dept'"
          v-model:value="model[field]"
          :api="getDeptTree"
          allowClear
          treeCheckable
          :showCheckedStrategy="TreeSelect.SHOW_ALL"
          :field-names="{ label: 'label', value: 'id', options: 'children' }"
        />
        <ApiSelect
          v-else-if="model['targetType'] === 'user'"
          v-model:value="model[field]"
          :api="listAccountNoCheck"
          allowClear
          showSearch
          mode="multiple"
          optionFilterProp="label"
          resultField="list"
          labelField="username"
          valueField="username"
        />
        <ApiSelect
          v-else-if="model['targetType'] === 'role'"
          v-model:value="model[field]"
          :api="listRoleNoCheck"
          allowClear
          showSearch
          mode="multiple"
          optionFilterProp="label"
          resultField="list"
          labelField="roleName"
          valueField="roleKey"
        />
      </template>
    </BasicForm>
  </BasicModal>
</template>
<script lang="ts" setup>
  import { ref, unref } from 'vue';
  import { TreeSelect } from 'ant-design-vue';
  import { BasicForm, useForm } from '/@/components/Form/index';
  import { formSchema, TargetTypeOptionData } from './data';
  import { BasicModal, useModalInner } from '/@/components/Modal';
  import ApiSelect from '/@/components/Form/src/components/ApiSelect.vue';
  import ApiTreeSelect from '/@/components/Form/src/components/ApiTreeSelect.vue';

  import { useI18n } from '/@/hooks/web/useI18n';
  import {
    getSysNotificationByKey,
    addSysNotificationEntry,
    updateSysNotificationEntry,
  } from '/@/api/admin/sys-notification';
  import { listAccountNoCheck, listRoleNoCheck, getDeptTree } from '/@/api/admin/system';

  const { t } = useI18n();
  const emit = defineEmits(['register', 'success']);

  const isUpdate = ref(true);
  const optionData = ref<TargetTypeOptionData>();
  const title = ref('');

  // const targetOptions = computed(() => {
  //   switch (unref(currentTargetType)) {
  //     case 'all':
  //     default:
  //       return [];
  //     case 'role':
  //       return unref(optionData)?.sysRoles;
  //     case 'user':
  //       return unref(optionData)?.sysUsers;
  //     case 'dept':
  //       return unref(optionData)?.sysDeptList;
  //   }
  // });

  const [registerForm, { resetFields, setFieldsValue, validate, updateSchema }] = useForm({
    labelWidth: 90,
    baseColProps: { span: 24 },
    schemas: formSchema,
    showActionButtonGroup: false,
  });

  const [registerModal, { setModalProps, closeModal }] = useModalInner(async (data) => {
    await resetFields();
    isUpdate.value = !!data?.isUpdate;
    optionData.value = data.optionData;

    if (unref(isUpdate)) {
      setModalProps({ loading: true });
      data.record = await getSysNotificationByKey(data.record.id);
      await setFieldsValue({
        ...data.record,
      });
      setModalProps({ loading: false });
      title.value = t('common.modalEditText') + data.record.id;
    } else {
      title.value = t('common.modalAddText');
    }

    await updateSchema([
      {
        field: 'id',
        dynamicDisabled: unref(isUpdate),
      },
      {
        field: 'targetType',
        dynamicDisabled: unref(isUpdate),
      },
    ]);
  });

  async function handleSubmit() {
    try {
      const values = await validate();
      // console.log('submit', values);
      values.targets = JSON.stringify(values.targets);

      setModalProps({ confirmLoading: true });
      if (!unref(isUpdate)) {
        await addSysNotificationEntry(values);
      } else {
        await updateSysNotificationEntry(values);
      }

      closeModal();
      emit('success', unref(isUpdate), values);
    } finally {
      setModalProps({ confirmLoading: false });
    }
  }
</script>
