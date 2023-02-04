import type { AppRouteModule } from '/@/router/types';

import { getParentLayout, LAYOUT } from '/@/router/constant';
import { t } from '/@/hooks/web/useI18n';

const comp: AppRouteModule = {
  path: '/comp',
  name: 'Comp',
  component: LAYOUT,
  redirect: '/comp/basic',
  meta: {
    orderNo: 30,
    icon: 'ion:layers-outline',
    title: t('routes.admin.comp.comp'),
  },

  children: [
    {
      path: 'basic',
      name: 'BasicDemo',
      component: () => import('/@/views/demo/comp/button/index.vue'),
      meta: {
        title: t('routes.admin.comp.basic'),
      },
    },

    {
      path: 'form',
      name: 'FormDemo',
      redirect: '/comp/form/basic',
      component: getParentLayout('FormDemo'),
      meta: {
        // icon: 'mdi:form-select',
        title: t('routes.admin.form.form'),
      },
      children: [
        {
          path: 'basic',
          name: 'FormBasicDemo',
          component: () => import('/@/views/demo/form/index.vue'),
          meta: {
            title: t('routes.admin.form.basic'),
          },
        },
        {
          path: 'useForm',
          name: 'UseFormDemo',
          component: () => import('/@/views/demo/form/UseForm.vue'),
          meta: {
            title: t('routes.admin.form.useForm'),
          },
        },
        {
          path: 'refForm',
          name: 'RefFormDemo',
          component: () => import('/@/views/demo/form/RefForm.vue'),
          meta: {
            title: t('routes.admin.form.refForm'),
          },
        },
        {
          path: 'advancedForm',
          name: 'AdvancedFormDemo',
          component: () => import('/@/views/demo/form/AdvancedForm.vue'),
          meta: {
            title: t('routes.admin.form.advancedForm'),
          },
        },
        {
          path: 'ruleForm',
          name: 'RuleFormDemo',
          component: () => import('/@/views/demo/form/RuleForm.vue'),
          meta: {
            title: t('routes.admin.form.ruleForm'),
          },
        },
        {
          path: 'dynamicForm',
          name: 'DynamicFormDemo',
          component: () => import('/@/views/demo/form/DynamicForm.vue'),
          meta: {
            title: t('routes.admin.form.dynamicForm'),
          },
        },
        {
          path: 'customerForm',
          name: 'CustomerFormDemo',
          component: () => import('/@/views/demo/form/CustomerForm.vue'),
          meta: {
            title: t('routes.admin.form.customerForm'),
          },
        },
        {
          path: 'appendForm',
          name: 'appendFormDemo',
          component: () => import('/@/views/demo/form/AppendForm.vue'),
          meta: {
            title: t('routes.admin.form.appendForm'),
          },
        },
        {
          path: 'tabsForm',
          name: 'tabsFormDemo',
          component: () => import('/@/views/demo/form/TabsForm.vue'),
          meta: {
            title: t('routes.admin.form.tabsForm'),
          },
        },
      ],
    },
    {
      path: 'table',
      name: 'TableDemo',
      redirect: '/comp/table/basic',
      component: getParentLayout('TableDemo'),
      meta: {
        // icon: 'carbon:table-split',
        title: t('routes.admin.table.table'),
      },

      children: [
        {
          path: 'basic',
          name: 'TableBasicDemo',
          component: () => import('/@/views/demo/table/Basic.vue'),
          meta: {
            title: t('routes.admin.table.basic'),
          },
        },
        {
          path: 'treeTable',
          name: 'TreeTableDemo',
          component: () => import('/@/views/demo/table/TreeTable.vue'),
          meta: {
            title: t('routes.admin.table.treeTable'),
          },
        },
        {
          path: 'fetchTable',
          name: 'FetchTableDemo',
          component: () => import('/@/views/demo/table/FetchTable.vue'),
          meta: {
            title: t('routes.admin.table.fetchTable'),
          },
        },
        {
          path: 'fixedColumn',
          name: 'FixedColumnDemo',
          component: () => import('/@/views/demo/table/FixedColumn.vue'),
          meta: {
            title: t('routes.admin.table.fixedColumn'),
          },
        },
        {
          path: 'customerCell',
          name: 'CustomerCellDemo',
          component: () => import('/@/views/demo/table/CustomerCell.vue'),
          meta: {
            title: t('routes.admin.table.customerCell'),
          },
        },
        {
          path: 'formTable',
          name: 'FormTableDemo',
          component: () => import('/@/views/demo/table/FormTable.vue'),
          meta: {
            title: t('routes.admin.table.formTable'),
          },
        },
        {
          path: 'useTable',
          name: 'UseTableDemo',
          component: () => import('/@/views/demo/table/UseTable.vue'),
          meta: {
            title: t('routes.admin.table.useTable'),
          },
        },
        {
          path: 'refTable',
          name: 'RefTableDemo',
          component: () => import('/@/views/demo/table/RefTable.vue'),
          meta: {
            title: t('routes.admin.table.refTable'),
          },
        },
        {
          path: 'multipleHeader',
          name: 'MultipleHeaderDemo',
          component: () => import('/@/views/demo/table/MultipleHeader.vue'),
          meta: {
            title: t('routes.admin.table.multipleHeader'),
          },
        },
        {
          path: 'mergeHeader',
          name: 'MergeHeaderDemo',
          component: () => import('/@/views/demo/table/MergeHeader.vue'),
          meta: {
            title: t('routes.admin.table.mergeHeader'),
          },
        },
        {
          path: 'expandTable',
          name: 'ExpandTableDemo',
          component: () => import('/@/views/demo/table/ExpandTable.vue'),
          meta: {
            title: t('routes.admin.table.expandTable'),
          },
        },
        {
          path: 'fixedHeight',
          name: 'FixedHeightDemo',
          component: () => import('/@/views/demo/table/FixedHeight.vue'),
          meta: {
            title: t('routes.admin.table.fixedHeight'),
          },
        },
        {
          path: 'footerTable',
          name: 'FooterTableDemo',
          component: () => import('/@/views/demo/table/FooterTable.vue'),
          meta: {
            title: t('routes.admin.table.footerTable'),
          },
        },
        {
          path: 'editCellTable',
          name: 'EditCellTableDemo',
          component: () => import('/@/views/demo/table/EditCellTable.vue'),
          meta: {
            title: t('routes.admin.table.editCellTable'),
          },
        },
        {
          path: 'editRowTable',
          name: 'EditRowTableDemo',
          component: () => import('/@/views/demo/table/EditRowTable.vue'),
          meta: {
            title: t('routes.admin.table.editRowTable'),
          },
        },
        {
          path: 'authColumn',
          name: 'AuthColumnDemo',
          component: () => import('/@/views/demo/table/AuthColumn.vue'),
          meta: {
            title: t('routes.admin.table.authColumn'),
          },
        },
        {
          path: 'resizeParentHeightTable',
          name: 'ResizeParentHeightTable',
          component: () => import('/@/views/demo/table/ResizeParentHeightTable.vue'),
          meta: {
            title: t('routes.admin.table.resizeParentHeightTable'),
          },
        },
      ],
    },
    {
      path: 'transition',
      name: 'transitionDemo',
      component: () => import('/@/views/demo/comp/transition/index.vue'),
      meta: {
        title: t('routes.admin.comp.transition'),
      },
    },
    {
      path: 'cropper',
      name: 'CropperDemo',
      component: () => import('/@/views/demo/comp/cropper/index.vue'),
      meta: {
        title: t('routes.admin.comp.cropperImage'),
      },
    },

    {
      path: 'timestamp',
      name: 'TimeDemo',
      component: () => import('/@/views/demo/comp/time/index.vue'),
      meta: {
        title: t('routes.admin.comp.time'),
      },
    },
    {
      path: 'countTo',
      name: 'CountTo',
      component: () => import('/@/views/demo/comp/count-to/index.vue'),
      meta: {
        title: t('routes.admin.comp.countTo'),
      },
    },
    {
      path: 'tree',
      name: 'TreeDemo',
      redirect: '/comp/tree/basic',
      component: getParentLayout('TreeDemo'),
      meta: {
        // icon: 'clarity:tree-view-line',
        title: t('routes.admin.comp.tree'),
      },
      children: [
        {
          path: 'basic',
          name: 'BasicTreeDemo',
          component: () => import('/@/views/demo/tree/index.vue'),
          meta: {
            title: t('routes.admin.comp.treeBasic'),
          },
        },
        {
          path: 'editTree',
          name: 'EditTreeDemo',
          component: () => import('/@/views/demo/tree/EditTree.vue'),
          meta: {
            title: t('routes.admin.comp.editTree'),
          },
        },
        {
          path: 'actionTree',
          name: 'ActionTreeDemo',
          component: () => import('/@/views/demo/tree/ActionTree.vue'),
          meta: {
            title: t('routes.admin.comp.actionTree'),
          },
        },
      ],
    },
    {
      path: 'editor',
      name: 'EditorDemo',
      redirect: '/comp/editor/markdown',
      component: getParentLayout('EditorDemo'),
      meta: {
        // icon: 'carbon:table-split',
        title: t('routes.admin.editor.editor'),
      },
      children: [
        {
          path: 'json',
          component: () => import('/@/views/demo/editor/json/index.vue'),
          name: 'JsonEditorDemo',
          meta: {
            title: t('routes.admin.editor.jsonEditor'),
          },
        },
        {
          path: 'markdown',
          component: getParentLayout('MarkdownDemo'),
          name: 'MarkdownDemo',
          meta: {
            title: t('routes.admin.editor.markdown'),
          },
          redirect: '/comp/editor/markdown/index',
          children: [
            {
              path: 'index',
              name: 'MarkDownBasicDemo',
              component: () => import('/@/views/demo/editor/markdown/index.vue'),
              meta: {
                title: t('routes.admin.editor.tinymceBasic'),
              },
            },
            {
              path: 'editor',
              name: 'MarkDownFormDemo',
              component: () => import('/@/views/demo/editor/markdown/Editor.vue'),
              meta: {
                title: t('routes.admin.editor.tinymceForm'),
              },
            },
          ],
        },

        {
          path: 'tinymce',
          component: getParentLayout('TinymceDemo'),
          name: 'TinymceDemo',
          meta: {
            title: t('routes.admin.editor.tinymce'),
          },
          redirect: '/comp/editor/tinymce/index',
          children: [
            {
              path: 'index',
              name: 'TinymceBasicDemo',
              component: () => import('/@/views/demo/editor/tinymce/index.vue'),
              meta: {
                title: t('routes.admin.editor.tinymceBasic'),
              },
            },
            {
              path: 'editor',
              name: 'TinymceFormDemo',
              component: () => import('/@/views/demo/editor/tinymce/Editor.vue'),
              meta: {
                title: t('routes.admin.editor.tinymceForm'),
              },
            },
          ],
        },
      ],
    },
    {
      path: 'scroll',
      name: 'ScrollDemo',
      redirect: '/comp/scroll/basic',
      component: getParentLayout('ScrollDemo'),
      meta: {
        title: t('routes.admin.comp.scroll'),
      },
      children: [
        {
          path: 'basic',
          name: 'BasicScrollDemo',
          component: () => import('/@/views/demo/comp/scroll/index.vue'),
          meta: {
            title: t('routes.admin.comp.scrollBasic'),
          },
        },
        {
          path: 'action',
          name: 'ActionScrollDemo',
          component: () => import('/@/views/demo/comp/scroll/Action.vue'),
          meta: {
            title: t('routes.admin.comp.scrollAction'),
          },
        },
        {
          path: 'virtualScroll',
          name: 'VirtualScrollDemo',
          component: () => import('/@/views/demo/comp/scroll/VirtualScroll.vue'),
          meta: {
            title: t('routes.admin.comp.virtualScroll'),
          },
        },
      ],
    },

    {
      path: 'modal',
      name: 'ModalDemo',
      component: () => import('/@/views/demo/comp/modal/index.vue'),
      meta: {
        title: t('routes.admin.comp.modal'),
      },
    },
    {
      path: 'drawer',
      name: 'DrawerDemo',
      component: () => import('/@/views/demo/comp/drawer/index.vue'),
      meta: {
        title: t('routes.admin.comp.drawer'),
      },
    },
    {
      path: 'desc',
      name: 'DescDemo',
      component: () => import('/@/views/demo/comp/desc/index.vue'),
      meta: {
        title: t('routes.admin.comp.desc'),
      },
    },

    {
      path: 'lazy',
      name: 'LazyDemo',
      component: getParentLayout('LazyDemo'),
      redirect: '/comp/lazy/basic',
      meta: {
        title: t('routes.admin.comp.lazy'),
      },
      children: [
        {
          path: 'basic',
          name: 'BasicLazyDemo',
          component: () => import('/@/views/demo/comp/lazy/index.vue'),
          meta: {
            title: t('routes.admin.comp.lazyBasic'),
          },
        },
        {
          path: 'transition',
          name: 'BasicTransitionDemo',
          component: () => import('/@/views/demo/comp/lazy/Transition.vue'),
          meta: {
            title: t('routes.admin.comp.lazyTransition'),
          },
        },
      ],
    },
    {
      path: 'verify',
      name: 'VerifyDemo',
      component: getParentLayout('VerifyDemo'),
      redirect: '/comp/verify/drag',
      meta: {
        title: t('routes.admin.comp.verify'),
      },
      children: [
        {
          path: 'drag',
          name: 'VerifyDragDemo',
          component: () => import('/@/views/demo/comp/verify/index.vue'),
          meta: {
            title: t('routes.admin.comp.verifyDrag'),
          },
        },
        {
          path: 'rotate',
          name: 'VerifyRotateDemo',
          component: () => import('/@/views/demo/comp/verify/Rotate.vue'),
          meta: {
            title: t('routes.admin.comp.verifyRotate'),
          },
        },
      ],
    },
    //

    {
      path: 'qrcode',
      name: 'QrCodeDemo',
      component: () => import('/@/views/demo/comp/qrcode/index.vue'),
      meta: {
        title: t('routes.admin.comp.qrcode'),
      },
    },
    {
      path: 'strength-meter',
      name: 'StrengthMeterDemo',
      component: () => import('/@/views/demo/comp/strength-meter/index.vue'),
      meta: {
        title: t('routes.admin.comp.strength'),
      },
    },
    {
      path: 'upload',
      name: 'UploadDemo',
      component: () => import('/@/views/demo/comp/upload/index.vue'),
      meta: {
        title: t('routes.admin.comp.upload'),
      },
    },
    {
      path: 'loading',
      name: 'LoadingDemo',
      component: () => import('/@/views/demo/comp/loading/index.vue'),
      meta: {
        title: t('routes.admin.comp.loading'),
      },
    },
    {
      path: 'cardList',
      name: 'CardListDemo',
      component: () => import('/@/views/demo/comp/card-list/index.vue'),
      meta: {
        title: t('routes.admin.comp.cardList'),
      },
    },
  ],
};

export default comp;
