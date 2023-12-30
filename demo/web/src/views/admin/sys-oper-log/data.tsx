import { BasicColumn, FormSchema } from '/@/components/Table';
import { Tag } from 'ant-design-vue';
import dayjs from 'dayjs';

export const columns: BasicColumn[] = [
  {
    title: '编号',
    dataIndex: 'id',
    fixed: 'left',
    width: 50,
  },
  {
    title: '信息',
    dataIndex: 'operUrl',
    width: 200,
  },
  {
    title: '请求方法',
    dataIndex: 'requestMethod',
    width: 60,
  },
  {
    title: 'IP地址',
    dataIndex: 'operIp',
    width: 90,
  },
  {
    title: '状态码',
    dataIndex: 'statusCode',
    width: 50,
  },
  {
    title: 'API Code',
    dataIndex: 'apiCode',
    width: 60,
    customRender: ({ value }) => {
      const color = value === 200 ? 'success' : 'error';
      return <Tag color={color}>{() => value}</Tag>;
    },
  },
  {
    title: '延迟',
    dataIndex: 'latencyTime',
    width: 90,
  },
  {
    title: '操作人员',
    dataIndex: 'operName',
    width: 70,
  },
  {
    title: '操作时间',
    dataIndex: 'operTime',
    width: 150,
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
    field: 'apiCode',
    label: 'API Code',
    component: 'Input',
  },
  {
    field: 'operName',
    label: '操作人员',
    component: 'Input',
  },
  {
    field: '[beginTime, endTime]',
    label: '操作时间',
    component: 'RangePicker',
    componentProps: {
      showTime: {
        defaultValue: [dayjs('00:00:00', 'HH:mm:ss'), dayjs('23:59:59', 'HH:mm:ss')],
      },
    },
    colProps: { span: 12 },
  },
];
