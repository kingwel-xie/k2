import { BasicColumn, FormSchema } from '/@/components/Table';

export const columns: BasicColumn[] = [
  {
    title: '编号',
    dataIndex: 'dictCode',
    width: 50,
    sorter: true,
  },
  {
    title: '数据类型',
    dataIndex: 'dictType',
    width: 80,
  },
  {
    title: '数据标签',
    dataIndex: 'dictLabel',
    width: 80,
  },
  {
    title: '数据键值',
    dataIndex: 'dictValue',
    width: 80,
  },
  // {
  //   title: '显示排序',
  //   dataIndex: 'dictSort',
  //   width: 80,
  // },
  {
    title: '状态',
    dataIndex: 'status',
    width: 60,
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
    field: 'dictType',
    label: '字典类型',
    component: 'Input',
    show: false,
  },
  {
    field: 'dictName',
    label: '字典名称',
    component: 'Input',
  },
  {
    field: 'dictLabel',
    label: '字典标签',
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
    field: 'dictCode',
    label: 'dictCode',
    component: 'InputNumber',
    show: false,
  },
  {
    field: 'dictType',
    label: '字典类型',
    component: 'Input',
    dynamicDisabled: true,
  },
  {
    field: 'dictLabel',
    label: '数据标签',
    component: 'Input',
    required: true,
  },
  {
    field: 'dictValue',
    label: '数据键值',
    component: 'Input',
    required: true,
  },
  // {
  //   field: 'dictSort',
  //   label: '显示排序',
  //   component: 'InputNumber',
  // },
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
