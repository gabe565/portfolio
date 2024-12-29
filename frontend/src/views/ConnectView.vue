<template>
  <section class="container max-w-2xl text-center space-y-4">
    <h1 class="font-display font-medium text-4xl">Connect</h1>
    <div>Here are some ways to reach out or see what I'm up to!</div>
    <h2 class="font-medium text-2xl">Accounts</h2>
    <ul class="flex flex-row gap-3 list-none justify-center">
      <li>
        <a :href="ApiPath('/to/github')" class="btn btn-primary" target="_blank">
          <github-icon />
          Github
        </a>
      </li>
      <li>
        <a :href="ApiPath('/to/linkedin')" class="btn btn-primary" target="_blank">
          <linkedin-icon />
          LinkedIn
        </a>
      </li>
    </ul>

    <h2 class="font-medium text-2xl">Contact</h2>
    <form
      id="needs-validation"
      ref="form"
      data-focus="false"
      method="post"
      :action="ApiPath('/api/mail')"
      role="form"
      novalidate
      class="flex flex-col gap-3 [&.validate_.input:has(input:invalid)]:input-error [&.validate_textarea:invalid]:input-error"
      @submit.prevent="submit"
    >
      <div v-if="error" class="alert alert-error" role="alert">
        <error-icon />
        {{ error }}
      </div>
      <div v-else-if="success" class="alert alert-success" role="alert">
        <check-icon />
        Your message has been sent successfully. Thank you!
      </div>

      <div class="join join-vertical md:join-horizontal">
        <!--Name-->
        <label class="input input-bordered join-item flex items-center gap-2 w-full">
          <account-icon />
          <span class="sr-only">Name</span>
          <input
            v-model="formData.name"
            name="name"
            placeholder="Name"
            class="grow"
            type="text"
            required
          />
        </label>

        <!--Email-->
        <label class="input input-bordered join-item flex items-center gap-2 w-full">
          <at-icon />
          <span class="sr-only">Email</span>
          <input
            id="emailInput"
            v-model="formData.email"
            name="email"
            placeholder="Email"
            class="grow"
            type="email"
            required
          />
        </label>
      </div>

      <!--Message-->
      <label class="form-control">
        <span class="sr-only">Message</span>
        <textarea
          id="messageInput"
          v-model="formData.message"
          class="textarea textarea-bordered max-h-96"
          name="text"
          required
          placeholder="Message"
          :style="{ height: textareaHeight }"
          @input="textareaInput"
        />
      </label>

      <div ref="captcha" class="h-[65px]" />

      <!-- Button -->
      <button class="btn btn-primary place-self-center">
        <loading-icon v-if="loading" />
        <send-icon v-else />
        Send
      </button>
    </form>
  </section>
</template>

<script setup>
import { nextTick, onActivated, onMounted, ref } from "vue";
import { ApiPath } from "@/config/api";
import pb from "@/plugins/pocketbase";
import { TurnstileEnabled, TurnstileKey, TurnstileReady, loadTurnstile } from "@/plugins/turnstile";
import AtIcon from "~icons/material-symbols/alternate-email-rounded";
import CheckIcon from "~icons/material-symbols/check-circle-outline-rounded";
import ErrorIcon from "~icons/material-symbols/error-outline-rounded";
import AccountIcon from "~icons/material-symbols/person-rounded";
import SendIcon from "~icons/material-symbols/send-rounded";
import GithubIcon from "~icons/simple-icons/github";
import LinkedinIcon from "~icons/simple-icons/linkedin";
import LoadingIcon from "~icons/svg-spinners/ring-resize";

const formData = ref({
  name: "",
  email: "",
  message: "",
});

const error = ref(null);
const success = ref(false);
const loading = ref(false);

const form = ref(null);
const textareaHeight = ref(null);
let captchaValue;

const textareaInput = async ($event) => {
  textareaHeight.value = "";
  await nextTick();
  textareaHeight.value = `${$event.target.scrollHeight + 2}px`;
};

const submit = async () => {
  const valid = form.value.checkValidity();
  if (valid) {
    loading.value = true;
    form.value.classList.remove("validate");
    success.value = false;
    error.value = null;
    const wait = new Promise((resolve) => setTimeout(resolve, 1000));
    try {
      await pb.collection("contact_form").create(formData.value, {
        headers: { "X-Captcha": captchaValue },
      });
      await wait;
      success.value = true;
    } catch (e) {
      console.error(e);
      error.value = e;
    } finally {
      loading.value = false;
    }
  } else {
    form.value.classList.add("validate");
    error.value = Error("Please fix invalid fields");
  }
};

const captcha = ref(null);

if (TurnstileEnabled) {
  let captchaId;

  onMounted(async () => {
    loadTurnstile();

    await TurnstileReady;
    captchaId = window.turnstile.render(captcha.value, {
      sitekey: TurnstileKey,
      action: "contact",
      callback: (token) => (captchaValue = token),
    });
  });

  onActivated(() => {
    if (captchaId) {
      window.turnstile.reset(captchaId);
    }
  });
}
</script>
