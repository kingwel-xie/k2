import { BasicColumn, FormSchema } from '/@/components/Table';
import { DescItem } from '/@/components/Description';

export const columns: BasicColumn[] = [
  {
    title: '代码',
    dataIndex: 'code',
  },
  {
    title: '三字符代码',
    dataIndex: 'code2',
  },
  {
    title: '数字代码',
    dataIndex: 'digitCode',
  },
  {
    title: '电话代码',
    dataIndex: 'teleCode',
  },
  {
    title: '分组',
    dataIndex: 'group',
  },
  {
    title: '从属',
    dataIndex: 'belongTo',
  },
  {
    title: '中文简称',
    dataIndex: 'nameCN',
  },
  {
    title: '英文简称',
    dataIndex: 'nameEN',
  },
  {
    title: '显示排序',
    dataIndex: 'displaySort',
  },
  {
    title: '描述',
    dataIndex: 'remark',
  },
];

export const searchFormSchema: FormSchema[] = [
  {
    field: 'code',
    label: '代码',
    component: 'Input',
  },
  {
    field: 'code2',
    label: '三字符代码',
    component: 'Input',
  },
  {
    field: 'digitCode',
    label: '数字代码',
    component: 'Input',
  },
  {
    field: 'teleCode',
    label: '电话代码',
    component: 'Input',
  },
  {
    field: 'group',
    label: '分组',
    component: 'Input',
  },
  {
    field: 'nameCN',
    label: '中文简称',
    component: 'Input',
  },
  {
    field: 'nameEN',
    label: '英文简称',
    component: 'Input',
  },
  {
    field: 'remark',
    label: '描述',
    component: 'Input',
  },
];

export const formSchema: FormSchema[] = [
  {
    field: 'code',
    label: '代码',
    component: 'Input',
    required: true,
  },
  {
    field: 'code2',
    label: '三字符代码',
    component: 'Input',
    required: true,
  },
  {
    field: 'digitCode',
    label: '数字代码',
    component: 'Input',
    required: true,
  },
  {
    field: 'teleCode',
    label: '电话代码',
    component: 'Input',
    required: true,
  },
  {
    field: 'group',
    label: '分组',
    component: 'Input',
    required: true,
  },
  {
    field: 'belongTo',
    label: '从属',
    component: 'Input',
  },
  {
    field: 'nameCN',
    label: '中文简称',
    component: 'Input',
    required: true,
  },
  {
    field: 'nameEN',
    label: '英文简称',
    helpMessage: 'Country name in English',
    component: 'Input',
    required: true,
  },
  {
    field: 'displaySort',
    label: '显示排序',
    component: 'InputNumber',
  },
  {
    field: 'remark',
    label: '描述',
    component: 'InputTextArea',
  },
];

export const descSchema: DescItem[] = [
  {
    label: '代码',
    field: 'code',
  },
  {
    label: '三字符代码',
    field: 'code2',
  },
  {
    label: '数字代码',
    field: 'digitCode',
  },
  {
    label: '电话代码',
    field: 'teleCode',
  },
  {
    label: '分组',
    field: 'group',
  },
  {
    label: '从属',
    field: 'belongTo',
  },
  {
    label: '中文简称',
    field: 'nameCN',
  },
  {
    label: '英文简称',
    field: 'nameEN',
  },
  {
    label: '显示排序',
    field: 'displaySort',
  },
  {
    label: '描述',
    field: 'remark',
  },
];

export const excelHeader = {
  code: '代码',
  code2: '三字符代码',
  digitCode: '数字代码',
  teleCode: '电话代码',
  group: '分组',
  belongTo: '从属',
  nameCN: '中文简称',
  nameEN: '英文简称',
  displaySort: '显示排序',
  remark: '描述',
};
