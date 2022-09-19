<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="utf-8">
        <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
        <meta name="description" content="Personal portfolio website for {{ config('app.name') }}">
        <meta name="author" content="{{ config('app.name') }}">
        <meta name="csrf-token" content="{{ csrf_token() }}">

        <link rel="apple-touch-icon" sizes="180x180" href="/apple-touch-icon.png">
        <link rel="icon" type="image/png" sizes="32x32" href="/favicon-32x32.png">
        <link rel="icon" type="image/png" sizes="16x16" href="/favicon-16x16.png">
        <link rel="manifest" href="/manifest.json">
        <link rel="mask-icon" href="/safari-pinned-tab.svg" color="#607d8b">
        <meta name="theme-color" content="#607d8b">

        <meta property='og:title' content='{{ config('app.name') }}'/>
        <meta property='og:image' content='{{ config('app.url') }}/images/thumb.png'/>
        <meta property='og:description' content='Personal portfolio website for {{ config('app.name') }}'/>
        <meta property='og:url' content='{{ config('app.url') }}' />
        <meta property="og:type" content="website" />

        <title>{{ config('app.name') }}</title>

        <!-- Bootstrap core CSS -->
        <link href="{{ mix('css/bootstrap.css') }}" rel="stylesheet">
        <link href="{{ mix('css/main.css') }}" rel="stylesheet">

        <!-- Global site tag (gtag.js) - Google Analytics -->
        <script async src="https://www.googletagmanager.com/gtag/js?id=UA-81536605-1"></script>
        <script>
            window.dataLayer = window.dataLayer || [];
            function gtag(){dataLayer.push(arguments);}
            gtag('js', new Date());

            gtag('config', 'UA-81536605-1');
        </script>
    </head>
    <body>
        <div id="app" :class="$route.meta.title">
            <nav class="navbar navbar-expand-md fixed-top navbar-dark">
                <div class="container">
                    <router-link to="/" class="navbar-brand">&lt; {{ str_replace(' ', '.', strtolower(config('app.name'))) }} &gt;</router-link>
                    <button class="navbar-toggler navbar-toggler-right" type="button" data-toggle="collapse" data-target="#navbarResponsive" aria-controls="navbarResponsive" aria-expanded="false" aria-label="Toggle navigation">
                        <font-awesome-icon icon="fas fa-bars"></font-awesome-icon>
                        &nbsp;Menu
                    </button>
                    <div class="collapse navbar-collapse" id="navbarResponsive">
                        <ul class="navbar-nav ms-auto">
                            <li class="nav-item">
                                <router-link to="/about" class="nav-link">
                                    <font-awesome-icon icon="fas fa-info-circle" fixed-width></font-awesome-icon>
                                    About
                                </router-link>
                            </li>
                            <li class="nav-item">
                                <router-link to="/skills" class="nav-link">
                                    <font-awesome-icon icon="fas fa-list-ul" fixed-width></font-awesome-icon>
                                    Skills
                                </router-link>
                            </li>
                            <li class="nav-item">
                                <router-link to="/projects" class="nav-link">
                                    <font-awesome-icon icon="fas fa-code" fixed-width></font-awesome-icon>
                                    Projects
                                </router-link>
                            </li>
                            <li class="nav-item">
                                <router-link to="/connect" class="nav-link">
                                    <font-awesome-icon icon="fas fa-envelope" fixed-width></font-awesome-icon>
                                    Connect
                                </router-link>
                            </li>
                        </ul>
                    </div>
                </div>
            </nav>

            <transition :name="transitionName" mode="out-in" appear>
                <keep-alive>
                    <router-view class="child-view"></router-view>
                </keep-alive>
            </transition>

            <vue-progress-bar></vue-progress-bar>

            <footer>
                <div class="container">
                    <div class="row text-center">
                        <div class="col-sm text-sm-start">
                            <span>&lt;/ {{ str_replace(' ', '.', strtolower(config('app.name'))) }} &gt;</span>
                        </div>
                        <div class="col-sm">
                            <a href="//github.com/gabe565/gabecook.com" target="_blank">
                                <span>
                                    <font-awesome-icon icon="fab fa-github"></font-awesome-icon>
                                    &nbsp;View on GitHub
                                </span>
                            </a>
                        </div>
                        <div class="col-sm text-sm-end">
                            <a href="//github.com/gabe565/gabecook.com/blob/master/LICENSE" target="_blank">
                                &copy; {{ now()->year }} {{ config('app.name') }}
                            </a>
                        </div>
                    </div>
                </div>
            </footer>

        </div>

        <script src="//cdn.polyfill.io/v2/polyfill.js?unknown=polyfill&features=Object.assign|gated,Promise|gated"></script>
        <script src="{{ mix('js/manifest.js') }}"></script>
        <script src="{{ mix('js/vendor.js') }}"></script>
        <script src="{{ mix('js/app.js') }}"></script>
    </body>
</html>
