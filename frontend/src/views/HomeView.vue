<template>
  <div>
    <div>
      <transition name="fade" appear>
        <div
          v-if="index !== null"
          :key="index"
          class="background-image position-absolute vw-100 vh-100 top-0"
          :style="{ backgroundImage: `url(${backgrounds[index]})` }"
        ></div>
      </transition>
    </div>
    <div class="row position-absolute g-0 align-items-center intro-body vw-100 vh-100">
      <div class="col-lg-8 mx-auto">
        <transition name="slide-right-slow" appear>
          <h1 class="display-1" style="transition-delay: 250ms">Gabe Cook</h1>
        </transition>
        <transition name="slide-right-slow" appear>
          <h2 class="h5" style="transition-delay: 500ms">DevOps Engineer / Software Developer</h2>
        </transition>
      </div>
    </div>
  </div>
</template>

<script setup>
import { onActivated, onDeactivated, ref } from "vue";
import pb from "../plugins/pocketbase";
import loadImage from "../util/loadImage";

const index = ref(null);
let timeout = null;

onActivated(() => {
  if (index.value !== null && !timeout) {
    startTimeout();
  }
});

onDeactivated(() => {
  stopTimeout();
  nextBackground();
});

let backgrounds = [];
const fetchBackgrounds = async () => {
  try {
    const response = await pb.collection("backgrounds").getFullList();
    backgrounds = response
      .map((background) => ({ sortKey: Math.random(), background }))
      .sort((a, b) => a.sortKey - b.sortKey)
      .map(({ background }) => background.url);
  } catch (error) {
    console.error(error);
  }
};

const nextBackground = async () => {
  const newIndex = (index.value + 1) % backgrounds.length;
  try {
    await loadImage(backgrounds[newIndex]);
    index.value = newIndex;
  } catch (error) {
    console.error(error);
  }
};

const startTimeout = () => {
  timeout = setTimeout(async () => {
    await nextBackground();
    startTimeout();
  }, 7500);
};

const stopTimeout = () => {
  if (timeout) {
    clearTimeout(timeout);
  }
  timeout = null;
};

(async () => {
  await fetchBackgrounds();
  await nextBackground();
  startTimeout();
})();
</script>

<style lang="scss" scoped>
.background-image {
  background-size: cover;
  background-position: center;
  opacity: 0.4;
}
</style>
