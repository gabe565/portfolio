import { createRouter, createWebHistory } from "vue-router";
import { ApiPath } from "@/config/api";
import AboutView from "@/views/AboutView.vue";
import ConnectView from "@/views/ConnectView.vue";
import HomeView from "@/views/HomeView.vue";
import NotFoundView from "@/views/NotFoundView.vue";
import ProjectsView from "@/views/ProjectsView.vue";
import SkillsView from "@/views/SkillsView.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  linkActiveClass: "active",
  scrollBehavior(to, from, savedPosition) {
    return new Promise((resolve) => {
      setTimeout(() => {
        if (savedPosition) {
          resolve(savedPosition);
        } else {
          resolve({ top: 0 });
        }
      }, 200);
    });
  },
  routes: [
    {
      name: "Home",
      path: "/",
      component: HomeView,
    },
    {
      name: "About",
      path: "/about",
      component: AboutView,
    },
    {
      name: "Skills",
      path: "/skills",
      component: SkillsView,
    },
    {
      name: "Projects",
      path: "/projects",
      component: ProjectsView,
    },
    {
      name: "Connect",
      path: "/connect",
      component: ConnectView,
    },
    {
      path: "/_",
      redirect() {
        window.location.href = ApiPath("/_");
      },
    },
    {
      name: "Not Found",
      path: "/:wildcard(.*)",
      component: NotFoundView,
    },
  ],
});

router.beforeEach((to, from, next) => {
  if (to.name && to.name !== "Home") {
    document.title = to.name + " · " + " Gabe Cook";
  } else {
    document.title = "Gabe Cook";
  }
  next();
});

export default router;
