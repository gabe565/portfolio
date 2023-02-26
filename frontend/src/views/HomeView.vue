<template>
  <div class="text-center bg-black text-white">
    <transition name="fade" appear>
      <div
        v-if="index !== null"
        :key="index"
        class="position-absolute vw-100 vh-100 top-0 opacity-50"
        :style="{
          backgroundImage: `url(${backgrounds[index]})`,
          backgroundSize: 'cover',
          backgroundPosition: 'center',
        }"
      ></div>
    </transition>
    <div
      class="row position-absolute g-0 align-items-center intro-body vw-100 vh-100"
    >
      <div class="col-lg-8 mx-auto">
        <h1 class="display-1">Gabe Cook</h1>
        <h2 class="h5">DevOps Engineer / Software Developer</h2>
      </div>
    </div>
  </div>
</template>

<script>
import pb from "@/plugins/pocketbase";
import loadImage from "@/util/loadImage";

export default {
  data: () => ({
    index: null,
    backgrounds: null,
    timeout: null,
  }),
  async created() {
    await this.fetchBackgrounds();
    await this.randomBackground();
    this.startTimeout();
  },
  activated() {
    if (this.index !== null && !this.timeout) {
      this.startTimeout();
    }
  },
  deactivated() {
    this.stopTimeout();
    this.randomBackground();
  },
  methods: {
    async fetchBackgrounds() {
      try {
        const response = await pb.collection("backgrounds").getFullList();
        this.backgrounds = response.map((e) => e.url);
      } catch (error) {
        console.error(error);
      }
    },
    async randomBackground() {
      let index = Math.floor(Math.random() * this.backgrounds.length);
      if (index === this.index) {
        index += 1;
        index %= this.backgrounds.length;
      }
      try {
        await loadImage(this.backgrounds[index]);
        this.index = index;
      } catch (error) {
        console.error(error);
      }
    },
    startTimeout() {
      this.timeout = setTimeout(async () => {
        await this.randomBackground();
        this.startTimeout();
      }, 7500);
    },
    stopTimeout() {
      if (this.timeout) {
        clearTimeout(this.timeout);
      }
      this.timeout = null;
    },
  },
};
</script>
