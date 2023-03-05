import { defHttp } from '/@/utils/http/axios';
import { BasicFetchResult, BasicPageParams } from '/@/api/model/baseModel';

enum Api {
  SysNotificationOp = '/v1/notification',
}

export const getSysNotificationList = (params?: Required<BasicPageParams>) =>
  defHttp.get<BasicFetchResult<any>>({ url: Api.SysNotificationOp, params });

export const getSysNotificationByKey = (id: string | number) =>
  defHttp.get<any>({ url: Api.SysNotificationOp + '/' + id });

export const addSysNotificationEntry = (params: any) =>
  defHttp.post<any>({ url: Api.SysNotificationOp, params });

export const updateSysNotificationEntry = (params: any) =>
  defHttp.put<any>({ url: Api.SysNotificationOp + '/' + params.id, params });

export const deleteSysNotificationEntry = (params: any) =>
  defHttp.delete<any>({ url: Api.SysNotificationOp, params: { ids: [params.id] } });
