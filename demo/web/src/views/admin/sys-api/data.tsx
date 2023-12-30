import { BasicColumn, FormSchema } from '/@/components/Table';
import { Tag, Tooltip } from 'ant-design-vue';
import { Icon } from '/@/components/Icon';

export const columns: BasicColumn[] = [
  {
    title: '编号',
    dataIndex: 'id',
    fixed: 'left',
    width: 50,
    ifShow: false,
  },
  {
    title: '标题',
    dataIndex: 'title',
    customRender: ({ record }) => {
      const mapping = {
        NOCHECK: ['success', 'ant-design:unlock-outlined'],
        CHECK: ['blue', 'ant-design:lock-outlined'],
      };
      const color = mapping[record.type] || ['default', ''];
      return (
        <div>
          <Tooltip title={<Icon icon={color[1]} size={32} />} color="gray">
            <Tag color={color[0]}>{() => record.title}</Tag>
          </Tooltip>
        </div>
      );
    },
    width: 160,
  },
  {
    title: 'handle',
    dataIndex: 'handle',
    width: 180,
  },
  {
    title: '方法与路径',
    dataIndex: 'path',
    customRender: ({ record }) => {
      const mapping = {
        GET: 'blue',
        POST: 'success',
        DELETE: 'error',
        PUT: 'warning',
      };
      const color = mapping[record.action] || 'default';
      return (
        <div>
          <Tag color={color}>{() => record.action}</Tag>
          <span class="ml-2">{record.path}</span>
        </div>
      );
    },
    width: 180,
  },
  {
    title: '类型',
    dataIndex: 'type',
    width: 50,
    ifShow: false,
  },
  {
    title: '创建时间',
    dataIndex: 'createdAt',
    width: 140,
    format: 'datetime|flex',
  },
  {
    title: '创建者',
    dataIndex: 'createBy',
    width: 80,
  },
  {
    title: '更新时间',
    dataIndex: 'updatedAt',
    width: 140,
    format: 'datetime|flex',
  },
  {
    title: '更新者',
    dataIndex: 'updateBy',
    width: 80,
  },
];

const options = [
  {
    label: 'GET',
    value: 'GET',
  },
  {
    label: 'POST',
    value: 'POST',
  },
  {
    label: 'PUT',
    value: 'PUT',
  },
  {
    label: 'DELETE',
    value: 'DELETE',
  },
];

export const searchFormSchema: FormSchema[] = [
  {
    label: '标题',
    field: 'title',
    component: 'Input',
  },
  {
    label: '路径',
    field: 'path',
    component: 'Input',
  },
  {
    label: '方法',
    field: 'action',
    component: 'Select',
    componentProps: {
      options: options,
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
    field: 'title',
    label: '标题',
    component: 'Input',
    required: true,
  },
  {
    field: 'handle',
    label: 'handle',
    component: 'Input',
    required: true,
  },
  {
    field: 'action',
    label: '方法',
    component: 'Select',
    componentProps: {
      options: options,
    },
    required: true,
  },
  {
    field: 'path',
    label: '路径',
    component: 'Input',
    required: true,
  },
  {
    label: '类型',
    field: 'type',
    component: 'Select',
    componentProps: {
      options: [
        {
          label: 'CHECK',
          value: 'CHECK',
        },
        {
          label: 'NOCHECK',
          value: 'NOCHECK',
        },
      ],
    },
  },
];
