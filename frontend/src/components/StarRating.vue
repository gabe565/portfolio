<template>
  <span>
    <span class="visually-hidden">{{ props.modelValue }} of {{ props.total }} stars</span>
    <span
      aria-hidden="true"
      :title="`${props.modelValue} star${props.modelValue === 1 ? '' : 's'}`"
    >
      <template v-for="n in props.total" :key="n">
        <star-filled-icon v-if="n <= props.modelValue" fixed-width fill />
        <star-outline-icon v-else class="text-white opacity-15" fixed-width />
      </template>
    </span>
  </span>
</template>

<script setup>
import StarFilledIcon from "~icons/material-symbols/star-rounded";
import StarOutlineIcon from "~icons/material-symbols/star-outline-rounded";

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
