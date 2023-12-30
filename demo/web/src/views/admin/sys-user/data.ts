import { listRoleNoCheck, getDeptTree, setAccountStatus } from '/@/api/admin/system';
import { BasicColumn } from '/@/components/Table';
import { FormSchema } from '/@/components/Table';
import { h } from 'vue';
import { Switch } from 'ant-design-vue';
import { useMessage } from '/@/hooks/web/useMessage';

export const columns: BasicColumn[] = [
  {
    title: '用户名',
    dataIndex: 'username',
    fixed: 'left',
    width: 100,
    sorter: true,
  },
  {
    title: '昵称',
    dataIndex: 'nickName',
    width: 120,
  },
  {
    title: '手机',
    dataIndex: 'phone',
    width: 110,
  },
  {
    title: '邮箱',
    dataIndex: 'email',
    width: 140,
  },
  {
    title: '部门',
    //dataIndex: ['dept', 'deptName'],
    dataIndex: 'deptId',
    width: 120,
    sorter: true,
    customRender: ({ record }) => {
      return record.dept?.deptName || '-';
    },
  },
  {
    title: '角色',
    // dataIndex: ['role', 'roleName'],
    dataIndex: 'roleId',
    width: 110,
    sorter: true,
    customRender: ({ record }) => {
      return record.role?.roleName || '-';
    },
  },
  {
    title: 'API Token',
    dataIndex: 'token',
    width: 210,
  },
  {
    title: '状态',
    dataIndex: 'status',
    width: 100,
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
          setAccountStatus(record.userId, newStatus)
            .then(() => {
              record.status = newStatus;
              createMessage.success(`已成功修改状态`);
            })
            .catch(() => {
              createMessage.error('修改状态失败');
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
    width: 140,
    sorter: true,
  },
  {
    title: '备注',
    dataIndex: 'remark',
    width: 110,
  },
];

export const searchFormSchema: FormSchema[] = [
  {
    field: 'username',
    label: '用户名',
    component: 'Input',
  },
  {
    field: 'nickName',
    label: '昵称',
    component: 'Input',
  },
  {
    field: 'roleKey',
    label: '角色',
    component: 'ApiSelect',
    componentProps: {
      api: listRoleNoCheck,
      resultField: 'list',
      labelField: 'roleName',
      valueField: 'roleKey',
      mode: 'multiple',
      getPopupContainer: () => document.body,
    },
  },
];

export const accountFormSchema: FormSchema[] = [
  {
    field: 'userId',
    label: 'userId',
    component: 'InputNumber',
    show: false,
  },
  {
    field: 'username',
    label: '用户名',
    component: 'Input',
    helpMessage: '不能输入已被占用的用户名',
  },
  {
    field: 'nickName',
    label: '昵称',
    component: 'Input',
    required: true,
  },
  {
    field: 'password',
    label: '密码',
    component: 'InputPassword',
    required: true,
    show: false,
  },
  {
    label: '角色',
    field: 'roleId',
    component: 'ApiSelect',
    componentProps: {
      api: listRoleNoCheck,
      resultField: 'list',
      labelField: 'roleName',
      valueField: 'roleId',
      getPopupContainer: () => document.body,
    },
    required: true,
  },
  {
    label: '所属部门',
    field: 'deptId',
    component: 'ApiTreeSelect',
    componentProps: {
      api: getDeptTree,
      fieldNames: {
        label: 'label',
        key: 'id',
        value: 'id',
      },
      getPopupContainer: () => document.body,
    },
    required: true,
  },
  {
    label: '邮箱',
    field: 'email',
    component: 'Input',
    required: true,
  },
  {
    label: '手机',
    field: 'phone',
    component: 'Input',
    required: true,
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
    label: 'API Token',
    field: 'token',
    component: 'InputTextArea',
    dynamicDisabled: true,
    componentProps: { rows: 3 },
    colProps: { span: 18 },
    ifShow: ({ values }) => values.userId > 0,
  },
  {
    label: ' ',
    labelWidth: '8px',
    field: 'tokenEnabled',
    slot: 'tokenEnabled',
    component: 'Checkbox',
    colProps: { span: 6 },
    ifShow: ({ values }) => values.userId > 0,
  },
];

export const excelHeader = {
  username: '用户名',
  nickName: '昵称',
  phone: '手机',
  email: '邮箱',
  status: '状态',
  createdAt: '创建时间',
};
