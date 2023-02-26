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
          <div class="card text-white bg-dark border-dark h-100">
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
                View at
                <span class="font-monospace">{{ project.pretty_url }}</span>
              </div>
            </a>
          </div>
        </div>
      </transition-group>
    </div>
  </section>
</template>

<script>
import pb from "@/plugins/pocketbase";

export default {
  data: () => ({
    projects: {},
    loading: true,
    error: null,
  }),
  async created() {
    await this.fetchData();
  },
  methods: {
    async fetchData() {
      try {
        const response = await pb.collection("projects").getFullList({
          expand: "tags",
        });
        this.projects = response.map((project) => {
          return {
            ...project,
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
        this.error = "Failed to fetch projects. Please try again later.";
      } finally {
        this.loading = false;
      }
    },
  },
};
</script>

<style scoped>
.card-description:deep(p) {
  margin: 0;
}
</style>
