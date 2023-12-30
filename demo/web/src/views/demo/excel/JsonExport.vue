<template>
  <PageWrapper title="导出示例" content="根据JSON格式的数据进行导出">
    <BasicTable title="基础表格" :columns="columns" :dataSource="data">
      <template #toolbar>
        <a-button @click="customHeader"> 导出：自定义头部 </a-button>
      </template>
    </BasicTable>
  </PageWrapper>
</template>

<script lang="ts">
  import { defineComponent } from 'vue';
  import { BasicTable } from '/@/components/Table';
  import { columns, data } from './data';
  import { PageWrapper } from '/@/components/Page';
  import { export2SingleSheetExcel } from '/@/utils/export2';

  export default defineComponent({
    components: { BasicTable, PageWrapper },
    setup() {
      function customHeader() {
        export2SingleSheetExcel(
          {
            data,
            header: {
              id: 'ID',
              name: '姓名',
              age: '年龄',
              no: '编号',
              address: '地址',
              beginTime: '开始时间',
              endTime: '结束时间',
            },
            name: 'sheet1',
          },
          '自定义头部.xlsx',
        );
      }

      return {
        customHeader,
        columns,
        data,
      };
    },
  });
</script>
