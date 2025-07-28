import { createRouter, createWebHistory } from 'vue-router';
import Login from '../views/Auth/Login.vue';
import Register from '../views/Auth/Register.vue';
import Dashboard from '../views/Dashboard.vue';
import BudgetList from '../views/Budget/BudgetList.vue';
import BudgetForm from '../views/Budget/BudgetForm.vue';
import ExpenseList from '../views/Expense/ExpenseList.vue';
import ExpenseForm from '../views/Expense/ExpenseForm.vue';
import { checkAuth } from '../services/auth';

const routes = [
  { path: '/login', name: 'Login', component: Login },
  { path: '/register', name: 'Register', component: Register },
  {
    path: '/dashboard',
    name: 'Dashboard',
    component: Dashboard,
    meta: { requiresAuth: true }
  },
  {
    path: '/budgets',
    name: 'BudgetList',
    component: BudgetList,
    meta: { requiresAuth: true }
  },
  {
    path: '/budgets/add',
    name: 'AddBudget',
    component: BudgetForm,
    meta: { requiresAuth: true }
  },
  {
    path: '/budgets/edit/:id',
    name: 'EditBudget',
    component: BudgetForm,
    props: true,
    meta: { requiresAuth: true }
  },
  {
    path: '/expenses',
    name: 'ExpenseList',
    component: ExpenseList,
    meta: { requiresAuth: true }
  },
  {
    path: '/expenses/add',
    name: 'AddExpense',
    component: ExpenseForm,
    meta: { requiresAuth: true }
  },
  {
    path: '/expenses/edit/:id',
    name: 'EditExpense',
    component: ExpenseForm,
    props: true,
    meta: { requiresAuth: true }
  },
  { path: '/', redirect: '/dashboard' } // Redirect root to dashboard
];

const router = createRouter({
  history: createWebHistory(),
  routes
});

router.beforeEach(async (to, from, next) => {
  const isAuthenticated = await checkAuth();
  if (to.meta.requiresAuth && !isAuthenticated) {
    next('/login');
  } else if ((to.name === 'Login' || to.name === 'Register') && isAuthenticated) {
    next('/dashboard');
  } else {
    next();
  }
});

export default router;