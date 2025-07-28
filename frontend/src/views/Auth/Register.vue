<template>
  <div class="flex items-center justify-center min-h-[calc(100vh-120px)]">
    <div class="bg-white p-8 rounded-lg shadow-xl w-full max-w-md">
      <h2 class="text-3xl font-bold text-center text-gray-800 mb-6">Daftar</h2>
      <form @submit.prevent="handleRegister">
        <div class="mb-4">
          <label for="username" class="block text-gray-700 text-sm font-semibold mb-2">Nama Pengguna:</label>
          <input
            type="text"
            id="username"
            v-model="username"
            class="shadow-sm appearance-none border rounded-md w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            required
          />
        </div>
        <div class="mb-6">
          <label for="password" class="block text-gray-700 text-sm font-semibold mb-2">Kata Sandi:</label>
          <input
            type="password"
            id="password"
            v-model="password"
            class="shadow-sm appearance-none border rounded-md w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            required
          />
        </div>
        <div v-if="error" class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded-md relative mb-4" role="alert">
          <span class="block sm:inline">{{ error }}</span>
        </div>
        <button
          type="submit"
          class="w-full bg-green-600 hover:bg-green-700 text-white font-bold py-2 px-4 rounded-md focus:outline-none focus:shadow-outline transition duration-300"
        >
          Daftar
        </button>
        <p class="text-center text-gray-600 text-sm mt-4">
          Sudah punya akun?
          <router-link to="/login" class="text-blue-600 hover:underline">Login di sini</router-link>
        </p>
      </form>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { registerUser } from '../../services/auth';

const username = ref('');
const password = ref('');
const error = ref('');
const router = useRouter();

const handleRegister = async () => {
  try {
    error.value = '';
    await registerUser(username.value, password.value);
    window.dispatchEvent(new Event('auth-status-changed')); // Notify App.vue
    router.push('/login'); // Redirect to login after successful registration
  } catch (err) {
    error.value = err.message || 'Registrasi gagal. Silakan coba lagi.';
  }
};
</script>