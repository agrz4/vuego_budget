import { createApp } from 'vue';
import App from './App.vue';
import router from './router';
import './assets/main.css'; // Import Tailwind CSS

const app = createApp(App);
app.use(router);
app.mount('#app');