/**
 * @description: Login interface parameters
 */
export interface LoginParams {
  username: string;
  password: string;
  // kingwel, for verification
  code?: string;
  uuid?: string;
  // as customer
  role?: string;
}

export interface RoleInfo {
  roleName: string;
  value: string;
}

/**
 * @description: Login interface return value
 */
export interface LoginResultModel {
  code: number;
  token: string;
  expired: string;
  msg?: string;
  // role: RoleInfo;
}

/**
 * @description: Get user information return value
 */
export interface GetUserInfoModel {
  roles: string[];
  // 用户id
  userId: string | number;
  // 用户名
  userName: string;
  // 真实名字
  name: string;
  // 头像
  avatar: string;
  // 权限
  permissions: string[];
  // 介绍
  introduction?: string;
  // when it is a customer
  customer?: any;
}

export type SetUserPasswordParams = {
  oldPassword: string;
  newPassword: string;
};

export interface CaptchaResModel {
  code: number;
  data: string;
  id: string;
  msg?: string;
}
