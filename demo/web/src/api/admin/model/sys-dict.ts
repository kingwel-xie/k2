// DictType
export interface SysDictTypeListItem {
  id: number;
  dictName: string;
  dictType: string;
  status: number;
  createdAt: string;
  remark: string;
}

// DictData
export interface SysDictDataListItem {
  dictCode: number;
  dictLabel: string;
  dictValue: string;
  dictSort: number;
  status: number;
  createdAt: string;
  remark: string;
}
