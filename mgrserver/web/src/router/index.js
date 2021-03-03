
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
      path: '/',
      name: 'menus',
      component: () => import('../pages/system/menus.vue'),
      children:[
				{
					path: 'punishment/info',
					name: 'PunishmentInfo',
					component: () => import('../pages/punishment/punishment.info.list.vue')
				},
				{
					path: 'punishment/info/detail',
					name: 'PunishmentInfoDetail',
					component: () => import('../pages/punishment/punishment.info.detail.vue')
				},
				{
					path: 'reward/info',
					name: 'RewardInfo',
					component: () => import('../pages/reward/reward.info.list.vue')
				},
				{
					path: 'reward/info/detail',
					name: 'RewardInfoDetail',
					component: () => import('../pages/reward/reward.info.detail.vue')
				},
				{
					path: 'score/record',
					name: 'ScoreRecord',
					component: () => import('../pages/score/score.record.list.vue')
				},
				{
					path: 'score/record/detail',
					name: 'ScoreRecordDetail',
					component: () => import('../pages/score/score.record.detail.vue')
				},
				{
					path: 'user/info',
					name: 'UserInfo',
					component: () => import('../pages/user/user.info.list.vue')
				},
				{
					path: 'user/info/detail',
					name: 'UserInfoDetail',
					component: () => import('../pages/user/user.info.detail.vue')
				},
				{
					path: 'dict/info',
					name: 'DictInfo',
					component: () => import('../pages/dict/dict.info.list.vue')
				},
				{
					path: 'dict/info/detail',
					name: 'DictInfoDetail',
					component: () => import('../pages/dict/dict.info.detail.vue')
				},
				{
					path: 'gif/info',
					name: 'GifInfo',
					component: () => import('../pages/gif/gif.info.list.vue')
				},
				{
					path: 'gif/info/detail',
					name: 'GifInfoDetail',
					component: () => import('../pages/gif/gif.info.detail.vue')
				},
      ]
    }
  ]
})

