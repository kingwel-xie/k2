<template>
  <div>
    <BasicTable @register="registerTable">
      <template #toolbar>
        <a-button
          v-auth="'admin:sysNotification:add'"
          type="primary"
          preIcon="ant-design:file-add-outlined"
          @click="handleCreate"
        >
          {{ t('common.addText') }}
        </a-button>
        <Dropdown
          v-if="false"
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
        <template v-if="column.key === 'id' && false">
          <TableAction
            :actions="[
              {
                icon: 'ant-design:info-circle-outlined',
                label: record.id,
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
              // {
              //   icon: 'ant-design:form-outlined',
              //   tooltip: t('common.editText'),
              //   onClick: handleEdit.bind(null, record),
              //   disabled: false,
              //   auth: 'admin:sysNotification:edit',
              // },
              {
                icon: 'ant-design:delete-outlined',
                tooltip: t('common.delText'),
                color: 'error',
                disabled: false,
                auth: 'admin:sysNotification:remove',
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
<script lang="ts" setup name="SysNotificationManage">
  import { useI18n } from '/@/hooks/web/useI18n';
  import { BasicTable, useTable, TableAction } from '/@/components/Table';
  import { computed, onMounted, reactive } from 'vue';
  import { useModal } from '/@/components/Modal';
  import { useDrawer } from '/@/components/Drawer';
  import { Dropdown, DropMenu } from '/@/components/Dropdown';
  import { cloneDeep } from 'lodash-es';
  import { export2ExcelV2 } from '/@/utils/export2';
  import { columns, excelHeader, searchFormSchema, TargetTypeOptionData } from './data';
  import EditModal from './EditModal.vue';
  import DetailDrawer from './DetailDrawer.vue';
  import {
    deleteSysNotificationEntry,
    getSysNotificationList,
    getSysNotificationByKey,
  } from '/@/api/admin/sys-notification';
  import { listAccountNoCheck, listDeptNoCheck, listRoleNoCheck } from '/@/api/admin/system';
  import { tryParseJson } from '/@/utils/formatUtil';

  const { t } = useI18n();

  const optionData = reactive<TargetTypeOptionData>({
    sysUsers: [],
    sysRoles: [],
    sysDeptList: [],
  });

  onMounted(async () => {
    optionData.sysUsers = (await listAccountNoCheck()).list.map((x) => ({
      label: x.username,
      value: x.username,
    }));
    optionData.sysRoles = (await listRoleNoCheck()).list.map((x) => ({
      label: x.roleName,
      value: x.roleKey,
    }));
    optionData.sysDeptList = (await listDeptNoCheck()).list.map((x) => ({
      label: x.deptName,
      value: x.deptId,
    }));
  });

  const clonedColumns = cloneDeep(columns);
  clonedColumns[2].customRender = ({ record }) => {
    switch (record.targetType) {
      case 'all':
        return '-';
      case 'role':
        const roles = tryParseJson(record.targets) || [];
        return roles
          .map((role) => optionData.sysRoles.find((x) => x.value === role)?.label || '?')
          .join(', ');
      case 'user':
        const users = tryParseJson(record.targets) || [];
        return users.join(', ');
      case 'dept':
        const deptList = tryParseJson(record.targets) || [];
        return deptList
          .map((dept) => optionData.sysDeptList.find((x) => x.value === dept)?.label || '?')
          .join(', ');
      default:
        return '?';
    }
  };

  const [registerDrawer, { openDrawer }] = useDrawer();
  const [registerModal, { openModal }] = useModal();
  const [
    registerTable,
    { reload, getDataSource, getSelectRows, fetchOnly, updateTableDataRecord },
  ] = useTable({
    name: 'SysNotificationManage.MainTable',
    title: '通知列表',
    api: getSysNotificationList,
    columns: clonedColumns,
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
      width: 90,
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
      optionData: optionData,
    });
  }

  // function handleEdit(record: Recordable) {
  //   openModal(true, {
  //     record,
  //     isUpdate: true,
  //   });
  // }

  async function handleDelete(record: Recordable) {
    await deleteSysNotificationEntry(record);
    await reload();
  }

  async function handleSuccess(isUpdate: boolean, record: Recordable) {
    if (isUpdate) {
      const data = await getSysNotificationByKey(record.id);
      updateTableDataRecord(record.id, data);
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
    export2ExcelV2(excelHeader, columns, data, '通知.xlsx');
  }
</script>
