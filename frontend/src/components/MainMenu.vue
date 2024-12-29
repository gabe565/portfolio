<template>
  <ul class="menu">
    <li>
      <router-link to="/about">
        <info-icon />
        About
      </router-link>
    </li>
    <li>
      <router-link to="/skills">
        <list-icon />
        Skills
      </router-link>
    </li>
    <li>
      <router-link to="/projects">
        <code-icon />
        Projects
      </router-link>
    </li>
    <li>
      <router-link to="/connect">
        <mail-icon />
        Connect
      </router-link>
    </li>
  </ul>
</template>

<script setup>
import CodeIcon from "~icons/material-symbols/code-rounded";
import InfoIcon from "~icons/material-symbols/info";
import ListIcon from "~icons/material-symbols/list-rounded";
import MailIcon from "~icons/material-symbols/mail-rounded";
</script>
