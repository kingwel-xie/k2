import { isString } from '/@/utils/is';
import { formatValue } from '/@/utils/formatUtil';
import { jsonToSheetXlsx } from '/@/components/Excel';
import { BasicColumn } from '/@/components/Table';

/**
 * export json data to Excel file, as described by header and columns
 * @param header, 传入
 * @param columns, 传入
 * @param data, 传入
 * @param filename, 传入
 */
export function export2Excel<T = any>(
  header: T,
  columns: BasicColumn[],
  data: T[],
  filename: string,
) {
  const headerFields = Object.keys(header);
  const formattedData = data.map((x) => {
    const xx = {};
    Object.keys(x)
      .filter((key) => headerFields.includes(key))
      .forEach((key) => {
        const col = columns.find((col) => col.dataIndex === key);
        // hard code, dont' export 'flex' datetime to excel
        let format = col && col.format;
        if (format === 'datetime|flex') format = 'datetime|';
        xx[key] = format && isString(format) ? formatValue(format, x[key]) : x[key];
      });
    return xx;
  });

  jsonToSheetXlsx({
    data: formattedData,
    header,
    filename,
  });
}
