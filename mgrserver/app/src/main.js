

import "jquery"
import "bootstrap"
 
import Vue from 'vue'
import App from './App'
import router from './router'

import VueCookies from 'vue-cookies'
Vue.use(VueCookies);

import Vant from 'vant'
import 'vant/lib/index.css'

Vue.use(Vant)

import utility from './utility'
Vue.use(utility,false);

Vue.config.productionTip = false;

router.beforeEach((to, from, next) => {
  /* 路由发生变化修改页面title */
  Vue.prototype.$sys.checkAuthCode(to)
  if (to.path != "/") {
      document.title = Vue.prototype.$sys.getTitle(to.path)
  }
  next()
})

  /* eslint-disable no-new */
new Vue({
    el: '#app',
    router,
    components: {
        App
    },
    template: '<App/>'
});

