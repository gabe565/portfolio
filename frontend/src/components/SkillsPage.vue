<template>
  <section id="skills" class="content-section text-center">
    <div class="container">
      <h1>Skills</h1>
      <div class="col-md-9 mx-auto">
        My development skills on a scale of 1-5 stars. Note that the scale is
        relative to the other languages that I have used, so a 5 star rating
        means that I am most familiar with that language (Not that I am perfect
        at it)!
      </div>
      <div
        v-if="loading"
        class="spinner-border text-primary mt-5"
        role="status"
      >
        <span class="visually-hidden">Loading...</span>
      </div>
      <transition-group v-else name="fade" appear>
        <div v-if="error" class="col mt-5">
          <div class="alert alert-danger">
            {{ error }}
          </div>
        </div>
        <div v-for="{ title, skills } in skills" v-else :key="title">
          <hr />
          <h3>{{ title }}</h3>
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
                      <div class="sr-only">{{ skill.rating }} of 5</div>
                      <div
                        class="float-end"
                        aria-hidden="true"
                        :title="`${skill.rating} Stars`"
                      >
                        <template v-for="n in 5" :key="n">
                          <template v-if="n <= skill.rating">
                            <font-awesome-icon
                              icon="fas fa-star"
                              fixed-width
                              class="text-primary"
                            />
                          </template>
                          <template v-else>
                            <font-awesome-icon
                              icon="far fa-star"
                              fixed-width
                              class="text-white opacity-15"
                            />
                          </template>
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
};
</script>

<style lang="scss">
.opacity-15 {
  opacity: 0.15;
}
</style>
