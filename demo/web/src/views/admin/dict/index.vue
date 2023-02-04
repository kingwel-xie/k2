<template>
  <div>
    <BasicTable @register="registerTable">
      <template #toolbar>
        <a-button
          v-auth="'admin:sysDictType:add'"
          preIcon="ant-design:file-add-outlined"
          type="primary"
          @click="handleCreate"
          >{{ t('common.addText') }}</a-button
        >
      </template>
      <template #bodyCell="{ column, record }">
        <template v-if="column.key === 'dictType'">
          <TableAction
            :actions="[
              {
                icon: 'ant-design:info-circle-outlined',
                label: record.dictType,
                tooltip: t('common.detailText'),
                onClick: handleView.bind(null, record),
              },
            ]"
          />
        </template>
        <template v-if="column.key === 'action'">
          <TableAction
            :actions="[
              {
                icon: 'ant-design:info-circle-outlined',
                tooltip: t('common.detailText'),
                onClick: handleView.bind(null, record),
              },
              {
                icon: 'ant-design:form-outlined',
                tooltip: t('common.editText'),
                onClick: handleEdit.bind(null, record),
                disabled: false,
                auth: 'admin:sysDictType:edit',
              },
              {
                icon: 'ant-design:delete-outlined',
                tooltip: t('common.delText'),
                color: 'error',
                disabled: false,
                auth: 'admin:sysDictType:remove',
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
<script lang="ts" setup name="SysDictTypeManage">
  import { useI18n } from '/@/hooks/web/useI18n';
  import { BasicTable, useTable, TableAction } from '/@/components/Table';
  import { useModal } from '/@/components/Modal';
  import EditModal from './EditModal.vue';
  import { columns, searchFormSchema } from './typeData';
  import {
    deleteSysDictTypeEntry,
    getSysDictTypeByKey,
    getSysDictTypeList,
  } from '/@/api/admin/sys-dict';
  import { useGo } from '/@/hooks/web/usePage';

  const { t } = useI18n();
  const go = useGo();

  // const [registerDrawer, { openDrawer }] = useDrawer();
  const [registerModal, { openModal }] = useModal();
  const [registerTable, { reload, updateTableDataRecord }] = useTable({
    title: '字典类型列表',
    api: getSysDictTypeList,
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
    await deleteSysDictTypeEntry(record);
    await reload();
  }

  async function handleSuccess(isUpdate: boolean, record: Recordable) {
    if (isUpdate) {
      const data = await getSysDictTypeByKey(record.id);
      updateTableDataRecord(record.id, data);
    } else {
      await reload();
    }
  }

  function handleView(record: Recordable) {
    go('dict/data/' + record.dictType);
  }
</script>
