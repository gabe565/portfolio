import { library } from "@fortawesome/fontawesome-svg-core";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";

import {
  faBars,
  faCode,
  faEnvelope,
  faInfoCircle,
  faListUl,
  faStar,
} from "@fortawesome/pro-solid-svg-icons";
library.add(faBars, faCode, faEnvelope, faInfoCircle, faListUl, faStar);

import {
  faAt,
  faComment,
  faPaperPlane,
  faStar as faStarEmpty,
  faSync,
  faUserAlt,
} from "@fortawesome/pro-regular-svg-icons";
library.add(faAt, faComment, faPaperPlane, faStarEmpty, faSync, faUserAlt);

import {
  faGithub,
  faLinkedin,
  faTwitter,
} from "@fortawesome/free-brands-svg-icons";
library.add(faAt, faGithub, faLinkedin, faTwitter);

export default FontAwesomeIcon;
