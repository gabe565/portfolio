let mix = require('laravel-mix');

/*
 |--------------------------------------------------------------------------
 | Mix Asset Management
 |--------------------------------------------------------------------------
 |
 | Mix provides a clean, fluent API for defining some Webpack build steps
 | for your Laravel application. By default, we are compiling the Sass
 | file for the application as well as bundling up all the JS files.
 |
 */

mix.autoload({
    jquery: ['$', 'jQuery', 'window.jQuery'],
    'popper.js/dist/umd/popper.js': ['Popper']
})
    .js('resources/js/app.js', 'public/js')
    .vue()
    .extract()
    .sass('resources/sass/bootstrap.scss', 'public/css')
    .sass('resources/sass/main.scss', 'public/css')
    .version()
    .sourceMaps();
