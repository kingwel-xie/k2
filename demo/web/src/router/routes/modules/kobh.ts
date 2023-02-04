import type { AppRouteModule } from '/@/router/types';
import { LAYOUT } from '/@/router/constant';
import { t } from '/@/hooks/web/useI18n';

const dashboardKobh: AppRouteModule = {
  path: '/dashboard',
  name: 'Dashboard',
  component: LAYOUT,
  redirect: '/dashboard/index',
  meta: {
    orderNo: 10,
    icon: 'ion:grid-outline',
    title: '主页',
  },
  children: [
    {
      path: 'analysis',
      name: 'Analysis',
      component: () => import('/@/views/dashboard/analysis/index.vue'),
      meta: {
        // affix: true,
        title: t('routes.dashboard.analysis'),
        hideBreadcrumb: true,
        hideMenu: true,
      },
    },
    {
      path: 'index',
      name: 'Workbench',
      component: () => import('/@/views/dashboard/workbench/index.vue'),
      meta: {
        // affix: true,
        title: t('routes.dashboard.workbench'),
      },
    },
    {
      path: 'profile',
      name: 'ProfileIndex',
      component: () => import('/@/views/sys/profile/index.vue'),
      meta: {
        title: t('routes.basic.profile'),
      },
    },
  ],
};

export default dashboardKobh;
