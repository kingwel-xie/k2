<template>
  <div>
    <BasicTable @register="registerTable">
      <template #toolbar>
        <a-button
          v-auth="'kobh:tbxCountry:add'"
          type="primary"
          preIcon="ant-design:file-add-outlined"
          @click="handleCreate"
        >
          {{ t('common.addText') }}
        </a-button>
        <Dropdown
          v-if="true"
          :dropMenuList="[
            {
              icon: 'ant-design:select-outlined',
              text: '选中项',
              event: 'selected',
              disabled: !hasSelected,
            },
            {
              icon: 'ant-design:profile-outlined',
              text: '当前列表',
              event: 'current',
            },
            {
              icon: 'ant-design:carry-out-outlined',
              text: '满足条件全部',
              event: 'all',
              popConfirm: {
                title: '是否导出满足当前查询条件的数据项（最多10000项）?',
                confirm: handleExport.bind(null, 'all'),
              },
            },
          ]"
          :trigger="['click']"
          @menu-event="handleMenuEvent"
          popconfirm
        >
          <a-button preIcon="ant-design:export-outlined">{{ t('common.exportText') }}...</a-button>
        </Dropdown>
      </template>
      <template #bodyCell="{ column, record }">
        <template v-if="column.key === 'code' && false">
          <TableAction
            :actions="[
              {
                icon: 'ant-design:info-circle-outlined',
                label: record.code,
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
                ifShow: false,
              },
              {
                icon: 'ant-design:form-outlined',
                tooltip: t('common.editText'),
                onClick: handleEdit.bind(null, record),
                disabled: false,
                auth: 'kobh:tbxCountry:edit',
              },
              {
                icon: 'ant-design:delete-outlined',
                tooltip: t('common.delText'),
                color: 'error',
                disabled: false,
                auth: 'kobh:tbxCountry:remove',
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
    <DetailDrawer @register="registerDrawer" />
  </div>
</template>
<script lang="ts" setup name="TbxCountry">
  import { useI18n } from '/@/hooks/web/useI18n';
  import { computed } from 'vue';
  import { BasicTable, useTable, TableAction } from '/@/components/Table';
  import { useModal } from '/@/components/Modal';
  import { useDrawer } from '/@/components/Drawer';
  import { Dropdown, DropMenu } from '/@/components/Dropdown';
  import { cloneDeep } from 'lodash-es';
  import { export2Excel } from '/@/utils/export';
  import { columns, excelHeader, searchFormSchema } from './data';
  import EditModal from './EditModal.vue';
  import DetailDrawer from './DetailDrawer.vue';
  import {
    deleteTbxCountryEntry,
    getTbxCountryList,
    getTbxCountryByKey,
  } from '/@/api/kobh/tbx-country';

  const { t } = useI18n();
  const [registerDrawer, { openDrawer }] = useDrawer();
  const [registerModal, { openModal }] = useModal();
  const [
    registerTable,
    { reload, getDataSource, getSelectRows, fetchOnly, updateTableDataRecord },
  ] = useTable({
    title: '国家编码列表',
    api: getTbxCountryList,
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
    rowKey: 'code',
    actionColumn: {
      width: 120,
      title: t('common.actionText'),
      dataIndex: 'action',
    },
  });

  const hasSelected = computed<boolean>(() => {
    try {
      return getSelectRows().length > 0;
    } catch {
      return false;
    }
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
    await deleteTbxCountryEntry(record);
    await reload();
  }

  async function handleSuccess(isUpdate: boolean, record: Recordable) {
    if (isUpdate) {
      const data = await getTbxCountryByKey(record.code);
      updateTableDataRecord(record.code, data);
    } else {
      await reload();
    }
  }

  function handleView(record: Recordable) {
    openDrawer(true, {
      record,
    });
  }

  function handleMenuEvent(menu: DropMenu) {
    if (menu.event === 'selected' || menu.event === 'current') {
      handleExport(menu.event);
    }
  }

  async function handleExport(event) {
    let data = [];
    switch (event) {
      case 'selected':
        data = cloneDeep(getSelectRows());
        break;
      case 'current':
        data = cloneDeep(getDataSource());
        break;
      case 'all':
        data = (await fetchOnly()) || [];
        break;
    }
    export2Excel(excelHeader, columns, data, '国家编码.xlsx');
  }
</script>
