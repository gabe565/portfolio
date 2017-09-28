
/**
 * First we will load all of this project's JavaScript dependencies which
 * includes Vue and other libraries. It is a great starting point when
 * building robust, powerful web applications using Vue and Laravel.
 */


require('./bootstrap')

require('./main')
require('./map')

import Vue from 'vue'
import VueRouter from 'vue-router'
Vue.use(VueRouter)

import Home from './components/Home.vue'
import About from './components/About.vue'
import Skills from './components/Skills.vue'
import Projects from './components/Projects.vue'
import Connect from './components/Connect.vue'
import NotFound from './components/NotFound.vue'

const routes = [
    { path: '/', component: Home, meta: { title: '' } },
    { path: '/about', component: About, meta: { title: 'About' } },
    { path: '/skills', component: Skills, meta: { title: 'Skills' } },
    { path: '/projects', component: Projects, meta: { title: 'Projects' } },
    { path: '/connect', component: Connect, meta: { title: 'Connect' } },
    { path: '*', component: NotFound, meta: { title: 'Not Found' } },
]

const router = new VueRouter({
    routes,
    mode: 'history',
    linkActiveClass: 'active'
})

router.beforeEach((to, from, next) => {
    document.title = to.meta.title ? to.meta.title + ' · Gabe Cook' : 'Gabe Cook'
    next()
})

const app = new Vue({
    el: '#app',
    router
})
