<template>
  <div>
    <BasicTable @register="registerTable">
      <template #toolbar>
        <a-button
          v-auth="'admin:sysDept:add'"
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
                auth: 'admin:sysDept:edit',
              },
              {
                icon: 'ant-design:delete-outlined',
                tooltip: t('common.delText'),
                color: 'error',
                auth: 'admin:sysDept:remove',
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
  import { deleteDeptEntry, getDeptByKey, getDeptList } from '/@/api/admin/system';
  import { useModal } from '/@/components/Modal';
  import EditModal from './EditModal.vue';
  import { columns, searchFormSchema } from './data';
  import { useI18n } from '/@/hooks/web/useI18n';

  export default defineComponent({
    name: 'SysDeptManage',
    components: { BasicTable, EditModal, TableAction },
    setup() {
      const { t } = useI18n();
      const [registerModal, { openModal }] = useModal();
      const [registerTable, { reload, updateTableDataRecord }] = useTable({
        title: '部门列表',
        api: getDeptList,
        columns,
        formConfig: {
          labelWidth: 100,
          baseColProps: { span: 6 },
          schemas: searchFormSchema,
        },
        rowKey: 'deptId',
        pagination: false,
        striped: false,
        useSearchForm: true,
        showTableSetting: true,
        bordered: true,
        showIndexColumn: false,
        canResize: false,
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
        await deleteDeptEntry(record);
        await reload();
      }

      async function handleSuccess(isUpdate: boolean, record: Recordable) {
        if (isUpdate) {
          const data = await getDeptByKey(record.deptId);
          updateTableDataRecord(record.deptId, data);
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
