<script setup>
import { onMounted, onUnmounted, ref } from "vue";
import { useRouter } from "vue-router";
import { Collapse } from "bootstrap";
import MenuIcon from "~icons/material-symbols/menu-rounded";
import InfoIcon from "~icons/material-symbols/info";
import ListIcon from "~icons/material-symbols/list-rounded";
import CodeIcon from "~icons/material-symbols/code-rounded";
import MailIcon from "~icons/material-symbols/mail-rounded";
import GithubIcon from "~icons/simple-icons/github";

const transitionName = ref("fade");
const navbarShrink = ref(false);
const minimal = ref(true);
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
  <a class="visually-hidden-focusable btn btn-primary position-absolute" href="#content"
    >Skip to main content</a
  >

  <nav
    class="navbar navbar-expand-md fixed-top navbar-dark"
    :class="{ 'navbar-shrink': navbarShrink }"
  >
    <div class="container">
      <router-link to="/" class="navbar-brand" @click="toggleNav(false)">
        <span aria-hidden="true">&lt; gabe.cook &gt;</span>
        <span class="visually-hidden">Gabe Cook</span>
      </router-link>
      <button
        class="navbar-toggler navbar-toggler-right"
        type="button"
        aria-controls="navbarResponsive"
        :aria-expanded="showNav"
        aria-label="Toggle navigation"
        @click="toggleNav()"
      >
        <menu-icon />
        Menu
      </button>
      <div id="navbarResponsive" ref="navbarResponsive" class="collapse navbar-collapse">
        <ul class="navbar-nav ms-auto">
          <li class="nav-item">
            <router-link to="/about" class="nav-link rounded" @click="toggleNav(false)">
              <info-icon />
              About
            </router-link>
          </li>
          <li class="nav-item">
            <router-link to="/skills" class="nav-link rounded" @click="toggleNav(false)">
              <list-icon />
              Skills
            </router-link>
          </li>
          <li class="nav-item">
            <router-link to="/projects" class="nav-link rounded" @click="toggleNav(false)">
              <code-icon />
              Projects
            </router-link>
          </li>
          <li class="nav-item">
            <router-link to="/connect" class="nav-link rounded" @click="toggleNav(false)">
              <mail-icon />
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

  <div class="d-flex flex-column min-vh-100">
    <router-view id="content" v-slot="{ Component }" class="text-center">
      <transition :name="transitionName" mode="out-in" appear>
        <keep-alive>
          <component :is="Component" />
        </keep-alive>
      </transition>
    </router-view>

    <transition name="fade-in">
      <footer v-show="!minimal" class="mt-auto">
        <div class="container">
          <div class="row text-center">
            <div class="col-sm text-sm-start" aria-hidden="true">&lt;/ gabe.cook &gt;</div>
            <div class="col-sm">
              <a href="//github.com/gabe565/portfolio" target="_blank"
                ><github-icon class="me-1" fixed-width fill />View on GitHub</a
              >
            </div>
            <div class="col-sm text-sm-end">
              <a href="//github.com/gabe565/portfolio/blob/master/LICENSE" target="_blank">
                &copy; {{ new Date().getFullYear() }} Gabe Cook
              </a>
            </div>
          </div>
        </div>
      </footer>
    </transition>
  </div>
</template>
