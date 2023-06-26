<template>
  <section class="text-center pt-page pb-5">
    <div class="container">
      <div class="row mb-4">
        <div class="col">
          <h1>Skills</h1>
        </div>
      </div>
      <div class="row mb-4">
        <div class="col-md-9 mx-auto">
          A list of skills and languages that I'm proficient in on a scale of 1-5 stars. Note that
          the scale is relative to the other languages that I have used, so a 5 star rating means
          that I'm most familiar with that language, not that I'm perfect at it.
        </div>
      </div>
      <div v-if="loading" class="row">
        <div class="col">
          <div class="spinner-border text-primary mt-5" role="status">
            <span class="visually-hidden">Loading...</span>
          </div>
        </div>
      </div>
      <div v-if="error" class="col">
        <div class="alert alert-danger" role="alert">
          {{ error }}
        </div>
      </div>
      <section v-for="{ title, skills } in skills" v-else :key="title">
        <div class="row mb-2">
          <div class="col">
            <h2 class="h3">{{ title }}</h2>
          </div>
        </div>
        <div class="row mb-4">
          <template
            v-for="(chunk, i) in [
              skills.slice(0, Math.ceil(skills.length / 2)),
              skills.slice(Math.ceil(skills.length / 2)),
            ]"
            :key="i"
          >
            <div
              class="col-sm-8 col-lg-5 col-xl-4 mx-auto"
              :class="[i === 0 ? 'ms-lg-auto me-lg-0' : 'me-md-auto ms-lg-0']"
            >
              <ul class="list-group list-group-inverse">
                <li v-for="(skill, j) in chunk" :key="skill.title" class="list-group-item">
                  <transition name="fade" appear>
                    <div class="row" :style="{ transitionDelay: `${i * 250 + j * 50}ms` }">
                      <div class="col text-start">
                        {{ skill.title }}
                      </div>
                      <div class="col-auto">
                        <div class="float-end">
                          <StarRating v-model="skill.rating" class="text-primary" />
                        </div>
                      </div>
                    </div>
                  </transition>
                </li>
              </ul>
            </div>
            <div
              v-if="i === 0"
              class="col-auto g-0 d-none d-lg-block border-end border-opacity-10 border-light"
            />
          </template>
        </div>
      </section>
    </div>
  </section>
</template>

<script setup>
import { ref } from "vue";
import pb from "../plugins/pocketbase";
import StarRating from "../components/StarRating.vue";

const skills = ref([]);
const loading = ref(true);
const error = ref();

const fetchData = async () => {
  try {
    const response = await pb.collection("skill_headings").getFullList({
      expand: "skills(heading)",
    });
    skills.value = response
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
  } catch (err) {
    console.error(err);
    error.value = "Failed to fetch skills. Please try again later.";
  } finally {
    loading.value = false;
  }
};

fetchData();
</script>

<style lang="scss">
.opacity-15 {
  opacity: 0.15;
}
</style>
