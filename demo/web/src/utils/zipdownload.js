import axios from 'axios'
import { getToken } from '@/utils/auth'

const mimeMap = {
  xlsx: 'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet',
  zip: 'application/zip'
}

const baseUrl = process.env.VUE_APP_BASE_API
export function downLoadZip(str, filename) {
  var url = baseUrl + str
  axios({
    method: 'get',
    url: url,
    responseType: 'blob',
    headers: { 'Authorization': 'Bearer ' + getToken() }
  }).then(res => {
    resolveBlob(res, mimeMap.zip)
  })
}

export function downLoadFile(url, blank, filename) {
  const aLink = document.createElement('a')
  aLink.href = url
  if (blank) {
    aLink.target = '_blank'
  }
  if (filename) {
    aLink.setAttribute('download', filename) // 设置下载文件名称
  }
  document.body.appendChild(aLink)
  aLink.click()
  document.body.removeChild(aLink)
}
/**
 * 解析blob响应内容并下载
 * @param {*} res blob响应内容
 * @param {String} mimeType MIME类型
 */
export function resolveBlob(res, mimeType) {
  var blob = new Blob([res.data], { type: mimeType })
  // //从response的headers中获取filename, 后端response.setHeader("Content-disposition", "attachment; filename=xxxx.docx") 设置的文件名;
  var patt = new RegExp('filename=([^;]+\\.[^\\.;]+);*')
  var contentDisposition = decodeURI(res.headers['content-disposition'])
  var result = patt.exec(contentDisposition)
  var fileName = result[1]

  fileName = fileName.replace(/\"/g, '')
  const url = URL.createObjectURL(blob)

  downLoadFile(url, false, fileName)
}

export function InlineDownloadPdf(str) {
  const url = baseUrl + str
  return axios({
    method: 'get',
    url: url,
    responseType: 'blob',
    headers: { 'Authorization': 'Bearer ' + getToken() }
  }).then(res => {
    const blob = new Blob([res.data], { type: 'application/pdf' })
    const url = URL.createObjectURL(blob)
    downLoadFile(url, true)
  })
}
