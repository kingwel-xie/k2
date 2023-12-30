import {
  AccountParams,
  DeptParams,
  MenuParams,
  RoleParams,
  RolePageParams,
  MenuListGetResultModel,
  DeptListGetResultModel,
  AccountListGetResultModel,
  RolePageListGetResultModel,
  RoleListGetResultModel,
  DeptPageParams,
  DeptTreeGetResultModel,
  MenuTreeGetResultModel,
  RoleListItem,
  DeptListGetResultModel2,
  DeptListItem,
  MenuListItem,
} from './model/system';
import { defHttp } from '/@/utils/http/axios';

enum Api {
  AccountOp = '/v1/sys-user',
  setAccountStatus = '/v1/sys-user/status',
  AccountPwdReset = '/v1/sys-user/pwd/reset',
  AccountTokenReset = '/v1/sys-user/token/reset',
  IsAccountExist = '/v1/check-existence',
  DeptOp = '/v1/dept',
  DeptTree = '/v1/deptTree',
  setRoleStatus = '/v1/role-status',
  MenuOp = '/v1/menu',
  MenuTree = '/v1/roleMenuTreeSelect/0',
  RoleOp = '/v1/role',
  ListRoleNoCheckOp = '/v1/role-list',
  ListDeptNoCheckOp = '/v1/dept-list',
  ListAccountNoCheckOp = '/v1/user/list',
}

// sys-dept ////
export const getDeptList = (params?: DeptPageParams) =>
  defHttp.get<DeptListGetResultModel>({ url: Api.DeptOp, params });

export const getDeptTree = (params?: DeptPageParams) =>
  defHttp.get<DeptTreeGetResultModel>({ url: Api.DeptTree, params });

export const listDeptNoCheck = (params?: DeptParams) =>
  defHttp.get<DeptListGetResultModel2>({ url: Api.ListDeptNoCheckOp, params });

export const addDeptEntry = (params: any) => defHttp.post<any>({ url: Api.DeptOp, params });

export const getDeptByKey = (deptId: any) =>
  defHttp.get<DeptListItem>({ url: Api.DeptOp + '/' + deptId });

export const updateDeptEntry = (params: any) =>
  defHttp.put<any>({ url: Api.DeptOp + '/' + params.deptId, params });

export const deleteDeptEntry = (params: any) =>
  defHttp.delete<any>({ url: Api.DeptOp, params: { ids: [params.deptId] } });

// menu ////
export const getMenuList = (params?: MenuParams) =>
  defHttp.get<MenuListGetResultModel>({ url: Api.MenuOp, params });

export const getMenuTree = (params?: MenuParams) =>
  defHttp.get<MenuTreeGetResultModel>({ url: Api.MenuTree, params });

export const addMenuEntry = (params: any) => defHttp.post<any>({ url: Api.MenuOp, params });

export const getMenuByKey = (menuId: any) =>
  defHttp.get<MenuListItem>({ url: Api.MenuOp + '/' + menuId });

export const updateMenuEntry = (params: any) =>
  defHttp.put<any>({ url: Api.MenuOp + '/' + params.menuId, params });

export const deleteMenuEntry = (params: any) =>
  defHttp.delete<any>({ url: Api.MenuOp, params: { ids: [params.menuId] } });

// role ////
export const getRoleListByPage = (params?: RolePageParams) =>
  defHttp.get<RolePageListGetResultModel>({ url: Api.RoleOp, params });

export const listRoleNoCheck = (params?: RoleParams) =>
  defHttp.get<RoleListGetResultModel>({ url: Api.ListRoleNoCheckOp, params });

export const getRoleByKey = (id: number) =>
  defHttp.get<RoleListItem>({ url: Api.RoleOp + '/' + id });

export const addRoleEntry = (params?: any) => defHttp.post<any>({ url: Api.RoleOp, params });

export const updateRoleEntry = (params?: any) =>
  defHttp.put<any>({ url: Api.RoleOp + '/' + params.roleId, params });

export const deleteRoleEntry = (params: any) =>
  defHttp.delete<any>({ url: Api.RoleOp, params: { ids: [params.roleId] } });

export const setRoleStatus = (id: number, status: string) =>
  defHttp.put({ url: Api.setRoleStatus, params: { roleId: id, status } });

// user ////
export const getAccountList = (params: AccountParams) =>
  defHttp.get<AccountListGetResultModel>({ url: Api.AccountOp, params });

export const listAccountNoCheck = (params?: AccountParams) =>
  defHttp.get<AccountListGetResultModel>({ url: Api.ListAccountNoCheckOp, params });

export const isAccountExist = (username: string) =>
  defHttp.get(
    { url: Api.IsAccountExist + '/' + username },
    { isTransformResponse: false, withToken: false, errorMessageMode: 'none' },
  );

export const setAccountStatus = (id: number, status: string) =>
  defHttp.put({ url: Api.setAccountStatus, params: { userId: id, status } });

export const getAccountByKey = (userId: any) =>
  defHttp.get<any>({ url: Api.AccountOp + '/' + userId });

export const addAccountEntry = (params?: any) => defHttp.post<any>({ url: Api.AccountOp, params });

// Note: updateAccountEntry is kind of special, key 'userId' is not specified in URI... but in json
export const updateAccountEntry = (params?: any) =>
  defHttp.put<any>({ url: Api.AccountOp, params });

export const deleteAccountEntry = (params: any) =>
  defHttp.delete<any>({ url: Api.AccountOp, params: { ids: [params.userId] } });

export const resetUserPwd = (userId: number, password: string) =>
  defHttp.put<any>({ url: Api.AccountPwdReset, params: { userId, password } });

export const resetUserToken = (userId: number) =>
  defHttp.put<any>({ url: Api.AccountTokenReset, params: { userId } });
