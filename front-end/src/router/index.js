import Vue from 'vue'
import Router from 'vue-router'
import index from '../components/index'
import temp from '../components/temp'
import order from '../components/order'
import other from '../components/other'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'index',
      component: index
    },
    {
      path: '/马原',
      name: 'mayuan',
      component: temp
    },
    {
      path: '/毛概',
      name: 'maogai',
      component: temp
    },
    {
      path: '/马原/顺序刷题',
      name: 'mayuanorder',
      component: order
    },
    {
      path: '/毛概/顺序刷题',
      name: 'maogaiorder',
      component: order
    },
    {
      path: '/马原/随机刷题',
      name: 'mayuanrandom',
      component: other
    },
    {
      path: '/毛概/随机刷题',
      name: 'maogairandom',
      component: other
    },
    {
      path: '/马原/考试模拟',
      name: 'mayuanexam',
      component: other
    },
    {
      path: '/毛概/考试模拟',
      name: 'maogaiexam',
      component: other
    }
  ]
})
