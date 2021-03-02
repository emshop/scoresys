
import Vue from 'vue';
import Router from 'vue-router';

Vue.use(Router);

const VueRouterPush = Router.prototype.push
Router.prototype.push = function push (to) {
  return VueRouterPush.call(this, to).catch(err => err)
}

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/home',
      name: 'home',
      component: () => import('../pages/system/menus.vue'),
      children:[
        {
        path: 'index',
        name: 'index',
        component: () => import('../pages/home/index.vue'),
        meta: { title: "首页" }
        }
      ]
    },
    {
      path: '/score',
      name: '/score',
      component: () => import('../pages/score/index.vue'),
      meta: { title: "积分" }
      },
  ]
})

