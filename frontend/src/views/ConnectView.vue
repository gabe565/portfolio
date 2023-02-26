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
                <font-awesome-icon icon="fab fa-twitter" fixed-width />
                <span class="network-name">Twitter</span>
              </a>
            </li>
            <li class="list-inline-item">
              <a
                href="/to/github"
                class="btn btn-outline-primary btn-lg"
                target="_blank"
              >
                <font-awesome-icon icon="fab fa-github" fixed-width />
                <span class="network-name">Github</span>
              </a>
            </li>
            <li class="list-inline-item">
              <a
                href="/to/linkedin"
                class="btn btn-outline-primary btn-lg"
                target="_blank"
              >
                <font-awesome-icon icon="fab fa-linkedin" fixed-width />
                <span class="network-name">LinkedIn</span>
              </a>
            </li>
          </ul>
        </div>
      </div>
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
            @submit.prevent="mail"
          >
            <fieldset>
              <div class="col mb-3">
                <div
                  v-if="message"
                  class="alert"
                  :class="[success ? 'alert-success' : 'alert-danger']"
                >
                  {{ message }}
                </div>
              </div>
              <!--Name-->
              <div class="form-group mb-3">
                <label for="nameInput" class="form-label">Name</label>
                <div class="input-group">
                  <div class="input-group-text">
                    <font-awesome-icon icon="far fa-user-alt" fixed-width />
                  </div>
                  <input
                    id="nameInput"
                    v-model="name"
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
                <label for="emailInput" class="form-label">Email Address</label>
                <div class="input-group">
                  <div class="input-group-text">
                    <font-awesome-icon icon="far fa-at" fixed-width />
                  </div>
                  <input
                    id="emailInput"
                    v-model="email"
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
                    <font-awesome-icon icon="far fa-comment" fixed-width />
                  </div>
                  <textarea
                    id="messageInput"
                    v-model="text"
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
                    icon="far fa-sync"
                    fixed-width
                    spin
                  />
                  <font-awesome-icon
                    v-else
                    icon="far fa-paper-plane"
                    fixed-width
                  />
                  Send
                </button>
              </div>
            </fieldset>
          </form>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
import pb from "@/plugins/pocketbase";

export default {
  data: () => ({
    name: "",
    email: "",
    text: "",
    error: null,
    success: false,
    loading: false,
  }),
  computed: {
    message() {
      if (this.error) {
        return "There was an error during submission.";
      } else if (this.success) {
        return "The message has been sent successfully. Thank you!";
      } else {
        return null;
      }
    },
  },
  methods: {
    async mail() {
      const valid = this.$refs.form.checkValidity();
      this.$refs.form.classList.add("was-validated");
      if (valid) {
        this.loading = true;
        const wait = new Promise((resolve) => setTimeout(resolve, 1000));
        try {
          await pb.collection("contact_form").create({
            name: this.name,
            email: this.email,
            message: this.text,
          });
          await wait;
          this.success = true;
        } catch (error) {
          console.error(error);
          this.error = error;
        } finally {
          this.loading = false;
        }
      }
    },
  },
};
</script>
