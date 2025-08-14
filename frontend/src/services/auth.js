import api from './api';

export const loginUser = async (username, password) => {
  try {
    const response = await api.post('/login', { username, password });
    localStorage.setItem('token', response.data.token);
    if (response.data.user) {
      localStorage.setItem('name', response.data.user.name || '');
      localStorage.setItem('username', response.data.user.username || '');
    }
    return response.data;
  } catch (error) {
    throw new Error(error.response?.data?.error || 'Login gagal');
  }
};

export const registerUser = async (username, password, name) => {
  try {
    const response = await api.post('/register', { username, password, name });
    return response.data;
  } catch (error) {
    throw new Error(error.response?.data?.error || 'Registrasi gagal');
  }
};

export const logoutUser = () => {
  localStorage.removeItem('token');
  localStorage.removeItem('name');
  localStorage.removeItem('username');
  // No backend call needed for JWT logout, just clear client-side token
};

export const checkAuth = async () => {
  const token = localStorage.getItem('token');
  // Simple check: if token exists, assume authenticated.
  // A more robust check would involve a backend endpoint to validate the token.
  return !!token;
};

export const getCurrentUser = async () => {
  try {
    const response = await api.get('/me');
    const user = response.data;
    if (user) {
      localStorage.setItem('name', user.name || '');
      localStorage.setItem('username', user.username || '');
    }
    return user;
  } catch (error) {
    return null;
  }
};