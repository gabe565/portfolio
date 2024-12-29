<template>
  <span>
    <span class="sr-only">{{ props.modelValue }} of {{ props.total }} stars</span>
    <span
      aria-hidden="true"
      :title="`${props.modelValue} star${props.modelValue === 1 ? '' : 's'}`"
      class="whitespace-nowrap"
    >
      <template v-for="n in props.total" :key="n">
        <star-filled-icon v-if="n <= props.modelValue" fixed-width fill />
        <star-outline-icon v-else class="text-white opacity-15" fixed-width />
      </template>
    </span>
  </span>
</template>

<script setup>
import StarOutlineIcon from "~icons/material-symbols/star-outline-rounded";
import StarFilledIcon from "~icons/material-symbols/star-rounded";

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
