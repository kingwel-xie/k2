import dayjs from 'dayjs';
import { BasicColumn, FormSchema } from '/@/components/Table';
import { listAccountNoCheck } from '/@/api/admin/system';

export const columns: BasicColumn[] = [
  {
    title: 'ID',
    dataIndex: 'id',
    fixed: 'left',
  },
  {
    title: '类型',
    dataIndex: 'type',
    format: 'dict|sys_notice_type',
  },
  {
    title: '发件人',
    dataIndex: 'sender',
  },
  {
    title: '收件人',
    dataIndex: 'receiver',
  },
  {
    title: '标题',
    dataIndex: 'title',
  },
  {
    title: '内容',
    dataIndex: 'content',
  },
  {
    title: '已读',
    dataIndex: 'read',
    format: 'bool|✔,',
  },
  {
    title: '时间',
    dataIndex: 'createdAt',
    format: 'datetime|flex',
  },
];

export const searchFormSchema: FormSchema[] = [
  {
    field: 'type',
    label: '类型',
    component: 'DictSelect',
    componentProps: { dictName: 'sys_notice_type' },
  },
  {
    field: 'sender',
    label: '发件人',
    component: 'Input',
  },
  {
    field: 'receiver',
    label: '收件人',
    component: 'Input',
  },
  {
    field: 'title',
    label: '标题',
    component: 'Input',
  },
  {
    field: 'content',
    label: '内容',
    component: 'Input',
  },
  {
    field: '[beginTime, endTime]',
    label: '时间范围',
    component: 'RangePicker',
    componentProps: {
      showTime: {
        defaultValue: [dayjs('00:00:00', 'HH:mm:ss'), dayjs('23:59:59', 'HH:mm:ss')],
      },
    },
    colProps: { span: 12 },
  },
];

export const formSchema: FormSchema[] = [
  {
    field: 'id',
    label: 'id',
    component: 'InputNumber',
    show: false,
  },
  // {
  //   field: 'receiver',
  //   slot: 'receiver',
  //   label: '接收人',
  //   component: 'Select',
  //   required: true,
  // },
  {
    field: 'receiver',
    label: '接收人',
    component: 'ApiSelect',
    componentProps: {
      api: listAccountNoCheck,
      allClear: true,
      mode: 'multiple',
      optionFilterProp: 'label',
      resultField: 'list',
      labelField: 'username',
      valueField: 'username',
      getPopupContainer: () => document.body,
    },
    required: true,
  },
  {
    field: 'title',
    label: '标题',
    component: 'Input',
    componentProps: {
      showCount: true,
      maxlength: 127,
    },
    required: true,
  },
  {
    field: 'content',
    slot: 'content',
    label: '内容',
    component: 'InputTextArea',
    required: true,
  },
];
