<template>
  <section id="about" class="content-section text-center">
    <div class="container">
      <h1>About</h1>
      <hr />
      <div class="row">
        <div class="col-3 col-lg-2 mb-3 mx-auto ms-lg-auto me-lg-0">
          <transition name="fade">
            <img
              class="img-fluid rounded"
              src="../assets/me.jpg"
              alt="Photo of Gabe Cook"
            />
          </transition>
        </div>
        <div class="col-lg-9 me-auto text-left">
          <p>
            I'm a {{ age }} year old DevOps Engineer and Software Developer
            based in Oklahoma City, Oklahoma, USA.
            <br />
            I love automating things, building things with Golang and Vue,
            tinkering with self-hosted home automation, and learning new things!
            See my <router-link to="/projects">projects page</router-link> for a
            list of projects.
            <br />
            I build and deploy using containers, GitHub Actions, Kubernetes, and
            GitOps.
          </p>
          <p>
            <router-link to="/connect">Feel free to reach out!</router-link>
          </p>
        </div>
      </div>
      <hr />
    </div>
    <transition name="fade">
      <GMapMap
        :center="mapData.center"
        :options="mapData.options"
        style="width: 100%; height: 100%"
      ></GMapMap>
    </transition>
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
