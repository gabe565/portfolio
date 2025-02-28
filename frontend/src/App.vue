<script setup>
import { ref } from "vue";
import { useRouter } from "vue-router";
import MenuIcon from "~icons/material-symbols/menu-rounded";
import GithubIcon from "~icons/simple-icons/github";
import FooterMap from "@/components/FooterMap.vue";
import MainMenu from "@/components/MainMenu.vue";

const transitionName = ref("fade");
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

  showNav.value = false;
  minimal.value = to.path === "/";
});
</script>

<template>
  <div class="flex flex-col place-items-center min-h-lvh overflow-x-hidden">
    <a class="sr-only focus:not-sr-only btn btn-primary z-50" href="#content"
      >Skip to main content</a
    >

    <nav
      class="navbar fixed z-30 border-white border-b border-opacity-5 bg-base-200 bg-opacity-20 shadow-xl transition duration-700 before:backdrop-blur before:backdrop-hack"
      :class="{
        'bg-opacity-80': !minimal,
        'shadow-transparent': minimal,
      }"
    >
      <div class="flex-1">
        <router-link to="/" class="btn btn-ghost font-display font-medium text-xl">
          <span aria-hidden="true">&lt; gabe.cook &gt;</span>
          <span class="sr-only">Gabe Cook</span>
        </router-link>
      </div>

      <main-menu class="menu-horizontal hidden md:flex gap-2" />

      <details class="flex-none md:hidden dropdown dropdown-end" :open="showNav">
        <summary class="btn btn-square btn-ghost" @click.prevent="showNav = !showNav">
          <menu-icon />
          <span class="sr-only">Main menu</span>
        </summary>
        <main-menu
          class="dropdown-content bg-base-200 bg-opacity-20 overflow-hidden rounded-box z-40 mt-3 w-52 p-2 shadow-xl before:backdrop-blur before:backdrop-hack"
          :class="{ 'bg-opacity-80': !minimal }"
          @click="showNav = false"
        />
      </details>
    </nav>

    <div
      v-if="showNav"
      class="absolute w-full h-full opacity-0 z-20"
      aria-hidden="true"
      @click="showNav = false"
    />

    <router-view
      id="content"
      v-slot="{ Component }"
      class="flex-grow pt-24 md:pt-32 pb-10 px-5 sm:px-10"
    >
      <transition :name="transitionName" mode="out-in" appear>
        <keep-alive>
          <component :is="Component" />
        </keep-alive>
      </transition>
    </router-view>

    <div class="flex-grow" />
    <transition name="fade">
      <footer v-if="!minimal" class="w-full">
        <footer-map />
        <div
          class="footer gap-y-5 sm:grid-flow-col bg-base-200 dark:bg-neutral p-10 justify-center sm:justify-between items-center shadow-inner"
        >
          <div class="font-display font-medium text-xl" aria-hidden="true">
            &lt;/ gabe.cook &gt;
          </div>
          <div>
            <a href="//github.com/gabe565/portfolio" target="_blank" class="link link-hover">
              <github-icon />
              Source on GitHub
            </a>
          </div>
          <div>
            <a
              href="//github.com/gabe565/portfolio/blob/master/LICENSE"
              target="_blank"
              class="link link-hover"
            >
              &copy; {{ new Date().getFullYear() }} Gabe Cook
            </a>
          </div>
        </div>
      </footer>
    </transition>
  </div>
</template>
