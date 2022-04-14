// 查询列表
export function getItems(f, query) {
  query = Object.assign(query || {}, { pageIndex: 1, pageSize: -1 })
  return f(query)
}
