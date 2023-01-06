ARG PHP_VERSION=8.1
ARG ROADRUNNER_VERSION=2.12.1
ARG COMPOSER_VERSION=2
ARG NODE_VERSION=18

ARG INSTALL_BCMATH=false
ARG INSTALL_EXIF=true
ARG INSTALL_PGSQL=true

ARG DEPS=tzdata

FROM composer:$COMPOSER_VERSION as local-composer
FROM ghcr.io/roadrunner-server/roadrunner:$ROADRUNNER_VERSION AS roadrunner
FROM php:$PHP_VERSION-cli-alpine as base-image
WORKDIR /app
COPY --from=local-composer /usr/bin/composer /usr/bin/composer
COPY --from=roadrunner /usr/bin/rr /usr/local/bin/rr

# composer install
FROM base-image as php-builder

COPY composer.json composer.lock ./
RUN composer install \
        --ignore-platform-reqs \
        --no-autoloader \
        --no-interaction \
        --no-progress \
        --no-suggest

COPY . .
RUN composer dump-autoload \
        --classmap-authoritative \
        --no-interaction


# npm install
FROM node:$NODE_VERSION-alpine as node-builder
WORKDIR /app

COPY artisan package.json package-lock.json webpack.* .npmrc ./
ARG FONTAWESOME_NPM_AUTH_TOKEN
RUN npm ci

COPY public/ public/
COPY resources/ resources/

ARG MIX_GOOGLE_API_KEY
RUN npm run production


# local image
FROM base-image as local-image
RUN curl -sSL https://github.com/mlocati/docker-php-extension-installer/releases/latest/download/install-php-extensions -o - | sh -s \
      exif \
      opcache \
      pdo_pgsql \
      pgsql \
      sockets
COPY --chown=root docker/entrypoint /
CMD ["/entrypoint"]


# prod image
FROM local-image as prod-image
LABEL org.opencontainers.image.source="https://github.com/gabe565/portfolio"
COPY --from=php-builder --chown=82:82 /app .
COPY --from=node-builder --chown=82:82 /app/public public/
