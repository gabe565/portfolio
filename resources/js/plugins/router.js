import Vue from 'vue';
import VueRouter from 'vue-router'

Vue.use(VueRouter);

import Home from '../components/Home.vue';
import About from '../components/About.vue';
import Skills from '../components/Skills.vue';
import Projects from '../components/Projects.vue';
import Connect from '../components/Connect.vue';
import NotFound from '../components/NotFound.vue';

export const routes = [
    {
        path: '/',
        component: Home,
        meta: { title: 'Home' }
    },
    {
        path: '/about',
        component: About,
        meta: { title: 'About' }
    },
    {
        path: '/skills',
        component: Skills,
        meta: { title: 'Skills' }
    },
    {
        path: '/projects',
        component: Projects,
        meta: { title: 'Projects' }
    },
    {
        path: '/connect',
        component: Connect,
        meta: { title: 'Connect' }
    },
    {
        path: '*',
        component: NotFound,
        meta: { title: 'Not Found' }
    },
];

export const router = new VueRouter({
    routes,
    mode: 'history',
    linkActiveClass: 'active',
    scrollBehavior(to, from, savedPosition) {
        return new Promise((resolve, reject) => {
            setTimeout(() => {
                if (savedPosition)
                    return savedPosition
                else
                    return { x: 0, y: 0 }
            }, 200)
        })
    }
});

export const initialTitle = document.title;

router.beforeEach((to, from, next) => {
    document.title = (to.meta.title ? to.meta.title + ' | ' : '') + (router.app.appName || initialTitle)
    next()
});
