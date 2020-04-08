# Composer
FROM gabe565/prestissimo as vendor

ENV COMPOSER_ALLOW_SUPERUSER 1

COPY composer.json composer.lock ./
RUN set -x \
    && composer install \
        --ignore-platform-reqs \
        --no-autoloader \
        --no-interaction \
        --no-progress \
        --no-suggest

COPY . ./
RUN set -x \
    && composer dump-autoload \
        --classmap-authoritative \
        --no-interaction \
    && php artisan route:cache

# Node
FROM node:8-alpine as frontend

ARG NODE_ENV=production

WORKDIR /app

COPY artisan package.json webpack.mix.js yarn.lock ./

RUN set -x \
    && yarn install \
        --frozen-lockfile \
        --no-cache \
        --no-progress

COPY public/ public/
COPY resources/ resources/

RUN set -x \
    && node_modules/webpack/bin/webpack.js \
        --no-progress \
        --hide-modules \
        --config=node_modules/laravel-mix/setup/webpack.config.js

# Final Image
FROM php:7.4-fpm-alpine
LABEL maintainer="Gabe Cook <gabe565@gmail.com>"

RUN set -x \
    && apk add --no-cache \
        fcgi \
    && docker-php-ext-install -j"$(nproc)" mysqli pdo_mysql opcache \
    && rm -rf /tmp/* /var/cache/apk/* \
    && sed -ri \
        -e 's/;(ping\.path)/\1/' /usr/local/etc/php-fpm.d/www.conf \
    && cd "$PHP_INI_DIR" \
    && ln -s php.ini-production php.ini \
    && sed -ri \
        -e 's/^(expose_php).*$/\1 = Off/' \
        php.ini-production

WORKDIR /var/www/html
VOLUME /var/www/html

COPY --chown=82:82 ./ .
COPY --from=vendor --chown=82:82 /app/vendor vendor/
COPY --from=frontend --chown=82:82 /app/public public/
COPY --chown=root docker/entrypoint docker/health-check /

HEALTHCHECK --interval=1m --timeout=5s \
    CMD /health-check

CMD ["/entrypoint"]
