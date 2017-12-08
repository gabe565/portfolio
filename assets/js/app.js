import { has } from 'lodash'
window._ = _

import jQuery from "jquery"
window.$ = window.jQuery = jQuery

import 'popper.js/dist/umd/popper'
import 'bootstrap'

import axios from 'axios'
window.axios = axios

import { TweenMax, Power2, TimelineMax } from "gsap"
window.TweenMax = TweenMax
window.TimelineMax = TimelineMax
import * as ScrollMagic from 'scrollmagic'
import 'imports-loader?define=>false!scrollmagic/scrollmagic/uncompressed/plugins/animation.gsap.js'
import 'imports-loader?define=>false!scrollmagic/scrollmagic/uncompressed/plugins/debug.addIndicators.js'

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

var selectors = [
    '#about',
    '#skills',
    '#projects',
    '#connect',
]
var locations = []
$('.navbar-nav .nav-item').each(function(i) {
    locations.push({
        "left": $(this).position().left,
        "width": $(this).outerWidth()
    })
})

var controller = new ScrollMagic.Controller()
_.forEach(locations, function(value, key) {
    if (key === 0) return;

    var tween = TweenMax.to('.navbar-nav .current', 1, { left: value.left, width: value.width, ease: Ease.easeInOut }, 0.15)
    
    new ScrollMagic.Scene({triggerElement: selectors[key - 1], duration: 300})
        .setTween(tween)
        .addIndicators({name: 'staggering'}) // add indicators (requires plugin)
        .addTo(controller)
})
