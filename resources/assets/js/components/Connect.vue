<template>
    <section id="about" class="content-section text-center">
        <div class="container">
            <h1>Connect</h1>
            <hr>
            <h3>Accounts</h3>
            <div class="row">
                <div class="col-lg-8 mx-auto">
                    <ul class="list-inline banner-social-buttons">
                        <li class="list-inline-item">
                            <a href="/to/twitter" class="btn btn-default btn-lg">
                                <i class="fab fa-twitter fa-fw"></i>
                                <span class="network-name">Twitter</span>
                            </a>
                        </li>
                        <li class="list-inline-item">
                            <a href="/to/github" class="btn btn-default btn-lg">
                                <i class="fab fa-github fa-fw"></i>
                                <span class="network-name">Github</span>
                            </a>
                        </li>
                        <li class="list-inline-item">
                            <a href="/to/linkedin" class="btn btn-default btn-lg">
                                <i class="fab fa-linkedin fa-fw"></i>
                                <span class="network-name">LinkedIn</span>
                            </a>
                        </li>
                    </ul>
                </div>
            </div>
            <br>
            <hr>
            <h3>Email</h3>
            <div class="col-lg-6 col-md-8 mx-auto">
                <form id="needs-validation" data-focus="false" method="post" action="/api/mail" role="form" novalidate v-on:submit.prevent="sendmail">
                    <fieldset>
                        <div class="col">
                            <div class="alert" v-if="showResponse" :class="[ successful ? 'alert-success' : 'alert-danger']">{{ message }}</div>
                        </div>
                        <!--Name-->
                        <div class="form-group">
                            <div class="input-group">
                                <span class="input-group-addon">
                                    <i class="far fa-user fa-fw" aria-hidden="true"></i>
                                </span>
                                <input v-model="name" name="name" placeholder="Name" class="form-control" type="text" required>
                            </div>
                        </div>
                        <!--Email-->
                        <div class="form-group">
                            <div class="input-group">
                                <span class="input-group-addon">
                                    <i class="far fa-at fa-fw" aria-hidden="true"></i>
                                </span>
                                <input v-model="email" name="email" placeholder="Email Address" class="form-control" type="email" required>
                            </div>
                        </div>
                        <!--Message-->
                        <div class="form-group">
                            <div class="input-group">
                                <span class="input-group-addon">
                                    <i class="far fa-comment fa-fw" aria-hidden="true"></i>
                                </span>
                                <textarea v-model="text" class="form-control vertical" name="text" placeholder="Message" required></textarea>
                            </div>
                        </div>
                        <!-- Button -->
                        <div class="form-group">
                            <button class="btn btn-default">
                                <i class="far fa-fw" :class="[ loading ? 'fa-sync fa-spin' : 'fa-paper-plane' ]" aria-hidden="true"></i>
                                &nbsp;Send
                            </button>
                        </div>
                    </fieldset>
                </form>
            </div>
        </div>
    </section>
</template>

<script>
export default {
    data: function() {
        return {
            form: null,
            name: '',
            email: '',
            text: '',
            loading: false,
            response: null,
            showResponse: false,
        }
    },
    computed: {
        message: function() {
            if (this.response != null && this.response.status == 200)
                return 'The message has been sent successfully. Thank you!'
            else
                return 'There was an error during submission'
        },
        successful: function() {
            if (this.response != null && this.response.status == 200)
                return true
            else
                return false
        }
    },
    methods: {
        sendmail: function() {
            var valid = this.form.checkValidity()
            this.form.classList.add('was-validated')
            if (valid) {
                this.loading = true
                var vue = this
                axios.post('/api/mail', {
                    name: this.name,
                    email: this.email,
                    text: this.text
                }).then(function(response) {
                    vue.response = response
                    vue.loading = false
                    vue.showResponse = true
                }).catch(function(error) {
                    vue.respose = error
                    vue.loading = false
                    vue.showResponse = true
                })
            }
        }
    },
    mounted: function() {
        this.form = document.getElementById("needs-validation")
    }
}
</script>
