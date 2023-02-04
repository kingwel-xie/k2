<template>
  <div>
    <BasicTable @register="registerTable">
      <template #toolbar>
        <a-button
          v-auth="'admin:sysRole:add'"
          preIcon="ant-design:file-add-outlined"
          type="primary"
          @click="handleCreate"
          >{{ t('common.addText') }}</a-button
        >
      </template>
      <template #bodyCell="{ column, record }">
        <template v-if="column.key === 'action'">
          <TableAction
            :actions="[
              {
                icon: 'ant-design:form-outlined',
                tooltip: t('common.editText'),
                onClick: handleEdit.bind(null, record),
                auth: 'admin:sysRole:edit',
              },
              {
                icon: 'ant-design:delete-outlined',
                tooltip: t('common.delText'),
                color: 'error',
                auth: 'admin:sysRole:remove',
                popConfirm: {
                  title: t('sys.app.deletePrompt'),
                  placement: 'left',
                  confirm: handleDelete.bind(null, record),
                },
              },
            ]"
          />
        </template>
      </template>
    </BasicTable>
    <EditModal @register="registerModal" @success="handleSuccess" />
  </div>
</template>
<script lang="ts">
  import { defineComponent } from 'vue';
  import { BasicTable, useTable, TableAction } from '/@/components/Table';
  import { deleteRoleEntry, getRoleByKey, getRoleListByPage } from '/@/api/admin/system';
  import { useModal } from '/@/components/Modal';
  import EditModal from './EditModal.vue';
  import { columns, searchFormSchema } from './data';
  import { useI18n } from '/@/hooks/web/useI18n';

  export default defineComponent({
    name: 'SysRoleManage',
    components: { BasicTable, EditModal, TableAction },
    setup() {
      const { t } = useI18n();
      const [registerModal, { openModal }] = useModal();
      const [registerTable, { reload, updateTableDataRecord }] = useTable({
        title: '角色列表',
        api: getRoleListByPage,
        columns,
        formConfig: {
          labelWidth: 100,
          baseColProps: { span: 6 },
          schemas: searchFormSchema,
        },
        rowKey: 'roleId',
        useSearchForm: true,
        showTableSetting: true,
        bordered: true,
        showIndexColumn: false,
        actionColumn: {
          width: 80,
          title: t('common.actionText'),
          dataIndex: 'action',
          fixed: undefined,
        },
      });

      function handleCreate() {
        openModal(true, {
          isUpdate: false,
        });
      }

      function handleEdit(record: Recordable) {
        openModal(true, {
          record,
          isUpdate: true,
        });
      }

      async function handleDelete(record: Recordable) {
        await deleteRoleEntry(record);
        await reload();
      }

      async function handleSuccess(isUpdate: boolean, record: Recordable) {
        if (isUpdate) {
          const data = await getRoleByKey(record.roleId);
          updateTableDataRecord(record.roleId, data);
        } else {
          await reload();
        }
      }

      return {
        registerTable,
        registerModal,
        handleCreate,
        handleEdit,
        handleDelete,
        handleSuccess,
        t,
      };
    },
  });
</script>
