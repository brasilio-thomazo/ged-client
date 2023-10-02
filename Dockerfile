ARG EXT="pdo pdo_sqlite pdo_pgsql pdo_odbc pdo_mysql pdo_dblib"
ARG PECL_EXT="redis protobuf mongodb memcache"
ARG USER_UID=1000
ARG USER_GID=1000

#
# NODEJS
#
FROM node:alpine as node
RUN mkdir -p /home/app/public_html
WORKDIR /home/app/public_html
COPY . .
RUN npm install && npm run build && rm -r node_modules

FROM devoptimus/php-fpm as fpm
ARG EXT
ARG PECL_EXT
ARG USER_UID
ARG USER_GID
ENV UID=${USER_GID}
ENV GID=${USER_GID}

USER root
RUN install-php-ext ${EXT} \
    && install-php-pecl-ext ${PECL_EXT}
COPY resources/docker/www.ini /etc/php/fpm/pool.d/
COPY resources/docker/fpm-entrypoint /usr/local/bin/entrypoint


USER app
WORKDIR /home/app/public_html
COPY --chown=app composer.json ./
RUN composer install --no-dev --no-scripts --no-autoloader --prefer-dist
COPY --chown=app . .
RUN composer install --no-dev --prefer-dist
COPY --from=node /home/app/public_html/public /home/app/public_html/public

RUN php artisan event:cache \
    && php artisan route:cache \
    && php artisan view:cache


#
# NGINX
#
FROM devoptimus/nginx as nginx
COPY resources/docker/default.template /etc/nginx/template.d/
COPY --from=node /home/app/public_html/public /home/app/public_html/public
WORKDIR /home/app/public_html/public

#
# Build server protobufer
#
FROM alpine as build
COPY resources/golang /golang
WORKDIR /golang
RUN apk add --no-cache go \
    && go mod tidy \
    && go build -v -o duat-server .

#
# Server protobufer
#
FROM alpine as server
ENV UID=1000
ENV GID=1000
ENV SERVER_PORT=50051

COPY --from=build /golang/duat-server /usr/local/bin/

RUN apk add --no-cache \
    && addgroup app \
    && adduser -G app -D app \
    &&  mkdir -p /home/app/public_html/storage/app/images \
    && chown app:app /home/app -R

USER app
WORKDIR /home/app/public_html
CMD [ "duat-server" ]
