/**
 * @description: SysDictEntry interface return value
 */
export interface SysDictEntryModel {
  dictLabel: string;
  dictValue: string;
  dictType: string;
  dictSort: number;
  status: number;
  remark: string;
}

/**
 * @description: ExtDictEntry interface return value
 */
export interface ExtDictEntryModel {
  code: string;
  name: string;
}
