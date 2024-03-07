import axios from 'axios'
import { ElMessage as Message } from 'element-plus'
//import store from '@/store'

axios.defaults.timeout = 10000
//设置post的请求头
axios.defaults.headers.post['Content-Type'] = 'application/json;charset=utf-8'

// 请求拦截器
axios.interceptors.request.use(
  (config) => {
    // 每次发送请求之前判断vuex中是否存在token
    // 如果存在，则统一在http请求的header都加上token，这样后台根据token判断你的登录情况
    // 即使本地存在token，也有可能token是过期的，所以在响应拦截器中要对返回状态进行判断
    //const token = store.state.token;
    // token && (config.headers.Authorization = token);
    return config
  },
  (error) => {
    return Promise.error(error)
  }
)

// 响应拦截器
axios.interceptors.response.use(
  (response) => {
    if (response.data.code == 200) {
      return Promise.resolve(response)
    } else {
      Message.error(response.data.msg)
      return Promise.reject(response)
    }
  },
  (error) => {
    if (error.response.status) {
      switch (error.response.status) {
        // 401: 未登录
        // 未登录则跳转登录页面，并携带当前页面的路径
        // 在登录成功后返回当前页面，这一步需要在登录页操作。
        case 401:
          window.location.href = '/login'
          break
        default:
          //Message.error(error.response.data.message)
      }
    }
    Message.error({ showClose: true, message: '连接服务器失败', type: 'error' })
    return Promise.reject(error.response)
  }
)

export function get(url, params) {
  return new Promise((resolve, reject) => {
    axios
      .get(url, { params: params })
      .then((res) => {
        resolve(res.data)
      })
      .catch((res) => {
        reject(res.data)
      })
  })
}

export function del(url, params) {
  return new Promise((resolve, reject) => {
    axios
      .delete(url, { params: params })
      .then((res) => {
        resolve(res.data)
      })
      .catch((res) => {
        reject(res.data)
      })
  })
}

export function post(url, params) {
  return new Promise((resolve, reject) => {
    axios
      .post(url, params)
      .then((res) => {
        resolve(res.data)
      })
      .catch((err) => {
        reject(err.data)
      })
  })
}

export function put(url, params) {
  return new Promise((resolve, reject) => {
    axios
      .put(url, params)
      .then((res) => {
        resolve(res.data)
      })
      .catch((err) => {
        reject(err.data)
      })
  })
}

export function postfile(url, params, f) {
  return new Promise((resolve, reject) => {
    axios
      .post(url, params, {
        headers: { 'Content-Type': 'multipart/form-data' },
        onUploadProgress: (progressEvent) => {
          let complete =
            ((progressEvent.loaded / progressEvent.total) * 100) | 0
          //self.uploadMessage = '上传 ' + complete
          f(complete)
        },
      })
      .then((res) => {
        resolve(res.data)
      })
      .catch((err) => {
        reject(err.data)
      })
  })
}
