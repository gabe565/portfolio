<template>
  <section class="container max-w-3xl text-center space-y-4">
    <h1 class="font-display font-medium text-4xl">Skills</h1>
    <div>
      A list of skills and languages that I'm proficient in on a scale of 1-5 stars. Note that the
      scale is relative to the other languages that I have used, so a 5 star rating means that I'm
      most familiar with that language, not that I'm perfect at it.
    </div>
    <div v-if="loading">
      <loading-icon class="inline" />
      <span class="sr-only">Loading...</span>
    </div>
    <div v-else-if="error" class="alert alert-error" role="alert">
      <error-icon />
      {{ error }}
    </div>
    <section v-for="{ title, skills } in skills" v-else :key="title" class="mb-4 space-y-4">
      <h2 class="font-medium text-xl">{{ title }}</h2>
      <ul
        class="grid grid-flow-row md:grid-flow-col auto-cols-fr list-none text-left"
        :style="{
          gridTemplateRows: `repeat(${Math.ceil(skills.length / 2)}, minmax(0, 1fr))`,
        }"
      >
        <li
          v-for="(skill, i) in skills"
          :key="skill.title"
          class="border-white border-opacity-10 px-8 py-3"
          :class="{ 'md:border-r': i < skills.length / 2 }"
        >
          <transition name="fade" appear>
            <div
              class="flex justify-between items-center"
              :style="{ transitionDelay: `${i * 40}ms` }"
            >
              <div>{{ skill.title }}</div>
              <StarRating v-model="skill.rating" class="text-primary" />
            </div>
          </transition>
        </li>
      </ul>
    </section>
  </section>
</template>

<script setup>
import { ref } from "vue";
import StarRating from "@/components/StarRating.vue";
import pb from "@/plugins/pocketbase";
import ErrorIcon from "~icons/material-symbols/error-outline-rounded";
import LoadingIcon from "~icons/svg-spinners/ring-resize";

const skills = ref([]);
const loading = ref(true);
const error = ref();

const fetchData = async () => {
  try {
    const response = await pb.collection("skill_headings").getFullList({
      expand: "skills_via_heading",
      sort: "order,title",
      fields: "title,expand.skills_via_heading.title,expand.skills_via_heading.rating",
    });
    skills.value = response
      .map(({ title, expand }) => {
        return {
          title: title,
          skills: expand["skills_via_heading"]?.sort(
            (a, b) => b.rating - a.rating || a.title.localeCompare(b.title),
          ),
        };
      })
      .filter(({ skills }) => skills);
  } catch (err) {
    console.error(err);
    error.value = "Failed to fetch skills. Please try again later.";
  } finally {
    loading.value = false;
  }
};

fetchData();
</script>
