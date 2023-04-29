import request from '@/utils/request'

// 查询Book列表
export function listBook(query) {
  return request({
    url: '/api/v1/book',
    method: 'get',
    params: query
  })
}

// 查询Book详细
export function getBook(id) {
  return request({
    url: '/api/v1/book/' + id,
    method: 'get'
  })
}

// 新增Book
export function addBook(data) {
  return request({
    url: '/api/v1/book',
    method: 'post',
    data: data
  })
}

// 修改Book
export function updateBook(data) {
  return request({
    url: '/api/v1/book/' + data.id,
    method: 'put',
    data: data
  })
}

// 删除Book
export function delBook(data) {
  return request({
    url: '/api/v1/book',
    method: 'delete',
    data: data
  })
}
