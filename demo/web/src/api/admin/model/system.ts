import { BasicPageParams, BasicFetchResult } from '/@/api/model/baseModel';

export type AccountParams = BasicPageParams & {
  account?: string;
  nickName?: string;
};

export type RoleParams = {
  roleName?: string;
  status?: string;
};

export type RolePageParams = BasicPageParams & RoleParams;

export type DeptParams = {
  deptName?: string;
  status?: string;
};

export type DeptPageParams = BasicPageParams & DeptParams;

export type MenuParams = {
  menuName?: string;
  status?: string;
};

export interface AccountListItem {
  userId: number;
  username: string;
  email: string;
  nickName: string;
  role: number;
  createdAt: string;
  remark: string;
  status: string;
}

export interface DeptListItem {
  deptId: number;
  deptName: string;
  sort: string;
  createdAt: string;
  remark: string;
  status: number;
}

export interface GenericTreeItem {
  id: number;
  label: string;
  children?: GenericTreeItem[];
}

export interface MenuListItem {
  menuId: number;
  menuName: string;
  title: string;
  sort: string;
  createdAt: string;
  visible: string;
  icon: string;
  component: string;
  path: string;
  permission: string;
}

export interface RoleListItem {
  roleId: number;
  roleName: string;
  roleKey: string;
  status: string;
  dataScope?: string;
  roleSort: string;
  createdAt: string;
}

/**
 * @description: Request list return value
 */
export type AccountListGetResultModel = BasicFetchResult<AccountListItem>;

export type DeptListGetResultModel = DeptListItem[];

export type DeptListGetResultModel2 = BasicFetchResult<DeptListItem>;

export type DeptTreeGetResultModel = GenericTreeItem[];

export type MenuListGetResultModel = MenuListItem[];

export type MenuTreeGetResultModel = { menus: GenericTreeItem[] };

export type RolePageListGetResultModel = BasicFetchResult<RoleListItem>;

export type RoleListGetResultModel = BasicFetchResult<RoleListItem>;
