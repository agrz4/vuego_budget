<template>
  <div class="p-6 bg-white rounded-lg shadow-md">
    <div class="flex justify-between items-center mb-6">
      <h1 class="text-3xl font-bold text-gray-800">Daftar Anggaran</h1>
      <router-link to="/budgets/add" class="bg-blue-600 hover:bg-blue-700 text-white font-semibold py-2 px-4 rounded-md shadow-md transition duration-300">
        Tambah Anggaran
      </router-link>
    </div>

    <div v-if="loading" class="text-center text-gray-600">Memuat anggaran...</div>
    <div v-else-if="error" class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded-md relative mb-4" role="alert">
      <span class="block sm:inline">{{ error }}</span>
    </div>
    <div v-else-if="budgets.length === 0" class="text-center text-gray-500 py-8">
      Belum ada anggaran yang ditambahkan.
    </div>
    <div v-else class="overflow-x-auto">
      <table class="min-w-full bg-white border border-gray-200 rounded-lg">
        <thead class="bg-gray-50">
          <tr>
            <th class="py-3 px-6 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Kategori</th>
            <th class="py-3 px-6 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Jumlah</th>
            <th class="py-3 px-6 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Deskripsi</th>
            <th class="py-3 px-6 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Dibuat Pada</th>
            <th class="py-3 px-6 text-center text-xs font-medium text-gray-500 uppercase tracking-wider">Aksi</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="budget in budgets" :key="budget.ID">
            <td class="py-4 px-6 whitespace-nowrap text-gray-800">{{ budget.category }}</td>
            <td class="py-4 px-6 whitespace-nowrap text-gray-800">{{ formatCurrency(budget.amount) }}</td>
            <td class="py-4 px-6 whitespace-nowrap text-gray-800">{{ budget.description || '-' }}</td>
            <td class="py-4 px-6 whitespace-nowrap text-gray-800">{{ formatDate(budget.created_at || budget.CreatedAt) }}</td>
            <td class="py-4 px-6 whitespace-nowrap text-center">
              <router-link :to="`/budgets/edit/${budget.ID}`" class="text-blue-600 hover:text-blue-900 mr-3">Edit</router-link>
              <button @click="confirmDelete(budget.ID)" class="text-red-600 hover:text-red-900">Hapus</button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- Delete Confirmation Modal -->
    <div v-if="showModal" class="fixed inset-0 bg-gray-600 bg-opacity-75 flex items-center justify-center z-50">
      <div class="bg-white p-6 rounded-lg shadow-xl w-full max-w-sm">
        <h3 class="text-lg font-semibold text-gray-800 mb-4">Konfirmasi Hapus</h3>
        <p class="text-gray-700 mb-6">Apakah Anda yakin ingin menghapus anggaran ini?</p>
        <div class="flex justify-end space-x-4">
          <button @click="showModal = false" class="bg-gray-300 hover:bg-gray-400 text-gray-800 font-bold py-2 px-4 rounded-md">Batal</button>
          <button @click="deleteBudgetConfirmed" class="bg-red-600 hover:bg-red-700 text-white font-bold py-2 px-4 rounded-md">Hapus</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { getBudgets, deleteBudget } from '../../services/budget';

const budgets = ref([]);
const loading = ref(true);
const error = ref('');
const showModal = ref(false);
const budgetToDelete = ref(null);

const fetchBudgets = async () => {
  try {
    loading.value = true;
    error.value = '';
    const response = await getBudgets();
    console.log('Response getBudgets:', response); // Debug log
    budgets.value = response && response.budgets ? response.budgets : [];
  } catch (err) {
    error.value = err.message || 'Gagal memuat daftar anggaran.';
    budgets.value = [];
  } finally {
    loading.value = false;
  }
};

const confirmDelete = (id) => {
  budgetToDelete.value = id;
  showModal.value = true;
};

const deleteBudgetConfirmed = async () => {
  try {
    await deleteBudget(budgetToDelete.value);
    budgets.value = budgets.value.filter(b => b.ID !== budgetToDelete.value);
    showModal.value = false;
    budgetToDelete.value = null;
  } catch (err) {
    error.value = err.message || 'Gagal menghapus anggaran.';
  }
};

const formatCurrency = (value) => {
  return new Intl.NumberFormat('id-ID', {
    style: 'currency',
    currency: 'IDR',
    minimumFractionDigits: 0,
    maximumFractionDigits: 0,
  }).format(value);
};

const formatDate = (dateString) => {
  if (!dateString) {
    return '-';
  }
  
  try {
    const date = new Date(dateString);
    if (isNaN(date.getTime())) {
      return '-';
    }
    
    const options = { year: 'numeric', month: 'long', day: 'numeric' };
    return date.toLocaleDateString('id-ID', options);
  } catch (error) {
    console.error('Error formatting date:', dateString, error);
    return '-';
  }
};

onMounted(() => {
  fetchBudgets();
});
</script>