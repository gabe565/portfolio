<template>
    <section id="about" class="content-section text-center">
        <div class="container">
            <h2>About</h2>
            <form id="email-form" data-toggle="validator" data-focus="false" method="post"
                                                                             action="email" role="form">
                <fieldset>
                    <div class="row">
                        <div class="col-md-8 mx-sm-auto">
                            <div class="messages"></div>
                        </div>
                    </div>
                    <!--Name-->
                    <div class="row">
                        <div class="form-group">
                            <div class="col-md-6 mx-sm-auto inputGroupContainer">
                                <div class="input-group"> <span class="input-group-addon"><i class="fa fa-user fa-fw" aria-hidden="true"></i></span>
                                    <input
                                                                                                                      name="name" placeholder="Name" class="form-control" type="text" required>
                                </div>
                            </div>
                        </div>
                    </div>
                    <br>
                    <!--Email-->
                    <div class="row">
                        <div class="form-group">
                            <div class="col-md-6 mx-sm-auto inputGroupContainer">
                                <div class="input-group"> <span class="input-group-addon"><i class="fa fa-envelope fa-fw" aria-hidden="true"></i></span>
                                    <input
                                                                                                                          name="email" placeholder="Email Address" class="form-control" type="email"
                                                                                                                                                                                        required>
                                </div>
                            </div>
                        </div>
                    </div>
                    <br>
                    <!--Message-->
                    <div class="row">
                        <div class="form-group">
                            <div class="col-md-6 mx-sm-auto inputGroupContainer">
                                <div class="input-group"> <span class="input-group-addon"><i class="fa fa-comment fa-fw" aria-hidden="true"></i></span>
                                    <textarea
                                                                                                                         class="form-control vertical" name="message" placeholder="Message" required></textarea>
                                </div>
                            </div>
                        </div>
                    </div>
                    <br>
                    <!-- Button -->
                    <div class="row">
                        <div class="form-group">
                            <div class="col-md-6 mx-sm-auto">
                                <button type="submit" class="btn btn-primary btn-raised"><i id="send-icon" class="fa fa-paper-plane fa-fw" aria-hidden="true"></i>&#xA0;&#xA0;Send</button>
                            </div>
                        </div>
                    </div>
                </fieldset>
            </form>
        </div>
    </section>
</template>
