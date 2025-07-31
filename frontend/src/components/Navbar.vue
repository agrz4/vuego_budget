<template>
  <nav class="bg-white/95 backdrop-blur-md border-b border-gray-200/50 shadow-lg sticky top-0 z-50">
    <div class="container mx-auto px-4 sm:px-6 lg:px-8">
      <div class="flex justify-between items-center h-16 lg:h-18">
        <!-- Logo and Brand -->
        <div class="flex items-center space-x-3">
          <router-link to="/" class="flex items-center space-x-3 group">
            <div class="w-10 h-10 lg:w-12 lg:h-12 bg-gradient-to-r from-emerald-500 to-blue-600 rounded-xl flex items-center justify-center shadow-lg group-hover:shadow-xl transition-all duration-300 group-hover:scale-105">
              <svg class="w-6 h-6 lg:w-7 lg:h-7 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1"></path>
              </svg>
            </div>
            <div class="hidden sm:block">
              <h1 class="text-xl lg:text-2xl font-bold bg-gradient-to-r from-emerald-600 to-blue-600 bg-clip-text text-transparent">
                BudgetApp
              </h1>
              <p class="text-xs lg:text-sm text-gray-500">Kelola keuangan dengan mudah</p>
            </div>
            <div class="sm:hidden">
              <h1 class="text-lg font-bold bg-gradient-to-r from-emerald-600 to-blue-600 bg-clip-text text-transparent">
                BudgetApp
              </h1>
            </div>
          </router-link>
        </div>

        <!-- Desktop Navigation -->
        <div class="hidden lg:flex items-center space-x-2">
          <template v-if="isAuthenticated">
            <router-link 
              to="/dashboard" 
              class="flex items-center px-4 py-2.5 text-gray-700 hover:text-emerald-600 hover:bg-emerald-50 rounded-xl transition-all duration-300 font-medium group"
              :class="{ 'text-emerald-600 bg-emerald-50 shadow-sm': $route.path === '/dashboard' }"
            >
              <svg class="w-5 h-5 mr-2 group-hover:scale-110 transition-transform duration-200" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 7v10a2 2 0 002 2h14a2 2 0 002-2V9a2 2 0 00-2-2H5a2 2 0 00-2-2z"></path>
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 5a2 2 0 012-2h4a2 2 0 012 2v6H8V5z"></path>
              </svg>
              Dashboard
            </router-link>
            
            <router-link 
              to="/budgets" 
              class="flex items-center px-4 py-2.5 text-gray-700 hover:text-emerald-600 hover:bg-emerald-50 rounded-xl transition-all duration-300 font-medium group"
              :class="{ 'text-emerald-600 bg-emerald-50 shadow-sm': $route.path === '/budgets' }"
            >
              <svg class="w-5 h-5 mr-2 group-hover:scale-110 transition-transform duration-200" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
              </svg>
              Anggaran
            </router-link>
            
            <router-link 
              to="/expenses" 
              class="flex items-center px-4 py-2.5 text-gray-700 hover:text-emerald-600 hover:bg-emerald-50 rounded-xl transition-all duration-300 font-medium group"
              :class="{ 'text-emerald-600 bg-emerald-50 shadow-sm': $route.path === '/expenses' }"
            >
              <svg class="w-5 h-5 mr-2 group-hover:scale-110 transition-transform duration-200" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1"></path>
              </svg>
              Pengeluaran
            </router-link>
          </template>
        </div>

        <!-- User Menu / Mobile Menu -->
        <div class="flex items-center space-x-2">
          <template v-if="isAuthenticated">
            <!-- User Profile Dropdown -->
            <div class="relative" ref="userDropdown">
              <button 
                @click="toggleUserMenu"
                class="flex items-center space-x-2 p-2 lg:px-3 lg:py-2 rounded-xl hover:bg-gray-100 transition-all duration-300 focus:outline-none focus:ring-2 focus:ring-emerald-500 group"
              >
                <div class="w-8 h-8 lg:w-9 lg:h-9 bg-gradient-to-r from-emerald-500 to-blue-600 rounded-full flex items-center justify-center shadow-md group-hover:shadow-lg transition-all duration-300">
                  <svg class="w-4 h-4 lg:w-5 lg:h-5 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"></path>
                  </svg>
                </div>
                <span class="text-sm lg:text-base font-medium text-gray-700 hidden md:block">{{ username || 'Pengguna' }}</span>
                <svg class="w-4 h-4 text-gray-500 transition-transform duration-200" :class="{ 'rotate-180': showUserMenu }" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"></path>
                </svg>
              </button>

              <!-- Dropdown Menu -->
              <transition
                enter-active-class="transition ease-out duration-200"
                enter-from-class="transform opacity-0 scale-95"
                enter-to-class="transform opacity-100 scale-100"
                leave-active-class="transition ease-in duration-150"
                leave-from-class="transform opacity-100 scale-100"
                leave-to-class="transform opacity-0 scale-95"
              >
                <div 
                  v-if="showUserMenu"
                  class="absolute right-0 mt-2 w-56 bg-white rounded-xl shadow-xl border border-gray-200 py-2 z-50 backdrop-blur-sm"
                >
                  <div class="px-4 py-3 border-b border-gray-100">
                    <p class="text-sm font-semibold text-gray-900">{{ username || 'Pengguna' }}</p>
                    <p class="text-xs text-emerald-600">Pengguna Aktif</p>
                  </div>
                  
                  <!-- Mobile Navigation Links (shown only on mobile) -->
                  <div class="lg:hidden border-b border-gray-100 py-2">
                    <router-link 
                      to="/dashboard" 
                      @click="closeUserMenu"
                      class="flex items-center px-4 py-2 text-sm text-gray-700 hover:bg-emerald-50 hover:text-emerald-600 transition-colors duration-200"
                      :class="{ 'text-emerald-600 bg-emerald-50': $route.path === '/dashboard' }"
                    >
                      <svg class="w-4 h-4 mr-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 7v10a2 2 0 002 2h14a2 2 0 002-2V9a2 2 0 00-2-2H5a2 2 0 00-2-2z"></path>
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 5a2 2 0 012-2h4a2 2 0 012 2v6H8V5z"></path>
                      </svg>
                      Dashboard
                    </router-link>
                    
                    <router-link 
                      to="/budgets" 
                      @click="closeUserMenu"
                      class="flex items-center px-4 py-2 text-sm text-gray-700 hover:bg-emerald-50 hover:text-emerald-600 transition-colors duration-200"
                      :class="{ 'text-emerald-600 bg-emerald-50': $route.path === '/budgets' }"
                    >
                      <svg class="w-4 h-4 mr-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
                      </svg>
                      Anggaran
                    </router-link>
                    
                    <router-link 
                      to="/expenses" 
                      @click="closeUserMenu"
                      class="flex items-center px-4 py-2 text-sm text-gray-700 hover:bg-emerald-50 hover:text-emerald-600 transition-colors duration-200"
                      :class="{ 'text-emerald-600 bg-emerald-50': $route.path === '/expenses' }"
                    >
                      <svg class="w-4 h-4 mr-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8c-1.657 0-3 .895-3 2s1.343 2 3 2 3 .895 3 2-1.343 2-3 2m0-8c1.11 0 2.08.402 2.599 1M12 8V7m0 1v8m0 0v1m0-1c-1.11 0-2.08-.402-2.599-1"></path>
                      </svg>
                      Pengeluaran
                    </router-link>
                  </div>
                  
                  <button 
                    @click="logout"
                    class="w-full text-left px-4 py-2 text-sm text-red-600 hover:bg-red-50 transition-colors duration-200 flex items-center"
                  >
                    <svg class="w-4 h-4 mr-3" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1"></path>
                    </svg>
                    Keluar
                  </button>
                </div>
              </transition>
            </div>
          </template>
          
          <!-- Desktop Login/Register Buttons (Right Side) -->
          <template v-else>
            <div class="flex items-center space-x-2">
              <!-- Desktop Login/Register Buttons -->
              <router-link 
                to="/login" 
                class="hidden lg:flex items-center px-4 py-2.5 text-gray-700 hover:text-emerald-600 hover:bg-emerald-50 rounded-xl transition-all duration-300 font-medium group"
                :class="{ 'text-emerald-600 bg-emerald-50 shadow-sm': $route.path === '/login' }"
              >
                <svg class="w-5 h-5 mr-2 group-hover:scale-110 transition-transform duration-200" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 16l-4-4m0 0l4-4m-4 4h14m-5 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h7a3 3 0 013 3v1"></path>
                </svg>
                Masuk
              </router-link>
              
              <router-link 
                to="/register" 
                class="hidden lg:flex items-center px-4 py-2.5 bg-gradient-to-r from-emerald-500 to-blue-600 hover:from-emerald-600 hover:to-blue-700 text-white rounded-xl transition-all duration-300 font-medium shadow-lg hover:shadow-xl transform hover:scale-105"
                :class="{ 'from-emerald-600 to-blue-700 shadow-xl': $route.path === '/register' }"
              >
                <svg class="w-5 h-5 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 9v3m0 0v3m0-3h3m-3 0h-3m-2-5a4 4 0 11-8 0 4 4 0 018 0zM3 20a6 6 0 0112 0v1H3v-1z"></path>
                </svg>
                Daftar
              </router-link>
              
              <!-- Mobile Login/Register Buttons -->
              <div class="flex items-center space-x-2 lg:hidden">
                <router-link 
                  to="/login" 
                  class="flex items-center px-3 py-2 text-sm text-gray-700 hover:text-emerald-600 hover:bg-emerald-50 rounded-lg transition-all duration-300 font-medium"
                  :class="{ 'text-emerald-600 bg-emerald-50': $route.path === '/login' }"
                >
                  <svg class="w-4 h-4 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 16l-4-4m0 0l4-4m-4 4h14m-5 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h7a3 3 0 013 3v1"></path>
                  </svg>
                  Masuk
                </router-link>
                
                <router-link 
                  to="/register" 
                  class="flex items-center px-3 py-2 bg-gradient-to-r from-emerald-500 to-blue-600 hover:from-emerald-600 hover:to-blue-700 text-white rounded-lg transition-all duration-300 font-medium text-sm shadow-md hover:shadow-lg transform hover:scale-105"
                >
                  <svg class="w-4 h-4 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 9v3m0 0v3m0-3h3m-3 0h-3m-2-5a4 4 0 11-8 0 4 4 0 018 0zM3 20a6 6 0 0112 0v1H3v-1z"></path>
                  </svg>
                  Daftar
                </router-link>
              </div>
            </div>
          </template>
        </div>
      </div>
    </div>
  </nav>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue';
import { useRouter } from 'vue-router';
import { checkAuth, logoutUser } from '../services/auth';

const router = useRouter();
const isAuthenticated = ref(false);
const username = ref('');
const showUserMenu = ref(false);
const userDropdown = ref(null);

const checkAuthenticationStatus = async () => {
  isAuthenticated.value = await checkAuth();
  // Get username from localStorage or API
  username.value = localStorage.getItem('username') || '';
};

const logout = async () => {
  await logoutUser();
  isAuthenticated.value = false;
  username.value = '';
  showUserMenu.value = false;
  router.push('/login');
};

const toggleUserMenu = () => {
  showUserMenu.value = !showUserMenu.value;
};

const closeUserMenu = () => {
  showUserMenu.value = false;
};

// Close dropdown when clicking outside
const handleClickOutside = (event) => {
  if (userDropdown.value && !userDropdown.value.contains(event.target)) {
    showUserMenu.value = false;
  }
};

onMounted(() => {
  checkAuthenticationStatus();
  window.addEventListener('auth-status-changed', checkAuthenticationStatus);
  document.addEventListener('click', handleClickOutside);
});

onUnmounted(() => {
  window.removeEventListener('auth-status-changed', checkAuthenticationStatus);
  document.removeEventListener('click', handleClickOutside);
});
</script>

<style scoped>
/* Custom animations */
@keyframes slideDown {
  from {
    opacity: 0;
    transform: translateY(-10px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.dropdown-enter-active {
  animation: slideDown 0.2s ease-out;
}

/* Smooth hover effects */
.nav-link {
  position: relative;
  overflow: hidden;
}

.nav-link::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(255,255,255,0.2), transparent);
  transition: left 0.5s;
}

.nav-link:hover::before {
  left: 100%;
}

/* Mobile optimization */
@media (max-width: 640px) {
  .container {
    padding-left: 1rem;
    padding-right: 1rem;
  }
}

/* Tablet optimization */
@media (min-width: 768px) and (max-width: 1023px) {
  .nav-item {
    padding: 0.5rem 0.75rem;
  }
}
</style>