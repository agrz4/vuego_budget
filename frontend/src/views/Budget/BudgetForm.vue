<template>
  <div class="p-6 bg-white rounded-lg shadow-md max-w-md mx-auto">
    <h1 class="text-3xl font-bold text-gray-800 mb-6 text-center">{{ isEdit ? 'Edit Anggaran' : 'Tambah Anggaran Baru' }}</h1>
    <form @submit.prevent="handleSubmit">
      <div class="mb-4">
        <label for="category" class="block text-gray-700 text-sm font-semibold mb-2">Kategori:</label>
        <input
          type="text"
          id="category"
          v-model="budget.category"
          class="shadow-sm appearance-none border rounded-md w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
          required
        />
      </div>
      <div class="mb-4">
        <label for="amount" class="block text-gray-700 text-sm font-semibold mb-2">Jumlah (IDR):</label>
        <input
          type="number"
          id="amount"
          v-model.number="budget.amount"
          class="shadow-sm appearance-none border rounded-md w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
          required
          min="0.01"
          step="0.01"
        />
      </div>
      <div class="mb-6">
        <label for="description" class="block text-gray-700 text-sm font-semibold mb-2">Deskripsi (Opsional):</label>
        <textarea
          id="description"
          v-model="budget.description"
          rows="3"
          class="shadow-sm appearance-none border rounded-md w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
        ></textarea>
      </div>
      <div v-if="error" class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded-md relative mb-4" role="alert">
        <span class="block sm:inline">{{ error }}</span>
      </div>
      <button
        type="submit"
        class="w-full bg-blue-600 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded-md focus:outline-none focus:shadow-outline transition duration-300"
      >
        {{ isEdit ? 'Perbarui Anggaran' : 'Tambah Anggaran' }}
      </button>
      <button
        type="button"
        @click="router.back()"
        class="w-full mt-3 bg-gray-300 hover:bg-gray-400 text-gray-800 font-bold py-2 px-4 rounded-md focus:outline-none focus:shadow-outline transition duration-300"
      >
        Batal
      </button>
    </form>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { createBudget, getBudgetByID, updateBudget } from '../../services/budget';

const route = useRoute();
const router = useRouter();
const isEdit = ref(false);
const budget = ref({
  category: '',
  amount: 0,
  description: '',
});
const error = ref('');

onMounted(async () => {
  if (route.params.id) {
    isEdit.value = true;
    try {
      const response = await getBudgetByID(route.params.id);
      
      // Validasi response dengan lebih detail
      if (response && typeof response === 'object') {
        if (response.budget && typeof response.budget === 'object') {
          // Pastikan semua field yang diperlukan ada
          if (response.budget.category && response.budget.amount !== undefined) {
            budget.value = {
              category: response.budget.category || '',
              amount: response.budget.amount || 0,
              description: response.budget.description || ''
            };
          } else {
            error.value = 'Data anggaran tidak lengkap';
          }
        } else {
          error.value = 'Format data anggaran tidak valid';
        }
      } else {
        error.value = 'Response dari server tidak valid';
      }
    } catch (err) {
      if (err.response) {
        // Error dari axios dengan response
        if (err.response.status === 404) {
          error.value = 'Anggaran tidak ditemukan';
        } else if (err.response.status === 401) {
          error.value = 'Anda tidak memiliki akses ke anggaran ini';
        } else {
          error.value = err.response.data?.error || 'Gagal memuat detail anggaran';
        }
      } else {
        error.value = err.message || 'Gagal memuat detail anggaran';
      }
    }
  }
});

const handleSubmit = async () => {
  try {
    error.value = '';
    if (isEdit.value) {
      await updateBudget(route.params.id, budget.value);
    } else {
      await createBudget(budget.value);
    }
    router.push('/budgets');
  } catch (err) {
    error.value = err.message || 'Gagal menyimpan anggaran.';
  }
};
</script>