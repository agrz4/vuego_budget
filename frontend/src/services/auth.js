import api from './api';

export const loginUser = async (username, password) => {
  try {
    const response = await api.post('/login', { username, password });
    localStorage.setItem('token', response.data.token);
    return response.data;
  } catch (error) {
    throw new Error(error.response?.data?.error || 'Login gagal');
  }
};

export const registerUser = async (username, password) => {
  try {
    const response = await api.post('/register', { username, password });
    return response.data;
  } catch (error) {
    throw new Error(error.response?.data?.error || 'Registrasi gagal');
  }
};

export const logoutUser = () => {
  localStorage.removeItem('token');
  // No backend call needed for JWT logout, just clear client-side token
};

export const checkAuth = async () => {
  const token = localStorage.getItem('token');
  // Simple check: if token exists, assume authenticated.
  // A more robust check would involve a backend endpoint to validate the token.
  return !!token;
};