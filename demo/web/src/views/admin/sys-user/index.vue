<template>
  <PageWrapper dense contentFullHeight contentClass="flex">
    <DeptTree class="w-1/4 xl:w-1/5" @select="handleSelect" />
    <BasicTable @register="registerTable" class="w-3/4 xl:w-4/5" :searchInfo="searchInfo">
      <template #toolbar>
        <a-button
          v-auth="'admin:sysUser:add'"
          type="primary"
          preIcon="ant-design:user-add-outlined"
          @click="handleCreate"
        >
          {{ t('common.addText') }}
        </a-button>
        <Dropdown
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
        <template v-if="column.key === 'action'">
          <TableAction
            :actions="[
              // {
              //   icon: 'ant-design:info-circle-outlined',
              //   tooltip: t('common.detailText'),
              //   onClick: handleView.bind(null, record),
              // },
              {
                icon: 'ant-design:form-outlined',
                tooltip: t('common.editText'),
                onClick: handleEdit.bind(null, record),
                auth: 'admin:sysUser:edit',
              },
              {
                icon: 'ant-design:key-outlined',
                tooltip: t('common.resetPwdText'),
                onClick: handleResetPwd.bind(null, record),
                auth: 'admin:sysUser:resetPassword',
              },
              {
                icon: 'ant-design:delete-outlined',
                tooltip: t('common.delText'),
                color: 'error',
                auth: 'admin:sysUser:remove',
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
    <ModalInput ref="passwordInputRef" title="请输入新密码" label="密码" inputType="Input" />
  </PageWrapper>
</template>
<script lang="ts" setup name="SysUserManage">
  import { reactive, ref, Ref, unref } from 'vue';
  import { computed } from 'vue';
  import { cloneDeep } from 'lodash-es';
  import { BasicTable, useTable, TableAction } from '/@/components/Table';
  import { Dropdown, DropMenu } from '/@/components/Dropdown';
  import { useMessage } from '/@/hooks/web/useMessage';
  import { useModal } from '/@/components/Modal';
  import { PageWrapper } from '/@/components/Page';
  import { useGo } from '/@/hooks/web/usePage';
  import { useI18n } from '/@/hooks/web/useI18n';
  import { export2ExcelV2 } from '/@/utils/export2';
  import { ModalInput, PromptModalAction } from '/@/components/ModalInput';
  import DeptTree from './DeptTree.vue';
  import EditModal from './EditModal.vue';
  import { columns, excelHeader, searchFormSchema } from './data';
  import {
    deleteAccountEntry,
    getAccountByKey,
    getAccountList,
    resetUserPwd,
  } from '/@/api/admin/system';

  const { t } = useI18n();
  const go = useGo();
  const { createMessage } = useMessage();
  const [registerModal, { openModal }] = useModal();
  const searchInfo = reactive<Recordable>({});
  const passwordInputRef: Ref<Nullable<PromptModalAction>> = ref(null);

  const [
    registerTable,
    { reload, updateTableDataRecord, getDataSource, getSelectRows, fetchOnly },
  ] = useTable({
    name: 'SysUserManage.MainTable',
    title: '用户账号列表',
    api: getAccountList,
    rowKey: 'userId',
    columns,
    formConfig: {
      labelWidth: 100,
      baseColProps: { span: 6 },
      schemas: searchFormSchema,
      autoSubmitOnEnter: true,
    },
    useSearchForm: true,
    showTableSetting: true,
    bordered: true,
    handleSearchInfoFn(info) {
      console.log('handleSearchInfoFn', info);
      return info;
    },
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

  function handleResetPwd(record: Recordable) {
    unref(passwordInputRef)?.openModal(async (v) => {
      await resetUserPwd(record.userId, v.input);
      createMessage.success('密码重置成功');
    });
  }

  async function handleDelete(record: Recordable) {
    await deleteAccountEntry(record);
    await reload();
  }

  async function handleSuccess(isUpdate: boolean, record: Recordable) {
    if (isUpdate) {
      // 演示不刷新表格直接更新内部数据。
      // 注意：updateTableDataRecord要求表格的rowKey属性为string并且存在于每一行的record的keys中
      const data = await getAccountByKey(record.userId);
      updateTableDataRecord(record.userId, data);
    } else {
      await reload();
    }
  }

  function handleSelect(deptId = '') {
    searchInfo.deptId = deptId;
    reload();
  }

  function _handleView(record: Recordable) {
    go('/admin/account_detail/' + record.username);
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
    export2ExcelV2(excelHeader, columns, data, '用户账号.xlsx');
  }
</script>
