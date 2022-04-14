
// 文件上传
import request from '@/utils/request'

export function uploadFile(data) {
  return request({
    url: '/api/v1/public/uploadFile',
    method: 'post',
    data: data
  })
}
