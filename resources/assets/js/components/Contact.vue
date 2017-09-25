<template>
    <section id="about" class="content-section text-center">
        <div class="container">
            <h1>Contact</h1>
            <div class="col-lg-6 col-md-8 mx-auto">
                <form id="needs-validation" data-focus="false" method="post" action="email" role="form" novalidate>
                    <fieldset>
                        <div class="col-md-8 mx-sm-auto">
                            <div class="messages"></div>
                        </div>
                        <!--Name-->
                        <div class="form-group">
                            <div class="input-group">
                                <span class="input-group-addon">
                                    <i class="far fa-user fa-fw" aria-hidden="true"></i>
                                </span>
                                <input name="name" placeholder="Name" class="form-control" type="text" required>
                            </div>
                        </div>
                        <!--Email-->
                        <div class="form-group">
                            <div class="input-group">
                                <span class="input-group-addon">
                                    <i class="far fa-envelope fa-fw" aria-hidden="true"></i>
                                </span>
                                <input name="email" placeholder="Email Address" class="form-control" type="email" required>
                            </div>
                        </div>
                        <!--Message-->
                        <div class="form-group">
                            <div class="input-group">
                                <span class="input-group-addon">
                                    <i class="far fa-comment fa-fw" aria-hidden="true"></i>
                                </span>
                                <textarea class="form-control vertical" name="message" placeholder="Message" required></textarea>
                            </div>
                        </div>
                        <!-- Button -->
                        <div class="form-group">
                            <button type="submit" class="btn btn-primary btn-raised">
                                <i id="send-icon" class="far fa-paper-plane fa-fw" aria-hidden="true"></i>
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
    mounted: function() {
        var form = document.getElementById("needs-validation")
        form.addEventListener("submit", function(e) {
            if (form.checkValidity() == false) {
                e.preventDefault()
                e.stopPropagation()
            }
            form.classList.add('was-validated');
        })
    }
}
</script>
