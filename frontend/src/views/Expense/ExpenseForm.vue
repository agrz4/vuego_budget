<template>
  <div class="p-6 bg-white rounded-lg shadow-md max-w-md mx-auto">
    <h1 class="text-3xl font-bold text-gray-800 mb-6 text-center">{{ isEdit ? 'Edit Pengeluaran' : 'Tambah Pengeluaran Baru' }}</h1>
    <form @submit.prevent="handleSubmit">
      <div class="mb-4">
        <label for="category" class="block text-gray-700 text-sm font-semibold mb-2">Kategori:</label>
        <input
          type="text"
          id="category"
          v-model="expense.category"
          class="shadow-sm appearance-none border rounded-md w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
          required
        />
      </div>
      <div class="mb-4">
        <label for="amount" class="block text-gray-700 text-sm font-semibold mb-2">Jumlah (IDR):</label>
        <input
          type="number"
          id="amount"
          v-model.number="expense.amount"
          class="shadow-sm appearance-none border rounded-md w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
          required
          min="0.01"
          step="0.01"
        />
      </div>
      <div class="mb-4">
        <label for="date" class="block text-gray-700 text-sm font-semibold mb-2">Tanggal:</label>
        <input
          type="date"
          id="date"
          v-model="expense.date"
          class="shadow-sm appearance-none border rounded-md w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
          required
        />
      </div>
      <div class="mb-6">
        <label for="description" class="block text-gray-700 text-sm font-semibold mb-2">Deskripsi (Opsional):</label>
        <textarea
          id="description"
          v-model="expense.description"
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
        {{ isEdit ? 'Perbarui Pengeluaran' : 'Tambah Pengeluaran' }}
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
import { createExpense, getExpenseByID, updateExpense } from '../../services/expense';

const route = useRoute();
const router = useRouter();
const isEdit = ref(false);
const expense = ref({
  category: '',
  amount: 0,
  description: '',
  date: new Date().toISOString().split('T')[0], // Default to today's date
});
const error = ref('');

onMounted(async () => {
  if (route.params.id) {
    isEdit.value = true;
    try {
      const response = await getExpenseByID(route.params.id);
      // Format date for input type="date"
      expense.value = {
        ...response.data.expense,
        date: new Date(response.data.expense.Date).toISOString().split('T')[0],
      };
    } catch (err) {
      error.value = err.message || 'Gagal memuat detail pengeluaran.';
    }
  }
});

const handleSubmit = async () => {
  try {
    error.value = '';
    // Backend expects ISO string for date, but date input provides YYYY-MM-DD
    // Convert back to ISO string if needed, or ensure backend handles YYYY-MM-DD
    const payload = {
      ...expense.value,
      date: new Date(expense.value.date).toISOString(), // Convert to ISO string for backend
    };

    if (isEdit.value) {
      await updateExpense(route.params.id, payload);
    } else {
      await createExpense(payload);
    }
    router.push('/expenses');
  } catch (err) {
    error.value = err.message || 'Gagal menyimpan pengeluaran.';
  }
};
</script>