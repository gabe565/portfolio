/* Set up using Vue 2 */
import Vue from 'vue';

/* import the fontawesome core */
import { library } from '@fortawesome/fontawesome-svg-core';

/* import font awesome icon component */
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';

/* import specific icons */
import {
    faBars,
    faCode,
    faEnvelope,
    faInfoCircle,
    faListUl,
    faStar,
} from '@fortawesome/pro-solid-svg-icons';
library.add(
    faBars,
    faCode,
    faEnvelope,
    faInfoCircle,
    faListUl,
    faStar,
);

import {
    faAt,
    faComment,
    faPaperPlane,
    faStar as faStarEmpty,
    faSync,
    faUserAlt
} from '@fortawesome/pro-regular-svg-icons';
library.add(
    faAt,
    faComment,
    faPaperPlane,
    faStarEmpty,
    faSync,
    faUserAlt,
);

import {
    faGithub,
    faLinkedin,
    faTwitter
} from '@fortawesome/free-brands-svg-icons';
library.add(faAt,
    faGithub,
    faLinkedin,
    faTwitter,
);

/* add font awesome icon component */
Vue.component('font-awesome-icon', FontAwesomeIcon);
