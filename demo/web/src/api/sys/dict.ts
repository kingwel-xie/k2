import { defHttp } from '/@/utils/http/axios';

enum Api {
  LoadAll = '/v1/all-dict',
}

/**
 * @description: getUserInfo
 */
export function getAllDicts() {
  return defHttp.get<any>({ url: Api.LoadAll }, { errorMessageMode: 'none' });
}
