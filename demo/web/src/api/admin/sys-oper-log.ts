import { defHttp } from '/@/utils/http/axios';
import { BasicFetchResult, BasicPageParams } from '/@/api/model/baseModel';

enum Api {
  SysOperLogOp = '/v1/sys-opera-log',
}

export const getSysOperLogList = (params?: Required<BasicPageParams>) =>
  defHttp.get<BasicFetchResult<any>>({ url: Api.SysOperLogOp, params });

export const getSysOperLogById = (id: number) =>
  defHttp.get<any>({ url: Api.SysOperLogOp + '/' + id });

export const deleteSysOperLogEntry = (params: any) =>
  defHttp.delete<any>({ url: Api.SysOperLogOp, params: { ids: [params.id] } });
