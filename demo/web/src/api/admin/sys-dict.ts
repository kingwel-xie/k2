import { SysDictDataListItem, SysDictTypeListItem } from '/@/api/admin/model/sys-dict';
import { defHttp } from '/@/utils/http/axios';
import { BasicFetchResult, BasicPageParams } from '/@/api/model/baseModel';

enum Api {
  SysDictTypeOp = '/v1/dict/type',
  SysDictDataOp = '/v1/dict/data',
}

// DictType
export const getSysDictTypeList = (params?: Required<BasicPageParams>) =>
  defHttp.get<BasicFetchResult<SysDictTypeListItem>>({ url: Api.SysDictTypeOp, params });

export const getSysDictTypeByKey = (id: number) =>
  defHttp.get<SysDictTypeListItem>({ url: Api.SysDictTypeOp + '/' + id });

export const addSysDictTypeEntry = (params: any) =>
  defHttp.post<any>({ url: Api.SysDictTypeOp, params });

export const updateSysDictTypeEntry = (params: any) =>
  defHttp.put<any>({ url: Api.SysDictTypeOp + '/' + params.id, params });

export const deleteSysDictTypeEntry = (params: any) =>
  defHttp.delete<any>({ url: Api.SysDictTypeOp, params: { ids: [params.id] } });

// DictData
export const getSysDictDataList = (params?: Required<BasicPageParams>) =>
  defHttp.get<BasicFetchResult<SysDictDataListItem>>({ url: Api.SysDictDataOp, params });

export const getSysDictDataByKey = (dictCode: number) =>
  defHttp.get<SysDictDataListItem>({ url: Api.SysDictDataOp + '/' + dictCode });

export const addSysDictDataEntry = (params: any) =>
  defHttp.post<any>({ url: Api.SysDictDataOp, params });

export const updateSysDictDataEntry = (params: any) =>
  defHttp.put<any>({ url: Api.SysDictDataOp + '/' + params.dictCode, params });

export const deleteSysDictDataEntry = (params: any) =>
  defHttp.delete<any>({ url: Api.SysDictDataOp, params: { ids: [params.dictCode] } });
