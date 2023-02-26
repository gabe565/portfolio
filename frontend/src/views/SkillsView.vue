<template>
  <section id="skills" class="content-section text-center">
    <div class="container">
      <div class="row">
        <div class="col">
          <h1>Skills</h1>
        </div>
      </div>
      <div class="row">
        <div class="col-md-9 mx-auto">
          A list of skills and languages that I'm proficient in on a scale of
          1-5 stars. Note that the scale is relative to the other languages that
          I have used, so a 5 star rating means that I'm most familiar with that
          language, not that I'm perfect at it.
        </div>
      </div>
      <div v-if="loading" class="row">
        <div class="col">
          <div class="spinner-border text-primary mt-5" role="status">
            <span class="visually-hidden">Loading...</span>
          </div>
        </div>
      </div>
      <transition-group v-else name="fade" tag="div" class="row" appear>
        <div v-if="error" class="col mt-5">
          <div class="alert alert-danger">
            {{ error }}
          </div>
        </div>
        <div v-for="{ title, skills } in skills" v-else :key="title">
          <div class="row">
            <div class="col">
              <h2 class="h3">{{ title }}</h2>
            </div>
          </div>
          <div class="row">
            <div
              v-for="(chunk, i) in [
                skills.slice(0, Math.ceil(skills.length / 2)),
                skills.slice(Math.ceil(skills.length / 2)),
              ]"
              :key="i"
              class="col-sm-8 col-lg-4 mx-auto"
              :class="[i === 0 ? 'ms-lg-auto me-lg-0' : 'me-lg-auto ms-lg-0']"
            >
              <ul class="list-group list-group-inverse">
                <li
                  v-for="skill in chunk"
                  :key="skill.title"
                  class="list-group-item"
                >
                  <div class="row">
                    <div class="col text-start">
                      {{ skill.title }}
                    </div>
                    <div class="col-auto">
                      <div class="sr-only">{{ skill.rating }} of 5 stars</div>
                      <div
                        class="float-end"
                        aria-hidden="true"
                        :title="`${skill.rating} star${
                          skill.rating === 1 ? '' : 's'
                        }`"
                      >
                        <template v-for="n in 5" :key="n">
                          <font-awesome-icon
                            v-if="n <= skill.rating"
                            icon="fas fa-star"
                            fixed-width
                            class="text-primary"
                          />
                          <font-awesome-layers
                            v-else-if="n === Math.ceil(skill.rating)"
                            fixed-width
                          >
                            <font-awesome-icon
                              icon="fas fa-star-half"
                              fixed-width
                              class="text-primary"
                            />
                            <font-awesome-icon
                              icon="far fa-star"
                              fixed-width
                              class="text-white opacity-15"
                            />
                          </font-awesome-layers>
                          <font-awesome-icon
                            v-else
                            icon="far fa-star"
                            fixed-width
                            class="text-white opacity-15"
                          />
                        </template>
                      </div>
                    </div>
                  </div>
                </li>
              </ul>
            </div>
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
    skills: null,
    error: null,
    loading: true,
  }),
  async created() {
    await this.fetchData();
  },
  methods: {
    async fetchData() {
      try {
        const response = await pb.collection("skill_headings").getFullList({
          expand: "skills(heading)",
        });
        this.skills = response
          .sort((a, b) => a.order - b.order || a.title.localeCompare(b.title))
          .map(({ title, expand }) => {
            return {
              title: title,
              skills: expand["skills(heading)"]?.sort(
                (a, b) => b.rating - a.rating || a.title.localeCompare(b.title),
              ),
            };
          })
          .filter(({ skills }) => skills);
        this.ready = true;
      } catch (error) {
        console.error(error);
        this.error = "Failed to fetch skills. Please try again later.";
      } finally {
        this.loading = false;
      }
    },
  },
};
</script>

<style lang="scss">
.opacity-15 {
  opacity: 0.15;
}
</style>
