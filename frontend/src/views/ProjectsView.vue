<template>
  <section class="container text-center space-y-4">
    <h1 class="font-display font-medium text-4xl">Projects</h1>
    <p>Here are some of my best projects!</p>
    <div v-if="loading">
      <loading-icon class="inline" />
      <span class="sr-only">Loading...</span>
    </div>
    <div v-else-if="error" class="alert alert-error max-w-3xl place-self-center" role="alert">
      <error-icon />
      {{ error }}
    </div>
    <div v-else class="grid gap-7 grid-cols-1 md:grid-cols-2 lg:grid-cols-3 items-stretch">
      <transition
        v-for="(project, i) in projects"
        :key="`project-${project.name}`"
        name="fade"
        appear
      >
        <section
          class="card shadow-xl dark:shadow-2xl flex flex-auto"
          :style="{ transitionDelay: `${i * 70}ms` }"
        >
          <figure v-if="project.image">
            <a :href="project.url" class="overflow-hidden w-full" target="_blank">
              <img :src="project.image" class="w-full" :alt="`Screenshot of ${project.name}`" />
            </a>
          </figure>
          <div class="card-body text-start">
            <div class="flex flex-wrap items-center gap-2">
              <h2 class="card-title">{{ project.name }}</h2>
              <div v-if="project.archived" class="badge badge-outline badge-sm text-yellow-600">
                Archived
              </div>
            </div>
            <!-- eslint-disable vue/no-v-html -->
            <p class="card-description mt-auto" v-html="project.description" />
          </div>
          <div class="pb-4">
            <h3 class="sr-only">Tags</h3>
            <ul class="list-none">
              <li
                v-for="tag in project.expand.tags"
                :key="tag.id"
                class="badge badge-outline mx-1 text-white"
                :style="{ color: tag.color }"
              >
                {{ tag.title }}
              </li>
            </ul>
          </div>

          <a
            :href="project.url"
            class="btn dark:bg-neutral btn-block rounded-t-none border-0 border-t"
            target="_blank"
          >
            <component
              :is="project.icon"
              v-if="project.icon"
              class="px-1"
              style="font-size: 1.25em"
            />
            <span class="font-mono">
              {{ project.pretty_url }}
            </span>
          </a>
        </section>
      </transition>
    </div>
  </section>
</template>

<script setup>
import { ref, shallowRef } from "vue";
import pb from "@/plugins/pocketbase";
import ErrorIcon from "~icons/material-symbols/error-outline-rounded";
import GlobeIcon from "~icons/mdi/web";
import GitHubIcon from "~icons/simple-icons/github";
import LoadingIcon from "~icons/svg-spinners/ring-resize";

const projects = shallowRef([]);
const loading = ref(true);
const error = ref();

const fetchData = async () => {
  try {
    const response = await pb.collection("projects").getFullList({
      sort: "archived,-priority,created",
      expand: "tags",
      fields:
        "collectionId,id,name,url,description,image,archived,expand.tags.title,expand.tags.color",
    });
    projects.value = response.map((project) => {
      let image = "";
      if (project.image) {
        image = pb.files.getURL(project, project.image);
      }
      let icon = GlobeIcon;
      if (project.url.startsWith("https://github.com/")) {
        icon = GitHubIcon;
      }
      return {
        ...project,
        icon,
        pretty_url: project.url.replace(/^(\w+:)?\/\//, "").replace(/^github\.com\//, ""),
        image,
        expand: {
          tags: project.expand.tags?.sort((a, b) => a.title.localeCompare(b.title)),
        },
      };
    });
  } catch (err) {
    console.error(err);
    error.value = "Failed to fetch projects. Please try again later.";
  } finally {
    loading.value = false;
  }
};

fetchData();
</script>
