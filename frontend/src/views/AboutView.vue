<template>
  <div class="flex flex-col w-full">
    <section class="container flex-grow max-w-7xl mx-auto text-center space-y-4">
      <h1 class="font-display font-medium text-4xl">About</h1>
      <div class="text-start space-y-4">
        <div class="p-3 md:float-end block mx-auto" style="width: 190px">
          <img class="rounded-lg" src="@/assets/me.jpg" alt="Gabe's Avatar" />
        </div>
        <p>Hi! I'm Gabe.</p>
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

      <h2 class="font-medium text-2xl">Stats</h2>
      <div class="flex flex-col md:flex-row justify-center items-center">
        <img
          :src="ApiPath('/api/github-stats/stats')"
          alt="Gabe Cook's GitHub stats"
          class="mw-100"
          style="min-height: 195px"
        />
        <img
          :src="ApiPath('/api/github-stats/top-langs')"
          alt="Most Used Languages"
          class="mw-100"
          style="min-height: 165px"
        />
      </div>
    </section>
  </div>
</template>

<script setup>
import { ApiPath } from "@/config/api";
</script>
