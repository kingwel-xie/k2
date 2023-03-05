import { BasicColumn } from '/@/components/Table';
import { FormSchema } from '/@/components/Table';
import { h } from 'vue';
import { Switch } from 'ant-design-vue';
import { setRoleStatus } from '/@/api/admin/system';
import { useMessage } from '/@/hooks/web/useMessage';

export const columns: BasicColumn[] = [
  {
    title: '角色名称',
    dataIndex: 'roleName',
    width: 200,
  },
  {
    title: '角色值',
    dataIndex: 'roleKey',
    width: 180,
  },
  {
    title: '排序',
    dataIndex: 'roleSort',
    width: 70,
  },
  {
    title: '状态',
    dataIndex: 'status',
    width: 120,
    format: 'dict|sys_normal_disable',
    customRender: ({ record }) => {
      if (!Reflect.has(record, 'pendingStatus')) {
        record.pendingStatus = false;
      }
      return h(Switch, {
        checked: record.status === '2',
        checkedChildren: '启用',
        unCheckedChildren: '禁用',
        loading: record.pendingStatus,
        onChange(checked: boolean) {
          record.pendingStatus = true;
          const newStatus = checked ? '2' : '1';
          const { createMessage } = useMessage();
          setRoleStatus(record.roleId, newStatus)
            .then(() => {
              record.status = newStatus;
              createMessage.success(`已成功修改角色状态`);
            })
            .catch(() => {
              createMessage.error('修改角色状态失败');
            })
            .finally(() => {
              record.pendingStatus = false;
            });
        },
      });
    },
  },
  {
    title: '创建时间',
    dataIndex: 'createdAt',
    format: 'datetime|flex',
    width: 180,
  },
  {
    title: '备注',
    dataIndex: 'remark',
  },
];

export const searchFormSchema: FormSchema[] = [
  {
    field: 'roleName',
    label: '角色名称',
    component: 'Input',
  },
  {
    field: 'status',
    label: '状态',
    component: 'DictSelect',
    componentProps: {
      dictName: 'sys_normal_disable',
    },
  },
];

export const formSchema: FormSchema[] = [
  {
    field: 'roleId',
    label: 'id',
    component: 'InputNumber',
    show: false,
  },
  {
    field: 'roleName',
    label: '角色名称',
    required: true,
    component: 'Input',
  },
  {
    field: 'roleKey',
    label: '角色值',
    required: true,
    component: 'Input',
  },
  {
    label: '状态',
    field: 'status',
    defaultValue: '2',
    component: 'DictRadioGroup',
    componentProps: {
      dictName: 'sys_normal_disable',
      isBtn: true,
    },
    required: true,
  },
  {
    label: '备注',
    field: 'remark',
    component: 'InputTextArea',
  },
  {
    label: '',
    field: 'menu',
    slot: 'menu',
    component: 'TreeSelect',
  },
];
