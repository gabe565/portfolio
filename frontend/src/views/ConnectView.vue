<template>
  <section class="pt-page">
    <div class="container pb-5">
      <div class="row mb-4">
        <div class="col">
          <h1>Connect</h1>
        </div>
      </div>
      <div class="row">
        <div class="col">
          Here are some ways to reach out or see what I'm up to!
        </div>
      </div>
      <section>
        <div class="row mt-4 mb-3">
          <div class="col">
            <h2 class="h3">Accounts</h2>
          </div>
        </div>
        <div class="row">
          <div class="col-lg-8 mx-auto">
            <ul class="list-inline banner-social-buttons">
              <li class="list-inline-item">
                <a
                  href="/to/github"
                  class="btn btn-outline-primary btn-lg"
                  target="_blank"
                >
                  <GitHubIcon class="icon-inline icon-fw me-2" />
                  <span class="network-name">Github</span>
                </a>
              </li>
              <li class="list-inline-item">
                <a
                  href="/to/linkedin"
                  class="btn btn-outline-primary btn-lg"
                  target="_blank"
                >
                  <LinkedinIcon class="icon-inline icon-fw me-2" />
                  <span class="network-name">LinkedIn</span>
                </a>
              </li>
            </ul>
          </div>
        </div>
      </section>
      <section>
        <div class="row mt-4 mb-1">
          <div class="col">
            <h2 class="h3">Contact</h2>
          </div>
        </div>
        <div class="row">
          <div class="col-lg-6 col-md-8 mx-auto">
            <form
              id="needs-validation"
              ref="form"
              data-focus="false"
              method="post"
              action="/api/mail"
              role="form"
              novalidate
              @submit.prevent="submit"
            >
              <fieldset>
                <div class="col mb-3">
                  <div v-if="error" class="alert alert-danger">
                    Failed to send message. Please try again later.
                  </div>
                  <div v-else-if="success" class="alert alert-success">
                    Your message has been sent successfully. Thank you!
                  </div>
                </div>
                <!--Name-->
                <div class="form-group mb-3">
                  <label for="nameInput" class="form-label">Name</label>
                  <div class="input-group">
                    <div class="input-group-text">
                      <AccountIcon class="icon-inline icon-fw" />
                    </div>
                    <input
                      id="nameInput"
                      v-model="formData.name"
                      name="name"
                      placeholder="Jane Doe"
                      class="form-control"
                      type="text"
                      required
                    />
                  </div>
                </div>
                <!--Email-->
                <div class="form-group mb-3">
                  <label for="emailInput" class="form-label"
                    >Email Address</label
                  >
                  <div class="input-group">
                    <div class="input-group-text">
                      <AtIcon class="icon-inline icon-fw" />
                    </div>
                    <input
                      id="emailInput"
                      v-model="formData.email"
                      name="email"
                      placeholder="name@example.com"
                      class="form-control"
                      type="email"
                      required
                    />
                  </div>
                </div>
                <!--Message-->
                <div class="form-group mb-3">
                  <label for="messageInput" class="form-label">Message</label>
                  <div class="input-group">
                    <div class="input-group-text">
                      <CommentIcon class="icon-inline icon-fw" />
                    </div>
                    <textarea
                      id="messageInput"
                      v-model="formData.message"
                      class="form-control vertical"
                      name="text"
                      required
                      style="min-height: 62px"
                      @input="
                        $event.target.style.height = '';
                        $event.target.style.height = `${$event.target.scrollHeight}px`;
                      "
                    ></textarea>
                  </div>
                </div>
                <!-- Button -->
                <div class="form-group mb-3">
                  <button class="btn btn-outline-primary">
                    <RefreshIcon
                      v-if="loading"
                      class="icon-inline icon-fw icon-spin"
                    />
                    <SendIcon v-else class="icon-inline icon-fw" />
                    Send
                  </button>
                </div>
              </fieldset>
            </form>
          </div>
        </div>
      </section>
    </div>
  </section>
</template>

<script setup>
import pb from "@/plugins/pocketbase";
import LinkedinIcon from "feather-icons/dist/icons/linkedin.svg";
import GitHubIcon from "feather-icons/dist/icons/github.svg";
import AccountIcon from "feather-icons/dist/icons/user.svg";
import AtIcon from "feather-icons/dist/icons/at-sign.svg";
import CommentIcon from "feather-icons/dist/icons/message-circle.svg";
import SendIcon from "feather-icons/dist/icons/send.svg";
import RefreshIcon from "feather-icons/dist/icons/refresh-cw.svg";

const formData = ref({
  name: "",
  email: "",
  message: "",
});

const error = ref(null);
const success = ref(false);
const loading = ref(false);

const form = ref(null);

const submit = async () => {
  const valid = form.value.checkValidity();
  form.value.classList.add("was-validated");
  if (valid) {
    loading.value = true;
    const wait = new Promise((resolve) => setTimeout(resolve, 1000));
    try {
      await pb.collection("contact_form").create(formData.value);
      await wait;
      success.value = true;
    } catch (e) {
      console.error(e);
      error.value = e;
    } finally {
      loading.value = false;
    }
  }
};
</script>
