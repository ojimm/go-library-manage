import request from '@/utils/request'

// 查询Lend列表
export function listLend(query) {
  return request({
    url: '/api/v1/lend',
    method: 'get',
    params: query
  })
}

// 查询Lend详细
export function getLend(id) {
  return request({
    url: '/api/v1/lend/' + id,
    method: 'get'
  })
}

// 新增Lend
export function addLend(data) {
  return request({
    url: '/api/v1/lend',
    method: 'post',
    data: data
  })
}

// 修改Lend
export function updateLend(data) {
  return request({
    url: '/api/v1/lend/' + data.id,
    method: 'put',
    data: data
  })
}

// 删除Lend
export function delLend(data) {
  return request({
    url: '/api/v1/lend',
    method: 'delete',
    data: data
  })
}
