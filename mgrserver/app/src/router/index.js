
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
        }, {
          path: 'score',
          name: 'score',
          component: () => import('../pages/home/score.vue'),
          meta: { title: "首页" }
          }, {
            path: 'exchange',
            name: 'exchange',
            component: () => import('../pages/home/exchange.vue'),
            meta: { title: "首页" }
            }
      ]
    },
    {
      path: '/score/add/:id',
      name: '/score/add',
      component: () => import('../pages/score/add.vue'),
      meta: { title: "积分" }
      },
      {
        path: '/score/index/:id',
        name: '/score/index',
        component: () => import('../pages/score/index.vue'),
        meta: { title: "积分" }
        },
  ]
})

