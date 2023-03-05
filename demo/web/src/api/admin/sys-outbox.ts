import { defHttp } from '/@/utils/http/axios';
import { BasicFetchResult, BasicPageParams } from '/@/api/model/baseModel';

enum Api {
  SysOutboxOp = '/v1/outbox',
}

export const getSysOutboxList = (params?: Required<BasicPageParams>) =>
  defHttp.get<BasicFetchResult<any>>({ url: Api.SysOutboxOp, params });

export const getSysOutboxByKey = (id: string | number) =>
  defHttp.get<any>({ url: Api.SysOutboxOp + '/' + id });

export const deleteSysOutboxEntry = (params: any) =>
  defHttp.delete<any>({ url: Api.SysOutboxOp, params: { ids: [params.id] } });

export const deleteSysOutboxMany = (ids: number[]) =>
  defHttp.delete<any>({ url: Api.SysOutboxOp, params: { ids } });
