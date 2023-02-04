import type { RouteMeta } from 'vue-router';
export interface RouteItem {
  menuId: number;
  menuName: string;
  title: string;
  path: string;
  component: string;
  sort: number;
  visible: string;
  noCache: boolean;
  icon: string;
  iconAntd: string;
  action: string;
  meta: RouteMeta;
  redirect?: string;
  permission: string;
  children?: RouteItem[];
}

/**
 * @description: Get menu return value
 */
export type getMenuListResultModel = RouteItem[];
