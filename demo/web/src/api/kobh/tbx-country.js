import request from '@/utils/request'

// 查询TbxCountry列表
export function listTbxCountry(query) {
  return request({
    url: '/api/v1/country',
    method: 'get',
    params: query
  })
}

// 查询TbxCountry详细
export function getTbxCountry(code) {
  return request({
    url: '/api/v1/country/' + code,
    method: 'get'
  })
}

// 新增TbxCountry
export function addTbxCountry(data) {
  return request({
    url: '/api/v1/country',
    method: 'post',
    data: data
  })
}

// 修改TbxCountry
export function updateTbxCountry(data) {
  return request({
    url: '/api/v1/country/' + data.code,
    method: 'put',
    data: data
  })
}

// 删除TbxCountry
export function delTbxCountry(data) {
  return request({
    url: '/api/v1/country',
    method: 'delete',
    data: data
  })
}

