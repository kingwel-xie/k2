import { defHttp } from '/@/utils/http/axios';
import { BasicFetchResult, BasicPageParams } from '/@/api/model/baseModel';

enum Api {
  SysApiOp = '/v1/sys-api',
}

export const getSysApiList = (params?: Required<BasicPageParams>) =>
  defHttp.get<BasicFetchResult<any>>({ url: Api.SysApiOp, params });

export const getSysApiByKey = (id: number) => defHttp.get<any>({ url: Api.SysApiOp + '/' + id });

export const addSysApiEntry = (params: any) => defHttp.post<any>({ url: Api.SysApiOp, params });

export const updateSysApiEntry = (params: any) =>
  defHttp.put<any>({ url: Api.SysApiOp + '/' + params.id, params });

export const deleteSysApiEntry = (params: any) =>
  defHttp.delete<any>({ url: Api.SysApiOp, params: { ids: [params.id] } });
