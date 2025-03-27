<template>
  <section class="container max-w-3xl text-center space-y-7">
    <h1 class="font-display font-medium text-4xl">Skills</h1>
    <div>
      A list of skills and languages that I'm proficient in on a scale of 1-5 stars. Note that the
      scale is relative to the other languages that I have used, so a 5 star rating means that I'm
      most familiar with that language, not that I'm perfect at it.
    </div>
    <div v-if="isLoading" class="h-screen">
      <loading-icon />
      <span class="sr-only">Loading...</span>
    </div>
    <div v-else-if="error" class="alert alert-error" role="alert">
      <error-icon />
      {{ error }}
    </div>
    <portfolio-card v-for="{ title, skills } in state" v-else :key="title" class="mb-7">
      <template #title>
        <h2 class="card-title self-center">{{ title }}</h2>
      </template>

      <ul
        class="grid grid-flow-row md:grid-flow-col auto-cols-fr list-none text-left"
        :style="{
          gridTemplateRows: `repeat(${Math.ceil(skills.length / 2)}, minmax(0, 1fr))`,
        }"
      >
        <li
          v-for="(skill, i) in skills"
          :key="skill.title"
          class="border-black/10 dark:border-white/10 px-8 py-3"
          :class="{ 'md:border-r': i < skills.length / 2 }"
        >
          <transition name="fade" appear>
            <div
              class="flex justify-between items-center"
              :style="{ transitionDelay: `${i * 40}ms` }"
            >
              <div>{{ skill.title }}</div>
              <StarRating v-model="skill.rating" class="text-blue-500 dark:text-blue-400" />
            </div>
          </transition>
        </li>
      </ul>
    </portfolio-card>
  </section>
</template>

<script setup>
import { useAsyncState } from "@vueuse/core";
import ErrorIcon from "~icons/material-symbols/error-outline-rounded";
import LoadingIcon from "~icons/svg-spinners/ring-resize";
import PortfolioCard from "@/components/PortfolioCard.vue";
import StarRating from "@/components/StarRating.vue";
import pb from "@/plugins/pocketbase";

const { state, isLoading, error } = useAsyncState(async () => {
  try {
    const response = await pb.collection("skill_headings").getFullList({
      expand: "skills_via_heading",
      sort: "order,title",
      fields: "title,expand.skills_via_heading.title,expand.skills_via_heading.rating",
    });

    return response
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
    throw "Failed to fetch skills. Please try again later.";
  }
});
</script>
