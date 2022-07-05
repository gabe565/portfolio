
/**
 * First we will load all of this project's JavaScript dependencies which
 * includes Vue and other libraries. It is a great starting point when
 * building robust, powerful web applications using Vue and Laravel.
 */

require('./bootstrap');

import Vue from 'vue';

import './plugins/google-maps';
import './plugins/progress-bar';
import './plugins/svgicon';
import { initialTitle, router, routes } from './plugins/router';

const app = new Vue({
    el: '#app',
    router,
    data() {
        return {
            transitionName: 'fade',
            appName: initialTitle
        }
    },
    watch: {
        '$route'(to, from) {
            let fromIndex = routes.findIndex(o => o.path == from.path)
            let toIndex = routes.findIndex(o => o.path == to.path)
            this.transitionName = (fromIndex < toIndex) ? 'slide-left' : 'slide-right'
        }
    }
});

require('./main');
