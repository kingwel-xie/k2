import { defHttp } from '/@/utils/http/axios';
import { BasicFetchResult, BasicPageParams } from '/@/api/model/baseModel';

enum Api {
  SysInboxOp = '/v1/inbox',
}

export const getSysInboxList = (params?: Required<BasicPageParams>) =>
  defHttp.get<BasicFetchResult<any>>({ url: Api.SysInboxOp, params });

export const getSysInboxByKey = (id: string | number) =>
  defHttp.get<any>({ url: Api.SysInboxOp + '/' + id });

export const addSysInboxEntry = (params: any) => defHttp.post<any>({ url: Api.SysInboxOp, params });

export const updateSysInboxEntry = (params: any) =>
  defHttp.put<any>({ url: Api.SysInboxOp + '/' + params.id, params });

export const deleteSysInboxEntry = (params: any) =>
  defHttp.delete<any>({ url: Api.SysInboxOp, params: { ids: [params.id] } });
