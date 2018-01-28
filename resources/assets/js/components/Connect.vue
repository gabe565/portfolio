<template>
    <section id="about" class="content-section text-center">
        <div class="container">
            <h1>Connect</h1>
            <p>
                Here are links to my profiles on other relevant sites.<br>
                The email form below will directly send to my personal email address. Feel free to contact me!
            </p>
            <hr>
            <h3>Accounts</h3>
            <div class="row">
                <div class="col-lg-8 mx-auto">
                    <ul class="list-inline banner-social-buttons">
                        <li class="list-inline-item">
                            <a href="/to/twitter" class="btn btn-default btn-lg" target="_blank">
                                <svgicon name="twitter"></svgicon>
                                <span class="network-name">Twitter</span>
                            </a>
                        </li>
                        <li class="list-inline-item">
                            <a href="/to/github" class="btn btn-default btn-lg" target="_blank">
                                <svgicon name="github"></svgicon>
                                <span class="network-name">Github</span>
                            </a>
                        </li>
                        <li class="list-inline-item">
                            <a href="/to/linkedin" class="btn btn-default btn-lg" target="_blank">
                                <svgicon name="linkedin"></svgicon>
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
                                <div class="input-group-prepend">
                                    <div class="input-group-text">
                                        <svgicon name="user-alt"></svgicon>
                                    </div>
                                </div>
                                <input v-model="name" name="name" placeholder="Name" class="form-control" type="text" required>
                            </div>
                        </div>
                        <!--Email-->
                        <div class="form-group">
                            <div class="input-group">
                                <div class="input-group-prepend">
                                    <div class="input-group-text">
                                        <svgicon name="at"></svgicon>
                                    </div>
                                </div>
                                <input v-model="email" name="email" placeholder="Email Address" class="form-control" type="email" required>
                            </div>
                        </div>
                        <!--Message-->
                        <div class="form-group">
                            <div class="input-group">
                                <div class="input-group-prepend">
                                    <div class="input-group-text">
                                        <svgicon name="comment"></svgicon>
                                    </div>
                                </div>
                                <textarea v-model="text" class="form-control vertical" name="text" placeholder="Message" required></textarea>
                            </div>
                        </div>
                        <!-- Button -->
                        <div class="form-group">
                            <button class="btn btn-default">
                                <svgicon name="paper-plane" v-if="!loading"></svgicon>
                                <svgicon name="sync" class="svg-spin" v-if="loading"></svgicon>
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
    data() {
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
        message() {
            if (this.successful)
                return 'The message has been sent successfully. Thank you!'
            else
                return 'There was an error during submission'
        },
        successful() {
            return this.response != null && this.response.status == 200
        }
    },
    methods: {
        sendmail() {
            let valid = this.form.checkValidity()
            this.form.classList.add('was-validated')
            if (valid) {
                this.loading = true
                axios.post('/api/mail', {
                    name: this.name,
                    email: this.email,
                    text: this.text
                }).then(response => {
                    this.response = response
                    this.loading = false
                    this.showResponse = true
                }).catch(error => {
                    this.respose = error
                    this.loading = false
                    this.showResponse = true
                })
            }
        }
    },
    mounted() {
        this.form = document.getElementById("needs-validation")
    }
}
</script>
