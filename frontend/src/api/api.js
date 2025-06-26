// frontend/src/api/api.js
import axios from 'axios'

const api = axios.create({
  baseURL: 'http://localhost:8080/api',
})

// Injects JWT if it exists
api.interceptors.request.use(config => {
  const token = localStorage.getItem('token')
  if (token && config.headers) {
    config.headers['Authorization'] = `Bearer ${token}`
  }
  return config
})

export default api