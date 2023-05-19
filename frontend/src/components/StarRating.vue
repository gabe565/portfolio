<template>
  <span>
    <span class="visually-hidden"
      >{{ props.modelValue }} of {{ props.total }} stars</span
    >
    <span
      aria-hidden="true"
      :title="`${props.modelValue} star${props.modelValue === 1 ? '' : 's'}`"
    >
      <template v-for="n in props.total" :key="n">
        <SvgIcon
          v-if="n <= props.modelValue"
          :icon="StarIcon"
          fixed-width
          fill
        />
        <SvgIcon
          v-else
          :icon="StarIcon"
          class="text-white opacity-15"
          fixed-width
        />
      </template>
    </span>
  </span>
</template>

<script setup>
import StarIcon from "feather-icons/dist/icons/star.svg";

const props = defineProps({
  modelValue: {
    type: Number,
    required: true,
  },
  total: {
    type: Number,
    default: 5,
  },
});
</script>

<style scoped>
svg {
  height: 1.15em;
}
</style>
