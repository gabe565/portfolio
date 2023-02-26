<template>
  <section id="about" class="content-section text-center">
    <div class="container">
      <div class="row">
        <div class="col">
          <h1>About</h1>
        </div>
      </div>
      <div class="row">
        <div class="col-3 col-lg-2 mx-auto ms-lg-auto me-lg-0">
          <img
            class="img-fluid rounded"
            src="../assets/me.jpg"
            alt="Photo of Gabe Cook"
          />
        </div>
        <div class="col-lg-9 me-auto text-left">
          <p>
            I'm a {{ age }} year old DevOps Engineer and Software Developer
            based in Oklahoma City, Oklahoma, USA.
          </p>
          <p>
            I love automating things, building things with Golang and Vue,
            tinkering with self-hosted home automation, and learning new things!
            See my <router-link to="/projects">projects page</router-link> for a
            list of projects.
          </p>
          <p>
            I build and deploy using containers, GitHub Actions, Kubernetes, and
            GitOps.
          </p>
          <p>
            <router-link to="/connect">Feel free to reach out!</router-link>
          </p>
        </div>
      </div>
      <div class="row g-0 justify-content-center">
        <div class="col-auto">
          <img
            src="https://github-readme-stats.vercel.app/api?username=gabe565&show_icons=true&theme=transparent&hide_border=true&count_private=true"
            alt="Gabe Cook's GitHub stats"
          />
        </div>
        <div class="col-auto">
          <img
            src="https://github-readme-stats.vercel.app/api/top-langs?username=gabe565&theme=transparent&hide_border=true&layout=compact"
            alt="Most Used Languages"
          />
        </div>
      </div>
      <hr />
    </div>
    <GMapMap
      :center="mapData.center"
      :options="mapData.options"
      style="width: 100%; height: 100%"
    ></GMapMap>
  </section>
</template>

<script>
import mapData from "@/data/mapOptions";

export default {
  data: () => ({
    age: "",
    mapData,
  }),
  async created() {
    const diff = Date.now() - new Date("1995-05-26");
    this.age = new Date(diff).getUTCFullYear() - 1970;
  },
};
</script>
