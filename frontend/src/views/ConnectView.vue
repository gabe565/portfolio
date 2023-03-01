<template>
  <section>
    <div class="container">
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
                  href="/to/twitter"
                  class="btn btn-outline-primary btn-lg"
                  target="_blank"
                >
                  <font-awesome-icon :icon="faTwitter" fixed-width />
                  <span class="network-name">Twitter</span>
                </a>
              </li>
              <li class="list-inline-item">
                <a
                  href="/to/github"
                  class="btn btn-outline-primary btn-lg"
                  target="_blank"
                >
                  <font-awesome-icon :icon="faGithub" fixed-width />
                  <span class="network-name">Github</span>
                </a>
              </li>
              <li class="list-inline-item">
                <a
                  href="/to/linkedin"
                  class="btn btn-outline-primary btn-lg"
                  target="_blank"
                >
                  <font-awesome-icon :icon="faLinkedin" fixed-width />
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
                    There was an error during submission.
                  </div>
                  <div v-else-if="success" class="alert alert-success">
                    The message has been sent successfully. Thank you!
                  </div>
                </div>
                <!--Name-->
                <div class="form-group mb-3">
                  <label for="nameInput" class="form-label">Name</label>
                  <div class="input-group">
                    <div class="input-group-text">
                      <font-awesome-icon :icon="faUserAlt" fixed-width />
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
                      <font-awesome-icon :icon="faAt" fixed-width />
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
                      <font-awesome-icon :icon="faComment" fixed-width />
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
                    <font-awesome-icon
                      v-if="loading"
                      :icon="faSync"
                      fixed-width
                      spin
                    />
                    <font-awesome-icon
                      v-else
                      :icon="faPaperPlane"
                      fixed-width
                    />
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
import {
  faTwitter,
  faGithub,
  faLinkedin,
} from "@fortawesome/free-brands-svg-icons";
import {
  faAt,
  faComment,
  faSync,
  faPaperPlane,
  faUserAlt,
} from "@fortawesome/pro-regular-svg-icons";

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
