import { defHttp } from '/@/utils/http/axios';
import { BasicFetchResult, BasicPageParams } from '/@/api/model/baseModel';

enum Api {
  SysLoginLogOp = '/v1/sys-login-log',
}

export const getSysLoginLogList = (params?: Required<BasicPageParams>) =>
  defHttp.get<BasicFetchResult<any>>({ url: Api.SysLoginLogOp, params });

export const getSysLoginLogById = (id: number) =>
  defHttp.get<any>({ url: Api.SysLoginLogOp + '/' + id });

export const deleteSysLoginLogEntry = (params: any) =>
  defHttp.delete<any>({ url: Api.SysLoginLogOp, params: { ids: [params.id] } });
