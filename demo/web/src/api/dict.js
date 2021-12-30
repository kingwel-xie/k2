import request from '@/utils/request'

// getAllDict 获取字典数据
export function getAllDict() {
  return request({
    url: '/api/v1/all-dict',
    method: 'get'
  })
}
