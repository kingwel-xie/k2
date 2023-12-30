<template>
  <PageWrapper title="excel数据导入示例">
    <ImpExcel2 @success="loadDataSuccess" dateFormat="YYYY-MM-DD">
      <a-button class="m-3"> 导入Excel </a-button>
    </ImpExcel2>
    <BasicTable
      v-for="(table, index) in tableListRef"
      :key="index"
      :title="table.title"
      :columns="table.columns"
      :dataSource="table.dataSource"
    />
  </PageWrapper>
</template>
<script lang="ts">
  import { defineComponent, ref } from 'vue';

  import { ImpExcel2, ImportedExcelData } from '/@/components/Excel2';
  import { BasicTable, BasicColumn } from '/@/components/Table';
  import { PageWrapper } from '/@/components/Page';

  export default defineComponent({
    components: { BasicTable, ImpExcel2, PageWrapper },

    setup() {
      const tableListRef = ref<
        {
          title: string;
          columns?: any[];
          dataSource?: any[];
        }[]
      >([]);

      function loadDataSuccess(excelDataList: ImportedExcelData) {
        tableListRef.value = [];
        console.log(excelDataList);
        for (const sheet of excelDataList.sheets) {
          const { header, results, name } = sheet;
          const columns: BasicColumn[] = [];
          for (const title of header) {
            columns.push({ title, dataIndex: title });
          }
          tableListRef.value.push({ title: name, dataSource: results, columns });
        }
      }

      return {
        loadDataSuccess,
        tableListRef,
      };
    },
  });
</script>
