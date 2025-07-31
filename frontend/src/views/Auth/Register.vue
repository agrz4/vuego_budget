<template>
  <div class="min-h-screen bg-gradient-to-br from-emerald-50 via-blue-50 to-indigo-100 flex items-center justify-center">
    <div class="relative w-full h-full flex items-center justify-center">
      <!-- Background decorative elements -->
      <div class="absolute top-8 left-8 w-32 h-32 bg-emerald-200 rounded-full opacity-20 animate-pulse"></div>
      <div class="absolute bottom-6 right-6 w-24 h-24 bg-blue-200 rounded-full opacity-30 animate-pulse delay-1000"></div>
      
      <!-- Main register card -->
      <div class="bg-white/80 backdrop-blur-sm p-8 rounded-2xl shadow-2xl w-full max-w-md border border-white/20 relative overflow-hidden">
        <!-- Header with icon -->
        <div class="text-center mb-8">
          <div class="inline-flex items-center justify-center w-16 h-16 bg-gradient-to-r from-emerald-500 to-blue-600 rounded-full mb-4 shadow-lg">
            <svg class="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 9v3m0 0v3m0-3h3m-3 0h-3m-2-5a4 4 0 11-8 0 4 4 0 018 0zM3 20a6 6 0 0112 0v1H3v-1z"></path>
            </svg>
          </div>
          <h2 class="text-3xl font-bold bg-gradient-to-r from-emerald-600 to-blue-600 bg-clip-text text-transparent mb-2">
            BudgetApp
          </h2>
          <p class="text-gray-600 text-sm">Bergabunglah dengan kami untuk mengelola keuangan</p>
        </div>

        <form @submit.prevent="handleRegister" class="space-y-6">
          <!-- Username field -->
          <div class="relative">
            <label for="username" class="text-gray-700 text-sm font-semibold mb-2 flex items-center">
              <svg class="w-4 h-4 mr-2 text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"></path>
              </svg>
              Nama Pengguna
            </label>
            <input
              type="text"
              id="username"
              v-model="username"
              class="w-full px-4 py-3 rounded-xl border border-gray-200 focus:outline-none focus:ring-2 focus:ring-emerald-500 focus:border-transparent transition-all duration-300 bg-gray-50 hover:bg-white"
              placeholder="Masukkan nama pengguna"
              required
            />
          </div>

          <!-- Password field -->
          <div class="relative">
            <label for="password" class="text-gray-700 text-sm font-semibold mb-2 flex items-center">
              <svg class="w-4 h-4 mr-2 text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z"></path>
              </svg>
              Kata Sandi
            </label>
            <input
              type="password"
              id="password"
              v-model="password"
              class="w-full px-4 py-3 rounded-xl border border-gray-200 focus:outline-none focus:ring-2 focus:ring-emerald-500 focus:border-transparent transition-all duration-300 bg-gray-50 hover:bg-white password-input"
              placeholder="Masukkan kata sandi"
              required
            />
          </div>

          <!-- Error message -->
          <div v-if="error" class="bg-red-50 border border-red-200 text-red-700 px-4 py-3 rounded-xl relative animate-shake" role="alert">
            <div class="flex items-center">
              <svg class="w-5 h-5 mr-2" fill="currentColor" viewBox="0 0 20 20">
                <path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7 4a1 1 0 11-2 0 1 1 0 012 0zm-1-9a1 1 0 00-1 1v4a1 1 0 102 0V6a1 1 0 00-1-1z" clip-rule="evenodd"></path>
              </svg>
              <span>{{ error }}</span>
            </div>
          </div>

          <!-- Register button -->
          <button
            type="submit"
            class="w-full bg-gradient-to-r from-emerald-500 to-blue-600 hover:from-emerald-600 hover:to-blue-700 text-white font-bold py-3 px-4 rounded-xl focus:outline-none focus:ring-2 focus:ring-emerald-500 focus:ring-offset-2 transition-all duration-300 transform hover:scale-[1.02] active:scale-[0.98] shadow-lg hover:shadow-xl"
          >
            <span class="flex items-center justify-center">
              <svg class="w-5 h-5 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 9v3m0 0v3m0-3h3m-3 0h-3m-2-5a4 4 0 11-8 0 4 4 0 018 0zM3 20a6 6 0 0112 0v1H3v-1z"></path>
              </svg>
              Daftar Sekarang
            </span>
          </button>

          <!-- Login link -->
          <div class="text-center pt-4 border-t border-gray-100">
            <p class="text-gray-600 text-sm">
              Sudah punya akun?
              <router-link 
                to="/login" 
                class="text-emerald-600 hover:text-emerald-700 font-semibold hover:underline transition-colors duration-200 ml-1"
              >
                Login di sini
              </router-link>
            </p>
          </div>
        </form>

        <!-- Feature highlights -->
        <div class="mt-8 pt-6 border-t border-gray-100">
          <div class="grid grid-cols-3 gap-4 text-center">
            <div class="flex flex-col items-center">
              <div class="w-8 h-8 bg-emerald-100 rounded-full flex items-center justify-center mb-2">
                <svg class="w-4 h-4 text-emerald-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
                </svg>
              </div>
              <span class="text-xs text-gray-600">Analisis</span>
            </div>
            <div class="flex flex-col items-center">
              <div class="w-8 h-8 bg-blue-100 rounded-full flex items-center justify-center mb-2">
                <svg class="w-4 h-4 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z"></path>
                </svg>
              </div>
              <span class="text-xs text-gray-600">Aman</span>
            </div>
            <div class="flex flex-col items-center">
              <div class="w-8 h-8 bg-indigo-100 rounded-full flex items-center justify-center mb-2">
                <svg class="w-4 h-4 text-indigo-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"></path>
                </svg>
              </div>
              <span class="text-xs text-gray-600">Cepat</span>
            </div>
          </div>
        </div>
      </div>
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

<style scoped>
@keyframes shake {
  0%, 100% { transform: translateX(0); }
  10%, 30%, 50%, 70%, 90% { transform: translateX(-2px); }
  20%, 40%, 60%, 80% { transform: translateX(2px); }
}

.animate-shake {
  animation: shake 0.5s ease-in-out;
}

/* Jarak antar karakter untuk placeholder kata sandi */
.password-input::placeholder {
  letter-spacing: 0.5px;
  font-family: 'Courier New', monospace;
}
</style>