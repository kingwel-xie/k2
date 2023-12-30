<template>
  <PageWrapper title="导出示例" content="根据数组格式的数据进行导出">
    <BasicTable title="基础表格" :columns="columns" :dataSource="data">
      <template #toolbar>
        <a-button @click="aoaToExcel"> 导出 </a-button>
      </template>
    </BasicTable>
  </PageWrapper>
</template>

<script lang="ts">
  import { defineComponent } from 'vue';
  import { BasicTable } from '/@/components/Table';
  import { arrHeader, arrData, columns, data } from './data';
  import { PageWrapper } from '/@/components/Page';
  import { export2SingleSheetExcel } from '/@/utils/export2';

  export default defineComponent({
    components: { BasicTable, PageWrapper },
    setup() {
      function aoaToExcel() {
        // 保证data顺序与header一致
        export2SingleSheetExcel(
          {
            data: arrData,
            header: arrHeader,
            name: 'sheet1',
          },
          '二维数组方式导出excel.xlsx',
        );
      }

      return {
        aoaToExcel,
        columns,
        data,
      };
    },
  });
</script>
