import api from './api';

export const getExpenses = async (params = {}) => {
  try {
    const response = await api.get('/expenses', { params });
    return response.data;
  } catch (error) {
    throw new Error(error.response?.data?.error || 'Gagal mengambil pengeluaran');
  }
};

export const getExpenseByID = async (id) => {
  try {
    const response = await api.get(`/expenses/${id}`);
    return response.data;
  } catch (error) {
    throw new Error(error.response?.data?.error || 'Gagal mengambil pengeluaran');
  }
};

export const createExpense = async (expenseData) => {
  try {
    const response = await api.post('/expenses', expenseData);
    return response.data;
  } catch (error) {
    throw new Error(error.response?.data?.error || 'Gagal membuat pengeluaran');
  }
};

export const updateExpense = async (id, expenseData) => {
  try {
    const response = await api.put(`/expenses/${id}`, expenseData);
    return response.data;
  } catch (error) {
    throw new Error(error.response?.data?.error || 'Gagal memperbarui pengeluaran');
  }
};

export const deleteExpense = async (id) => {
  try {
    const response = await api.delete(`/expenses/${id}`);
    return response.data;
  } catch (error) {
    throw new Error(error.response?.data?.error || 'Gagal menghapus pengeluaran');
  }
};