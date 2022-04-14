import Vue from 'vue'

/**
 * get the first invalid field of the form
 * @param form Element UI form component
 * @returns {Vue} Invalid form item or null
 */
export function getFirstElementUIFormInvalidField(form) {
  return form.fields.find((f) => f.validateState === 'error')
}

/**
 * find the 'tag' ancestor of the component
 * @param component Vue component
 * @param ancestorTag ancestor vue tag
 * @returns {*|Vue}
 */
export function findVueAncestorComponent(component, ancestorTag) {
  let parent = component.$parent
  while (parent && parent.$options._componentTag !== ancestorTag) {
    parent = parent.$parent
  }
  return parent
}

/**
 * focus element ui form field
 * @param field ElementUI form item component
 * @param nextTick Vue nextTick function
 */
export function focusElementUIFormField(field, nextTick) {
  if (field.$children && field.$children.length === 2) {
    const input = field.$children[1]
    if (input && input.focus) {
      nextTick(() => {
        input.focus()
      })
    }
  }
}

/**
 * get vue component the prop value
 * @param component Vue component
 * @param name property name
 * @returns {*} property value
 */
export function getVueComponentProperty(component, name) {
  return component.$options.propsData[name]
}

/**
 * Focus on the first invalid field
 * Handle the el-tabs case: locate and switch to the tab pane and do focus
 * @param {Object} form, 传入
 * @return {*}
 */
export function focusOnInvalidField(form) {
  const invalidField = getFirstElementUIFormInvalidField(form)
  if (invalidField) {
    const tabPane = findVueAncestorComponent(invalidField, 'el-tab-pane')
    if (tabPane) {
      const pageName = getVueComponentProperty(tabPane, 'name')
      // FIXME: hardcoded, setCurrentName
      tabPane.$parent.setCurrentName(pageName)
    }
    if (invalidField.$children && invalidField.$children.length === 2) {
      focusElementUIFormField(invalidField, Vue.nextTick)
    }
  }
}
