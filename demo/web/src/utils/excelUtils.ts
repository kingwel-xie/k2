import { Column, DataValidation } from 'exceljs';

export interface KobhWorksheet {
  name: string;
  hidden?: boolean;
  autoWidth?: boolean;
  columns: Partial<Column>[];
  validations?: {
    header: string;
    v: DataValidation | ((r: number, c: number) => DataValidation);
  }[];
  rows: Recordable[];
}

export interface KobhWorkbook {
  filename: string;
  version?: string;
  csvFormat?: boolean;
  sheets: KobhWorksheet[];
}

/**
 * return Excel column name for the given column index
 * @param index column index, 1 based
 */
export function excelColumnIndex2Name(index: number): string {
  let columnName = '';
  while (index > 0) {
    const remainder = (index - 1) % 26;
    columnName = String.fromCharCode(65 + remainder) + columnName;
    index = Math.floor((index - 1) / 26);
  }
  return columnName;
}
