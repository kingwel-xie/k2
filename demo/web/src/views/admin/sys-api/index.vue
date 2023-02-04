<template>
  <div>
    <BasicTable @register="registerTable">
      <template #bodyCell="{ column, record }">
        <template v-if="column.key === 'action'">
          <TableAction
            :actions="[
              {
                icon: 'ant-design:form-outlined',
                tooltip: t('common.editText'),
                onClick: handleEdit.bind(null, record),
                disabled: false,
                auth: 'admin:sysApi:edit',
              },
            ]"
          />
        </template>
      </template>
    </BasicTable>
    <EditModal @register="registerModal" @success="handleSuccess" />
  </div>
</template>
<script lang="ts" setup name="SysApiManage">
  import { useI18n } from '/@/hooks/web/useI18n';
  import { BasicTable, useTable, TableAction } from '/@/components/Table';
  import { useModal } from '/@/components/Modal';
  import EditModal from './EditModal.vue';
  import { columns, searchFormSchema } from './data';
  import { getSysApiByKey, getSysApiList } from '/@/api/admin/sys-api';

  const { t } = useI18n();
  const [registerModal, { openModal }] = useModal();
  const [registerTable, { reload, updateTableDataRecord }] = useTable({
    title: 'API列表',
    api: getSysApiList,
    columns,
    formConfig: {
      labelWidth: 100,
      baseColProps: { span: 6 },
      schemas: searchFormSchema,
    },
    striped: true,
    bordered: true,
    canResize: true,
    useSearchForm: true,
    showTableSetting: true,
    showIndexColumn: false,
    rowKey: 'id',
    actionColumn: {
      width: 60,
      title: t('common.actionText'),
      dataIndex: 'action',
    },
  });

  // function handleCreate() {
  //   openModal(true, {
  //     isUpdate: false,
  //   });
  // }

  function handleEdit(record: Recordable) {
    openModal(true, {
      record,
      isUpdate: true,
    });
  }

  async function handleSuccess(isUpdate: boolean, record: Recordable) {
    if (isUpdate) {
      const data = await getSysApiByKey(record.id);
      updateTableDataRecord(record.id, data);
    } else {
      await reload();
    }
  }
</script>
