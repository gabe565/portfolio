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
        <div class="col-lg-6 me-auto">
          <p>
            I am a {{ age }} year old DevOps Engineer and Software Developer
            from Oklahoma City, OK. I graduated with my B.S. in Computer Science
            from the University of Central Oklahoma.
          </p>
          <p>
            I enjoy scripting things together, writing websites which have an
            easy-to-use interface, tinkering with self-hosted home automation,
            and just overall playing with technology!
          </p>
          <p>
            <router-link to="/connect"
              >Let me know if you have any questions!</router-link
            >
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
  data() {
    return {
      age: "",
      mapData,
    };
  },
  async created() {
    const diff = Date.now() - new Date("1995-05-26");
    this.age = new Date(diff).getUTCFullYear() - 1970;
  },
};
</script>
