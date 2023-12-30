<template>
  <div>
    <input
      ref="inputRef"
      type="file"
      v-show="false"
      accept=".xlsx, .xls"
      @change="handleInputClick"
    />
    <div @click="handleUpload" @dragover.prevent @drop="handleDrop">
      <slot></slot>
    </div>
  </div>
</template>
<script lang="ts">
  import { defineComponent, ref, unref } from 'vue';
  import * as exceljs from 'exceljs';
  import { dateUtil } from '/@/utils/dateUtil';
  import type { ImportedExcelData } from './typing';
  import { HIDDEN_DICT_SHEET } from '/@/store/modules/dictionary';

  export default defineComponent({
    name: 'ImportExcel2',
    props: {
      // 日期时间格式。如果不提供或者提供空值，将返回原始Date对象
      dateFormat: {
        type: String,
      },
      // 是否直接返回选中文件
      isReturnFile: {
        type: Boolean,
        default: false,
      },
    },
    emits: ['success', 'error'],
    setup(props, { emit }) {
      const inputRef = ref<HTMLInputElement | null>(null);
      const loadingRef = ref<Boolean>(false);

      /**
       * @description: 第一行作为头部
       */
      function getHeaderRow(sheet: exceljs.Worksheet) {
        const header: string[] = [];
        const row = sheet.getRow(1);
        row.eachCell((cell, _colNumber) => {
          // console.log('Cell ' + colNumber + ' = ' + cell.value, cell);
          header.push(cell.value as string);
        });
        return header;
      }

      /**
       * @description: 获得excel数据
       */
      function getExcelData(workbook: exceljs.Workbook) {
        // console.log(workbook);
        const excelData: ImportedExcelData = {
          creator: workbook.creator,
          created: workbook.created,
          version: workbook.description,
          sheets: [],
        };
        const { dateFormat } = props;
        for (const sheet of workbook.worksheets) {
          // bypass 'dict' sheet
          if (sheet.name === HIDDEN_DICT_SHEET) continue;

          const header = getHeaderRow(sheet);
          const excelList: Record<typeof header[number], string>[] = [];
          sheet
            .getSheetValues()
            // 移除空行
            .filter((temp) => !!temp?.length)
            .forEach((item) => {
              // 移除每行首个空元素
              (item as string[]).shift();
              // 定义临时对象存储每一行内容
              let tempObj: Record<typeof header[number], string> = {};
              (item as any[]).forEach((item2, index2) => {
                // handle date object, when dateFormat presents
                if (item2 instanceof Date) {
                  if (dateFormat) {
                    item2 = dateUtil(item2).format(dateFormat);
                  }
                }
                tempObj[header[index2]] = item2;
              });
              excelList.push(tempObj);
            });

          excelData.sheets.push({
            header,
            results: excelList.slice(1),
            name: sheet.name,
          });
        }
        return excelData;
      }

      /**
       * @description: 读取excel数据
       */
      function readerData(rawFile: File) {
        loadingRef.value = true;
        return new Promise((resolve, reject) => {
          const reader = new FileReader();
          reader.onload = async (e) => {
            try {
              const workbook = new exceljs.Workbook();
              await workbook.xlsx.load(e.target!.result as ArrayBuffer);
              // console.log(workbook);
              /* DO SOMETHING WITH workbook HERE */
              const excelData = getExcelData(workbook);
              emit('success', excelData);
              resolve('');
            } catch (error) {
              reject(error);
              emit('error');
            } finally {
              loadingRef.value = false;
            }
          };
          reader.readAsArrayBuffer(rawFile);
        });
      }

      async function upload(rawFile: File) {
        const inputRefDom = unref(inputRef);
        if (inputRefDom) {
          // fix can't select the same excel
          inputRefDom.value = '';
        }
        await readerData(rawFile);
      }

      /**
       * @description: 触发选择文件管理器
       */
      function handleInputClick(e: Event) {
        const files = e && (e.target as HTMLInputElement).files;
        const rawFile = files && files[0]; // only setting files[0]
        if (!rawFile) return;
        if (props.isReturnFile) {
          emit('success', rawFile);
          return;
        }
        upload(rawFile);
      }

      /**
       * @description: 点击上传按钮
       */
      function handleUpload() {
        const inputRefDom = unref(inputRef);
        inputRefDom && inputRefDom.click();
      }

      function handleDrop(e: DragEvent) {
        // 阻止默认行为
        e.preventDefault();
        // 获取拖放的文件
        const rawFile = e.dataTransfer?.files[0];
        if (!rawFile) return;

        if (props.isReturnFile) {
          emit('success', rawFile);
          return;
        }
        upload(rawFile);
      }

      return { handleDrop, handleUpload, handleInputClick, inputRef };
    },
  });
</script>
