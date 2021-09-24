import request from '@/utils/request'

// 角色数据权限
export function dataScope(data) {
  return request({
    url: '/api/v1/role-datascope',
    method: 'put',
    data: data
  })
}

export function getRoutes() {
  return request({
    url: '/api/v1/menu-role',
    method: 'get'
  })
}
