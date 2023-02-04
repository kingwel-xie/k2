<template>
  <Popover trigger="click" title="API明细" :overlayStyle="{ width: '550px', height: '200px' }">
    <template #content>
      <Table size="small" bordered :dataSource="data.sysApi" :columns="columns" />
    </template>
    <div class="cursor-pointer underline">
      <div>✲ {{ data.permission }}</div>
    </div>
  </Popover>
</template>

<script lang="tsx">
  export default {
    inheritAttrs: false,
  };
</script>

<script lang="tsx" setup>
  import { PropType } from 'vue';
  import { Popover, Table, Tag } from 'ant-design-vue';

  const columns = [
    {
      title: 'API 名称',
      dataIndex: 'title',
      customRender: ({ text }) => {
        return <Tag color="success">{() => text}</Tag>;
      },
    },
    {
      title: 'URL 及方法',
      dataIndex: 'path',
      customRender: ({ record }) => {
        const mapping = {
          GET: 'processing',
          POST: 'success',
          PUT: 'warning',
          DELETE: 'error',
        };
        const color = mapping[record.action] || 'default';
        return (
          <span>
            <Tag color={color}>{() => record.action}</Tag>
            <span>{record.path}</span>
          </span>
        );
      },
    },
  ];

  defineProps({
    data: {
      type: Object as PropType<Recordable>,
      default: () => {},
    },
  });
</script>
