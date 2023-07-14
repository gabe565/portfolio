<template>
  <section class="pt-page">
    <div class="container">
      <div class="row mb-4">
        <div class="col">
          <h1>About</h1>
        </div>
      </div>
      <div class="row mb-3">
        <div class="col-auto me-auto text-start">
          <div class="p-3 float-md-end d-block mx-auto" style="width: 190px">
            <img class="img-fluid rounded-4" src="../assets/me.jpg" alt="Gabe's Avatar" />
          </div>
          <p>Hi! I'm Gabe.</p>
          <p>
            I'm a DevOps Engineer and Software Developer based in Oklahoma City, Oklahoma, USA. I'm
            passionate about automating processes, crafting innovative solutions, and staying
            up-to-date with the latest technologies. I specialize in building robust systems using
            Golang and Vue.js, and I enjoy exploring new ideas and tinkering with self-hosted home
            infrastructure. My
            <router-link to="/projects">project</router-link> page showcases some of my most notable
            works.
          </p>
          <p>
            To ensure efficient development and deployment, I utilize containers, GitHub Actions,
            Kubernetes, and GitOps. These tools allow me to streamline the development lifecycle and
            ensure that my projects run seamlessly. If you're looking for a reliable, results-driven
            DevOps Engineer and Software Developer, look no further! I'm always excited to
            collaborate on new projects and leverage my skills to help organizations achieve their
            goals.
          </p>
          <p>
            <router-link to="/connect">Feel free to reach out!</router-link>
          </p>
        </div>
      </div>
      <section>
        <div class="row">
          <div class="col">
            <h2 class="h3">Stats</h2>
          </div>
        </div>
        <div class="row g-0 mb-3 justify-content-center">
          <div class="col-auto">
            <img
              :src="ApiPath('/api/github-stats/stats')"
              alt="Gabe Cook's GitHub stats"
              style="min-height: 195px"
            />
          </div>
          <div class="col-auto">
            <img
              :src="ApiPath('/api/github-stats/top-langs')"
              alt="Most Used Languages"
              style="min-height: 165px"
            />
          </div>
        </div>
      </section>
    </div>
    <section>
      <h2 class="h3 visually-hidden">Map of Oklahoma City</h2>
      <GMapMap
        :center="mapData.center"
        :options="mapData.options"
        style="width: 100%; height: 100%"
      ></GMapMap>
    </section>
  </section>
</template>

<script setup>
import mapData from "../data/mapOptions";
import { ApiPath } from "../config/api";
</script>
