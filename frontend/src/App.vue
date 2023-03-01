<script setup>
import { Collapse } from "bootstrap";
import {
  faBars,
  faCode,
  faEnvelope,
  faInfoCircle,
  faListUl,
} from "@fortawesome/pro-solid-svg-icons";
import { faGithub } from "@fortawesome/free-brands-svg-icons";

const transitionName = ref("fade");
const navbarShrink = ref(false);
const minimal = ref(false);
const showNav = ref(false);

const router = useRouter();
router.afterEach((to, from) => {
  if (from.href) {
    const routes = router.getRoutes();
    const fromIndex = routes.findIndex((o) => o.path === from.path);
    const toIndex = routes.findIndex((o) => o.path === to.path);
    transitionName.value = fromIndex < toIndex ? "slide-left" : "slide-right";
  }

  minimal.value = to.path === "/";
});

let collapse = null;
const navbarResponsive = ref(null);
onMounted(() => {
  window.addEventListener("scroll", onScroll);
  collapse = new Collapse(navbarResponsive.value, {
    toggle: showNav.value,
  });
});

onUnmounted(() => {
  window.removeEventListener("scroll", onScroll);
});

const onScroll = () => {
  navbarShrink.value = window.scrollY > 0;
};

const toggleNav = (show) => {
  if (show === undefined) {
    show = !showNav.value;
  }
  if (show) {
    collapse.show();
    showNav.value = true;
  } else {
    collapse.hide();
    showNav.value = false;
  }
};
</script>

<template>
  <a class="sr-only-focusable btn btn-primary position-absolute" href="#content"
    >Skip to main content</a
  >

  <nav
    class="navbar navbar-expand-md fixed-top navbar-dark"
    :class="{ 'navbar-shrink': navbarShrink }"
  >
    <div class="container">
      <router-link to="/" class="navbar-brand" @click="toggleNav(false)">
        <span aria-hidden="true">&lt; gabe.cook &gt;</span>
        <span class="sr-only">Gabe Cook</span>
      </router-link>
      <button
        class="navbar-toggler navbar-toggler-right"
        type="button"
        aria-controls="navbarResponsive"
        :aria-expanded="showNav"
        aria-label="Toggle navigation"
        @click="toggleNav()"
      >
        <font-awesome-icon :icon="faBars" fixed-width />
        Menu
      </button>
      <div
        id="navbarResponsive"
        ref="navbarResponsive"
        class="collapse navbar-collapse"
      >
        <ul class="navbar-nav ms-auto">
          <li class="nav-item">
            <router-link
              to="/about"
              class="nav-link rounded"
              @click="toggleNav(false)"
            >
              <font-awesome-icon
                :icon="faInfoCircle"
                fixed-width
              ></font-awesome-icon>
              About
            </router-link>
          </li>
          <li class="nav-item">
            <router-link
              to="/skills"
              class="nav-link rounded"
              @click="toggleNav(false)"
            >
              <font-awesome-icon
                :icon="faListUl"
                fixed-width
              ></font-awesome-icon>
              Skills
            </router-link>
          </li>
          <li class="nav-item">
            <router-link
              to="/projects"
              class="nav-link rounded"
              @click="toggleNav(false)"
            >
              <font-awesome-icon :icon="faCode" fixed-width></font-awesome-icon>
              Projects
            </router-link>
          </li>
          <li class="nav-item">
            <router-link
              to="/connect"
              class="nav-link rounded"
              @click="toggleNav(false)"
            >
              <font-awesome-icon
                :icon="faEnvelope"
                fixed-width
              ></font-awesome-icon>
              Connect
            </router-link>
          </li>
        </ul>
      </div>
    </div>
  </nav>

  <div
    v-if="showNav"
    class="position-absolute w-100 h-100 opacity-0"
    aria-hidden="true"
    style="z-index: 1000"
    @click="toggleNav(false)"
  />

  <router-view
    id="content"
    v-slot="{ Component }"
    class="text-center"
    :class="{ 'content-section': !minimal }"
  >
    <transition :name="transitionName" mode="out-in" appear>
      <keep-alive>
        <component :is="Component" />
      </keep-alive>
    </transition>
  </router-view>

  <footer :class="{ minimal }">
    <div class="container">
      <div class="row text-center">
        <div class="col-sm text-sm-start" aria-hidden="true">
          &lt;/ gabe.cook &gt;
        </div>
        <div class="col-sm">
          <a href="//github.com/gabe565/portfolio" target="_blank"
            ><font-awesome-icon :icon="faGithub" fixed-width></font-awesome-icon
            >View on GitHub</a
          >
        </div>
        <div class="col-sm text-sm-end">
          <a
            href="//github.com/gabe565/portfolio/blob/master/LICENSE"
            target="_blank"
          >
            &copy; {{ new Date().getFullYear() }} Gabe Cook
          </a>
        </div>
      </div>
    </div>
  </footer>
</template>
