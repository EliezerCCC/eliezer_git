import Vue from 'vue'
import App from './App.vue'
import router from './router'
import axios from 'axios'
import VueAxios from 'vue-axios'
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
import global from './global/global.vue'
import HeaderTop from './components/HeaderTop.vue'
import EditArticle from './components/EditArticle.vue'

Vue.use(ElementUI)
Vue.use(VueAxios, axios)
Vue.config.productionTip = false
Vue.prototype.global =global
Vue.component('HeaderTop',HeaderTop);
Vue.component('EditArticle',EditArticle);

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
