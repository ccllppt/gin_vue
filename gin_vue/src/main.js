import Vue from 'vue';
import { BootstrapVue, IconsPlugin } from 'bootstrap-vue';
// Import Bootstrap and BootstrapVue CSS files (order is important)
import 'bootstrap/dist/css/bootstrap.css';
import 'bootstrap-vue/dist/bootstrap-vue.css';
import './assets/scss/index.scss';
import axios from 'axios';
import VueAxios from 'vue-axios';
import Vuelidate from 'vuelidate';

import App from './App.vue';
import router from './router';
import store from './store';

// Make BootstrapVue available throughout your project
Vue.use(BootstrapVue);
// Optionally install the BootstrapVue icon components plugin
Vue.use(IconsPlugin);

Vue.config.productionTip = false;
// Make BootstrapVue available throughout your project
Vue.use(BootstrapVue);
// Optionally install the BootstrapVue icon components plugin
Vue.use(IconsPlugin);
// vuelidate
Vue.use(Vuelidate);
// axios
Vue.use(VueAxios, axios);

new Vue({
  router,
  store,
  render: (h) => h(App),
}).$mount('#app');
