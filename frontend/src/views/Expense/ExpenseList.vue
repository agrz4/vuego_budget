<template>
  <div class="p-6 bg-white rounded-lg shadow-md">
    <div class="flex flex-col md:flex-row justify-between items-center mb-6">
      <h1 class="text-3xl font-bold text-gray-800 mb-4 md:mb-0">Daftar Pengeluaran</h1>
      <router-link to="/expenses/add" class="bg-blue-600 hover:bg-blue-700 text-white font-semibold py-2 px-4 rounded-md shadow-md transition duration-300">
        Tambah Pengeluaran
      </router-link>
    </div>

    <div class="mb-6 p-4 bg-gray-50 rounded-lg shadow-sm">
      <h2 class="text-xl font-semibold text-gray-700 mb-3">Filter & Pencarian</h2>
      <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
        <div>
          <label for="filterCategory" class="block text-gray-700 text-sm font-semibold mb-2">Kategori:</label>
          <input
            type="text"
            id="filterCategory"
            v-model="filter.category"
            @input="debouncedFetchExpenses"
            class="shadow-sm appearance-none border rounded-md w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            placeholder="Filter berdasarkan kategori"
          />
        </div>
        <div>
          <label for="startDate" class="block text-gray-700 text-sm font-semibold mb-2">Dari Tanggal:</label>
          <input
            type="date"
            id="startDate"
            v-model="filter.startDate"
            @change="fetchExpenses"
            class="shadow-sm appearance-none border rounded-md w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
          />
        </div>
        <div>
          <label for="endDate" class="block text-gray-700 text-sm font-semibold mb-2">Sampai Tanggal:</label>
          <input
            type="date"
            id="endDate"
            v-model="filter.endDate"
            @change="fetchExpenses"
            class="shadow-sm appearance-none border rounded-md w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
          />
        </div>
        <div class="md:col-span-3">
          <label for="search" class="block text-gray-700 text-sm font-semibold mb-2">Pencarian (Deskripsi/Kategori):</label>
          <input
            type="text"
            id="search"
            v-model="filter.search"
            @input="debouncedFetchExpenses"
            class="shadow-sm appearance-none border rounded-md w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            placeholder="Cari berdasarkan deskripsi atau kategori"
          />
        </div>
        <div class="md:col-span-3 flex justify-end">
          <button @click="resetFilters" class="bg-gray-200 hover:bg-gray-300 text-gray-800 font-bold py-2 px-4 rounded-md shadow-sm transition duration-300">
            Reset Filter
          </button>
        </div>
      </div>
    </div>

    <div v-if="loading" class="text-center text-gray-600">Memuat pengeluaran...</div>
    <div v-else-if="error" class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded-md relative mb-4" role="alert">
      <span class="block sm:inline">{{ error }}</span>
    </div>
    <div v-else-if="expenses.length === 0" class="text-center text-gray-500 py-8">
      Belum ada pengeluaran yang ditambahkan atau tidak ada pengeluaran yang cocok dengan filter Anda.
    </div>
    <div v-else class="overflow-x-auto">
      <table class="min-w-full bg-white border border-gray-200 rounded-lg">
        <thead class="bg-gray-50">
          <tr>
            <th class="py-3 px-6 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Tanggal</th>
            <th class="py-3 px-6 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Kategori</th>
            <th class="py-3 px-6 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Jumlah</th>
            <th class="py-3 px-6 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Deskripsi</th>
            <th class="py-3 px-6 text-center text-xs font-medium text-gray-500 uppercase tracking-wider">Aksi</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-200">
          <tr v-for="expense in expenses" :key="expense.ID">
            <td class="py-4 px-6 whitespace-nowrap text-gray-800">{{ formatDate(expense.date || expense.Date) }}</td>
            <td class="py-4 px-6 whitespace-nowrap text-gray-800">{{ expense.category }}</td>
            <td class="py-4 px-6 whitespace-nowrap text-gray-800">{{ formatCurrency(expense.amount) }}</td>
            <td class="py-4 px-6 whitespace-nowrap text-gray-800">{{ expense.description || '-' }}</td>
            <td class="py-4 px-6 whitespace-nowrap text-center">
              <router-link :to="`/expenses/edit/${expense.ID}`" class="text-blue-600 hover:text-blue-900 mr-3">Edit</router-link>
              <button @click="confirmDelete(expense.ID)" class="text-red-600 hover:text-red-900">Hapus</button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- Delete Confirmation Modal -->
    <div v-if="showModal" class="fixed inset-0 bg-gray-600 bg-opacity-75 flex items-center justify-center z-50">
      <div class="bg-white p-6 rounded-lg shadow-xl w-full max-w-sm">
        <h3 class="text-lg font-semibold text-gray-800 mb-4">Konfirmasi Hapus</h3>
        <p class="text-gray-700 mb-6">Apakah Anda yakin ingin menghapus pengeluaran ini?</p>
        <div class="flex justify-end space-x-4">
          <button @click="showModal = false" class="bg-gray-300 hover:bg-gray-400 text-gray-800 font-bold py-2 px-4 rounded-md">Batal</button>
          <button @click="deleteExpenseConfirmed" class="bg-red-600 hover:bg-red-700 text-white font-bold py-2 px-4 rounded-md">Hapus</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { getExpenses, deleteExpense } from '../../services/expense';
import { debounce } from 'lodash'; // You'll need to install lodash: npm install lodash

const expenses = ref([]);
const loading = ref(true);
const error = ref('');
const showModal = ref(false);
const expenseToDelete = ref(null);

const filter = ref({
  category: '',
  startDate: '',
  endDate: '',
  search: '',
});

const fetchExpenses = async () => {
  try {
    loading.value = true;
    error.value = '';
    const params = {
      category: filter.value.category,
      start_date: filter.value.startDate,
      end_date: filter.value.endDate,
      search: filter.value.search,
    };
    const response = await getExpenses(params);
    console.log('Response getExpenses:', response); // Debug log
    expenses.value = response && response.expenses ? response.expenses : [];
  } catch (err) {
    error.value = err.message || 'Gagal memuat daftar pengeluaran.';
    expenses.value = [];
  } finally {
    loading.value = false;
  }
};
// Debounce the fetchExpenses call for search and category input
const debouncedFetchExpenses = debounce(fetchExpenses, 500);

const resetFilters = () => {
  filter.value = {
    category: '',
    startDate: '',
    endDate: '',
    search: '',
  };
  fetchExpenses();
};

const confirmDelete = (id) => {
  expenseToDelete.value = id;
  showModal.value = true;
};

const deleteExpenseConfirmed = async () => {
  try {
    await deleteExpense(expenseToDelete.value);
    expenses.value = expenses.value.filter(e => e.ID !== expenseToDelete.value);
    showModal.value = false;
    expenseToDelete.value = null;
  } catch (err) {
    error.value = err.message || 'Gagal menghapus pengeluaran.';
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
  fetchExpenses();
});
</script>