<template>
  <component
    :is="props.icon"
    :class="{
      'icon-inline': true,
      'icon-fw': props.fixedWidth,
      'icon-fill': props.fill,
    }"
    aria-hidden="true"
  />
</template>

<script setup>
const props = defineProps({
  icon: {
    type: String,
    required: true,
  },
  fixedWidth: Boolean,
  fill: Boolean,
});
</script>
