import { otherHttp } from '/@/utils/http/axios';

enum Api {
  TREE_OPTIONS_LIST = '/tree/getDemoOptions',
}

/**
 * @description: Get sample options value
 */
export const treeOptionsListApi = (params?: Recordable) =>
  otherHttp.get<Recordable[]>({ url: Api.TREE_OPTIONS_LIST, params });
