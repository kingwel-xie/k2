import { defHttp } from '/@/utils/http/axios';
import { getMenuListResultModel } from './model/menuModel';

enum Api {
  GetRoleMenus = '/v1/menu-role',
}

/**
 * @description: Get user menu based on id
 */

export const getMenuRoleList = () => {
  return defHttp.get<getMenuListResultModel>({ url: Api.GetRoleMenus });
};
