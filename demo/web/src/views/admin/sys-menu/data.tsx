import { BasicColumn } from '/@/components/Table';
import { FormSchema } from '/@/components/Table';
import { h } from 'vue';
import { Tag } from 'ant-design-vue';
import { Icon } from '/@/components/Icon';
import ApiDetail from '/@/views/admin/sys-menu/components/ApiDetail.vue';

export const columns: BasicColumn[] = [
  {
    title: '菜单名称',
    dataIndex: 'title',
    width: 180,
    align: 'left',
  },
  {
    title: '图标',
    dataIndex: 'icon',
    width: 70,
    customRender: ({ record }) => {
      return h(Icon, { icon: record.iconAntd });
    },
  },
  {
    title: '排序',
    dataIndex: 'sort',
    width: 60,
  },
  {
    title: '权限标识',
    dataIndex: 'permission',
    width: 220,
    align: 'left',
    customRender: ({ record }) => {
      if (!record.permission || !record.sysApi || record.sysApi.length === 0)
        return record.permission;
      // @ts-ignore
      return <ApiDetail data={record} />;
    },
  },
  {
    title: '组件路径',
    dataIndex: 'component',
    align: 'left',
    width: 180,
  },
  {
    title: '路由地址',
    dataIndex: 'path',
    align: 'left',
    width: 180,
  },
  {
    title: '可见',
    dataIndex: 'visible',
    width: 80,
    customRender: ({ record }) => {
      if (isButton(record.menuType)) return '-';
      const status = record.visible;
      const enable = status === '0';
      const color = enable ? 'green' : 'red';
      const text = enable ? '显示' : '隐藏';
      return h(Tag, { color: color }, () => text);
    },
  },
  {
    title: '缓存',
    dataIndex: 'noCache',
    format: 'bool|否,是',
    width: 60,
  },
  {
    title: '创建时间',
    dataIndex: 'createdAt',
    width: 180,
    format: 'datetime|flex',
  },
];

// const isDir = (type: string) => type === 'M';
const isMenu = (type: string) => type === 'C';
const isButton = (type: string) => type === 'F';

export const searchFormSchema: FormSchema[] = [
  {
    field: 'menuName',
    label: '菜单名称',
    component: 'Input',
  },
  {
    field: 'status',
    label: '状态',
    component: 'Select',
    componentProps: {
      options: [
        { label: '启用', value: '0' },
        { label: '停用', value: '1' },
      ],
    },
  },
];

export const formSchema: FormSchema[] = [
  {
    field: 'menuId',
    label: '菜单Id',
    component: 'InputNumber',
    show: false,
  },
  {
    field: 'menuType',
    label: '菜单类型',
    component: 'RadioButtonGroup',
    defaultValue: 'C',
    componentProps: {
      options: [
        { label: '目录', value: 'M' },
        { label: '菜单', value: 'C' },
        { label: '按钮', value: 'F' },
      ],
    },
    helpMessage: '包含目录：以及菜单或者菜单组，菜单：具体对应某一个页面，按钮：功能按钮',
    colProps: { lg: 24, md: 24 },
  },
  {
    field: 'title',
    label: '菜单名称',
    component: 'Input',
    required: true,
  },

  {
    field: 'parentId',
    label: '上级菜单',
    component: 'TreeSelect',
    componentProps: {
      fieldNames: {
        label: 'label',
        key: 'id',
        value: 'id',
      },
      getPopupContainer: () => document.body,
    },
    helpMessage: '指当前菜单停靠的菜单归属',
    required: true,
  },

  {
    field: 'sort',
    label: '显示排序',
    component: 'InputNumber',
    required: true,
  },
  {
    field: 'iconAntd',
    label: '菜单图标',
    component: 'IconPicker',
    ifShow: ({ values }) => !isButton(values.menuType),
  },

  {
    field: 'menuName',
    label: '菜单标识',
    component: 'Input',
    required: true,
    helpMessage: '需要和页面 component name 保持一致，页面才可以被缓存',
    ifShow: ({ values }) => !isButton(values.menuType),
  },
  {
    field: 'path',
    label: '路由地址',
    helpMessage: '访问此页面自定义的url地址，建议以 / 开头书写，例如 /app-name/menu-name',
    component: 'Input',
    required: true,
    ifShow: ({ values }) => !isButton(values.menuType),
  },
  {
    field: 'component',
    label: '组件路径',
    helpMessage:
      '菜单对应的具体vue页面文件路径views的下级路径，如：/admin/sys-api/index；目录类型：请填写Layout，如何有二级目录请参照日志目录填写',
    component: 'Input',
    ifShow: ({ values }) => !isButton(values.menuType),
  },

  {
    field: 'isFrame',
    label: '是否外链',
    component: 'RadioButtonGroup',
    defaultValue: '1',
    componentProps: {
      options: [
        { label: '是', value: '0' },
        { label: '否', value: '1' },
      ],
    },
    ifShow: ({ values }) => !isButton(values.menuType),
  },

  {
    field: 'visible',
    label: '是否显示',
    helpMessage: '出现在菜单列表的菜单项设置为显示，否则设置为隐藏',
    component: 'RadioButtonGroup',
    defaultValue: '0',
    componentProps: {
      options: [
        { label: '是', value: '0' },
        { label: '否', value: '1' },
      ],
    },
    ifShow: ({ values }) => !isButton(values.menuType),
  },

  {
    field: 'noCache',
    label: '是否缓存',
    helpMessage: '启用 keep-alive 机制缓存页面',
    component: 'RadioButtonGroup',
    defaultValue: false,
    componentProps: {
      options: [
        { label: '是', value: false },
        { label: '否', value: true },
      ],
    },
    ifShow: ({ values }) => isMenu(values.menuType),
  },
  {
    field: 'paths',
    label: 'paths',
    component: 'Input',
    show: false,
  },
  {
    field: 'icon',
    label: 'icon',
    component: 'Input',
    show: false,
  },
  {
    field: 'action',
    label: 'action',
    component: 'Input',
    show: false,
  },

  {
    field: 'permission',
    label: '权限标识',
    helpMessage: '前端权限控制按钮是否显示',
    component: 'Input',
    ifShow: ({ values }) => isButton(values.menuType),
  },
  {
    field: 'divider0',
    component: 'Divider',
    label: '❁',
    colProps: { span: 24 },
    ifShow: ({ values }) => isButton(values.menuType),
  },
  {
    field: 'apis',
    label: 'API 权限',
    valueField: 'targetKeys',
    helpMessage: '配置此条目需要使用到的API接口，否则在设置用户角色时，接口将无权访问',
    component: 'Transfer',
    componentProps: {
      // :dataSource="availableItems"
      showSearch: true,
      listStyle: { width: '280px', height: '390px' },
      titles: [' 未选中', ' 已选中'],
      rowKey: (item) => String(item.id),
      render: (item) => item.title,
    },
    colProps: { span: 24 },
    ifShow: ({ values }) => isButton(values.menuType),
  },
];
