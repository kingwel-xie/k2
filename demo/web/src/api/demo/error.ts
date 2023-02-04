import { otherHttp } from '/@/utils/http/axios';

enum Api {
  // The address does not exist
  Error = '/error',
}

/**
 * @description: Trigger ajax error
 */

export const fireErrorApi = () => otherHttp.get({ url: Api.Error });
