<template>
  <div class="p-6 bg-white rounded-lg shadow-md">
    <h1 class="text-3xl font-bold text-gray-800 mb-6">Dashboard</h1>

    <div v-if="loading" class="text-center text-gray-600">Memuat data...</div>
    <div v-else-if="error" class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded-md relative mb-4" role="alert">
      <span class="block sm:inline">{{ error }}</span>
    </div>
    <div v-else>
      <div class="grid grid-cols-1 md:grid-cols-3 gap-6 mb-8">
        <div class="bg-blue-100 p-6 rounded-lg shadow-sm">
          <h2 class="text-xl font-semibold text-blue-800 mb-2">Total Anggaran</h2>
          <p class="text-3xl font-bold text-blue-900">{{ formatCurrency(summary?.total_budget ?? 0) }}</p>
        </div>
        <div class="bg-red-100 p-6 rounded-lg shadow-sm">
          <h2 class="text-xl font-semibold text-red-800 mb-2">Total Pengeluaran</h2>
          <p class="text-3xl font-bold text-red-900">{{ formatCurrency(summary.total_expense) }}</p>
        </div>
        <div :class="['p-6 rounded-lg shadow-sm', summary.remaining_budget >= 0 ? 'bg-green-100' : 'bg-orange-100']">
          <h2 :class="['text-xl font-semibold mb-2', summary.remaining_budget >= 0 ? 'text-green-800' : 'text-orange-800']">Sisa Anggaran</h2>
          <p :class="['text-3xl font-bold', summary.remaining_budget >= 0 ? 'text-green-900' : 'text-orange-900']">{{ formatCurrency(summary.remaining_budget) }}</p>
        </div>
      </div>

      <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <div class="bg-white p-6 rounded-lg shadow-md">
          <h2 class="text-xl font-semibold text-gray-800 mb-4">Ringkasan Pengeluaran per Kategori</h2>
          <ul class="space-y-2">
            <li v-for="item in summary.expenses_by_category" :key="item.category" class="flex justify-between items-center py-2 border-b border-gray-200 last:border-b-0">
              <span class="text-gray-700 font-medium">{{ item.category }}</span>
              <span class="font-semibold text-gray-900">{{ formatCurrency(item.total) }}</span>
            </li>
            <li v-if="summary.expenses_by_category.length === 0" class="text-center text-gray-500 py-4">Belum ada pengeluaran.</li>
          </ul>
        </div>
        <Chart v-if="summary.expenses_by_category.length > 0" :chartData="summary.expenses_by_category" class="h-80" />
        <div v-else class="bg-white p-6 rounded-lg shadow-md flex items-center justify-center h-80">
          <p class="text-gray-500">Tidak ada data untuk grafik pengeluaran.</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { getDashboardSummary } from '../services/budget';
import Chart from '../components/Chart.vue';

const summary = ref({
  total_budget: 0,
  total_expense: 0,
  remaining_budget: 0,
  expenses_by_category: [],
});
const loading = ref(true);
const error = ref('');

const fetchDashboardSummary = async () => {
  try {
    loading.value = true;
    error.value = '';
    const response = await getDashboardSummary();
    console.log('Dashboard summary response:', response); // Debug log
    summary.value = response;
  } catch (err) {
    console.log('Dashboard summary error:', err, err.response); // Debug log error
    error.value = (err.response && err.response.data && err.response.data.error) ? err.response.data.error : (err.message || 'Gagal memuat ringkasan dashboard.');
    summary.value = {
      total_budget: 0,
      total_expense: 0,
      remaining_budget: 0,
      expenses_by_category: [],
    };
  } finally {
    loading.value = false;
  }
};

const formatCurrency = (value) => {
  return new Intl.NumberFormat('id-ID', {
    style: 'currency',
    currency: 'IDR',
    minimumFractionDigits: 0, // No decimal places for whole numbers
    maximumFractionDigits: 0,
  }).format(value);
};

onMounted(() => {
  fetchDashboardSummary();
});
</script>