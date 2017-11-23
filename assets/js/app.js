import { has } from 'lodash'
window._ = _

import jQuery from "jquery"
window.$ = window.jQuery = jQuery

import 'popper.js/dist/umd/popper'
import 'bootstrap'

import axios from 'axios'
window.axios = axios

import {TweenMax, Power2, TimelineLite} from "gsap";
import ScrollMagic from 'ScrollMagic';

import Vue from 'vue'
import VueProgressBar from 'vue-progressbar'
import 'vue-svgicon/dist/polyfill'
import VueSVGIcon from 'vue-svgicon'
import * as VueGoogleMaps from 'vue2-google-maps'

//Vue.use(VueRouter)
Vue.use(VueProgressBar, {
    color: '#8cb9df',
    failedColor: 'red',
    height: '2px'
})
Vue.use(VueSVGIcon)
Vue.use(VueGoogleMaps, {
    load: {
        key: 'AIzaSyD_X6EOD9nVWa8YDdb4_-8qIeyZUUjcGho'
    }
})

import Home from './components/Home.vue'
import About from './components/About.vue'
import Skills from './components/Skills.vue'
import Projects from './components/Projects.vue'
import Connect from './components/Connect.vue'
import Gmap from './components/Gmap.vue'
import NotFound from './components/NotFound.vue'

import './svg/bars'
import './svg/info-circle'
import './svg/list-ul'
import './svg/terminal'
import './svg/envelope'

const app = new Vue({
    el: '#app',
    components: {
        Home,
        About,
        Skills,
        Projects,
        Connect,
        Gmap,
        NotFound
    }
})

require('./main')
