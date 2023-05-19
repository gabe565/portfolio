<template>
  <section class="pt-page">
    <div class="container pb-5">
      <div class="row mb-4">
        <div class="col">
          <h1>Projects</h1>
        </div>
      </div>
      <div class="row mb-3">
        <div class="col">
          <p>Here are some of my best projects!</p>
        </div>
      </div>
      <div v-if="loading" class="row">
        <div class="col">
          <div class="spinner-border text-primary" role="status">
            <span class="visually-hidden">Loading...</span>
          </div>
        </div>
      </div>
      <div v-else class="row justify-content-center">
        <div v-if="error" class="col">
          <div class="alert alert-danger">
            {{ error }}
          </div>
        </div>
        <div
          v-for="(project, i) in projects"
          v-else
          :key="`project-${project.name}`"
          class="col-10 col-lg-6 col-xl-4 pb-4"
        >
          <transition name="fade" appear>
            <section
              class="card text-white bg-dark border-dark h-100"
              :style="{ transitionDelay: `${i * 70}ms` }"
            >
              <template v-if="project.image">
                <a
                  :href="project.url"
                  class="card-link overflow-hidden"
                  target="_blank"
                >
                  <img
                    :src="project.image"
                    class="card-img-top"
                    :alt="`Screenshot of ${project.name}`"
                    :style="{ padding: `${project.image_padding}px` }"
                  />
                </a>
              </template>
              <div class="card-body d-flex flex-column">
                <h2 class="card-title h4">{{ project.name }}</h2>
                <!-- eslint-disable vue/no-v-html -->
                <p
                  class="card-description mb-auto text-start"
                  v-html="project.description"
                />
                <div>
                  <h3 class="visually-hidden">Tags</h3>
                  <span
                    v-for="tag in project.expand.tags"
                    :key="tag.id"
                    class="badge rounded-pill mx-1"
                    :class="{ 'bg-info': !tag.color }"
                    :style="{ backgroundColor: tag.color }"
                  >
                    {{ tag.title }}
                  </span>
                </div>
              </div>
              <a :href="project.url" class="card-link" target="_blank">
                <div class="card-footer">
                  <component
                    :is="project.icon"
                    v-if="project.icon"
                    class="icon-inline px-1"
                  />
                  <span class="font-monospace">
                    {{ project.pretty_url }}
                  </span>
                </div>
              </a>
            </section>
          </transition>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup>
import pb from "@/plugins/pocketbase";
import GitHubIcon from "feather-icons/dist/icons/github.svg";
import GlobeIcon from "feather-icons/dist/icons/globe.svg";

const projects = ref([]);
const loading = ref(true);
const error = ref(null);

const fetchData = async () => {
  try {
    const response = await pb.collection("projects").getFullList({
      expand: "tags",
    });
    projects.value = response.map((project) => {
      let image = "";
      if (project.image) {
        image = pb.getFileUrl(project, project.image);
      }
      return {
        ...project,
        icon: project.url.match(/^https:\/\/github\.com/)
          ? GitHubIcon
          : GlobeIcon,
        pretty_url: project.url
          .replace(/^(\w+:)?\/\//, "")
          .replace(/^github\.com\//, ""),
        image,
        expand: {
          tags: project.expand.tags?.sort((a, b) =>
            a.title.localeCompare(b.title),
          ),
        },
      };
    });
  } catch (error) {
    console.error(error);
    error.value = "Failed to fetch projects. Please try again later.";
  } finally {
    loading.value = false;
  }
};

fetchData();
</script>
