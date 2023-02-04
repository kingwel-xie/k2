import * as xlsx from 'xlsx';
import type { WorkBook } from 'xlsx';
import type { JsonToSheet, AoAToSheet } from './typing';
import { MultiJsonToSheet } from './typing';

const { utils, writeFile } = xlsx;

const DEF_FILE_NAME = 'excel-list.xlsx';

/**
 * @param data source data
 * @param worksheet worksheet object
 * @param min min width
 */
function setColumnWidth(data, worksheet, min = 10) {
  const obj = {};
  worksheet['!cols'] = [];
  data.forEach((item) => {
    Object.keys(item).forEach((key) => {
      let length = min;
      const cur = item[key];
      if (cur) {
        /*判断是否为中文*/
        if (cur.toString().charCodeAt(0) > 255) {
          length = cur.toString().length * 2;
        } else {
          length = cur.toString().length;
        }
      }
      obj[key] = Math.max(min, length);
    });
  });
  Object.keys(obj).forEach((key) => {
    worksheet['!cols'].push({
      wch: obj[key],
    });
  });
}

export function multiJsonToSheetXlsx<T = any>({
  multiData,
  filename = DEF_FILE_NAME,
  json2sheetOpts = {},
  write2excelOpts = { bookType: 'xlsx' },
}: MultiJsonToSheet<T>) {
  // construct the workbook
  const workbook: WorkBook = {
    SheetNames: [],
    Sheets: {},
  };
  multiData.forEach(({ sheetName, data, header }) => {
    const arrData = [...data];
    if (header) {
      arrData.unshift(header);
      json2sheetOpts.skipHeader = true;
    }
    const worksheet = utils.json_to_sheet(arrData, json2sheetOpts);
    setColumnWidth(arrData, worksheet);
    /* add worksheet to workbook */
    workbook.SheetNames.push(sheetName);
    workbook.Sheets[sheetName] = worksheet;
  });

  /* output format determined by filename */
  writeFile(workbook, filename, write2excelOpts);
  /* at this point, out.xlsb will have been downloaded */
}

export function jsonToSheetXlsx<T = any>({
  data,
  header,
  filename = DEF_FILE_NAME,
  json2sheetOpts = {},
  write2excelOpts = { bookType: 'xlsx' },
}: JsonToSheet<T>) {
  const arrData = [...data];
  if (header) {
    arrData.unshift(header);
    json2sheetOpts.skipHeader = true;
  }

  const worksheet = utils.json_to_sheet(arrData, json2sheetOpts);
  setColumnWidth(arrData, worksheet);
  /* add worksheet to workbook */
  const workbook: WorkBook = {
    SheetNames: [filename],
    Sheets: {
      [filename]: worksheet,
    },
  };
  /* output format determined by filename */
  writeFile(workbook, filename, write2excelOpts);
  /* at this point, out.xlsb will have been downloaded */
}

export function aoaToSheetXlsx<T = any>({
  data,
  header,
  filename = DEF_FILE_NAME,
  write2excelOpts = { bookType: 'xlsx' },
}: AoAToSheet<T>) {
  const arrData = [...data];
  if (header) {
    arrData.unshift(header);
  }

  const worksheet = utils.aoa_to_sheet(arrData);

  /* add worksheet to workbook */
  const workbook: WorkBook = {
    SheetNames: [filename],
    Sheets: {
      [filename]: worksheet,
    },
  };
  /* output format determined by filename */
  writeFile(workbook, filename, write2excelOpts);
  /* at this point, out.xlsb will have been downloaded */
}
