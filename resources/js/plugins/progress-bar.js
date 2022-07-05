import Vue from 'vue';
import VueProgressBar from 'vue-progressbar';

Vue.use(VueProgressBar, {
    autoFinish: false,
    color: '#8cb9df',
    failedColor: 'red',
    thickness: '3px'
});
