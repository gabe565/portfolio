ARG PHP_VERSION=8.1
ARG COMPOSER_VERSION=2
ARG NODE_VERSION=16

ARG INSTALL_BCMATH=false
ARG INSTALL_EXIF=true
ARG INSTALL_PGSQL=true

# composer install
FROM ghcr.io/clevyr/php:$PHP_VERSION-composer$COMPOSER_VERSION-base as php-builder
WORKDIR /app

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
FROM ghcr.io/clevyr/php:$PHP_VERSION-composer$COMPOSER_VERSION as local-image
WORKDIR /app

COPY --chown=root docker/app/rootfs /
RUN crontab /etc/cron.d/scheduler

CMD ["s6-svscan", "/etc/services.d/app"]


# prod image
FROM local-image as prod-image
LABEL org.opencontainers.image.source="https://github.com/gabe565/portfolio"
COPY --from=php-builder --chown=82:82 /app .
COPY --from=node-builder --chown=82:82 /app/public public/
