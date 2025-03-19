<template>
  <div
    class="flex flex-col absolute justify-center text-center z-10 w-full h-full text-black dark:text-white"
    v-bind="$attrs"
    @click="nextBackground(true)"
  >
    <transition name="title-slide" appear>
      <h1 class="font-display font-light text-5xl sm:text-7xl" style="transition-delay: 250ms">
        Gabe Cook
      </h1>
    </transition>
    <transition name="title-slide" appear>
      <h2 class="font-light sm:text-2xl" style="transition-delay: 500ms">
        DevOps Engineer / Software Developer
      </h2>
    </transition>

    <teleport to="body">
      <transition name="fade" appear>
        <div
          v-if="active && index !== null"
          class="absolute top-0 w-full h-lvh bg-white dark:bg-black"
        />
      </transition>
      <transition name="fade" appear>
        <div
          v-if="active && index !== null"
          :key="index"
          class="absolute top-0 w-full h-lvh bg-cover bg-center opacity-65 dark:opacity-50"
          :style="{ backgroundImage: `url(${backgrounds[index]})` }"
        />
      </transition>
    </teleport>
  </div>
</template>

<script setup>
import { onActivated, onDeactivated, ref } from "vue";
import pb from "@/plugins/pocketbase";
import loadImage from "@/util/loadImage";

const index = ref(null);
let timeout = null;
const active = ref(true);
const transitionName = ref("fade");

onActivated(() => {
  transitionName.value = "fade";
  active.value = true;
  if (index.value !== null && !timeout) {
    startTimeout();
  }
});

onDeactivated(() => {
  transitionName.value = "slide-left";
  active.value = false;
  stopTimeout();
  nextBackground();
});

let backgrounds = [];
const fetchBackgrounds = async () => {
  try {
    const response = await pb.collection("backgrounds").getFullList({
      fields: "url",
    });
    backgrounds = response
      .map((background) => ({ sortKey: Math.random(), background }))
      .sort((a, b) => a.sortKey - b.sortKey)
      .map(({ background }) => background.url);
  } catch (error) {
    console.error(error);
  }
};

const nextBackground = async (resetTimeout = false) => {
  if (backgrounds.length !== 0) {
    const newIndex = (index.value + 1) % backgrounds.length;
    try {
      await loadImage(backgrounds[newIndex]);
      index.value = newIndex;
    } catch (error) {
      console.error(error);
    }
  }
  if (resetTimeout) {
    startTimeout();
  }
};

const startTimeout = () => {
  stopTimeout();
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
  await nextBackground(true);
})();
</script>
