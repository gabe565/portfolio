<template>
  <ul class="menu">
    <li>
      <router-link to="/about" active-class="menu-active">
        <info-icon />
        About
      </router-link>
    </li>
    <li>
      <router-link to="/skills" active-class="menu-active">
        <list-icon />
        Skills
      </router-link>
    </li>
    <li>
      <router-link to="/projects" active-class="menu-active">
        <code-icon />
        Projects
      </router-link>
    </li>
    <li>
      <router-link to="/connect" active-class="menu-active">
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
