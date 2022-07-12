import store from '@/store'

/**
 * @param {Array} value
 * @returns {Boolean}
 * @example see @/views/permission/directive.vue
 */
export default function checkPermission(value) {
  if (value && value instanceof Array && value.length > 0) {
    const roles = store.getters && store.getters.roles
    const permissionRoles = value

    const hasPermission = roles.some(role => {
      return permissionRoles.includes(role)
    })

    if (!hasPermission) {
      return false
    }
    return true
  } else {
    console.error(`need roles! Like v-permission="['admin','editor']"`)
    return false
  }
}

const ALL_PERMISSION = '*:*:*'

export function checkPermissionAction(permission) {
  permission = typeof permission === 'string' ? [permission] : permission

  if (permission && permission instanceof Array && permission.length > 0) {
    const permissions = store.getters && store.getters.permisaction
    return permissions.some(perm => {
      return ALL_PERMISSION === perm || permission.includes(perm)
    })
  } else {
    throw new Error(`请设置操作权限标签值`)
  }
}
