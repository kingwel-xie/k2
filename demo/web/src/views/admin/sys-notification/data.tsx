import { BasicColumn, FormSchema } from '/@/components/Table';
import { DescItem } from '/@/components/Description';
import { tryParseJson } from '/@/utils/formatUtil';

function formatMultiLabel(data: string, separator = ', ') {
  return (tryParseJson(data) || []).join(separator);
}

interface TargetTypeOptionEntry {
  label: string;
  value: string | number;
}

export interface TargetTypeOptionData {
  sysUsers: TargetTypeOptionEntry[];
  sysRoles: TargetTypeOptionEntry[];
  sysDeptList: TargetTypeOptionEntry[];
}

export const columns: BasicColumn[] = [
  {
    title: 'ID',
    dataIndex: 'id',
  },
  {
    title: '接收人类别',
    dataIndex: 'targetType',
    format: 'dict|sys_notification_target_type',
  },
  {
    title: '接收人',
    dataIndex: 'targets',
    customRender: ({ record }) => {
      switch (record.targetType) {
        case 'all':
          return '-';
        case 'role':
        case 'user':
        case 'dept':
          return formatMultiLabel(record.targets);
        default:
          return '?';
      }
    },
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
    title: '备注',
    dataIndex: 'remark',
  },
  {
    title: '创建时间',
    dataIndex: 'createdAt',
    format: 'datetime|flex',
  },
];

export const searchFormSchema: FormSchema[] = [
  {
    field: 'targetType',
    label: '接收人类别',
    component: 'DictSelect',
    componentProps: { dictName: 'sys_notification_target_type' },
  },
  {
    field: 'targets',
    label: '接收人',
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
];

export const formSchema: FormSchema[] = [
  {
    field: 'id',
    label: 'id',
    component: 'InputNumber',
    show: false,
  },
  {
    field: 'targetType',
    label: '接收人类别',
    component: 'DictSelect',
    componentProps: ({ formModel }) => {
      return {
        showSearch: false,
        dictName: 'sys_notification_target_type',
        onChange: (_val) => {
          // clear targets when target type is changed
          formModel['targets'] = undefined;
        },
      };
    },
    required: true,
  },
  {
    field: 'targets',
    slot: 'targets',
    label: '接收人',
    helpMessage: '接收人，由接收人类别决定：角色/用户/部门',
    component: 'Select',
    required: true,
    ifShow: ({ values }) => ['role', 'dept', 'user'].includes(values.targetType),
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
    label: '内容',
    component: 'InputTextArea',
    componentProps: {
      showCount: true,
      maxlength: 511,
      rows: 4,
    },
    required: true,
  },
  {
    field: 'remark',
    label: '备注',
    component: 'InputTextArea',
    colProps: { span: 24 },
  },
];

export const descSchema: DescItem[] = [
  {
    label: 'ID',
    field: 'id',
  },
  {
    label: '接收人类别',
    field: 'targetType',
    format: 'dict|sys_notification_target_type',
  },
  {
    label: '接收人',
    field: 'targets',
  },
  {
    label: '标题',
    field: 'title',
  },
  {
    label: '内容',
    field: 'content',
  },
  {
    label: '备注',
    field: 'remark',
  },
];

export const excelHeader = {
  id: 'ID',
  targetType: '接收人类别',
  targets: '接收人',
  title: '内容',
  content: '内容',
  remark: '备注',
};
