<template>
  <section class="container text-center space-y-7">
    <h1 class="font-display font-medium text-4xl">Projects</h1>
    <p>Here are some of my best projects!</p>
    <div v-if="loading" class="h-screen">
      <loading-icon />
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
        <portfolio-card
          :style="{ transitionDelay: `${i * 70}ms` }"
          :title="project.name"
          :description="project.description"
          :header-image="project.image"
          :header-alt="`Screenshot of ${project.name}`"
          :header-url="project.url"
          description-html
        >
          <template #title-append>
            <div
              v-if="project.archived"
              class="badge bg-transparent text-yellow-600 border-yellow-600 badge-sm"
            >
              Archived
            </div>
          </template>

          <template #footer>
            <div class="pb-4">
              <h3 class="sr-only">Tags</h3>
              <ul class="list-none">
                <li
                  v-for="tag in project.expand.tags"
                  :key="tag.id"
                  class="badge badge-outline mx-1"
                  :style="{ '--badge-color': tag.color }"
                >
                  {{ tag.title }}
                </li>
              </ul>
            </div>

            <a
              :href="project.url"
              class="btn btn-soft m-2 rounded-[calc(var(--radius-box)-.5rem)]"
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
          </template>
        </portfolio-card>
      </transition>
    </div>
  </section>
</template>

<script setup>
import { ref, shallowRef } from "vue";
import ErrorIcon from "~icons/material-symbols/error-outline-rounded";
import GlobeIcon from "~icons/mdi/web";
import GitHubIcon from "~icons/simple-icons/github";
import LoadingIcon from "~icons/svg-spinners/ring-resize";
import PortfolioCard from "@/components/PortfolioCard.vue";
import pb from "@/plugins/pocketbase";

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
