FROM composer:2 as local-composer
FROM ghcr.io/roadrunner-server/roadrunner:2.12.2 AS roadrunner
FROM mlocati/php-extension-installer:2.0.2 as php-extension-installer

FROM php:8.2-cli-alpine as base-image
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
FROM node:18-alpine as node-builder
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
RUN --mount=type=bind,from=php-extension-installer,source=/usr/bin/install-php-extensions,target=/usr/local/bin/install-php-extensions \
    install-php-extensions \
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
