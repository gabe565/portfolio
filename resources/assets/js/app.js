
/**
 * First we will load all of this project's JavaScript dependencies which
 * includes Vue and other libraries. It is a great starting point when
 * building robust, powerful web applications using Vue and Laravel.
 */

require('./bootstrap')
require('jquery.easing')

window.Vue = require('vue')

require('./main')
require('./map')

import VueRouter from 'vue-router';
Vue.use(VueRouter)

const Home = require('./components/Home.vue')
const About = require('./components/About.vue')
const Skills = require('./components/Skills.vue')
const Contact = require('./components/Contact.vue')
const NotFound = require('./components/NotFound.vue')

const routes = [
    { path: '/', component: Home },
    { path: '/about', component: About },
    { path: '/skills', component: Skills },
    { path: '/contact', component: Contact },
    { path: '*', component: NotFound },
]

const router = new VueRouter({
    routes,
    mode: 'history',
    linkActiveClass: 'active'
})

const app = new Vue({
    el: '#app',
    router
})

$(function(){ 
    var navMain = $(".navbar-collapse"); // avoid dependency on #id
    // "a:not([data-toggle])" - to avoid issues caused
    // when you have dropdown inside navbar
    navMain.on("click", "a:not([data-toggle])", null, function () {
        navMain.collapse('hide');
    });
});
