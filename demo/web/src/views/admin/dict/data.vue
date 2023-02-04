<template>
  <div>
    <BasicTable @register="registerTable">
      <template #toolbar>
        <a-button
          v-auth="'admin:sysDictData:add'"
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
                disabled: false,
                auth: 'admin:sysDictData:edit',
              },
              {
                icon: 'ant-design:delete-outlined',
                tooltip: t('common.delText'),
                color: 'error',
                disabled: false,
                auth: 'admin:sysDictData:remove',
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
<script lang="ts" setup name="SysDictDataManage">
  import { useI18n } from '/@/hooks/web/useI18n';
  import { BasicTable, useTable, TableAction } from '/@/components/Table';
  import { useModal } from '/@/components/Modal';
  import EditModal from './DataEditModal.vue';
  import { columns, searchFormSchema } from './dataData';
  import {
    deleteSysDictDataEntry,
    getSysDictDataByKey,
    getSysDictDataList,
  } from '/@/api/admin/sys-dict';
  import { useRoute } from 'vue-router';

  const { t } = useI18n();
  const route = useRoute();

  // 此处可以得到dictType, historical reason - dictId is in the URL
  const defaultDictType = route.params?.dictId;
  // console.log(route.params);

  const [registerModal, { openModal }] = useModal();
  const [registerTable, { reload, updateTableDataRecord }] = useTable({
    title: '字典数据列表',
    api: getSysDictDataList,
    columns,
    formConfig: {
      labelWidth: 120,
      baseColProps: { span: 6 },
      schemas: searchFormSchema,
    },
    striped: true,
    bordered: true,
    canResize: true,
    useSearchForm: true,
    showTableSetting: true,
    showIndexColumn: false,
    rowKey: 'dictCode',
    beforeFetch(params) {
      params.dictType = defaultDictType;
      return params;
    },
    actionColumn: {
      width: 80,
      title: t('common.actionText'),
      dataIndex: 'action',
    },
  });

  function handleCreate() {
    openModal(true, {
      isUpdate: false,
      dictType: defaultDictType,
    });
  }

  function handleEdit(record: Recordable) {
    openModal(true, {
      record,
      isUpdate: true,
    });
  }

  async function handleDelete(record: Recordable) {
    await deleteSysDictDataEntry(record);
    await reload();
  }

  async function handleSuccess(isUpdate: boolean, record: Recordable) {
    if (isUpdate) {
      const data = await getSysDictDataByKey(record.dictCode);
      updateTableDataRecord(record.dictCode, data);
    } else {
      await reload();
    }
  }
</script>
