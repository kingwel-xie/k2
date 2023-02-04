import { AccountListItem, RoleListItem } from '/@/api/admin/model/system';

export interface GetProfileModel {
  apiToken?: string;
  customerCode?: string;
  user: AccountListItem;
  roles: RoleListItem[];
}
