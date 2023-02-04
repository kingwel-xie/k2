import { BasicColumn } from '/@/components/Table';
import { FormSchema } from '/@/components/Table';

export const columns: BasicColumn[] = [
  {
    title: '部门名称',
    dataIndex: 'deptName',
    width: 160,
    align: 'left',
  },
  {
    title: '排序',
    dataIndex: 'sort',
    width: 60,
  },
  {
    title: '负责人',
    dataIndex: 'leader',
    width: 100,
  },
  {
    title: '联系电话',
    dataIndex: 'phone',
    width: 160,
  },
  {
    title: '邮箱',
    dataIndex: 'email',
    width: 160,
  },
  {
    title: '状态',
    dataIndex: 'status',
    width: 100,
    format: 'dict|sys_normal_disable',
  },
  {
    title: '创建时间',
    dataIndex: 'createdAt',
    width: 180,
    format: 'datetime|flex',
  },
  {
    title: '更新时间',
    dataIndex: 'updatedAt',
    width: 180,
    format: 'datetime|flex',
  },
  {
    title: '备注',
    dataIndex: 'remark',
  },
];

export const searchFormSchema: FormSchema[] = [
  {
    field: 'deptName',
    label: '部门名称',
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
    field: 'deptId',
    label: 'id',
    component: 'InputNumber',
    show: false,
  },
  {
    field: 'deptName',
    label: '部门名称',
    component: 'Input',
    required: true,
  },
  {
    field: 'parentId',
    label: '上级部门',
    component: 'TreeSelect',
    componentProps: {
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
    field: 'sort',
    label: '排序',
    component: 'InputNumber',
    required: true,
  },
  {
    field: 'leader',
    label: '负责人',
    component: 'Input',
    required: true,
  },
  {
    label: '邮箱',
    field: 'email',
    component: 'Input',
  },
  {
    label: '联系电话',
    field: 'phone',
    component: 'Input',
  },
  {
    label: '状态',
    field: 'status',
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
];
