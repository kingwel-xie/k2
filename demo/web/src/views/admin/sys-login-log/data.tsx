import { BasicColumn, FormSchema } from '/@/components/Table';
import dayjs from 'dayjs';

export const columns: BasicColumn[] = [
  {
    title: '编号',
    dataIndex: 'id',
    width: 50,
  },
  {
    title: '登录用户',
    dataIndex: 'username',
    width: 70,
  },
  {
    title: '信息',
    dataIndex: 'msg',
    width: 90,
  },
  {
    title: '状态',
    dataIndex: 'status',
    width: 70,
    format: 'dict|sys_normal_disable',
  },
  {
    title: '来源地址',
    dataIndex: 'ipaddr',
    width: 90,
  },
  {
    title: 'OS',
    dataIndex: 'os',
    width: 80,
  },
  {
    title: '浏览器',
    dataIndex: 'browser',
    width: 90,
  },
  {
    title: '登录时间',
    dataIndex: 'loginTime',
    width: 180,
    format: 'datetime|full',
  },
  {
    title: '备注',
    dataIndex: 'remark',
    defaultHidden: true,
  },
];

export const searchFormSchema: FormSchema[] = [
  {
    field: 'username',
    label: '登录用户',
    component: 'Input',
  },
  {
    field: 'ipaddr',
    label: '来源地址',
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
  {
    field: '[beginTime, endTime]',
    label: '登录时间',
    component: 'RangePicker',
    componentProps: {
      showTime: {
        defaultValue: [dayjs('00:00:00', 'HH:mm:ss'), dayjs('23:59:59', 'HH:mm:ss')],
      },
    },
    colProps: { span: 9 },
  },
];
