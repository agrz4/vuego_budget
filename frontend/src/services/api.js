import axios from 'axios';

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080/api';

const api = axios.create({
  baseURL: API_BASE_URL,
  withCredentials: true, // Send cookies (for JWT if stored in cookie)
});

// Interceptor untuk menambahkan token JWT ke header Authorization
api.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem('token');
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

// Interceptor untuk menangani respons error (misalnya, token kadaluarsa)
api.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error.response && error.response.status === 401) {
      // Token tidak valid atau kadaluarsa, arahkan ke halaman login
      localStorage.removeItem('token');
      window.dispatchEvent(new Event('auth-status-changed')); // Notify App.vue
      // Using window.location.href to force a full reload and clear state
      window.location.href = '/login';
    }
    return Promise.reject(error);
  }
);

export default api;