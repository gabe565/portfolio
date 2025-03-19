<template>
  <section class="card bg-base-100 card-border border-base-300 flex flex-auto shadow-xs">
    <figure v-if="headerImage" class="p-2">
      <a v-if="headerLink" :href="headerLink" class="overflow-hidden w-full" target="_blank">
        <img
          :src="headerImage"
          class="w-full rounded-[calc(var(--radius-box)-.5rem)]"
          :alt="headerAlt"
        />
      </a>
      <img
        v-else
        :src="headerImage"
        class="w-full rounded-[calc(var(--radius-box)-.5rem)]"
        :alt="headerAlt"
      />
    </figure>

    <div class="card-body text-start">
      <slot name="title">
        <h2 v-if="title" class="card-title">
          {{ title }}
          <slot name="title-append" />
        </h2>
      </slot>

      <slot name="description">
        <template v-if="description">
          <!-- eslint-disable vue/no-v-html -->
          <p v-if="descriptionHtml" v-html="description" />
          <p v-else>{{ description }}</p>
        </template>
      </slot>

      <slot />
    </div>

    <slot name="footer" />
  </section>
</template>

<script setup>
defineProps({
  headerImage: {
    type: String,
    default: "",
  },
  headerAlt: {
    type: String,
    default: "",
  },
  headerLink: {
    type: String,
    default: "",
  },
  title: {
    type: String,
    default: "",
  },
  description: {
    type: String,
    default: "",
  },
  descriptionHtml: {
    type: Boolean,
    default: false,
  },
});
</script>
