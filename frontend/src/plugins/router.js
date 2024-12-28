import { createRouter, createWebHistory } from "vue-router";
import { ApiPath } from "../config/api";
import AboutView from "../views/AboutView.vue";
import ConnectView from "../views/ConnectView.vue";
import HomeView from "../views/HomeView.vue";
import NotFoundView from "../views/NotFoundView.vue";
import ProjectsView from "../views/ProjectsView.vue";
import SkillsView from "../views/SkillsView.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  linkActiveClass: "active",
  routes: [
    {
      path: "/",
      component: HomeView,
      meta: { title: "Home" },
    },
    {
      path: "/about",
      component: AboutView,
      meta: { title: "About" },
    },
    {
      path: "/skills",
      component: SkillsView,
      meta: { title: "Skills" },
    },
    {
      path: "/projects",
      component: ProjectsView,
      meta: { title: "Projects" },
    },
    {
      path: "/connect",
      component: ConnectView,
      meta: { title: "Connect" },
    },
    {
      path: "/_",
      redirect() {
        window.location.href = ApiPath("/_");
      },
    },
    {
      path: "/:wildcard(.*)",
      component: NotFoundView,
      meta: { title: "Not Found" },
    },
  ],
});

router.beforeEach((to, from, next) => {
  if (to.meta.title && to.meta.title !== "Home") {
    document.title = to.meta.title + " · " + " Gabe Cook";
  } else {
    document.title = "Gabe Cook";
  }
  next();
});

export default router;
