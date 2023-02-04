import type { AppRouteModule } from '/@/router/types';

import { LAYOUT } from '/@/router/constant';
// import { t } from '/@/hooks/web/useI18n';

const data: AppRouteModule = {
  path: '/data',
  name: 'CommonData',
  component: LAYOUT,
  redirect: '/data/country',
  meta: {
    icon: 'ant-design:group-outlined',
    title: '公共数据',
    orderNo: 301,
  },
  children: [
    {
      path: 'country',
      name: 'TbxCountry',
      meta: {
        title: '国家地区',
      },
      component: () => import('/@/views/kobh/tbx-country/index.vue'),
    },
  ],
};

export default data;
