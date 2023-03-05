<template>
  <div>
    <BasicTable @register="registerTable">
      <template #bodyCell="{ column, record }">
        <template v-if="column.key === 'action'">
          <TableAction
            :actions="[
              {
                icon: 'ant-design:delete-outlined',
                tooltip: t('common.delText'),
                color: 'error',
                disabled: false,
                auth: 'admin:sysLoginLog:remove',
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
  </div>
</template>
<script lang="ts" setup name="SysLoginLogManage">
  import { useI18n } from '/@/hooks/web/useI18n';
  import { BasicTable, useTable, TableAction } from '/@/components/Table';
  import { columns, searchFormSchema } from './data';
  import { deleteSysLoginLogEntry, getSysLoginLogList } from '/@/api/admin/sys-login-log';

  const { t } = useI18n();

  const [registerTable, { reload }] = useTable({
    name: 'SysLoginLogManage.MainTable',
    title: '列表',
    api: getSysLoginLogList,
    columns,
    formConfig: {
      labelWidth: 100,
      baseColProps: { span: 5 },
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
      width: 120,
      title: t('common.actionText'),
      dataIndex: 'action',
    },
  });

  async function handleDelete(record: Recordable) {
    await deleteSysLoginLogEntry(record);
    await reload();
  }
</script>
