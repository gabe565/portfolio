<template>
  <section class="container max-w-5xl mx-auto text-center space-y-7">
    <h1 class="font-display font-medium text-4xl">About</h1>
    <portfolio-card>
      <div class="text-start space-y-4">
        <div class="p-3 md:float-end block mx-auto" style="width: 190px">
          <img class="rounded-lg" :src="headerImage" alt="Gabe's Avatar" />
        </div>

        <p class="text-lg">Hi! I'm Gabe.</p>
        <p>
          I'm a DevOps Engineer and Software Developer based in Oklahoma City, Oklahoma, USA. I'm
          passionate about automating processes, crafting innovative solutions, and staying
          up-to-date with the latest technologies. I specialize in building robust systems using
          Golang and Vue.js, and I enjoy exploring new ideas and tinkering with self-hosted home
          infrastructure. My
          <router-link to="/projects">project</router-link>
          page showcases some of my most notable works.
        </p>
        <p>
          To ensure efficient development and deployment, I utilize containers, GitHub Actions,
          Kubernetes, and GitOps. These tools allow me to streamline the development lifecycle and
          ensure that my projects run seamlessly. If you're looking for a reliable, results-driven
          DevOps Engineer and Software Developer, look no further! I'm always excited to collaborate
          on new projects and leverage my skills to help organizations achieve their goals.
        </p>
        <p>
          <router-link to="/connect">Feel free to reach out!</router-link>
        </p>
      </div>
    </portfolio-card>

    <div class="flex flex-wrap gap-7 w-full justify-center">
      <img
        :src="ApiPath('/api/github-stats/stats')"
        alt="Gabe Cook's GitHub stats"
        class="card bg-base-100 card-border border-base-300 shadow-xs"
      />
      <img
        :src="ApiPath('/api/github-stats/top-langs')"
        alt="Most Used Languages"
        class="card bg-base-100 card-border border-base-300 shadow-xs"
      />
    </div>
  </section>
</template>

<script setup>
import headerImage from "@/assets/me.jpg";
import PortfolioCard from "@/components/PortfolioCard.vue";
import { ApiPath } from "@/config/api";
</script>
