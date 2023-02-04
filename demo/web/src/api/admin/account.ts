import { defHttp } from '/@/utils/http/axios';
import { GetProfileModel } from './model/account';
import { UploadFileParams } from '/#/axios';
import { UploadApiResult } from '/@/api/sys/model/uploadModel';

enum Api {
  ACCOUNT_INFO = '/v1/user/profile',
  UPLOAD_AVATAR = '/v1/user/avatar',
  SESSION_TIMEOUT = '/user/sessionTimeout',
  TOKEN_EXPIRED = '/user/tokenExpired',
}

// Get personal center-basic settings

export const getAccountProfileApi = () => defHttp.get<GetProfileModel>({ url: Api.ACCOUNT_INFO });

export const setAccountProfileApi = (params?: any) =>
  defHttp.post<void>({ url: Api.ACCOUNT_INFO, params });

export const sessionTimeoutApi = () => defHttp.post<void>({ url: Api.SESSION_TIMEOUT });

export const tokenExpiredApi = () => defHttp.post<void>({ url: Api.TOKEN_EXPIRED });

export function uploadAvatar(
  params: UploadFileParams,
  onUploadProgress: (progressEvent: ProgressEvent) => void,
) {
  const formData = new FormData();
  formData.append('upload[]', params.file, params.filename);
  return defHttp.post<UploadApiResult>({
    url: Api.UPLOAD_AVATAR,
    onUploadProgress,
    params: formData,
  });
}
