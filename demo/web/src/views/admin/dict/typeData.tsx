import { BasicColumn, FormSchema } from '/@/components/Table';

export const columns: BasicColumn[] = [
  {
    title: '字典编号',
    dataIndex: 'id',
    width: 50,
    sorter: true,
  },
  {
    title: '字典名称',
    dataIndex: 'dictName',
    width: 90,
  },
  {
    title: '字典类型',
    dataIndex: 'dictType',
    width: 90,
  },
  {
    title: '状态',
    dataIndex: 'status',
    width: 70,
    format: 'dict|sys_normal_disable',
  },
  {
    title: '创建时间',
    dataIndex: 'createdAt',
    width: 180,
    format: 'datetime|flex',
  },
  {
    title: '备注',
    dataIndex: 'remark',
    defaultHidden: true,
  },
];

export const searchFormSchema: FormSchema[] = [
  {
    field: 'dictName',
    label: '字典名称',
    component: 'Input',
  },
  {
    field: 'dictType',
    label: '字典类型',
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
    field: 'id',
    label: 'id',
    component: 'InputNumber',
    show: false,
  },
  {
    field: 'dictName',
    label: '字典名称',
    component: 'Input',
    required: true,
  },
  {
    field: 'dictType',
    label: '字典类型',
    component: 'Input',
    required: true,
  },
  {
    field: 'status',
    label: '状态',
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
];
