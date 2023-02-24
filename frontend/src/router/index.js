import { createRouter, createWebHistory } from "vue-router";
import Home from "@/components/HomePage.vue";
import About from "@/components/AboutPage.vue";
import Skills from "@/components/SkillsPage.vue";
import Projects from "@/components/ProjectsPage.vue";
import Connect from "@/components/ConnectPage.vue";
import NotFound from "@/components/NotFoundPage.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  linkActiveClass: "active",
  routes: [
    {
      path: "/",
      component: Home,
      meta: { title: "Home" },
    },
    {
      path: "/about",
      component: About,
      meta: { title: "About" },
    },
    {
      path: "/skills",
      component: Skills,
      meta: { title: "Skills" },
    },
    {
      path: "/projects",
      component: Projects,
      meta: { title: "Projects" },
    },
    {
      path: "/connect",
      component: Connect,
      meta: { title: "Connect" },
    },
    {
      path: "/:wildcard(.*)",
      component: NotFound,
      meta: { title: "Not Found" },
    },
  ],
});

router.beforeEach((to, from, next) => {
  if (to.meta.title && to.meta.title !== "Home") {
    document.title = to.meta.title + " | " + " Gabe Cook";
  } else {
    document.title = "Gabe Cook";
  }
  next();
});

export default router;
