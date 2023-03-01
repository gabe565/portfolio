<template>
  <section>
    <div class="container">
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
      <transition-group
        v-else
        name="fade"
        tag="div"
        appear
        class="row justify-content-center"
      >
        <div v-if="error" class="col">
          <div class="alert alert-danger">
            {{ error }}
          </div>
        </div>
        <div
          v-for="project in projects"
          v-else
          :key="`project-${project.name}`"
          class="col-10 col-lg-6 col-xl-4 pb-4"
        >
          <section class="card text-white bg-dark border-dark h-100">
            <template v-if="project.image_path">
              <a
                :href="project.url"
                class="card-link overflow-hidden"
                target="_blank"
              >
                <img
                  :src="project.image_path"
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
                class="card-description my-auto"
                v-html="project.description"
              />
              <div>
                <h3 class="sr-only">Tags</h3>
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
                <font-awesome-icon
                  v-if="project.icon"
                  :icon="project.icon"
                  fixed-width
                  class="px-1"
                />
                <span class="font-monospace">
                  {{ project.pretty_url }}
                </span>
              </div>
            </a>
          </section>
        </div>
      </transition-group>
    </div>
  </section>
</template>

<script setup>
import pb from "@/plugins/pocketbase";
import { faGithub } from "@fortawesome/free-brands-svg-icons";
import { faGlobe } from "@fortawesome/pro-solid-svg-icons";

const projects = ref([]);
const loading = ref(true);
const error = ref(null);

const fetchData = async () => {
  try {
    const response = await pb.collection("projects").getFullList({
      expand: "tags",
    });
    projects.value = response.map((project) => {
      return {
        ...project,
        icon: project.url.match(/^https:\/\/github\.com/) ? faGithub : faGlobe,
        pretty_url: project.url
          .replace(/^(\w+:)?\/\//, "")
          .replace(/^github\.com\//, ""),
        expand: {
          tags: project.expand.tags.sort((a, b) =>
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

<style scoped>
.card-description:deep(p) {
  margin: 0;
}
</style>