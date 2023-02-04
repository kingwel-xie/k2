import type { AppRouteModule } from '/@/router/types';

import { getParentLayout, LAYOUT } from '/@/router/constant';
import { t } from '/@/hooks/web/useI18n';

const admin: AppRouteModule = {
  path: '/admin',
  name: 'Admin',
  component: LAYOUT,
  redirect: '/admin/user',
  meta: {
    icon: 'ant-design:setting-outline',
    title: '系统管理',
    orderNo: 201,
  },
  children: [
    {
      path: 'user',
      name: 'SysUserManage',
      meta: {
        title: t('routes.admin.system.account'),
        ignoreKeepAlive: false,
      },
      component: () => import('/@/views/admin/sys-user/index.vue'),
    },
    {
      path: 'detail/:id',
      name: 'UserDetail',
      meta: {
        hideMenu: true,
        title: t('routes.admin.system.account_detail'),
        ignoreKeepAlive: true,
        showMenu: false,
        currentActiveMenu: '/system/account',
      },
      component: () => import('/@/views/admin/sys-user/AccountDetail.vue'),
    },
    {
      path: 'role',
      name: 'SysRoleManage',
      meta: {
        title: t('routes.admin.system.role'),
        ignoreKeepAlive: false,
      },
      component: () => import('/@/views/admin/sys-role/index.vue'),
    },

    {
      path: 'menu',
      name: 'SysMenuManage',
      meta: {
        title: t('routes.admin.system.menu'),
        ignoreKeepAlive: false,
      },
      component: () => import('/@/views/admin/sys-menu/index.vue'),
    },
    {
      path: 'dept',
      name: 'SysDeptManage',
      meta: {
        title: t('routes.admin.system.dept'),
        ignoreKeepAlive: true,
      },
      component: () => import('/@/views/admin/sys-dept/index.vue'),
    },
    {
      path: 'changePassword',
      name: 'ChangePassword',
      meta: {
        title: t('routes.admin.system.password'),
        ignoreKeepAlive: true,
      },
      component: () => import('/@/views/admin/password/index.vue'),
    },
    {
      path: 'dict',
      name: 'SysDictTypeManage',
      component: () => import('/@/views/admin/dict/index.vue'),
      meta: {
        title: t('routes.admin.system.dict'),
      },
    },
    {
      path: 'data/:id',
      name: 'SysDictDataManage',
      meta: {
        hideMenu: true,
        title: t('routes.admin.system.dictData'),
        ignoreKeepAlive: true,
        currentActiveMenu: '/admin/dict',
      },
      component: () => import('/@/views/admin/dict/data.vue'),
    },
    {
      path: 'api',
      name: 'SysApi',
      component: () => import('/@/views/admin/sys-api/index.vue'),
      meta: {
        title: t('routes.admin.system.api'),
      },
    },
    {
      path: 'log',
      name: 'SysLog',
      component: getParentLayout('SysLog'),
      meta: {
        icon: 'ant-design:hdd-outline',
        title: t('routes.admin.system.log'),
      },
      redirect: '/admin/log/login',
      children: [
        {
          path: 'login',
          name: 'SysLoginLogManage',
          component: () => import('/@/views/admin/sys-login-log/index.vue'),
          meta: {
            title: t('routes.admin.system.loginLog'),
          },
        },
        {
          path: 'operation',
          name: 'SysOperLogManage',
          component: () => import('/@/views/admin/sys-oper-log/index.vue'),
          meta: {
            title: t('routes.admin.system.operLog'),
          },
        },
      ],
    },
  ],
};

export default admin;
