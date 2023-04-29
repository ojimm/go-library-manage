import request from '@/utils/request'

// 查询BookClass列表
export function listBookClass(query) {
  return request({
    url: '/api/v1/book-class',
    method: 'get',
    params: query
  })
}

// 查询BookClass详细
export function getBookClass(id) {
  return request({
    url: '/api/v1/book-class/' + id,
    method: 'get'
  })
}

// 新增BookClass
export function addBookClass(data) {
  return request({
    url: '/api/v1/book-class',
    method: 'post',
    data: data
  })
}

// 修改BookClass
export function updateBookClass(data) {
  return request({
    url: '/api/v1/book-class/' + data.id,
    method: 'put',
    data: data
  })
}

// 删除BookClass
export function delBookClass(data) {
  return request({
    url: '/api/v1/book-class',
    method: 'delete',
    data: data
  })
}
