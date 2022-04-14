import {
  findVueAncestorComponent,
  focusElementUIFormField,
  getFirstElementUIFormInvalidField,
  getVueComponentProperty
} from '@/utils/hack'

// const directiveId = 'focusFirstInvalidOnValidate'
const directiveId = 'ffiov'

export default {
  install(Vue) {
    /**
     * The purpose is that the first invalid input item automatically gets the focus when the
     * 'validate' method of the form instance is called and the validation fails.
     * When using the tab panel, first jump to the corresponding tab panel
     *
     * There are two usage scenarios
     * first:
     *  <el-form v-ffiov>...</el-form>
     *
     * second:
     *  <el-form v-ffiov.tab="activeTab">
     *    <el-tabs v-model="activeTab">
     *      <el-tab-pane name="tab1">...</el-tab-pane>
     *      <el-tab-pane name="tab2">...</el-tab-pane>
     *    </el-tabs>
     *  </el-form>
     */
    Vue.directive(directiveId, {
      bind: function(el, binding, vnode) {
        if (vnode.componentOptions.tag !== 'el-form') {
          throw new Error('focus-first-invalid-on-validation must use on <el-form>')
        }
        const { expression, modifiers: { tab }} = binding
        const { componentInstance, context } = vnode

        const originalValidate = componentInstance.validate
        componentInstance.validate = (callback) => {
          originalValidate((valid, data) => {
            callback(valid, data)
            if (valid) {
              return
            }
            const invalidField = getFirstElementUIFormInvalidField(componentInstance)
            if (!invalidField) {
              return
            }
            if (tab) {
              const tabPane = findVueAncestorComponent(invalidField, 'el-tab-pane')
              if (tabPane) {
                context[expression] = getVueComponentProperty(tabPane, 'name')
              }
            }
            if (invalidField.$children && invalidField.$children.length === 2) {
              focusElementUIFormField(invalidField, context.$nextTick)
            }
          })
        }
      }
    })
  }
}

