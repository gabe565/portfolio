<script>
import { Collapse } from "bootstrap";

let collapse = null;

export default {
  data: () => ({
    transitionName: "fade",
    navbarShrink: false,
    minimal: false,
    showNav: false,
  }),
  watch: {
    $route(to, from) {
      if (from.href) {
        const routes = this.$router.getRoutes();
        const fromIndex = routes.findIndex((o) => o.path === from.path);
        const toIndex = routes.findIndex((o) => o.path === to.path);
        this.transitionName =
          fromIndex < toIndex ? "slide-left" : "slide-right";
      }

      this.minimal = to.path === "/";
    },
  },
  mounted() {
    window.addEventListener("scroll", this.scroll);
    collapse = new Collapse(this.$refs.navbarResponsive, {
      toggle: this.showNav,
    });
  },
  unmounted() {
    window.removeEventListener("scroll", this.scroll);
  },
  methods: {
    scroll() {
      this.navbarShrink = window.scrollY > 0;
    },
    toggleNav(show) {
      if (show === undefined) {
        show = !this.showNav;
      }
      if (show) {
        collapse.show();
        this.showNav = true;
      } else {
        collapse.hide();
        this.showNav = false;
      }
    },
  },
};
</script>

<template>
  <a class="sr-only-focusable btn btn-primary position-absolute" href="#content"
    >Skip to main content</a
  >

  <nav
    class="navbar navbar-expand-md fixed-top navbar-dark"
    :class="{ 'navbar-shrink': navbarShrink, minimal }"
  >
    <div class="container">
      <router-link to="/" class="navbar-brand" @click="toggleNav(false)"
        >&lt; gabe.cook &gt;</router-link
      >
      <button
        class="navbar-toggler navbar-toggler-right"
        type="button"
        aria-controls="navbarResponsive"
        :aria-expanded="showNav"
        aria-label="Toggle navigation"
        @click="toggleNav()"
      >
        <font-awesome-icon icon="fas fa-bars fa-fw" fixed-width />
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
                icon="fas fa-info-circle"
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
                icon="fas fa-list-ul"
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
              <font-awesome-icon
                icon="fas fa-code"
                fixed-width
              ></font-awesome-icon>
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
                icon="fas fa-envelope"
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
    class="child-view text-center"
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
        <div class="col-sm text-sm-start">
          <span>&lt;/ gabe.cook &gt;</span>
        </div>
        <div class="col-sm">
          <a href="//github.com/gabe565/portfolio" target="_blank">
            <span>
              <font-awesome-icon icon="fab fa-github"></font-awesome-icon>
              &nbsp;View on GitHub
            </span>
          </a>
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
