<template>
  <BasicDrawer
    v-bind="$attrs"
    @register="registerDrawer"
    showFooter
    :title="title"
    width="62%"
    @ok="handleSubmit"
  >
    <BasicForm @register="registerForm" />
  </BasicDrawer>
</template>
<script lang="ts" setup>
  import { ref, unref } from 'vue';
  import { BasicForm, useForm } from '/@/components/Form';
  import { formSchema } from './data';
  import { BasicDrawer, useDrawerInner } from '/@/components/Drawer';

  import { addMenuEntry, getMenuByKey, getMenuTree, updateMenuEntry } from '/@/api/admin/system';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { getSysApiList } from '/@/api/admin/sys-api';

  const { t } = useI18n();
  const emit = defineEmits(['success', 'register']);

  const isUpdate = ref(true);
  const title = ref('');

  const [registerForm, { resetFields, setFieldsValue, updateSchema, validate }] = useForm({
    labelWidth: 110,
    schemas: formSchema,
    showActionButtonGroup: false,
    baseColProps: { span: 12 },
  });

  const [registerDrawer, { setDrawerProps, closeDrawer }] = useDrawerInner(async (data) => {
    await resetFields();
    isUpdate.value = !!data?.isUpdate;

    const treeData = (await getMenuTree()).menus;
    // fix sys-menu tree data, assign id 0 a label
    const root = { id: 0, label: '-' };
    treeData.unshift(root);
    const params = { pageIndex: 1, pageSize: -1, type: 'CHECK', pathOrder: 'ASC' };
    const sysApiList = (await getSysApiList(params)).list;

    if (unref(isUpdate)) {
      setDrawerProps({ loading: true });
      data.record = await getMenuByKey(data.record.menuId);
      data.record.apis = data.record.sysApi.map((x) => String(x.id));
      await setFieldsValue({
        ...data.record,
      });
      setDrawerProps({ loading: false });
      title.value = t('common.modalEditText') + data.record.title;
    } else {
      // set parent id
      await setFieldsValue({
        parentId: data.parentId,
      });
      title.value = t('common.modalAddText');
    }
    await updateSchema([
      {
        field: 'parentId',
        componentProps: { treeData },
      },
      {
        field: 'apis',
        componentProps: { dataSource: sysApiList },
      },
    ]);
  });

  async function handleSubmit() {
    try {
      const values = await validate();
      setDrawerProps({ confirmLoading: true });
      values.apis = (values.apis && values.apis.map((x) => Number(x))) || [];
      // console.log('submir', values);

      if (!unref(isUpdate)) {
        await addMenuEntry(values);
      } else {
        await updateMenuEntry(values);
      }

      closeDrawer();
      emit('success');
    } finally {
      setDrawerProps({ confirmLoading: false });
    }
  }
</script>
