<template>
  <PageWrapper title="导出示例" content="可以选择导出格式">
    <BasicTable title="基础表格" :columns="columns" :dataSource="data">
      <template #toolbar>
        <a-button @click="openModal"> 导出 </a-button>
      </template>
    </BasicTable>
    <ExpExcelModal @register="register" @success="defaultHeader" />
  </PageWrapper>
</template>

<script lang="ts">
  import { defineComponent } from 'vue';
  import { BasicTable } from '/@/components/Table';
  import { ExpExcelModal, ExportModalResult } from '/@/components/Excel2';
  import { columns, header, data } from './data';
  import { useModal } from '/@/components/Modal';
  import { PageWrapper } from '/@/components/Page';
  import { export2SingleSheetExcel } from '/@/utils/export2';

  export default defineComponent({
    components: { BasicTable, ExpExcelModal, PageWrapper },
    setup() {
      function defaultHeader({ filename, bookType }: ExportModalResult) {
        // 默认Object.keys(data[0])作为header
        export2SingleSheetExcel(
          {
            header,
            data,
            name: 'sheet1',
          },
          filename,
          bookType === 'csv',
        );
      }
      const [register, { openModal }] = useModal();

      return {
        defaultHeader,
        columns,
        data,
        register,
        openModal,
      };
    },
  });
</script>
