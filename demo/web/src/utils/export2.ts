import * as exceljs from 'exceljs';
import type { Column } from 'exceljs';
import { downloadByData } from '/@/utils/file/download';
import { BasicColumn } from '/@/components/Table';
import { isBoolean, isFunction, isNumber, isString } from '/@/utils/is';
import {
  dictValidationRef,
  getSupplierOptions,
  useDictStoreWithOut,
} from '/@/store/modules/dictionary';
import { KobhWorkbook } from '/@/utils/excelUtils';
import { formatValue } from '/@/utils/formatUtil';

export interface JsonSheet<T> {
  header: T;
  data: T[];
  name: string;
}

function header2Columns<T = any>(header: T) {
  const columns = [] as any[];
  Object.entries(header).forEach(([key, header]) => {
    columns.push({ header, key });
  });
  return columns;
}

export function formatJsonData<T = any>(header: T, columns: BasicColumn[], data: T[]) {
  const headerFields = Object.keys(header);
  return data.map((x) => {
    const xx = {};
    headerFields.forEach((key) => {
      const col = columns.find((col) => col.dataIndex === key);
      // firstly, try column.exportFunc
      if (col && col.exportFunc) {
        xx[key] = col.exportFunc(x[key], x);
        return;
      }
      // otherwise, try column.format
      // hard code, don't export 'flex' datetime to excel
      let format = col && col.format;
      if (format === 'datetime|flex') format = 'datetime|';
      // only care about 'string | boolean' values
      xx[key] =
        format && isString(format) && (isString(x[key]) || isBoolean(x[key]) || isNumber(x[key]))
          ? formatValue(format, x[key])
          : x[key];
    });
    return xx;
  });
}

/**
 * @param data source data
 * @param worksheet worksheet object
 * @param min min width
 */
function setColumnWidth(data: any[], columns: Partial<Column>[], min = 8) {
  const columnLengths = [] as any[];
  data.forEach((item) => {
    Object.keys(item).forEach((key) => {
      let length = min;
      const cur = item[key];
      if (cur) {
        /*判断是否为中文*/
        if (cur.toString().charCodeAt(0) > 255) {
          length = cur.toString().length * 2;
        } else if (cur.toString().length > 11) {
          /*判断字符串长度是否大于11*/
          length = cur.toString().length * 1.2;
        } else {
          length = cur.toString().length;
        }
      }
      if (key == 'length') {
        key = '__length__';
      }
      /*数组默认有 length 属性，在判断时为假，就不会运行 columnLengths[key] = [];*/
      if (!columnLengths[key]) {
        columnLengths[key] = [];
      }
      /*数组默认得 length 属性，没有 push 方法*/
      columnLengths[key].push(Math.max(min, length));
    });
  });
  Object.values(columnLengths).map((column) => {
    column.maxWidth = Math.max(...column);
  });

  columns.forEach((col) => {
    if (col.key) {
      col.width = columnLengths[col.key].maxWidth;
    }
  });
}

/**
 * export to Excel file, as described by workbook
 * @param w, 传入
 */
export function workbook2Excel(w: KobhWorkbook) {
  const workbook = new exceljs.Workbook();
  workbook.creator = 'kobh';
  workbook.lastModifiedBy = 'kobh';
  workbook.created = new Date();
  workbook.description = w.version || '';

  // a hardcode for __dict__ and definedNames
  const sheets = [...w.sheets];
  if (w.sheets.some((x) => !!x.validations)) {
    sheets.push(useDictStoreWithOut().validationSheet.dictSheet!);
    const options = getSupplierOptions();
    options.forEach((o) => {
      workbook.definedNames.add(dictValidationRef(o.value)[0], o.label);
    });
  }

  sheets.forEach((s) => {
    const sheet = workbook.addWorksheet(s.name);
    if (s.columns) {
      sheet.columns = s.columns;
      sheet.columns.forEach((c) => {
        // header as key if key is not specified
        if (!c.key && isString(c.header)) {
          c.key = c.header;
        }
      });
    }
    if (s.rows) {
      sheet.addRows(s.rows);
      if (s.autoWidth && sheet.columns) {
        setColumnWidth(s.rows, sheet.columns);
      }
    }
    if (s.hidden) {
      sheet.state = 'hidden';
    }
    s.validations?.forEach(({ header, v }) => {
      const _index = sheet.columns.findIndex((x) => x.header === header);
      if (_index !== -1) {
        // locate the cell and set data validation
        for (let i = 2; i < 20; i++) {
          sheet.getCell(i, _index + 1).dataValidation = isFunction(v) ? v(i, _index + 1) : v;
        }
      }
    });
  });

  const engine = w.csvFormat ? workbook.csv : workbook.xlsx;
  engine.writeBuffer().then((data) => {
    downloadByData(data, w.filename);
  });
}

/**
 * export json data to Excel file, as described by header and columns
 * @param header, 传入
 * @param columns, 传入
 * @param data, 传入
 * @param filename, 传入
 */
export function export2ExcelV2<T = any>(
  header: T,
  columns: BasicColumn[],
  data: T[],
  filename: string,
) {
  const formattedData = formatJsonData(header, columns, data);
  export2SingleSheetExcel({ header, data: formattedData, name: 'sheet1' }, filename);
}

/**
 * export json data to Excel single sheet file, as described by JsonSheet
 * @param s, 传入
 * @param filename, 传入
 */
export function export2SingleSheetExcel<T = any>(
  s: JsonSheet<T>,
  filename: string,
  csvFormat = false,
) {
  const workbook: KobhWorkbook = {
    filename,
    csvFormat,
    sheets: [
      {
        name: s.name,
        autoWidth: true,
        columns: header2Columns(s.header),
        rows: s.data,
      },
    ],
  };
  workbook2Excel(workbook);
}

/**
 * export json data to Excel multi sheet file, as described by [JsonSheet]
 * @param ss, 传入
 * @param filename, 传入
 */
export function export2ExcelMultiSheet<T = any>(
  ss: JsonSheet<T>[],
  filename: string,
  csvFormat = false,
) {
  const sheets = ss.map(({ header, data, name }) => {
    return {
      name,
      autoWidth: true,
      columns: header2Columns(header),
      rows: data,
    };
  });
  const workbook: KobhWorkbook = {
    filename,
    csvFormat,
    sheets,
  };
  workbook2Excel(workbook);
}
