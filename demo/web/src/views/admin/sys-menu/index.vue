<template>
  <div>
    <BasicTable @register="registerTable" @fetch-success="onFetchSuccess">
      <template #toolbar>
        <a-button
          v-auth="'admin:sysMenu:add'"
          preIcon="ant-design:file-add-outlined"
          type="primary"
          @click="handleCreate(0)"
          >{{ t('common.addText') }}</a-button
        >
      </template>
      <template #bodyCell="{ column, record }">
        <template v-if="column.key === 'action'">
          <TableAction
            :actions="[
              {
                icon: 'ant-design:file-add-outlined',
                tooltip: t('common.addText'),
                onClick: handleCreate.bind(null, record.menuId),
                auth: 'admin:sysMenu:add',
              },
              {
                icon: 'ant-design:form-outlined',
                color: 'warning',
                tooltip: t('common.editText'),
                onClick: handleEdit.bind(null, record),
                auth: 'admin:sysMenu:edit',
              },
              {
                icon: 'ant-design:delete-outlined',
                tooltip: t('common.delText'),
                color: 'error',
                auth: 'admin:sysMenu:remove',
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
    <EditDrawer @register="registerDrawer" @success="handleSuccess" />
  </div>
</template>
<script lang="ts">
  import { defineComponent } from 'vue';

  import { BasicTable, useTable, TableAction } from '/@/components/Table';
  import { deleteMenuEntry, getMenuList } from '/@/api/admin/system';

  import { useDrawer } from '/@/components/Drawer';
  import EditDrawer from './EditDrawer.vue';

  import { columns, searchFormSchema } from './data';
  import { useI18n } from '/@/hooks/web/useI18n';

  export default defineComponent({
    name: 'SysMenuManage',
    components: { BasicTable, EditDrawer, TableAction },
    setup() {
      const { t } = useI18n();
      const [registerDrawer, { openDrawer }] = useDrawer();
      const [registerTable, { reload, expandAll }] = useTable({
        name: 'SysMenuManage.MainTable',
        title: '菜单列表',
        api: getMenuList,
        columns,
        formConfig: {
          labelWidth: 100,
          baseColProps: { span: 6 },
          schemas: searchFormSchema,
        },
        rowKey: 'menuId',
        isTreeTable: true,
        pagination: false,
        striped: false,
        useSearchForm: true,
        showTableSetting: true,
        bordered: true,
        showIndexColumn: false,
        canResize: false,
        actionColumn: {
          width: 110,
          title: t('common.actionText'),
          dataIndex: 'action',
          fixed: undefined,
        },
      });

      function handleCreate(parentId: number) {
        openDrawer(true, {
          isUpdate: false,
          parentId,
        });
      }

      function handleEdit(record: Recordable) {
        openDrawer(true, {
          record,
          isUpdate: true,
        });
      }

      async function handleDelete(record: Recordable) {
        await deleteMenuEntry(record);
        await reload();
      }

      function handleSuccess() {
        reload();
      }

      function onFetchSuccess() {
        // 演示默认展开所有表项
        // nextTick(expandAll);
      }

      return {
        registerTable,
        registerDrawer,
        handleCreate,
        handleEdit,
        handleDelete,
        handleSuccess,
        onFetchSuccess,
        t,
      };
    },
  });
</script>
