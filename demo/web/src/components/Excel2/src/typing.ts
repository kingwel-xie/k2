export interface ImportedSheet<T = any> {
  header: string[];
  results: T[];
  name: string;
}

export interface ImportedExcelData<T = any> {
  version?: string;
  creator: string;
  created: Date;
  sheets: ImportedSheet<T>[];
}

export interface ExportModalResult {
  filename: string;
  bookType: string;
}
