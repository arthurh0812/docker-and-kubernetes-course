FROM php:7.4-fpm-alpine

WORKDIR /var/www/html

COPY src .

RUN docker-php-ext-install pdo pdo_mysql

# the www-data user is the default user created inside the php:7.4-fpm-alpine image
RUN chown -R www-data:www-data /var/www/html