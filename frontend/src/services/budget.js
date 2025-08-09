import api from './api';

export const getBudgets = async () => {
  try {
    const response = await api.get('/budgets');
    return response.data;
  } catch (error) {
    throw new Error(error.response?.data?.error || 'Gagal mengambil anggaran');
  }
};

export const getBudgetByID = async (id) => {
  try {
    console.log('Service: Fetching budget with ID:', id);
    const response = await api.get(`/budgets/${id}`);
    console.log('Service: Raw API response:', response);
    console.log('Service: Response data:', response.data);
    return response.data;
  } catch (error) {
    console.error('Service: Error in getBudgetByID:', error);
    throw new Error(error.response?.data?.error || 'Gagal mengambil anggaran');
  }
};

export const createBudget = async (budgetData) => {
  try {
    const response = await api.post('/budgets', budgetData);
    return response.data;
  } catch (error) {
    throw new Error(error.response?.data?.error || 'Gagal membuat anggaran');
  }
};

export const updateBudget = async (id, budgetData) => {
  try {
    const response = await api.put(`/budgets/${id}`, budgetData);
    return response.data;
  } catch (error) {
    throw new Error(error.response?.data?.error || 'Gagal memperbarui anggaran');
  }
};

export const deleteBudget = async (id) => {
  try {
    const response = await api.delete(`/budgets/${id}`);
    return response.data;
  } catch (error) {
    throw new Error(error.response?.data?.error || 'Gagal menghapus anggaran');
  }
};

export const getDashboardSummary = async () => {
  try {
    const response = await api.get('/dashboard');
    return response.data;
  } catch (error) {
    throw new Error(error.response?.data?.error || 'Gagal mengambil ringkasan dashboard');
  }
};