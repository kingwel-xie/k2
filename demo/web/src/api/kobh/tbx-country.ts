import { defHttp } from '/@/utils/http/axios';
import { BasicFetchResult, BasicPageParams } from '/@/api/model/baseModel';

enum Api {
  TbxCountryOp = '/v1/country',
}

export const getTbxCountryList = (params?: Required<BasicPageParams>) =>
  defHttp.get<BasicFetchResult<any>>({ url: Api.TbxCountryOp, params });

export const getTbxCountryByKey = (code: string | number) =>
  defHttp.get<any>({ url: Api.TbxCountryOp + '/' + code });

export const addTbxCountryEntry = (params: any) =>
  defHttp.post<any>({ url: Api.TbxCountryOp, params });

export const updateTbxCountryEntry = (params: any) =>
  defHttp.put<any>({ url: Api.TbxCountryOp + '/' + params.code, params });

export const deleteTbxCountryEntry = (params: any) =>
  defHttp.delete<any>({ url: Api.TbxCountryOp, params: { ids: [params.code] } });

