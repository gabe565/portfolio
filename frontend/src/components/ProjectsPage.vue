<template>
  <section id="skills" class="content-section text-center">
    <div class="container">
      <h1>Projects</h1>
      <p>
        Below are links to the personal or collaborative projects that I can
        claim a piece of.
      </p>
      <transition-group
        name="fade"
        tag="div"
        appear
        class="row justify-content-center"
      >
        <div
          v-for="project in projects"
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
              <h5 class="card-title">{{ project.name }}</h5>
              <!-- eslint-disable vue/no-v-html -->
              <p
                class="card-description my-auto"
                v-html="project.description"
              />
              <div>
                <h6 class="sr-only">Tags</h6>
                <span
                  v-for="tag in project.expand.tags"
                  :key="tag.id"
                  class="badge rounded-pill bg-info mx-1"
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
  data() {
    return {
      projects: {},
    };
  },
  async created() {
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
    }
  },
};
</script>

<style scoped>
.card-description:deep(p) {
  margin: 0;
}
</style>
