import { withInstall } from '/@/utils';
import impExcel from './src/ImportExcel.vue';
import expExcelModal from './src/ExportExcelModal.vue';

export const ImpExcel2 = withInstall(impExcel);
export const ExpExcelModal = withInstall(expExcelModal);
export * from './src/typing';
