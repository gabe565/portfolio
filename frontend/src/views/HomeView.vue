<template>
  <div>
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

<script setup>
import pb from "@/plugins/pocketbase";
import loadImage from "@/util/loadImage";

const index = ref(null);
const backgrounds = ref([]);
const timeout = ref(null);

onActivated(() => {
  if (index.value !== null && !timeout.value) {
    startTimeout();
  }
});

onDeactivated(() => {
  stopTimeout();
  randomBackground();
});

const fetchBackgrounds = async () => {
  try {
    const response = await pb.collection("backgrounds").getFullList();
    backgrounds.value = response.map((e) => e.url);
  } catch (error) {
    console.error(error);
  }
};

const randomBackground = async () => {
  let newIndex = Math.floor(Math.random() * backgrounds.value.length);
  if (newIndex === index.value) {
    newIndex += 1;
    newIndex %= backgrounds.value.length;
  }
  try {
    await loadImage(backgrounds.value[newIndex]);
    index.value = newIndex;
  } catch (error) {
    console.error(error);
  }
};

const startTimeout = () => {
  timeout.value = setTimeout(async () => {
    await randomBackground();
    startTimeout();
  }, 7500);
};

const stopTimeout = () => {
  if (timeout.value) {
    clearTimeout(timeout.value);
  }
  timeout.value = null;
};

(async () => {
  await fetchBackgrounds();
  await randomBackground();
  startTimeout();
})();
</script>
