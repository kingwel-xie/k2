<template>
  <div>
    <BasicTable @register="registerTable">
      <template #bodyCell="{ column, record }">
        <template v-if="column.key === 'action'">
          <TableAction
            :actions="[
              {
                icon: 'ant-design:info-circle-outlined',
                tooltip: t('common.detailText'),
                onClick: handleView.bind(null, record),
              },
              {
                icon: 'ant-design:delete-outlined',
                tooltip: t('common.delText'),
                color: 'error',
                disabled: false,
                auth: 'admin:sysOperLog:remove',
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
    <DetailDrawer @register="registerDrawer" />
  </div>
</template>
<script lang="ts" setup name="SysOperLogManage">
  import { useI18n } from '/@/hooks/web/useI18n';
  import { BasicTable, useTable, TableAction } from '/@/components/Table';
  import { columns, searchFormSchema } from './data';
  import { getSysOperLogList, deleteSysOperLogEntry } from '/@/api/admin/sys-oper-log';
  import DetailDrawer from './DetailDrawer.vue';
  import { useDrawer } from '/@/components/Drawer';

  const { t } = useI18n();

  const [registerDrawer, { openDrawer }] = useDrawer();
  const [registerTable, { reload }] = useTable({
    title: '列表',
    api: getSysOperLogList,
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
      width: 120,
      title: t('common.actionText'),
      dataIndex: 'action',
    },
  });

  async function handleDelete(record: Recordable) {
    await deleteSysOperLogEntry(record);
    await reload();
  }

  function handleView(record: Recordable) {
    openDrawer(true, {
      record,
    });
  }
</script>
