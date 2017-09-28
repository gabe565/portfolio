<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="utf-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
        <meta name="description" content="">
        <meta name="author" content="">
        <meta name="csrf-token" content="{{ csrf_token() }}">

        <title>{{ config('app.name') }}</title>

        <!-- Bootstrap core CSS -->
        <link href="{{ mix('css/bootstrap.css') }}" rel="stylesheet">
        <link href="{{ mix('css/font-awesome.css') }}" rel="stylesheet">
        <link href="{{ mix('css/main.css') }}" rel="stylesheet">
    </head>
    <body>
        <div id="app">
            <nav class="navbar navbar-expand-lg fixed-top navbar-dark">
                <div class="container">
                    <router-link to="/" class="navbar-brand">{{ config('app.name') }}</router-link>
                    <button class="navbar-toggler navbar-toggler-right" type="button" data-toggle="collapse" data-target="#navbarResponsive" aria-controls="navbarResponsive" aria-expanded="false" aria-label="Toggle navigation">
                        <i class="far fa-bars"></i>
                        &nbsp;Menu
                    </button>
                    <div class="collapse navbar-collapse" id="navbarResponsive">
                        <ul class="navbar-nav ml-auto">
                            <li class="nav-item">
                                <router-link to="/about" class="nav-link">
                                    <i class="far fa-info-circle fa-fw"></i>
                                    About
                                </router-link>
                            </li>
                            <li class="nav-item">
                                <router-link to="/skills" class="nav-link">
                                    <i class="far fa-list fa-fw"></i>
                                    Skills
                                </router-link>
                            </li>
                            <li class="nav-item">
                                <router-link to="/projects" class="nav-link">
                                    <i class="far fa-file fa-fw"></i>
                                    Projects
                                </router-link>
                            </li>
                            <li class="nav-item">
                                <router-link to="/connect" class="nav-link">
                                    <i class="far fa-envelope fa-fw"></i>
                                    Connect
                                </router-link>
                            </li>
                        </ul>
                    </div>
                </div>
            </nav>

            <transition name="fade" mode="out-in" appear>
            <router-view></router-view>
            </transition>

            <footer>
                <div class="container">
                    <a href="https://github.com/gabe565/gabecook.com/blob/master/LICENSE" target="_blank" class="float-left">
                        <span>&copy; 2017 Gabe Cook</span>
                    </a>
                    <a href="https://github.com/gabe565/gabecook.com" target="_blank" class="float-right">
                        <span>
                            <i class="fab fa-github fa-fw"></i>&nbsp;View on GitHub
                        </span>
                    </a>
                </div>
            </footer>

        </div>

        <!-- Google Maps API Key - Use your own API key to enable the map feature. More information on the Google Maps API can be found at https://developers.google.com/maps/ -->
        <script src="https://maps.googleapis.com/maps/api/js?key=AIzaSyCRngKslUGJTlibkQ3FkfTxj3Xss1UlZDA&sensor=false"></script>

        <script src="{{ mix('js/manifest.js') }}"></script>
        <script src="{{ mix('js/vendor.js') }}"></script>
        <script src="{{ mix('js/app.js') }}"></script>
    </body>
</html>
