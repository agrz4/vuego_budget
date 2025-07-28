<template>
  <div class="min-h-screen bg-gray-100 font-inter">
    <nav class="bg-gradient-to-r from-blue-600 to-blue-800 p-4 shadow-md">
      <div class="container mx-auto flex justify-between items-center">
        <router-link to="/" class="text-white text-2xl font-bold rounded-md px-3 py-1 hover:bg-blue-700 transition duration-300">
          BudgetApp
        </router-link>
        <div class="space-x-4">
          <template v-if="isAuthenticated">
            <router-link to="/dashboard" class="text-white hover:text-blue-200 transition duration-300 rounded-md px-3 py-1 hover:bg-blue-700">Dashboard</router-link>
            <router-link to="/budgets" class="text-white hover:text-blue-200 transition duration-300 rounded-md px-3 py-1 hover:bg-blue-700">Anggaran</router-link>
            <router-link to="/expenses" class="text-white hover:text-blue-200 transition duration-300 rounded-md px-3 py-1 hover:bg-blue-700">Pengeluaran</router-link>
            <button @click="logout" class="bg-red-500 hover:bg-red-600 text-white font-semibold py-2 px-4 rounded-md shadow-md transition duration-300">
              Logout
            </button>
          </template>
          <template v-else>
            <router-link to="/login" class="text-white hover:text-blue-200 transition duration-300 rounded-md px-3 py-1 hover:bg-blue-700">Login</router-link>
            <router-link to="/register" class="text-white hover:text-blue-200 transition duration-300 rounded-md px-3 py-1 hover:bg-blue-700">Register</router-link>
          </template>
        </div>
      </div>
    </nav>

    <main class="container mx-auto p-4 mt-4">
      <router-view />
    </main>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { checkAuth, logoutUser } from './services/auth';

const router = useRouter();
const isAuthenticated = ref(false);

const checkAuthenticationStatus = async () => {
  isAuthenticated.value = await checkAuth();
};

const logout = async () => {
  await logoutUser();
  isAuthenticated.value = false;
  router.push('/login');
};

onMounted(() => {
  checkAuthenticationStatus();
  // Listen for custom event after login/register to update auth status
  window.addEventListener('auth-status-changed', checkAuthenticationStatus);
});
</script>