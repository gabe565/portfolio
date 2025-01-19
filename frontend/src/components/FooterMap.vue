<script setup>
import { ApiPath } from "@/config/api.js";
</script>

<template>
  <picture class="w-full">
    <source
      :srcset="ApiPath('/api/map/xxl-dark.png')"
      media="(min-width: 2000px) and (prefers-color-scheme: dark)"
    />
    <source
      :srcset="ApiPath('/api/map/xl-dark.png')"
      media="(min-width: 1400px) and (prefers-color-scheme: dark)"
    />
    <source
      :srcset="ApiPath('/api/map/lg-dark.png')"
      media="(min-width: 992px) and (prefers-color-scheme: dark)"
    />
    <source
      :srcset="ApiPath('/api/map/md-dark.png')"
      media="(min-width: 768px) and (prefers-color-scheme: dark)"
    />
    <source
      :srcset="ApiPath('/api/map/sm-dark.png')"
      media="(min-width: 576px) and (prefers-color-scheme: dark)"
    />
    <source
      :srcset="ApiPath('/api/map/xxl-light.png')"
      media="(min-width: 2000px) and (prefers-color-scheme: light)"
    />
    <source
      :srcset="ApiPath('/api/map/xl-light.png')"
      media="(min-width: 1400px) and (prefers-color-scheme: light)"
    />
    <source
      :srcset="ApiPath('/api/map/lg-light.png')"
      media="(min-width: 992px) and (prefers-color-scheme: light)"
    />
    <source
      :srcset="ApiPath('/api/map/md-light.png')"
      media="(min-width: 768px) and (prefers-color-scheme: light)"
    />
    <source
      :srcset="ApiPath('/api/map/sm-light.png')"
      media="(min-width: 576px) and (prefers-color-scheme: light)"
    />
    <source :srcset="ApiPath('/api/map/xs-light.png')" media="(prefers-color-scheme: light)" />
    <source :srcset="ApiPath('/api/map/xs-dark.png')" />
    <img
      :src="ApiPath(`/api/map/lg-dark.png`)"
      alt="Map of Oklahoma City"
      class="w-full h-[200px] lg:h-[300px] object-cover"
    />
  </picture>
</template>
