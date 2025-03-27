<template>
  <section class="container max-w-2xl text-center space-y-7">
    <h1 class="font-display font-medium text-4xl">Connect</h1>
    <div>Here are some ways to reach out or see what I'm up to!</div>

    <portfolio-card>
      <div class="flex flex-row flex-wrap gap-4 justify-end">
        <h2 class="card-title grow">Accounts</h2>
        <ul class="flex flex-row flex-wrap gap-3 list-none justify-center">
          <li>
            <a :href="ApiPath('/api/to/github')" class="btn btn-primary" target="_blank">
              <github-icon />
              Github
            </a>
          </li>
          <li>
            <a :href="ApiPath('/api/to/linkedin')" class="btn btn-primary" target="_blank">
              <linkedin-icon />
              LinkedIn
            </a>
          </li>
        </ul>
      </div>
    </portfolio-card>

    <portfolio-card title="Contact">
      <form
        id="needs-validation"
        ref="form"
        data-focus="false"
        method="post"
        :action="ApiPath('/api/mail')"
        role="form"
        novalidate
        class="flex flex-col gap-3 [&.validate_.input:has(input:invalid)]:input-error [&.validate_textarea:invalid]:input-error"
        :class="{ validate: !formValid }"
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
              v-model="name"
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
              v-model="email"
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
            ref="textarea"
            v-model="message"
            class="textarea textarea-bordered max-h-96 w-full resize-none"
            name="text"
            required
            placeholder="Message"
          />
        </label>

        <div v-if="TurnstileEnabled" ref="captcha" class="self-center h-[65px]" />

        <!-- Button -->
        <button class="btn btn-primary self-end">
          <loading-icon v-if="loading" />
          <send-icon v-else />
          Send
        </button>
      </form>
    </portfolio-card>
  </section>
</template>

<script setup>
import { useTextareaAutosize } from "@vueuse/core";
import { onActivated, onDeactivated, ref } from "vue";
import AtIcon from "~icons/material-symbols/alternate-email-rounded";
import CheckIcon from "~icons/material-symbols/check-circle-outline-rounded";
import ErrorIcon from "~icons/material-symbols/error-outline-rounded";
import AccountIcon from "~icons/material-symbols/person-rounded";
import SendIcon from "~icons/material-symbols/send-rounded";
import GithubIcon from "~icons/simple-icons/github";
import LinkedinIcon from "~icons/simple-icons/linkedin";
import LoadingIcon from "~icons/svg-spinners/ring-resize";
import PortfolioCard from "@/components/PortfolioCard.vue";
import { ApiPath } from "@/config/api";
import pb from "@/plugins/pocketbase";
import { TurnstileEnabled, TurnstileKey, TurnstileReady, useTurnstile } from "@/plugins/turnstile";

const name = ref("");
const email = ref("");

const error = ref(null);
const success = ref(false);
const loading = ref(false);

const form = ref(null);
const formValid = ref(true);
let captchaValue;

const { textarea, input: message } = useTextareaAutosize();

const submit = async () => {
  formValid.value = form.value.checkValidity();
  if (formValid.value) {
    loading.value = true;
    success.value = false;
    error.value = null;

    const params = {
      name: name.value,
      email: email.value,
      message: message.value,
    };
    const wait = new Promise((resolve) => setTimeout(resolve, 1000));

    try {
      await Promise.all([
        pb.collection("contact_form").create(params, {
          headers: { "X-Captcha": captchaValue },
        }),
        wait,
      ]);

      name.value = "";
      email.value = "";
      message.value = "";
      success.value = true;
    } catch (e) {
      console.error(e);
      error.value = e;
    } finally {
      loading.value = false;
    }
  } else {
    error.value = Error("Please fix invalid fields");
  }
};

const captcha = ref(null);
if (TurnstileEnabled) {
  useTurnstile();

  onActivated(async () => {
    await TurnstileReady;
    window.turnstile.render(captcha.value, {
      sitekey: TurnstileKey,
      action: "contact",
      callback: (token) => (captchaValue = token),
    });
  });

  onDeactivated(() => {
    window.turnstile.remove();
  });
}
</script>
