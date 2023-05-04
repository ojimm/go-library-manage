import request from '@/utils/request'

// 查询Render列表
export function listRender(query) {
  return request({
    url: '/api/v1/render',
    method: 'get',
    params: query
  })
}

// 查询Render详细
export function getRender(id) {
  return request({
    url: '/api/v1/render/' + id,
    method: 'get'
  })
}

// 新增Render
export function addRender(data) {
  return request({
    url: '/api/v1/render',
    method: 'post',
    data: data
  })
}

// 修改Render
export function updateRender(data) {
  return request({
    url: '/api/v1/render/' + data.id,
    method: 'put',
    data: data
  })
}

// 删除Render
export function delRender(data) {
  return request({
    url: '/api/v1/render',
    method: 'delete',
    data: data
  })
}
