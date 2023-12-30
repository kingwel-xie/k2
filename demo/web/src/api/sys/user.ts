import { defHttp } from '/@/utils/http/axios';
import {
  LoginParams,
  LoginResultModel,
  GetUserInfoModel,
  SetUserPasswordParams,
  CaptchaResModel,
} from './model/userModel';

import { ErrorMessageMode } from '/#/axios';
import { GetUnreadInfo } from '/@/api/sys/model/messageModel';

enum Api {
  GetCaptcha = '/v1/captcha',
  Login = '/v1/login',
  Logout = '/v1/logout',
  GetUserInfo = '/v1/getinfo',
  SysInboxOpUnread = '/v1/inbox/unread',
  SysInboxOpRead = '/v1/inbox/read',
  UpdatePassword = '/v1/user/pwd/set',
  RegenerateToken = '/v1/user/token',
  TestRetry = '/testRetry',
}

export function getCaptcha() {
  return defHttp.get<CaptchaResModel>(
    { url: Api.GetCaptcha },
    { errorMessageMode: 'none', isTransformResponse: false },
  );
}

/**
 * @description: user login api
 */
export function loginApi(params: LoginParams, mode: ErrorMessageMode = 'modal') {
  return defHttp.post<LoginResultModel>(
    {
      url: Api.Login,
      params,
    },
    {
      // login api is kind of special
      isTransformResponse: false,
      errorMessageMode: mode,
    },
  );
}

/**
 * @description: getUserInfo
 */
export function getUserInfo() {
  return defHttp.get<GetUserInfoModel>({ url: Api.GetUserInfo }, { errorMessageMode: 'none' });
}

export function doLogout() {
  return defHttp.post({ url: Api.Logout });
}

export function updatePassword(params: SetUserPasswordParams) {
  return defHttp.put({ url: Api.UpdatePassword, params });
}

export function regenerateToken() {
  return defHttp.put({ url: Api.RegenerateToken });
}

export function testRetry() {
  return defHttp.get(
    { url: Api.TestRetry },
    {
      retryRequest: {
        isOpenRetry: true,
        count: 5,
        waitTime: 1000,
      },
    },
  );
}

/**
 * @description: getMessageUnreadApi
 */
export const getMessageUnreadApi = () => defHttp.get<GetUnreadInfo>({ url: Api.SysInboxOpUnread });

export const readMessageApi = (ids: number[], read = true) =>
  defHttp.post<any>({ url: Api.SysInboxOpRead, params: { ids, read: read } });
