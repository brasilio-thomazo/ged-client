ARG UID=1000
ARG GID=1000
####################################################################################
#                                      CACHE                                       #
#                                                                                  #
####################################################################################
FROM alpine:latest as cache
RUN apk add --no-cache redis memcached bash \
    && echo "vm.overcommit_memory = 1" >> /etc/sysctl.conf
COPY resources/docker/redis-entrypoint /usr/local/bin/entrypoint
ENTRYPOINT [ "entrypoint" ]
CMD [ "redis-server", "/etc/redis.conf", "--appendonly", "yes", "--protected-mode", "no" ]

####################################################################################
#                                      PHP                                         #
#                                   COMPOSER                                       #
####################################################################################
FROM alpine:latest as composer
ARG GID
ARG UID
RUN addgroup -g ${GID} app \
    && adduser -G app -u ${UID} -s /bin/bash -D app \
    && mkdir -p /home/app/public_html \
    && chown app:app /home/app -R \
    && apk add --no-cache --no-interactive bash icu-data-full doas shadow \
    php82 php82-zip php82-xmlwriter php82-xml php82-tokenizer php82-sodium \
    php82-sockets php82-bz2 php82-session php82-phar php82-pdo_sqlite \
    php82-pdo php82-pdo_pgsql php82-pdo_mysql php82-openssl php82-opcache \
    php82-mbstring php82-intl php82-iconv php82-gettext php82-gd php82-ftp \
    php82-fileinfo php82-dom php82-curl php82-ctype php82-calendar php82-bcmath \
    php82-pecl-redis php82-pecl-imagick php82-pecl-protobuf php82-pecl-mongodb \
    php82-pecl-memcached \
    && usermod -G wheel app \
    && echo 'permit nopass :wheel as root' >> /etc/doas.d/doas.conf \
    && cp /usr/bin/php82 /usr/bin/php \
    && php -r "copy('https://getcomposer.org/installer', 'composer-setup.php');" \
    && php composer-setup.php --install-dir=/usr/local/bin --filename=composer \
    && rm -rf composer-setup.php
USER app
WORKDIR /home/app/public_html
COPY --chown=app:app . .
RUN composer install --no-dev --no-scripts --no-autoloader --prefer-dist
RUN composer install --no-dev --prefer-dist
RUN php artisan cache:clear

####################################################################################
#                                     NODEJS                                       #
#                                                                                  #
####################################################################################
FROM alpine as node
ARG GID
ARG UID
RUN addgroup -g ${UID} app \
    && adduser -h /home/app -G app -u ${GID} -D app \
    && mkdir -p /home/app/public_html \
    && apk add --no-cache nodejs npm bash curl \
    && chown app:app /home/app -R
USER app
WORKDIR /home/app/public_html
COPY --chown=app:app package.json package-lock.json tsconfig.json vite.config.ts tsconfig.node.json ./
COPY --chown=app:app resources resources
COPY --chown=app:app public public
RUN npm install && npm run build

####################################################################################
#                                       PHP                                        #
#                                       CLI                                        #
####################################################################################
FROM alpine:latest as cli
ARG GID
ARG UID

RUN addgroup -g ${GID} app \
    && adduser -G app -u ${UID} -s /bin/bash -D app \
    && mkdir -p /home/app/public_html \
    && chown app:app /home/app -R \
    && apk add --no-cache --no-interactive bash icu-data-full doas shadow \
    php82 php82-zip php82-xmlwriter php82-xml php82-tokenizer php82-sodium \
    php82-sockets php82-bz2 php82-session php82-phar php82-pdo_sqlite \
    php82-pdo php82-pdo_pgsql php82-pdo_mysql php82-openssl php82-opcache \
    php82-mbstring php82-intl php82-iconv php82-gettext php82-gd php82-ftp \
    php82-fileinfo php82-dom php82-curl php82-ctype php82-calendar php82-bcmath \
    php82-pecl-redis php82-pecl-imagick php82-pecl-protobuf php82-pecl-mongodb \
    php82-pecl-memcached \
    && echo 'permit nopass :wheel as root' >> /etc/doas.d/doas.conf \
    && usermod -G wheel app \
    && cp /usr/bin/php82 /usr/bin/php
USER app
WORKDIR /home/app/public_html
COPY --from=composer /home/app/public_html .
COPY --from=node /home/app/public_html/public public
RUN php artisan cache:clear

####################################################################################
#                                       PHP                                        #
#                                       FPM                                        #
####################################################################################
FROM alpine:latest as fpm
ARG GID
ARG UID
RUN addgroup -g ${GID} app \
    && adduser -G app -u ${UID} -s /bin/bash -D app \
    && mkdir -p /home/app/public_html \
    && chown app:app /home/app -R \
    && apk add --no-cache --no-interactive bash icu-data-full doas shadow \
    php82 php82-zip php82-xmlwriter php82-xml php82-tokenizer php82-sodium \
    php82-sockets php82-bz2 php82-session php82-phar php82-pdo_sqlite \
    php82-pdo php82-pdo_pgsql php82-pdo_mysql php82-openssl php82-opcache \
    php82-mbstring php82-intl php82-iconv php82-gettext php82-gd php82-ftp \
    php82-fileinfo php82-dom php82-curl php82-ctype php82-calendar php82-bcmath \
    php82-pecl-redis php82-pecl-imagick php82-pecl-protobuf php82-pecl-mongodb \
    php82-pecl-memcached php82-fpm curl \
    && echo 'permit nopass :wheel as root' >> /etc/doas.d/doas.conf \
    && usermod -G wheel app \
    && cp /usr/bin/php82 /usr/bin/php \
    && cp /usr/sbin/php-fpm82 /usr/bin/php-fpm

COPY resources/docker/php-fpm.ini /etc/php/fpm/
COPY resources/docker/www.ini /etc/php/fpm/pool.d/
COPY resources/docker/fpm-entrypoint /usr/local/bin/entrypoint
USER app
WORKDIR /home/app/public_html
COPY --chown=app:app --from=composer /home/app/public_html .
COPY --chown=app:app --from=node /home/app/public_html/public public
RUN php artisan config:clear\
    && php artisan cache:clear \
    && php artisan event:cache \
    && php artisan route:cache \
    && php artisan view:cache \
    && php artisan config:cache
ENTRYPOINT [ "entrypoint" ]
CMD [ "doas", "php-fpm", "-e", "-y", "/etc/php/fpm/php-fpm.ini" ]
EXPOSE 9000

####################################################################################
#                                      NGINX                                       #
#                                                                                  #
####################################################################################
FROM alpine:latest as nginx
RUN apk add --no-cache nginx gettext-envsubst bash curl \
    && mkdir -p /etc/nginx/template.d /etc/nginx/vhost.d /home/app/public_html
COPY resources/docker/default.template /etc/nginx/template.d/
COPY resources/docker/nginx.conf /etc/nginx/nginx.conf
COPY resources/docker/nginx-entrypoint /usr/local/bin/entrypoint
WORKDIR /home/app/public_html
COPY --from=node /home/app/public_html/public public
ENTRYPOINT [ "entrypoint" ]
CMD [ "nginx" ]
EXPOSE 80

####################################################################################
#                                   BUILD GRPC                                     #
#                                     SERVER                                       #
####################################################################################
FROM alpine as build
COPY resources/golang /golang
WORKDIR /golang
RUN apk add --no-cache go \
    && go mod tidy \
    && go build -v -o duat-server .

####################################################################################
#                                      GRPC                                        #
#                                     SERVER                                       #
####################################################################################
FROM alpine as grpc
ENV SERVER_PORT=50051
ARG GID
ARG UID

COPY --from=build /golang/duat-server /usr/local/bin/

RUN apk add --no-cache \
    && addgroup -g ${GID} app \
    && adduser -G app -u ${UID} -D app \
    &&  mkdir -p /home/app/public_html/storage/app \
    && chown app:app /home/app -R

USER app
WORKDIR /home/app/public_html/storage/app
CMD [ "duat-server" ]
