<template>
  <div class="min-h-screen bg-gray-100 font-inter">
    <!-- Use the sophisticated Navbar component -->
    <Navbar />

    <main class="container mx-auto p-4 mt-4">
      <router-view />
    </main>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { checkAuth, logoutUser } from './services/auth';
import Navbar from './components/Navbar.vue';

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