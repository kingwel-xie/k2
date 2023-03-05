<template>
  <BasicModal
    v-bind="$attrs"
    @register="registerModal"
    :title="title"
    width="550px"
    @ok="handleSubmit"
  >
    <BasicForm @register="registerForm">
      <template #menu="{ model, field }">
        <BasicTree
          v-if="model['roleKey'] !== 'admin'"
          ref="treeRef"
          v-model:value="model[field]"
          :treeData="treeData"
          :fieldNames="{ title: 'label', key: 'id' }"
          checkable
          toolbar
          @check="onCheck"
          title="菜单分配"
        />
      </template>
    </BasicForm>
  </BasicModal>
</template>
<script lang="ts" setup>
  import { ref, unref, toRaw } from 'vue';
  import { BasicForm, useForm } from '/@/components/Form';
  import { formSchema } from './data';
  import { BasicModal, useModalInner } from '/@/components/Modal';
  import { BasicTree, TreeActionType, TreeItem } from '/@/components/Tree';
  import { addRoleEntry, getMenuTree, getRoleByKey, updateRoleEntry } from '/@/api/admin/system';
  import { useI18n } from '/@/hooks/web/useI18n';

  const { t } = useI18n();
  const emit = defineEmits(['success', 'register']);

  const isUpdate = ref(true);
  const title = ref('');

  const treeRef = ref<Nullable<TreeActionType>>(null);
  const treeData = ref<TreeItem[]>([]);

  const [registerForm, { resetFields, setFieldsValue, validate }] = useForm({
    labelWidth: 90,
    baseColProps: { span: 24 },
    schemas: formSchema,
    showActionButtonGroup: false,
  });

  const [registerModal, { setModalProps, closeModal }] = useModalInner(async (data) => {
    // 需要在setFieldsValue之前先填充treeData，否则Tree组件可能会报key not exist警告
    if (unref(treeData).length === 0) {
      treeData.value = (await getMenuTree()).menus as any as TreeItem[];
    }
    await resetFields();
    isUpdate.value = !!data?.isUpdate;

    if (unref(isUpdate)) {
      setModalProps({ loading: true });
      data.record = await getRoleByKey(data.record.roleId);
      await setFieldsValue({
        ...data.record,
      });
      setModalProps({ loading: false });
      const tree = unref(treeRef);
      if (!tree && data.record.roleKey !== 'admin') {
        throw new Error('tree is null!');
      }
      tree?.setCheckedKeys(unref(data.record.menuIds));
      title.value = t('common.modalEditText') + data.record.roleName;
    } else {
      title.value = t('common.modalAddText');
    }
    // updateSchema({
    //   field: 'menu',
    //   dynamicDisabled: ({ values }) => values.roleKey === 'admin',
    // });
  });

  async function handleSubmit() {
    try {
      const values = await validate();

      // 'admin' doesn't have menu, so just ignore the menuIds
      if (values.menu) {
        // 目前被选中的菜单节点, v-model of the 'menu' field
        const checkedKeys = toRaw(values.menu);
        // kingwel: have to figure out all halfCheckedKeys, backend needs them
        const findNode = (items, xx, path) => {
          for (let i = 0; i < items.length; i++) {
            if (items[i].children) {
              path.push(items[i].id);
              const rr = findNode(items[i].children, xx, path);
              if (rr) {
                return true;
              }
              path.pop();
            } else {
              if (items[i].id === xx) {
                return true;
              }
            }
          }
          return false;
        };

        const halfCheckedKeys = [];
        (checkedKeys as number[]).forEach((id) => {
          const pp = [];
          if (findNode(treeData.value, id, pp)) {
            pp.forEach((pid) => {
              if (halfCheckedKeys.indexOf(pid) === -1 && checkedKeys.indexOf(pid) === -1) {
                halfCheckedKeys.push(pid);
              }
            });
          }
        });
        // console.log('xxx', halfCheckedKeys, checkedKeys);
        values.menuIds = halfCheckedKeys.concat(checkedKeys as any);
        delete values.menu;
      }

      setModalProps({ confirmLoading: true });
      if (!unref(isUpdate)) {
        await addRoleEntry(values);
      } else {
        await updateRoleEntry(values);
      }

      closeModal();
      emit('success', unref(isUpdate), values);
    } finally {
      setModalProps({ confirmLoading: false });
    }
  }

  function onCheck(checkedKeys, e) {
    console.log(checkedKeys, e);
  }
</script>
